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

	// Protected routes
	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.JWTAuthMiddleware)
	protectedRoutes.HandleFunc("/user/set-address", handler.SetAddressHandler).Methods("POST")
	protectedRoutes.HandleFunc("/user/distance", handler.CalculateDistance).Methods("GET")
	protectedRoutes.HandleFunc("/GetAllrestaurant", handler.ListAllRestaurants).Methods("GET")
	protectedRoutes.HandleFunc("/GetAllDishes", handler.GetAllDishesHandler).Methods("GET")
	protectedRoutes.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")

	// SubAdmin routes
	subAdminRoutes := protectedRoutes.PathPrefix("/subadmin").Subrouter()
	subAdminRoutes.Use(middleware.RequireRole("subAdmin"))
	subAdminRoutes.HandleFunc("/create-restaurant", handler.CreateRestaurantHandler).Methods("POST")
	subAdminRoutes.HandleFunc("/create-user", handler.CreateUserByAdmins).Methods("POST")
	subAdminRoutes.HandleFunc("/create-dish", handler.AddDish).Methods("POST")
	subAdminRoutes.HandleFunc("/dishes-by-rest-id", handler.GetDishesByRestaurant).Methods("GET")
	subAdminRoutes.HandleFunc("/restaurants-by-creator", handler.GetMyRestaurantsByCreatorId).Methods("GET")
	subAdminRoutes.HandleFunc("/dishes-by-creator", handler.GetMyDishesHandler).Methods("GET")
	subAdminRoutes.HandleFunc("/users-by-creator", handler.GetUsersCreatedById).Methods("GET")

	// Admin routes
	adminRoutes := protectedRoutes.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(middleware.RequireRole("admin"))
	adminRoutes.HandleFunc("/create-subadmin", handler.CreateSubAdmins).Methods("POST")
	adminRoutes.HandleFunc("/create-restaurant", handler.CreateRestaurantHandler).Methods("POST")
	adminRoutes.HandleFunc("/create-dish", handler.AddDish).Methods("POST")
	adminRoutes.HandleFunc("/create-user", handler.CreateUserByAdmins).Methods("POST")
	adminRoutes.HandleFunc("/dishes-by-creator", handler.GetMyDishesHandler).Methods("GET")
	adminRoutes.HandleFunc("/dishes-by-rest-id", handler.GetDishesByRestaurant).Methods("GET")
	adminRoutes.HandleFunc("/users-by-creator", handler.GetUsersCreatedById).Methods("GET")
	adminRoutes.HandleFunc("/restaurantsByCreator", handler.GetMyRestaurantsByCreatorId).Methods("GET")
	adminRoutes.HandleFunc("/create-user-with-roles", handler.CreateUserWithRoles).Methods("POST")

	return router
}
