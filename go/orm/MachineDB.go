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
var dummy_Machine sql.NullBool
var __Machine_time__dummyDeclaration time.Duration
var dummy_Machine_sort sort.Float64Slice

// MachineAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model machineAPI
type MachineAPI struct {
	gorm.Model

	models.Machine

	// encoding of pointers
	MachinePointersEnconding
}

// MachinePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type MachinePointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// MachineDB describes a machine in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model machineDB
type MachineDB struct {
	gorm.Model

	// insertion for basic fields declaration
	// Declation for basic field machineDB.TechName {{BasicKind}} (to be completed)
	TechName_Data sql.NullString

	// Declation for basic field machineDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field machineDB.DrumLoad {{BasicKind}} (to be completed)
	DrumLoad_Data sql.NullFloat64

	// Declation for basic field machineDB.RemainingTime {{BasicKind}} (to be completed)
	RemainingTime_Data sql.NullInt64

	// Declation for basic field machineDB.Cleanedlaundry bool (to be completed)
	// provide the sql storage for the boolan
	Cleanedlaundry_Data sql.NullBool

	// Declation for basic field machineDB.State {{BasicKind}} (to be completed)
	State_Data sql.NullString

	// encoding of pointers
	MachinePointersEnconding
}

// MachineDBs arrays machineDBs
// swagger:response machineDBsResponse
type MachineDBs []MachineDB

// MachineDBResponse provides response
// swagger:response machineDBResponse
type MachineDBResponse struct {
	MachineDB
}

// MachineWOP is a Machine without pointers
// it holds the same basic fields but pointers are encoded into uint
type MachineWOP struct {
	ID int

	// insertion for WOP basic fields

	TechName string

	Name string

	DrumLoad float64

	RemainingTime time.Duration

	Cleanedlaundry bool

	State models.MachineStateEnum
	// insertion for WOP pointer fields
}

var Machine_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"TechName",
	"Name",
	"DrumLoad",
	"RemainingTime",
	"Cleanedlaundry",
	"State",
}


type BackRepoMachineStruct struct {
	// stores MachineDB according to their gorm ID
	Map_MachineDBID_MachineDB *map[uint]*MachineDB

	// stores MachineDB ID according to Machine address
	Map_MachinePtr_MachineDBID *map[*models.Machine]uint

	// stores Machine according to their gorm ID
	Map_MachineDBID_MachinePtr *map[uint]*models.Machine

	db *gorm.DB
}

func (backRepoMachine *BackRepoMachineStruct) GetDB() *gorm.DB {
	return backRepoMachine.db
}

// GetMachineDBFromMachinePtr is a handy function to access the back repo instance from the stage instance
func (backRepoMachine *BackRepoMachineStruct) GetMachineDBFromMachinePtr(machine *models.Machine) (machineDB *MachineDB) {
	id := (*backRepoMachine.Map_MachinePtr_MachineDBID)[machine]
	machineDB = (*backRepoMachine.Map_MachineDBID_MachineDB)[id]
	return
}

// BackRepoMachine.Init set up the BackRepo of the Machine
func (backRepoMachine *BackRepoMachineStruct) Init(db *gorm.DB) (Error error) {

	if backRepoMachine.Map_MachineDBID_MachinePtr != nil {
		err := errors.New("In Init, backRepoMachine.Map_MachineDBID_MachinePtr should be nil")
		return err
	}

	if backRepoMachine.Map_MachineDBID_MachineDB != nil {
		err := errors.New("In Init, backRepoMachine.Map_MachineDBID_MachineDB should be nil")
		return err
	}

	if backRepoMachine.Map_MachinePtr_MachineDBID != nil {
		err := errors.New("In Init, backRepoMachine.Map_MachinePtr_MachineDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Machine, 0)
	backRepoMachine.Map_MachineDBID_MachinePtr = &tmp

	tmpDB := make(map[uint]*MachineDB, 0)
	backRepoMachine.Map_MachineDBID_MachineDB = &tmpDB

	tmpID := make(map[*models.Machine]uint, 0)
	backRepoMachine.Map_MachinePtr_MachineDBID = &tmpID

	backRepoMachine.db = db
	return
}

// BackRepoMachine.CommitPhaseOne commits all staged instances of Machine to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMachine *BackRepoMachineStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for machine := range stage.Machines {
		backRepoMachine.CommitPhaseOneInstance(machine)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, machine := range *backRepoMachine.Map_MachineDBID_MachinePtr {
		if _, ok := stage.Machines[machine]; !ok {
			backRepoMachine.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoMachine.CommitDeleteInstance commits deletion of Machine to the BackRepo
func (backRepoMachine *BackRepoMachineStruct) CommitDeleteInstance(id uint) (Error error) {

	machine := (*backRepoMachine.Map_MachineDBID_MachinePtr)[id]

	// machine is not staged anymore, remove machineDB
	machineDB := (*backRepoMachine.Map_MachineDBID_MachineDB)[id]
	query := backRepoMachine.db.Unscoped().Delete(&machineDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoMachine.Map_MachinePtr_MachineDBID), machine)
	delete((*backRepoMachine.Map_MachineDBID_MachinePtr), id)
	delete((*backRepoMachine.Map_MachineDBID_MachineDB), id)

	return
}

// BackRepoMachine.CommitPhaseOneInstance commits machine staged instances of Machine to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMachine *BackRepoMachineStruct) CommitPhaseOneInstance(machine *models.Machine) (Error error) {

	// check if the machine is not commited yet
	if _, ok := (*backRepoMachine.Map_MachinePtr_MachineDBID)[machine]; ok {
		return
	}

	// initiate machine
	var machineDB MachineDB
	machineDB.CopyBasicFieldsFromMachine(machine)

	query := backRepoMachine.db.Create(&machineDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoMachine.Map_MachinePtr_MachineDBID)[machine] = machineDB.ID
	(*backRepoMachine.Map_MachineDBID_MachinePtr)[machineDB.ID] = machine
	(*backRepoMachine.Map_MachineDBID_MachineDB)[machineDB.ID] = &machineDB

	return
}

// BackRepoMachine.CommitPhaseTwo commits all staged instances of Machine to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMachine *BackRepoMachineStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, machine := range *backRepoMachine.Map_MachineDBID_MachinePtr {
		backRepoMachine.CommitPhaseTwoInstance(backRepo, idx, machine)
	}

	return
}

// BackRepoMachine.CommitPhaseTwoInstance commits {{structname }} of models.Machine to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMachine *BackRepoMachineStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, machine *models.Machine) (Error error) {

	// fetch matching machineDB
	if machineDB, ok := (*backRepoMachine.Map_MachineDBID_MachineDB)[idx]; ok {

		machineDB.CopyBasicFieldsFromMachine(machine)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoMachine.db.Save(&machineDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Machine intance %s", machine.Name))
		return err
	}

	return
}

// BackRepoMachine.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One is the creation of instance in the stage
//
// NOTE: the is supposed to have been reset before
//
func (backRepoMachine *BackRepoMachineStruct) CheckoutPhaseOne() (Error error) {

	machineDBArray := make([]MachineDB, 0)
	query := backRepoMachine.db.Find(&machineDBArray)
	if query.Error != nil {
		return query.Error
	}

	// copy orm objects to the the map
	for _, machineDB := range machineDBArray {
		backRepoMachine.CheckoutPhaseOneInstance(&machineDB)
	}

	return
}

// CheckoutPhaseOneInstance takes a machineDB that has been found in the DB, updates the backRepo and stages the
// models version of the machineDB
func (backRepoMachine *BackRepoMachineStruct) CheckoutPhaseOneInstance(machineDB *MachineDB) (Error error) {

	machine, ok := (*backRepoMachine.Map_MachineDBID_MachinePtr)[machineDB.ID]
	if !ok {
		machine = new(models.Machine)

		(*backRepoMachine.Map_MachineDBID_MachinePtr)[machineDB.ID] = machine
		(*backRepoMachine.Map_MachinePtr_MachineDBID)[machine] = machineDB.ID

		// append model store with the new element
		machine.Name = machineDB.Name_Data.String
		machine.Stage()
	}
	machineDB.CopyBasicFieldsToMachine(machine)

	// preserve pointer to machineDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_MachineDBID_MachineDB)[machineDB hold variable pointers
	machineDB_Data := *machineDB
	preservedPtrToMachine := &machineDB_Data
	(*backRepoMachine.Map_MachineDBID_MachineDB)[machineDB.ID] = preservedPtrToMachine

	return
}

// BackRepoMachine.CheckoutPhaseTwo Checkouts all staged instances of Machine to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMachine *BackRepoMachineStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, machineDB := range *backRepoMachine.Map_MachineDBID_MachineDB {
		backRepoMachine.CheckoutPhaseTwoInstance(backRepo, machineDB)
	}
	return
}

// BackRepoMachine.CheckoutPhaseTwoInstance Checkouts staged instances of Machine to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMachine *BackRepoMachineStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, machineDB *MachineDB) (Error error) {

	machine := (*backRepoMachine.Map_MachineDBID_MachinePtr)[machineDB.ID]
	_ = machine // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitMachine allows commit of a single machine (if already staged)
func (backRepo *BackRepoStruct) CommitMachine(machine *models.Machine) {
	backRepo.BackRepoMachine.CommitPhaseOneInstance(machine)
	if id, ok := (*backRepo.BackRepoMachine.Map_MachinePtr_MachineDBID)[machine]; ok {
		backRepo.BackRepoMachine.CommitPhaseTwoInstance(backRepo, id, machine)
	}
}

// CommitMachine allows checkout of a single machine (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutMachine(machine *models.Machine) {
	// check if the machine is staged
	if _, ok := (*backRepo.BackRepoMachine.Map_MachinePtr_MachineDBID)[machine]; ok {

		if id, ok := (*backRepo.BackRepoMachine.Map_MachinePtr_MachineDBID)[machine]; ok {
			var machineDB MachineDB
			machineDB.ID = id

			if err := backRepo.BackRepoMachine.db.First(&machineDB, id).Error; err != nil {
				log.Panicln("CheckoutMachine : Problem with getting object with id:", id)
			}
			backRepo.BackRepoMachine.CheckoutPhaseOneInstance(&machineDB)
			backRepo.BackRepoMachine.CheckoutPhaseTwoInstance(backRepo, &machineDB)
		}
	}
}

// CopyBasicFieldsFromMachine
func (machineDB *MachineDB) CopyBasicFieldsFromMachine(machine *models.Machine) {
	// insertion point for fields commit
	machineDB.TechName_Data.String = machine.TechName
	machineDB.TechName_Data.Valid = true

	machineDB.Name_Data.String = machine.Name
	machineDB.Name_Data.Valid = true

	machineDB.DrumLoad_Data.Float64 = machine.DrumLoad
	machineDB.DrumLoad_Data.Valid = true

	machineDB.RemainingTime_Data.Int64 = int64(machine.RemainingTime)
	machineDB.RemainingTime_Data.Valid = true

	machineDB.Cleanedlaundry_Data.Bool = machine.Cleanedlaundry
	machineDB.Cleanedlaundry_Data.Valid = true

	machineDB.State_Data.String = string(machine.State)
	machineDB.State_Data.Valid = true

}

// CopyBasicFieldsFromMachineWOP
func (machineDB *MachineDB) CopyBasicFieldsFromMachineWOP(machine *MachineWOP) {
	// insertion point for fields commit
	machineDB.TechName_Data.String = machine.TechName
	machineDB.TechName_Data.Valid = true

	machineDB.Name_Data.String = machine.Name
	machineDB.Name_Data.Valid = true

	machineDB.DrumLoad_Data.Float64 = machine.DrumLoad
	machineDB.DrumLoad_Data.Valid = true

	machineDB.RemainingTime_Data.Int64 = int64(machine.RemainingTime)
	machineDB.RemainingTime_Data.Valid = true

	machineDB.Cleanedlaundry_Data.Bool = machine.Cleanedlaundry
	machineDB.Cleanedlaundry_Data.Valid = true

	machineDB.State_Data.String = string(machine.State)
	machineDB.State_Data.Valid = true

}

// CopyBasicFieldsToMachine
func (machineDB *MachineDB) CopyBasicFieldsToMachine(machine *models.Machine) {
	// insertion point for checkout of basic fields (back repo to stage)
	machine.TechName = machineDB.TechName_Data.String
	machine.Name = machineDB.Name_Data.String
	machine.DrumLoad = machineDB.DrumLoad_Data.Float64
	machine.RemainingTime = time.Duration(machineDB.RemainingTime_Data.Int64)
	machine.Cleanedlaundry = machineDB.Cleanedlaundry_Data.Bool
	machine.State = models.MachineStateEnum(machineDB.State_Data.String)
}

// CopyBasicFieldsToMachineWOP
func (machineDB *MachineDB) CopyBasicFieldsToMachineWOP(machine *MachineWOP) {
	machine.ID = int(machineDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	machine.TechName = machineDB.TechName_Data.String
	machine.Name = machineDB.Name_Data.String
	machine.DrumLoad = machineDB.DrumLoad_Data.Float64
	machine.RemainingTime = time.Duration(machineDB.RemainingTime_Data.Int64)
	machine.Cleanedlaundry = machineDB.Cleanedlaundry_Data.Bool
	machine.State = models.MachineStateEnum(machineDB.State_Data.String)
}

// Backup generates a json file from a slice of all MachineDB instances in the backrepo
func (backRepoMachine *BackRepoMachineStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "MachineDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MachineDB, 0)
	for _, machineDB := range *backRepoMachine.Map_MachineDBID_MachineDB {
		forBackup = append(forBackup, machineDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Machine ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Machine file", err.Error())
	}
}

// Backup generates a json file from a slice of all MachineDB instances in the backrepo
func (backRepoMachine *BackRepoMachineStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MachineDB, 0)
	for _, machineDB := range *backRepoMachine.Map_MachineDBID_MachineDB {
		forBackup = append(forBackup, machineDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Machine")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Machine_Fields, -1)
	for _, machineDB := range forBackup {

		var machineWOP MachineWOP
		machineDB.CopyBasicFieldsToMachineWOP(&machineWOP)

		row := sh.AddRow()
		row.WriteStruct(&machineWOP, -1)
	}
}

// RestorePhaseOne read the file "MachineDB.json" in dirPath that stores an array
// of MachineDB and stores it in the database
// the map BackRepoMachineid_atBckpTime_newID is updated accordingly
func (backRepoMachine *BackRepoMachineStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoMachineid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "MachineDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Machine file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*MachineDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_MachineDBID_MachineDB
	for _, machineDB := range forRestore {

		machineDB_ID_atBackupTime := machineDB.ID
		machineDB.ID = 0
		query := backRepoMachine.db.Create(machineDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoMachine.Map_MachineDBID_MachineDB)[machineDB.ID] = machineDB
		BackRepoMachineid_atBckpTime_newID[machineDB_ID_atBackupTime] = machineDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Machine file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Machine>id_atBckpTime_newID
// to compute new index
func (backRepoMachine *BackRepoMachineStruct) RestorePhaseTwo() {

	for _, machineDB := range (*backRepoMachine.Map_MachineDBID_MachineDB) {

		// next line of code is to avert unused variable compilation error
		_ = machineDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoMachine.db.Model(machineDB).Updates(*machineDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoMachineid_atBckpTime_newID map[uint]uint
