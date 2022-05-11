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
import { GongsimCommandDB } from '../gongsimcommand-db'
import { GongsimCommandService } from '../gongsimcommand.service'

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
  selector: 'app-gongsimcommandstable',
  templateUrl: './gongsimcommands-table.component.html',
  styleUrls: ['./gongsimcommands-table.component.css'],
})
export class GongsimCommandsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of GongsimCommand instances
  selection: SelectionModel<GongsimCommandDB> = new (SelectionModel)
  initialSelection = new Array<GongsimCommandDB>()

  // the data source for the table
  gongsimcommands: GongsimCommandDB[] = []
  matTableDataSource: MatTableDataSource<GongsimCommandDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.gongsimcommands
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
    this.matTableDataSource.sortingDataAccessor = (gongsimcommandDB: GongsimCommandDB, property: string) => {
      switch (property) {
        case 'ID':
          return gongsimcommandDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return gongsimcommandDB.Name;

        case 'Command':
          return gongsimcommandDB.Command;

        case 'CommandDate':
          return gongsimcommandDB.CommandDate;

        case 'SpeedCommandType':
          return gongsimcommandDB.SpeedCommandType;

        case 'DateSpeedCommand':
          return gongsimcommandDB.DateSpeedCommand;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (gongsimcommandDB: GongsimCommandDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the gongsimcommandDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += gongsimcommandDB.Name.toLowerCase()
      mergedContent += gongsimcommandDB.Command.toLowerCase()
      mergedContent += gongsimcommandDB.CommandDate.toLowerCase()
      mergedContent += gongsimcommandDB.SpeedCommandType.toLowerCase()
      mergedContent += gongsimcommandDB.DateSpeedCommand.toLowerCase()

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
    private gongsimcommandService: GongsimCommandService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongsimcommand instances
    public dialogRef: MatDialogRef<GongsimCommandsTableComponent>,
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
    this.gongsimcommandService.GongsimCommandServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongsimCommands()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Command",
        "CommandDate",
        "SpeedCommandType",
        "DateSpeedCommand",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Command",
        "CommandDate",
        "SpeedCommandType",
        "DateSpeedCommand",
      ]
      this.selection = new SelectionModel<GongsimCommandDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongsimCommands()
    this.matTableDataSource = new MatTableDataSource(this.gongsimcommands)
  }

  getGongsimCommands(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.gongsimcommands = this.frontRepo.GongsimCommands_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let gongsimcommand of this.gongsimcommands) {
            let ID = this.dialogData.ID
            let revPointer = gongsimcommand[this.dialogData.ReversePointer as keyof GongsimCommandDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(gongsimcommand)
            }
            this.selection = new SelectionModel<GongsimCommandDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongsimCommandDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as GongsimCommandDB[]
          for (let associationInstance of sourceField) {
            let gongsimcommand = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongsimCommandDB
            this.initialSelection.push(gongsimcommand)
          }

          this.selection = new SelectionModel<GongsimCommandDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.gongsimcommands
      }
    )
  }

  // newGongsimCommand initiate a new gongsimcommand
  // create a new GongsimCommand objet
  newGongsimCommand() {
  }

  deleteGongsimCommand(gongsimcommandID: number, gongsimcommand: GongsimCommandDB) {
    // list of gongsimcommands is truncated of gongsimcommand before the delete
    this.gongsimcommands = this.gongsimcommands.filter(h => h !== gongsimcommand);

    this.gongsimcommandService.deleteGongsimCommand(gongsimcommandID).subscribe(
      gongsimcommand => {
        this.gongsimcommandService.GongsimCommandServiceChanged.next("delete")
      }
    );
  }

  editGongsimCommand(gongsimcommandID: number, gongsimcommand: GongsimCommandDB) {

  }

  // display gongsimcommand in router
  displayGongsimCommandInRouter(gongsimcommandID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongsim_go-" + "gongsimcommand-display", gongsimcommandID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongsimcommandID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongsim_go_editor: ["github_com_fullstack_lang_gongsim_go-" + "gongsimcommand-detail", gongsimcommandID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongsimcommandID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongsim_go_presentation: ["github_com_fullstack_lang_gongsim_go-" + "gongsimcommand-presentation", gongsimcommandID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongsimcommands.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongsimcommands.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<GongsimCommandDB>()

      // reset all initial selection of gongsimcommand that belong to gongsimcommand
      for (let gongsimcommand of this.initialSelection) {
        let index = gongsimcommand[this.dialogData.ReversePointer as keyof GongsimCommandDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(gongsimcommand)

      }

      // from selection, set gongsimcommand that belong to gongsimcommand
      for (let gongsimcommand of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = gongsimcommand[this.dialogData.ReversePointer as keyof GongsimCommandDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(gongsimcommand)
      }


      // update all gongsimcommand (only update selection & initial selection)
      for (let gongsimcommand of toUpdate) {
        this.gongsimcommandService.updateGongsimCommand(gongsimcommand)
          .subscribe(gongsimcommand => {
            this.gongsimcommandService.GongsimCommandServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongsimCommandDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedGongsimCommand = new Set<number>()
      for (let gongsimcommand of this.initialSelection) {
        if (this.selection.selected.includes(gongsimcommand)) {
          // console.log("gongsimcommand " + gongsimcommand.Name + " is still selected")
        } else {
          console.log("gongsimcommand " + gongsimcommand.Name + " has been unselected")
          unselectedGongsimCommand.add(gongsimcommand.ID)
          console.log("is unselected " + unselectedGongsimCommand.has(gongsimcommand.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let gongsimcommand = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongsimCommandDB
      if (unselectedGongsimCommand.has(gongsimcommand.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<GongsimCommandDB>) = new Array<GongsimCommandDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          gongsimcommand => {
            if (!this.initialSelection.includes(gongsimcommand)) {
              // console.log("gongsimcommand " + gongsimcommand.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + gongsimcommand.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = gongsimcommand.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = gongsimcommand.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("gongsimcommand " + gongsimcommand.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<GongsimCommandDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
