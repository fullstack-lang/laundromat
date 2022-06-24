// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MachineDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DrumLoad: number = 0
	RemainingTime: number = 0
	Cleanedlaundry: boolean = false
	State: string = ""

	// insertion point for other declarations
	RemainingTime_string?: string
}
