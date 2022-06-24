// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/laundromat/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Simulation_sql sql.NullBool
var dummy_Simulation_time time.Duration
var dummy_Simulation_sort sort.Float64Slice

// SimulationAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model simulationAPI
type SimulationAPI struct {
	gorm.Model

	models.Simulation

	// encoding of pointers
	SimulationPointersEnconding
}

// SimulationPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SimulationPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field Machine is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	MachineID sql.NullInt64

	// field Washer is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	WasherID sql.NullInt64
}

// SimulationDB describes a simulation in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model simulationDB
type SimulationDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field simulationDB.Name
	Name_Data sql.NullString

	// Declation for basic field simulationDB.LastCommitNb
	LastCommitNb_Data sql.NullInt64
	// encoding of pointers
	SimulationPointersEnconding
}

// SimulationDBs arrays simulationDBs
// swagger:response simulationDBsResponse
type SimulationDBs []SimulationDB

// SimulationDBResponse provides response
// swagger:response simulationDBResponse
type SimulationDBResponse struct {
	SimulationDB
}

// SimulationWOP is a Simulation without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SimulationWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	LastCommitNb int `xlsx:"2"`
	// insertion for WOP pointer fields
}

var Simulation_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"LastCommitNb",
}

type BackRepoSimulationStruct struct {
	// stores SimulationDB according to their gorm ID
	Map_SimulationDBID_SimulationDB *map[uint]*SimulationDB

	// stores SimulationDB ID according to Simulation address
	Map_SimulationPtr_SimulationDBID *map[*models.Simulation]uint

	// stores Simulation according to their gorm ID
	Map_SimulationDBID_SimulationPtr *map[uint]*models.Simulation

	db *gorm.DB
}

func (backRepoSimulation *BackRepoSimulationStruct) GetDB() *gorm.DB {
	return backRepoSimulation.db
}

// GetSimulationDBFromSimulationPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSimulation *BackRepoSimulationStruct) GetSimulationDBFromSimulationPtr(simulation *models.Simulation) (simulationDB *SimulationDB) {
	id := (*backRepoSimulation.Map_SimulationPtr_SimulationDBID)[simulation]
	simulationDB = (*backRepoSimulation.Map_SimulationDBID_SimulationDB)[id]
	return
}

// BackRepoSimulation.Init set up the BackRepo of the Simulation
func (backRepoSimulation *BackRepoSimulationStruct) Init(db *gorm.DB) (Error error) {

	if backRepoSimulation.Map_SimulationDBID_SimulationPtr != nil {
		err := errors.New("In Init, backRepoSimulation.Map_SimulationDBID_SimulationPtr should be nil")
		return err
	}

	if backRepoSimulation.Map_SimulationDBID_SimulationDB != nil {
		err := errors.New("In Init, backRepoSimulation.Map_SimulationDBID_SimulationDB should be nil")
		return err
	}

	if backRepoSimulation.Map_SimulationPtr_SimulationDBID != nil {
		err := errors.New("In Init, backRepoSimulation.Map_SimulationPtr_SimulationDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Simulation, 0)
	backRepoSimulation.Map_SimulationDBID_SimulationPtr = &tmp

	tmpDB := make(map[uint]*SimulationDB, 0)
	backRepoSimulation.Map_SimulationDBID_SimulationDB = &tmpDB

	tmpID := make(map[*models.Simulation]uint, 0)
	backRepoSimulation.Map_SimulationPtr_SimulationDBID = &tmpID

	backRepoSimulation.db = db
	return
}

// BackRepoSimulation.CommitPhaseOne commits all staged instances of Simulation to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSimulation *BackRepoSimulationStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for simulation := range stage.Simulations {
		backRepoSimulation.CommitPhaseOneInstance(simulation)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, simulation := range *backRepoSimulation.Map_SimulationDBID_SimulationPtr {
		if _, ok := stage.Simulations[simulation]; !ok {
			backRepoSimulation.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSimulation.CommitDeleteInstance commits deletion of Simulation to the BackRepo
func (backRepoSimulation *BackRepoSimulationStruct) CommitDeleteInstance(id uint) (Error error) {

	simulation := (*backRepoSimulation.Map_SimulationDBID_SimulationPtr)[id]

	// simulation is not staged anymore, remove simulationDB
	simulationDB := (*backRepoSimulation.Map_SimulationDBID_SimulationDB)[id]
	query := backRepoSimulation.db.Unscoped().Delete(&simulationDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoSimulation.Map_SimulationPtr_SimulationDBID), simulation)
	delete((*backRepoSimulation.Map_SimulationDBID_SimulationPtr), id)
	delete((*backRepoSimulation.Map_SimulationDBID_SimulationDB), id)

	return
}

// BackRepoSimulation.CommitPhaseOneInstance commits simulation staged instances of Simulation to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSimulation *BackRepoSimulationStruct) CommitPhaseOneInstance(simulation *models.Simulation) (Error error) {

	// check if the simulation is not commited yet
	if _, ok := (*backRepoSimulation.Map_SimulationPtr_SimulationDBID)[simulation]; ok {
		return
	}

	// initiate simulation
	var simulationDB SimulationDB
	simulationDB.CopyBasicFieldsFromSimulation(simulation)

	query := backRepoSimulation.db.Create(&simulationDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoSimulation.Map_SimulationPtr_SimulationDBID)[simulation] = simulationDB.ID
	(*backRepoSimulation.Map_SimulationDBID_SimulationPtr)[simulationDB.ID] = simulation
	(*backRepoSimulation.Map_SimulationDBID_SimulationDB)[simulationDB.ID] = &simulationDB

	return
}

// BackRepoSimulation.CommitPhaseTwo commits all staged instances of Simulation to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSimulation *BackRepoSimulationStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, simulation := range *backRepoSimulation.Map_SimulationDBID_SimulationPtr {
		backRepoSimulation.CommitPhaseTwoInstance(backRepo, idx, simulation)
	}

	return
}

// BackRepoSimulation.CommitPhaseTwoInstance commits {{structname }} of models.Simulation to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSimulation *BackRepoSimulationStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, simulation *models.Simulation) (Error error) {

	// fetch matching simulationDB
	if simulationDB, ok := (*backRepoSimulation.Map_SimulationDBID_SimulationDB)[idx]; ok {

		simulationDB.CopyBasicFieldsFromSimulation(simulation)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value simulation.Machine translates to updating the simulation.MachineID
		simulationDB.MachineID.Valid = true // allow for a 0 value (nil association)
		if simulation.Machine != nil {
			if MachineId, ok := (*backRepo.BackRepoMachine.Map_MachinePtr_MachineDBID)[simulation.Machine]; ok {
				simulationDB.MachineID.Int64 = int64(MachineId)
				simulationDB.MachineID.Valid = true
			}
		}

		// commit pointer value simulation.Washer translates to updating the simulation.WasherID
		simulationDB.WasherID.Valid = true // allow for a 0 value (nil association)
		if simulation.Washer != nil {
			if WasherId, ok := (*backRepo.BackRepoWasher.Map_WasherPtr_WasherDBID)[simulation.Washer]; ok {
				simulationDB.WasherID.Int64 = int64(WasherId)
				simulationDB.WasherID.Valid = true
			}
		}

		query := backRepoSimulation.db.Save(&simulationDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Simulation intance %s", simulation.Name))
		return err
	}

	return
}

// BackRepoSimulation.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoSimulation *BackRepoSimulationStruct) CheckoutPhaseOne() (Error error) {

	simulationDBArray := make([]SimulationDB, 0)
	query := backRepoSimulation.db.Find(&simulationDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	simulationInstancesToBeRemovedFromTheStage := make(map[*models.Simulation]any)
	for key, value := range models.Stage.Simulations {
		simulationInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, simulationDB := range simulationDBArray {
		backRepoSimulation.CheckoutPhaseOneInstance(&simulationDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		simulation, ok := (*backRepoSimulation.Map_SimulationDBID_SimulationPtr)[simulationDB.ID]
		if ok {
			delete(simulationInstancesToBeRemovedFromTheStage, simulation)
		}
	}

	// remove from stage and back repo's 3 maps all simulations that are not in the checkout
	for simulation := range simulationInstancesToBeRemovedFromTheStage {
		simulation.Unstage()

		// remove instance from the back repo 3 maps
		simulationID := (*backRepoSimulation.Map_SimulationPtr_SimulationDBID)[simulation]
		delete((*backRepoSimulation.Map_SimulationPtr_SimulationDBID), simulation)
		delete((*backRepoSimulation.Map_SimulationDBID_SimulationDB), simulationID)
		delete((*backRepoSimulation.Map_SimulationDBID_SimulationPtr), simulationID)
	}

	return
}

// CheckoutPhaseOneInstance takes a simulationDB that has been found in the DB, updates the backRepo and stages the
// models version of the simulationDB
func (backRepoSimulation *BackRepoSimulationStruct) CheckoutPhaseOneInstance(simulationDB *SimulationDB) (Error error) {

	simulation, ok := (*backRepoSimulation.Map_SimulationDBID_SimulationPtr)[simulationDB.ID]
	if !ok {
		simulation = new(models.Simulation)

		(*backRepoSimulation.Map_SimulationDBID_SimulationPtr)[simulationDB.ID] = simulation
		(*backRepoSimulation.Map_SimulationPtr_SimulationDBID)[simulation] = simulationDB.ID

		// append model store with the new element
		simulation.Name = simulationDB.Name_Data.String
		simulation.Stage()
	}
	simulationDB.CopyBasicFieldsToSimulation(simulation)

	// preserve pointer to simulationDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SimulationDBID_SimulationDB)[simulationDB hold variable pointers
	simulationDB_Data := *simulationDB
	preservedPtrToSimulation := &simulationDB_Data
	(*backRepoSimulation.Map_SimulationDBID_SimulationDB)[simulationDB.ID] = preservedPtrToSimulation

	return
}

// BackRepoSimulation.CheckoutPhaseTwo Checkouts all staged instances of Simulation to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSimulation *BackRepoSimulationStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, simulationDB := range *backRepoSimulation.Map_SimulationDBID_SimulationDB {
		backRepoSimulation.CheckoutPhaseTwoInstance(backRepo, simulationDB)
	}
	return
}

// BackRepoSimulation.CheckoutPhaseTwoInstance Checkouts staged instances of Simulation to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSimulation *BackRepoSimulationStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, simulationDB *SimulationDB) (Error error) {

	simulation := (*backRepoSimulation.Map_SimulationDBID_SimulationPtr)[simulationDB.ID]
	_ = simulation // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// Machine field
	if simulationDB.MachineID.Int64 != 0 {
		simulation.Machine = (*backRepo.BackRepoMachine.Map_MachineDBID_MachinePtr)[uint(simulationDB.MachineID.Int64)]
	}
	// Washer field
	if simulationDB.WasherID.Int64 != 0 {
		simulation.Washer = (*backRepo.BackRepoWasher.Map_WasherDBID_WasherPtr)[uint(simulationDB.WasherID.Int64)]
	}
	return
}

// CommitSimulation allows commit of a single simulation (if already staged)
func (backRepo *BackRepoStruct) CommitSimulation(simulation *models.Simulation) {
	backRepo.BackRepoSimulation.CommitPhaseOneInstance(simulation)
	if id, ok := (*backRepo.BackRepoSimulation.Map_SimulationPtr_SimulationDBID)[simulation]; ok {
		backRepo.BackRepoSimulation.CommitPhaseTwoInstance(backRepo, id, simulation)
	}
}

// CommitSimulation allows checkout of a single simulation (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSimulation(simulation *models.Simulation) {
	// check if the simulation is staged
	if _, ok := (*backRepo.BackRepoSimulation.Map_SimulationPtr_SimulationDBID)[simulation]; ok {

		if id, ok := (*backRepo.BackRepoSimulation.Map_SimulationPtr_SimulationDBID)[simulation]; ok {
			var simulationDB SimulationDB
			simulationDB.ID = id

			if err := backRepo.BackRepoSimulation.db.First(&simulationDB, id).Error; err != nil {
				log.Panicln("CheckoutSimulation : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSimulation.CheckoutPhaseOneInstance(&simulationDB)
			backRepo.BackRepoSimulation.CheckoutPhaseTwoInstance(backRepo, &simulationDB)
		}
	}
}

// CopyBasicFieldsFromSimulation
func (simulationDB *SimulationDB) CopyBasicFieldsFromSimulation(simulation *models.Simulation) {
	// insertion point for fields commit

	simulationDB.Name_Data.String = simulation.Name
	simulationDB.Name_Data.Valid = true

	simulationDB.LastCommitNb_Data.Int64 = int64(simulation.LastCommitNb)
	simulationDB.LastCommitNb_Data.Valid = true
}

// CopyBasicFieldsFromSimulationWOP
func (simulationDB *SimulationDB) CopyBasicFieldsFromSimulationWOP(simulation *SimulationWOP) {
	// insertion point for fields commit

	simulationDB.Name_Data.String = simulation.Name
	simulationDB.Name_Data.Valid = true

	simulationDB.LastCommitNb_Data.Int64 = int64(simulation.LastCommitNb)
	simulationDB.LastCommitNb_Data.Valid = true
}

// CopyBasicFieldsToSimulation
func (simulationDB *SimulationDB) CopyBasicFieldsToSimulation(simulation *models.Simulation) {
	// insertion point for checkout of basic fields (back repo to stage)
	simulation.Name = simulationDB.Name_Data.String
	simulation.LastCommitNb = int(simulationDB.LastCommitNb_Data.Int64)
}

// CopyBasicFieldsToSimulationWOP
func (simulationDB *SimulationDB) CopyBasicFieldsToSimulationWOP(simulation *SimulationWOP) {
	simulation.ID = int(simulationDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	simulation.Name = simulationDB.Name_Data.String
	simulation.LastCommitNb = int(simulationDB.LastCommitNb_Data.Int64)
}

// Backup generates a json file from a slice of all SimulationDB instances in the backrepo
func (backRepoSimulation *BackRepoSimulationStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SimulationDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SimulationDB, 0)
	for _, simulationDB := range *backRepoSimulation.Map_SimulationDBID_SimulationDB {
		forBackup = append(forBackup, simulationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Simulation ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Simulation file", err.Error())
	}
}

// Backup generates a json file from a slice of all SimulationDB instances in the backrepo
func (backRepoSimulation *BackRepoSimulationStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SimulationDB, 0)
	for _, simulationDB := range *backRepoSimulation.Map_SimulationDBID_SimulationDB {
		forBackup = append(forBackup, simulationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Simulation")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Simulation_Fields, -1)
	for _, simulationDB := range forBackup {

		var simulationWOP SimulationWOP
		simulationDB.CopyBasicFieldsToSimulationWOP(&simulationWOP)

		row := sh.AddRow()
		row.WriteStruct(&simulationWOP, -1)
	}
}

// RestoreXL from the "Simulation" sheet all SimulationDB instances
func (backRepoSimulation *BackRepoSimulationStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSimulationid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Simulation"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSimulation.rowVisitorSimulation)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoSimulation *BackRepoSimulationStruct) rowVisitorSimulation(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var simulationWOP SimulationWOP
		row.ReadStruct(&simulationWOP)

		// add the unmarshalled struct to the stage
		simulationDB := new(SimulationDB)
		simulationDB.CopyBasicFieldsFromSimulationWOP(&simulationWOP)

		simulationDB_ID_atBackupTime := simulationDB.ID
		simulationDB.ID = 0
		query := backRepoSimulation.db.Create(simulationDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoSimulation.Map_SimulationDBID_SimulationDB)[simulationDB.ID] = simulationDB
		BackRepoSimulationid_atBckpTime_newID[simulationDB_ID_atBackupTime] = simulationDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SimulationDB.json" in dirPath that stores an array
// of SimulationDB and stores it in the database
// the map BackRepoSimulationid_atBckpTime_newID is updated accordingly
func (backRepoSimulation *BackRepoSimulationStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSimulationid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SimulationDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Simulation file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SimulationDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SimulationDBID_SimulationDB
	for _, simulationDB := range forRestore {

		simulationDB_ID_atBackupTime := simulationDB.ID
		simulationDB.ID = 0
		query := backRepoSimulation.db.Create(simulationDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoSimulation.Map_SimulationDBID_SimulationDB)[simulationDB.ID] = simulationDB
		BackRepoSimulationid_atBckpTime_newID[simulationDB_ID_atBackupTime] = simulationDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Simulation file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Simulation>id_atBckpTime_newID
// to compute new index
func (backRepoSimulation *BackRepoSimulationStruct) RestorePhaseTwo() {

	for _, simulationDB := range *backRepoSimulation.Map_SimulationDBID_SimulationDB {

		// next line of code is to avert unused variable compilation error
		_ = simulationDB

		// insertion point for reindexing pointers encoding
		// reindexing Machine field
		if simulationDB.MachineID.Int64 != 0 {
			simulationDB.MachineID.Int64 = int64(BackRepoMachineid_atBckpTime_newID[uint(simulationDB.MachineID.Int64)])
			simulationDB.MachineID.Valid = true
		}

		// reindexing Washer field
		if simulationDB.WasherID.Int64 != 0 {
			simulationDB.WasherID.Int64 = int64(BackRepoWasherid_atBckpTime_newID[uint(simulationDB.WasherID.Int64)])
			simulationDB.WasherID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoSimulation.db.Model(simulationDB).Updates(*simulationDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSimulationid_atBckpTime_newID map[uint]uint
