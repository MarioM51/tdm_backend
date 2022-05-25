package blog

import (
	"strconv"
	"time"
)

type BlogModel struct {
	Id        int       `json:"id"`
	Title     string    `json:"title" gorm:"unique;not null"`
	Body      string    `json:"body,omitempty" gorm:"not null"`
	Thumbnail string    `json:"thumbnail,omitempty"`
	Author    string    `json:"author" gorm:"not null"`
	Abstract  string    `json:"abstract" gorm:"not null;size:130"`
	Likes     int       `json:"likes" gorm:"-:all"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

// json-ld

func BlogModelToArrayJSONLD(blogsM []BlogModel) BlogsWrapperJSONLD {

	var blogListA = []BlogJSONLD{}
	for _, b := range blogsM {
		newBlog := blogModelToJSONLD(b)
		blogListA = append(blogListA, newBlog)
	}

	return BlogsWrapperJSONLD{
		Val: BlogsJSONLD{
			Context:         "https://schema.org",
			Type:            "ItemList",
			ItemListElement: blogListA,
		},
	}
}

func blogModelToJSONLD(b BlogModel) BlogJSONLD {

	blog := BlogJSONLD{
		Type:          "BlogPosting",
		Identifier:    strconv.Itoa(b.Id),
		Headline:      b.Title,
		Image:         strconv.Itoa(b.Id),
		Abstract:      b.Abstract,
		ArticleBody:   b.Body,
		Likes:         b.Likes,
		DatePublished: b.CreatedAt.Format(time.RFC3339),
		DateModified:  b.UpdatedAt.Format(time.RFC3339),
		Author: Author{
			Type: "Person",
			Name: b.Author,
		},
	}

	return blog

}

func BlogModelToJSONLDWrapped(b BlogModel) BlogWrapperJSONLD {

	newBlog := blogModelToJSONLD(b)

	wrap := BlogWrapperJSONLD{
		Val: newBlog,
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
	Type          string `json:"@type"`
	Identifier    string `json:"identifier"`
	Headline      string `json:"headline"`
	Abstract      string `json:"abstract"`
	ArticleBody   string `json:"articleBody"`
	DatePublished string `json:"datePublished"`
	DateModified  string `json:"dateModified"`
	Image         string `json:"image"`
	Author        Author `json:"autor"`
	Likes         int    `json:"likes"`
}

type Author struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}
