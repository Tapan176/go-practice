package entityschema

type SessionUser struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Role         string `json:"role"`
	IsVerified   bool   `json:"isVerified"`
	PasswordHash string `json:"-"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type Post struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	UserId     string `json:"userId"`
	CategoryId string `json:"categoryId"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	Likes      int    `json:"likes"`
	Dislikes   int    `json:"dislikes"`
}

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type LikedBlogs struct {
	BlogID int `json:"blogId"`
	UserID int `json:"userId"`
}

type DislikedBlogs struct {
	BlogID int `json:"blogId"`
	UserID int `json:"userId"`
}

type LikedAndDisliked struct {
	UserID        int `json:"userId"`
	BlogsLiked    int `json:"blogsLiked"`
	BlogsDisliked int `json:"blogsDisliked"`
}

type AuthorRating struct {
	ID       int    `json:"id"`
	AuthorID string `json:"authorId"`
	UserID   string `json:"userId"`
	Rating   int    `json:"rating"`
}

type Comment struct {
	ID        int          `json:"id"`
	ArticleId int          `json:"articleId"`
	UserId    int          `json:"userId"`
	Comment   string       `json:"comment"`
	CreatedAt string       `json:"createdAt"`
	UpdatedAt string       `json:"updatedAt"`
	Likes     int          `json:"likes"`
	Dislikes  int          `json:"dislikes"`
	Reply     CommentReply `json:"reply"`
}

type CommentReply struct {
	ID        int    `json:"id"`
	CommentId int    `json:"commentId"`
	UserId    int    `json:"userId"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Likes     int    `json:"likes"`
	Dislikes  int    `json:"dislikes"`
}
