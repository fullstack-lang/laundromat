import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { MachinesTableComponent } from './machines-table/machines-table.component'
import { MachineDetailComponent } from './machine-detail/machine-detail.component'
import { MachinePresentationComponent } from './machine-presentation/machine-presentation.component'

import { SimulationsTableComponent } from './simulations-table/simulations-table.component'
import { SimulationDetailComponent } from './simulation-detail/simulation-detail.component'
import { SimulationPresentationComponent } from './simulation-presentation/simulation-presentation.component'

import { WashersTableComponent } from './washers-table/washers-table.component'
import { WasherDetailComponent } from './washer-detail/washer-detail.component'
import { WasherPresentationComponent } from './washer-presentation/washer-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_laundromat_go-machines', component: MachinesTableComponent, outlet: 'github_com_fullstack_lang_laundromat_go_table' },
	{ path: 'github_com_fullstack_lang_laundromat_go-machine-adder', component: MachineDetailComponent, outlet: 'github_com_fullstack_lang_laundromat_go_editor' },
	{ path: 'github_com_fullstack_lang_laundromat_go-machine-adder/:id/:association', component: MachineDetailComponent, outlet: 'github_com_fullstack_lang_laundromat_go_editor' },
	{ path: 'github_com_fullstack_lang_laundromat_go-machine-detail/:id', component: MachineDetailComponent, outlet: 'github_com_fullstack_lang_laundromat_go_editor' },
	{ path: 'github_com_fullstack_lang_laundromat_go-machine-presentation/:id', component: MachinePresentationComponent, outlet: 'github_com_fullstack_lang_laundromat_go_presentation' },
	{ path: 'github_com_fullstack_lang_laundromat_go-machine-presentation-special/:id', component: MachinePresentationComponent, outlet: 'github_com_fullstack_lang_laundromat_gomachinepres' },

	{ path: 'github_com_fullstack_lang_laundromat_go-simulations', component: SimulationsTableComponent, outlet: 'github_com_fullstack_lang_laundromat_go_table' },
	{ path: 'github_com_fullstack_lang_laundromat_go-simulation-adder', component: SimulationDetailComponent, outlet: 'github_com_fullstack_lang_laundromat_go_editor' },
	{ path: 'github_com_fullstack_lang_laundromat_go-simulation-adder/:id/:association', component: SimulationDetailComponent, outlet: 'github_com_fullstack_lang_laundromat_go_editor' },
	{ path: 'github_com_fullstack_lang_laundromat_go-simulation-detail/:id', component: SimulationDetailComponent, outlet: 'github_com_fullstack_lang_laundromat_go_editor' },
	{ path: 'github_com_fullstack_lang_laundromat_go-simulation-presentation/:id', component: SimulationPresentationComponent, outlet: 'github_com_fullstack_lang_laundromat_go_presentation' },
	{ path: 'github_com_fullstack_lang_laundromat_go-simulation-presentation-special/:id', component: SimulationPresentationComponent, outlet: 'github_com_fullstack_lang_laundromat_gosimulationpres' },

	{ path: 'github_com_fullstack_lang_laundromat_go-washers', component: WashersTableComponent, outlet: 'github_com_fullstack_lang_laundromat_go_table' },
	{ path: 'github_com_fullstack_lang_laundromat_go-washer-adder', component: WasherDetailComponent, outlet: 'github_com_fullstack_lang_laundromat_go_editor' },
	{ path: 'github_com_fullstack_lang_laundromat_go-washer-adder/:id/:association', component: WasherDetailComponent, outlet: 'github_com_fullstack_lang_laundromat_go_editor' },
	{ path: 'github_com_fullstack_lang_laundromat_go-washer-detail/:id', component: WasherDetailComponent, outlet: 'github_com_fullstack_lang_laundromat_go_editor' },
	{ path: 'github_com_fullstack_lang_laundromat_go-washer-presentation/:id', component: WasherPresentationComponent, outlet: 'github_com_fullstack_lang_laundromat_go_presentation' },
	{ path: 'github_com_fullstack_lang_laundromat_go-washer-presentation-special/:id', component: WasherPresentationComponent, outlet: 'github_com_fullstack_lang_laundromat_gowasherpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
