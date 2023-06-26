// This program demonstrates linear regression calculations using the montanaflynn/stats package.
// It performs regression analysis on multiple datasets and calculates the slope and intercept of each regression.
//
// The input data is provided in separate arrays for x and y values. The program creates stats.Coordinate
// data from these arrays and performs linear regression using the stats.LinearRegression function.
// The slope and intercept of each regression line are then calculated using the calcSlopeIntercept function.
//
// The produceAllResults function calls produceSlopeIntercept for each dataset and prints the results to the console.
//
// The main function calls produceAllResults and handles any errors that occur during the calculations.
//
// Note: The montanaflynn/stats package is used for convenience, but alternative implementations can be used as well.

package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

var (
	// Input data
	x1 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x2 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x3 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x4 = []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y1 = []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	y2 = []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	y3 = []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
	y4 = []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}
)

// makeStatsData creates stats.Coordinate data from input arrays.
func makeStatsData() ([]stats.Coordinate, []stats.Coordinate, []stats.Coordinate, []stats.Coordinate) {
	data1 := make([]stats.Coordinate, len(x1))
	data2 := make([]stats.Coordinate, len(x2))
	data3 := make([]stats.Coordinate, len(x3))
	data4 := make([]stats.Coordinate, len(x4))

	for i := 0; i < len(x1); i++ {
		data1[i] = stats.Coordinate{X: x1[i], Y: y1[i]}
		data2[i] = stats.Coordinate{X: x2[i], Y: y2[i]}
		data3[i] = stats.Coordinate{X: x3[i], Y: y3[i]}
		data4[i] = stats.Coordinate{X: x4[i], Y: y4[i]}
	}

	return data1, data2, data3, data4
}

// calcSlopeIntercept calculates the slope and intercept of a linear regression.
func calcSlopeIntercept(reg stats.Series) (float64, float64, error) {
	if len(reg) == 0 {
		return 0, 0, fmt.Errorf("empty series")
	}

	// Note: this function is designed to be used with the LinearRegression function from the montanaflynn/stats package.
	// We know that the y values are the predicted values from a linear equation, so we know definitively the structure
	// of the data generation equation.
	var x1, y1, x2, y2 float64

	x1 = reg[0].X
	y1 = reg[0].Y

	// This bundle of code iterates over the regression data to find a point with a different x value.
	// One of the datasets provided has n-1 data elements with identical x-values, so this error check is critical
	matched := false
	for i := 1; i < len(reg); i++ {
		if reg[i].X != x1 {
			matched = true
			x2 = reg[i].X
			y2 = reg[i].Y
			break
		}
	}

	if !matched {
		return 0, 0, fmt.Errorf("no match found")
	}

	// These are standard calculations for slope and intercept
	slope := (y2 - y1) / (x2 - x1)
	intercept := y1 - slope*x1

	return intercept, slope, nil
}

// produceSlopeIntercept produces the slope and intercept for a regression number.
func produceSlopeIntercept(regNum int, data []stats.Coordinate) (float64, float64, error) {
	reg, err := stats.LinearRegression(data)
	if err != nil {
		return 0, 0, err
	}

	intercept, slope, err := calcSlopeIntercept(reg)
	if err != nil {
		return 0, 0, err
	}

	fmt.Printf("Regression %d, Int: %f, Slope: %f\n\n", regNum, intercept, slope)
	return intercept, slope, nil
}

// produceAllResults produces the slope and intercept for all regressions.
func produceAllResults() error {
	data1, data2, data3, data4 := makeStatsData()

	_, _, err := produceSlopeIntercept(1, data1)
	if err != nil {
		return err
	}

	_, _, err = produceSlopeIntercept(2, data2)
	if err != nil {
		return err
	}

	_, _, err = produceSlopeIntercept(3, data3)
	if err != nil {
		return err
	}

	_, _, err = produceSlopeIntercept(4, data4)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	startTime := time.Now()
	err := produceAllResults()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	endTime := time.Now()
	timeElapsed := endTime.Sub(startTime)
	fmt.Printf("Time running regressions: %v\n", timeElapsed)
}
