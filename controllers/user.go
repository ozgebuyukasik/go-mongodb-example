package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/ozgebuyukasik/mongo-golang/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

func (userController UserController) GetUser(responseWriter http.ResponseWriter, req *http.Request, params httprouter.Params) {

	userID := params.ByName("id")

	if !bson.IsObjectIdHex(userID) {
		responseWriter.WriteHeader(http.StatusNotFound)
	}

	objectId := bson.ObjectIdHex(userID)

	user := models.User{}

	if error := userController.session.DB("mongo-golang").C("users").FindId(objectId).One(&user); error != nil {
		responseWriter.WriteHeader(404)
		return
	}

	userJsonData, error := json.Marshal(user)

	if error != nil {
		fmt.Println(error)
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(responseWriter, "/s\n", userJsonData)
}

func (userController UserController) CreateUser(responseWriter http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	newUser := models.User{}
	json.NewDecoder(req.Body).Decode(&newUser)
	newUser.Id = bson.NewObjectId()
	userController.session.DB("mongo-golang").C("users").Insert(newUser)
	userJsonData, error := json.Marshal(newUser)

	if error != nil {
		fmt.Println(error)
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(responseWriter, "%s\n", userJsonData)
}

func (userController UserController) DeleteUser(responseWriter http.ResponseWriter, req *http.Request, params httprouter.Params) {

	userID := params.ByName("id")

	if !bson.IsObjectIdHex(userID) {
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}

	objectId := bson.ObjectIdHex(userID)

	if err := userController.session.DB("mongo-golang").C("users").RemoveId(objectId); err != nil {
		responseWriter.WriteHeader(404)
	}

	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(responseWriter, "Deleted user with id %s\n", objectId)

}
