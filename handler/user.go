package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/models"
	"rmssystem_1/utils"
	"strings"
)

// register new public user
func RegisterPublicUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := utils.ParseJSONBody(r, &user); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
	}
	if user.Email == "" || user.Password == "" || user.Name == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "username, email and password are required")
		return
	}

	userID, err := dbHelper.CreatePublicUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			utils.RespondError(w, http.StatusConflict, err, "email is already registered")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, err, "Failed to create user")
		}
		return
	}
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"user_id": userID,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginReq
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
		return
	}
	if req.Email == "" || req.Password == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "email and password are required")
		return
	}
	userID, hashedPassword, err := dbHelper.GetUserByEmail(req.Email)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "invalid email or password")
		return
	}
	if !utils.CheckHashPassword(req.Password, hashedPassword) {
		utils.RespondError(w, http.StatusUnauthorized, nil, "invalid email or password")
		return
	}

	roles, err := dbHelper.GetUserRoles(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to get user roles")
		return
	}

	accessToken, err := middleware.GenerateJWT(userID, roles)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate access token")
		return
	}

	refreshToken, err := middleware.GenerateRefreshToken(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate refresh token")
		return
	}

	// no need for session
	//err = dbHelper.SaveSession(userID, refreshToken)
	//if err != nil {
	//	utils.RespondError(w, http.StatusInternalServerError, err, "failed to save session")
	//	return
	//}

	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "User logged-in successfully",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	_, _, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	//no need or session table records
	//err = dbHelper.DeleteSessionRecord(userID)
	//if err != nil {
	//	http.Error(w, "failed to delete records ", http.StatusInternalServerError)
	//	return
	//}
	jsoniter.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}
