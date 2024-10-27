package routers

import "tcc-project/src/controllers"

func init() {
	Routers = append(Routers,
		Router{
			Method:       "POST",
			Path:         "/login",
			Handler:      controllers.Login,
			AuthRequired: false,
		},
		Router{
			Method:       "POST",
			Path:         "/logout",
			Handler:      controllers.Logout,
			AuthRequired: true,
		},
	)
}
