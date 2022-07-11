package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Routes{
	{
		Uri:     "/users", // Create User
		Methods: http.MethodPost,
		Func:    controllers.CreatedUser,
		Auth:    false,
	},
	{
		Uri:     "/users", // Get All User
		Methods: http.MethodGet,
		Func:    controllers.GetAllUsers,
		Auth:    false,
	},
	{
		Uri:     "/users/{id}", // Get user
		Methods: http.MethodGet,
		Func:    controllers.GetUser,
		Auth:    false,
	},
	{
		Uri:     "/users/{id}", // Update
		Methods: http.MethodPut,
		Func:    controllers.UpdateUser,
		Auth:    false,
	},
	{
		Uri:     "/users/{id}", // Delete
		Methods: http.MethodDelete,
		Func:    controllers.DeleteUser,
		Auth:    false,
	},
}
