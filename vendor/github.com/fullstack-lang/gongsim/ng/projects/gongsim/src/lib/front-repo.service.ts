import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { DummyAgentDB } from './dummyagent-db'
import { DummyAgentService } from './dummyagent.service'

import { EngineDB } from './engine-db'
import { EngineService } from './engine.service'

import { EventDB } from './event-db'
import { EventService } from './event.service'

import { GongsimCommandDB } from './gongsimcommand-db'
import { GongsimCommandService } from './gongsimcommand.service'

import { GongsimStatusDB } from './gongsimstatus-db'
import { GongsimStatusService } from './gongsimstatus.service'

import { UpdateStateDB } from './updatestate-db'
import { UpdateStateService } from './updatestate.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  DummyAgents_array = new Array<DummyAgentDB>(); // array of repo instances
  DummyAgents = new Map<number, DummyAgentDB>(); // map of repo instances
  DummyAgents_batch = new Map<number, DummyAgentDB>(); // same but only in last GET (for finding repo instances to delete)
  Engines_array = new Array<EngineDB>(); // array of repo instances
  Engines = new Map<number, EngineDB>(); // map of repo instances
  Engines_batch = new Map<number, EngineDB>(); // same but only in last GET (for finding repo instances to delete)
  Events_array = new Array<EventDB>(); // array of repo instances
  Events = new Map<number, EventDB>(); // map of repo instances
  Events_batch = new Map<number, EventDB>(); // same but only in last GET (for finding repo instances to delete)
  GongsimCommands_array = new Array<GongsimCommandDB>(); // array of repo instances
  GongsimCommands = new Map<number, GongsimCommandDB>(); // map of repo instances
  GongsimCommands_batch = new Map<number, GongsimCommandDB>(); // same but only in last GET (for finding repo instances to delete)
  GongsimStatuss_array = new Array<GongsimStatusDB>(); // array of repo instances
  GongsimStatuss = new Map<number, GongsimStatusDB>(); // map of repo instances
  GongsimStatuss_batch = new Map<number, GongsimStatusDB>(); // same but only in last GET (for finding repo instances to delete)
  UpdateStates_array = new Array<UpdateStateDB>(); // array of repo instances
  UpdateStates = new Map<number, UpdateStateDB>(); // map of repo instances
  UpdateStates_batch = new Map<number, UpdateStateDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private dummyagentService: DummyAgentService,
    private engineService: EngineService,
    private eventService: EventService,
    private gongsimcommandService: GongsimCommandService,
    private gongsimstatusService: GongsimStatusService,
    private updatestateService: UpdateStateService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<DummyAgentDB[]>,
    Observable<EngineDB[]>,
    Observable<EventDB[]>,
    Observable<GongsimCommandDB[]>,
    Observable<GongsimStatusDB[]>,
    Observable<UpdateStateDB[]>,
  ] = [ // insertion point sub template 
      this.dummyagentService.getDummyAgents(),
      this.engineService.getEngines(),
      this.eventService.getEvents(),
      this.gongsimcommandService.getGongsimCommands(),
      this.gongsimstatusService.getGongsimStatuss(),
      this.updatestateService.getUpdateStates(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            dummyagents_,
            engines_,
            events_,
            gongsimcommands_,
            gongsimstatuss_,
            updatestates_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var dummyagents: DummyAgentDB[]
            dummyagents = dummyagents_ as DummyAgentDB[]
            var engines: EngineDB[]
            engines = engines_ as EngineDB[]
            var events: EventDB[]
            events = events_ as EventDB[]
            var gongsimcommands: GongsimCommandDB[]
            gongsimcommands = gongsimcommands_ as GongsimCommandDB[]
            var gongsimstatuss: GongsimStatusDB[]
            gongsimstatuss = gongsimstatuss_ as GongsimStatusDB[]
            var updatestates: UpdateStateDB[]
            updatestates = updatestates_ as UpdateStateDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.DummyAgents_array = dummyagents

            // clear the map that counts DummyAgent in the GET
            FrontRepoSingloton.DummyAgents_batch.clear()

            dummyagents.forEach(
              dummyagent => {
                FrontRepoSingloton.DummyAgents.set(dummyagent.ID, dummyagent)
                FrontRepoSingloton.DummyAgents_batch.set(dummyagent.ID, dummyagent)
              }
            )

            // clear dummyagents that are absent from the batch
            FrontRepoSingloton.DummyAgents.forEach(
              dummyagent => {
                if (FrontRepoSingloton.DummyAgents_batch.get(dummyagent.ID) == undefined) {
                  FrontRepoSingloton.DummyAgents.delete(dummyagent.ID)
                }
              }
            )

            // sort DummyAgents_array array
            FrontRepoSingloton.DummyAgents_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Engines_array = engines

            // clear the map that counts Engine in the GET
            FrontRepoSingloton.Engines_batch.clear()

            engines.forEach(
              engine => {
                FrontRepoSingloton.Engines.set(engine.ID, engine)
                FrontRepoSingloton.Engines_batch.set(engine.ID, engine)
              }
            )

            // clear engines that are absent from the batch
            FrontRepoSingloton.Engines.forEach(
              engine => {
                if (FrontRepoSingloton.Engines_batch.get(engine.ID) == undefined) {
                  FrontRepoSingloton.Engines.delete(engine.ID)
                }
              }
            )

            // sort Engines_array array
            FrontRepoSingloton.Engines_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Events_array = events

            // clear the map that counts Event in the GET
            FrontRepoSingloton.Events_batch.clear()

            events.forEach(
              event => {
                FrontRepoSingloton.Events.set(event.ID, event)
                FrontRepoSingloton.Events_batch.set(event.ID, event)
              }
            )

            // clear events that are absent from the batch
            FrontRepoSingloton.Events.forEach(
              event => {
                if (FrontRepoSingloton.Events_batch.get(event.ID) == undefined) {
                  FrontRepoSingloton.Events.delete(event.ID)
                }
              }
            )

            // sort Events_array array
            FrontRepoSingloton.Events_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.GongsimCommands_array = gongsimcommands

            // clear the map that counts GongsimCommand in the GET
            FrontRepoSingloton.GongsimCommands_batch.clear()

            gongsimcommands.forEach(
              gongsimcommand => {
                FrontRepoSingloton.GongsimCommands.set(gongsimcommand.ID, gongsimcommand)
                FrontRepoSingloton.GongsimCommands_batch.set(gongsimcommand.ID, gongsimcommand)
              }
            )

            // clear gongsimcommands that are absent from the batch
            FrontRepoSingloton.GongsimCommands.forEach(
              gongsimcommand => {
                if (FrontRepoSingloton.GongsimCommands_batch.get(gongsimcommand.ID) == undefined) {
                  FrontRepoSingloton.GongsimCommands.delete(gongsimcommand.ID)
                }
              }
            )

            // sort GongsimCommands_array array
            FrontRepoSingloton.GongsimCommands_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.GongsimStatuss_array = gongsimstatuss

            // clear the map that counts GongsimStatus in the GET
            FrontRepoSingloton.GongsimStatuss_batch.clear()

            gongsimstatuss.forEach(
              gongsimstatus => {
                FrontRepoSingloton.GongsimStatuss.set(gongsimstatus.ID, gongsimstatus)
                FrontRepoSingloton.GongsimStatuss_batch.set(gongsimstatus.ID, gongsimstatus)
              }
            )

            // clear gongsimstatuss that are absent from the batch
            FrontRepoSingloton.GongsimStatuss.forEach(
              gongsimstatus => {
                if (FrontRepoSingloton.GongsimStatuss_batch.get(gongsimstatus.ID) == undefined) {
                  FrontRepoSingloton.GongsimStatuss.delete(gongsimstatus.ID)
                }
              }
            )

            // sort GongsimStatuss_array array
            FrontRepoSingloton.GongsimStatuss_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.UpdateStates_array = updatestates

            // clear the map that counts UpdateState in the GET
            FrontRepoSingloton.UpdateStates_batch.clear()

            updatestates.forEach(
              updatestate => {
                FrontRepoSingloton.UpdateStates.set(updatestate.ID, updatestate)
                FrontRepoSingloton.UpdateStates_batch.set(updatestate.ID, updatestate)
              }
            )

            // clear updatestates that are absent from the batch
            FrontRepoSingloton.UpdateStates.forEach(
              updatestate => {
                if (FrontRepoSingloton.UpdateStates_batch.get(updatestate.ID) == undefined) {
                  FrontRepoSingloton.UpdateStates.delete(updatestate.ID)
                }
              }
            )

            // sort UpdateStates_array array
            FrontRepoSingloton.UpdateStates_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            dummyagents.forEach(
              dummyagent => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Engine redeeming
                {
                  let _engine = FrontRepoSingloton.Engines.get(dummyagent.EngineID.Int64)
                  if (_engine) {
                    dummyagent.Engine = _engine
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            engines.forEach(
              engine => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            events.forEach(
              event => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            gongsimcommands.forEach(
              gongsimcommand => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            gongsimstatuss.forEach(
              gongsimstatus => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            updatestates.forEach(
              updatestate => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // DummyAgentPull performs a GET on DummyAgent of the stack and redeem association pointers 
  DummyAgentPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.dummyagentService.getDummyAgents()
        ]).subscribe(
          ([ // insertion point sub template 
            dummyagents,
          ]) => {
            // init the array
            FrontRepoSingloton.DummyAgents_array = dummyagents

            // clear the map that counts DummyAgent in the GET
            FrontRepoSingloton.DummyAgents_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            dummyagents.forEach(
              dummyagent => {
                FrontRepoSingloton.DummyAgents.set(dummyagent.ID, dummyagent)
                FrontRepoSingloton.DummyAgents_batch.set(dummyagent.ID, dummyagent)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Engine redeeming
                {
                  let _engine = FrontRepoSingloton.Engines.get(dummyagent.EngineID.Int64)
                  if (_engine) {
                    dummyagent.Engine = _engine
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear dummyagents that are absent from the GET
            FrontRepoSingloton.DummyAgents.forEach(
              dummyagent => {
                if (FrontRepoSingloton.DummyAgents_batch.get(dummyagent.ID) == undefined) {
                  FrontRepoSingloton.DummyAgents.delete(dummyagent.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // EnginePull performs a GET on Engine of the stack and redeem association pointers 
  EnginePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.engineService.getEngines()
        ]).subscribe(
          ([ // insertion point sub template 
            engines,
          ]) => {
            // init the array
            FrontRepoSingloton.Engines_array = engines

            // clear the map that counts Engine in the GET
            FrontRepoSingloton.Engines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            engines.forEach(
              engine => {
                FrontRepoSingloton.Engines.set(engine.ID, engine)
                FrontRepoSingloton.Engines_batch.set(engine.ID, engine)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear engines that are absent from the GET
            FrontRepoSingloton.Engines.forEach(
              engine => {
                if (FrontRepoSingloton.Engines_batch.get(engine.ID) == undefined) {
                  FrontRepoSingloton.Engines.delete(engine.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // EventPull performs a GET on Event of the stack and redeem association pointers 
  EventPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.eventService.getEvents()
        ]).subscribe(
          ([ // insertion point sub template 
            events,
          ]) => {
            // init the array
            FrontRepoSingloton.Events_array = events

            // clear the map that counts Event in the GET
            FrontRepoSingloton.Events_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            events.forEach(
              event => {
                FrontRepoSingloton.Events.set(event.ID, event)
                FrontRepoSingloton.Events_batch.set(event.ID, event)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear events that are absent from the GET
            FrontRepoSingloton.Events.forEach(
              event => {
                if (FrontRepoSingloton.Events_batch.get(event.ID) == undefined) {
                  FrontRepoSingloton.Events.delete(event.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // GongsimCommandPull performs a GET on GongsimCommand of the stack and redeem association pointers 
  GongsimCommandPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongsimcommandService.getGongsimCommands()
        ]).subscribe(
          ([ // insertion point sub template 
            gongsimcommands,
          ]) => {
            // init the array
            FrontRepoSingloton.GongsimCommands_array = gongsimcommands

            // clear the map that counts GongsimCommand in the GET
            FrontRepoSingloton.GongsimCommands_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongsimcommands.forEach(
              gongsimcommand => {
                FrontRepoSingloton.GongsimCommands.set(gongsimcommand.ID, gongsimcommand)
                FrontRepoSingloton.GongsimCommands_batch.set(gongsimcommand.ID, gongsimcommand)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear gongsimcommands that are absent from the GET
            FrontRepoSingloton.GongsimCommands.forEach(
              gongsimcommand => {
                if (FrontRepoSingloton.GongsimCommands_batch.get(gongsimcommand.ID) == undefined) {
                  FrontRepoSingloton.GongsimCommands.delete(gongsimcommand.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // GongsimStatusPull performs a GET on GongsimStatus of the stack and redeem association pointers 
  GongsimStatusPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongsimstatusService.getGongsimStatuss()
        ]).subscribe(
          ([ // insertion point sub template 
            gongsimstatuss,
          ]) => {
            // init the array
            FrontRepoSingloton.GongsimStatuss_array = gongsimstatuss

            // clear the map that counts GongsimStatus in the GET
            FrontRepoSingloton.GongsimStatuss_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongsimstatuss.forEach(
              gongsimstatus => {
                FrontRepoSingloton.GongsimStatuss.set(gongsimstatus.ID, gongsimstatus)
                FrontRepoSingloton.GongsimStatuss_batch.set(gongsimstatus.ID, gongsimstatus)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear gongsimstatuss that are absent from the GET
            FrontRepoSingloton.GongsimStatuss.forEach(
              gongsimstatus => {
                if (FrontRepoSingloton.GongsimStatuss_batch.get(gongsimstatus.ID) == undefined) {
                  FrontRepoSingloton.GongsimStatuss.delete(gongsimstatus.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // UpdateStatePull performs a GET on UpdateState of the stack and redeem association pointers 
  UpdateStatePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.updatestateService.getUpdateStates()
        ]).subscribe(
          ([ // insertion point sub template 
            updatestates,
          ]) => {
            // init the array
            FrontRepoSingloton.UpdateStates_array = updatestates

            // clear the map that counts UpdateState in the GET
            FrontRepoSingloton.UpdateStates_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            updatestates.forEach(
              updatestate => {
                FrontRepoSingloton.UpdateStates.set(updatestate.ID, updatestate)
                FrontRepoSingloton.UpdateStates_batch.set(updatestate.ID, updatestate)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear updatestates that are absent from the GET
            FrontRepoSingloton.UpdateStates.forEach(
              updatestate => {
                if (FrontRepoSingloton.UpdateStates_batch.get(updatestate.ID) == undefined) {
                  FrontRepoSingloton.UpdateStates.delete(updatestate.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getDummyAgentUniqueID(id: number): number {
  return 31 * id
}
export function getEngineUniqueID(id: number): number {
  return 37 * id
}
export function getEventUniqueID(id: number): number {
  return 41 * id
}
export function getGongsimCommandUniqueID(id: number): number {
  return 43 * id
}
export function getGongsimStatusUniqueID(id: number): number {
  return 47 * id
}
export function getUpdateStateUniqueID(id: number): number {
  return 53 * id
}
