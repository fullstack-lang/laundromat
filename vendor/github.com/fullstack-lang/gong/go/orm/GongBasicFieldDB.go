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

	"github.com/fullstack-lang/gong/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongBasicField_sql sql.NullBool
var dummy_GongBasicField_time time.Duration
var dummy_GongBasicField_sort sort.Float64Slice

// GongBasicFieldAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongbasicfieldAPI
type GongBasicFieldAPI struct {
	gorm.Model

	models.GongBasicField

	// encoding of pointers
	GongBasicFieldPointersEnconding
}

// GongBasicFieldPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongBasicFieldPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field GongEnum is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	GongEnumID sql.NullInt64

	// Implementation of a reverse ID for field GongStruct{}.GongBasicFields []*GongBasicField
	GongStruct_GongBasicFieldsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	GongStruct_GongBasicFieldsDBID_Index sql.NullInt64
}

// GongBasicFieldDB describes a gongbasicfield in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongbasicfieldDB
type GongBasicFieldDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongbasicfieldDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field gongbasicfieldDB.BasicKindName {{BasicKind}} (to be completed)
	BasicKindName_Data sql.NullString

	// Declation for basic field gongbasicfieldDB.DeclaredType {{BasicKind}} (to be completed)
	DeclaredType_Data sql.NullString

	// Declation for basic field gongbasicfieldDB.Index {{BasicKind}} (to be completed)
	Index_Data sql.NullInt64
	// encoding of pointers
	GongBasicFieldPointersEnconding
}

// GongBasicFieldDBs arrays gongbasicfieldDBs
// swagger:response gongbasicfieldDBsResponse
type GongBasicFieldDBs []GongBasicFieldDB

// GongBasicFieldDBResponse provides response
// swagger:response gongbasicfieldDBResponse
type GongBasicFieldDBResponse struct {
	GongBasicFieldDB
}

// GongBasicFieldWOP is a GongBasicField without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongBasicFieldWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	BasicKindName string `xlsx:"2"`

	DeclaredType string `xlsx:"3"`

	Index int `xlsx:"4"`
	// insertion for WOP pointer fields
}

var GongBasicField_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"BasicKindName",
	"DeclaredType",
	"Index",
}

type BackRepoGongBasicFieldStruct struct {
	// stores GongBasicFieldDB according to their gorm ID
	Map_GongBasicFieldDBID_GongBasicFieldDB *map[uint]*GongBasicFieldDB

	// stores GongBasicFieldDB ID according to GongBasicField address
	Map_GongBasicFieldPtr_GongBasicFieldDBID *map[*models.GongBasicField]uint

	// stores GongBasicField according to their gorm ID
	Map_GongBasicFieldDBID_GongBasicFieldPtr *map[uint]*models.GongBasicField

	db *gorm.DB
}

func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) GetDB() *gorm.DB {
	return backRepoGongBasicField.db
}

// GetGongBasicFieldDBFromGongBasicFieldPtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) GetGongBasicFieldDBFromGongBasicFieldPtr(gongbasicfield *models.GongBasicField) (gongbasicfieldDB *GongBasicFieldDB) {
	id := (*backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID)[gongbasicfield]
	gongbasicfieldDB = (*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB)[id]
	return
}

// BackRepoGongBasicField.Init set up the BackRepo of the GongBasicField
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) Init(db *gorm.DB) (Error error) {

	if backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr != nil {
		err := errors.New("In Init, backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr should be nil")
		return err
	}

	if backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB != nil {
		err := errors.New("In Init, backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB should be nil")
		return err
	}

	if backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID != nil {
		err := errors.New("In Init, backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.GongBasicField, 0)
	backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr = &tmp

	tmpDB := make(map[uint]*GongBasicFieldDB, 0)
	backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB = &tmpDB

	tmpID := make(map[*models.GongBasicField]uint, 0)
	backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID = &tmpID

	backRepoGongBasicField.db = db
	return
}

// BackRepoGongBasicField.CommitPhaseOne commits all staged instances of GongBasicField to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for gongbasicfield := range stage.GongBasicFields {
		backRepoGongBasicField.CommitPhaseOneInstance(gongbasicfield)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongbasicfield := range *backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr {
		if _, ok := stage.GongBasicFields[gongbasicfield]; !ok {
			backRepoGongBasicField.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongBasicField.CommitDeleteInstance commits deletion of GongBasicField to the BackRepo
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) CommitDeleteInstance(id uint) (Error error) {

	gongbasicfield := (*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr)[id]

	// gongbasicfield is not staged anymore, remove gongbasicfieldDB
	gongbasicfieldDB := (*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB)[id]
	query := backRepoGongBasicField.db.Unscoped().Delete(&gongbasicfieldDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID), gongbasicfield)
	delete((*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr), id)
	delete((*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB), id)

	return
}

// BackRepoGongBasicField.CommitPhaseOneInstance commits gongbasicfield staged instances of GongBasicField to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) CommitPhaseOneInstance(gongbasicfield *models.GongBasicField) (Error error) {

	// check if the gongbasicfield is not commited yet
	if _, ok := (*backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID)[gongbasicfield]; ok {
		return
	}

	// initiate gongbasicfield
	var gongbasicfieldDB GongBasicFieldDB
	gongbasicfieldDB.CopyBasicFieldsFromGongBasicField(gongbasicfield)

	query := backRepoGongBasicField.db.Create(&gongbasicfieldDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID)[gongbasicfield] = gongbasicfieldDB.ID
	(*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr)[gongbasicfieldDB.ID] = gongbasicfield
	(*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB)[gongbasicfieldDB.ID] = &gongbasicfieldDB

	return
}

// BackRepoGongBasicField.CommitPhaseTwo commits all staged instances of GongBasicField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongbasicfield := range *backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr {
		backRepoGongBasicField.CommitPhaseTwoInstance(backRepo, idx, gongbasicfield)
	}

	return
}

// BackRepoGongBasicField.CommitPhaseTwoInstance commits {{structname }} of models.GongBasicField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongbasicfield *models.GongBasicField) (Error error) {

	// fetch matching gongbasicfieldDB
	if gongbasicfieldDB, ok := (*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB)[idx]; ok {

		gongbasicfieldDB.CopyBasicFieldsFromGongBasicField(gongbasicfield)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value gongbasicfield.GongEnum translates to updating the gongbasicfield.GongEnumID
		gongbasicfieldDB.GongEnumID.Valid = true // allow for a 0 value (nil association)
		if gongbasicfield.GongEnum != nil {
			if GongEnumId, ok := (*backRepo.BackRepoGongEnum.Map_GongEnumPtr_GongEnumDBID)[gongbasicfield.GongEnum]; ok {
				gongbasicfieldDB.GongEnumID.Int64 = int64(GongEnumId)
				gongbasicfieldDB.GongEnumID.Valid = true
			}
		}

		query := backRepoGongBasicField.db.Save(&gongbasicfieldDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongBasicField intance %s", gongbasicfield.Name))
		return err
	}

	return
}

// BackRepoGongBasicField.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) CheckoutPhaseOne() (Error error) {

	gongbasicfieldDBArray := make([]GongBasicFieldDB, 0)
	query := backRepoGongBasicField.db.Find(&gongbasicfieldDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongbasicfieldInstancesToBeRemovedFromTheStage := make(map[*models.GongBasicField]struct{})
	for key, value := range models.Stage.GongBasicFields {
		gongbasicfieldInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongbasicfieldDB := range gongbasicfieldDBArray {
		backRepoGongBasicField.CheckoutPhaseOneInstance(&gongbasicfieldDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongbasicfield, ok := (*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr)[gongbasicfieldDB.ID]
		if ok {
			delete(gongbasicfieldInstancesToBeRemovedFromTheStage, gongbasicfield)
		}
	}

	// remove from stage and back repo's 3 maps all gongbasicfields that are not in the checkout
	for gongbasicfield := range gongbasicfieldInstancesToBeRemovedFromTheStage {
		gongbasicfield.Unstage()

		// remove instance from the back repo 3 maps
		gongbasicfieldID := (*backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID)[gongbasicfield]
		delete((*backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID), gongbasicfield)
		delete((*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB), gongbasicfieldID)
		delete((*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr), gongbasicfieldID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongbasicfieldDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongbasicfieldDB
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) CheckoutPhaseOneInstance(gongbasicfieldDB *GongBasicFieldDB) (Error error) {

	gongbasicfield, ok := (*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr)[gongbasicfieldDB.ID]
	if !ok {
		gongbasicfield = new(models.GongBasicField)

		(*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr)[gongbasicfieldDB.ID] = gongbasicfield
		(*backRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID)[gongbasicfield] = gongbasicfieldDB.ID

		// append model store with the new element
		gongbasicfield.Name = gongbasicfieldDB.Name_Data.String
		gongbasicfield.Stage()
	}
	gongbasicfieldDB.CopyBasicFieldsToGongBasicField(gongbasicfield)

	// preserve pointer to gongbasicfieldDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongBasicFieldDBID_GongBasicFieldDB)[gongbasicfieldDB hold variable pointers
	gongbasicfieldDB_Data := *gongbasicfieldDB
	preservedPtrToGongBasicField := &gongbasicfieldDB_Data
	(*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB)[gongbasicfieldDB.ID] = preservedPtrToGongBasicField

	return
}

// BackRepoGongBasicField.CheckoutPhaseTwo Checkouts all staged instances of GongBasicField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongbasicfieldDB := range *backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB {
		backRepoGongBasicField.CheckoutPhaseTwoInstance(backRepo, gongbasicfieldDB)
	}
	return
}

// BackRepoGongBasicField.CheckoutPhaseTwoInstance Checkouts staged instances of GongBasicField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongbasicfieldDB *GongBasicFieldDB) (Error error) {

	gongbasicfield := (*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr)[gongbasicfieldDB.ID]
	_ = gongbasicfield // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// GongEnum field
	if gongbasicfieldDB.GongEnumID.Int64 != 0 {
		gongbasicfield.GongEnum = (*backRepo.BackRepoGongEnum.Map_GongEnumDBID_GongEnumPtr)[uint(gongbasicfieldDB.GongEnumID.Int64)]
	}
	return
}

// CommitGongBasicField allows commit of a single gongbasicfield (if already staged)
func (backRepo *BackRepoStruct) CommitGongBasicField(gongbasicfield *models.GongBasicField) {
	backRepo.BackRepoGongBasicField.CommitPhaseOneInstance(gongbasicfield)
	if id, ok := (*backRepo.BackRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID)[gongbasicfield]; ok {
		backRepo.BackRepoGongBasicField.CommitPhaseTwoInstance(backRepo, id, gongbasicfield)
	}
}

// CommitGongBasicField allows checkout of a single gongbasicfield (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongBasicField(gongbasicfield *models.GongBasicField) {
	// check if the gongbasicfield is staged
	if _, ok := (*backRepo.BackRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID)[gongbasicfield]; ok {

		if id, ok := (*backRepo.BackRepoGongBasicField.Map_GongBasicFieldPtr_GongBasicFieldDBID)[gongbasicfield]; ok {
			var gongbasicfieldDB GongBasicFieldDB
			gongbasicfieldDB.ID = id

			if err := backRepo.BackRepoGongBasicField.db.First(&gongbasicfieldDB, id).Error; err != nil {
				log.Panicln("CheckoutGongBasicField : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongBasicField.CheckoutPhaseOneInstance(&gongbasicfieldDB)
			backRepo.BackRepoGongBasicField.CheckoutPhaseTwoInstance(backRepo, &gongbasicfieldDB)
		}
	}
}

// CopyBasicFieldsFromGongBasicField
func (gongbasicfieldDB *GongBasicFieldDB) CopyBasicFieldsFromGongBasicField(gongbasicfield *models.GongBasicField) {
	// insertion point for fields commit

	gongbasicfieldDB.Name_Data.String = gongbasicfield.Name
	gongbasicfieldDB.Name_Data.Valid = true

	gongbasicfieldDB.BasicKindName_Data.String = gongbasicfield.BasicKindName
	gongbasicfieldDB.BasicKindName_Data.Valid = true

	gongbasicfieldDB.DeclaredType_Data.String = gongbasicfield.DeclaredType
	gongbasicfieldDB.DeclaredType_Data.Valid = true

	gongbasicfieldDB.Index_Data.Int64 = int64(gongbasicfield.Index)
	gongbasicfieldDB.Index_Data.Valid = true
}

// CopyBasicFieldsFromGongBasicFieldWOP
func (gongbasicfieldDB *GongBasicFieldDB) CopyBasicFieldsFromGongBasicFieldWOP(gongbasicfield *GongBasicFieldWOP) {
	// insertion point for fields commit

	gongbasicfieldDB.Name_Data.String = gongbasicfield.Name
	gongbasicfieldDB.Name_Data.Valid = true

	gongbasicfieldDB.BasicKindName_Data.String = gongbasicfield.BasicKindName
	gongbasicfieldDB.BasicKindName_Data.Valid = true

	gongbasicfieldDB.DeclaredType_Data.String = gongbasicfield.DeclaredType
	gongbasicfieldDB.DeclaredType_Data.Valid = true

	gongbasicfieldDB.Index_Data.Int64 = int64(gongbasicfield.Index)
	gongbasicfieldDB.Index_Data.Valid = true
}

// CopyBasicFieldsToGongBasicField
func (gongbasicfieldDB *GongBasicFieldDB) CopyBasicFieldsToGongBasicField(gongbasicfield *models.GongBasicField) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongbasicfield.Name = gongbasicfieldDB.Name_Data.String
	gongbasicfield.BasicKindName = gongbasicfieldDB.BasicKindName_Data.String
	gongbasicfield.DeclaredType = gongbasicfieldDB.DeclaredType_Data.String
	gongbasicfield.Index = int(gongbasicfieldDB.Index_Data.Int64)
}

// CopyBasicFieldsToGongBasicFieldWOP
func (gongbasicfieldDB *GongBasicFieldDB) CopyBasicFieldsToGongBasicFieldWOP(gongbasicfield *GongBasicFieldWOP) {
	gongbasicfield.ID = int(gongbasicfieldDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongbasicfield.Name = gongbasicfieldDB.Name_Data.String
	gongbasicfield.BasicKindName = gongbasicfieldDB.BasicKindName_Data.String
	gongbasicfield.DeclaredType = gongbasicfieldDB.DeclaredType_Data.String
	gongbasicfield.Index = int(gongbasicfieldDB.Index_Data.Int64)
}

// Backup generates a json file from a slice of all GongBasicFieldDB instances in the backrepo
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongBasicFieldDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongBasicFieldDB, 0)
	for _, gongbasicfieldDB := range *backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB {
		forBackup = append(forBackup, gongbasicfieldDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json GongBasicField ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json GongBasicField file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongBasicFieldDB instances in the backrepo
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongBasicFieldDB, 0)
	for _, gongbasicfieldDB := range *backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB {
		forBackup = append(forBackup, gongbasicfieldDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongBasicField")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongBasicField_Fields, -1)
	for _, gongbasicfieldDB := range forBackup {

		var gongbasicfieldWOP GongBasicFieldWOP
		gongbasicfieldDB.CopyBasicFieldsToGongBasicFieldWOP(&gongbasicfieldWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongbasicfieldWOP, -1)
	}
}

// RestoreXL from the "GongBasicField" sheet all GongBasicFieldDB instances
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongBasicFieldid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongBasicField"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongBasicField.rowVisitorGongBasicField)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) rowVisitorGongBasicField(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongbasicfieldWOP GongBasicFieldWOP
		row.ReadStruct(&gongbasicfieldWOP)

		// add the unmarshalled struct to the stage
		gongbasicfieldDB := new(GongBasicFieldDB)
		gongbasicfieldDB.CopyBasicFieldsFromGongBasicFieldWOP(&gongbasicfieldWOP)

		gongbasicfieldDB_ID_atBackupTime := gongbasicfieldDB.ID
		gongbasicfieldDB.ID = 0
		query := backRepoGongBasicField.db.Create(gongbasicfieldDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB)[gongbasicfieldDB.ID] = gongbasicfieldDB
		BackRepoGongBasicFieldid_atBckpTime_newID[gongbasicfieldDB_ID_atBackupTime] = gongbasicfieldDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongBasicFieldDB.json" in dirPath that stores an array
// of GongBasicFieldDB and stores it in the database
// the map BackRepoGongBasicFieldid_atBckpTime_newID is updated accordingly
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongBasicFieldid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongBasicFieldDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json GongBasicField file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongBasicFieldDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongBasicFieldDBID_GongBasicFieldDB
	for _, gongbasicfieldDB := range forRestore {

		gongbasicfieldDB_ID_atBackupTime := gongbasicfieldDB.ID
		gongbasicfieldDB.ID = 0
		query := backRepoGongBasicField.db.Create(gongbasicfieldDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB)[gongbasicfieldDB.ID] = gongbasicfieldDB
		BackRepoGongBasicFieldid_atBckpTime_newID[gongbasicfieldDB_ID_atBackupTime] = gongbasicfieldDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json GongBasicField file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongBasicField>id_atBckpTime_newID
// to compute new index
func (backRepoGongBasicField *BackRepoGongBasicFieldStruct) RestorePhaseTwo() {

	for _, gongbasicfieldDB := range *backRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB {

		// next line of code is to avert unused variable compilation error
		_ = gongbasicfieldDB

		// insertion point for reindexing pointers encoding
		// reindexing GongEnum field
		if gongbasicfieldDB.GongEnumID.Int64 != 0 {
			gongbasicfieldDB.GongEnumID.Int64 = int64(BackRepoGongEnumid_atBckpTime_newID[uint(gongbasicfieldDB.GongEnumID.Int64)])
			gongbasicfieldDB.GongEnumID.Valid = true
		}

		// This reindex gongbasicfield.GongBasicFields
		if gongbasicfieldDB.GongStruct_GongBasicFieldsDBID.Int64 != 0 {
			gongbasicfieldDB.GongStruct_GongBasicFieldsDBID.Int64 =
				int64(BackRepoGongStructid_atBckpTime_newID[uint(gongbasicfieldDB.GongStruct_GongBasicFieldsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoGongBasicField.db.Model(gongbasicfieldDB).Updates(*gongbasicfieldDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongBasicFieldid_atBckpTime_newID map[uint]uint
