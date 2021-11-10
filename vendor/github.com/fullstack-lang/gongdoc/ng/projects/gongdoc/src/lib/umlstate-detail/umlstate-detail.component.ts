// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { UmlStateDB } from '../umlstate-db'
import { UmlStateService } from '../umlstate.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { UmlscDB } from '../umlsc-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// UmlStateDetailComponent is initilizaed from different routes
// UmlStateDetailComponentState detail different cases 
enum UmlStateDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Umlsc_States_SET,
}

@Component({
	selector: 'app-umlstate-detail',
	templateUrl: './umlstate-detail.component.html',
	styleUrls: ['./umlstate-detail.component.css'],
})
export class UmlStateDetailComponent implements OnInit {

	// insertion point for declarations

	// the UmlStateDB of interest
	umlstate: UmlStateDB = new UmlStateDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: UmlStateDetailComponentState = UmlStateDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private umlstateService: UmlStateService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id')!;
		this.originStruct = this.route.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = UmlStateDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = UmlStateDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "States":
						// console.log("UmlState" + " is instanciated with back pointer to instance " + this.id + " Umlsc association States")
						this.state = UmlStateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Umlsc_States_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getUmlState()

		// observable for changes in structs
		this.umlstateService.UmlStateServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getUmlState()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getUmlState(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case UmlStateDetailComponentState.CREATE_INSTANCE:
						this.umlstate = new (UmlStateDB)
						break;
					case UmlStateDetailComponentState.UPDATE_INSTANCE:
						let umlstate = frontRepo.UmlStates.get(this.id)
						console.assert(umlstate != undefined, "missing umlstate with id:" + this.id)
						this.umlstate = umlstate!
						break;
					// insertion point for init of association field
					case UmlStateDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Umlsc_States_SET:
						this.umlstate = new (UmlStateDB)
						this.umlstate.Umlsc_States_reverse = frontRepo.Umlscs.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.umlstate.Umlsc_States_reverse != undefined) {
			if (this.umlstate.Umlsc_StatesDBID == undefined) {
				this.umlstate.Umlsc_StatesDBID = new NullInt64
			}
			this.umlstate.Umlsc_StatesDBID.Int64 = this.umlstate.Umlsc_States_reverse.ID
			this.umlstate.Umlsc_StatesDBID.Valid = true
			if (this.umlstate.Umlsc_StatesDBID_Index == undefined) {
				this.umlstate.Umlsc_StatesDBID_Index = new NullInt64
			}
			this.umlstate.Umlsc_StatesDBID_Index.Valid = true
			this.umlstate.Umlsc_States_reverse = new UmlscDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case UmlStateDetailComponentState.UPDATE_INSTANCE:
				this.umlstateService.updateUmlState(this.umlstate)
					.subscribe(umlstate => {
						this.umlstateService.UmlStateServiceChanged.next("update")
					});
				break;
			default:
				this.umlstateService.postUmlState(this.umlstate).subscribe(umlstate => {
					this.umlstateService.UmlStateServiceChanged.next("post")
					this.umlstate = new (UmlStateDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.umlstate.ID!
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
			dialogData.ID = this.umlstate.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "UmlState"
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
			ID: this.umlstate.ID,
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

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.umlstate.Name == undefined) {
			this.umlstate.Name = event.value.Name
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
			return this.mapFields_displayAsTextArea.get(fieldName)!
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
