# FILET 

A customizable cellullar automata.

## Disclaimer

Before any statements about cryptography or anythng related to security, i'm not a cryptography expert in any way. So do not use this program for any production or professionnal tooling if you don't know chat you're doing. I'm not responsible for any damage made with this program.

## What this library is all about ?

This library is a proof of concept project part that is basically a customizable cellular automata.
The main purpose of this project was originally to make a hashing algorithm but it was not simple without the right tool. So i wanted to do ciphering before hashing because it's simpler and more accessible.This tool can also be used to make basic cellullar automatas like game of life or other simulations so it's not only for ciphering.

## How it works

You throw a Grid Type into a 'filet.CatchOne()' or a 'filet.CatchNthGen()' function. Then the program will check for any cells from top left to right, top to bottom, and for any rules that you set into a 'Rules' Type, apply an 'Opcode' Type.

The library has some files made specifically to be modified like 'filet/config.go'. This file contains all the opcodes you want to use in the program. Any opcode can do anything, some tooling is provided but you can do whatever you want with them.

There is also a 'Rules' Type provided to make custom rules that the program will check before applying an opcode. The rules are applied if and only if:

- a cell 'State' is equal to a rule cell 'Value' AND
- a cell 'State' is equal to a rule cell 'State' AND
   - So this way we check if the rule really should be considered for this rule application
- a target 'Value' is equal to a rule 'TargetValue' AND
   - this 'TargetValue' represent the 'Value' of the target to compare with
- a target 'State' is equal to a rule 'TargetState' AND
   - this 'TargetState' represent the 'State' of the target to compare with
- the target 'IsIn' value is equal to the rule 'ShouldBeTargeted' value
   - this is a boolean to simply tell if the rule should be applied or not in the previous conditions already checked

## Provided tooling within the library

The library provides some accessible methods and functions, the list is below:

> you can generate the documentation of them with the 'godoc' or 'go doc' commands

- cell.go
   - Methods
      - Equal
      - IsDeadOrAlive
   - Functions
      - IsAlive
      - IsDead
      - PrintDetailedState
- coordinates.go
   - Methods
      - FindRealTargetLocations
- filet.go
   - Methods
      - CatchOne
      - CatchNthGen
- grid.go
   - Methods
      - ActualCellState
         - NextGeneration
         - Reverse
         - FormatState
         - PrintState
         - PrintDetailedState
   - Functions
      - GenerateTwoDimSlice
- opcodes.go
   - Methods
      - ProcessRule
- set.go
   - Functions
      - NewSet
- targets.go
   - Functions
      - NewTarget
