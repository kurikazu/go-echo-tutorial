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
    }
}

func (f *Fantasy) Post(sess *dbr.Session) error {

    _, err := sess.InsertInto("fantasies").Columns("fantasy").Record(f).Exec()
    return err
}
