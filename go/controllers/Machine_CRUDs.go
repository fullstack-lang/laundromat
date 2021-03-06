// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/laundromat/go/models"
	"github.com/fullstack-lang/laundromat/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Machine__dummysDeclaration__ models.Machine
var __Machine_time__dummyDeclaration time.Duration

// An MachineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMachine updateMachine deleteMachine
type MachineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MachineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMachine updateMachine
type MachineInput struct {
	// The Machine to submit or modify
	// in: body
	Machine *orm.MachineAPI
}

// GetMachines
//
// swagger:route GET /machines machines getMachines
//
// Get all machines
//
// Responses:
//    default: genericError
//        200: machineDBsResponse
func GetMachines(c *gin.Context) {
	db := orm.BackRepo.BackRepoMachine.GetDB()

	// source slice
	var machineDBs []orm.MachineDB
	query := db.Find(&machineDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	machineAPIs := make([]orm.MachineAPI, 0)

	// for each machine, update fields from the database nullable fields
	for idx := range machineDBs {
		machineDB := &machineDBs[idx]
		_ = machineDB
		var machineAPI orm.MachineAPI

		// insertion point for updating fields
		machineAPI.ID = machineDB.ID
		machineDB.CopyBasicFieldsToMachine(&machineAPI.Machine)
		machineAPI.MachinePointersEnconding = machineDB.MachinePointersEnconding
		machineAPIs = append(machineAPIs, machineAPI)
	}

	c.JSON(http.StatusOK, machineAPIs)
}

// PostMachine
//
// swagger:route POST /machines machines postMachine
//
// Creates a machine
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: machineDBResponse
func PostMachine(c *gin.Context) {
	db := orm.BackRepo.BackRepoMachine.GetDB()

	// Validate input
	var input orm.MachineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create machine
	machineDB := orm.MachineDB{}
	machineDB.MachinePointersEnconding = input.MachinePointersEnconding
	machineDB.CopyBasicFieldsFromMachine(&input.Machine)

	query := db.Create(&machineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, machineDB)
}

// GetMachine
//
// swagger:route GET /machines/{ID} machines getMachine
//
// Gets the details for a machine.
//
// Responses:
//    default: genericError
//        200: machineDBResponse
func GetMachine(c *gin.Context) {
	db := orm.BackRepo.BackRepoMachine.GetDB()

	// Get machineDB in DB
	var machineDB orm.MachineDB
	if err := db.First(&machineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var machineAPI orm.MachineAPI
	machineAPI.ID = machineDB.ID
	machineAPI.MachinePointersEnconding = machineDB.MachinePointersEnconding
	machineDB.CopyBasicFieldsToMachine(&machineAPI.Machine)

	c.JSON(http.StatusOK, machineAPI)
}

// UpdateMachine
//
// swagger:route PATCH /machines/{ID} machines updateMachine
//
// Update a machine
//
// Responses:
//    default: genericError
//        200: machineDBResponse
func UpdateMachine(c *gin.Context) {
	db := orm.BackRepo.BackRepoMachine.GetDB()

	// Get model if exist
	var machineDB orm.MachineDB

	// fetch the machine
	query := db.First(&machineDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.MachineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	machineDB.CopyBasicFieldsFromMachine(&input.Machine)
	machineDB.MachinePointersEnconding = input.MachinePointersEnconding

	query = db.Model(&machineDB).Updates(machineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the machineDB
	c.JSON(http.StatusOK, machineDB)
}

// DeleteMachine
//
// swagger:route DELETE /machines/{ID} machines deleteMachine
//
// Delete a machine
//
// Responses:
//    default: genericError
func DeleteMachine(c *gin.Context) {
	db := orm.BackRepo.BackRepoMachine.GetDB()

	// Get model if exist
	var machineDB orm.MachineDB
	if err := db.First(&machineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&machineDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
