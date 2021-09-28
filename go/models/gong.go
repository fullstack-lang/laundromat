// generated by ModelGongFileTemplate
package models

import "sort"

// swagger:ignore
type __void struct{}

// needed for creating set of instances in the stage
var __member __void

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct { // insertion point for definition of arrays registering instances
	Machines           map[*Machine]struct{}
	Machines_mapString map[string]*Machine

	Simulations           map[*Simulation]struct{}
	Simulations_mapString map[string]*Simulation

	Washers           map[*Washer]struct{}
	Washers_mapString map[string]*Washer

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback OnInitCommitInterface
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitMachine(machine *Machine)
	CheckoutMachine(machine *Machine)
	CommitSimulation(simulation *Simulation)
	CheckoutSimulation(simulation *Simulation)
	CommitWasher(washer *Washer)
	CheckoutWasher(washer *Washer)
	GetLastCommitNb() uint
	GetLastPushFromFrontNb() uint
}

// swagger:ignore instructs the gong compiler (gongc) to avoid this particular struct
var Stage StageStruct = StageStruct{ // insertion point for array initiatialisation
	Machines:           make(map[*Machine]struct{}),
	Machines_mapString: make(map[string]*Machine),

	Simulations:           make(map[*Simulation]struct{}),
	Simulations_mapString: make(map[string]*Simulation),

	Washers:           make(map[*Washer]struct{}),
	Washers_mapString: make(map[string]*Washer),

	// end of insertion point
}

func (stage *StageStruct) Commit() {
	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
func (stage *StageStruct) getMachineOrderedStructWithNameField() []*Machine {
	// have alphabetical order generation
	machineOrdered := []*Machine{}
	for machine := range stage.Machines {
		machineOrdered = append(machineOrdered, machine)
	}
	sort.Slice(machineOrdered[:], func(i, j int) bool {
		return machineOrdered[i].Name < machineOrdered[j].Name
	})
	return machineOrdered
}

// Stage puts machine to the model stage
func (machine *Machine) Stage() *Machine {
	Stage.Machines[machine] = __member
	Stage.Machines_mapString[machine.Name] = machine

	return machine
}

// Unstage removes machine off the model stage
func (machine *Machine) Unstage() *Machine {
	delete(Stage.Machines, machine)
	delete(Stage.Machines_mapString, machine.Name)
	return machine
}

// commit machine to the back repo (if it is already staged)
func (machine *Machine) Commit() *Machine {
	if _, ok := Stage.Machines[machine]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitMachine(machine)
		}
	}
	return machine
}

// Checkout machine to the back repo (if it is already staged)
func (machine *Machine) Checkout() *Machine {
	if _, ok := Stage.Machines[machine]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutMachine(machine)
		}
	}
	return machine
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of machine to the model stage
func (machine *Machine) StageCopy() *Machine {
	_machine := new(Machine)
	*_machine = *machine
	_machine.Stage()
	return _machine
}

// StageAndCommit appends machine to the model stage and commit to the orm repo
func (machine *Machine) StageAndCommit() *Machine {
	machine.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMMachine(machine)
	}
	return machine
}

// DeleteStageAndCommit appends machine to the model stage and commit to the orm repo
func (machine *Machine) DeleteStageAndCommit() *Machine {
	machine.Unstage()
	DeleteORMMachine(machine)
	return machine
}

// StageCopyAndCommit appends a copy of machine to the model stage and commit to the orm repo
func (machine *Machine) StageCopyAndCommit() *Machine {
	_machine := new(Machine)
	*_machine = *machine
	_machine.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMMachine(machine)
	}
	return _machine
}

// CreateORMMachine enables dynamic staging of a Machine instance
func CreateORMMachine(machine *Machine) {
	machine.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMMachine(machine)
	}
}

// DeleteORMMachine enables dynamic staging of a Machine instance
func DeleteORMMachine(machine *Machine) {
	machine.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMMachine(machine)
	}
}

func (stage *StageStruct) getSimulationOrderedStructWithNameField() []*Simulation {
	// have alphabetical order generation
	simulationOrdered := []*Simulation{}
	for simulation := range stage.Simulations {
		simulationOrdered = append(simulationOrdered, simulation)
	}
	sort.Slice(simulationOrdered[:], func(i, j int) bool {
		return simulationOrdered[i].Name < simulationOrdered[j].Name
	})
	return simulationOrdered
}

// Stage puts simulation to the model stage
func (simulation *Simulation) Stage() *Simulation {
	Stage.Simulations[simulation] = __member
	Stage.Simulations_mapString[simulation.Name] = simulation

	return simulation
}

// Unstage removes simulation off the model stage
func (simulation *Simulation) Unstage() *Simulation {
	delete(Stage.Simulations, simulation)
	delete(Stage.Simulations_mapString, simulation.Name)
	return simulation
}

// commit simulation to the back repo (if it is already staged)
func (simulation *Simulation) Commit() *Simulation {
	if _, ok := Stage.Simulations[simulation]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitSimulation(simulation)
		}
	}
	return simulation
}

// Checkout simulation to the back repo (if it is already staged)
func (simulation *Simulation) Checkout() *Simulation {
	if _, ok := Stage.Simulations[simulation]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutSimulation(simulation)
		}
	}
	return simulation
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of simulation to the model stage
func (simulation *Simulation) StageCopy() *Simulation {
	_simulation := new(Simulation)
	*_simulation = *simulation
	_simulation.Stage()
	return _simulation
}

// StageAndCommit appends simulation to the model stage and commit to the orm repo
func (simulation *Simulation) StageAndCommit() *Simulation {
	simulation.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMSimulation(simulation)
	}
	return simulation
}

// DeleteStageAndCommit appends simulation to the model stage and commit to the orm repo
func (simulation *Simulation) DeleteStageAndCommit() *Simulation {
	simulation.Unstage()
	DeleteORMSimulation(simulation)
	return simulation
}

// StageCopyAndCommit appends a copy of simulation to the model stage and commit to the orm repo
func (simulation *Simulation) StageCopyAndCommit() *Simulation {
	_simulation := new(Simulation)
	*_simulation = *simulation
	_simulation.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMSimulation(simulation)
	}
	return _simulation
}

// CreateORMSimulation enables dynamic staging of a Simulation instance
func CreateORMSimulation(simulation *Simulation) {
	simulation.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMSimulation(simulation)
	}
}

// DeleteORMSimulation enables dynamic staging of a Simulation instance
func DeleteORMSimulation(simulation *Simulation) {
	simulation.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMSimulation(simulation)
	}
}

func (stage *StageStruct) getWasherOrderedStructWithNameField() []*Washer {
	// have alphabetical order generation
	washerOrdered := []*Washer{}
	for washer := range stage.Washers {
		washerOrdered = append(washerOrdered, washer)
	}
	sort.Slice(washerOrdered[:], func(i, j int) bool {
		return washerOrdered[i].Name < washerOrdered[j].Name
	})
	return washerOrdered
}

// Stage puts washer to the model stage
func (washer *Washer) Stage() *Washer {
	Stage.Washers[washer] = __member
	Stage.Washers_mapString[washer.Name] = washer

	return washer
}

// Unstage removes washer off the model stage
func (washer *Washer) Unstage() *Washer {
	delete(Stage.Washers, washer)
	delete(Stage.Washers_mapString, washer.Name)
	return washer
}

// commit washer to the back repo (if it is already staged)
func (washer *Washer) Commit() *Washer {
	if _, ok := Stage.Washers[washer]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitWasher(washer)
		}
	}
	return washer
}

// Checkout washer to the back repo (if it is already staged)
func (washer *Washer) Checkout() *Washer {
	if _, ok := Stage.Washers[washer]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutWasher(washer)
		}
	}
	return washer
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of washer to the model stage
func (washer *Washer) StageCopy() *Washer {
	_washer := new(Washer)
	*_washer = *washer
	_washer.Stage()
	return _washer
}

// StageAndCommit appends washer to the model stage and commit to the orm repo
func (washer *Washer) StageAndCommit() *Washer {
	washer.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMWasher(washer)
	}
	return washer
}

// DeleteStageAndCommit appends washer to the model stage and commit to the orm repo
func (washer *Washer) DeleteStageAndCommit() *Washer {
	washer.Unstage()
	DeleteORMWasher(washer)
	return washer
}

// StageCopyAndCommit appends a copy of washer to the model stage and commit to the orm repo
func (washer *Washer) StageCopyAndCommit() *Washer {
	_washer := new(Washer)
	*_washer = *washer
	_washer.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMWasher(washer)
	}
	return _washer
}

// CreateORMWasher enables dynamic staging of a Washer instance
func CreateORMWasher(washer *Washer) {
	washer.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMWasher(washer)
	}
}

// DeleteORMWasher enables dynamic staging of a Washer instance
func DeleteORMWasher(washer *Washer) {
	washer.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMWasher(washer)
	}
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMMachine(Machine *Machine)
	CreateORMSimulation(Simulation *Simulation)
	CreateORMWasher(Washer *Washer)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMMachine(Machine *Machine)
	DeleteORMSimulation(Simulation *Simulation)
	DeleteORMWasher(Washer *Washer)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Machines = make(map[*Machine]struct{})
	stage.Machines_mapString = make(map[string]*Machine)

	stage.Simulations = make(map[*Simulation]struct{})
	stage.Simulations_mapString = make(map[string]*Simulation)

	stage.Washers = make(map[*Washer]struct{})
	stage.Washers_mapString = make(map[string]*Washer)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Machines = nil
	stage.Machines_mapString = nil

	stage.Simulations = nil
	stage.Simulations_mapString = nil

	stage.Washers = nil
	stage.Washers_mapString = nil

}
