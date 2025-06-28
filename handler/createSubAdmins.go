package handler

import (
	"encoding/json"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/models"
	"rmssystem_1/utils"
)

func CreateSubAdmins(w http.ResponseWriter, r *http.Request) {
	adminID, role, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	var req models.SubAdminReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
		return
	}

	err = dbHelper.CreateNewSubAdmin(req, adminID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "internal server error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "subAdmin created successfully",
		"crated by": role,
	})
}
