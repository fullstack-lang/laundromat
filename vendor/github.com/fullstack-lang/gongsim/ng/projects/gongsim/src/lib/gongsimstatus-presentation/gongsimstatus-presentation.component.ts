import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongsimStatusDB } from '../gongsimstatus-db'
import { GongsimStatusService } from '../gongsimstatus.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface gongsimstatusDummyElement {
}

const ELEMENT_DATA: gongsimstatusDummyElement[] = [
];

@Component({
	selector: 'app-gongsimstatus-presentation',
	templateUrl: './gongsimstatus-presentation.component.html',
	styleUrls: ['./gongsimstatus-presentation.component.css'],
})
export class GongsimStatusPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongsimstatus: GongsimStatusDB = new (GongsimStatusDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongsimstatusService: GongsimStatusService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongsimStatus();

		// observable for changes in 
		this.gongsimstatusService.GongsimStatusServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongsimStatus()
				}
			}
		)
	}

	getGongsimStatus(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongsimstatus = this.frontRepo.GongsimStatuss.get(id)!

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
				github_com_fullstack_lang_gongsim_go_editor: ["github_com_fullstack_lang_gongsim_go-" + "gongsimstatus-detail", ID]
			}
		}]);
	}
}
