import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { EngineDB } from '../engine-db'
import { EngineService } from '../engine.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface engineDummyElement {
}

const ELEMENT_DATA: engineDummyElement[] = [
];

@Component({
	selector: 'app-engine-presentation',
	templateUrl: './engine-presentation.component.html',
	styleUrls: ['./engine-presentation.component.css'],
})
export class EnginePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	engine: EngineDB = new (EngineDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private engineService: EngineService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getEngine();

		// observable for changes in 
		this.engineService.EngineServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getEngine()
				}
			}
		)
	}

	getEngine(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.engine = this.frontRepo.Engines.get(id)!

				// insertion point for recovery of durations
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
				github_com_fullstack_lang_gongsim_go_editor: ["github_com_fullstack_lang_gongsim_go-" + "engine-detail", ID]
			}
		}]);
	}
}
