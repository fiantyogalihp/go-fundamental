package main

import (
	"fmt"

	"go-fundamental/database"
	"go-fundamental/helper"
	"go-fundamental/other"

	/*
		  ? terkadang kita perlu me-import sesuatu tpi tidak perlu declare func dan variable nya, hanya proses program nya saja
			? hal ini sering terjadi pada proses connection database menggunakan func init

			? untuk handle issue seperti itu untuk menghindari unused import, dkk. Kita hanya perlu menggunalan blank identifier, yaitu tanda `_` di awal sebelum package tersebut di import

			* contoh nya seperti dibawah ini
	*/
	_ "go-fundamental/engine"
)

func ImportModule() {
	fmt.Println(database.GetDatabase())

	helper.SayHello("eko")
	other.SayHi("Budi")
}
