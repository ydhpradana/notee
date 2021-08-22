package request

import "notee/business/transactions"

type Transactions struct {
	Code      string `json:"code"`
	NoteId int    `json:"note_id"`
	UserId     int    `json:"user_id"`
	TotalPrice int `json:"total_price"`
}

func (req *Transactions) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		Code:      req.Code,
		NoteId:       req.NoteId,
		UserId:     req.UserId,
		TotalPrice:     req.TotalPrice,
	}
}
