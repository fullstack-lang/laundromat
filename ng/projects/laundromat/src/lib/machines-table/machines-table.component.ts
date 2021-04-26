// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-machines-table',
  templateUrl: './machines-table.component.html',
  styleUrls: ['./machines-table.component.css'],
})
export class MachinesTableComponent implements OnInit {

  // used if the component is called as a selection component of Machine instances
  selection: SelectionModel<MachineDB>;
  initialSelection = new Array<MachineDB>();

  // the data source for the table
  machines: MachineDB[];
  matTableDataSource: MatTableDataSource<MachineDB>


  // front repo, that will be referenced by this.machines
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort) sort: MatSort;
  @ViewChild(MatPaginator) paginator: MatPaginator;

  ngAfterViewInit() {
    this.matTableDataSource.sort = this.sort;
    this.matTableDataSource.paginator = this.paginator;
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
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.machineService.MachineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getMachines()
        }
      }
    )
    if (dialogData == undefined) {
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
        console.log("front repo pull returned")

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
        if (this.dialogData != undefined) {
          this.machines.forEach(
            machine => {
              let ID = this.dialogData.ID
              let revPointer = machine[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(machine)
              }
            }
          )
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

        console.log("machine deleted")
      }
    );
  }

  editMachine(machineID: number, machine: MachineDB) {

  }

  // display machine in router
  displayMachineInRouter(machineID: number) {
    this.router.navigate(["machine-display", machineID])
  }

  // set editor outlet
  setEditorRouterOutlet(machineID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["machine-detail", machineID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(machineID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["machine-presentation", machineID]
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

    let toUpdate = new Set<MachineDB>()

    // reset all initial selection of machine that belong to machine through Anarrayofb
    this.initialSelection.forEach(
      machine => {
        machine[this.dialogData.ReversePointer].Int64 = 0
        machine[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(machine)
      }
    )

    // from selection, set machine that belong to machine through Anarrayofb
    this.selection.selected.forEach(
      machine => {
        console.log("selection ID " + machine.ID)
        let ID = +this.dialogData.ID
        machine[this.dialogData.ReversePointer].Int64 = ID
        machine[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(machine)
      }
    )

    // update all machine (only update selection & initial selection)
    toUpdate.forEach(
      machine => {
        this.machineService.updateMachine(machine)
          .subscribe(machine => {
            this.machineService.MachineServiceChanged.next("update")
            console.log("machine saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
