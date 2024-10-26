package routers

import "tcc-project/src/controllers"

func init() {
	Routers = append(Routers,
		Router{
			Method:  "GET",
			Path:    "/units",
			Handler: controllers.ReadAllUnits,
		},
		Router{
			Method:  "GET",
			Path:    "/unit/{id:[0-9]+}",
			Handler: controllers.ReadUnitByID,
		},
	)
}
