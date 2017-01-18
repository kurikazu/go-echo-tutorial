package models

import (
	"time"
	"github.com/gocraft/dbr"
)

type (
	Fantasy struct {
		ID        int64     `db:"id"`
		Fantasy   string    `db:"fantasy"`
		Createdat time.Time `db:"createdAt"`
	}
)

func New(fantasy string) *Fantasy {

	return &Fantasy {
		Fantasy:   fantasy,
		Createdat: time.Now(),
	}
}

func (f *Fantasy) Post(sess *dbr.Session) error {

	_, err := sess.InsertInto("fantasies").Columns("fantasy").Record(f).Exec()
	return err
}

func (f *Fantasy) Load(sess *dbr.Session, id int64) error {

	return sess.Select("*").From("fantasies").Where("id = ?", id).LoadStruct(f)
}