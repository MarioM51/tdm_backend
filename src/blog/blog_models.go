package blog

import (
	"fmt"
	"strconv"
	"time"
	"users_api/src/helpers"

	"gorm.io/gorm"
)

type BlogModel struct {
	Id             int           `json:"id"`
	Title          string        `json:"title" gorm:"unique;not null"`
	Body           string        `json:"body,omitempty" gorm:"not null"`
	Thumbnail      string        `json:"thumbnail,omitempty"`
	Author         string        `json:"author" gorm:"not null"`
	Abstract       string        `json:"abstract" gorm:"not null;size:130"`
	Likes          int           `json:"likes" gorm:"-:all"`
	Comments       []BlogComment `json:"comment,omitempty" gorm:"-:all"`
	CommentCount   int           `json:"commentCount" gorm:"-:all"`
	CommentsRating float32       `json:"comments_rating" gorm:"-:all"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

func (BlogModel) TableName() string {
	return "blogs"
}

func (b BlogModel) validate() string {
	if b.Title == "" {
		return "title is required"
	}

	if b.Body == "" {
		return "body is required"
	}

	return ""
}

type LikeBlog struct {
	FkBlog    int       `json:"fk_blog"`
	FKUser    int       `json:"fk_user"`
	CreatedAt time.Time `json:"created_at"`
}

func (LikeBlog) TableName() string {
	return "blog_likes"
}

type BlogComment struct {
	Id        int            `json:"identifier" gorm:"primaryKey"`
	IdUser    int            `json:"idUser"`
	IdBlog    int            `json:"IdBlog"`
	Text      string         `json:"text"`
	Rating    int            `json:"rating"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (BlogComment) TableName() string {
	return "blog_comments"
}

// json-ld

func BlogModelToArrayJSONLD(blogsM []BlogModel) BlogsWrapperJSONLD {

	var blogListA = []BlogJSONLD{}
	for _, b := range blogsM {
		newBlog := new(BlogJSONLD)
		newBlog.Init(b, true)
		blogListA = append(blogListA, *newBlog)
	}

	return BlogsWrapperJSONLD{
		Val: BlogsJSONLD{
			Context:         "https://schema.org",
			Type:            "ItemList",
			ItemListElement: blogListA,
		},
	}
}

type CommentJSONLD struct {
	Type          string       `json:"@type"`
	Identifier    string       `json:"identifier"`
	IdUser        int          `json:"idUser"`
	Text          string       `json:"text"`
	DatePublished string       `json:"datePublished"`
	ReviewRating  ReviewRating `json:"reviewRating"`
}

func (b BlogComment) toJsonLD() CommentJSONLD {
	reviewTemp := CommentJSONLD{
		Type:          "comment",
		Identifier:    fmt.Sprint(b.Id),
		IdUser:        b.IdUser,
		Text:          b.Text,
		DatePublished: b.CreatedAt.String(),
		ReviewRating: ReviewRating{
			Type:        "Rating",
			BestRating:  5,
			WorstRating: 1,
			RatingValue: uint8(b.Rating),
		},
	}
	return reviewTemp
}

func BlogModelToJSONLDWrapped(b BlogModel) BlogWrapperJSONLD {
	newBlog := new(BlogJSONLD)
	newBlog.Init(b, false)

	wrap := BlogWrapperJSONLD{
		Val: *newBlog,
	}

	return wrap
}

type BlogsWrapperJSONLD struct {
	Val BlogsJSONLD
}

type BlogWrapperJSONLD struct {
	Val BlogJSONLD
}

type BlogsJSONLD struct {
	Context         string       `json:"@context"`
	Type            string       `json:"@type"`
	ItemListElement []BlogJSONLD `json:"itemListElement"`
}

type BlogJSONLD struct {
	Type           string          `json:"@type"`
	Context        string          `json:"@context,omitempty"`
	Identifier     string          `json:"identifier"`
	Headline       string          `json:"headline"`
	Abstract       string          `json:"abstract"`
	ArticleBody    string          `json:"articleBody"`
	DatePublished  string          `json:"datePublished"`
	DateModified   string          `json:"dateModified"`
	Image          string          `json:"image"`
	Author         Author          `json:"author"`
	Comments       []CommentJSONLD `json:"comment"`
	CommentCount   int             `json:"commentCount"`
	CommentsRating float32         `json:"comments_rating"`
	Likes          int             `json:"likes"`
}

func (blogJsonLD *BlogJSONLD) Init(b BlogModel, isInList bool) {
	if !isInList {
		blogJsonLD.Context = "https://schema.org"
	}
	blogJsonLD.Type = "article"
	blogJsonLD.Identifier = strconv.Itoa(b.Id)
	blogJsonLD.Headline = b.Title
	blogJsonLD.Image = fmt.Sprintf(helpers.URL_BLOG_IMG, strconv.Itoa(b.Id), b.UpdatedAt.Format(time.RFC3339))
	blogJsonLD.Abstract = b.Abstract
	blogJsonLD.ArticleBody = b.Body
	blogJsonLD.Likes = b.Likes
	blogJsonLD.DatePublished = b.CreatedAt.Format(time.RFC3339)
	blogJsonLD.DateModified = b.UpdatedAt.Format(time.RFC3339)
	blogJsonLD.CommentCount = b.CommentCount
	blogJsonLD.CommentsRating = b.CommentsRating
	blogJsonLD.Author = Author{
		Type: "Person",
		Name: b.Author,
	}

	commentsJsonLD := []CommentJSONLD{}
	for i := range b.Comments {
		reviewTemp := b.Comments[i].toJsonLD()
		commentsJsonLD = append(commentsJsonLD, reviewTemp)
	}
	blogJsonLD.Comments = commentsJsonLD
}

type Author struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}

type ReviewRating struct {
	Type        string `json:"@type"`
	BestRating  uint8  `json:"bestRating"`
	RatingValue uint8  `json:"ratingValue"`
	WorstRating uint8  `json:"worstRating"`
}
