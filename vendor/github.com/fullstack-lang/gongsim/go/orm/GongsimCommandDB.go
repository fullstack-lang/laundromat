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

	"github.com/fullstack-lang/gongsim/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongsimCommand_sql sql.NullBool
var dummy_GongsimCommand_time time.Duration
var dummy_GongsimCommand_sort sort.Float64Slice

// GongsimCommandAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongsimcommandAPI
type GongsimCommandAPI struct {
	gorm.Model

	models.GongsimCommand

	// encoding of pointers
	GongsimCommandPointersEnconding
}

// GongsimCommandPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongsimCommandPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// GongsimCommandDB describes a gongsimcommand in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongsimcommandDB
type GongsimCommandDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongsimcommandDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field gongsimcommandDB.Command {{BasicKind}} (to be completed)
	Command_Data sql.NullString

	// Declation for basic field gongsimcommandDB.CommandDate {{BasicKind}} (to be completed)
	CommandDate_Data sql.NullString

	// Declation for basic field gongsimcommandDB.SpeedCommandType {{BasicKind}} (to be completed)
	SpeedCommandType_Data sql.NullString

	// Declation for basic field gongsimcommandDB.DateSpeedCommand {{BasicKind}} (to be completed)
	DateSpeedCommand_Data sql.NullString
	// encoding of pointers
	GongsimCommandPointersEnconding
}

// GongsimCommandDBs arrays gongsimcommandDBs
// swagger:response gongsimcommandDBsResponse
type GongsimCommandDBs []GongsimCommandDB

// GongsimCommandDBResponse provides response
// swagger:response gongsimcommandDBResponse
type GongsimCommandDBResponse struct {
	GongsimCommandDB
}

// GongsimCommandWOP is a GongsimCommand without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongsimCommandWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Command models.GongsimCommandType `xlsx:"2"`

	CommandDate string `xlsx:"3"`

	SpeedCommandType models.SpeedCommandType `xlsx:"4"`

	DateSpeedCommand string `xlsx:"5"`
	// insertion for WOP pointer fields
}

var GongsimCommand_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Command",
	"CommandDate",
	"SpeedCommandType",
	"DateSpeedCommand",
}

type BackRepoGongsimCommandStruct struct {
	// stores GongsimCommandDB according to their gorm ID
	Map_GongsimCommandDBID_GongsimCommandDB *map[uint]*GongsimCommandDB

	// stores GongsimCommandDB ID according to GongsimCommand address
	Map_GongsimCommandPtr_GongsimCommandDBID *map[*models.GongsimCommand]uint

	// stores GongsimCommand according to their gorm ID
	Map_GongsimCommandDBID_GongsimCommandPtr *map[uint]*models.GongsimCommand

	db *gorm.DB
}

func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) GetDB() *gorm.DB {
	return backRepoGongsimCommand.db
}

// GetGongsimCommandDBFromGongsimCommandPtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) GetGongsimCommandDBFromGongsimCommandPtr(gongsimcommand *models.GongsimCommand) (gongsimcommandDB *GongsimCommandDB) {
	id := (*backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID)[gongsimcommand]
	gongsimcommandDB = (*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB)[id]
	return
}

// BackRepoGongsimCommand.Init set up the BackRepo of the GongsimCommand
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) Init(db *gorm.DB) (Error error) {

	if backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr != nil {
		err := errors.New("In Init, backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr should be nil")
		return err
	}

	if backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB != nil {
		err := errors.New("In Init, backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB should be nil")
		return err
	}

	if backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID != nil {
		err := errors.New("In Init, backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.GongsimCommand, 0)
	backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr = &tmp

	tmpDB := make(map[uint]*GongsimCommandDB, 0)
	backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB = &tmpDB

	tmpID := make(map[*models.GongsimCommand]uint, 0)
	backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID = &tmpID

	backRepoGongsimCommand.db = db
	return
}

// BackRepoGongsimCommand.CommitPhaseOne commits all staged instances of GongsimCommand to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for gongsimcommand := range stage.GongsimCommands {
		backRepoGongsimCommand.CommitPhaseOneInstance(gongsimcommand)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongsimcommand := range *backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr {
		if _, ok := stage.GongsimCommands[gongsimcommand]; !ok {
			backRepoGongsimCommand.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongsimCommand.CommitDeleteInstance commits deletion of GongsimCommand to the BackRepo
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) CommitDeleteInstance(id uint) (Error error) {

	gongsimcommand := (*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr)[id]

	// gongsimcommand is not staged anymore, remove gongsimcommandDB
	gongsimcommandDB := (*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB)[id]
	query := backRepoGongsimCommand.db.Unscoped().Delete(&gongsimcommandDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID), gongsimcommand)
	delete((*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr), id)
	delete((*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB), id)

	return
}

// BackRepoGongsimCommand.CommitPhaseOneInstance commits gongsimcommand staged instances of GongsimCommand to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) CommitPhaseOneInstance(gongsimcommand *models.GongsimCommand) (Error error) {

	// check if the gongsimcommand is not commited yet
	if _, ok := (*backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID)[gongsimcommand]; ok {
		return
	}

	// initiate gongsimcommand
	var gongsimcommandDB GongsimCommandDB
	gongsimcommandDB.CopyBasicFieldsFromGongsimCommand(gongsimcommand)

	query := backRepoGongsimCommand.db.Create(&gongsimcommandDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID)[gongsimcommand] = gongsimcommandDB.ID
	(*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr)[gongsimcommandDB.ID] = gongsimcommand
	(*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB)[gongsimcommandDB.ID] = &gongsimcommandDB

	return
}

// BackRepoGongsimCommand.CommitPhaseTwo commits all staged instances of GongsimCommand to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongsimcommand := range *backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr {
		backRepoGongsimCommand.CommitPhaseTwoInstance(backRepo, idx, gongsimcommand)
	}

	return
}

// BackRepoGongsimCommand.CommitPhaseTwoInstance commits {{structname }} of models.GongsimCommand to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongsimcommand *models.GongsimCommand) (Error error) {

	// fetch matching gongsimcommandDB
	if gongsimcommandDB, ok := (*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB)[idx]; ok {

		gongsimcommandDB.CopyBasicFieldsFromGongsimCommand(gongsimcommand)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoGongsimCommand.db.Save(&gongsimcommandDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongsimCommand intance %s", gongsimcommand.Name))
		return err
	}

	return
}

// BackRepoGongsimCommand.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) CheckoutPhaseOne() (Error error) {

	gongsimcommandDBArray := make([]GongsimCommandDB, 0)
	query := backRepoGongsimCommand.db.Find(&gongsimcommandDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongsimcommandInstancesToBeRemovedFromTheStage := make(map[*models.GongsimCommand]struct{})
	for key, value := range models.Stage.GongsimCommands {
		gongsimcommandInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongsimcommandDB := range gongsimcommandDBArray {
		backRepoGongsimCommand.CheckoutPhaseOneInstance(&gongsimcommandDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongsimcommand, ok := (*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr)[gongsimcommandDB.ID]
		if ok {
			delete(gongsimcommandInstancesToBeRemovedFromTheStage, gongsimcommand)
		}
	}

	// remove from stage and back repo's 3 maps all gongsimcommands that are not in the checkout
	for gongsimcommand := range gongsimcommandInstancesToBeRemovedFromTheStage {
		gongsimcommand.Unstage()

		// remove instance from the back repo 3 maps
		gongsimcommandID := (*backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID)[gongsimcommand]
		delete((*backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID), gongsimcommand)
		delete((*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB), gongsimcommandID)
		delete((*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr), gongsimcommandID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongsimcommandDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongsimcommandDB
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) CheckoutPhaseOneInstance(gongsimcommandDB *GongsimCommandDB) (Error error) {

	gongsimcommand, ok := (*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr)[gongsimcommandDB.ID]
	if !ok {
		gongsimcommand = new(models.GongsimCommand)

		(*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr)[gongsimcommandDB.ID] = gongsimcommand
		(*backRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID)[gongsimcommand] = gongsimcommandDB.ID

		// append model store with the new element
		gongsimcommand.Name = gongsimcommandDB.Name_Data.String
		gongsimcommand.Stage()
	}
	gongsimcommandDB.CopyBasicFieldsToGongsimCommand(gongsimcommand)

	// preserve pointer to gongsimcommandDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongsimCommandDBID_GongsimCommandDB)[gongsimcommandDB hold variable pointers
	gongsimcommandDB_Data := *gongsimcommandDB
	preservedPtrToGongsimCommand := &gongsimcommandDB_Data
	(*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB)[gongsimcommandDB.ID] = preservedPtrToGongsimCommand

	return
}

// BackRepoGongsimCommand.CheckoutPhaseTwo Checkouts all staged instances of GongsimCommand to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongsimcommandDB := range *backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB {
		backRepoGongsimCommand.CheckoutPhaseTwoInstance(backRepo, gongsimcommandDB)
	}
	return
}

// BackRepoGongsimCommand.CheckoutPhaseTwoInstance Checkouts staged instances of GongsimCommand to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongsimcommandDB *GongsimCommandDB) (Error error) {

	gongsimcommand := (*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr)[gongsimcommandDB.ID]
	_ = gongsimcommand // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitGongsimCommand allows commit of a single gongsimcommand (if already staged)
func (backRepo *BackRepoStruct) CommitGongsimCommand(gongsimcommand *models.GongsimCommand) {
	backRepo.BackRepoGongsimCommand.CommitPhaseOneInstance(gongsimcommand)
	if id, ok := (*backRepo.BackRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID)[gongsimcommand]; ok {
		backRepo.BackRepoGongsimCommand.CommitPhaseTwoInstance(backRepo, id, gongsimcommand)
	}
}

// CommitGongsimCommand allows checkout of a single gongsimcommand (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongsimCommand(gongsimcommand *models.GongsimCommand) {
	// check if the gongsimcommand is staged
	if _, ok := (*backRepo.BackRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID)[gongsimcommand]; ok {

		if id, ok := (*backRepo.BackRepoGongsimCommand.Map_GongsimCommandPtr_GongsimCommandDBID)[gongsimcommand]; ok {
			var gongsimcommandDB GongsimCommandDB
			gongsimcommandDB.ID = id

			if err := backRepo.BackRepoGongsimCommand.db.First(&gongsimcommandDB, id).Error; err != nil {
				log.Panicln("CheckoutGongsimCommand : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongsimCommand.CheckoutPhaseOneInstance(&gongsimcommandDB)
			backRepo.BackRepoGongsimCommand.CheckoutPhaseTwoInstance(backRepo, &gongsimcommandDB)
		}
	}
}

// CopyBasicFieldsFromGongsimCommand
func (gongsimcommandDB *GongsimCommandDB) CopyBasicFieldsFromGongsimCommand(gongsimcommand *models.GongsimCommand) {
	// insertion point for fields commit

	gongsimcommandDB.Name_Data.String = gongsimcommand.Name
	gongsimcommandDB.Name_Data.Valid = true

	gongsimcommandDB.Command_Data.String = string(gongsimcommand.Command)
	gongsimcommandDB.Command_Data.Valid = true

	gongsimcommandDB.CommandDate_Data.String = gongsimcommand.CommandDate
	gongsimcommandDB.CommandDate_Data.Valid = true

	gongsimcommandDB.SpeedCommandType_Data.String = string(gongsimcommand.SpeedCommandType)
	gongsimcommandDB.SpeedCommandType_Data.Valid = true

	gongsimcommandDB.DateSpeedCommand_Data.String = gongsimcommand.DateSpeedCommand
	gongsimcommandDB.DateSpeedCommand_Data.Valid = true
}

// CopyBasicFieldsFromGongsimCommandWOP
func (gongsimcommandDB *GongsimCommandDB) CopyBasicFieldsFromGongsimCommandWOP(gongsimcommand *GongsimCommandWOP) {
	// insertion point for fields commit

	gongsimcommandDB.Name_Data.String = gongsimcommand.Name
	gongsimcommandDB.Name_Data.Valid = true

	gongsimcommandDB.Command_Data.String = string(gongsimcommand.Command)
	gongsimcommandDB.Command_Data.Valid = true

	gongsimcommandDB.CommandDate_Data.String = gongsimcommand.CommandDate
	gongsimcommandDB.CommandDate_Data.Valid = true

	gongsimcommandDB.SpeedCommandType_Data.String = string(gongsimcommand.SpeedCommandType)
	gongsimcommandDB.SpeedCommandType_Data.Valid = true

	gongsimcommandDB.DateSpeedCommand_Data.String = gongsimcommand.DateSpeedCommand
	gongsimcommandDB.DateSpeedCommand_Data.Valid = true
}

// CopyBasicFieldsToGongsimCommand
func (gongsimcommandDB *GongsimCommandDB) CopyBasicFieldsToGongsimCommand(gongsimcommand *models.GongsimCommand) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongsimcommand.Name = gongsimcommandDB.Name_Data.String
	gongsimcommand.Command = models.GongsimCommandType(gongsimcommandDB.Command_Data.String)
	gongsimcommand.CommandDate = gongsimcommandDB.CommandDate_Data.String
	gongsimcommand.SpeedCommandType = models.SpeedCommandType(gongsimcommandDB.SpeedCommandType_Data.String)
	gongsimcommand.DateSpeedCommand = gongsimcommandDB.DateSpeedCommand_Data.String
}

// CopyBasicFieldsToGongsimCommandWOP
func (gongsimcommandDB *GongsimCommandDB) CopyBasicFieldsToGongsimCommandWOP(gongsimcommand *GongsimCommandWOP) {
	gongsimcommand.ID = int(gongsimcommandDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongsimcommand.Name = gongsimcommandDB.Name_Data.String
	gongsimcommand.Command = models.GongsimCommandType(gongsimcommandDB.Command_Data.String)
	gongsimcommand.CommandDate = gongsimcommandDB.CommandDate_Data.String
	gongsimcommand.SpeedCommandType = models.SpeedCommandType(gongsimcommandDB.SpeedCommandType_Data.String)
	gongsimcommand.DateSpeedCommand = gongsimcommandDB.DateSpeedCommand_Data.String
}

// Backup generates a json file from a slice of all GongsimCommandDB instances in the backrepo
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongsimCommandDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongsimCommandDB, 0)
	for _, gongsimcommandDB := range *backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB {
		forBackup = append(forBackup, gongsimcommandDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json GongsimCommand ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json GongsimCommand file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongsimCommandDB instances in the backrepo
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongsimCommandDB, 0)
	for _, gongsimcommandDB := range *backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB {
		forBackup = append(forBackup, gongsimcommandDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongsimCommand")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongsimCommand_Fields, -1)
	for _, gongsimcommandDB := range forBackup {

		var gongsimcommandWOP GongsimCommandWOP
		gongsimcommandDB.CopyBasicFieldsToGongsimCommandWOP(&gongsimcommandWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongsimcommandWOP, -1)
	}
}

// RestoreXL from the "GongsimCommand" sheet all GongsimCommandDB instances
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongsimCommandid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongsimCommand"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongsimCommand.rowVisitorGongsimCommand)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) rowVisitorGongsimCommand(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongsimcommandWOP GongsimCommandWOP
		row.ReadStruct(&gongsimcommandWOP)

		// add the unmarshalled struct to the stage
		gongsimcommandDB := new(GongsimCommandDB)
		gongsimcommandDB.CopyBasicFieldsFromGongsimCommandWOP(&gongsimcommandWOP)

		gongsimcommandDB_ID_atBackupTime := gongsimcommandDB.ID
		gongsimcommandDB.ID = 0
		query := backRepoGongsimCommand.db.Create(gongsimcommandDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB)[gongsimcommandDB.ID] = gongsimcommandDB
		BackRepoGongsimCommandid_atBckpTime_newID[gongsimcommandDB_ID_atBackupTime] = gongsimcommandDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongsimCommandDB.json" in dirPath that stores an array
// of GongsimCommandDB and stores it in the database
// the map BackRepoGongsimCommandid_atBckpTime_newID is updated accordingly
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongsimCommandid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongsimCommandDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json GongsimCommand file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongsimCommandDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongsimCommandDBID_GongsimCommandDB
	for _, gongsimcommandDB := range forRestore {

		gongsimcommandDB_ID_atBackupTime := gongsimcommandDB.ID
		gongsimcommandDB.ID = 0
		query := backRepoGongsimCommand.db.Create(gongsimcommandDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB)[gongsimcommandDB.ID] = gongsimcommandDB
		BackRepoGongsimCommandid_atBckpTime_newID[gongsimcommandDB_ID_atBackupTime] = gongsimcommandDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json GongsimCommand file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongsimCommand>id_atBckpTime_newID
// to compute new index
func (backRepoGongsimCommand *BackRepoGongsimCommandStruct) RestorePhaseTwo() {

	for _, gongsimcommandDB := range *backRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandDB {

		// next line of code is to avert unused variable compilation error
		_ = gongsimcommandDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoGongsimCommand.db.Model(gongsimcommandDB).Updates(*gongsimcommandDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongsimCommandid_atBckpTime_newID map[uint]uint
