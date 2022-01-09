package calc

type ICalc interface {
	Add(x, y float64) float64
	Subs(x, y float64) float64
}
type Calc struct{}

func (c *Calc) Add(x, y float64) float64 {
	return x + y
}

func (c *Calc) Subs(x, y float64) float64 {
	return x - y
}
