package json

import (
	"encoding/json"
	"log"
)

type t struct {
	ID      int64
	Message string
}

func OK(id int64, msg string) []byte {
	r := t{}
	r.ID = id
	r.Message = msg
	j, err := json.Marshal(r)
	if err != nil {
		log.Panic(err)
	}
	return j
}

func KO(msg string) []byte {
	r := t{}
	r.Message = msg
	j, err := json.Marshal(r)
	if err != nil {
		log.Panic(err)
	}
	return j
}
