package enum

import "fmt"

type OrderState int

const (
	StateCreated OrderState = iota
	StatePaid
	StatePacked
	StateShipped
	StateDelivered
	StateCancelled
	StateReturned
)

type OrderAction string

const (
	ActionPay     OrderAction = "pay"
	ActionPack    OrderAction = "pack"
	ActionShip    OrderAction = "ship"
	ActionDeliver OrderAction = "deliver"
	ActionCancel  OrderAction = "cancel"
	ActionReturn  OrderAction = "return"
)

var stateActionMappings = map[OrderState]string{
	StateCreated:   "created",
	StatePaid:      "paid",
	StatePacked:    "packed",
	StateShipped:   "shipped",
	StateDelivered: "delivered",
	StateCancelled: "cancelled",
	StateReturned:  "returned",
}

func (s OrderState) String() string {
	return stateActionMappings[s]
}

func NextState(current OrderState, action string) OrderState {
	switch current {
	case StateCreated:
		if action == string(ActionPay) {
			return StatePaid
		} else if action == string(ActionCancel) {
			return StateCancelled
		}
		return current
	case StatePaid:
		if action == string(ActionPack) {
			return StatePacked
		} else if action == string(ActionCancel) {
			return StateCancelled
		}
		return current
	case StatePacked:
		if action == string(ActionShip) {
			return StateShipped
		}
		return current
	case StateShipped:
		if action == string(ActionDeliver) {
			return StateDelivered
		} else if action == string(ActionReturn) {
			return StateReturned
		}
		return current
	case StateDelivered:
		if action == string(ActionReturn) {
			return StateReturned
		}
		return current
	case StateCancelled, StateReturned:
		return current
	default:
		panic(fmt.Errorf("unknown state: %s", current))
	}
}
