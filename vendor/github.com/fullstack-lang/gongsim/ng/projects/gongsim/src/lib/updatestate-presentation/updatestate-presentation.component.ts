import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { UpdateStateDB } from '../updatestate-db'
import { UpdateStateService } from '../updatestate.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface updatestateDummyElement {
}

const ELEMENT_DATA: updatestateDummyElement[] = [
];

@Component({
	selector: 'app-updatestate-presentation',
	templateUrl: './updatestate-presentation.component.html',
	styleUrls: ['./updatestate-presentation.component.css'],
})
export class UpdateStatePresentationComponent implements OnInit {

	// insertion point for declarations
	// fields from Duration
	Duration_Hours: number = 0
	Duration_Minutes: number = 0
	Duration_Seconds: number = 0
	// fields from Period
	Period_Hours: number = 0
	Period_Minutes: number = 0
	Period_Seconds: number = 0

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	updatestate: UpdateStateDB = new (UpdateStateDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private updatestateService: UpdateStateService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getUpdateState();

		// observable for changes in 
		this.updatestateService.UpdateStateServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getUpdateState()
				}
			}
		)
	}

	getUpdateState(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.updatestate = this.frontRepo.UpdateStates.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for Duration
				this.Duration_Hours = Math.floor(this.updatestate.Duration / (3600 * 1000 * 1000 * 1000))
				this.Duration_Minutes = Math.floor(this.updatestate.Duration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration_Seconds = this.updatestate.Duration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
				// computation of Hours, Minutes, Seconds for Period
				this.Period_Hours = Math.floor(this.updatestate.Period / (3600 * 1000 * 1000 * 1000))
				this.Period_Minutes = Math.floor(this.updatestate.Period % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Period_Seconds = this.updatestate.Period % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongsim_go_presentation: ["github_com_fullstack_lang_gongsim_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongsim_go_editor: ["github_com_fullstack_lang_gongsim_go-" + "updatestate-detail", ID]
			}
		}]);
	}
}
