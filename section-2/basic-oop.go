package main

import "fmt"

// Class/Object Account
type Account struct {
	nama    string // private
	Username string // public
}


func (u *Account) SetNama(n string) {
	u.nama = n
}


func (u Account) GetNama() string {
	return u.nama
}

func main() {
	u := Account{Username: "rakahikmah"}
	u.SetNama("Raka Hikmah Ramadhan")

	fmt.Println("Username:", u.Username)
	fmt.Println("Nama lengkap:", u.GetNama()) 
	fmt.Println(u.nama)                     
}