package main

import "fmt"

type maps map[string]string

type people struct {
	Name map[int]int
}

func main()  {
	var people people
	fmt.Println("%t",people.Name)
}