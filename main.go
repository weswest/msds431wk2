// Hat tip to https://gitlab.com/devthoughts/code/tree/master/linear-regression-with-go

package main

import (
	"fmt"
	"math"

	"github.com/sajari/regression"
)

var x1 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
var x2 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
var x3 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
var x4 = []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
var y1 = []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
var y2 = []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
var y3 = []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
var y4 = []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

func runRegression(verbose bool, rNum int, x []float64, y []float64) (float64, float64) {

	// Build the regression
	var r regression.Regression
	r.SetObserved("y")
	r.SetVar(0, "x")

	// Capture the data points in the regression item
	for i := range x {
		r.Train(regression.DataPoint(y[i], []float64{x[i]}))
	}

	// Train/fit the regression model.
	r.Run()

	b0, b1 := r.Coeff(0), r.Coeff(1)
	b0Round := math.Round(b0*100) / 100
	b1Round := math.Round(b1*100) / 100

	// Output the trained model parameters.
	if verbose {

		fmt.Printf("\nRegression %d Formula:\n%v\n\n", rNum, r.Formula)
		fmt.Printf("Int: %f, Slope: %f\n\n", b0Round, b1Round)
	}

	return b0Round, b1Round

}
func runAllRegressions(verbose bool) {
	runRegression(verbose, 1, x1, y1)
	runRegression(verbose, 2, x2, y2)
	runRegression(verbose, 3, x3, y3)
	runRegression(verbose, 4, x4, y4)
}

func main() {
	runAllRegressions(true)
}
