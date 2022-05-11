import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongsimCommandDB } from '../gongsimcommand-db'
import { GongsimCommandService } from '../gongsimcommand.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongsimcommandDummyElement {
}

const ELEMENT_DATA: gongsimcommandDummyElement[] = [
];

@Component({
	selector: 'app-gongsimcommand-presentation',
	templateUrl: './gongsimcommand-presentation.component.html',
	styleUrls: ['./gongsimcommand-presentation.component.css'],
})
export class GongsimCommandPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongsimcommand: GongsimCommandDB = new (GongsimCommandDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongsimcommandService: GongsimCommandService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongsimCommand();

		// observable for changes in 
		this.gongsimcommandService.GongsimCommandServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongsimCommand()
				}
			}
		)
	}

	getGongsimCommand(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongsimcommand = this.frontRepo.GongsimCommands.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
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
				github_com_fullstack_lang_gongsim_go_editor: ["github_com_fullstack_lang_gongsim_go-" + "gongsimcommand-detail", ID]
			}
		}]);
	}
}
