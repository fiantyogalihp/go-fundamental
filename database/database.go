package database

var connection string

// * func init adalah func yang pertama kali dan selalu dijalankan, sebelum perintah lainnya dieksekusi
// * sama seperti hal nya contructor dalam class
func init() {
	connection = "MySQL"
	// fmt.Println("init dipanggil") // ! hal ini dilakukan ketika kita ingin tahu func init nya berjalan atau tidak
}

func GetDatabase() string {
	return connection
}
