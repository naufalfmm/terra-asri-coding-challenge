package main

import "math"

func CalculateDiagonalSum(gridSize uint) uint {
	if gridSize%2 == 0 {
		return 0
	}

	diagonalNumber := (gridSize / 2) + 1

	midTopRightSum := (gridSize * (gridSize + 1) * (gridSize + 2)) / 6
	midBottomLeftSum := (4*(uint(math.Pow(float64(diagonalNumber), 3)))-6*(uint(math.Pow(float64(diagonalNumber), 2)))+5*diagonalNumber)/3 - 1
	midTopLeftSum := (4*(uint(math.Pow(float64(diagonalNumber), 3)))-3*(uint(math.Pow(float64(diagonalNumber), 2)))+2*diagonalNumber)/3 - 1
	midBottomRightSum := (4*(uint(math.Pow(float64(diagonalNumber), 3)))-9*(uint(math.Pow(float64(diagonalNumber), 2)))+8*diagonalNumber)/3 - 1

	return midTopRightSum + midBottomLeftSum + midTopLeftSum + midBottomRightSum
}
