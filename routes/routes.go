package routes

import (
	"rmssystem_1/handler"
	"rmssystem_1/middleware"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	mainRouter := mux.NewRouter()
	api := mainRouter.PathPrefix("/api/v1").Subrouter()

	// --- Public Routes ---
	api.HandleFunc("/register", handler.RegisterPublicUser).Methods("POST")
	api.HandleFunc("/login", handler.Login).Methods("POST")
	api.HandleFunc("/restaurant", handler.ListAllRestaurants).Methods("GET")
	api.HandleFunc("/dishes", handler.GetAllDishesHandler).Methods("GET")

	//
	protectedRoutes := api.NewRoute().Subrouter()
	protectedRoutes.Use(middleware.JWTAuthMiddleware)

	protectedRoutes.HandleFunc("/user/address", handler.SetAddressHandler).Methods("POST")
	protectedRoutes.HandleFunc("/user/distance", handler.CalculateDistance).Methods("GET")
	protectedRoutes.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")

	//shared route admin and subadmin
	sharedRoutes := protectedRoutes.PathPrefix("/common").Subrouter()
	sharedRoutes.Use(middleware.RequireRole("admin", "subAdmin"))

	sharedRoutes.HandleFunc("/restaurant", handler.CreateRestaurantHandler).Methods("POST")
	sharedRoutes.HandleFunc("/dish", handler.AddDish).Methods("POST")
	sharedRoutes.HandleFunc("/restaurant/dishes", handler.GetDishesByRestaurant).Methods("GET")
	sharedRoutes.HandleFunc("/user/restaurant", handler.GetMyRestaurantsByCreatorId).Methods("GET")
	sharedRoutes.HandleFunc("/user/dishes", handler.GetMyDishesHandler).Methods("GET")
	sharedRoutes.HandleFunc("/users-by-creator", handler.GetUsersCreatedById).Methods("GET")
	sharedRoutes.HandleFunc("/create-user", handler.CreateUserWithRolesByAdmins).Methods("POST")

	//only admim
	adminRoutes := protectedRoutes.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(middleware.RequireRole("admin"))
	adminRoutes.HandleFunc("/sub-admins", handler.GetAllSubAdmins).Methods("GET")

	return mainRouter
}
