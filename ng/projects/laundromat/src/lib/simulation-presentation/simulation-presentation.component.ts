import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { SimulationDB } from '../simulation-db'
import { SimulationService } from '../simulation.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface simulationDummyElement {
}

const ELEMENT_DATA: simulationDummyElement[] = [
];

@Component({
	selector: 'app-simulation-presentation',
	templateUrl: './simulation-presentation.component.html',
	styleUrls: ['./simulation-presentation.component.css'],
})
export class SimulationPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	simulation: SimulationDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private simulationService: SimulationService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getSimulation();

		// observable for changes in 
		this.simulationService.SimulationServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getSimulation()
				}
			}
		)
	}

	getSimulation(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.simulation = this.frontRepo.Simulations.get(id)

				// insertion point for recovery of durations
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
				github_com_fullstack_lang_laundromat_go_editor: ["github_com_fullstack_lang_laundromat_go-" + "simulation-detail", ID]
			}
		}]);
	}
}
