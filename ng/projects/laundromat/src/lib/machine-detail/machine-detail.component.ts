// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { MachineStateEnumSelect, MachineStateEnumList } from '../MachineStateEnum'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

// MachineDetailComponent is initilizaed from different routes
// MachineDetailComponentState detail different cases 
enum MachineDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
}

@Component({
	selector: 'app-machine-detail',
	templateUrl: './machine-detail.component.html',
	styleUrls: ['./machine-detail.component.css'],
})
export class MachineDetailComponent implements OnInit {

	// insertion point for declarations
	RemainingTime_Hours: number
	RemainingTime_Minutes: number
	RemainingTime_Seconds: number
	CleanedlaundryFormControl = new FormControl(false);
	MachineStateEnumList: MachineStateEnumSelect[]

	// the MachineDB of interest
	machine: MachineDB;

	// front repo
	frontRepo: FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: MachineDetailComponentState

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string
	originStructFieldName: string

	constructor(
		private machineService: MachineService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id');
		this.originStruct = this.route.snapshot.paramMap.get('originStruct');
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName');

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = MachineDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = MachineDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getMachine()

		// observable for changes in structs
		this.machineService.MachineServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getMachine()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.MachineStateEnumList = MachineStateEnumList
	}

	getMachine(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case MachineDetailComponentState.CREATE_INSTANCE:
						this.machine = new (MachineDB)
						break;
					case MachineDetailComponentState.UPDATE_INSTANCE:
						this.machine = frontRepo.Machines.get(this.id)
						break;
					// insertion point for init of association field
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				// computation of Hours, Minutes, Seconds for RemainingTime
				this.RemainingTime_Hours = Math.floor(this.machine.RemainingTime / (3600 * 1000 * 1000 * 1000))
				this.RemainingTime_Minutes = Math.floor(this.machine.RemainingTime % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.RemainingTime_Seconds = this.machine.RemainingTime % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
				this.CleanedlaundryFormControl.setValue(this.machine.Cleanedlaundry)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.machine.RemainingTime =
			this.RemainingTime_Hours * (3600 * 1000 * 1000 * 1000) +
			this.RemainingTime_Minutes * (60 * 1000 * 1000 * 1000) +
			this.RemainingTime_Seconds * (1000 * 1000 * 1000)
		this.machine.Cleanedlaundry = this.CleanedlaundryFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers

		switch (this.state) {
			case MachineDetailComponentState.UPDATE_INSTANCE:
				this.machineService.updateMachine(this.machine)
					.subscribe(machine => {
						this.machineService.MachineServiceChanged.next("update")
					});
				break;
			default:
				this.machineService.postMachine(this.machine).subscribe(machine => {
					this.machineService.MachineServiceChanged.next("post")
					this.machine = {} // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string ) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.machine.ID
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.machine.ID
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Machine"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.machine.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event) {
		if (this.machine.Name == undefined) {
			this.machine.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
