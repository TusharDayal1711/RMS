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

	tx, err := database.DB.Beginx()
	if err != nil {
		log.Println("failed to start transaction:", err)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Println("transaction panicked and rolled back")
			panic(p)
		} else if err != nil {
			tx.Rollback()
			log.Println("transaction rolled back due to error:", err)
		} else {
			err = tx.Commit()
			if err != nil {
				log.Println("failed to commit transaction:", err)
			}
		}
	}()

	var adminID uuid.UUID
	err = tx.Get(&adminID, `
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
		err = tx.Get(&roleID, `
			INSERT INTO roles (role_name)
			VALUES ('admin')
			RETURNING id
		`)
		if err != nil {
			log.Println("failed to create admin role:", err)
			return
		}
	}

	_, err = tx.Exec(`
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
