// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class WasherDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DirtyLaundryWeight: number = 0
	State: string = ""
	CleanedLaundryWeight: number = 0

	// insertion point for other declarations
}
