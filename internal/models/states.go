package models

type GuestState string

var (
	StateWaitingName        GuestState = "waiting_name"
	StateWaitingAttending   GuestState = "waiting_attending"
	StateWaitingPlusOne     GuestState = "waiting_plus_one"
	StateWaitingPlusOneName GuestState = "waiting_plus_one_name"
	StateWaitingMeal        GuestState = "waiting_meal"
	StateWaitingDrinks      GuestState = "waiting_drinks"
	StateWaitingDrinkType   GuestState = "waiting_drink_type"
	StateDone               GuestState = "done"
)
