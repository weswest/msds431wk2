# msds431wk2 - Should Data Scientists use Go?

Class assignment for Week 2 of Northwestern's MSDS-431 Intro to Go class, Summer 2023 term.

## How to Use This Program

This is a deterministic program that takes no inputs and produces a standard set of outputs.  When you run the executable the code performs the following steps:

1. Reads in the anscombe datasets
2. Implicitly generates an OLS regression to calculate predicted y values for each x,y pair
3. Reverse-engineers the coefficients of the OLS regression
4. Prints the results of the regression.

Spoiler alert: all four of the Anscombe datasets are predicted using Y = 3 + 0.5*X, although they all look very different.

The python and R scripts are provided for the purposes of execution time benchmarking and to confirm model accuracy.

The program is dependent on the montanaflynn/stats project.

Note that msds431wk2 was compiled to run on a mac, but the course submission form won't let me submit it.  msds431wk2.exe was compiled to run on a Windows machine, but hasn't been tested for completeness.  Seeing as it was compiled from identical source I have to assume it is working.

## tl;dr: Let's not Use Go for Data Science

The purpose of this assignment is to run linear regressions on four standard datasets using Go and compare the performance (e.g., time to execute) and the accuracy (e.g., regression coefficients) against similar work performed in Python and R.  This gets us in the mind of data scientists at an organization that wants to see employees using Go as their primary programming language, and allows us to provide a recommendation: should data scientists use Go?

Answer: **We should not use Go exclusively for our Data Science work.**

The regressions in Go executed significantly faster than R or Python, and the regression coefficients were accurate enough compared to the Python and R results. However, Python and R are mature platforms with a wealth of tools designed to simplify and streamline the data science part of a data science pipeline. Go lacks such tools. The statistics package used for benchmarking had limited functionality, lacking important features such as coefficient confidence intervals, R-squared, and the ability to generate new predicted y values for unseen x values. Even basic model evaluation metrics and tests like heteroskedasticity and kurtosis require building them from scratch in Go.

If we want our data science group to spend their time recreating from scratch the baseline metrics for model evaluation, then we can consider Go. However, starting from scratch in Go would require significant effort and likely thousands of person-hours to code a foundation on which we can build. Furthermore, our data scientists are not primarily skilled in programming, which may result in a suboptimal foundation being built.

## Background on the Data

The datasets used for testing are the Anscombe data, which contain the interesting property that while they all look very different, they are all equally predicted by the linear regression Y = 3 + 0.5*X:

![anscombe](/img/fig_anscombe.png)

## Verifying Accuracy of Go

Going into this exercise I did not expect the Go code to provide a precisely consistent OLS equation, since all of these approaches rely on gradient descent to find a "good enough" equation within some boundary of error.

With that as a premise, all three systems land on Y = 3 + 0.5*X with a rounding precision of two decimal places.

### Comparison of Intercepts

|Dataset|Python|R|Go|
|:----|:----|:----|:----|
|Set I|3.0001|3.0001|3.0001|
|Set II|3.0009|3.0001|3.0009|
|Set III|3.0025|3.0025|3.0025|
|Set IV|3.0017|3.0017|3.0017|

### Comparison of Slopes

|Dataset|Python|R|Go|
|:----|:----|:----|:----|
|Set I|0.5001|0.5001|0.5001|
|Set II|0.5000|0.5000|0.5000|
|Set III|0.4997|0.4997|0.4997|
|Set IV|0.4999|0.4999|0.4999|

## Verifying Speed of Go

All three programs were speed benchmarked by capturing the system time as early as possible in the main program, capturing the system time after execution / printing, and then calculating the difference.  This provides the closest apples-to-apples comparison of performance speed.

As we can see, Go is exponentially faster than Python or R.

Per the formal requirements, the program was also benchmarked using the go test command.  Full execution of the program in Go is comparable to the slimmed down speedtests of Python or R.

|Program|Execution time (sec)|
|:----|:----|
|Python|3.11|
|R|0.01|
|Go - in-program|0.000073|
|Go - go test| 0.19

## Je Ne Sais Quoi of Go

How did it feel to program in Go?  Terrible.  Go is commonly used in high performant functions under very heavy loads and it excels at doing that.  But Go is not good for data science:
* Go is not good for analytical exploration.  There are rigorous rules to ensure all imported packages and declard variables are used, and that everything works well.  Data science - as shown by the popularity of Jupyter notebooks - is substantially more freeform.  The formalization of needing to produce an optimized package prior to testing a bit of code will ultimately slow down the data scientist
* Go doesn't have the infrastructure to streamline a data scientist's workflows.  
    * For example, even for the incredibly basic case of validating a regression equation I had to create a bespoke method simply to reverse-engineer the linear equation the recommended stats package generated
    * And that's before we get into all of the other statistical testing for which I would have to write my own functions (adjusted R-squared, kurtosis, heteroskedasticity, confidence intervals, etc).  There is no way a data science team could switch to Go without many thousands of coder-hours invested to create a development suite
    * This point gets exponentially harder as we move out of simple univariate OLS into multivariate linear regressions, more sophisticated regressions, and machine learning applications
* Data scientists are selected for their analytical ability, not their coding ability, which makes Go a non-starter.  I like to think I'm a moderately talented amateur coder and I still spent more than half of my time on this assignment fighting with github and VS Code to manage packages and dependencies to get otherwise functional code to execute.  The level of computer science understanding to write a basic "hello, world" script in Go is higher than the majority of data scientists.

There's a thesis of using the right tool for the job.  If we consider the three stages of model development as first managing data, then exploring the data and testing models, then deploying the models, then Go is great at the first and third items.  But if this hypothetical company's leadership wants to use only one language then it shouldn't be Go.
