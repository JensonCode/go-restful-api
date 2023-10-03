package main

import (
	"fmt"
	"net/http"

	"github.com/JensonCode/go-restful-api/internal/auth"
	"github.com/JensonCode/go-restful-api/internal/database"
)

func main() {

    //connect to planet scale database
    db := database.ConnectPlanetScale()
    defer db.Close()

    auth.AuthRoutes()

    //listen to port 8080
    port := ":8080"

    fmt.Printf("Server listening on port %s\n", port)

    err:= http.ListenAndServe(port, nil)

    if err != nil{
        panic(err)
    }

}