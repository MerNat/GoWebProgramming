package data

import (
	"time"
	_"math/big"
)
// User defines a user structure
type User struct{
	Id	uint32
	Uuid	string
	Name	string
	Email	string
	Password	string
	CreatedAt	time.Time	
}

type Session struct{
	Id 	uint32
	Uuid	string
	Email	string
	UserId	uint32
	CreatedAt	time.Time
}

// Create a new session for an existing user
func (user *User) createSession()(session Session, err error){
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id,&session.Uuid,&session.Email,&session.CreatedAt)
	return
}

// Session Method Gets the session for an existing user
func (user *User) Session()(session Session, err error){
	session = Session{}
	statement := "SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1"
	err = Db.QueryRow(statement,user.Id).Scan(&session.Id,&session.UserId,&session.Email,&session.UserId,&session.CreatedAt)
	return
}

// Check if session is valid in the db
func (session *Session) Check()(valid bool, err error){
	statement := "SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1"
	err = Db.QueryRow(statement,session.Uuid).Scan(&session.Id,&session.UserId,&session.Email,&session.UserId,&session.CreatedAt)
	if err!= nil{
		valid = false
		return
	}
	if session.Id != 0{
		valid = true
	}
	return
}

// Delete session from database

func (session *Session) DeleteByUUID()(err error){
	statement := "delete from sessions where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

// Get the user from the session
func (session *Session) User() (user User, err error){
	statement := "SELECT id, uuid, name, email, created_at FROM users WHERE id=$1"
	err = Db.QueryRow(statement,session.UserId).Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.CreatedAt)
	return
}

// Delete all sessions from database
func SessionDeleteAll()(err error){
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error){
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)

	if err != nil{
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(createUUID, user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id,&user.Uuid,&user.CreatedAt)
	return
}