package routes

import (
	"rmssystem_1/handler"
	"rmssystem_1/middleware"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	server := mux.NewRouter()
	router := server.PathPrefix("/api/v1").Subrouter()
	// Public routes
	router.HandleFunc("/register", handler.RegisterPublicUser).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/restaurant", handler.ListAllRestaurants).Methods("GET")
	router.HandleFunc("/dishes", handler.GetAllDishesHandler).Methods("GET")

	// Protected base route
	protectedRoutes := router.NewRoute().Subrouter()
	protectedRoutes.Use(middleware.JWTAuthMiddleware)
	protectedRoutes.HandleFunc("/user/address", handler.SetAddressHandler).Methods("POST")
	protectedRoutes.HandleFunc("/user/distance", handler.CalculateDistance).Methods("GET")
	//protectedRoutes.HandleFunc("/GetAllrestaurant", handler.ListAllRestaurants).Methods("GET")
	//protectedRoutes.HandleFunc("/GetAllDishes", handler.GetAllDishesHandler).Methods("GET")
	protectedRoutes.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")

	// Shared routes, admin or subAdmin
	sharedRoutes := protectedRoutes.PathPrefix("/common").Subrouter()
	sharedRoutes.Use(middleware.RequireRole("admin", "subAdmin"))
	sharedRoutes.HandleFunc("/restaurant", handler.CreateRestaurantHandler).Methods("POST")
	sharedRoutes.HandleFunc("/dish", handler.AddDish).Methods("POST")
	sharedRoutes.HandleFunc("/restaurant/{restaurant}/dishes", handler.GetDishesByRestaurant).Methods("GET")
	sharedRoutes.HandleFunc("/user/restaurant", handler.GetMyRestaurantsByCreatorId).Methods("GET")
	sharedRoutes.HandleFunc("/user/dishes", handler.GetMyDishesHandler).Methods("GET")
	sharedRoutes.HandleFunc("/users-by-creator", handler.GetUsersCreatedById).Methods("GET")
	sharedRoutes.HandleFunc("/create-user", handler.CreateUserWithRolesByAdmins).Methods("POST")

	// Admin routes
	adminRoutes := protectedRoutes.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(middleware.RequireRole("admin"))
	adminRoutes.HandleFunc("/get-subAdmins", handler.GetAllSubAdmins).Methods("GET")

	return router
}
