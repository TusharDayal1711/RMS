package routes

import (
	"rmssystem_1/handler"
	"rmssystem_1/middleware"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/public-register", handler.RegisterPublicUser).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/GetAllrestaurant", handler.ListAllRestaurants).Methods("GET")
	router.HandleFunc("/GetAllDishes", handler.GetAllDishesHandler).Methods("GET")

	// Protected base route
	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.JWTAuthMiddleware)
	protectedRoutes.HandleFunc("/user/set-address", handler.SetAddressHandler).Methods("POST")
	protectedRoutes.HandleFunc("/user/distance", handler.CalculateDistance).Methods("GET")
	protectedRoutes.HandleFunc("/GetAllrestaurant", handler.ListAllRestaurants).Methods("GET")
	protectedRoutes.HandleFunc("/GetAllDishes", handler.GetAllDishesHandler).Methods("GET")
	protectedRoutes.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")

	// Shared routes: admin or subAdmin
	sharedRoutes := protectedRoutes.PathPrefix("/common").Subrouter()
	sharedRoutes.Use(middleware.RequireRole("admin", "subAdmin"))
	sharedRoutes.HandleFunc("/create-restaurant", handler.CreateRestaurantHandler).Methods("POST")
	sharedRoutes.HandleFunc("/create-dish", handler.AddDish).Methods("POST")
	sharedRoutes.HandleFunc("/dishes-by-rest-id", handler.GetDishesByRestaurant).Methods("GET")
	sharedRoutes.HandleFunc("/restaurants-by-creator", handler.GetMyRestaurantsByCreatorId).Methods("GET")
	sharedRoutes.HandleFunc("/dishes-by-creator", handler.GetMyDishesHandler).Methods("GET")
	sharedRoutes.HandleFunc("/users-by-creator", handler.GetUsersCreatedById).Methods("GET")
	sharedRoutes.HandleFunc("/create-user", handler.CreateUserWithRolesByAdmins).Methods("POST")

	// SubAdmin routes
	subAdminRoutes := protectedRoutes.PathPrefix("/subadmin").Subrouter()
	subAdminRoutes.Use(middleware.RequireRole("subAdmin"))

	// Admin routes
	adminRoutes := protectedRoutes.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(middleware.RequireRole("admin"))
	adminRoutes.HandleFunc("/create-user-with-roles", handler.CreateUserWithRolesByAdmins).Methods("POST")

	return router
}
