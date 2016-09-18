package models

import (
	"time"

	"github.com/gocraft/dbr"
)

type (
	Fantasy struct {
		ID        int    `josn:"id"`
		Fantasy   string `json:"name"`
		CreatedAt int64  `json:"createdAt"`
	}
)

func New(fantasy string) *Fantasy {
	return &Fantasy{
		Fantasy:   fantasy,
		CreatedAt: time.Now().Unix(),
	}
}

func (f *Fantasy) Post(tx *dbr.Tx) error {
	_, err := tx.InsertInto("Fantasy").Columns("fantasy", "created_at").Record(f).Exec()
	return err
}

/*
func (f *Fantasy) Get(tx *dbr.Tx, id int) error {
	return err
}

func (f *Fantasy) Put(tx *dbr.Tx, id int) error {
	return err
}

func (f *Fantasy) Delete(tx *dbr.Tx, id int) error {
	return err
}
*/