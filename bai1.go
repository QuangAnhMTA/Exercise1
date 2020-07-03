package main

import (
	"io"
	"os"
)

// 1. Tạo 1 file có nội dung như sau: ghi ra file có tên là serve.json
/*[
    {
      "name": "cofaxAdmin",
      "class": "org.cofax.cds.AdminServlet"
    },
    {
      "name": "fileServlet",
      "class": "org.cofax.cds.FileServlet"
    }
]
*/
func writetofile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return nil
}
