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
import { GongBasicFieldDB } from '../gongbasicfield-db'
import { GongBasicFieldService } from '../gongbasicfield.service'

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-gongbasicfieldstable',
  templateUrl: './gongbasicfields-table.component.html',
  styleUrls: ['./gongbasicfields-table.component.css'],
})
export class GongBasicFieldsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of GongBasicField instances
  selection: SelectionModel<GongBasicFieldDB> = new (SelectionModel)
  initialSelection = new Array<GongBasicFieldDB>()

  // the data source for the table
  gongbasicfields: GongBasicFieldDB[] = []
  matTableDataSource: MatTableDataSource<GongBasicFieldDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.gongbasicfields
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
    this.matTableDataSource.sortingDataAccessor = (gongbasicfieldDB: GongBasicFieldDB, property: string) => {
      switch (property) {
        case 'ID':
          return gongbasicfieldDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return gongbasicfieldDB.Name;

        case 'BasicKindName':
          return gongbasicfieldDB.BasicKindName;

        case 'GongEnum':
          return (gongbasicfieldDB.GongEnum ? gongbasicfieldDB.GongEnum.Name : '');

        case 'DeclaredType':
          return gongbasicfieldDB.DeclaredType;

        case 'Index':
          return gongbasicfieldDB.Index;

        case 'GongStruct_GongBasicFields':
          return this.frontRepo.GongStructs.get(gongbasicfieldDB.GongStruct_GongBasicFieldsDBID.Int64)!.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (gongbasicfieldDB: GongBasicFieldDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the gongbasicfieldDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += gongbasicfieldDB.Name.toLowerCase()
      mergedContent += gongbasicfieldDB.BasicKindName.toLowerCase()
      if (gongbasicfieldDB.GongEnum) {
        mergedContent += gongbasicfieldDB.GongEnum.Name.toLowerCase()
      }
      mergedContent += gongbasicfieldDB.DeclaredType.toLowerCase()
      mergedContent += gongbasicfieldDB.Index.toString()
      if (gongbasicfieldDB.GongStruct_GongBasicFieldsDBID.Int64 != 0) {
        mergedContent += this.frontRepo.GongStructs.get(gongbasicfieldDB.GongStruct_GongBasicFieldsDBID.Int64)!.Name.toLowerCase()
      }


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
    private gongbasicfieldService: GongBasicFieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongbasicfield instances
    public dialogRef: MatDialogRef<GongBasicFieldsTableComponent>,
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
    this.gongbasicfieldService.GongBasicFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongBasicFields()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "BasicKindName",
        "GongEnum",
        "DeclaredType",
        "Index",
        "GongStruct_GongBasicFields",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "BasicKindName",
        "GongEnum",
        "DeclaredType",
        "Index",
        "GongStruct_GongBasicFields",
      ]
      this.selection = new SelectionModel<GongBasicFieldDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongBasicFields()
    this.matTableDataSource = new MatTableDataSource(this.gongbasicfields)
  }

  getGongBasicFields(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.gongbasicfields = this.frontRepo.GongBasicFields_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let gongbasicfield of this.gongbasicfields) {
            let ID = this.dialogData.ID
            let revPointer = gongbasicfield[this.dialogData.ReversePointer as keyof GongBasicFieldDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(gongbasicfield)
            }
            this.selection = new SelectionModel<GongBasicFieldDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongBasicFieldDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as GongBasicFieldDB[]
          for (let associationInstance of sourceField) {
            let gongbasicfield = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongBasicFieldDB
            this.initialSelection.push(gongbasicfield)
          }

          this.selection = new SelectionModel<GongBasicFieldDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.gongbasicfields
      }
    )
  }

  // newGongBasicField initiate a new gongbasicfield
  // create a new GongBasicField objet
  newGongBasicField() {
  }

  deleteGongBasicField(gongbasicfieldID: number, gongbasicfield: GongBasicFieldDB) {
    // list of gongbasicfields is truncated of gongbasicfield before the delete
    this.gongbasicfields = this.gongbasicfields.filter(h => h !== gongbasicfield);

    this.gongbasicfieldService.deleteGongBasicField(gongbasicfieldID).subscribe(
      gongbasicfield => {
        this.gongbasicfieldService.GongBasicFieldServiceChanged.next("delete")
      }
    );
  }

  editGongBasicField(gongbasicfieldID: number, gongbasicfield: GongBasicFieldDB) {

  }

  // display gongbasicfield in router
  displayGongBasicFieldInRouter(gongbasicfieldID: number) {
    this.router.navigate(["github_com_fullstack_lang_gong_go-" + "gongbasicfield-display", gongbasicfieldID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongbasicfieldID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gongbasicfield-detail", gongbasicfieldID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongbasicfieldID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gong_go_presentation: ["github_com_fullstack_lang_gong_go-" + "gongbasicfield-presentation", gongbasicfieldID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongbasicfields.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongbasicfields.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<GongBasicFieldDB>()

      // reset all initial selection of gongbasicfield that belong to gongbasicfield
      for (let gongbasicfield of this.initialSelection) {
        let index = gongbasicfield[this.dialogData.ReversePointer as keyof GongBasicFieldDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(gongbasicfield)

      }

      // from selection, set gongbasicfield that belong to gongbasicfield
      for (let gongbasicfield of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = gongbasicfield[this.dialogData.ReversePointer as keyof GongBasicFieldDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(gongbasicfield)
      }


      // update all gongbasicfield (only update selection & initial selection)
      for (let gongbasicfield of toUpdate) {
        this.gongbasicfieldService.updateGongBasicField(gongbasicfield)
          .subscribe(gongbasicfield => {
            this.gongbasicfieldService.GongBasicFieldServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongBasicFieldDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedGongBasicField = new Set<number>()
      for (let gongbasicfield of this.initialSelection) {
        if (this.selection.selected.includes(gongbasicfield)) {
          // console.log("gongbasicfield " + gongbasicfield.Name + " is still selected")
        } else {
          console.log("gongbasicfield " + gongbasicfield.Name + " has been unselected")
          unselectedGongBasicField.add(gongbasicfield.ID)
          console.log("is unselected " + unselectedGongBasicField.has(gongbasicfield.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let gongbasicfield = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongBasicFieldDB
      if (unselectedGongBasicField.has(gongbasicfield.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<GongBasicFieldDB>) = new Array<GongBasicFieldDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          gongbasicfield => {
            if (!this.initialSelection.includes(gongbasicfield)) {
              // console.log("gongbasicfield " + gongbasicfield.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + gongbasicfield.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = gongbasicfield.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = gongbasicfield.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("gongbasicfield " + gongbasicfield.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<GongBasicFieldDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
