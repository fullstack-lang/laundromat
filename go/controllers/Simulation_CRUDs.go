// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"net/http"
	"time"

	"github.com/fullstack-lang/laundromat/go/models"
	"github.com/fullstack-lang/laundromat/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Simulation__dummysDeclaration__ models.Simulation
var __Simulation_time__dummyDeclaration time.Duration

// An SimulationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSimulation updateSimulation deleteSimulation
type SimulationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SimulationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSimulation updateSimulation
type SimulationInput struct {
	// The Simulation to submit or modify
	// in: body
	Simulation *orm.SimulationAPI
}

// GetSimulations
//
// swagger:route GET /simulations simulations getSimulations
//
// Get all simulations
//
// Responses:
//    default: genericError
//        200: simulationDBsResponse
func GetSimulations(c *gin.Context) {
	db := orm.BackRepo.BackRepoSimulation.GetDB()
	
	// source slice
	var simulationDBs []orm.SimulationDB
	query := db.Find(&simulationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	simulationAPIs := make([]orm.SimulationAPI, 0)

	// for each simulation, update fields from the database nullable fields
	for idx := range simulationDBs {
		simulationDB := &simulationDBs[idx]
		_ = simulationDB
		var simulationAPI orm.SimulationAPI

		// insertion point for updating fields
		simulationAPI.ID = simulationDB.ID
		simulationDB.CopyBasicFieldsToSimulation(&simulationAPI.Simulation)
		simulationAPI.SimulationPointersEnconding = simulationDB.SimulationPointersEnconding
		simulationAPIs = append(simulationAPIs, simulationAPI)
	}

	c.JSON(http.StatusOK, simulationAPIs)
}

// PostSimulation
//
// swagger:route POST /simulations simulations postSimulation
//
// Creates a simulation
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: simulationDBResponse
func PostSimulation(c *gin.Context) {
	db := orm.BackRepo.BackRepoSimulation.GetDB()

	// Validate input
	var input orm.SimulationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create simulation
	simulationDB := orm.SimulationDB{}
	simulationDB.SimulationPointersEnconding = input.SimulationPointersEnconding
	simulationDB.CopyBasicFieldsFromSimulation(&input.Simulation)

	query := db.Create(&simulationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	c.JSON(http.StatusOK, simulationDB)
}

// GetSimulation
//
// swagger:route GET /simulations/{ID} simulations getSimulation
//
// Gets the details for a simulation.
//
// Responses:
//    default: genericError
//        200: simulationDBResponse
func GetSimulation(c *gin.Context) {
	db := orm.BackRepo.BackRepoSimulation.GetDB()

	// Get simulationDB in DB
	var simulationDB orm.SimulationDB
	if err := db.First(&simulationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var simulationAPI orm.SimulationAPI
	simulationAPI.ID = simulationDB.ID
	simulationAPI.SimulationPointersEnconding = simulationDB.SimulationPointersEnconding
	simulationDB.CopyBasicFieldsToSimulation(&simulationAPI.Simulation)

	c.JSON(http.StatusOK, simulationAPI)
}

// UpdateSimulation
//
// swagger:route PATCH /simulations/{ID} simulations updateSimulation
//
// Update a simulation
//
// Responses:
//    default: genericError
//        200: simulationDBResponse
func UpdateSimulation(c *gin.Context) {
	db := orm.BackRepo.BackRepoSimulation.GetDB()

	// Get model if exist
	var simulationDB orm.SimulationDB

	// fetch the simulation
	query := db.First(&simulationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.SimulationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	simulationDB.CopyBasicFieldsFromSimulation(&input.Simulation)
	simulationDB.SimulationPointersEnconding = input.SimulationPointersEnconding

	query = db.Model(&simulationDB).Updates(simulationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	// return status OK with the marshalling of the the simulationDB
	c.JSON(http.StatusOK, simulationDB)
}

// DeleteSimulation
//
// swagger:route DELETE /simulations/{ID} simulations deleteSimulation
//
// Delete a simulation
//
// Responses:
//    default: genericError
func DeleteSimulation(c *gin.Context) {
	db := orm.BackRepo.BackRepoSimulation.GetDB()

	// Get model if exist
	var simulationDB orm.SimulationDB
	if err := db.First(&simulationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&simulationDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementCommitNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
