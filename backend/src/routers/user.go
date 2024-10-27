package routers

import (
	"tcc-project/src/controllers"
)

func init() {
	Routers = append(Routers,
		Router{
			Method:       "GET",
			Path:         "/users",
			Handler:      controllers.ReadAllUsers,
			AuthRequired: true,
		},
		Router{
			Method:       "GET",
			Path:         "/user/{id:[0-9]+}",
			Handler:      controllers.ReadUserByID,
			AuthRequired: true,
		},
		Router{
			Method:       "POST",
			Path:         "/user",
			Handler:      controllers.CreateUser,
			AuthRequired: true,
		},
		Router{
			Method:       "PUT",
			Path:         "/user/{id:[0-9]+}",
			Handler:      controllers.UpdateUser,
			AuthRequired: true,
		},
		Router{
			Method:       "DELETE",
			Path:         "/user/{id:[0-9]+}",
			Handler:      controllers.DeleteUser,
			AuthRequired: true,
		},
	)
}
