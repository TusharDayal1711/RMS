package handler

//not required
//func CreateUserByAdmins(w http.ResponseWriter, r *http.Request) {
//	creatorID, roles, err := middleware.GetUserAndRolesFromContext(r)
//	if err != nil {
//		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
//		return
//	}
//	isAllow := false
//
//	for _, role := range roles {
//		if role == "admin" || role == "subAdmin" {
//			isAllow = true
//			break
//		}
//	}
//	if !isAllow {
//		utils.RespondError(w, http.StatusForbidden, nil, "permission denied")
//		return
//	}
//
//	var req models.CreateUserReq
//	if err := utils.ParseJSONBody(r, &req); err != nil {
//		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
//	}
//
//	err = dbHelper.CreateUser(req, creatorID)
//	if err != nil {
//		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create user")
//		return
//	}
//
//	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
//		"message": "user created successfully",
//	})
//}
