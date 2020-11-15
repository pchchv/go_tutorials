package main

import ("fmt"; "math")

type Circle struct {
    x, y, r float64
}
func (c *Circle) area() float64 {
    return math.Pi * math.Pow(c.r, 2)
}
func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - y1
    b := y2 - y1
    return math.Sqrt(a * a + b * b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}/*
func circleArea(c *Circle) float64 {
    return math.Pi * math.Pow(c.r, 2)
}*/
func main() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10
    c := Circle{0, 0, 5}

    fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
    fmt.Println(c.area())
}