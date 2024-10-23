package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tcc-project/src/router"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	port := os.Getenv("SERVER_PORT")

	router.InitilizarRouters()

	fmt.Printf("Up and running in port %v\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(err)
	}

}
