package usecases

import (
	"domain"
)

type Topic struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
	Like int64  `json:"like"`
}

type Topics struct {
	Collection []Topic
}

type TopicsRepository interface {
	Find() Topics
}

type TopicInteractor struct {
	TopicRepository domain.TopicRepository
}

type TopicsInteractor struct {
	TopicsRepository TopicsRepository
}

func (interactor TopicInteractor) Add(text string) (Topic, error) {
	td := domain.Topic{}
	td, err := td.Create(text)
	if err != nil {
		return Topic{}, err
	}
	id, err := interactor.TopicRepository.Store(td)
	if err != nil {
		return Topic{}, err
	}
	t := Topic{}
	t.ID = id
	t.Text = td.Text()
	t.Like = 0
	return t, err
}

func (interactor TopicInteractor) Like(id int64) (Topic, error) {
	td, err := interactor.TopicRepository.IncLike(id)
	if err != nil {
		return Topic{}, err
	}
	t := Topic{}
	t.ID = td.ID()
	t.Text = td.Text()
	t.Like = td.Like()
	return t, err
}

func (interactor TopicsInteractor) List() Topics {
	topics := interactor.TopicsRepository.Find()
	return topics
}
