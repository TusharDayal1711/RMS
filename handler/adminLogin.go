package handler

import (
	"net/http"
	"rmssystem_1/middleware"
	"rmssystem_1/utils"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	adminId, role, err := middleware.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

}
