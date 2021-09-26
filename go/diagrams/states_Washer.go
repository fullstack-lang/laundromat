package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/laundromat/go/models"
)

var states_Washer uml.Umlsc = uml.Umlsc{
	Name:        "Washer States",
	Activestate: string(models.WASHER_IDLE),
	States: []*uml.UmlState{
		{
			X:    10.000000,
			Y:    10.000000,
			Name: string(models.WASHER_IDLE),
		},
		{
			X:    10.000000,
			Y:    60.000000,
			Name: string(models.WASHER_OPEN_DOOR),
		},
		{
			X:    10.000000,
			Y:    110.000000,
			Name: string(models.WASHER_LOAD_DRUM),
		},
		{
			X:    10.000000,
			Y:    160.000000,
			Name: string(models.WASHER_CLOSE_DOOR),
		},
		{
			X:    10.000000,
			Y:    210.000000,
			Name: string(models.WASHER_START_PROGRAM),
		},
		{
			X:    10.000000,
			Y:    260.000000,
			Name: string(models.WASHER_WAIT_PROGRAM_END),
		},
		{
			X:    10.000000,
			Y:    310.000000,
			Name: string(models.WASHER_UNLOAD_DRUM),
		},
	},
}
