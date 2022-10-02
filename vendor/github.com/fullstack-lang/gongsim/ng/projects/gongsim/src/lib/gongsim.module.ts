import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { DummyAgentsTableComponent } from './dummyagents-table/dummyagents-table.component'
import { DummyAgentSortingComponent } from './dummyagent-sorting/dummyagent-sorting.component'
import { DummyAgentDetailComponent } from './dummyagent-detail/dummyagent-detail.component'
import { DummyAgentPresentationComponent } from './dummyagent-presentation/dummyagent-presentation.component'

import { EnginesTableComponent } from './engines-table/engines-table.component'
import { EngineSortingComponent } from './engine-sorting/engine-sorting.component'
import { EngineDetailComponent } from './engine-detail/engine-detail.component'
import { EnginePresentationComponent } from './engine-presentation/engine-presentation.component'

import { EventsTableComponent } from './events-table/events-table.component'
import { EventSortingComponent } from './event-sorting/event-sorting.component'
import { EventDetailComponent } from './event-detail/event-detail.component'
import { EventPresentationComponent } from './event-presentation/event-presentation.component'

import { GongsimCommandsTableComponent } from './gongsimcommands-table/gongsimcommands-table.component'
import { GongsimCommandSortingComponent } from './gongsimcommand-sorting/gongsimcommand-sorting.component'
import { GongsimCommandDetailComponent } from './gongsimcommand-detail/gongsimcommand-detail.component'
import { GongsimCommandPresentationComponent } from './gongsimcommand-presentation/gongsimcommand-presentation.component'

import { GongsimStatussTableComponent } from './gongsimstatuss-table/gongsimstatuss-table.component'
import { GongsimStatusSortingComponent } from './gongsimstatus-sorting/gongsimstatus-sorting.component'
import { GongsimStatusDetailComponent } from './gongsimstatus-detail/gongsimstatus-detail.component'
import { GongsimStatusPresentationComponent } from './gongsimstatus-presentation/gongsimstatus-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		DummyAgentsTableComponent,
		DummyAgentSortingComponent,
		DummyAgentDetailComponent,
		DummyAgentPresentationComponent,

		EnginesTableComponent,
		EngineSortingComponent,
		EngineDetailComponent,
		EnginePresentationComponent,

		EventsTableComponent,
		EventSortingComponent,
		EventDetailComponent,
		EventPresentationComponent,

		GongsimCommandsTableComponent,
		GongsimCommandSortingComponent,
		GongsimCommandDetailComponent,
		GongsimCommandPresentationComponent,

		GongsimStatussTableComponent,
		GongsimStatusSortingComponent,
		GongsimStatusDetailComponent,
		GongsimStatusPresentationComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		DummyAgentsTableComponent,
		DummyAgentSortingComponent,
		DummyAgentDetailComponent,
		DummyAgentPresentationComponent,

		EnginesTableComponent,
		EngineSortingComponent,
		EngineDetailComponent,
		EnginePresentationComponent,

		EventsTableComponent,
		EventSortingComponent,
		EventDetailComponent,
		EventPresentationComponent,

		GongsimCommandsTableComponent,
		GongsimCommandSortingComponent,
		GongsimCommandDetailComponent,
		GongsimCommandPresentationComponent,

		GongsimStatussTableComponent,
		GongsimStatusSortingComponent,
		GongsimStatusDetailComponent,
		GongsimStatusPresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongsimModule { }
