package interfaces

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"

	"infrastructure"
	"usecases"
)

func TestCreate3ThenLikeId2(test *testing.T) {
	var err error
	textLike := "Topic 2 - To Like"
	dbHandler := infrastructure.NewSqliteHandler("./../../tests.db")
	topicInteractor := new(usecases.TopicInteractor)
	topicInteractor.TopicRepository = NewDbTopicRepo(dbHandler)
	topicsInteractor := new(usecases.TopicsInteractor)
	topicsInteractor.TopicsRepository = NewDbTopicsRepo(dbHandler)

	t, err := topicInteractor.Add("Topic 1")
	if err != nil {
		test.Error(fmt.Sprintf("ERROR: Add Topic 1: %s", err))
	}
	if t.ID != 1 {
		test.Error(fmt.Sprintf("ERROR: ID Topic 1 != 1: %v", t))
	}
	t, err = topicInteractor.Add(textLike)
	if err != nil {
		test.Error(fmt.Sprintf("ERROR: Add %s: %s", textLike, err))
	}
	if t.ID != 2 {
		test.Error(fmt.Sprintf("ERROR: ID %s != 2: %v", textLike, t))
	}
	t, err = topicInteractor.Add("Topic 3")
	if err != nil {
		test.Error(fmt.Sprintf("ERROR: Add Topic 3: %s", err))
	}
	if t.ID != 3 {
		test.Error(fmt.Sprintf("ERROR: ID Topic 3 != 3: %v", t))
	}
	_, err = topicInteractor.Like(2)
	if err != nil {
		test.Error(fmt.Sprintf("ERROR: Like %s: %s", textLike, err))
	}
	topics := topicsInteractor.List()
	if len(topics.Collection) != 3 {
		test.Error(fmt.Sprintf("ERROR: Collection length != 3: %d", len(topics.Collection)))
	}
	/* First because more liked than the others */
	t = (topics.Collection)[0]
	if t.ID != 2 {
		test.Error(fmt.Sprintf("ERROR: ID %s != 2: %v", textLike, t))
	}
	if t.Text != textLike {
		test.Error(fmt.Sprintf("ERROR: Text of like (expected %s): %v", textLike, t))
	}
	if t.Like != 1 {
		test.Error(fmt.Sprintf("ERROR: Number of like (expected 1): %v", t))
	}
}
