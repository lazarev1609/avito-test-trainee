package models

type BannerTags struct {
	Id       uint `json:"id"`
	BannerId int  `json:"banner_id"`
	TagId    int  `json:"tag-id"`
}
