package routers

import (
	"tcc-project/src/controllers"
)

func init() {
	Routers = append(Routers,
		Router{
			Method:  "GET",
			Path:    "/users",
			Handler: controllers.ReadAllUsers,
		},
		Router{
			Method:  "GET",
			Path:    "/user/{id:[0-9]+}",
			Handler: controllers.ReadUserByID,
		},
		Router{
			Method:  "POST",
			Path:    "/user",
			Handler: controllers.CreateUser,
		},
		Router{
			Method:  "PUT",
			Path:    "/user/{id:[0-9]+}",
			Handler: controllers.UpdateUser,
		},
		Router{
			Method:  "DELETE",
			Path:    "/user/{id:[0-9]+}",
			Handler: controllers.DeleteUser,
		},
	)
}
