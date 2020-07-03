/*Đọc file `serve.json` đó ra và tạo 1 struct có tên là `Serve` để đưa nội dung trên vào dạng cấu trúc đã định sẵn. in ra màn hình kết quả với `log`.
Gợi ý nó là dạng slice
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type serve struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func readfile(filename string) []serve {
	dat, _ := ioutil.ReadFile(filename)
	sv := make([]serve, 2)
	_ = json.Unmarshal(dat, &sv)
	for _, i := range sv {
		log.Println("name:", i.Name)
		log.Println("class:", i.Class)
	}
	return sv
}

//3. Thực hiện lọc kết quả với class có chứa từ admin in ra đối tượng phù hợp.
func locAdmin(ad string, sv []serve) {
	for _, i := range sv {
		if i.Class == ad {
			log.Println("name:", i.Name)
			log.Println("class:", i.Class)
		}

	}
}

/*4. Thêm một đối tượng mới vào slice vừa trên:
{
	"name": "fileCustome",
	"class": "org.cofax.cds.FileServlet.Custome"
}
*/
func addserve(name string, class string, sv []serve) {
	s := serve{}
	s.Class = class
	s.Name = name
	sv = append(sv, s)
	log.Println(sv)
}

// 5 .In ra màn hình các địa chỉ của các item strong slice trên

func inaddress(sv []serve) {
	for i := range sv {
		fmt.Println(&i)
	}
}
