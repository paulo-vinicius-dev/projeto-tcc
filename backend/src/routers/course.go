package routers

import "tcc-project/src/controllers"

// AppendCourseRouters ...
func AppendCourseRouters() {
	Routers = append(Routers,
		Router{
			Method:  "GET",
			Path:    "/courses",
			Handler: controllers.ReadAllCourses,
		},
	)
}
