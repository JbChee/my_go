package model

type ArticleTag struct {
	Model     *Model
	TagID     uint64 `json:"tag_id"`
	ArticleID uint64 `json:"article_id"`
}
