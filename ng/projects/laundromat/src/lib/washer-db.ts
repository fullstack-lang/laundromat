// insertion point for imports
import { MachineDB } from './machine-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class WasherDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	TechName: string = ""
	Name: string = ""
	DirtyLaundryWeight: number = 0
	State: string = ""
	CleanedLaundryWeight: number = 0

	// insertion point for other declarations
	Machine?: MachineDB
	MachineID: NullInt64 = new NullInt64 // if pointer is null, Machine.ID = 0

}
