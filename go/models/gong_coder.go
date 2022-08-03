package models

import "time"

// GongfieldCoder return an instance of Type where each field 
// encodes the index of the field
//
// This allows for refactorable field names
// 
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case Machine:
		fieldCoder := Machine{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.DrumLoad = 1.000000
		fieldCoder.RemainingTime = 2
		fieldCoder.Cleanedlaundry = false
		fieldCoder.State = "4"
		return (any)(fieldCoder).(Type)
	case Simulation:
		fieldCoder := Simulation{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.LastCommitNb = 3
		return (any)(fieldCoder).(Type)
	case Washer:
		fieldCoder := Washer{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.DirtyLaundryWeight = 1.000000
		fieldCoder.State = "2"
		fieldCoder.CleanedLaundryWeight = 3.000000
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *Machine | []*Machine | *Simulation | []*Simulation | *Washer | []*Washer
}

// GongfieldName provides the name of the field by passing the instance of the coder to
// the fonction.
//
// This allows for refactorable field name
//
// fieldCoder := models.GongfieldCoder[models.Astruct]()
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Name))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Booleanfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Intfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Floatfield))
// 
// limitations:
// 1. cannot encode boolean fields
// 2. for associations (pointer to gongstruct or slice of pointer to gongstruct, uses GetAssociationName)
func GongfieldName[Type PointerToGongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case *Machine:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "4" {
				return "State"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 2 {
				return "RemainingTime"
			}
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "DrumLoad"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "Cleanedlaundry"
			}
		}
	case *Simulation:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 3 {
				return "LastCommitNb"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Washer:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "2" {
				return "State"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "DirtyLaundryWeight"
			}
			if field == 3.000000 {
				return "CleanedLaundryWeight"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	default:
		return ""
	}
	_ = field
	return ""
}
