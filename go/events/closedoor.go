package events

import "github.com/fullstack-lang/gongsim/go/models"

// CloseDoor is an event whose role is close the door
// of the machine
type CloseDoor struct {
	models.Event
}
