package problem3

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const SPIRAL_SIZE = 51 //Must be odd in order for spiral to have proper centre

type Direction int

const (
	RIGHT = iota
	UP
	LEFT
	DOWN
)

type Location struct {
	x, y int
}

type Problem3 struct{}

func (p *Problem3) Solve(data string, part int) int {
	intData, _ := strconv.Atoi(strings.TrimSpace(data))
	if part == 1 {
		return solvePart1(intData)
	}
	return solvePart2(intData)
}

func solvePart1(data int) int {
	finalLocation, _ := createSpiral(data, 1)
	return getDistance(finalLocation)
}

func solvePart2(data int) int {
	_, greaterVal := createSpiral(data, 2)
	return greaterVal
}

func createSpiral(maxVal, part int) (Location, int) {
	spiral := make([][]int, SPIRAL_SIZE)
	for i := 0; i < SPIRAL_SIZE; i++ {
		spiral[i] = make([]int, SPIRAL_SIZE)
	}
	location, value := fillSpiral(spiral, maxVal, part)

	return location, value
}

func fillSpiral(spiral [][]int, maxVal, part int) (Location, int) {
	location := Location{x: (SPIRAL_SIZE - 1) / 2, y: (SPIRAL_SIZE - 1) / 2}
	currentVal := 1
	currentDirection := DOWN

	for (part == 1 && currentVal < maxVal) || (part == 2 && currentVal <= maxVal) {
		spiral[location.y][location.x] = currentVal
		switch currentDirection {
		case RIGHT:
			if spiral[location.y-1][location.x] == 0 {
				currentDirection = UP
				location.y -= 1
			} else {
				location.x += 1
			}
		case UP:
			if spiral[location.y][location.x-1] == 0 {
				currentDirection = LEFT
				location.x -= 1
			} else {
				location.y -= 1
			}
		case LEFT:
			if spiral[location.y+1][location.x] == 0 {
				currentDirection = DOWN
				location.y += 1
			} else {
				location.x -= 1
			}
		case DOWN:
			if spiral[location.y][location.x+1] == 0 {
				currentDirection = RIGHT
				location.x += 1
			} else {
				location.y += 1
			}
		default:
			fmt.Println("Unknown direction, something went wrong")
		}

		if part == 1 {
			currentVal += 1
		} else if part == 2 {
			currentVal = getSumOfSurrounding(location, spiral)
		}
	}
	return location, currentVal
}

func getDistance(location Location) int {
	xDistance := math.Abs(float64(location.x - (SPIRAL_SIZE-1)/2))
	yDistance := math.Abs(float64(location.y - (SPIRAL_SIZE-1)/2))
	return int(xDistance + yDistance)
}

func getSumOfSurrounding(location Location, spiral [][]int) int {
	total := 0
	for y := location.y - 1; y <= location.y+1; y++ {
		for x := location.x - 1; x <= location.x+1; x++ {
			if location.y != y || location.x != x {
				total += spiral[y][x]
			}
		}
	}
	return total
}
