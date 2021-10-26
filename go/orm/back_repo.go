// generated by genORMTranslation.go
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"github.com/fullstack-lang/laundromat/go/models"

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoMachine BackRepoMachineStruct

	BackRepoSimulation BackRepoSimulationStruct

	BackRepoWasher BackRepoWasherStruct

	CommitNb uint // this ng is updated at the BackRepo level but also at the BackRepo<GongStruct> level

	PushFromFrontNb uint // records increments from push from front
}

func (backRepo *BackRepoStruct) GetLastCommitNb() uint {
	return backRepo.CommitNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	backRepo.CommitNb = backRepo.CommitNb + 1
	return backRepo.CommitNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitNb
}

// Init the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) init(db *gorm.DB) {
	// insertion point for per struct back repo declarations
	backRepo.BackRepoMachine.Init(db)
	backRepo.BackRepoSimulation.Init(db)
	backRepo.BackRepoWasher.Init(db)

	models.Stage.BackRepo = backRepo
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoMachine.CommitPhaseOne(stage)
	backRepo.BackRepoSimulation.CommitPhaseOne(stage)
	backRepo.BackRepoWasher.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoMachine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSimulation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoWasher.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoMachine.CheckoutPhaseOne()
	backRepo.BackRepoSimulation.CheckoutPhaseOne()
	backRepo.BackRepoWasher.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoMachine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSimulation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoWasher.CheckoutPhaseTwo(backRepo)
}

var BackRepo BackRepoStruct

func GetLastCommitNb() uint {
	return BackRepo.GetLastCommitNb()
}

func GetLastPushFromFrontNb() uint {
	return BackRepo.GetLastPushFromFrontNb()
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoMachine.Backup(dirPath)
	backRepo.BackRepoSimulation.Backup(dirPath)
	backRepo.BackRepoWasher.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoMachine.BackupXL(file)
	backRepo.BackRepoSimulation.BackupXL(file)
	backRepo.BackRepoWasher.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	models.Stage.Commit()
	models.Stage.Reset()
	models.Stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoMachine.RestorePhaseOne(dirPath)
	backRepo.BackRepoSimulation.RestorePhaseOne(dirPath)
	backRepo.BackRepoWasher.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoMachine.RestorePhaseTwo()
	backRepo.BackRepoSimulation.RestorePhaseTwo()
	backRepo.BackRepoWasher.RestorePhaseTwo()

	models.Stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {
}
