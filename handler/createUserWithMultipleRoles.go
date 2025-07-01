package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/models"
	"rmssystem_1/utils"
)

func CreateUserWithRoles(w http.ResponseWriter, r *http.Request) {
	var MultiRolereq models.MultiRole

	if err := utils.ParseJSONBody(r, &MultiRolereq); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
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
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"Password ":      MultiRolereq.Password,
		"User Email Id ": MultiRolereq.Email,
		"created by":     role,
		"message":        "successfully created user",
	})
}

func CreateUserWithRolesByAdmins(w http.ResponseWriter, r *http.Request) {
	creatorId, roles, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	var req models.MultiRole
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
	}
	if len(req.Roles) == 0 {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
		return
	}
	isAdmin := false
	for _, role := range roles {
		if role == "admin" {
			isAdmin = true
			break
		}
	}
	if !isAdmin {
		for _, role := range req.Roles {
			if role != "user" {
				utils.RespondError(w, http.StatusBadRequest, err, "not authorized to create subAdmin with this permisssion")
				return
			}
		}
	}
	err = dbHelper.CreateUserWithMultiRole(req, creatorId)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create user")
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"Password ":      req.Password,
		"User Email Id ": req.Email,
		"created by":     roles,
		"message":        "successfully created user",
	})
}
