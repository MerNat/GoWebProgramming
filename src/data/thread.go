package data

import(
	"time"
)

type Thread struct{
	Id	uint32
	Uuid	string
	Topic	string
	UserId	uint32
	CreatedAt	time.Time
}

type Post struct{
	Id	uint32
	Uuid	string
	Body	string
	UserId	uint32
	ThreadId	uint32
	CreatedAt	time.Time
}

// Formatting CreatedAt field 
func(thread *Thread) CreatedAtDate() string{
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// CreatedAtDate formats date of a Post
func (post *Post) CreatedAtDate() string{
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// CreateThread creates a Thread for a user
func (user *User) CreateThread(topic string) (conv Thread, err error){
	statement := "insert into threads (uuid, topic, user_id, created_at) values ($1,$2,$3,$4) returning id, uuid, topic, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err!=nil{
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(),topic,user.Id,time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// CreatePost creates a new post to a thread
func (user *User) CreatePost(conv Thread, body string)(thePost Post,err error){
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1,$2,$3,$4,$5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(),body,user.Id,conv.Id,time.Now()).Scan(&thePost.Id,&thePost.Uuid,&thePost.Body,&thePost.UserId,&thePost.ThreadId,&thePost.CreatedAt)
	return
}
// Threads returns all threads from the database
func Threads()(threads []Thread, err error){
	rows, err := Db.Query("SELECT id, uuid, topic,user_id,created_at FROM threads ORDER BY created_at DESC")
	if err != nil{
		return
	}
	for rows.Next(){
		conv :=  Thread{}
		if err = rows.Scan(&conv.Id,&conv.Uuid,&conv.Topic,&conv.UserId,&conv.CreatedAt); err != nil{
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

// User method of Thread returns the User who created the thread
func (thread *Thread) User() (user User){
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id=$1",thread.UserId).
	Scan(&user.Id ,&user.Uuid ,&user.Name ,&user.Email ,&user.CreatedAt)
	return
}
// User method of Post returns the User who created the post
func (post *Post) User()(user User){
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id=$1",post.UserId).
	Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.CreatedAt)
	return
}

// ThreadByUUID gets a thread by the UUID
func ThreadByUUID(uuid string)(conv Thread, err error){
	conv = Thread{}
	Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid=$1",uuid).
	Scan(&conv.Id,&conv.Uuid,&conv.Topic,&conv.UserId,&conv.CreatedAt)
	return
}

// Posts Method returns posts associated with a thread
func (thread *Thread) Posts()(posts []Post, err error){
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_id=$1",thread.Id)
	if err != nil{
		return
	}
	defer rows.Close()
	for rows.Next(){
		post := Post{}
		if err = rows.Scan(&post.Id,&post.Uuid,&post.Body,&post.UserId,&post.ThreadId,&post.CreatedAt); err != nil{
			return
		}
		posts = append(posts, post)
	}
	return
}


// NumReplies Method returns number of posts of a specific thread
func (thread *Thread) NumReplies()(count int){
	rows, err := Db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", thread.Id)
	if err != nil{
		return
	}
	defer rows.Close()
	for rows.Next(){
		if err = rows.Scan(&count); err!=nil{
			return
		}
	}
	return
}