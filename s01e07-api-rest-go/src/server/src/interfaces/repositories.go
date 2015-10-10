package interfaces

import (
	"database/sql"
	"log"

	"domain"
	"usecases"
)

type DbHandler interface {
	Open() error
	Close()
	Db() *sql.DB
}

type DbRepo struct {
	dbHandler DbHandler
}

type DbTopicRepo DbRepo
type DbTopicsRepo DbRepo

func NewDbTopicRepo(dbHandler DbHandler) DbTopicRepo {
	dbTopicRepo := DbTopicRepo{}
	dbTopicRepo.dbHandler = dbHandler
	return dbTopicRepo
}

func NewDbTopicsRepo(dbHandler DbHandler) DbTopicsRepo {
	dbTopicsRepo := DbTopicsRepo{}
	dbTopicsRepo.dbHandler = dbHandler
	return dbTopicsRepo
}

func (repo DbTopicRepo) Store(t domain.Topic) (int64, error) {
	err := repo.dbHandler.Open()
	if err != nil {
		log.Panic(err)
	}
	defer repo.dbHandler.Close()
	db := repo.dbHandler.Db()
	stmt, err := db.Prepare("INSERT INTO topic(text, like) values (?, 0)")
	if err != nil {
		log.Panic(err)
	}
	r, err := stmt.Exec(t.Text())
	if err != nil {
		log.Panic(err)
	}
	id, err := r.LastInsertId()
	if err != nil {
		log.Panic(err)
	}
	return id, err
}

func (repo DbTopicRepo) IncLike(pid int64) (domain.Topic, error) {
	err := repo.dbHandler.Open()
	if err != nil {
		log.Panic(err)
	}
	defer repo.dbHandler.Close()
	db := repo.dbHandler.Db()
	stmt, err := db.Prepare("UPDATE topic SET like=like+1 WHERE id=?")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(pid)
	if err != nil {
		log.Panic(err)
	}
	stmt, err = db.Prepare("SELECT * FROM topic WHERE id=?")
	if err != nil {
		log.Panic(err)
	}
	var id int64
	var text string
	var like int64
	err = stmt.QueryRow(pid).Scan(&id, &text, &like)
	if err != nil {
		log.Panic(err)
	}
	t := domain.Topic{}
	t = t.SetID(id)
	t = t.SetText(text)
	t = t.SetLike(like)
	return t, err
}

func (repo DbTopicsRepo) Find() usecases.Topics {
	err := repo.dbHandler.Open()
	if err != nil {
		log.Panic(err)
	}
	defer repo.dbHandler.Close()
	db := repo.dbHandler.Db()
	r, err := db.Query("SELECT * FROM topic ORDER BY like DESC, id")
	if err != nil {
		log.Panic(err)
	}
	defer r.Close()
	topics := usecases.Topics{}
	for r.Next() {
		var id int64
		var text string
		var like int64
		err = r.Scan(&id, &text, &like)
		if err != nil {
			log.Panic(err)
		}
		t := usecases.Topic{}
		t.ID = id
		t.Text = text
		t.Like = like
		topics.Collection = append(topics.Collection, t)
	}
	err = r.Err()
	if err != nil {
		log.Panic(err)
	}
	return topics
}
