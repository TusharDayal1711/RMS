package handler

import (
	"encoding/json"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/models"
	"rmssystem_1/utils"
)

func CreateUserByAdmins(w http.ResponseWriter, r *http.Request) {
	creatorID, roles, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	isAllow := false

	for _, role := range roles {
		if role == "admin" || role == "subAdmin" {
			isAllow = true
			break
		}
	}
	if !isAllow {
		utils.RespondError(w, http.StatusForbidden, nil, "permission denied")
		return
	}

	var req models.CreateUserReq
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
	}

	err = dbHelper.CreateUser(req, creatorID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create user")
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "user created successfully",
	})
}
