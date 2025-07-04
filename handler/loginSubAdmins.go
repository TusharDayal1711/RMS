package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/models"
	"rmssystem_1/utils"
)

func LoginSubAdmins(w http.ResponseWriter, r *http.Request) {
	var req models.LoginReq
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
		return
	}
	if req.Email == "" || req.Password == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "email and password are required")
		return
	}
	subAdminId, hashedPassword, err := dbHelper.GetUserByEmail(req.Email)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	if !utils.CheckHashPassword(req.Password, hashedPassword) {
		utils.RespondError(w, http.StatusUnauthorized, nil, "unauthorized")
		return
	}
	roles, err := dbHelper.GetUserRoles(subAdminId)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to get user roles")
		return
	}
	isSubAdmin := false
	for _, role := range roles {
		if role == "subAdmin" {
			isSubAdmin = true
		}
	}
	if !isSubAdmin {
		utils.RespondError(w, http.StatusUnauthorized, nil, "unauthorized")
		return
	}
	accessToken, err := middleware.GenerateJWT(subAdminId, roles)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate access token")
		return
	}

	refreshToken, err := middleware.GenerateRefreshToken(subAdminId)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate refresh token")
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "subAdmin log-in successfully ",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
