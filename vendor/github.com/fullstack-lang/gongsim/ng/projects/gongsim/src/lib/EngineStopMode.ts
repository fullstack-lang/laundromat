// generated from ng_file_enum.ts.go
export enum EngineStopMode {
	// insertion point	
	STATE_CHANGED = 1,
	TEN_MINUTES = 0,
}

export interface EngineStopModeSelect {
	value: number;
	viewValue: string;
}

export const EngineStopModeList: EngineStopModeSelect[] = [ // insertion point	
	{ value: EngineStopMode.STATE_CHANGED, viewValue: "STATE_CHANGED" },
	{ value: EngineStopMode.TEN_MINUTES, viewValue: "TEN_MINUTES" },
];
