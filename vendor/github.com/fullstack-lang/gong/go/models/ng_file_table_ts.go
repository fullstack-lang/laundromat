package models

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "embed"
)

//go:embed ng_file_table.css
var NgFileTableCssTmpl string

const NgTableTemplateTS = `// generated by gong
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
import { {{Structname}}DB } from '../{{structname}}-db'
import { {{Structname}}Service } from '../{{structname}}.service'

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-{{structname}}stable',
  templateUrl: './{{structname}}s-table.component.html',
  styleUrls: ['./{{structname}}s-table.component.css'],
})
export class {{Structname}}sTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of {{Structname}} instances
  selection: SelectionModel<{{Structname}}DB> = new (SelectionModel)
  initialSelection = new Array<{{Structname}}DB>()

  // the data source for the table
  {{structname}}s: {{Structname}}DB[] = []
  matTableDataSource: MatTableDataSource<{{Structname}}DB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.{{structname}}s
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
    this.matTableDataSource.sortingDataAccessor = ({{structname}}DB: {{Structname}}DB, property: string) => {
      switch (property) {
        case 'ID':
          return {{structname}}DB.ID

        // insertion point for specific sorting accessor{{` + string(rune(NgTableTsInsertionPerStructColumnsSorting)) + `}}
        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = ({{structname}}DB: {{Structname}}DB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the {{structname}}DB properties
      let mergedContent = ""

      // insertion point for merging of fields{{` + string(rune(NgTableTsInsertionPerStructColumnsFiltering)) + `}}

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
    private {{structname}}Service: {{Structname}}Service,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of {{structname}} instances
    public dialogRef: MatDialogRef<{{Structname}}sTableComponent>,
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
    this.{{structname}}Service.{{Structname}}ServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.get{{Structname}}s()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display{{` + string(rune(NgTableTsInsertionPerStructColumns)) + `}}
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display{{` + string(rune(NgTableTsInsertionPerStructColumns)) + `}}
      ]
      this.selection = new SelectionModel<{{Structname}}DB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.get{{Structname}}s()
    this.matTableDataSource = new MatTableDataSource(this.{{structname}}s)
  }

  get{{Structname}}s(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.{{structname}}s = this.frontRepo.{{Structname}}s_array;

        // insertion point for variables Recoveries{{` + string(rune(NgTableTsInsertionPerStructRecoveries)) + `}}

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let {{structname}} of this.{{structname}}s) {
            let ID = this.dialogData.ID
            let revPointer = {{structname}}[this.dialogData.ReversePointer as keyof {{Structname}}DB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push({{structname}})
            }
            this.selection = new SelectionModel<{{Structname}}DB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, {{Structname}}DB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as {{Structname}}DB[]
          for (let associationInstance of sourceField) {
            let {{structname}} = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as {{Structname}}DB
            this.initialSelection.push({{structname}})
          }

          this.selection = new SelectionModel<{{Structname}}DB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.{{structname}}s
      }
    )
  }

  // new{{Structname}} initiate a new {{structname}}
  // create a new {{Structname}} objet
  new{{Structname}}() {
  }

  delete{{Structname}}({{structname}}ID: number, {{structname}}: {{Structname}}DB) {
    // list of {{structname}}s is truncated of {{structname}} before the delete
    this.{{structname}}s = this.{{structname}}s.filter(h => h !== {{structname}});

    this.{{structname}}Service.delete{{Structname}}({{structname}}ID).subscribe(
      {{structname}} => {
        this.{{structname}}Service.{{Structname}}ServiceChanged.next("delete")
      }
    );
  }

  edit{{Structname}}({{structname}}ID: number, {{structname}}: {{Structname}}DB) {

  }

  // display {{structname}} in router
  display{{Structname}}InRouter({{structname}}ID: number) {
    this.router.navigate(["{{PkgPathRootWithoutSlashes}}-" + "{{structname}}-display", {{structname}}ID])
  }

  // set editor outlet
  setEditorRouterOutlet({{structname}}ID: number) {
    this.router.navigate([{
      outlets: {
        {{PkgPathRootWithoutSlashes}}_editor: ["{{PkgPathRootWithoutSlashes}}-" + "{{structname}}-detail", {{structname}}ID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet({{structname}}ID: number) {
    this.router.navigate([{
      outlets: {
        {{PkgPathRootWithoutSlashes}}_presentation: ["{{PkgPathRootWithoutSlashes}}-" + "{{structname}}-presentation", {{structname}}ID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.{{structname}}s.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.{{structname}}s.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<{{Structname}}DB>()

      // reset all initial selection of {{structname}} that belong to {{structname}}
      for (let {{structname}} of this.initialSelection) {
        let index = {{structname}}[this.dialogData.ReversePointer as keyof {{Structname}}DB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add({{structname}})

      }

      // from selection, set {{structname}} that belong to {{structname}}
      for (let {{structname}} of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = {{structname}}[this.dialogData.ReversePointer as keyof {{Structname}}DB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add({{structname}})
      }


      // update all {{structname}} (only update selection & initial selection)
      for (let {{structname}} of toUpdate) {
        this.{{structname}}Service.update{{Structname}}({{structname}})
          .subscribe({{structname}} => {
            this.{{structname}}Service.{{Structname}}ServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, {{Structname}}DB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselected{{Structname}} = new Set<number>()
      for (let {{structname}} of this.initialSelection) {
        if (this.selection.selected.includes({{structname}})) {
          // console.log("{{structname}} " + {{structname}}.Name + " is still selected")
        } else {
          console.log("{{structname}} " + {{structname}}.Name + " has been unselected")
          unselected{{Structname}}.add({{structname}}.ID)
          console.log("is unselected " + unselected{{Structname}}.has({{structname}}.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let {{structname}} = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as {{Structname}}DB
      if (unselected{{Structname}}.has({{structname}}.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<{{Structname}}DB>) = new Array<{{Structname}}DB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          {{structname}} => {
            if (!this.initialSelection.includes({{structname}})) {
              // console.log("{{structname}} " + {{structname}}.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + {{structname}}.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = {{structname}}.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = {{structname}}.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("{{structname}} " + {{structname}}.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<{{Structname}}DB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
`

// insertion points in the main template
type NgTableTsInsertionPoint int

const (
	NgTableTsInsertionPerStructRecoveries NgTableTsInsertionPoint = iota
	NgTableTsInsertionPerStructColumns
	NgTableTsInsertionPerStructColumnsSorting
	NgTableTsInsertionPerStructColumnsFiltering
	NgTableTsInsertionsNb
)

type NgTableSubTemplate int

const (
	NgTableTSPerStructTimeDurationRecoveries NgTableSubTemplate = iota

	NgTableTSBasicFieldSorting
	NgTableTSTimeFieldSorting

	NgTableTSPointerToStructSorting
	NgTableTSSliceOfPointerToStructSorting

	NgTableTSNonNumberFieldFiltering
	NgTableTSNumberFieldFiltering
	NgTableTSTimeFieldFiltering
	NgTableTSPointerToStructFiltering
	NgTableTSSliceOfPointerToStructFiltering

	NgTableTSPerStructColumn
	NgTableTSSliceOfPointerToStructPerStructColumn
)

var NgTablelSubTemplateCode map[NgTableSubTemplate]string = map[NgTableSubTemplate]string{

	NgTableTSPerStructTimeDurationRecoveries: `
        // compute strings for durations
        for (let {{structname}} of this.{{structname}}s) {
          {{structname}}.{{FieldName}}_string =
            Math.floor({{structname}}.{{FieldName}} / (3600 * 1000 * 1000 * 1000)) + "H " +
            Math.floor({{structname}}.{{FieldName}} % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000)) + "M " +
            {{structname}}.{{FieldName}} % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000) + "S"
        }`,

	NgTableTSBasicFieldSorting: `
        case '{{FieldName}}':
          return {{structname}}DB.{{FieldName}}{{TranslationIntoString}};
`,
	NgTableTSTimeFieldSorting: `
        case '{{FieldName}}':
          return {{structname}}DB.{{FieldName}}.getDate();
`,
	NgTableTSPointerToStructSorting: `
        case '{{FieldName}}':
          return ({{structname}}DB.{{FieldName}} ? {{structname}}DB.{{FieldName}}.Name : '');
`,
	NgTableTSSliceOfPointerToStructSorting: `
        case '{{AssocStructName}}_{{FieldName}}':
          return this.frontRepo.{{AssocStructName}}s.get({{structname}}DB.{{AssocStructName}}_{{FieldName}}DBID.Int64)!.Name;
`,

	NgTableTSNonNumberFieldFiltering: `
      mergedContent += {{structname}}DB.{{FieldName}}.toLowerCase()`,
	NgTableTSNumberFieldFiltering: `
      mergedContent += {{structname}}DB.{{FieldName}}.toString()`,
	NgTableTSTimeFieldFiltering: `
`,
	NgTableTSPointerToStructFiltering: `
      if ({{structname}}DB.{{FieldName}}) {
        mergedContent += {{structname}}DB.{{FieldName}}.Name.toLowerCase()
      }`,
	NgTableTSSliceOfPointerToStructFiltering: `
      if ({{structname}}DB.{{AssocStructName}}_{{FieldName}}DBID.Int64 != 0) {
        mergedContent += this.frontRepo.{{AssocStructName}}s.get({{structname}}DB.{{AssocStructName}}_{{FieldName}}DBID.Int64)!.Name.toLowerCase()
      }
`,

	NgTableTSPerStructColumn: `
        "{{FieldName}}",`,

	NgTableTSSliceOfPointerToStructPerStructColumn: `
        "{{AssocStructName}}_{{FieldName}}",`,
}

// MultiCodeGeneratorNgTable parses mdlPkg and generates the code for the
// Table component
func MultiCodeGeneratorNgTable(
	mdlPkg *ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string) {

	// have alphabetical order generation
	structList := []*GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range mdlPkg.GongStructs {
		if !_struct.HasNameField() {
			continue
		}

		// create the component directory
		dirPath := filepath.Join(matTargetPath, strings.ToLower(_struct.Name)+"s-table")
		errd := os.MkdirAll(dirPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + dirPath)
		}

		// generate the css file
		VerySimpleCodeGenerator(mdlPkg,
			pkgName,
			pkgGoPath,
			filepath.Join(dirPath, strings.ToLower(_struct.Name)+"s-table.component.css"),
			NgFileTableCssTmpl,
		)

		// generate the typescript file
		codeTS := NgTableTemplateTS

		TsInsertions := make(map[NgTableTsInsertionPoint]string)
		for insertion := NgTableTsInsertionPoint(0); insertion < NgTableTsInsertionsNb; insertion++ {
			TsInsertions[insertion] = ""
		}

		HtmlInsertions := make(map[NgTableHtmlInsertionPoint]string)
		for insertion := NgTableHtmlInsertionPoint(0); insertion < NgTableHtmlInsertionsNb; insertion++ {
			HtmlInsertions[insertion] = ""
		}

		codeHTML := NgTableTemplateHTML

		for _, field := range _struct.Fields {
			switch field := field.(type) {
			case *GongBasicField:

				// conversion form go type to ts type
				TypeInput := ""
				switch field.basicKind {
				case types.Int, types.Float64:
					TypeInput = "type=\"number\" [ngModelOptions]=\"{standalone: true}\" "
				case types.String:
					TypeInput = "name=\"\" [ngModelOptions]=\"{standalone: true}\"     "
				}

				switch field.basicKind {
				case types.Float64:
					HtmlInsertions[NgTableHtmlInsertionColumn] += Replace2(NgTableHTMLSubTemplateCode[NgTableHTMLBasicFloat64Field],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", TypeInput)
					TsInsertions[NgTableTsInsertionPerStructColumnsFiltering] += Replace1(NgTablelSubTemplateCode[NgTableTSNumberFieldFiltering],
						"{{FieldName}}", field.Name)
				case types.Int, types.Int64:
					if field.DeclaredType != "time.Duration" {
						HtmlInsertions[NgTableHtmlInsertionColumn] += Replace2(NgTableHTMLSubTemplateCode[NgTableHTMLBasicField],
							"{{FieldName}}", field.Name,
							"{{TypeInput}}", TypeInput)
						TsInsertions[NgTableTsInsertionPerStructColumnsFiltering] += Replace1(NgTablelSubTemplateCode[NgTableTSNumberFieldFiltering],
							"{{FieldName}}", field.Name)

					} else {
						HtmlInsertions[NgTableHtmlInsertionColumn] += Replace1(NgTableHTMLSubTemplateCode[NgTableHTMLBasicFieldTimeDuration],
							"{{FieldName}}", field.Name)
					}

				case types.Bool:
					HtmlInsertions[NgTableHtmlInsertionColumn] += Replace2(NgTableHTMLSubTemplateCode[NgTableHTMLBool],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", TypeInput)
				default:

					HtmlInsertions[NgTableHtmlInsertionColumn] += Replace1(NgTableHTMLSubTemplateCode[NgTableHTMLBasicField],
						"{{FieldName}}", field.Name)
					TsInsertions[NgTableTsInsertionPerStructColumnsFiltering] += Replace1(NgTablelSubTemplateCode[NgTableTSNonNumberFieldFiltering],
						"{{FieldName}}", field.Name)

				}
				TsInsertions[NgTableTsInsertionPerStructColumns] +=
					Replace1(NgTablelSubTemplateCode[NgTableTSPerStructColumn],
						"{{FieldName}}", field.Name)

				if field.DeclaredType == "time.Duration" {
					TsInsertions[NgTableTsInsertionPerStructRecoveries] +=
						Replace1(NgTablelSubTemplateCode[NgTableTSPerStructTimeDurationRecoveries],
							"{{FieldName}}", field.Name)
				}

				// sorting requires a string translation for each field
				// for boolean, one needs true or false
				translationString := ""
				if field.basicKind == types.Bool {
					translationString = "?\"true\":\"false\""
				}

				TsInsertions[NgTableTsInsertionPerStructColumnsSorting] +=
					Replace2(NgTablelSubTemplateCode[NgTableTSBasicFieldSorting],
						"{{FieldName}}", field.Name,
						"{{TranslationIntoString}}", translationString)

			case *GongTimeField:

				TsInsertions[NgTableTsInsertionPerStructColumns] +=
					Replace1(NgTablelSubTemplateCode[NgTableTSPerStructColumn],
						"{{FieldName}}", field.Name)

				TsInsertions[NgTableTsInsertionPerStructColumnsSorting] +=
					Replace1(NgTablelSubTemplateCode[NgTableTSTimeFieldSorting],
						"{{FieldName}}", field.Name)

				HtmlInsertions[NgTableHtmlInsertionColumn] += Replace1(NgTableHTMLSubTemplateCode[NgTableHTMLTimeField],
					"{{FieldName}}", field.Name)
			case *PointerToGongStructField:

				HtmlInsertions[NgTableHtmlInsertionColumn] += Replace3(NgTableHTMLSubTemplateCode[NgTablePointerToStructHTMLFormField],
					"{{FieldName}}", field.Name,
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))

				TsInsertions[NgTableTsInsertionPerStructColumns] +=
					Replace1(NgTablelSubTemplateCode[NgTableTSPerStructColumn],
						"{{FieldName}}", field.Name)

				TsInsertions[NgTableTsInsertionPerStructColumnsSorting] +=
					Replace1(NgTablelSubTemplateCode[NgTableTSPointerToStructSorting],
						"{{FieldName}}", field.Name)

				TsInsertions[NgTableTsInsertionPerStructColumnsFiltering] +=
					Replace1(NgTablelSubTemplateCode[NgTableTSPointerToStructFiltering],
						"{{FieldName}}", field.Name)
			}
		}

		//
		// Parse all fields from other structs that points to this struct
		//
		for _, __struct := range structList {
			for _, field := range __struct.Fields {
				switch field := field.(type) {
				case *SliceOfPointerToGongStructField:

					if field.GongStruct == _struct {
						HtmlInsertions[NgTableHtmlInsertionColumn] += Replace2(NgTableHTMLSubTemplateCode[NgTablePointerToSliceOfGongStructHTMLFormField],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", __struct.Name)

						TsInsertions[NgTableTsInsertionPerStructColumns] +=
							Replace2(NgTablelSubTemplateCode[NgTableTSSliceOfPointerToStructPerStructColumn],
								"{{FieldName}}", field.Name,
								"{{AssocStructName}}", __struct.Name)

						TsInsertions[NgTableTsInsertionPerStructColumnsSorting] +=
							Replace2(NgTablelSubTemplateCode[NgTableTSSliceOfPointerToStructSorting],
								"{{AssocStructName}}", __struct.Name,
								"{{FieldName}}", field.Name)

						TsInsertions[NgTableTsInsertionPerStructColumnsFiltering] +=
							Replace2(NgTablelSubTemplateCode[NgTableTSSliceOfPointerToStructFiltering],
								"{{AssocStructName}}", __struct.Name,
								"{{FieldName}}", field.Name)
					}
				}
			}
		}

		for insertion := NgTableTsInsertionPoint(0); insertion < NgTableTsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, TsInsertions[insertion])
		}

		for insertion := NgTableHtmlInsertionPoint(0); insertion < NgTableHtmlInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeHTML = strings.ReplaceAll(codeHTML, toReplace, HtmlInsertions[insertion])
		}

		pkgPathRootWithoutSlashes := strings.ReplaceAll(pkgGoPath, "/models", "")
		pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, "/", "_")
		pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, "-", "_")
		pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, ".", "_")

		// final replacement
		codeTS = Replace7(codeTS,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", strings.Title(pkgName),
			"{{pkgname}}", strings.ToLower(pkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name),
			"{{PkgPathRootWithoutSlashes}}", pkgPathRootWithoutSlashes)
		codeHTML = Replace6(codeHTML,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", strings.Title(pkgName),
			"{{pkgname}}", strings.ToLower(pkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))
		{
			file, err := os.Create(filepath.Join(dirPath, strings.ToLower(_struct.Name)+"s-table.component.ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeTS)
		}
		{
			file, err := os.Create(filepath.Join(dirPath, strings.ToLower(_struct.Name)+"s-table.component.html"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeHTML)
		}

	}
}
