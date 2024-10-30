package post

type User struct {
	Id       int
	Username string
	Password string
	Posts    []Post
}

type Post struct {
	Id      int
	Title   string
	Content string
	UserId  int
	User    User
}
