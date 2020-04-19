package req

import (
	"github.com/morteza-r/flexdb"
)

type QueryRequest struct {
	Table     string                 `json:"table"`
	Id        string                 `json:"id"`
	Doc       map[string]interface{} `json:"doc"`
	Where     []flexdb.Where         `json:"where"`
	WhereType string                 `json:"where_type"`
	Order     flexdb.Order           `json:"order"`
	Limit     int                    `json:"limit"`
	Type      string                 `json:"type"`
}

func (qr *QueryRequest) GetDoc() (doc *flexdb.Doc, err error) {
	doc = flexdb.NewDoc()
	for key, value := range qr.Doc {
		_ = doc.Set(key, value)
		//if err != nil {
		//	return
		//}
	}

	return
}

func (qr *QueryRequest) GetQuery() (query *flexdb.Query, err error) {
	doc, err := qr.GetDoc()
	if err != nil {
		return
	}
	doc.FixId()

	return &flexdb.Query{
		Table:     qr.Table,
		QId:       qr.Id,
		Doc:       doc,
		Where:     qr.Where,
		WhereType: qr.WhereType,
		Order:     qr.Order,
		Limit:     qr.Limit,
		Type:      qr.Type,
		Took:      0,
	}, nil
}
