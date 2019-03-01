package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
	}
type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
	}
type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
	}
	

func main(){
	fileJson, err := os.Open("post.json")
	if err!=nil{
		fmt.Println("Error openening file:", err)
		return
	}
	defer fileJson.Close()
	jsonData, err := ioutil.ReadAll(fileJson)
	if err!=nil{
		fmt.Println("Error getting the data", err)
		return
	}

	var post Post

	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
	data, err := json.Marshal(post)
	fmt.Println(string(data))
}