package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/models"
	"rmssystem_1/utils"
)

func SetAddressHandler(w http.ResponseWriter, r *http.Request) {
	userID, _, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	var req models.AddressReq
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
		return
	}
	if req.Address == "" || req.Longitude < -180 || req.Longitude > 180 || req.Latitude < -90 || req.Latitude > 90 {
		utils.RespondError(w, http.StatusBadRequest, nil, "address along with coordinates required")
		return
	}
	err = dbHelper.SetUserAddress(req, userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to save address")
		return
	}
	w.WriteHeader(http.StatusCreated)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Address Set Successfully",
	})
}
