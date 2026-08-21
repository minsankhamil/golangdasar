package helper

//fungsi SayHello dapat diakses package lain,
// jika ingin diprivate(tidak bisa diakses package lain) ubah menjadi sayHello(huruf pertama kecil)
func SayHello(name string) string {
	return "hello " + name
}
