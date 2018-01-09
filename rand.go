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



func pred()(r int){
	aa := rand.New(src)
	return aa.Intn(2)
}
