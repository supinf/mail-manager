package v1

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/supinf/supinf-mail/app-api/aws/utils"
	"github.com/supinf/supinf-mail/app-api/errors"
	apiModels "github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
	"github.com/supinf/supinf-mail/app-api/logs"
	"github.com/supinf/supinf-mail/app-api/misc"
	dbModels "github.com/supinf/supinf-mail/app-api/model"
)

func ListHistory(userName string, from, to *strfmt.Email, sendAtFrom, sendAtTo *strfmt.DateTime) (*apiModels.ListHistoryResponse, *errors.Error) {
	var mFrom, mTo *string
	if from != nil {
		mFrom = swag.String(from.String())
	}
	if to != nil {
		mTo = swag.String(to.String())
	}

	histories, err := dbModels.ListHistory(
		userName,
		mFrom,
		mTo,
		misc.StrfmtDateTimeToTimePtr(sendAtFrom),
		misc.StrfmtDateTimeToTimePtr(sendAtTo),
	)
	if err != nil {
		logs.Error("unable list history", err, nil)
		return nil, errors.InternalServerError
	}

	var response []*apiModels.History
	for _, history := range histories {
		if item, err := toAPIHistory(history); err == nil {
			response = append(response, item)
		} else {
			logs.Error("unable convert to response history", err, nil)
			return nil, errors.InternalServerError
		}
	}

	return &apiModels.ListHistoryResponse{
		List: response,
	}, nil
}

func toAPIHistory(history dbModels.History) (*apiModels.History, error) {
	templateMap, err := jsonStringToMap(history.TemplateData)
	if err != nil {
		return nil, err
	}

	var mimeType *string
	if history.MimeType != nil {
		mimeType = swag.String(history.MimeType.String())
	}

	return &apiModels.History{
		UserName: history.UserName,
		From:     toMailAddress(history.From),
		Destination: &apiModels.MailDestination{
			To:  toMailAddress(history.To),
			Cc:  toMailAddressSlice(history.Cc),
			Bcc: toMailAddressSlice(history.Bcc),
		},
		ContentType: &apiModels.MailContentType{
			MimeType: mimeType,
		},
		Content: &apiModels.MailContent{
			Subject: swag.String(history.Subject),
			Plain:   swag.StringValue(history.Plain),
			HTML:    swag.StringValue(history.HTML),
		},
		Map:    templateMap,
		SendAt: misc.TimeToStrfmtDateTimePtr(swag.Time(history.SendAt)),
	}, nil
}

func toMailAddress(mail string) *apiModels.MailAddress {
	email := strfmt.Email(mail)
	return &apiModels.MailAddress{
		Address: &email,
	}
}

func toMailAddressSlice(mails []string) []*apiModels.MailAddress {
	mailAddresses := make([]*apiModels.MailAddress, len(mails))
	for i, mail := range mails {
		mailAddresses[i] = toMailAddress(mail)
	}
	return mailAddresses
}

func jsonStringToMap(str *string) (*apiModels.Map, error) {
	m, err := utils.JSONStringToMap(str)
	if err != nil {
		return nil, err
	}

	mapItems := make([]*apiModels.MapItem, len(m))
	i := 0
	for k, v := range m {
		mapItems[i] = &apiModels.MapItem{
			Key:   swag.String(k),
			Value: v,
		}
		i++
	}

	return &apiModels.Map{
		Data: mapItems,
	}, nil
}
