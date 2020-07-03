package main

import (
	"fmt"
	"log"
)

func main() {
	str :=
		`
	[
		{
		  "name": "cofaxAdmin",
		  "class": "org.cofax.cds.AdminServlet"
		},
		{
		  "name": "fileServlet",
		  "class": "org.cofax.cds.FileServlet"
		}
	]`
	err := writetofile("serve.json", str)
	if err != nil {
		log.Fatal(err)
	}
	sv := make([]serve, 2)
	sv = readfile("serve.json")
	locAdmin("admin", sv)
	addserve("fileCustome", "org.cofax.cds.FileServlet.Custome", sv)
	inaddress(sv)
	sl(15)
	s := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	ls1, err1 := copyslices(s, 1, 7)
	fmt.Println(ls1)
	fmt.Println(err1)
	ls2, err2 := copyslices(s, 1, 15)
	fmt.Println(ls2)
	fmt.Println(err2)
	map_ls()
}
