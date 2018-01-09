package main

import (
	"os"
	"fmt"
)

var (
	count = 0
	option = []string{"n","f","l","d","c"}
	options = []string{"fname","lname","number","domaine","country"}
	arg []string
)

type Argv struct{
	nbr int
	lname []string
	fname []string
	domaine []string
	country []string
	mail []email
	Err error
}

func what(a string)(r string){
	return
}

func typer(a int,b string,c *Argv){
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("")
	case 5:
		fmt.Println("")
	}
}

func typers(a, b string,c *Argv){
	switch a {
	case "a":
		fmt.Println("one")
	case "b":
		fmt.Println("two")
	case "c":
		fmt.Println("three")
	case "d":
		fmt.Println("")
	case "e":
		fmt.Println("")
	}
	
}

func make()(t Argv,err error){
	act := ""
	var tt = 0;
	for _,ok := range arg {
		if ok[0] == '-'{
			act = what(ok)
		}
		if act == ""{
			typer(tt,ok,&t)
			tt = tt+1
		}else{
			typers(act,ok,&t)
			act = ""
		}
	}
	if act != "" {
		
	}
	return
}

func argvs(ar []string){
	arg = ar
	count = len(ar)
	make()
}

func main(){
	argvs(os.Args[1:])
	fmt.Println(len(arg),pred())
}
