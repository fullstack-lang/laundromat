import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { DummyAgentDB } from '../dummyagent-db'
import { DummyAgentService } from '../dummyagent.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface dummyagentDummyElement {
}

const ELEMENT_DATA: dummyagentDummyElement[] = [
];

@Component({
	selector: 'app-dummyagent-presentation',
	templateUrl: './dummyagent-presentation.component.html',
	styleUrls: ['./dummyagent-presentation.component.css'],
})
export class DummyAgentPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	dummyagent: DummyAgentDB = new (DummyAgentDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private dummyagentService: DummyAgentService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getDummyAgent();

		// observable for changes in 
		this.dummyagentService.DummyAgentServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getDummyAgent()
				}
			}
		)
	}

	getDummyAgent(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.dummyagent = this.frontRepo.DummyAgents.get(id)!

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
				github_com_fullstack_lang_gongsim_go_editor: ["github_com_fullstack_lang_gongsim_go-" + "dummyagent-detail", ID]
			}
		}]);
	}
}
