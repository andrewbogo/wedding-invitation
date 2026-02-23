package service

import (
	"log/slog"
	"wedding-invitation/internal/errors"
	"wedding-invitation/internal/models"
)

type GuestRepository interface {
	Save(g *models.Guest) error
	FindByTgID(tgID int64) (*models.Guest, error)
}

type GuestService struct {
	repo   GuestRepository
	logger *slog.Logger
}

func NewGuestService(repo GuestRepository, logger *slog.Logger) *GuestService {
	return &GuestService{repo: repo,
		logger: logger}
}

func (s *GuestService) GetGuest(tgID int64) (*models.Guest, error) {
	return s.repo.FindByTgID(tgID)
}

func (s *GuestService) UpdateState(tgID int64, state models.GuestState) error {
	guest, err := s.repo.FindByTgID(tgID)
	if err != nil {
		return errors.ErrNotFound
	}
	guest.State = state
	return s.repo.Save(guest)
}

func (s *GuestService) InitGuest(tgID int64, name, username string) error {
	guest := &models.Guest{
		TgID:     tgID,
		Name:     name,
		Username: username,
		State:    models.StateWaitingName,
	}

	return s.repo.Save(guest)
}

func (s *GuestService) SetName(tgID int64, fio string) error {
	guest, err := s.repo.FindByTgID(tgID)
	if err != nil {
		return errors.ErrNotFound
	}

	guest.FIO.Valid = true
	guest.FIO.String = fio
	guest.State = models.StateWaitingAttending

	return s.repo.Save(guest)
}

func (s *GuestService) SetPlusOneName(tgID int64, fio string) error {
	guest, err := s.repo.FindByTgID(tgID)
	if err != nil {
		return errors.ErrNotFound
	}

	guest.PlusOneFio.Valid = true
	guest.PlusOneFio.String = fio
	guest.State = models.StateWaitingMeal

	return s.repo.Save(guest)
}

func (s *GuestService) SetDrinkType(tgID int64, drink string) error {
	guest, err := s.repo.FindByTgID(tgID)
	if err != nil {
		return errors.ErrNotFound
	}

	guest.DrinkType = drink
	guest.State = models.StateDone

	return s.repo.Save(guest)
}

func (s *GuestService) SetDrinks(tgID int64, drinks bool) error {
	guest, err := s.repo.FindByTgID(tgID)
	if err != nil {
		return errors.ErrNotFound
	}

	guest.Drinks = drinks

	if drinks {
		guest.State = models.StateWaitingDrinkType
	} else {
		guest.State = models.StateDone
	}

	return s.repo.Save(guest)
}

func (s *GuestService) SetAttending(tgID int64, attending bool) error {
	guest, err := s.repo.FindByTgID(tgID)
	if err != nil {
		return errors.ErrNotFound
	}

	guest.Attending.Valid = true
	guest.Attending.Bool = attending

	if attending {
		guest.State = models.StateWaitingPlusOne
	} else {
		guest.State = models.StateDone
	}

	return s.repo.Save(guest)
}

func (s *GuestService) SetMeal(tgID int64, meal string) error {
	guest, err := s.repo.FindByTgID(tgID)
	if err != nil {
		return errors.ErrNotFound
	}

	guest.Meal = meal
	guest.State = models.StateWaitingDrinks

	return s.repo.Save(guest)
}

func (s *GuestService) SetPlusOne(tgID int64, plusOne bool) error {
	guest, err := s.repo.FindByTgID(tgID)
	if err != nil {
		return errors.ErrNotFound
	}

	guest.PlusOne.Valid = true
	guest.PlusOne.Bool = plusOne

	if plusOne {
		guest.State = models.StateWaitingPlusOneName
	} else {
		guest.State = models.StateWaitingDrinks
	}

	return s.repo.Save(guest)
}
