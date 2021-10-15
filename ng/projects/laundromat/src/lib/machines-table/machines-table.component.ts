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
import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-machinestable',
  templateUrl: './machines-table.component.html',
  styleUrls: ['./machines-table.component.css'],
})
export class MachinesTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Machine instances
  selection: SelectionModel<MachineDB> = new (SelectionModel)
  initialSelection = new Array<MachineDB>()

  // the data source for the table
  machines: MachineDB[] = []
  matTableDataSource: MatTableDataSource<MachineDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.machines
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
    this.matTableDataSource.sortingDataAccessor = (machineDB: MachineDB, property: string) => {
      switch (property) {
        // insertion point for specific sorting accessor
        case 'TechName':
          return machineDB.TechName;

        case 'Name':
          return machineDB.Name;

        case 'DrumLoad':
          return machineDB.DrumLoad;

        case 'RemainingTime':
          return machineDB.RemainingTime;

        case 'Cleanedlaundry':
          return machineDB.Cleanedlaundry?"true":"false";

        case 'State':
          return machineDB.State;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (machineDB: MachineDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the machineDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += machineDB.TechName.toLowerCase()
      mergedContent += machineDB.Name.toLowerCase()
      mergedContent += machineDB.DrumLoad.toString()
      mergedContent += machineDB.State.toLowerCase()

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
    private machineService: MachineService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of machine instances
    public dialogRef: MatDialogRef<MachinesTableComponent>,
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
    this.machineService.MachineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getMachines()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "TechName",
        "Name",
        "DrumLoad",
        "RemainingTime",
        "Cleanedlaundry",
        "State",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "TechName",
        "Name",
        "DrumLoad",
        "RemainingTime",
        "Cleanedlaundry",
        "State",
      ]
      this.selection = new SelectionModel<MachineDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getMachines()
    this.matTableDataSource = new MatTableDataSource(this.machines)
  }

  getMachines(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.machines = this.frontRepo.Machines_array;

        // insertion point for variables Recoveries
        // compute strings for durations
        for (let machine of this.machines) {
          machine.RemainingTime_string =
            Math.floor(machine.RemainingTime / (3600 * 1000 * 1000 * 1000)) + "H " +
            Math.floor(machine.RemainingTime % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000)) + "M " +
            machine.RemainingTime % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000) + "S"
        }

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          this.machines.forEach(
            machine => {
              let ID = this.dialogData.ID
              let revPointer = machine[this.dialogData.ReversePointer as keyof MachineDB] as unknown as NullInt64
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(machine)
              }
            }
          )
          this.selection = new SelectionModel<MachineDB>(allowMultiSelect, this.initialSelection);
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, MachineDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as MachineDB[]
          for (let associationInstance of sourceField) {
            let machine = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as MachineDB
            this.initialSelection.push(machine)
          }

          this.selection = new SelectionModel<MachineDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.machines
      }
    )
  }

  // newMachine initiate a new machine
  // create a new Machine objet
  newMachine() {
  }

  deleteMachine(machineID: number, machine: MachineDB) {
    // list of machines is truncated of machine before the delete
    this.machines = this.machines.filter(h => h !== machine);

    this.machineService.deleteMachine(machineID).subscribe(
      machine => {
        this.machineService.MachineServiceChanged.next("delete")
      }
    );
  }

  editMachine(machineID: number, machine: MachineDB) {

  }

  // display machine in router
  displayMachineInRouter(machineID: number) {
    this.router.navigate(["github_com_fullstack_lang_laundromat_go-" + "machine-display", machineID])
  }

  // set editor outlet
  setEditorRouterOutlet(machineID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_laundromat_go_editor: ["github_com_fullstack_lang_laundromat_go-" + "machine-detail", machineID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(machineID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_laundromat_go_presentation: ["github_com_fullstack_lang_laundromat_go-" + "machine-presentation", machineID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.machines.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.machines.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<MachineDB>()

      // reset all initial selection of machine that belong to machine
      this.initialSelection.forEach(
        machine => {
          let index = machine[this.dialogData.ReversePointer as keyof MachineDB] as unknown as NullInt64
          index.Int64 = 0
          index.Valid = true
          toUpdate.add(machine)
        }
      )

      // from selection, set machine that belong to machine
      this.selection.selected.forEach(
        machine => {
          let ID = this.dialogData.ID as number
          let reversePointer = machine[this.dialogData.ReversePointer  as keyof MachineDB] as unknown as NullInt64
          reversePointer.Int64 = ID
          toUpdate.add(machine)
        }
      )

      // update all machine (only update selection & initial selection)
      toUpdate.forEach(
        machine => {
          this.machineService.updateMachine(machine)
            .subscribe(machine => {
              this.machineService.MachineServiceChanged.next("update")
            });
        }
      )
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, MachineDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedMachine = new Set<number>()
      for (let machine of this.initialSelection) {
        if (this.selection.selected.includes(machine)) {
          // console.log("machine " + machine.Name + " is still selected")
        } else {
          console.log("machine " + machine.Name + " has been unselected")
          unselectedMachine.add(machine.ID)
          console.log("is unselected " + unselectedMachine.has(machine.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let machine = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as MachineDB
      if (unselectedMachine.has(machine.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<MachineDB>) = new Array<MachineDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          machine => {
            if (!this.initialSelection.includes(machine)) {
              // console.log("machine " + machine.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + machine.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField+"ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = machine.ID

              let indexDB = associationInstance[this.dialogData.IntermediateStructField+"DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = machine.ID

              this.frontRepoService.postService( this.dialogData.IntermediateStruct, associationInstance )

            } else {
              // console.log("machine " + machine.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<MachineDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
