// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Classshape__dummysDeclaration__ models.Classshape
var __Classshape_time__dummyDeclaration time.Duration

// An ClassshapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getClassshape updateClassshape deleteClassshape
type ClassshapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ClassshapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postClassshape updateClassshape
type ClassshapeInput struct {
	// The Classshape to submit or modify
	// in: body
	Classshape *orm.ClassshapeAPI
}

// GetClassshapes
//
// swagger:route GET /classshapes classshapes getClassshapes
//
// Get all classshapes
//
// Responses:
//    default: genericError
//        200: classshapeDBsResponse
func GetClassshapes(c *gin.Context) {
	db := orm.BackRepo.BackRepoClassshape.GetDB()

	// source slice
	var classshapeDBs []orm.ClassshapeDB
	query := db.Find(&classshapeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	classshapeAPIs := make([]orm.ClassshapeAPI, 0)

	// for each classshape, update fields from the database nullable fields
	for idx := range classshapeDBs {
		classshapeDB := &classshapeDBs[idx]
		_ = classshapeDB
		var classshapeAPI orm.ClassshapeAPI

		// insertion point for updating fields
		classshapeAPI.ID = classshapeDB.ID
		classshapeDB.CopyBasicFieldsToClassshape(&classshapeAPI.Classshape)
		classshapeAPI.ClassshapePointersEnconding = classshapeDB.ClassshapePointersEnconding
		classshapeAPIs = append(classshapeAPIs, classshapeAPI)
	}

	c.JSON(http.StatusOK, classshapeAPIs)
}

// PostClassshape
//
// swagger:route POST /classshapes classshapes postClassshape
//
// Creates a classshape
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: classshapeDBResponse
func PostClassshape(c *gin.Context) {
	db := orm.BackRepo.BackRepoClassshape.GetDB()

	// Validate input
	var input orm.ClassshapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create classshape
	classshapeDB := orm.ClassshapeDB{}
	classshapeDB.ClassshapePointersEnconding = input.ClassshapePointersEnconding
	classshapeDB.CopyBasicFieldsFromClassshape(&input.Classshape)

	query := db.Create(&classshapeDB)
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

	c.JSON(http.StatusOK, classshapeDB)
}

// GetClassshape
//
// swagger:route GET /classshapes/{ID} classshapes getClassshape
//
// Gets the details for a classshape.
//
// Responses:
//    default: genericError
//        200: classshapeDBResponse
func GetClassshape(c *gin.Context) {
	db := orm.BackRepo.BackRepoClassshape.GetDB()

	// Get classshapeDB in DB
	var classshapeDB orm.ClassshapeDB
	if err := db.First(&classshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var classshapeAPI orm.ClassshapeAPI
	classshapeAPI.ID = classshapeDB.ID
	classshapeAPI.ClassshapePointersEnconding = classshapeDB.ClassshapePointersEnconding
	classshapeDB.CopyBasicFieldsToClassshape(&classshapeAPI.Classshape)

	c.JSON(http.StatusOK, classshapeAPI)
}

// UpdateClassshape
//
// swagger:route PATCH /classshapes/{ID} classshapes updateClassshape
//
// Update a classshape
//
// Responses:
//    default: genericError
//        200: classshapeDBResponse
func UpdateClassshape(c *gin.Context) {
	db := orm.BackRepo.BackRepoClassshape.GetDB()

	// Get model if exist
	var classshapeDB orm.ClassshapeDB

	// fetch the classshape
	query := db.First(&classshapeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.ClassshapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	classshapeDB.CopyBasicFieldsFromClassshape(&input.Classshape)
	classshapeDB.ClassshapePointersEnconding = input.ClassshapePointersEnconding

	query = db.Model(&classshapeDB).Updates(classshapeDB)
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

	// return status OK with the marshalling of the the classshapeDB
	c.JSON(http.StatusOK, classshapeDB)
}

// DeleteClassshape
//
// swagger:route DELETE /classshapes/{ID} classshapes deleteClassshape
//
// Delete a classshape
//
// Responses:
//    default: genericError
func DeleteClassshape(c *gin.Context) {
	db := orm.BackRepo.BackRepoClassshape.GetDB()

	// Get model if exist
	var classshapeDB orm.ClassshapeDB
	if err := db.First(&classshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&classshapeDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
