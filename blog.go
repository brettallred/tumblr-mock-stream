package main

type Blog struct {
	Created     int64  `json:"created"`
	Updated     int64  `json:"updated"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsGroupBlog bool   `json:"is_group_blog"`
	IsPrimary   bool   `json:"is_primary"`
	Timezone    string `json:"timezone"`
	PostCount   int    `json:"post_count"`
	LastPost    int64  `json:"last_post"`
	Language    string `json:"language"`
}
