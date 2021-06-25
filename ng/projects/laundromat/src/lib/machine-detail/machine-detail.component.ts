// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { MachineStateEnumSelect, MachineStateEnumList } from '../MachineStateEnum'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

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

	constructor(
		private machineService: MachineService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
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
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.machine = frontRepo.Machines.get(id)
				} else {
					this.machine = new (MachineDB)
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
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.machine.RemainingTime =
			this.RemainingTime_Hours * (3600 * 1000 * 1000 * 1000) +
			this.RemainingTime_Minutes * (60 * 1000 * 1000 * 1000) +
			this.RemainingTime_Seconds * (1000 * 1000 * 1000)
		this.machine.Cleanedlaundry = this.CleanedlaundryFormControl.value

		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
		}

		if (id != 0 && association == undefined) {

			this.machineService.updateMachine(this.machine)
				.subscribe(machine => {
					this.machineService.MachineServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.machineService.postMachine(this.machine).subscribe(machine => {

				this.machineService.MachineServiceChanged.next("post")

				this.machine = {} // reset fields
			});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		dialogConfig.data = {
			ID: this.machine.ID,
			ReversePointer: reverseField,
			OrderingMode: false,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
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
