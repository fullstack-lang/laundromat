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
import { MachinesTableComponent } from './machines-table/machines-table.component'
import { MachineSortingComponent } from './machine-sorting/machine-sorting.component'
import { MachineDetailComponent } from './machine-detail/machine-detail.component'
import { MachinePresentationComponent } from './machine-presentation/machine-presentation.component'

import { SimulationsTableComponent } from './simulations-table/simulations-table.component'
import { SimulationSortingComponent } from './simulation-sorting/simulation-sorting.component'
import { SimulationDetailComponent } from './simulation-detail/simulation-detail.component'
import { SimulationPresentationComponent } from './simulation-presentation/simulation-presentation.component'

import { WashersTableComponent } from './washers-table/washers-table.component'
import { WasherSortingComponent } from './washer-sorting/washer-sorting.component'
import { WasherDetailComponent } from './washer-detail/washer-detail.component'
import { WasherPresentationComponent } from './washer-presentation/washer-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		MachinesTableComponent,
		MachineSortingComponent,
		MachineDetailComponent,
		MachinePresentationComponent,

		SimulationsTableComponent,
		SimulationSortingComponent,
		SimulationDetailComponent,
		SimulationPresentationComponent,

		WashersTableComponent,
		WasherSortingComponent,
		WasherDetailComponent,
		WasherPresentationComponent,


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
		MachinesTableComponent,
		MachineSortingComponent,
		MachineDetailComponent,
		MachinePresentationComponent,

		SimulationsTableComponent,
		SimulationSortingComponent,
		SimulationDetailComponent,
		SimulationPresentationComponent,

		WashersTableComponent,
		WasherSortingComponent,
		WasherDetailComponent,
		WasherPresentationComponent,


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
export class LaundromatModule { }
