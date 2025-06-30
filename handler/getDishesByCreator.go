package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/utils"
)

func GetMyDishesHandler(w http.ResponseWriter, r *http.Request) {
	userID, _, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	dishes, err := dbHelper.GetDishesByCreator(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsoniter.NewEncoder(w).Encode(map[string]interface{}{
			"message": "failed to fetch dishes",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "dishes fetched successfully",
		"data":    dishes,
	})
}
