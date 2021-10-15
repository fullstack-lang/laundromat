// insertion point for imports
import { MachineDB } from './machine-db'
import { WasherDB } from './washer-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SimulationDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	LastCommitNb: number = 0

	// insertion point for other declarations
	Machine?: MachineDB
	MachineID: NullInt64 = new NullInt64 // if pointer is null, Machine.ID = 0

	Washer?: WasherDB
	WasherID: NullInt64 = new NullInt64 // if pointer is null, Washer.ID = 0

}
