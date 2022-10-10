package main

import (
	"fmt"
	"nicerslicermain/nicerslicer"
)


func main() {

	str := `Don't communicate by sharing memory - share memory by communicating`

	sl := nicerslicer.NewNicerSlicer(str)
	var ncs []nicerslicer.NicerSlicer = sl.Split("-")

	fmt.Printf(`%v`, ncs[3])
}
