package main

import (
    "fmt"
    "test/filet"
)

var Rules filet.Rules = filet.Rules{
    RuleSet: []filet.Set{
        filet.Set{
            CellValue: 0,
            CellState: false,
            TargetValue: 3,
            TargetState: false,
            ShouldBeTargeted: true,
            Opcode: filet.SWAP_CELLS,
        },
        filet.Set{
            CellValue: 1,
            CellState: true,
            TargetValue: 5,
            TargetState: false,
            ShouldBeTargeted: false,
            Opcode: filet.KILL_CELL,
        },
        filet.Set{
            CellValue: 2,
            CellState: true,
            TargetValue: 7,
            TargetState: false,
            ShouldBeTargeted: true,
            Opcode: filet.COPY_VALUE,
        },
        filet.Set{
            CellValue: 3,
            CellState: false,
            TargetValue: 1,
            TargetState: true,
            ShouldBeTargeted: true,
            Opcode: filet.SWAP_CELLS,
        },
        filet.Set{
            CellValue: 4,
            CellState: true,
            TargetValue: 0,
            TargetState: false,
            ShouldBeTargeted: false,
            Opcode: filet.SWAP_CELLS,
        },
        filet.Set{
            CellValue: 5,
            CellState: false,
            TargetValue: 2,
            TargetState: true,
            ShouldBeTargeted: true,
            Opcode: filet.SWAP_CELLS,
        },
        filet.Set{
            CellValue: 6,
            CellState: true,
            TargetValue: 4,
            TargetState: true,
            ShouldBeTargeted: true,
            Opcode: filet.SWAP_CELLS,
        },
        filet.Set{
            CellValue: 7,
            CellState: false,
            TargetValue: 6,
            TargetState: true,
            ShouldBeTargeted: true,
            Opcode: filet.COPY_VALUE,
        },
    },
    
    TargetCellsLocations: []filet.Coordinates{
        filet.Coordinates{X: -42, Y: 36},
        filet.Coordinates{X: -1, Y: 12},
        filet.Coordinates{X: 82, Y: 0},
        filet.Coordinates{X: 2, Y: 4},
        filet.Coordinates{X: 6, Y: -99},
        filet.Coordinates{X: 5, Y: -6},
        filet.Coordinates{X: -6, Y: 2},
        filet.Coordinates{X: 33, Y: 26},
        filet.Coordinates{X: 0, Y: 35},
        filet.Coordinates{X: 231, Y: -3},
        filet.Coordinates{X: 6, Y: 25},
        filet.Coordinates{X: -3, Y: -1},
    },
    TargetValues: filet.Target{
        AliveValues: []int{

        },
        DeadValues: []int{

        },
        TargetIfAlive: []int{

        },
        TargetIfDead: []int{

        },
    },
}

func main() {
    /*
    set := filet.Set{
        cellValue: ,
        cellState: ,
        targetValue: ,
        shouldBeTargeted: ,
        opcode: ,
    }
    */

    var grid [][]filet.Cell = [][]filet.Cell{
        []filet.Cell{
            filet.Cell{Value: 0},
            filet.Cell{Value: 1},
        },
        []filet.Cell{
            filet.Cell{Value: 2},
            filet.Cell{Value: 3},
        },
        []filet.Cell{
            filet.Cell{Value: 4},
            filet.Cell{Value: 5},
        },
        []filet.Cell{
            filet.Cell{Value: 6},
            filet.Cell{Value: 7},
        },
    }
    


}
