package misc

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
)

// Now 現在時刻を UTC で返します
func Now() time.Time {
	return time.Now().UTC()
}

func StrfmtDateTimeToTime(t strfmt.DateTime) time.Time {
	return time.Time(t)
}

func StrfmtDateTimeToTimePtr(t *strfmt.DateTime) *time.Time {
	if t == nil {
		return nil
	}
	tv := conv.DateTimeValue(t)
	return swag.Time(StrfmtDateTimeToTime(tv))
}

func TimeToStrfmtDateTime(t time.Time) strfmt.DateTime {
	return strfmt.DateTime(t)
}

func TimeToStrfmtDateTimePtr(t *time.Time) *strfmt.DateTime {
	if t == nil {
		return nil
	}
	tv := swag.TimeValue(t)
	st := TimeToStrfmtDateTime(tv)
	return &st
}
