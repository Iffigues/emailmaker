package main

import(
	"math/rand"
	"time"
)

var(
	src = rand.NewSource(int64(time.Now().Local().Nanosecond()))
)

func nbrl(a string)(b int){
	return
}


func much(b int)(r int){
	aa := rand.New(src)
	return aa.Intn(b)
}

func pred()(r int){
	aa := rand.New(src)
	return aa.Intn(2)
}

func prand(a []func(string)(string))(b int){
	aa := rand.New(src)
	b = aa.Intn(len(a))
	return
}

