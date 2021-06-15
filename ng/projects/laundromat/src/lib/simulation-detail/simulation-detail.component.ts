// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { SimulationDB } from '../simulation-db'
import { SimulationService } from '../simulation.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-simulation-detail',
	templateUrl: './simulation-detail.component.html',
	styleUrls: ['./simulation-detail.component.css'],
})
export class SimulationDetailComponent implements OnInit {

	// insertion point for declarations

	// the SimulationDB of interest
	simulation: SimulationDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private simulationService: SimulationService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getSimulation()

		// observable for changes in structs
		this.simulationService.SimulationServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getSimulation()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getSimulation(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.simulation = frontRepo.Simulations.get(id)
				} else {
					this.simulation = new (SimulationDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization
		
		// insertion point for translation/nullation of each field
		if (this.simulation.MachineID == undefined) {
			this.simulation.MachineID = new NullInt64
		}
		if (this.simulation.Machine != undefined) {
			this.simulation.MachineID.Int64 = this.simulation.Machine.ID
			this.simulation.MachineID.Valid = true
		} else {
			this.simulation.MachineID.Int64 = 0
			this.simulation.MachineID.Valid = true
		}
		if (this.simulation.WasherID == undefined) {
			this.simulation.WasherID = new NullInt64
		}
		if (this.simulation.Washer != undefined) {
			this.simulation.WasherID.Int64 = this.simulation.Washer.ID
			this.simulation.WasherID.Valid = true
		} else {
			this.simulation.WasherID.Int64 = 0
			this.simulation.WasherID.Valid = true
		}
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
		}

		if (id != 0 && association == undefined) {

			this.simulationService.updateSimulation(this.simulation)
				.subscribe(simulation => {
					this.simulationService.SimulationServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.simulationService.postSimulation(this.simulation).subscribe(simulation => {

				this.simulationService.SimulationServiceChanged.next("post")

				this.simulation = {} // reset fields
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
			ID: this.simulation.ID,
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
			ID: this.simulation.ID,
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
		if (this.simulation.Name == undefined) {
			this.simulation.Name = event.value.Name		
		}
	}
}
