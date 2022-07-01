package v1

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	awserr "github.com/supinf/supinf-mail/app-api/aws/errors"
	awsses "github.com/supinf/supinf-mail/app-api/aws/ses"
	"github.com/supinf/supinf-mail/app-api/errors"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
	"github.com/supinf/supinf-mail/app-api/logs"
	"github.com/supinf/supinf-mail/app-api/misc"
)

func ListSuppression(from, to *strfmt.DateTime, reasons []string, limit *int64, nextToken *string) (*models.ListSuppressionResponse, *errors.Error) {
	ses, err := awsses.New()
	if err != nil {
		logs.Error("unable create ses client", err, nil)
		return nil, errors.InternalServerError
	}

	suppressReasons := make([]awsses.SuppressedReason, len(reasons))
	for i, r := range reasons {
		suppressReasons[i] = awsses.ConvertSuppressedReason(r)
	}

	summary, token, err := ses.ListSuppressedDestination(
		misc.StrfmtDateTimeToTimePtr(from),
		misc.StrfmtDateTimeToTimePtr(to),
		suppressReasons,
		limit,
		nextToken,
	)
	if err != nil {
		logs.Error("unable list suppressed destinations", err, nil)
		return nil, errors.InternalServerError
	}

	list := make([]*models.Suppression, len(summary))
	for i, s := range summary {
		list[i] = toAPISuppression(s.EmailAddress, s.Reason, s.LastUpdateTime)
	}

	return &models.ListSuppressionResponse{
		List:      list,
		NextToken: swag.StringValue(token),
	}, nil
}

func GetSuppression(mail strfmt.Email) (*models.GetSuppressionResponse, *errors.Error) {
	ses, err := awsses.New()
	if err != nil {
		logs.Error("unable create ses client", err, nil)
		return nil, errors.InternalServerError
	}

	destination, err := ses.GetSuppressedDestination(swag.String(mail.String()))
	if err != nil {
		logs.Error("unable get suppressed destination", err, nil)
		return nil, errors.CodeToError(swag.IntValue(awserr.Code(err)))
	}

	return &models.GetSuppressionResponse{
		Item: toAPISuppression(
			destination.EmailAddress,
			destination.Reason,
			destination.LastUpdateTime,
		),
	}, nil
}

func PostSuppression(req *models.PostSuppressionRequest) *errors.Error {
	ses, err := awsses.New()
	if err != nil {
		logs.Error("unable create ses client", err, nil)
		return errors.InternalServerError
	}

	mail := swag.String(req.Content.Mail.Address.String())

	// すでに登録済みの場合はエラー
	destination, _ := ses.GetSuppressedDestination(mail)
	if destination != nil {
		logs.Error("already subscribed suppressed destination", nil, nil)
		return errors.AlreadySubscribed
	}

	suppressReason := awsses.ConvertSuppressedReason(swag.StringValue(req.Content.Reason.Type))

	if err := ses.PutSuppressedDestination(mail, suppressReason); err != nil {
		logs.Error("unable put suppressed destination", err, nil)
		return errors.CodeToError(swag.IntValue(awserr.Code(err)))
	}

	return nil
}

func DeleteSuppression(req *models.DeleteSuppressionRequest) *errors.Error {
	ses, err := awsses.New()
	if err != nil {
		logs.Error("unable create ses client", err, nil)
		return errors.InternalServerError
	}

	if err := ses.DeleteSuppressedDestination(swag.String(req.Mail.Address.String())); err != nil {
		// データが存在しない場合はここで 404 が返される
		logs.Error("unable delete suppressed destination", err, nil)
		return errors.CodeToError(swag.IntValue(awserr.Code(err)))
	}

	return nil
}

func toAPISuppression(mail, reason *string, t *time.Time) *models.Suppression {
	email := strfmt.Email(swag.StringValue(mail))
	return &models.Suppression{
		Content: &models.SuppressedDestination{
			Mail: &models.MailAddress{
				Address: &email,
			},
			Reason: &models.SuppressedReason{
				Type: reason,
			},
		},
		LastUpdateTime: misc.TimeToStrfmtDateTimePtr(t),
	}
}
