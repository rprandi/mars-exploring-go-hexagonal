# Mars Exploring
A repo to implement a programming challenge.

## The Challenge

### The problem

A set of probes translate was sent by NASA to Mars and will land on a plateau. 

The plateau is rectangular and must be explored by the probe with the built-in cameras.

The position and direction of the probe are represented by a set of coordinates (x,y) and a letter representing the direction:
- N is north
- S is south
- E is East
- W is west

To simplify the implementation, the plateau is a grid and (0,0,N) means the probe is in the bottom-left corner of the grid facing north

To control the probe, NASA will send a string of characters:
- L will make the probe turn 90 degrees left
- R will make the probe turn 90 degrees right
- M will make the probe move forward 1 step of the grid in the direction it is facing.

In this grid, the point north of (x,y) is always (x, y+1)

### Input

First line: The coordinates in the top-right of the grid. Reminder: bottom-left is always (0,0)

Second line (onwards): Probe information, 2 lines per probe:
- First line is the initial position
- Second line is the set of instructions to explore the plateau

#### Example of input

5 5

1 2 N

LMLMLMLMM

3 3 E

MMRMMRMRRM

### Output

1 line per probe (same order of the input) describing the final coordinates and direction

#### Example of output

1 3 N

5 1 E

### Rules
- No external framework (a framework for BDD testing is ok)
- Tests and validations are allowed
- Using Go language
- Input must come from a file or terminal/cli/STDIO
- Program must output to a file or terminal/cli 
- Any rules not specified here can be implemented as you wish (eg. What happens when 2 probes collide?) 

## Discussions and Considerations

### Why Golang ?

Goland, while not Object-oriented, has interfaces and could allow me to show how this problem could be implemented with a Hexagonal architecture. 

It is an overengineering for such a simple problem, however, it does fit it (as you have at least 2 inbound adapters for input and 2 outbound adapters for output)

It will also allow me to explicitly use SOLID principles as Goland has structs with interfaces. Other languages like Ruby are OO, but don't have explicit interfaces.

### About testing

I will use TDD because it will allow me to iterate over the problem adding new use cases and then refactor when those tests pass. 
This will allow me to refactor the code easily as I solve the problem.

I will use BDD because this in type of challenge is easy to come up with new use cases, and i will have the output well described. 
This will allow me to add new behaviour to the tests easily for TDD.ginkgo
I will use Ginkgo lib in Go because of my familiarity with this tool. It can be easily read by developers not used to it.

The focus of TDD is not the output, but the testing. The focus of BDD is the output. They complement each other well.

I will use unit and integration tests as I see fit throughout the development.

### Configuration

- I am using GoLand IDE
- I will leave goimport / golang-lint / gofmt running on the background (trigger = save) on my IDE