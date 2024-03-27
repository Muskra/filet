package main

import (
    "fmt"
    "crypto/rand"
    "os"
    "test/filet"
)

var Rules filet.Rules = filet.Rules{
    RuleSet: []filet.Set{
        {
            CellValue: 0,
            CellState: false,
            TargetValue: 3,
            TargetState: false,
            ShouldBeTargeted: true,
            Opcode: filet.SWAP_CELLS,
        },
        {
            CellValue: 1,
            CellState: true,
            TargetValue: 5,
            TargetState: false,
            ShouldBeTargeted: false,
            Opcode: filet.KILL_CELL,
        },
        {
            CellValue: 2,
            CellState: true,
            TargetValue: 7,
            TargetState: false,
            ShouldBeTargeted: true,
            Opcode: filet.COPY_VALUE,
        },
        {
            CellValue: 3,
            CellState: false,
            TargetValue: 1,
            TargetState: true,
            ShouldBeTargeted: true,
            Opcode: filet.SWAP_CELLS,
        },
        {
            CellValue: 4,
            CellState: true,
            TargetValue: 0,
            TargetState: false,
            ShouldBeTargeted: false,
            Opcode: filet.SWAP_CELLS,
        },
        {
            CellValue: 5,
            CellState: false,
            TargetValue: 2,
            TargetState: true,
            ShouldBeTargeted: true,
            Opcode: filet.SWAP_CELLS,
        },
        {
            CellValue: 6,
            CellState: true,
            TargetValue: 4,
            TargetState: true,
            ShouldBeTargeted: true,
            Opcode: filet.SWAP_CELLS,
        },
        {
            CellValue: 7,
            CellState: false,
            TargetValue: 6,
            TargetState: true,
            ShouldBeTargeted: true,
            Opcode: filet.COPY_VALUE,
        },
    },
    
    TargetCellsLocations: []filet.Coordinates{
        {X: -42, Y: 36},
        {X: -1,  Y: 12},
        {X: 82,  Y: 0},
        {X: 2,   Y: 4},
        {X: 6,   Y: -99},
        {X: 5,   Y: -6},
        {X: -6,  Y: 2},
        {X: 33,  Y: 26},
        {X: 0,   Y: 35},
        {X: 231, Y: -3},
        {X: 6,   Y: 25},
        {X: -3,  Y: -1},
    },
    TargetValues: filet.Target{
        AliveValues:   []int{ 1, 2, 5 },
        DeadValues:    []int{ 0, 3, 4, 6 },
        TargetIfAlive: []int{ 3, 4, 6 },
        TargetIfDead:  []int{ 0, 1, 2, 5 },
    },
}

func main() {
    lines, cols := 4, 4
    var grid filet.Grid = filet.Grid{State: filet.GenerateTwoDimArray(lines, cols)}
    for i := 0; i < len(grid.State); i = i + 1 {
        for j := 0; j < len(grid.State[i]); j = j + 1 {
            tmp := make([]byte, 1)
            _, err := rand.Read(tmp)
            if err != nil {
                fmt.Println("error:", err)
                os.Exit(1)
            }
            grid.State[i][j].Value = int(tmp[0])
            if grid.State[i][j].Value % 2 == 1 {
                grid.State[i][j].State = true
            }
            if grid.State[i][j].Value % 2 == 0 {
                grid.State[i][j].Value = -grid.State[i][j].Value
                grid.State[i][j].State = false
            }
            grid.State[i][j].Position.X = i
            grid.State[i][j].Position.Y = j
            grid.State[i][j].ValidatedLinkedCells = make([]int, 0)
            grid.State[i][j].IsIn = false
        }
    }

    for _, e := range grid.State {
        for _, v := range e {
            fmt.Printf("X: %d, Y: %d, V: %d, S: %v, VLC: %v, II: %v\n", 
                v.Position.X,
                v.Position.Y,
                v.Value,
                v.State,
                v.ValidatedLinkedCells,
                v.IsIn,
            )
        }
        fmt.Println()
    }
    newGrid, err := filet.Catch(grid.State, Rules)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, e := range newGrid {
        for _, v := range e {
            fmt.Printf("X: %d, Y: %d, V: %d, S: %v, VLC: %v, II: %v\n", 
                v.Position.X,
                v.Position.Y,
                v.Value,
                v.State,
                v.ValidatedLinkedCells,
                v.IsIn,
            )
        }
        fmt.Println()
    }
}
