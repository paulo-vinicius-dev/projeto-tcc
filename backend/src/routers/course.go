package routers

import "tcc-project/src/controllers"

func init() {
	Routers = append(Routers,
		Router{
			Method:       "GET",
			Path:         "/courses",
			Handler:      controllers.ReadAllCourses,
			AuthRequired: true,
		},
		Router{
			Method:       "GET",
			Path:         "/course/{id:[0-9]+}",
			Handler:      controllers.ReadCourseByID,
			AuthRequired: true,
		},
	)
}
