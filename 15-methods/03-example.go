/* creating a method for a non-local type */

package main

import "fmt"

/*
func (s string) Length() int {
	return len(s)
}

func main() {
	s := "Qui in enim quis laboris cillum deserunt voluptate eiusmod cupidatat cupidatat ullamco nisi occaecat culpa."
	fmt.Println(s.Length())
}
*/

type MyString string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	s := MyString("Qui in enim quis laboris cillum deserunt voluptate eiusmod cupidatat cupidatat ullamco nisi occaecat culpa.")
	fmt.Println(s.Length())
}
