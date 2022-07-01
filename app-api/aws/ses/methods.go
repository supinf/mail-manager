package ses

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sesv2"
	"github.com/google/uuid"
	awserr "github.com/supinf/supinf-mail/app-api/aws/errors"
	"github.com/supinf/supinf-mail/app-api/aws/utils"
)

// Send メール送信を依頼します
func (s *SES) Send(from *string, to, cc, bcc []*string, mime MimeType, subject, plain, html *string) error {
	body := &sesv2.Body{}
	switch mime {
	case MimeTypeTextPlain:
		body.Text = &sesv2.Content{
			Charset: aws.String(DefaultCharset),
			Data:    plain,
		}
	case MimeTypeTextHTML:
		body.Html = &sesv2.Content{
			Charset: aws.String(DefaultCharset),
			Data:    html,
		}
	case MimeTypeAuto:
		if plain != nil {
			body.Text = &sesv2.Content{
				Charset: aws.String(DefaultCharset),
				Data:    plain,
			}
		}
		if html != nil {
			body.Html = &sesv2.Content{
				Charset: aws.String(DefaultCharset),
				Data:    html,
			}
		}
	}

	input := &sesv2.SendEmailInput{
		FromEmailAddress: from,
		Destination: &sesv2.Destination{
			ToAddresses:  to,
			CcAddresses:  cc,
			BccAddresses: bcc,
		},
		Content: &sesv2.EmailContent{
			Simple: &sesv2.Message{
				Subject: &sesv2.Content{
					Charset: aws.String(DefaultCharset),
					Data:    subject,
				},
				Body: body,
			},
		},
	}

	if _, err := s.client.SendEmailWithContext(context.Background(), input); err != nil {
		return err
	}
	return nil
}

// BulkSend 一括メール送信を依頼します
func (s *SES) BulkSend(from *string, entries []*BulkEntry, subject, plain, html *string) (err error) {
	if entries == nil {
		return errors.New("entries is required")
	}
	if !validateBulkSendEntryCount(entries) {
		return fmt.Errorf("maximum number of destinations (=%d) has been exceeded", BulkSendEntryMaximumCount)
	}

	// テンプレート作成
	templateName := aws.String(uuid.New().String())
	if err = s.CreateTemplate(
		templateName,
		subject,
		plain,
		html,
	); err != nil {
		return err
	}

	// テンプレート削除
	defer func() {
		err = s.DeleteTemplate(templateName)
	}()

	// メール一括送信
	bulkEmailEntries := make([]*sesv2.BulkEmailEntry, 0)
	for _, entry := range entries {
		if entry == nil {
			continue
		}
		item := &sesv2.BulkEmailEntry{
			Destination: &sesv2.Destination{
				ToAddresses:  entry.To,
				CcAddresses:  entry.Cc,
				BccAddresses: entry.Bcc,
			},
		}
		if data := entry.TemplateData.JSONString(); data != nil {
			item.ReplacementEmailContent = &sesv2.ReplacementEmailContent{
				ReplacementTemplate: &sesv2.ReplacementTemplate{
					ReplacementTemplateData: data,
				},
			}
		}
		bulkEmailEntries = append(bulkEmailEntries, item)
	}

	// デフォルトテンプレートデータとして空の値を持つ JSON を作成
	templateDefaultData := aws.String("{}")
	for _, entry := range bulkEmailEntries {
		if entry.ReplacementEmailContent != nil {
			jsonStr := entry.ReplacementEmailContent.ReplacementTemplate.ReplacementTemplateData
			emptyJSONData, err := utils.ReplaceAllFieldValues(aws.StringValue(jsonStr), "")
			if err != nil {
				return err
			}
			templateDefaultData = emptyJSONData
			break
		}
	}

	input := &sesv2.SendBulkEmailInput{
		FromEmailAddress: from,
		BulkEmailEntries: bulkEmailEntries,
		DefaultContent: &sesv2.BulkEmailContent{
			Template: &sesv2.Template{
				TemplateName: templateName,
				TemplateData: templateDefaultData,
			},
		},
	}

	if _, err := s.client.SendBulkEmailWithContext(context.Background(), input); err != nil {
		return err
	}
	return nil
}

// BulkSendWithTemplate テンプレートを指定して一括メール送信を依頼します
func (s *SES) BulkSendWithTemplate(from *string, entries []*BulkEntry, template *Template) error {
	if entries == nil {
		return errors.New("entries is required")
	}
	if template == nil {
		return errors.New("template is required")
	}
	if !validateBulkSendEntryCount(entries) {
		return fmt.Errorf("maximum number of destinations (=%d) has been exceeded", BulkSendEntryMaximumCount)
	}

	bulkEmailEntries := make([]*sesv2.BulkEmailEntry, 0)
	for _, entry := range entries {
		if entry == nil {
			continue
		}
		bulkEmailEntries = append(bulkEmailEntries, &sesv2.BulkEmailEntry{
			Destination: &sesv2.Destination{
				ToAddresses:  entry.To,
				CcAddresses:  entry.Cc,
				BccAddresses: entry.Bcc,
			},
			ReplacementEmailContent: &sesv2.ReplacementEmailContent{
				ReplacementTemplate: &sesv2.ReplacementTemplate{
					ReplacementTemplateData: entry.TemplateData.JSONString(),
				},
			},
		})
	}

	// デフォルトテンプレートデータが未指定の場合は空の値を持つ JSON を作成
	templateDefaultData := template.DefaultData.JSONString()
	if templateDefaultData == nil && len(bulkEmailEntries) > 0 {
		jsonStr := bulkEmailEntries[0].ReplacementEmailContent.ReplacementTemplate.ReplacementTemplateData
		emptyJSONData, err := utils.ReplaceAllFieldValues(aws.StringValue(jsonStr), "")
		if err != nil {
			return err
		}
		templateDefaultData = emptyJSONData
	}

	input := &sesv2.SendBulkEmailInput{
		FromEmailAddress: from,
		BulkEmailEntries: bulkEmailEntries,
		DefaultContent: &sesv2.BulkEmailContent{
			Template: &sesv2.Template{
				TemplateName: template.Name,
				TemplateData: templateDefaultData,
			},
		},
	}

	if _, err := s.client.SendBulkEmailWithContext(context.Background(), input); err != nil {
		return err
	}
	return nil
}

// ListTemplate テンプレートを一覧で取得します
func (s *SES) ListTemplate(limit *int64, nextToken *string) ([]*sesv2.EmailTemplateMetadata, *string, error) {
	input := &sesv2.ListEmailTemplatesInput{
		PageSize:  limit,
		NextToken: nextToken,
	}

	out, err := s.client.ListEmailTemplates(input)
	if err != nil {
		return nil, nil, err
	}
	return out.TemplatesMetadata, out.NextToken, nil
}

// GetTemplate テンプレートを取得します
func (s *SES) GetTemplate(name *string) (*sesv2.EmailTemplateContent, error) {
	input := &sesv2.GetEmailTemplateInput{
		TemplateName: name,
	}

	out, err := s.client.GetEmailTemplate(input)
	if err != nil {
		return nil, err
	}
	return out.TemplateContent, nil
}

// CreateTemplate テンプレートを作成します
func (s *SES) CreateTemplate(name *string, subject, plain, html *string) error {
	input := &sesv2.CreateEmailTemplateInput{
		TemplateName: name,
		TemplateContent: &sesv2.EmailTemplateContent{
			Subject: subject,
			Text:    plain,
			Html:    html,
		},
	}

	_, err := s.client.CreateEmailTemplateWithContext(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTemplate テンプレートを更新します
func (s *SES) UpdateTemplate(name *string, subject, plain, html *string) error {
	input := &sesv2.UpdateEmailTemplateInput{
		TemplateName: name,
		TemplateContent: &sesv2.EmailTemplateContent{
			Subject: subject,
			Text:    plain,
			Html:    html,
		},
	}

	_, err := s.client.UpdateEmailTemplateWithContext(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}

// SaveTemplate テンプレートを作成または更新します
func (s *SES) SaveTemplate(name *string, subject, plain, html *string) error {
	content, err := s.GetTemplate(name)
	if err != nil {
		// テンプレートが存在しない場合は作成
		if aws.IntValue(awserr.Code(err)) == 404 {
			return s.CreateTemplate(name, subject, plain, html)
		}
		// それ以外はエラー返却
		return err
	}

	// テンプレートの内容が異なる場合は更新
	if (aws.StringValue(subject) != "" && !utils.EqualsValue(content.Subject, subject)) ||
		(aws.StringValue(plain) != "" && !utils.EqualsValue(content.Text, plain)) ||
		(aws.StringValue(html) != "" && !utils.EqualsValue(content.Html, html)) {
		return s.UpdateTemplate(name, subject, plain, html)
	}

	// 上記以外は何もしない
	return nil
}

// DeleteTemplate テンプレートを削除します
func (s *SES) DeleteTemplate(name *string) error {
	input := &sesv2.DeleteEmailTemplateInput{
		TemplateName: name,
	}

	_, err := s.client.DeleteEmailTemplateWithContext(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}

// ListSuppressedDestination サプレッションリストから登録データを一覧で取得します
func (s *SES) ListSuppressedDestination(from, to *time.Time, reasons []SuppressedReason, limit *int64, nextToken *string) ([]*sesv2.SuppressedDestinationSummary, *string, error) {
	var suppressReasons []*string
	for _, r := range reasons {
		switch r {
		case SuppressedReasonBounce:
			suppressReasons = append(suppressReasons, aws.String(SuppressedReasonBounce.String()))
		case SuppressedReasonComplaint:
			suppressReasons = append(suppressReasons, aws.String(SuppressedReasonComplaint.String()))
		}
	}

	input := &sesv2.ListSuppressedDestinationsInput{
		StartDate: from,
		EndDate:   to,
		Reasons:   suppressReasons,
		PageSize:  limit,
		NextToken: nextToken,
	}

	out, err := s.client.ListSuppressedDestinationsWithContext(context.Background(), input)
	if err != nil {
		return nil, nil, err
	}
	return out.SuppressedDestinationSummaries, out.NextToken, nil
}

// GetSuppressedDestination サプレッションリストから登録データを取得します
func (s *SES) GetSuppressedDestination(mail *string) (*sesv2.SuppressedDestination, error) {
	input := &sesv2.GetSuppressedDestinationInput{
		EmailAddress: mail,
	}

	out, err := s.client.GetSuppressedDestinationWithContext(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return out.SuppressedDestination, nil
}

// PutSuppressedDestination サプレッションリストにメールアドレスを登録します
func (s *SES) PutSuppressedDestination(mail *string, reason SuppressedReason) error {
	input := &sesv2.PutSuppressedDestinationInput{
		EmailAddress: mail,
		Reason:       aws.String(reason.String()),
	}

	_, err := s.client.PutSuppressedDestinationWithContext(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSuppressedDestination サプレッションリストからメールアドレスを削除します
func (s *SES) DeleteSuppressedDestination(mail *string) error {
	input := &sesv2.DeleteSuppressedDestinationInput{
		EmailAddress: mail,
	}

	_, err := s.client.DeleteSuppressedDestinationWithContext(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}
