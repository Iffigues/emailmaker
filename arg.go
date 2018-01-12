package main

type act struct{
	cond []string
	action []string
}

type file_fname struct{
	name string
	fors []string
	actions []act
}

type skelet struct {
	skel_fname []file_fname `json:"fname"`
}
