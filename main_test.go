package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkRunAllRegressions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runAllRegressions(false)
	}
}

func TestRegressionResults(t *testing.T) {
	// Define the expected values for b0 and b1
	expectedB0 := 3.0
	expectedB1 := 0.5

	// Test regression 1
	testRegression(t, 1, x1, y1, expectedB0, expectedB1)

	// Test regression 2
	testRegression(t, 2, x2, y2, expectedB0, expectedB1)

	// Test regression 3
	testRegression(t, 3, x3, y3, expectedB0, expectedB1)

	// Test regression 4
	testRegression(t, 4, x4, y4, expectedB0, expectedB1)
}

func testRegression(t *testing.T, rNum int, x []float64, y []float64, expectedB0, expectedB1 float64) {
	// Run the regression
	b0, b1 := runRegression(false, rNum, x, y)

	// Round the actual coefficients to compare with the expected values
	b0Round := math.Round(b0*100) / 100
	b1Round := math.Round(b1*100) / 100

	// Use the testify/assert package to perform assertions
	assert.Equal(t, expectedB0, b0Round, "Incorrect value for b0 in regression %d", rNum)
	assert.Equal(t, expectedB1, b1Round, "Incorrect value for b1 in regression %d", rNum)
}
