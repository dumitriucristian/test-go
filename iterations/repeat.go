package iterations

//Repeat the first param as many times the second params requires
func Repeat(repeat string, nrOfRepetitions int) string {

	var repeated string

	for i := 0; i < nrOfRepetitions; i++ {

		repeated += repeat
	}

	return repeated
}
