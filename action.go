package main

import (
	"fmt"
	"bytes"
)

var (
	voyelles = []string{"a","e","i","o","u","y"}
	action_fname = []func(string)(string){
		haha,
		be_letter,
		two_letter,
		third_letter,
		pi_letter,
		voyelle,
		hasard,
	}
)


func haha(a string)(b string){
	b = a
	return
}

func pi(a int)(vrai bool){
	b := a/2
	vrai = false
	if b == 1{
		vrai = true
	}
	return
}

func be_letter(a string)(b string){
	size := len(a)
	how  := much(size)
	b = a[:how]
	if len(b) == 0{
		b = a
	}
	return
}

func two_letter(a string)(b string){
	b = string(a[0]) + string(a[len(a)-1])
	return
}

func third_letter(a string)(b string){
	b = string(a[0])
	size := len(a)
	lol := pi(size)
	middle := size / 2
	if lol  == false && pred()> 0 && size > 5{
		middle = middle + 1;
	}
	if size > middle{
		if size > 3 {
			if pred() > 0 {
				middle = middle + 1
			}
		}
		b = b + string(a[middle])
		b = b + string(a[size-1])
		return
	}
	b = a
	return
}

func pi_letter(a string)(b string){
	if len(a)== 1{
		b = a
		return
	}
	var buffer bytes.Buffer
	var buf bytes.Buffer
	z:= string(a[0])
	zz := string(a[len(a)-1])
	for l := 1;l<len(a)-1;l++{
		buf.WriteString(string(a[l]))
	}
	zzz := buf.String()
	fmt.Println(zzz)
	buffer.WriteString(z)
	if pred() > 0 {
		for i := 0;i<len(zzz);i++{
			if pi(i){
				buffer.WriteString(string(zzz[i]))
			}
		}
	}else{
		for i := 0;i<len(zzz);i++{
			if !pi(i){
				buffer.WriteString(string(zzz[i]))
			}
		}
		
	}
	
	buffer.WriteString(zz)
	b = buffer.String()
	return
}

func ret_be(a string)(b,c string){
	b = string(a[0])
	c = string(a[len(a)-1])
	return
}

func isVoyelle(a string)(b bool){
	b = false
	for _,ok := range voyelles{
		if ok == a {
			b = true
			return
		}
	}
	return
}

func voyelle(a string)(b string){
	bb,e := ret_be(a)
	var buffer bytes.Buffer
	buffer.WriteString(bb)
	for i := 1;i<len(a)-1;i++{
		if !isVoyelle(string(a[i])){
			buffer.WriteString(string(a[i]))
		}
	}
	buffer.WriteString(e)
	b = buffer.String()
	return
}

func hasard(a string)(b string){
	if len(a)== 1{
		b = a
		return
	}
	var buffer bytes.Buffer
	var buf bytes.Buffer
	z:= string(a[0])
	zz := string(a[len(a)-1])
	for l := 1;l<len(a)-1;l++{
		buf.WriteString(string(a[l]))
	}
	zzz := buf.String()
	fmt.Println(zzz)
	buffer.WriteString(z)
	for i := 0;i<len(zzz);i++{
		if pred()>0{
			buffer.WriteString(string(zzz[i]))
		}
	}
	buffer.WriteString(zz)
	b = buffer.String()
	return
	
}
