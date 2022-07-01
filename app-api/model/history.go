package model

import (
	"time"

	"github.com/supinf/supinf-mail/app-api/aws/ses"
)

type History struct {
	UserName     *string       `json:"userName" dynamo:"user_name,hash"`
	From         string        `json:"from" dynamo:"from" localIndex:"from-index,range"`
	To           string        `json:"to" dynamo:"to" localIndex:"to-index,range"`
	Cc           []string      `json:"cc" dynamo:"cc"`
	Bcc          []string      `json:"bcc" dynamo:"bcc"`
	MimeType     *ses.MimeType `json:"mimeType" dynamo:"mime_type"`
	Subject      string        `json:"subject" dynamo:"subject"`
	Plain        *string       `json:"plain" dynamo:"plain"`
	HTML         *string       `json:"html" dynamo:"html"`
	TemplateData *string       `json:"templateData" dynamo:"template_data"`
	SendAt       time.Time     `json:"sendAt" dynamo:"send_at,range"`
}

func ListHistory(userName string, from, to *string, sendAtFrom, sendAtTo *time.Time) ([]History, error) {
	var list []History

	query := newQuery()
	query = query.Table(History{}).Get("user_name", userName)

	switch {
	case sendAtFrom != nil && sendAtTo != nil:
		query = query.Range("send_at", "between", sendAtFrom, sendAtTo)
	case sendAtFrom != nil:
		query = query.Range("send_at", ">=", sendAtFrom)
	case sendAtTo != nil:
		query = query.Range("send_at", "<=", sendAtTo)
	}

	if from != nil {
		query = query.IndexRange("from", "=", from)
	}
	if to != nil {
		query = query.IndexRange("to", "=", to)
	}

	if err := query.AllWithAutoFetch(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func (h *History) Create() error {
	query := newQuery()
	return query.Table(History{}).Put(h).Run()
}
