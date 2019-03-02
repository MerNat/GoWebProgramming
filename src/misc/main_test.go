package main

import "testing"


func TestDecode(t *testing.T){
	errMsg, err := decode("post.json")

	if err!=nil{
		t.Error("Error wasn't expected!", errMsg, err.Error())
	}
}