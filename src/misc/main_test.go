package main

import "testing"


func TestDecode(t *testing.T){
	post, errMsg, err := decode("post.json")

	if err!=nil{
		t.Error("Error wasn't expected!", errMsg, err.Error())
	}
	// post.Content = "Other"
	if post.Content != "Hello World!"{
		t.Error("post.Content must be >> Hello World!")
	}
}