package router

import (
	"tcc-project/src/controller"
)

// AppendUserRouters ...
func AppendUserRouters() {

	Routers = append(Routers,
		Router{
			Domain:  "user",
			Method:  "GET",
			Path:    "/users",
			Handler: controller.ReadAllUsers,
		},
		Router{
			Domain:  "user",
			Method:  "GET",
			Path:    "/user/{id:[0-9]+}",
			Handler: controller.ReadUserByID,
		},
		Router{
			Domain:  "user",
			Method:  "POST",
			Path:    "/user",
			Handler: controller.CreateUser,
		},
		Router{
			Domain:  "user",
			Method:  "PUT",
			Path:    "/user/{id:[0-9]+}",
			Handler: controller.UpdateUser,
		},
		Router{
			Domain:  "user",
			Method:  "DELETE",
			Path:    "/user/{id:[0-9]+}",
			Handler: controller.DeleteUser,
		},
	)
}
