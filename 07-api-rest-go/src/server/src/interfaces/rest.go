package interfaces

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	msgjson "json"
	"usecases"
)

type TopicInteractor interface {
	Add(text string) (usecases.Topic, error)
	Like(id int64) (usecases.Topic, error)
}

type TopicsInteractor interface {
	List() usecases.Topics
}

type RestHandler struct {
	TopicInteractor  TopicInteractor
	TopicsInteractor TopicsInteractor
}

func (handler RestHandler) Add() httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		type msg struct {
			Text string `json:"text"`
		}
		var m msg
		j, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Panic(err)
		}
		err = json.Unmarshal(j, &m)
		if err != nil {
			log.Panic(err)
		}
		topic, err := handler.TopicInteractor.Add(m.Text)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			j = msgjson.KO("Bad request, topic not added")
			rw.Write(j)
			return
		}
		j = msgjson.OK(topic.ID, "New topic added")
		rw.Header().Set("Location", "/topics/id/"+strconv.FormatInt(topic.ID, 10))
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(j)
	}
}

func (handler RestHandler) Like() httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var j []byte
		id, err := strconv.ParseInt(p.ByName("id"), 0, 64)
		if err != nil {
			log.Panic(err)
		}
		topic, err := handler.TopicInteractor.Like(id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			j = msgjson.KO("Bad request, topic not liked")
			rw.Write(j)
			return
		}
		j = msgjson.OK(topic.ID, "Topic liked")
		rw.Header().Set("Location", "/topics/id/"+strconv.FormatInt(topic.ID, 10))
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(j)
	}
}

func (handler RestHandler) List() httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		topics := handler.TopicsInteractor.List()
		j, err := json.Marshal(topics.Collection)
		if err != nil {
			log.Panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(j)
	}
}
