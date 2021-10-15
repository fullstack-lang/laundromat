import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface machineDummyElement {
}

const ELEMENT_DATA: machineDummyElement[] = [
];

@Component({
	selector: 'app-machine-presentation',
	templateUrl: './machine-presentation.component.html',
	styleUrls: ['./machine-presentation.component.css'],
})
export class MachinePresentationComponent implements OnInit {

	// insertion point for declarations
	// fields from RemainingTime
	RemainingTime_Hours: number = 0
	RemainingTime_Minutes: number = 0
	RemainingTime_Seconds: number = 0

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	machine: MachineDB = new (MachineDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private machineService: MachineService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getMachine();

		// observable for changes in 
		this.machineService.MachineServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getMachine()
				}
			}
		)
	}

	getMachine(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.machine = this.frontRepo.Machines.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for RemainingTime
				this.RemainingTime_Hours = Math.floor(this.machine.RemainingTime / (3600 * 1000 * 1000 * 1000))
				this.RemainingTime_Minutes = Math.floor(this.machine.RemainingTime % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.RemainingTime_Seconds = this.machine.RemainingTime % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_laundromat_go_presentation: ["github_com_fullstack_lang_laundromat_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_laundromat_go_editor: ["github_com_fullstack_lang_laundromat_go-" + "machine-detail", ID]
			}
		}]);
	}
}
