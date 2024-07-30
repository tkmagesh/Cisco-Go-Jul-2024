package main

import "fmt"

func main() {
	/*
		var pen struct {
			Id   int
			Name string
			Cost float64
		}
		// fmt.Println(pen)
		pen.Id = 100
		pen.Name = "Fountain Pen"
		pen.Cost = 10
		// fmt.Printf("%#v\n", pen)
		fmt.Printf("%+v\n", pen)
	*/

	var pen struct {
		Id   int
		Name string
		Cost float64
	} = struct {
		Id   int
		Name string
		Cost float64
	}{
		Id:   100,
		Name: "Fountain Pen",
		Cost: 10,
	}
	// fmt.Printf("%+v\n", pen)
	fmt.Println(Format(pen))

}

func Format(p struct {
	Id   int
	Name string
	Cost float64
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}
