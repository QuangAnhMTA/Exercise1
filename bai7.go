/*
Tạo 1 chương trình golang với 1 struct như sau.
type User struct {
  name string
  age int64
  gender bool
  address string
}
a. Tạo các getter của struct trên.(Vd: func (u*User) GetName() string) sử dụng method
b. Tạo 1 map[string]*User với key là name và value là đối tượng user.
Nhập 10 giá trị cho slice trên. Lấy ra giá trị của map trên nạp vào 1 slice. Inra kết quả.

*/
package main

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func (u *User) getname() string {
	return u.name
}
func (u *User) getAddress() string {
	return u.address
}
func map_ls() {
	m := make(map[string]*User)

}
