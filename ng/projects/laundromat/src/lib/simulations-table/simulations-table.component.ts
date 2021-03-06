// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { SimulationDB } from '../simulation-db'
import { SimulationService } from '../simulation.service'

// insertion point for additional imports

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-simulationstable',
  templateUrl: './simulations-table.component.html',
  styleUrls: ['./simulations-table.component.css'],
})
export class SimulationsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Simulation instances
  selection: SelectionModel<SimulationDB> = new (SelectionModel)
  initialSelection = new Array<SimulationDB>()

  // the data source for the table
  simulations: SimulationDB[] = []
  matTableDataSource: MatTableDataSource<SimulationDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.simulations
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (simulationDB: SimulationDB, property: string) => {
      switch (property) {
        case 'ID':
          return simulationDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return simulationDB.Name;

        case 'Machine':
          return (simulationDB.Machine ? simulationDB.Machine.Name : '');

        case 'Washer':
          return (simulationDB.Washer ? simulationDB.Washer.Name : '');

        case 'LastCommitNb':
          return simulationDB.LastCommitNb;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (simulationDB: SimulationDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the simulationDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += simulationDB.Name.toLowerCase()
      if (simulationDB.Machine) {
        mergedContent += simulationDB.Machine.Name.toLowerCase()
      }
      if (simulationDB.Washer) {
        mergedContent += simulationDB.Washer.Name.toLowerCase()
      }
      mergedContent += simulationDB.LastCommitNb.toString()

      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private simulationService: SimulationService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of simulation instances
    public dialogRef: MatDialogRef<SimulationsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.simulationService.SimulationServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getSimulations()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Machine",
        "Washer",
        "LastCommitNb",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Machine",
        "Washer",
        "LastCommitNb",
      ]
      this.selection = new SelectionModel<SimulationDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getSimulations()
    this.matTableDataSource = new MatTableDataSource(this.simulations)
  }

  getSimulations(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.simulations = this.frontRepo.Simulations_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let simulation of this.simulations) {
            let ID = this.dialogData.ID
            let revPointer = simulation[this.dialogData.ReversePointer as keyof SimulationDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(simulation)
            }
            this.selection = new SelectionModel<SimulationDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, SimulationDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as SimulationDB[]
          for (let associationInstance of sourceField) {
            let simulation = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as SimulationDB
            this.initialSelection.push(simulation)
          }

          this.selection = new SelectionModel<SimulationDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.simulations
      }
    )
  }

  // newSimulation initiate a new simulation
  // create a new Simulation objet
  newSimulation() {
  }

  deleteSimulation(simulationID: number, simulation: SimulationDB) {
    // list of simulations is truncated of simulation before the delete
    this.simulations = this.simulations.filter(h => h !== simulation);

    this.simulationService.deleteSimulation(simulationID).subscribe(
      simulation => {
        this.simulationService.SimulationServiceChanged.next("delete")
      }
    );
  }

  editSimulation(simulationID: number, simulation: SimulationDB) {

  }

  // display simulation in router
  displaySimulationInRouter(simulationID: number) {
    this.router.navigate(["github_com_fullstack_lang_laundromat_go-" + "simulation-display", simulationID])
  }

  // set editor outlet
  setEditorRouterOutlet(simulationID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_laundromat_go_editor: ["github_com_fullstack_lang_laundromat_go-" + "simulation-detail", simulationID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(simulationID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_laundromat_go_presentation: ["github_com_fullstack_lang_laundromat_go-" + "simulation-presentation", simulationID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.simulations.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.simulations.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<SimulationDB>()

      // reset all initial selection of simulation that belong to simulation
      for (let simulation of this.initialSelection) {
        let index = simulation[this.dialogData.ReversePointer as keyof SimulationDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(simulation)

      }

      // from selection, set simulation that belong to simulation
      for (let simulation of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = simulation[this.dialogData.ReversePointer as keyof SimulationDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(simulation)
      }


      // update all simulation (only update selection & initial selection)
      for (let simulation of toUpdate) {
        this.simulationService.updateSimulation(simulation)
          .subscribe(simulation => {
            this.simulationService.SimulationServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, SimulationDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedSimulation = new Set<number>()
      for (let simulation of this.initialSelection) {
        if (this.selection.selected.includes(simulation)) {
          // console.log("simulation " + simulation.Name + " is still selected")
        } else {
          console.log("simulation " + simulation.Name + " has been unselected")
          unselectedSimulation.add(simulation.ID)
          console.log("is unselected " + unselectedSimulation.has(simulation.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let simulation = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as SimulationDB
      if (unselectedSimulation.has(simulation.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<SimulationDB>) = new Array<SimulationDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          simulation => {
            if (!this.initialSelection.includes(simulation)) {
              // console.log("simulation " + simulation.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + simulation.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = simulation.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = simulation.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("simulation " + simulation.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<SimulationDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
