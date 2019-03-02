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

func decode(fileName string) (post Post, errMsg string, err error){
	fileJson, err := os.Open(fileName)
	if err!=nil{
		errMsg = "Error opening file"
		return
	}
	defer fileJson.Close()
	jsonData, err := ioutil.ReadAll(fileJson)
	if err!=nil{
		errMsg = "Error getting the data"
		return
	}

	// var post Post

	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
	data, err := json.Marshal(post)
	fmt.Println(string(data))

	return
}

func main(){
	_, errMsg, err := decode("post.json")

	if err!=nil{
		fmt.Printf("%s: %s\n", errMsg, err.Error())
	}
}