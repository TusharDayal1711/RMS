package routes

import (
	"rmssystem_1/handler"
	"rmssystem_1/middleware"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	mainRouter := mux.NewRouter()
	api := mainRouter.PathPrefix("/api").Subrouter()

	//public routes
	api.HandleFunc("/register", handler.RegisterPublicUser).Methods("POST")
	api.HandleFunc("/user/login", handler.Login).Methods("POST")
	api.HandleFunc("/subadmin/login", handler.LoginSubAdmins).Methods("POST")
	api.HandleFunc("/admin/login", handler.LoginAdmin).Methods("POST")
	api.HandleFunc("/restaurants", handler.ListAllRestaurants).Methods("GET")
	api.HandleFunc("/dishes", handler.GetAllDishesHandler).Methods("GET")
	api.HandleFunc("/restaurant", handler.GetRestaurantById).Methods("GET")
	api.HandleFunc("/restaurant/dishes", handler.GetDishesByRestaurant).Methods("GET")

	//protected routes
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
	sharedRoutes.HandleFunc("/user/restaurant", handler.GetMyRestaurantsByAdminId).Methods("GET")
	sharedRoutes.HandleFunc("/user/dishes", handler.GetMyDishesByAdminId).Methods("GET")
	sharedRoutes.HandleFunc("/user/users", handler.GetUsersCreatedById).Methods("GET")
	sharedRoutes.HandleFunc("/user/register", handler.CreateUserWithRolesByAdmins).Methods("POST")

	//only admim
	adminRoutes := protectedRoutes.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(middleware.RequireRole("admin"))
	adminRoutes.HandleFunc("/subadmins", handler.GetAllSubAdmins).Methods("GET")

	return mainRouter
}
