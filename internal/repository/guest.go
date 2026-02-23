package repository

import (
	"database/sql"
	"wedding-invitation/internal/models"
)

type GuestRepository interface {
	UpsertGuest(g models.Guest) error
	UpdateState(tgID int64, state string) error
	SetAttending(tgID int64, attending bool) error
	SetName(tgID int64, fio string) error
	SetPlusOne(tgID int64, plusOne bool) error
	SetPlusOneName(tgID int64, fio string) error
	SetMeal(tgID int64, meal string) error
	SetDrinks(tgID int64, drinks bool) error
	SetDrinkType(tgID int64, drink string) error
	GetGuest(tgID int64) (*models.Guest, error)
}

type guestRepository struct {
	db *sql.DB
}

func NewGuestRepository(db *sql.DB) *guestRepository {
	return &guestRepository{db: db}
}

func (r *guestRepository) Save(g *models.Guest) error {
	_, err := r.db.Exec(`
	INSERT INTO guests (tg_id, name, username, state, attending, plus_one, meal)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(tg_id) DO UPDATE SET
	name = excluded.name,
	username = excluded.username,
	state = excluded.state,
	attending = excluded.attending,
	plus_one = excluded.plus_one,
	meal = excluded.meal
	`,
		g.TgID,
		g.Name,
		g.Username,
		g.State,
		g.Attending,
		g.PlusOne,
		g.Meal,
	)
	return err
}

func (r *guestRepository) FindByTgID(tgID int64) (*models.Guest, error) {
	var g models.Guest

	err := r.db.QueryRow(`
	SELECT tg_id, name, username, state
	FROM guests
	WHERE tg_id = ?
	`, tgID).Scan(
		&g.TgID,
		&g.Name,
		&g.Username,
		&g.State,
	)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
