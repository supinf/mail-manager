package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/supinf/supinf-mail/app-api/aws/utils"
)

// BulkEntry 一括送信エントリ
type BulkEntry struct {
	To           []*string
	Cc           []*string
	Bcc          []*string
	TemplateData TemplateData
}

// Template テンプレート
type Template struct {
	Name        *string
	DefaultData TemplateData
}

// TemplateDataItem テンプレートマッピングアイテム
type TemplateDataItem struct {
	Key   *string
	Value interface{}
}

// TemplateData テンプレートマッピング
type TemplateData []*TemplateDataItem

// MimeType MimeType
type MimeType string

// SuppressedReason 追加された要因
type SuppressedReason string

const (
	// MimeTypeTextPlain プレーンテキスト
	MimeTypeTextPlain MimeType = "text/plain"
	// MimeTypeTextHTML HTML
	MimeTypeTextHTML MimeType = "text/html"
	// MimeTypeAuto SES 側で自動判別
	MimeTypeAuto MimeType = "auto"
)

const (
	// SuppressedReasonUndefined 未定義
	SuppressedReasonUndefined SuppressedReason = ""
	// SuppressedReasonBounce バウンス
	SuppressedReasonBounce SuppressedReason = "BOUNCE"
	// SuppressedReasonComplaint 苦情
	SuppressedReasonComplaint SuppressedReason = "COMPLAINT"
)

// JSONString TemplateData を JSON 文字列に変換します
func (td TemplateData) JSONString() *string {
	if td == nil {
		return nil
	}
	m := make(map[string]interface{})
	for _, item := range td {
		m[aws.StringValue(item.Key)] = item.Value
	}
	str, _ := utils.MapToJSONString(m)
	return str
}

// String MimeType を文字列に変換します
func (mt MimeType) String() string {
	return string(mt)
}

// ConvertMimeType 文字列を MimeType に変換します
func ConvertMimeType(value string) MimeType {
	switch value {
	case MimeTypeTextPlain.String():
		return MimeTypeTextPlain
	case MimeTypeTextHTML.String():
		return MimeTypeTextHTML
	case MimeTypeAuto.String():
		return MimeTypeAuto
	default:
		return MimeTypeAuto
	}
}

// String SuppressedReason を文字列に変換します
func (sr SuppressedReason) String() string {
	return string(sr)
}

// ConvertSuppressedReason 文字列を SuppressedReason に変換します
func ConvertSuppressedReason(value string) SuppressedReason {
	switch value {
	case SuppressedReasonBounce.String():
		return SuppressedReasonBounce
	case SuppressedReasonComplaint.String():
		return SuppressedReasonComplaint
	default:
		return SuppressedReasonUndefined
	}
}
