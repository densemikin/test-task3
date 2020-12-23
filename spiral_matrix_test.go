package main

import (
	"reflect"
	"testing"
)

func TestValidateMatrix(t *testing.T) {
	inputCorrect := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}

	matrix := Matrix{inputCorrect}
	isValidMatrix := matrix.ValidateMatrix()

	if !isValidMatrix {
		t.Error("Matrix are not valid")
	}

	// Incorrect matrix

	inputIncorrect := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9, 10},
	}

	matrixIncorrect := Matrix{inputIncorrect}
	isValidMatrix = matrixIncorrect.ValidateMatrix()

	if isValidMatrix {
		t.Error("Matrix are valid")
	}
}

func TestGenerateMatrixPath(t *testing.T) {
	inputCorrect3Dim := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}

	matrix3Dim := Matrix{inputCorrect3Dim}
	isValidMatrix := matrix3Dim.ValidateMatrix()

	if !isValidMatrix {
		t.Error("Matrix not valid")
	} else {
		outputArray := matrix3Dim.GenerateMatrixPath()
		if !reflect.DeepEqual([]int{1,2,3,6,9,8,7,4,5}, outputArray) {
			t.Error("Generated path are incorrect")
		}
	}

	inputCorrect4Dim := [][]int{
		[]int{1,2,3,4},
		[]int{5,6,7,8},
		[]int{9,10,11,12},
		[]int{13,14,15,16},
	}

	matrix4Dim := Matrix{inputCorrect4Dim}
	isValidMatrix = matrix4Dim.ValidateMatrix()

	if !isValidMatrix {
		t.Error("Matrix not valid")
	} else {
		outputArray := matrix4Dim.GenerateMatrixPath()
		if !reflect.DeepEqual([]int{1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10}, outputArray) {
			t.Error("Generated path are incorrect")
		}
	}
}