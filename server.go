package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"
	"os"

	//use full path for import statement, otherwise import statements in subdirectories fail for some reason
	"github.com/QianMason/drone-cloud-tracking/routes"

	//"github.com/QianMason/drone-backend/utils"
	_ "github.com/unrolled/render"
	_ "github.com/urfave/negroni"
)

func main() {
	port, success := os.LookupEnv("PORT")
	fmt.Println("port lookup")
	if !success {
		port = "8080"
	}
	//models.Init()
	//utils.LoadTemplates("templates/*html")
	fmt.Println("handling routes")
	r := routes.NewRouter()
	http.Handle("/", r)
	fmt.Println("listening on port:", port)
	http.ListenAndServe(":"+port, nil)
}
