// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { GongsimCommandDB } from '../gongsimcommand-db'
import { GongsimCommandService } from '../gongsimcommand.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-gongsimcommand-sorting',
  templateUrl: './gongsimcommand-sorting.component.html',
  styleUrls: ['./gongsimcommand-sorting.component.css']
})
export class GongsimCommandSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of GongsimCommand instances that are in the association
  associatedGongsimCommands = new Array<GongsimCommandDB>();

  constructor(
    private gongsimcommandService: GongsimCommandService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongsimcommand instances
    public dialogRef: MatDialogRef<GongsimCommandSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getGongsimCommands()
  }

  getGongsimCommands(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let gongsimcommand of this.frontRepo.GongsimCommands_array) {
          let ID = this.dialogData.ID
          let revPointerID = gongsimcommand[this.dialogData.ReversePointer as keyof GongsimCommandDB] as unknown as NullInt64
          let revPointerID_Index = gongsimcommand[this.dialogData.ReversePointer + "_Index" as keyof GongsimCommandDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedGongsimCommands.push(gongsimcommand)
          }
        }

        // sort associated gongsimcommand according to order
        this.associatedGongsimCommands.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
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
    moveItemInArray(this.associatedGongsimCommands, event.previousIndex, event.currentIndex);

    // set the order of GongsimCommand instances
    let index = 0

    for (let gongsimcommand of this.associatedGongsimCommands) {
      let revPointerID_Index = gongsimcommand[this.dialogData.ReversePointer + "_Index" as keyof GongsimCommandDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedGongsimCommands.forEach(
      gongsimcommand => {
        this.gongsimcommandService.updateGongsimCommand(gongsimcommand)
          .subscribe(gongsimcommand => {
            this.gongsimcommandService.GongsimCommandServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer +' done');
  }
}
