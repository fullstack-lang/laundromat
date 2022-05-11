import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { EventDB } from '../event-db'
import { EventService } from '../event.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface eventDummyElement {
}

const ELEMENT_DATA: eventDummyElement[] = [
];

@Component({
	selector: 'app-event-presentation',
	templateUrl: './event-presentation.component.html',
	styleUrls: ['./event-presentation.component.css'],
})
export class EventPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// fields from Duration
	Duration_Hours: number = 0
	Duration_Minutes: number = 0
	Duration_Seconds: number = 0
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	event: EventDB = new (EventDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private eventService: EventService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getEvent();

		// observable for changes in 
		this.eventService.EventServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getEvent()
				}
			}
		)
	}

	getEvent(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.event = this.frontRepo.Events.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for Duration
				this.Duration_Hours = Math.floor(this.event.Duration / (3600 * 1000 * 1000 * 1000))
				this.Duration_Minutes = Math.floor(this.event.Duration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration_Seconds = this.event.Duration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
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
				github_com_fullstack_lang_gongsim_go_editor: ["github_com_fullstack_lang_gongsim_go-" + "event-detail", ID]
			}
		}]);
	}
}
