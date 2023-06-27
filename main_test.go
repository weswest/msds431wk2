// This file contains tests for the regression calculations performed in the main.go file.
// It includes benchmark tests for running all regressions and tests for individual regression results.

package main

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

// BenchmarkRunAllRegressions benchmarks the performance of running all regressions.
func BenchmarkRunAllRegressions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		produceAllResults()
	}
}

// TestRegressionResults tests the regression results against expected values.
func TestRegressionResults(t *testing.T) {
	// Define the expected values for b0 and b1
	interceptExpected := 3.0
	slopeExpected := 0.5
	data1, data2, data3, data4 := makeStatsData()

	// Test regression 1
	testRegression(t, 1, data1, interceptExpected, slopeExpected)

	// Test regression 2
	testRegression(t, 2, data2, interceptExpected, slopeExpected)

	// Test regression 3
	testRegression(t, 3, data3, interceptExpected, slopeExpected)

	// Test regression 4
	testRegression(t, 4, data4, interceptExpected, slopeExpected)
}

// testRegression performs the regression calculation and compares the results with the expected values.
func testRegression(t *testing.T, regNum int, data []stats.Coordinate, interceptExpected, slopeExpected float64) {
	// Run the regression
	intercept, slope, err := produceSlopeIntercept(regNum, data)
	if err != nil {
		t.Errorf("Regression %d failed: %v", regNum, err)
	}

	// Round the actual coefficients to compare with the expected values
	interceptRound := math.Round(intercept*100) / 100
	slopeRound := math.Round(slope*100) / 100

	// Compare the actual and expected coefficients
	if interceptRound != interceptExpected {
		t.Errorf("Incorrect value for Intercept in regression %d. Expected: %f, Actual: %f", regNum, interceptExpected, interceptRound)
	}

	if slopeRound != slopeExpected {
		t.Errorf("Incorrect value for Slope in regression %d. Expected: %f, Actual: %f", regNum, slopeExpected, slopeRound)
	}
}
