// insertion point for imports
import { EngineDB } from './engine-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AgentDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	TechName: string = ""

	// insertion point for other declarations
	Engine?: EngineDB
	EngineID: NullInt64 = new NullInt64 // if pointer is null, Engine.ID = 0

}
