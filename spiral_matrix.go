package main

import "fmt"

type Matrix struct {
	inputMatrix [][]int
}

func (m* Matrix) ValidateMatrix() bool {
	nDim := len(m.inputMatrix)
	for i := 0; i < nDim; i++ {
		if len(m.inputMatrix[i]) != nDim {
			return false
		}
	}
	return true
}

func (m* Matrix) GenerateMatrixPath() []int {

	n := len(m.inputMatrix)

	topRow := 0
	leftColumn := 0

	bottomRow := n - 1
	rightColumn := bottomRow


	var outArray []int

	for topRow <= bottomRow && leftColumn <= rightColumn {

		// Move from left to right
		for i := topRow; i <= rightColumn; i++ {
			outArray = append(outArray, m.inputMatrix[leftColumn][i])
		}
		topRow++

		//Move from top right to bottom right
		for i := topRow; i <= bottomRow; i++ {
			outArray = append(outArray, m.inputMatrix[i][rightColumn])
		}
		rightColumn--

		// Move from bottom right to bottom left
		for i := rightColumn; i >= leftColumn; i-- {
			outArray = append(outArray, m.inputMatrix[bottomRow][i])
		}
		bottomRow--

		// Move from bottomb left to top left
		for i := bottomRow; i >= topRow; i-- {
			outArray = append(outArray, m.inputMatrix[i][leftColumn])
		}
		leftColumn++

	}
	return outArray
}

func main() {

	input := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}

	matrix := Matrix{input}
	if matrix.ValidateMatrix() {
		fmt.Println(matrix.GenerateMatrixPath())
	} else {
		fmt.Println("Incorrect matrix size")
	}
}
