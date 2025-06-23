package main

import "fmt"

const  statusOk = 200

const (
	statusNotFound = 404
	statusCreated = 201
	statusForbidden = 403
)

func main(){
	fmt.Println("Ok: ",statusOk)
	fmt.Println("Not Found : ",statusNotFound)
	fmt.Println("Created : ",statusCreated)
	fmt.Println("Forbidden : " ,statusForbidden)
}