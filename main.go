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
	err := writetofile("serve.js", str)
	if err != nil {
		log.Fatal(err)
	}
	sv := make([]serve, 2)
	sv = readfile("serve.js")
	locAdmin("org.cofax.cds.FileServlet", sv)
	addserve("fileCustome", "org.cofax.cds.FileServlet.Custome", sv)
	inaddress(sv)
	sl()
	map_ls()
}
