// generated from ng_file_enum.ts.go
export enum EngineDriverState {
	// insertion point	
	CHECKOUT_AGENT_STATES = 1,
	COMMIT_AGENT_STATES = 0,
	FIRE_ONE_EVENT = 2,
	RESET_SIMULATION = 4,
	SLEEP_100_MS = 3,
	UNKOWN = 5,
}

export interface EngineDriverStateSelect {
	value: number;
	viewValue: string;
}

export const EngineDriverStateList: EngineDriverStateSelect[] = [ // insertion point	
	{ value: EngineDriverState.CHECKOUT_AGENT_STATES, viewValue: "CHECKOUT_AGENT_STATES" },
	{ value: EngineDriverState.COMMIT_AGENT_STATES, viewValue: "COMMIT_AGENT_STATES" },
	{ value: EngineDriverState.FIRE_ONE_EVENT, viewValue: "FIRE_ONE_EVENT" },
	{ value: EngineDriverState.RESET_SIMULATION, viewValue: "RESET_SIMULATION" },
	{ value: EngineDriverState.SLEEP_100_MS, viewValue: "SLEEP_100_MS" },
	{ value: EngineDriverState.UNKOWN, viewValue: "UNKOWN" },
];
