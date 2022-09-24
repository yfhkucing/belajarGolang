// ini golang
// import ini wkwkwk
package main //nama package tu nama programnya?
import "fmt" //buat print
//mesti dalam function main baru jalan (mirip void loop di arduino)
//Print pake p besar
//sejauh ini mirip python campur cpp(?)
//di compile di terminal pake go build <nama program>, jadi file .exe (nice)

func main() {
	//contoh variabel
	//tipe data golang : https://www.geeksforgeeks.org/data-types-in-go/
	//dynamic jadi ga harus dideklarasiin tipe datanya ( kecuali pas jadi array )
	//setiap variabel yg dideklarasiin harus dipake, kalo gk error (forced to write a clean code,damn golang!)
	//deklarasi variabel sekaligus
	var (
		angka         = 1
		iniString     = "ini string"
		iniBoolean    = true
		iniStringJuga string
		iniArray      = [...]int{1, 2, 3, 4}
	)
	fmt.Println(angka)
	fmt.Println(iniString)
	fmt.Println(iniBoolean)
	fmt.Println(iniStringJuga)
	fmt.Println(iniArray)
	loop()
}

func loop() {
	sum := 0
	//range 1 sampai kurang dari 5, mirip loop di cpp tapi ga pake kurung
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)
}
