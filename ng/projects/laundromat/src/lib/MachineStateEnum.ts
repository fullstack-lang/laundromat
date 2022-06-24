// generated from ng_file_enum.ts.go
export enum MachineStateEnum {
	// insertion point	
	MACHINE_DOOR_OPEN = "MACHINE_DOOR_OPEN",
	MACHINE_DOOR_CLOSED_RUNNING = "MACHINE_DOOR_CLOSED_RUNNING",
	MACHINE_DOOR_CLOSED_IDLE = "MACHINE_DOOR_CLOSED_IDLE",
}

export interface MachineStateEnumSelect {
	value: string;
	viewValue: string;
}

export const MachineStateEnumList: MachineStateEnumSelect[] = [ // insertion point	
	{ value: MachineStateEnum.MACHINE_DOOR_OPEN, viewValue: "MACHINE_DOOR_OPEN" },
	{ value: MachineStateEnum.MACHINE_DOOR_CLOSED_RUNNING, viewValue: "MACHINE_DOOR_CLOSED_RUNNING" },
	{ value: MachineStateEnum.MACHINE_DOOR_CLOSED_IDLE, viewValue: "MACHINE_DOOR_CLOSED_IDLE" },
];
