package middleware

import (
	"context"
	"errors"
	"net/http"
	"rmssystem_1/database"
	"rmssystem_1/utils"
	"strings"
)

type contextKey string

const (
	userContextKey  contextKey = "user_key"
	rolesContextKey contextKey = "roles_key"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			utils.RespondError(w, http.StatusUnauthorized, errors.New("missing access token"), "missing access token")
			return
		}
		userID, roles, err := utils.ParseJWT(accessToken)
		if err != nil && strings.Contains(err.Error(), "invalid or expired token") {
			refreshToken := r.Header.Get("refresh_token")
			if refreshToken == "" {
				utils.RespondError(w, http.StatusUnauthorized, errors.New("missing refresh token"), "access token expired, and refresh token missing")
				return
			}
			userID, err = utils.ParseRefreshToken(refreshToken)
			if err != nil {
				utils.RespondError(w, http.StatusUnauthorized, err, "invalid or expired refresh token")
				return
			}

			// fetching roles from db using user_id
			err = database.DB.Select(&roles, `
			SELECT r.role_name FROM user_roles ur
		    JOIN roles r
			ON ur.role_id = r.id
 			WHERE ur.user_id = $1
			`, userID)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch roles")
				return
			}
			newToken, err := utils.GenerateJWT(userID, roles)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate new access token")
				return
			}
			//regenerating refresh token
			newRefreshToken, err := utils.GenerateRefreshToken(userID)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate new refresh token")
			}
			w.Header().Set("Authorization", newToken)
			w.Header().Set("Refresh_token", newRefreshToken)

		} else if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}
		ctx := context.WithValue(r.Context(), userContextKey, userID)
		ctx = context.WithValue(ctx, rolesContextKey, roles)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireRole(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, roles, err := GetUserAndRolesFromContext(r)
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			for _, role := range roles {
				if role == requiredRole {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "unauthorized user", http.StatusForbidden)
		})
	}
}

func GetUserAndRolesFromContext(r *http.Request) (string, []string, error) {
	userID, ok := r.Context().Value(userContextKey).(string)
	if !ok {
		return "", nil, errors.New("user ID not found in context")
	}
	roles, ok := r.Context().Value(rolesContextKey).([]string)
	if !ok {
		return "", nil, errors.New("roles not found in context")
	}
	return userID, roles, nil
}
