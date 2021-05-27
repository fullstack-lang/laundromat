// insertion point sub template for components imports 
  import { MachinesTableComponent } from './machines-table/machines-table.component'
  import { MachineSortingComponent } from './machine-sorting/machine-sorting.component'
  import { SimulationsTableComponent } from './simulations-table/simulations-table.component'
  import { SimulationSortingComponent } from './simulation-sorting/simulation-sorting.component'
  import { WashersTableComponent } from './washers-table/washers-table.component'
  import { WasherSortingComponent } from './washer-sorting/washer-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfMachinesComponents: Map<string, any> = new Map([["MachinesTableComponent", MachinesTableComponent],])
  export const MapOfMachineSortingComponents: Map<string, any> = new Map([["MachineSortingComponent", MachineSortingComponent],])
  export const MapOfSimulationsComponents: Map<string, any> = new Map([["SimulationsTableComponent", SimulationsTableComponent],])
  export const MapOfSimulationSortingComponents: Map<string, any> = new Map([["SimulationSortingComponent", SimulationSortingComponent],])
  export const MapOfWashersComponents: Map<string, any> = new Map([["WashersTableComponent", WashersTableComponent],])
  export const MapOfWasherSortingComponents: Map<string, any> = new Map([["WasherSortingComponent", WasherSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Machine", MapOfMachinesComponents],
      ["Simulation", MapOfSimulationsComponents],
      ["Washer", MapOfWashersComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Machine", MapOfMachineSortingComponents],
      ["Simulation", MapOfSimulationSortingComponents],
      ["Washer", MapOfWasherSortingComponents],
    ]
  )
