package main

import "log"

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
	map_ls()
}
