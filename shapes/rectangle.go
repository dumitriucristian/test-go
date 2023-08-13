package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height + r.Width
}

func (r Rectangle) Perimeter() float64 {
	return r.Height * r.Width
}
