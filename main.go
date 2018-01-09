package main

import (
	"os"
	"fmt"
	"strings"
	"errors"
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

func split(a string)(b []string){
	b = strings.Split(a," ")
	return
}

func trim(a string)(b string){
	b = strings.Trim(a,"-")
	return
}


func what(a string)(r string){
	return
}

func typer(a int,b string,c *Argv){
	switch a {
	case 0:
		c.fname = append(c.fname,b)
	case 1:
		c.lname = append(c.lname,b)
	case 2:
		c.fname = append(c.domaine,b)
	case 3:
		c.country = append(c.country,b)
	}
}

func typers(a, b string,c *Argv){
	g := trim(a)
	r := split(b)
	switch g {
	case "a":
		fmt.Println(r)
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

func output()(b []string){

}

func make()(t Argv,err error){
	act := ""
	var tt = 0;
	for _,ok := range arg {
		if ok[0] == '-'{
			if act != ""{
				zz := output()
			}
			act = ok
		}else{
			if act == ""{
				typer(tt,ok,&t)
				tt = tt+1
			}else{
				typers(act,ok,&t)
				act = ""
			}
		}
	}
	if act != "" {
		err = errors.New("must be no finish by options")
	}
	fmt.Println(t)
	return
}

func argvs(ar []string){
	arg = ar
	count = len(ar)
	_,err := make()
	fmt.Println(err)
}

func main(){
	argvs(os.Args[1:])
	fmt.Println(len(arg),pred())
}
