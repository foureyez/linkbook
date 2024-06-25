package persistence

import "time"

type Collection struct {
	Id           string    `db:"Id"`
	Name         string    `db:"Name"`
	CreatedDate  time.Time `db:"CreatedDate"`
	ModifiedDate time.Time `db:"ModifiedDate"`
}
