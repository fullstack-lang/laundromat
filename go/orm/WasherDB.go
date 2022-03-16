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
var dummy_Washer_sql sql.NullBool
var dummy_Washer_time time.Duration
var dummy_Washer_sort sort.Float64Slice

// WasherAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model washerAPI
type WasherAPI struct {
	gorm.Model

	models.Washer

	// encoding of pointers
	WasherPointersEnconding
}

// WasherPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type WasherPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field Machine is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	MachineID sql.NullInt64
}

// WasherDB describes a washer in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model washerDB
type WasherDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field washerDB.TechName {{BasicKind}} (to be completed)
	TechName_Data sql.NullString

	// Declation for basic field washerDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field washerDB.DirtyLaundryWeight {{BasicKind}} (to be completed)
	DirtyLaundryWeight_Data sql.NullFloat64

	// Declation for basic field washerDB.State {{BasicKind}} (to be completed)
	State_Data sql.NullString

	// Declation for basic field washerDB.CleanedLaundryWeight {{BasicKind}} (to be completed)
	CleanedLaundryWeight_Data sql.NullFloat64
	// encoding of pointers
	WasherPointersEnconding
}

// WasherDBs arrays washerDBs
// swagger:response washerDBsResponse
type WasherDBs []WasherDB

// WasherDBResponse provides response
// swagger:response washerDBResponse
type WasherDBResponse struct {
	WasherDB
}

// WasherWOP is a Washer without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type WasherWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	TechName string `xlsx:"1"`

	Name string `xlsx:"2"`

	DirtyLaundryWeight float64 `xlsx:"3"`

	State models.WasherStateEnum `xlsx:"4"`

	CleanedLaundryWeight float64 `xlsx:"5"`
	// insertion for WOP pointer fields
}

var Washer_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"TechName",
	"Name",
	"DirtyLaundryWeight",
	"State",
	"CleanedLaundryWeight",
}

type BackRepoWasherStruct struct {
	// stores WasherDB according to their gorm ID
	Map_WasherDBID_WasherDB *map[uint]*WasherDB

	// stores WasherDB ID according to Washer address
	Map_WasherPtr_WasherDBID *map[*models.Washer]uint

	// stores Washer according to their gorm ID
	Map_WasherDBID_WasherPtr *map[uint]*models.Washer

	db *gorm.DB
}

func (backRepoWasher *BackRepoWasherStruct) GetDB() *gorm.DB {
	return backRepoWasher.db
}

// GetWasherDBFromWasherPtr is a handy function to access the back repo instance from the stage instance
func (backRepoWasher *BackRepoWasherStruct) GetWasherDBFromWasherPtr(washer *models.Washer) (washerDB *WasherDB) {
	id := (*backRepoWasher.Map_WasherPtr_WasherDBID)[washer]
	washerDB = (*backRepoWasher.Map_WasherDBID_WasherDB)[id]
	return
}

// BackRepoWasher.Init set up the BackRepo of the Washer
func (backRepoWasher *BackRepoWasherStruct) Init(db *gorm.DB) (Error error) {

	if backRepoWasher.Map_WasherDBID_WasherPtr != nil {
		err := errors.New("In Init, backRepoWasher.Map_WasherDBID_WasherPtr should be nil")
		return err
	}

	if backRepoWasher.Map_WasherDBID_WasherDB != nil {
		err := errors.New("In Init, backRepoWasher.Map_WasherDBID_WasherDB should be nil")
		return err
	}

	if backRepoWasher.Map_WasherPtr_WasherDBID != nil {
		err := errors.New("In Init, backRepoWasher.Map_WasherPtr_WasherDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Washer, 0)
	backRepoWasher.Map_WasherDBID_WasherPtr = &tmp

	tmpDB := make(map[uint]*WasherDB, 0)
	backRepoWasher.Map_WasherDBID_WasherDB = &tmpDB

	tmpID := make(map[*models.Washer]uint, 0)
	backRepoWasher.Map_WasherPtr_WasherDBID = &tmpID

	backRepoWasher.db = db
	return
}

// BackRepoWasher.CommitPhaseOne commits all staged instances of Washer to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoWasher *BackRepoWasherStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for washer := range stage.Washers {
		backRepoWasher.CommitPhaseOneInstance(washer)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, washer := range *backRepoWasher.Map_WasherDBID_WasherPtr {
		if _, ok := stage.Washers[washer]; !ok {
			backRepoWasher.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoWasher.CommitDeleteInstance commits deletion of Washer to the BackRepo
func (backRepoWasher *BackRepoWasherStruct) CommitDeleteInstance(id uint) (Error error) {

	washer := (*backRepoWasher.Map_WasherDBID_WasherPtr)[id]

	// washer is not staged anymore, remove washerDB
	washerDB := (*backRepoWasher.Map_WasherDBID_WasherDB)[id]
	query := backRepoWasher.db.Unscoped().Delete(&washerDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoWasher.Map_WasherPtr_WasherDBID), washer)
	delete((*backRepoWasher.Map_WasherDBID_WasherPtr), id)
	delete((*backRepoWasher.Map_WasherDBID_WasherDB), id)

	return
}

// BackRepoWasher.CommitPhaseOneInstance commits washer staged instances of Washer to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoWasher *BackRepoWasherStruct) CommitPhaseOneInstance(washer *models.Washer) (Error error) {

	// check if the washer is not commited yet
	if _, ok := (*backRepoWasher.Map_WasherPtr_WasherDBID)[washer]; ok {
		return
	}

	// initiate washer
	var washerDB WasherDB
	washerDB.CopyBasicFieldsFromWasher(washer)

	query := backRepoWasher.db.Create(&washerDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoWasher.Map_WasherPtr_WasherDBID)[washer] = washerDB.ID
	(*backRepoWasher.Map_WasherDBID_WasherPtr)[washerDB.ID] = washer
	(*backRepoWasher.Map_WasherDBID_WasherDB)[washerDB.ID] = &washerDB

	return
}

// BackRepoWasher.CommitPhaseTwo commits all staged instances of Washer to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWasher *BackRepoWasherStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, washer := range *backRepoWasher.Map_WasherDBID_WasherPtr {
		backRepoWasher.CommitPhaseTwoInstance(backRepo, idx, washer)
	}

	return
}

// BackRepoWasher.CommitPhaseTwoInstance commits {{structname }} of models.Washer to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWasher *BackRepoWasherStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, washer *models.Washer) (Error error) {

	// fetch matching washerDB
	if washerDB, ok := (*backRepoWasher.Map_WasherDBID_WasherDB)[idx]; ok {

		washerDB.CopyBasicFieldsFromWasher(washer)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value washer.Machine translates to updating the washer.MachineID
		washerDB.MachineID.Valid = true // allow for a 0 value (nil association)
		if washer.Machine != nil {
			if MachineId, ok := (*backRepo.BackRepoMachine.Map_MachinePtr_MachineDBID)[washer.Machine]; ok {
				washerDB.MachineID.Int64 = int64(MachineId)
				washerDB.MachineID.Valid = true
			}
		}

		query := backRepoWasher.db.Save(&washerDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Washer intance %s", washer.Name))
		return err
	}

	return
}

// BackRepoWasher.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoWasher *BackRepoWasherStruct) CheckoutPhaseOne() (Error error) {

	washerDBArray := make([]WasherDB, 0)
	query := backRepoWasher.db.Find(&washerDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	washerInstancesToBeRemovedFromTheStage := make(map[*models.Washer]struct{})
	for key, value := range models.Stage.Washers {
		washerInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, washerDB := range washerDBArray {
		backRepoWasher.CheckoutPhaseOneInstance(&washerDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		washer, ok := (*backRepoWasher.Map_WasherDBID_WasherPtr)[washerDB.ID]
		if ok {
			delete(washerInstancesToBeRemovedFromTheStage, washer)
		}
	}

	// remove from stage and back repo's 3 maps all washers that are not in the checkout
	for washer := range washerInstancesToBeRemovedFromTheStage {
		washer.Unstage()

		// remove instance from the back repo 3 maps
		washerID := (*backRepoWasher.Map_WasherPtr_WasherDBID)[washer]
		delete((*backRepoWasher.Map_WasherPtr_WasherDBID), washer)
		delete((*backRepoWasher.Map_WasherDBID_WasherDB), washerID)
		delete((*backRepoWasher.Map_WasherDBID_WasherPtr), washerID)
	}

	return
}

// CheckoutPhaseOneInstance takes a washerDB that has been found in the DB, updates the backRepo and stages the
// models version of the washerDB
func (backRepoWasher *BackRepoWasherStruct) CheckoutPhaseOneInstance(washerDB *WasherDB) (Error error) {

	washer, ok := (*backRepoWasher.Map_WasherDBID_WasherPtr)[washerDB.ID]
	if !ok {
		washer = new(models.Washer)

		(*backRepoWasher.Map_WasherDBID_WasherPtr)[washerDB.ID] = washer
		(*backRepoWasher.Map_WasherPtr_WasherDBID)[washer] = washerDB.ID

		// append model store with the new element
		washer.Name = washerDB.Name_Data.String
		washer.Stage()
	}
	washerDB.CopyBasicFieldsToWasher(washer)

	// preserve pointer to washerDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_WasherDBID_WasherDB)[washerDB hold variable pointers
	washerDB_Data := *washerDB
	preservedPtrToWasher := &washerDB_Data
	(*backRepoWasher.Map_WasherDBID_WasherDB)[washerDB.ID] = preservedPtrToWasher

	return
}

// BackRepoWasher.CheckoutPhaseTwo Checkouts all staged instances of Washer to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWasher *BackRepoWasherStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, washerDB := range *backRepoWasher.Map_WasherDBID_WasherDB {
		backRepoWasher.CheckoutPhaseTwoInstance(backRepo, washerDB)
	}
	return
}

// BackRepoWasher.CheckoutPhaseTwoInstance Checkouts staged instances of Washer to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWasher *BackRepoWasherStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, washerDB *WasherDB) (Error error) {

	washer := (*backRepoWasher.Map_WasherDBID_WasherPtr)[washerDB.ID]
	_ = washer // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// Machine field
	if washerDB.MachineID.Int64 != 0 {
		washer.Machine = (*backRepo.BackRepoMachine.Map_MachineDBID_MachinePtr)[uint(washerDB.MachineID.Int64)]
	}
	return
}

// CommitWasher allows commit of a single washer (if already staged)
func (backRepo *BackRepoStruct) CommitWasher(washer *models.Washer) {
	backRepo.BackRepoWasher.CommitPhaseOneInstance(washer)
	if id, ok := (*backRepo.BackRepoWasher.Map_WasherPtr_WasherDBID)[washer]; ok {
		backRepo.BackRepoWasher.CommitPhaseTwoInstance(backRepo, id, washer)
	}
}

// CommitWasher allows checkout of a single washer (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutWasher(washer *models.Washer) {
	// check if the washer is staged
	if _, ok := (*backRepo.BackRepoWasher.Map_WasherPtr_WasherDBID)[washer]; ok {

		if id, ok := (*backRepo.BackRepoWasher.Map_WasherPtr_WasherDBID)[washer]; ok {
			var washerDB WasherDB
			washerDB.ID = id

			if err := backRepo.BackRepoWasher.db.First(&washerDB, id).Error; err != nil {
				log.Panicln("CheckoutWasher : Problem with getting object with id:", id)
			}
			backRepo.BackRepoWasher.CheckoutPhaseOneInstance(&washerDB)
			backRepo.BackRepoWasher.CheckoutPhaseTwoInstance(backRepo, &washerDB)
		}
	}
}

// CopyBasicFieldsFromWasher
func (washerDB *WasherDB) CopyBasicFieldsFromWasher(washer *models.Washer) {
	// insertion point for fields commit

	washerDB.TechName_Data.String = washer.TechName
	washerDB.TechName_Data.Valid = true

	washerDB.Name_Data.String = washer.Name
	washerDB.Name_Data.Valid = true

	washerDB.DirtyLaundryWeight_Data.Float64 = washer.DirtyLaundryWeight
	washerDB.DirtyLaundryWeight_Data.Valid = true

	washerDB.State_Data.String = washer.State.ToString()
	washerDB.State_Data.Valid = true

	washerDB.CleanedLaundryWeight_Data.Float64 = washer.CleanedLaundryWeight
	washerDB.CleanedLaundryWeight_Data.Valid = true
}

// CopyBasicFieldsFromWasherWOP
func (washerDB *WasherDB) CopyBasicFieldsFromWasherWOP(washer *WasherWOP) {
	// insertion point for fields commit

	washerDB.TechName_Data.String = washer.TechName
	washerDB.TechName_Data.Valid = true

	washerDB.Name_Data.String = washer.Name
	washerDB.Name_Data.Valid = true

	washerDB.DirtyLaundryWeight_Data.Float64 = washer.DirtyLaundryWeight
	washerDB.DirtyLaundryWeight_Data.Valid = true

	washerDB.State_Data.String = washer.State.ToString()
	washerDB.State_Data.Valid = true

	washerDB.CleanedLaundryWeight_Data.Float64 = washer.CleanedLaundryWeight
	washerDB.CleanedLaundryWeight_Data.Valid = true
}

// CopyBasicFieldsToWasher
func (washerDB *WasherDB) CopyBasicFieldsToWasher(washer *models.Washer) {
	// insertion point for checkout of basic fields (back repo to stage)
	washer.TechName = washerDB.TechName_Data.String
	washer.Name = washerDB.Name_Data.String
	washer.DirtyLaundryWeight = washerDB.DirtyLaundryWeight_Data.Float64
	washer.State.FromString(washerDB.State_Data.String)
	washer.CleanedLaundryWeight = washerDB.CleanedLaundryWeight_Data.Float64
}

// CopyBasicFieldsToWasherWOP
func (washerDB *WasherDB) CopyBasicFieldsToWasherWOP(washer *WasherWOP) {
	washer.ID = int(washerDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	washer.TechName = washerDB.TechName_Data.String
	washer.Name = washerDB.Name_Data.String
	washer.DirtyLaundryWeight = washerDB.DirtyLaundryWeight_Data.Float64
	washer.State.FromString(washerDB.State_Data.String)
	washer.CleanedLaundryWeight = washerDB.CleanedLaundryWeight_Data.Float64
}

// Backup generates a json file from a slice of all WasherDB instances in the backrepo
func (backRepoWasher *BackRepoWasherStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "WasherDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*WasherDB, 0)
	for _, washerDB := range *backRepoWasher.Map_WasherDBID_WasherDB {
		forBackup = append(forBackup, washerDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Washer ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Washer file", err.Error())
	}
}

// Backup generates a json file from a slice of all WasherDB instances in the backrepo
func (backRepoWasher *BackRepoWasherStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*WasherDB, 0)
	for _, washerDB := range *backRepoWasher.Map_WasherDBID_WasherDB {
		forBackup = append(forBackup, washerDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Washer")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Washer_Fields, -1)
	for _, washerDB := range forBackup {

		var washerWOP WasherWOP
		washerDB.CopyBasicFieldsToWasherWOP(&washerWOP)

		row := sh.AddRow()
		row.WriteStruct(&washerWOP, -1)
	}
}

// RestoreXL from the "Washer" sheet all WasherDB instances
func (backRepoWasher *BackRepoWasherStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoWasherid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Washer"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoWasher.rowVisitorWasher)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoWasher *BackRepoWasherStruct) rowVisitorWasher(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var washerWOP WasherWOP
		row.ReadStruct(&washerWOP)

		// add the unmarshalled struct to the stage
		washerDB := new(WasherDB)
		washerDB.CopyBasicFieldsFromWasherWOP(&washerWOP)

		washerDB_ID_atBackupTime := washerDB.ID
		washerDB.ID = 0
		query := backRepoWasher.db.Create(washerDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoWasher.Map_WasherDBID_WasherDB)[washerDB.ID] = washerDB
		BackRepoWasherid_atBckpTime_newID[washerDB_ID_atBackupTime] = washerDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "WasherDB.json" in dirPath that stores an array
// of WasherDB and stores it in the database
// the map BackRepoWasherid_atBckpTime_newID is updated accordingly
func (backRepoWasher *BackRepoWasherStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoWasherid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "WasherDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Washer file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*WasherDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_WasherDBID_WasherDB
	for _, washerDB := range forRestore {

		washerDB_ID_atBackupTime := washerDB.ID
		washerDB.ID = 0
		query := backRepoWasher.db.Create(washerDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoWasher.Map_WasherDBID_WasherDB)[washerDB.ID] = washerDB
		BackRepoWasherid_atBckpTime_newID[washerDB_ID_atBackupTime] = washerDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Washer file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Washer>id_atBckpTime_newID
// to compute new index
func (backRepoWasher *BackRepoWasherStruct) RestorePhaseTwo() {

	for _, washerDB := range *backRepoWasher.Map_WasherDBID_WasherDB {

		// next line of code is to avert unused variable compilation error
		_ = washerDB

		// insertion point for reindexing pointers encoding
		// reindexing Machine field
		if washerDB.MachineID.Int64 != 0 {
			washerDB.MachineID.Int64 = int64(BackRepoMachineid_atBckpTime_newID[uint(washerDB.MachineID.Int64)])
			washerDB.MachineID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoWasher.db.Model(washerDB).Updates(*washerDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoWasherid_atBckpTime_newID map[uint]uint
