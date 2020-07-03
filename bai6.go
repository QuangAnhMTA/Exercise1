/*
Khởi tạo slice có nội dung như sau:

11,34,56,77,99,109,66,20, 88, 34
Sắp xếp slice trên theo thứ tự không giảm. In ra màn hình kết quả.
Tạo 1 slice (copy slice) có index từ 1 đến 7.
Copy slice trên với index từ 1 đến 15.(có lỗi ko?)
*/
package main

import (
	"fmt"
	"sort"
)

func sl(leng int) {
	s := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	sort.Ints(s)
	for _, i := range s {
		fmt.Println(i)
	}
	if len(s) < leng {
		fmt.Printf("Ban khong the copy slice co length lon hon %d \n", len(s))
	} else {
		l := s[1:leng]
		fmt.Println("l:", l)
	}

}
