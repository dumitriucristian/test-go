package compare

type name string

type location struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

type dot struct {
	name
	location
	size
}

func getDots() []dot {
	var dot1 dot
	dot2 := dot{}

	dot2.name = "A"
	dot2.location.x = 5
	dot2.location.y = 6
	dot2.size.width = 1
	dot2.size.height = 2

	return []dot{dot1, dot2}

}
