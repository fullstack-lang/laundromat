package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/laundromat/go/models"
)

var states_Machine uml.Umlsc = uml.Umlsc{
	Name:        "Washing Machine States",
	Activestate: string(models.MACHINE_DOOR_OPEN),
	States: []*uml.State{
		{
			X:    10.000000,
			Y:    10.000000,
			Name: string(models.MACHINE_DOOR_OPEN),
		},
		{
			X:    10.000000,
			Y:    60.000000,
			Name: string(models.MACHINE_DOOR_CLOSED_IDLE),
		},
		{
			X:    10.000000,
			Y:    110.000000,
			Name: string(models.MACHINE_DOOR_CLOSED_RUNNING),
		},
	},
}
