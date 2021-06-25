// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'

import { FrontRepoService, FrontRepo, NullInt64 } from '../front-repo.service'
@Component({
  selector: 'lib-machine-sorting',
  templateUrl: './machine-sorting.component.html',
  styleUrls: ['./machine-sorting.component.css']
})
export class MachineSortingComponent implements OnInit {

  frontRepo: FrontRepo

  // array of Machine instances that are in the association
  associatedMachines = new Array<MachineDB>();

  constructor(
    private machineService: MachineService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of machine instances
    public dialogRef: MatDialogRef<MachineSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getMachines()
  }

  getMachines(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let machine of this.frontRepo.Machines_array) {
          let ID = this.dialogData.ID
          let revPointerID = machine[this.dialogData.ReversePointer]
          let revPointerID_Index = machine[this.dialogData.ReversePointer+"_Index"]
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedMachines.push(machine)
          }
        }

        // sort associated machine according to order
        this.associatedMachines.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer+"_Index"]
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer+"_Index"]
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }  
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedMachines, event.previousIndex, event.currentIndex);

    // set the order of Machine instances
    let index = 0
    
    for (let machine of this.associatedMachines) {
      let revPointerID_Index = machine[this.dialogData.ReversePointer+"_Index"]
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedMachines.forEach(
      machine => {
        this.machineService.updateMachine(machine)
          .subscribe(machine => {
            this.machineService.MachineServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of '+ this.dialogData.ReversePointer+' done');
  }
}