// insertion point sub template for components imports 
  import { DummyAgentsTableComponent } from './dummyagents-table/dummyagents-table.component'
  import { DummyAgentSortingComponent } from './dummyagent-sorting/dummyagent-sorting.component'
  import { EnginesTableComponent } from './engines-table/engines-table.component'
  import { EngineSortingComponent } from './engine-sorting/engine-sorting.component'
  import { EventsTableComponent } from './events-table/events-table.component'
  import { EventSortingComponent } from './event-sorting/event-sorting.component'
  import { GongsimCommandsTableComponent } from './gongsimcommands-table/gongsimcommands-table.component'
  import { GongsimCommandSortingComponent } from './gongsimcommand-sorting/gongsimcommand-sorting.component'
  import { GongsimStatussTableComponent } from './gongsimstatuss-table/gongsimstatuss-table.component'
  import { GongsimStatusSortingComponent } from './gongsimstatus-sorting/gongsimstatus-sorting.component'
  import { UpdateStatesTableComponent } from './updatestates-table/updatestates-table.component'
  import { UpdateStateSortingComponent } from './updatestate-sorting/updatestate-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfDummyAgentsComponents: Map<string, any> = new Map([["DummyAgentsTableComponent", DummyAgentsTableComponent],])
  export const MapOfDummyAgentSortingComponents: Map<string, any> = new Map([["DummyAgentSortingComponent", DummyAgentSortingComponent],])
  export const MapOfEnginesComponents: Map<string, any> = new Map([["EnginesTableComponent", EnginesTableComponent],])
  export const MapOfEngineSortingComponents: Map<string, any> = new Map([["EngineSortingComponent", EngineSortingComponent],])
  export const MapOfEventsComponents: Map<string, any> = new Map([["EventsTableComponent", EventsTableComponent],])
  export const MapOfEventSortingComponents: Map<string, any> = new Map([["EventSortingComponent", EventSortingComponent],])
  export const MapOfGongsimCommandsComponents: Map<string, any> = new Map([["GongsimCommandsTableComponent", GongsimCommandsTableComponent],])
  export const MapOfGongsimCommandSortingComponents: Map<string, any> = new Map([["GongsimCommandSortingComponent", GongsimCommandSortingComponent],])
  export const MapOfGongsimStatussComponents: Map<string, any> = new Map([["GongsimStatussTableComponent", GongsimStatussTableComponent],])
  export const MapOfGongsimStatusSortingComponents: Map<string, any> = new Map([["GongsimStatusSortingComponent", GongsimStatusSortingComponent],])
  export const MapOfUpdateStatesComponents: Map<string, any> = new Map([["UpdateStatesTableComponent", UpdateStatesTableComponent],])
  export const MapOfUpdateStateSortingComponents: Map<string, any> = new Map([["UpdateStateSortingComponent", UpdateStateSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["DummyAgent", MapOfDummyAgentsComponents],
      ["Engine", MapOfEnginesComponents],
      ["Event", MapOfEventsComponents],
      ["GongsimCommand", MapOfGongsimCommandsComponents],
      ["GongsimStatus", MapOfGongsimStatussComponents],
      ["UpdateState", MapOfUpdateStatesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["DummyAgent", MapOfDummyAgentSortingComponents],
      ["Engine", MapOfEngineSortingComponents],
      ["Event", MapOfEventSortingComponents],
      ["GongsimCommand", MapOfGongsimCommandSortingComponents],
      ["GongsimStatus", MapOfGongsimStatusSortingComponents],
      ["UpdateState", MapOfUpdateStateSortingComponents],
    ]
  )
