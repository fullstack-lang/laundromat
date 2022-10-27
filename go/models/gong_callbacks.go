package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Machine:
		if stage.OnAfterMachineCreateCallback != nil {
			stage.OnAfterMachineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Simulation:
		if stage.OnAfterSimulationCreateCallback != nil {
			stage.OnAfterSimulationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Washer:
		if stage.OnAfterWasherCreateCallback != nil {
			stage.OnAfterWasherCreateCallback.OnAfterCreate(stage, target)
		}
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Machine:
		newTarget := any(new).(*Machine)
		if stage.OnAfterMachineUpdateCallback != nil {
			stage.OnAfterMachineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Simulation:
		newTarget := any(new).(*Simulation)
		if stage.OnAfterSimulationUpdateCallback != nil {
			stage.OnAfterSimulationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Washer:
		newTarget := any(new).(*Washer)
		if stage.OnAfterWasherUpdateCallback != nil {
			stage.OnAfterWasherUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Machine:
		if stage.OnAfterMachineDeleteCallback != nil {
			staged := any(staged).(*Machine)
			stage.OnAfterMachineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Simulation:
		if stage.OnAfterSimulationDeleteCallback != nil {
			staged := any(staged).(*Simulation)
			stage.OnAfterSimulationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Washer:
		if stage.OnAfterWasherDeleteCallback != nil {
			staged := any(staged).(*Washer)
			stage.OnAfterWasherDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Machine:
		if stage.OnAfterMachineReadCallback != nil {
			stage.OnAfterMachineReadCallback.OnAfterRead(stage, target)
		}
	case *Simulation:
		if stage.OnAfterSimulationReadCallback != nil {
			stage.OnAfterSimulationReadCallback.OnAfterRead(stage, target)
		}
	case *Washer:
		if stage.OnAfterWasherReadCallback != nil {
			stage.OnAfterWasherReadCallback.OnAfterRead(stage, target)
		}
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Machine:
		stage.OnAfterMachineUpdateCallback = any(callback).(OnAfterUpdateInterface[Machine])
	
	case *Simulation:
		stage.OnAfterSimulationUpdateCallback = any(callback).(OnAfterUpdateInterface[Simulation])
	
	case *Washer:
		stage.OnAfterWasherUpdateCallback = any(callback).(OnAfterUpdateInterface[Washer])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Machine:
		stage.OnAfterMachineCreateCallback = any(callback).(OnAfterCreateInterface[Machine])
	
	case *Simulation:
		stage.OnAfterSimulationCreateCallback = any(callback).(OnAfterCreateInterface[Simulation])
	
	case *Washer:
		stage.OnAfterWasherCreateCallback = any(callback).(OnAfterCreateInterface[Washer])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Machine:
		stage.OnAfterMachineDeleteCallback = any(callback).(OnAfterDeleteInterface[Machine])
	
	case *Simulation:
		stage.OnAfterSimulationDeleteCallback = any(callback).(OnAfterDeleteInterface[Simulation])
	
	case *Washer:
		stage.OnAfterWasherDeleteCallback = any(callback).(OnAfterDeleteInterface[Washer])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Machine:
		stage.OnAfterMachineReadCallback = any(callback).(OnAfterReadInterface[Machine])
	
	case *Simulation:
		stage.OnAfterSimulationReadCallback = any(callback).(OnAfterReadInterface[Simulation])
	
	case *Washer:
		stage.OnAfterWasherReadCallback = any(callback).(OnAfterReadInterface[Washer])
	
	}
}
