package request

import "notee/business/notes"

type Notes struct {
	Title string `json:"title"`
	Note string `json:"note"`
	CategoryId int `json:"category_id"`
	UserId int `json:"user_id"`
	IsFree bool `json:"is_free"`
}

func (req *Notes) ToDomain() *notes.Domain {
	return &notes.Domain{
		Title: req.Title,
		Note: req.Note,
		CategoryId: req.CategoryId,
		UserId: req.UserId,
		IsFree: req.IsFree,
	}
}
