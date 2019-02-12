package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// Db is a Global Database connection
var Db *sql.DB

func init(){
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
}

func createUUID()(uuid string){
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil{
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Encrypt encypts a string with sha1 algorithm
func Encrypt(plaintext string)(cryptedtext string){
	cryptedtext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
