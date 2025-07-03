package handler

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"rmssystem_1/database/dbHelper"
	"rmssystem_1/middleware"
	"rmssystem_1/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	limit, offset := utils.GetPageLimitAndOffset(r)
	_, _, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	limit, offset = utils.GetPageLimitAndOffset(r)
	users, err := dbHelper.GetAllUsers(limit, offset)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "List of all users created :-",
		"data":    users,
	})
}
