# betterpath

The goal of this project is to develop a program that, given a series of points (latitude, longitude, timestamp) for a courier journey from A-B, will disregard potentially erroneous points

## Usage
```
go build cmd/betterpath.go
./betterpath assets/points.csv
```
## Issues
#### Understanding the problem
The first thing to do was to spot the erroneous points from the dataset given with the subject.<br/>
I plotted the coordinates on a map and noticed that some points were far from the main path:<br/>

  ![Alt text](assets//uncleaned-journey.png?raw=true "Uncleaned path")<br/>

#### Heuristics
  The second thing to do was to identify which aspects of these points made them erroneous.<br/>
  I chose to work on the distances between the points, rather than the speed of the courier because it would have required trigonometry computations and a greater overhead.<br/> 
  Since the "wrong" distances seemed significantly higher than the "normal" ones I decided to work on the standard deviation of the distances, inspired by the work I did on [linear](https://github.com/ljoly/ft_linear_regression) and [logistic](https://github.com/ljoly/DSLR) regressions and an interesting [article](https://www.kdnuggets.com/2017/02/removing-outliers-standard-deviation-python.html) about the concept of normal distribution.<br/>

#### Position of the points
Another matter was the position of the wrong points after being sorted by timestamp. Wether they were contiguous, at the beginning of the path or at the end, it required to update carefully the data stored.

## Design decisions
Two slices:
* for storing the coordinates and the timestamp of each point, using a structure Point
```go
type Point struct {
	x float64
	y float64
	t int64
}
```

* for storing the distances between points and to calculate the standard deviation

The two slices are browsed only once, and simultaneously. When a distance is considered too high (ie. superior to the standard deviation), the erroneous point of the pair is identified, deleted, and the distances are updated (and checked because two erroneous points can be contiguous).<br/>


