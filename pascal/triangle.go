package main

import "fmt"

type Row []int

type Triangle struct {
	Rows []Row
}

// triangle contains a bunch of rows. each row is a slice. function needs to find the all the rows that prceed N and do the pascal math to return the N row.

func main() {
	var TargetLine int
	fmt.Print("Enter the row number you want : ")
	fmt.Scan(&TargetLine)
	emptyTriangle := initTriangle(TargetLine)
	newTriangle := buildTriangle(TargetLine, emptyTriangle)
	for k, v := range newTriangle.Rows {
		if k == len(newTriangle.Rows)-1 {
			fmt.Printf("%v <------- What you Requested\n", v)
		} else {
			fmt.Println(v)
			fmt.Print("\n")
		}
	}

}

func buildTriangle(numRows int, triangle *Triangle) *Triangle {
	var val int = 1
	for i := 0; i < numRows; i++ {
		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				val = 1
			} else {
				val = val * (i - k + 1) / k
			}
			triangle.Rows[i] = append(triangle.Rows[i], val)
		}
	}
	return triangle
}

func initTriangle(numRows int) *Triangle {
	triangle := Triangle{}
	triangle.Rows = make([]Row, numRows)
	return &triangle
}
