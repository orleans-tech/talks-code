package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"infrastructure"
	"interfaces"
	"usecases"
)

func main() {
	dbHandler := infrastructure.NewSqliteHandler("./topics.db")
	topicInteractor := new(usecases.TopicInteractor)
	topicInteractor.TopicRepository = interfaces.NewDbTopicRepo(dbHandler)
	topicsInteractor := new(usecases.TopicsInteractor)
	topicsInteractor.TopicsRepository = interfaces.NewDbTopicsRepo(dbHandler)

	restHandler := interfaces.RestHandler{}
	restHandler.TopicInteractor = topicInteractor
	restHandler.TopicsInteractor = topicsInteractor

	router := httprouter.New()
	router.PanicHandler = func(rw http.ResponseWriter, r *http.Request, p interface{}) {
		log.Println(fmt.Sprintf("FATAL ERROR: rw:%v r:%v p:%v", rw, r, p))
	}
	router.GET("/topics", restHandler.List())
	router.POST("/topics", restHandler.Add())
	router.PUT("/topics/like/:id", restHandler.Like())
	log.Fatal(http.ListenAndServe("localhost:49200", router))
}
