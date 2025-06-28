package handler

import (
	"github.com/google/uuid"
	"log"
	"rmssystem_1/database"
	"rmssystem_1/utils"
)

func CreateSuperAdmin() {
	adminEmail := "admin@gmail.com"
	adminPassword := "admin123"
	adminName := "Super Admin"

	hashedPassword, err := utils.HashPassword(adminPassword)
	if err != nil {
		log.Println("failed to hash password:", err)
		return
	}

	var adminID uuid.UUID
	err = database.DB.Get(&adminID, `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`, adminName, adminEmail, hashedPassword)
	if err != nil {
		log.Println("admin already exists or error inserting user:", err)
		return
	}

	var roleID uuid.UUID
	err = database.DB.Get(&roleID, `SELECT id FROM roles WHERE role_name = 'admin'`)
	if err != nil {
		err = database.DB.Get(&roleID, `
			INSERT INTO roles (role_name)
			VALUES ('admin')
			RETURNING id
		`)
		if err != nil {
			log.Println("failed to create admin role:", err)
			return
		}
	}

	_, err = database.DB.Exec(`
		INSERT INTO user_roles (user_id, role_id, created_by)
		VALUES ($1, $2, $3)
	`, adminID, roleID, adminID)
	if err != nil {
		log.Println("failed to assign admin role:", err)
		return
	}

	log.Println("    Admin user created:")
	log.Println("    Email: admin@example.com")
	log.Println("    Password:", adminPassword)
}
