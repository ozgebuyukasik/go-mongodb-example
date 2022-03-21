package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	router := httprouter.New()
	userController := controllers.NewUserController(getSession())
	router.GET("/user/:id", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	http.ListenAndServe("localhost:9000", router)
}

func getSession() *mgo.Session {
	session, error := mgo.Dial("mongodb://localhost:27107")

	if error != nil {
		panic(error)
	}
	return session
}
