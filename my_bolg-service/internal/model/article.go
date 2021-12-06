package model

type Article struct {
	Model *Model
	// 文章标题
	Title string `json:"title"`
	// 文章简述
	Desc string `json:"desc"`
	// 文章内容
	Content string `json:"content"`
	// 封面图片地址
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}
