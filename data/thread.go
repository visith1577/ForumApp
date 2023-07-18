package data

import (
	"database/sql"
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      int
	Topic     string
	UserId    int
	CreatedAt time.Time
}

var db *sql.DB

func Threads() (threads []Thread, err error) {
	rows, err := db.Query(
		"SELECT id, uuid, topic, user_id, createdat FROM threads ORDER BY createdat DESC",
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	err = rows.Close()
	if err != nil {
		log.Fatal("rows did not close properly")
		return
	}
	return
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := db.Query(
		"SELECT count(*) FROM posts WHERE thread_id = $1",
		thread.Id,
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	err = rows.Close()
	if err != nil {
		log.Fatal("rows did not close properly")
		return
	}
	return
}
