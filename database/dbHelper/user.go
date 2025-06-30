package dbHelper

import (
	"database/sql"
	"fmt"
	"rmssystem_1/utils"
	"time"

	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"rmssystem_1/database"
	"rmssystem_1/models"
)

func CreatePublicUser(user models.User) (uuid.UUID, error) {
	user.Email = strings.ToLower(user.Email)

	tx, err := database.DB.Begin()
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Checking if user already exists
	var isExist uuid.UUID
	err = tx.QueryRow(`
		SELECT id FROM users 
		WHERE email = $1 AND archived_at IS NULL
	`, user.Email).Scan(&isExist)
	if err == nil {
		return uuid.Nil, fmt.Errorf("email already exists")
	} else if !errors.Is(err, sql.ErrNoRows) {
		return uuid.Nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	user.Password = strings.TrimSpace(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to hash password")
	}

	var newUserID uuid.UUID
	err = tx.QueryRow(`
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`, user.Name, user.Email, hashedPassword).Scan(&newUserID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %w", err)
	}

	const defaultRole = "user"
	var roleID uuid.UUID
	err = tx.QueryRow(`
		SELECT id FROM roles 
		WHERE role_name = $1
	`, defaultRole).Scan(&roleID)
	if err != nil {
		fmt.Println("failed to find role 'user' ", err)
		return uuid.Nil, fmt.Errorf("default user role not found: %w", err)
	}

	// Assign the role to the user
	_, err = tx.Exec(`
		INSERT INTO user_roles (user_id, role_id, created_by)
		VALUES ($1, $2, $1)
	`, newUserID, roleID)
	if err != nil {
		fmt.Println("failed to assign role to user:", err)
		return uuid.Nil, fmt.Errorf("failed to assign role to user: %w", err)
	}
	return newUserID, nil
}

func CreateUser(req models.CreateUserReq, createdBy string) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {

		return err
	}
	tx, err := database.DB.Begin()
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	createdByUUID, err := uuid.Parse(createdBy)
	if err != nil {
		return err
	}
	var userID uuid.UUID
	err = tx.QueryRow(`
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`, req.Name, req.Email, hashedPassword).Scan(&userID)
	if err != nil {
		return err
	}

	var roleID uuid.UUID
	err = tx.QueryRow(`
		SELECT id FROM roles 
		WHERE role_name = 'user'
	`).Scan(&roleID)
	if err != nil {
		return err
	}

	// Assign "user" role
	_, err = tx.Exec(`
		INSERT INTO user_roles (user_id, role_id, created_by)
		VALUES ($1, $2, $3)
	`, userID, roleID, createdByUUID)

	return err
}

func GetUserByEmail(email string) (string, []byte, error) {
	var userID string
	var hashedPassword []byte

	err := database.DB.QueryRow(`
		SELECT id, password FROM users
		WHERE email = $1 AND archived_at IS NULL
	`, email).Scan(&userID, &hashedPassword)

	if err != nil {
		return "", nil, err
	}
	return userID, hashedPassword, nil
}

func GetUserRoles(userID string) ([]string, error) {
	var roles []string
	err := database.DB.Select(&roles, `
		SELECT r.role_name FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = $1
	`, userID)
	return roles, err
}

func SaveSession(userID, refreshToken string) error {
	_, err := database.DB.Exec(`
		INSERT INTO sessions (user_id, refresh_token, expire_at)
		VALUES ($1, $2, $3)
	`, userID, refreshToken, time.Now().Add(7*24*time.Hour))
	return err
}
