package models

import (
	"database/sql"
)

type Guest struct {
	// Guest
	ID        int64
	TgID      int64
	Name      string
	Username  string
	FIO       sql.NullString
	Attending sql.NullBool
	// guest +1
	PlusOne    sql.NullBool
	PlusOneFio sql.NullString
	// Food
	Meal      string
	Drinks    bool
	DrinkType string

	TableID *int
	// states (wait meal, done etc.)
	State GuestState
}
