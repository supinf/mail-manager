package v1

import (
	"github.com/go-openapi/swag"
	awsses "github.com/supinf/supinf-mail/app-api/aws/ses"
	"github.com/supinf/supinf-mail/app-api/aws/utils"
	"github.com/supinf/supinf-mail/app-api/errors"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
	"github.com/supinf/supinf-mail/app-api/logs"
	"github.com/supinf/supinf-mail/app-api/misc"
	"github.com/supinf/supinf-mail/app-api/model"
)

func PostMail(req *models.PostMailRequest, userName string) *errors.Error {
	ses, err := awsses.New()
	if err != nil {
		logs.Error("unable create ses client", err, nil)
		return errors.InternalServerError
	}

	// メール送信
	if err := sendMail(ses, req); err != nil {
		return err
	}

	// 履歴追加
	if err := createHistory(req, userName); err != nil {
		return err
	}

	return nil
}

func PostBulkMail(req *models.PostBulkMailRequest, userName string) *errors.Error {
	ses, err := awsses.New()
	if err != nil {
		logs.Error("unable create ses client", err, nil)
		return errors.InternalServerError
	}

	// メール送信
	if err := bulkSendMail(ses, req); err != nil {
		return err
	}

	// 履歴追加
	if err := bulkCreateHistory(req, userName); err != nil {
		return err
	}

	return nil
}

func sendMail(client *awsses.SES, req *models.PostMailRequest) *errors.Error {
	if err := client.Send(
		swag.String(req.From.Address.String()),
		[]*string{swag.String(req.Destination.To.Address.String())},
		mailAddressPtrSliceToStringPtrSlice(req.Destination.Cc),
		mailAddressPtrSliceToStringPtrSlice(req.Destination.Bcc),
		awsses.ConvertMimeType(swag.StringValue(req.ContentType.MimeType)),
		req.Content.Subject,
		swag.String(req.Content.Plain),
		swag.String(req.Content.HTML),
	); err != nil {
		logs.Error("unable send mail", err, nil)
		return errors.InternalServerError
	}

	return nil
}

func bulkSendMail(client *awsses.SES, req *models.PostBulkMailRequest) *errors.Error {
	if err := client.BulkSend(
		swag.String(req.From.Address.String()),
		toSESBulkEntrySlice(req.Entries),
		req.Content.Subject,
		swag.String(req.Content.Plain),
		swag.String(req.Content.HTML),
	); err != nil {
		logs.Error("unable bulk send mail", err, nil)
		return errors.InternalServerError
	}

	return nil
}

func createHistory(req *models.PostMailRequest, userName string) *errors.Error {
	mt := awsses.ConvertMimeType(swag.StringValue(req.ContentType.MimeType))

	history := &model.History{
		UserName: swag.String(userName),
		From:     req.From.Address.String(),
		To:       req.Destination.To.Address.String(),
		Cc:       mailAddressPtrSliceToStringSlice(req.Destination.Cc),
		Bcc:      mailAddressPtrSliceToStringSlice(req.Destination.Bcc),
		MimeType: &mt,
		Subject:  swag.StringValue(req.Content.Subject),
		Plain:    swag.String(req.Content.Plain),
		HTML:     swag.String(req.Content.HTML),
		SendAt:   misc.Now(),
	}

	if err := history.Create(); err != nil {
		logs.Error("unable create history", err, nil)
		return errors.InternalServerError
	}

	return nil
}

func bulkCreateHistory(req *models.PostBulkMailRequest, userName string) *errors.Error {
	for _, entry := range req.Entries {
		templateData, err := mapPtrToJSONStringPtr(entry.Map)
		if err != nil {
			return errors.InternalServerError
		}

		history := &model.History{
			UserName:     swag.String(userName),
			From:         req.From.Address.String(),
			To:           entry.Destination.To.Address.String(),
			Cc:           mailAddressPtrSliceToStringSlice(entry.Destination.Cc),
			Bcc:          mailAddressPtrSliceToStringSlice(entry.Destination.Bcc),
			Subject:      swag.StringValue(req.Content.Subject),
			Plain:        swag.String(req.Content.Plain),
			HTML:         swag.String(req.Content.HTML),
			TemplateData: templateData,
			SendAt:       misc.Now(),
		}

		if err := history.Create(); err != nil {
			logs.Error("unable create history", err, nil)
			return errors.InternalServerError
		}
	}

	return nil
}

func mailAddressPtrSliceToStringPtrSlice(mailAddress []*models.MailAddress) []*string {
	slice := make([]*string, len(mailAddress))
	for i, mail := range mailAddress {
		slice[i] = swag.String(mail.Address.String())
	}
	return slice
}

func mailAddressPtrSliceToStringSlice(mailAddress []*models.MailAddress) []string {
	slice := make([]string, len(mailAddress))
	for i, mail := range mailAddress {
		slice[i] = mail.Address.String()
	}
	return slice
}

func mapPtrToJSONStringPtr(mapObj *models.Map) (*string, error) {
	if mapObj == nil {
		return nil, nil
	}
	m := make(map[string]interface{})
	for _, item := range mapObj.Data {
		m[swag.StringValue(item.Key)] = item.Value
	}
	return utils.MapToJSONString(m)
}

func mapPtrToTemplateData(mapObj *models.Map) awsses.TemplateData {
	if mapObj == nil {
		return nil
	}
	var templateData awsses.TemplateData
	for _, item := range mapObj.Data {
		templateData = append(templateData, &awsses.TemplateDataItem{
			Key:   item.Key,
			Value: item.Value,
		})
	}
	return templateData
}

func toSESBulkEntrySlice(entries []*models.PostBulkMailRequestEntriesItems0) []*awsses.BulkEntry {
	bulkEntries := make([]*awsses.BulkEntry, len(entries))
	for i, entry := range entries {
		bulkEntries[i] = &awsses.BulkEntry{
			To:           mailAddressPtrSliceToStringPtrSlice([]*models.MailAddress{entry.Destination.To}),
			Cc:           mailAddressPtrSliceToStringPtrSlice(entry.Destination.Cc),
			Bcc:          mailAddressPtrSliceToStringPtrSlice(entry.Destination.Bcc),
			TemplateData: mapPtrToTemplateData(entry.Map),
		}
	}
	return bulkEntries
}
