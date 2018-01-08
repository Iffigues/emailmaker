package main

import (
	"os"
	"fmt"
)

var (
	count = 0
	nbr = 1
	domain  string
	option = []string{"n","f","l","d","c"}
	options = []string{"fname","lname","number","domaine","country"}
	arg  []string
	mail  []email
)

type Argv struct{
	nbr int
	types int
	domaine string
	mail []email
}


func make(){
	for _,ok := range arg {
		fmt.Println(ok)
	}
}

func argvs(ar []string){
	arg = ar
	count = len(ar)
	make()
}

func main(){
	argvs(os.Args[1:])
	fmt.Println(len(arg))
}
