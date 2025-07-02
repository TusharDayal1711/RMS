package utils

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ParseJSONBody(r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

// HashPassword
func CheckHashPassword(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		log.Println("Password mismatch:", err)
	}
	return err == nil
}

func GetPageLimitAndOffset(r *http.Request) (int, int) {
	page := 1
	limit := 10
	if pageValue := r.URL.Query().Get("page"); pageValue != "" {
		if p, err := strconv.Atoi(pageValue); err == nil {
			page = p //if error ocurs keep the default valuse
		}
	}
	if limitValue := r.URL.Query().Get("limit"); limitValue != "" {
		if l, err := strconv.Atoi(limitValue); err == nil {
			limit = l //if error ocurs keep the default valuse
		}
	}
	offset := (page - 1) * limit
	return limit, offset
}
