package domain

import (
	"errors"
)

type TopicRepository interface {
	Store(t Topic) (int64, error)
	IncLike(id int64) (Topic, error)
}

type topic struct {
	id   int64
	text string
	like int64
}

type Topic struct {
	topic
}

func (t Topic) ID() int64 {
	return t.id
}

func (t Topic) Text() string {
	return t.text
}

func (t Topic) Like() int64 {
	return t.like
}

func (t Topic) SetID(id int64) Topic {
	t.id = id
	return t
}

func (t Topic) SetText(text string) Topic {
	t.text = text
	return t
}

func (t Topic) SetLike(like int64) Topic {
	t.like = like
	return t
}

func (t Topic) Create(text string) (Topic, error) {
	if text == "" {
		return t, errors.New("Empty topic")
	}
	t.text = text
	return t, nil
}
