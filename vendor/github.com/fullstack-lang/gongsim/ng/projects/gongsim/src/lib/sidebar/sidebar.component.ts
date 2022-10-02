import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { DummyAgentService } from '../dummyagent.service'
import { getDummyAgentUniqueID } from '../front-repo.service'
import { EngineService } from '../engine.service'
import { getEngineUniqueID } from '../front-repo.service'
import { EventService } from '../event.service'
import { getEventUniqueID } from '../front-repo.service'
import { GongsimCommandService } from '../gongsimcommand.service'
import { getGongsimCommandUniqueID } from '../front-repo.service'
import { GongsimStatusService } from '../gongsimstatus.service'
import { getGongsimStatusUniqueID } from '../front-repo.service'

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongsim-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNb: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbService: CommitNbService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private dummyagentService: DummyAgentService,
    private engineService: EngineService,
    private eventService: EventService,
    private gongsimcommandService: GongsimCommandService,
    private gongsimstatusService: GongsimStatusService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        this.setTableRouterOutlet(selectedStruct)
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.dummyagentService.DummyAgentServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.engineService.EngineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.eventService.EventServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongsimcommandService.GongsimCommandServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongsimstatusService.GongsimStatusServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the DummyAgent part of the mat tree
      */
      let dummyagentGongNodeStruct: GongNode = {
        name: "DummyAgent",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "DummyAgent",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(dummyagentGongNodeStruct)

      this.frontRepo.DummyAgents_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.DummyAgents_array.forEach(
        dummyagentDB => {
          let dummyagentGongNodeInstance: GongNode = {
            name: dummyagentDB.Name,
            type: GongNodeType.INSTANCE,
            id: dummyagentDB.ID,
            uniqueIdPerStack: getDummyAgentUniqueID(dummyagentDB.ID),
            structName: "DummyAgent",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          dummyagentGongNodeStruct.children!.push(dummyagentGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Engine part of the mat tree
      */
      let engineGongNodeStruct: GongNode = {
        name: "Engine",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Engine",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(engineGongNodeStruct)

      this.frontRepo.Engines_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Engines_array.forEach(
        engineDB => {
          let engineGongNodeInstance: GongNode = {
            name: engineDB.Name,
            type: GongNodeType.INSTANCE,
            id: engineDB.ID,
            uniqueIdPerStack: getEngineUniqueID(engineDB.ID),
            structName: "Engine",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          engineGongNodeStruct.children!.push(engineGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Event part of the mat tree
      */
      let eventGongNodeStruct: GongNode = {
        name: "Event",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Event",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(eventGongNodeStruct)

      this.frontRepo.Events_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Events_array.forEach(
        eventDB => {
          let eventGongNodeInstance: GongNode = {
            name: eventDB.Name,
            type: GongNodeType.INSTANCE,
            id: eventDB.ID,
            uniqueIdPerStack: getEventUniqueID(eventDB.ID),
            structName: "Event",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          eventGongNodeStruct.children!.push(eventGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the GongsimCommand part of the mat tree
      */
      let gongsimcommandGongNodeStruct: GongNode = {
        name: "GongsimCommand",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongsimCommand",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongsimcommandGongNodeStruct)

      this.frontRepo.GongsimCommands_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongsimCommands_array.forEach(
        gongsimcommandDB => {
          let gongsimcommandGongNodeInstance: GongNode = {
            name: gongsimcommandDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongsimcommandDB.ID,
            uniqueIdPerStack: getGongsimCommandUniqueID(gongsimcommandDB.ID),
            structName: "GongsimCommand",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongsimcommandGongNodeStruct.children!.push(gongsimcommandGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the GongsimStatus part of the mat tree
      */
      let gongsimstatusGongNodeStruct: GongNode = {
        name: "GongsimStatus",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongsimStatus",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongsimstatusGongNodeStruct)

      this.frontRepo.GongsimStatuss_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongsimStatuss_array.forEach(
        gongsimstatusDB => {
          let gongsimstatusGongNodeInstance: GongNode = {
            name: gongsimstatusDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongsimstatusDB.ID,
            uniqueIdPerStack: getGongsimStatusUniqueID(gongsimstatusDB.ID),
            structName: "GongsimStatus",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongsimstatusGongNodeStruct.children!.push(gongsimstatusGongNodeInstance)

          // insertion point for per field code
        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    });

    // fetch the number of commits
    this.commitNbService.getCommitNb().subscribe(
      commitNb => {
        this.commitNb = commitNb
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongsim_go_table: ["github_com_fullstack_lang_gongsim_go-" + path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongsim_go_table: ["github_com_fullstack_lang_gongsim_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongsim_go_presentation: ["github_com_fullstack_lang_gongsim_go-" + structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongsim_go_editor: ["github_com_fullstack_lang_gongsim_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongsim_go_editor: ["github_com_fullstack_lang_gongsim_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}
