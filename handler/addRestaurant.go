package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/models"
	"rmssystem_1/utils"
)

func CreateRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	userID, role, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	var req models.RestaurantReq
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
	}
	err = dbHelper.CreateNewRestaurant(req, userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create restaurant")
		return
	}
	w.WriteHeader(http.StatusCreated)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "restaurant created successfully",
		"created by": role,
	})
}
