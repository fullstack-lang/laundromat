// generated from ng_file_enum.ts.go
export enum GongsimCommandType {
	// insertion point	
	COMMAND_ADVANCE_10_MIN = "ADVANCE_10_MIN",
	COMMAND_FIRE_EVENT_TILL_STATES_CHANGE = "FIRE_EVENT_TILL_STATES_CHANGE",
	COMMAND_FIRE_NEXT_EVENT = "FIRE_NEXT_EVENT",
	COMMAND_PAUSE = "PAUSE",
	COMMAND_PLAY = "PLAY",
	COMMAND_RESET = "RESET",
}

export interface GongsimCommandTypeSelect {
	value: string;
	viewValue: string;
}

export const GongsimCommandTypeList: GongsimCommandTypeSelect[] = [ // insertion point	
	{ value: GongsimCommandType.COMMAND_ADVANCE_10_MIN, viewValue: "ADVANCE_10_MIN" },
	{ value: GongsimCommandType.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE, viewValue: "FIRE_EVENT_TILL_STATES_CHANGE" },
	{ value: GongsimCommandType.COMMAND_FIRE_NEXT_EVENT, viewValue: "FIRE_NEXT_EVENT" },
	{ value: GongsimCommandType.COMMAND_PAUSE, viewValue: "PAUSE" },
	{ value: GongsimCommandType.COMMAND_PLAY, viewValue: "PLAY" },
	{ value: GongsimCommandType.COMMAND_RESET, viewValue: "RESET" },
];
