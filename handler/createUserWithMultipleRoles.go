package handler

import (
	"encoding/json"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/models"
	"rmssystem_1/utils"
)

func CreateUserWithRoles(w http.ResponseWriter, r *http.Request) {
	var MultiRolereq models.MultiRole
	if err := json.NewDecoder(r.Body).Decode(&MultiRolereq); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
		return
	}

	CreatorID, role, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	err = dbHelper.CreateUserWithMultiRole(MultiRolereq, CreatorID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create user")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Password ":      MultiRolereq.Password,
		"User Email Id ": MultiRolereq.Email,
		"created by":     role,
		"message":        "successfully created user",
	})
}
