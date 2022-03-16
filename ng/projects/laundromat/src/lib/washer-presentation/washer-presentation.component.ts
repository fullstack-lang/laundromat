import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { WasherDB } from '../washer-db'
import { WasherService } from '../washer.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface washerDummyElement {
}

const ELEMENT_DATA: washerDummyElement[] = [
];

@Component({
	selector: 'app-washer-presentation',
	templateUrl: './washer-presentation.component.html',
	styleUrls: ['./washer-presentation.component.css'],
})
export class WasherPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	washer: WasherDB = new (WasherDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private washerService: WasherService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getWasher();

		// observable for changes in 
		this.washerService.WasherServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getWasher()
				}
			}
		)
	}

	getWasher(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.washer = this.frontRepo.Washers.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
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
				github_com_fullstack_lang_laundromat_go_editor: ["github_com_fullstack_lang_laundromat_go-" + "washer-detail", ID]
			}
		}]);
	}
}
