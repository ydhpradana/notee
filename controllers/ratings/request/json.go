package request

import "notee/business/ratings"

type Ratings struct {
	Rating      int `json:"rating"`
	Review       string `json:"review"`
	NoteId int    `json:"note_id"`
	UserId     int    `json:"user_id"`
}

func (req *Ratings) ToDomain() *ratings.Domain {
	return &ratings.Domain{
		Rating:      req.Rating,
		Review:       req.Review,
		NoteId: req.NoteId,
		UserId:     req.UserId,
	}
}
