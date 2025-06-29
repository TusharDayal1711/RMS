package handler

import (
	"encoding/json"
	"net/http"
	"rmssystem_1/database/dbHelper"
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

	userID, err := dbHelper.CreatePublicUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			utils.RespondError(w, http.StatusConflict, err, "Email is already registered")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, err, "Failed to create user")
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"user_id": userID,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginReq
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
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
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch user roles")
		return
	}

	accessToken, err := utils.GenerateJWT(userID, roles)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate access token")
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate refresh token")
		return
	}

	err = dbHelper.SaveSession(userID, refreshToken)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to save session")
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "User logged successfully",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
