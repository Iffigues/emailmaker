package main

import (
	"bufio"
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
	a = a % 4
	switch a {
	case 1:
		c.fname = append(c.fname,b)
	case 2:
		c.lname = append(c.lname,b)
	case 3:
		c.domaine = append(c.domaine,b)
	case 0:
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
	reader := bufio.NewReader(os.Stdin)
	var text string
	for text != "!"{
		text, _ = reader.ReadString('\n')
		text = strings.Trim(text,"\n")
		b = append(b,text)
	}
	return
}

func make()(t Argv,err error){
	act := ""
	var tt = 1;
	for _,ok := range arg {
		if ok[0] == '-'{
			if act != ""{
				zz := output()
				for _,ak := range zz{
					typers(act,ak,&t)
				}
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
	if err != nil{
		fmt.Println(err)
	}
}

func main(){
	argvs(os.Args[1:])
}
