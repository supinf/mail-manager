package v1

import (
	"fmt"

	"github.com/go-openapi/swag"
	awsapigw "github.com/supinf/supinf-mail/app-api/aws/apigw"
	"github.com/supinf/supinf-mail/app-api/config"
	"github.com/supinf/supinf-mail/app-api/errors"
	apiModels "github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
	"github.com/supinf/supinf-mail/app-api/logs"
	"github.com/supinf/supinf-mail/app-api/misc"
	dbModels "github.com/supinf/supinf-mail/app-api/model"
)

func PostUsagePlan(req *apiModels.PostUsagePlanRequest) (*string, *errors.Error) {
	apigw, err := awsapigw.New()
	if err != nil {
		logs.Error("unable create api gateway client", err, nil)
		return nil, errors.InternalServerError
	}

	usagePlan, err := apigw.CreateUsagePlan(
		usagePlanName(req.UsagePlan.Name),
		apiThrottleToAwsThrottle(req.UsagePlan.Throttle),
		apiQuotaToAwsQuota(req.UsagePlan.Quota),
		apiStageToAwsStage(req.UsagePlan.APIStage),
	)
	if err != nil {
		logs.Error("unable create api usage plan", err, nil)
		return nil, errors.InternalServerError
	}

	return usagePlan.Id, nil
}

func PostUser(req *apiModels.PostUserRequest) (*string, *errors.Error) {
	// バリデーション（以下の２パターンのみ許容）
	//   [ パターン１: ドメインのみ指定 ]
	//     name: supinf, mail: @example.com
	//   [ パターン２: メールアドレスが異なる複数（ただし、ドメインのみの指定は NG） ]
	//     name: supinf, mail: supinf1@example.com
	//     name: supinf, mail: supinf2@example.com
	//     name: supinf, mail: supinf3@example.com
	list, err := dbModels.ListUserByGSI(swag.StringValue(req.User.Name))
	if err != nil {
		logs.Error("unable list users", err, nil)
		return nil, errors.InternalServerError
	}
	if len(list) > 0 {
		// 既存データが パターン１ のため追加 NG
		if misc.IsMailDomainOnly(list[0].Mail) {
			logs.Error("already subscribed user", nil, &logs.Map{
				"Name": swag.StringValue(req.User.Name),
			})
			return nil, errors.AlreadySubscribed
		}
		// 既存データが パターン２ のためドメイン指定は NG
		if misc.IsMailDomainOnly(swag.StringValue(req.User.Mail)) {
			logs.Error("cannot specify a domain only", nil, &logs.Map{
				"Name": swag.StringValue(req.User.Name),
			})
			return nil, errors.InvalidParameters
		}
		// 既存データが パターン２ のため メールアドレスの重複は NG
		for _, item := range list {
			if item.Mail == swag.StringValue(req.User.Mail) {
				logs.Error("already subscribed user", nil, &logs.Map{
					"Name": swag.StringValue(req.User.Name),
					"Mail": swag.StringValue(req.User.Mail),
				})
				return nil, errors.AlreadySubscribed
			}
		}
	}

	apigw, err := awsapigw.New()
	if err != nil {
		logs.Error("unable create api gateway client", err, nil)
		return nil, errors.InternalServerError
	}

	// API Key 発行
	apiKey, err := apigw.CreateAPIKey(apiKeyName(req.User.Name), req.User.Mail)
	if err != nil {
		logs.Error("unable create api key", err, nil)
		return nil, errors.InternalServerError
	}

	// Usage Plan と紐付け
	_, err = apigw.CreateUsagePlanKey(apiKey.Id, req.UsagePlanID)
	if err != nil {
		logs.Error("unable create api usage plan key", err, nil)
		return nil, errors.InternalServerError
	}

	// ユーザ登録
	user := &dbModels.User{
		APIKey:      apiKey.Value,
		APIKeyID:    swag.StringValue(apiKey.Id),
		Name:        swag.StringValue(req.User.Name),
		UsagePlanID: swag.StringValue(req.UsagePlanID),
		Mail:        swag.StringValue(req.User.Mail),
		Role:        dbModels.RoleNone,
	}
	if err := user.Create(); err != nil {
		logs.Error("unable create user", err, nil)
		return nil, errors.InternalServerError
	}

	return apiKey.Value, nil
}

func PatchUserEnabled(req *apiModels.PatchUserEnabledRequest) *errors.Error {
	user, err := dbModels.FindUserByHash(swag.StringValue(req.APIKey))
	if err != nil {
		logs.Error("not found user", err, nil)
		return errors.NotFound
	}

	apigw, err := awsapigw.New()
	if err != nil {
		logs.Error("unable create api gateway client", err, nil)
		return errors.InternalServerError
	}

	if swag.BoolValue(req.Enabled) {
		err = apigw.EnableAPIKey(swag.String(user.APIKeyID))
	} else {
		err = apigw.DisableAPIKey(swag.String(user.APIKeyID))
	}
	if err != nil {
		logs.Error("unable update api key enabled", err, nil)
		return errors.InternalServerError
	}

	return nil
}

func apiThrottleToAwsThrottle(throttle *apiModels.Throttle) *awsapigw.Throttle {
	if throttle == nil {
		return nil
	}
	return &awsapigw.Throttle{
		RateLimit:  throttle.RateLimit,
		BurstLimit: throttle.BurstLimit,
	}
}

func apiQuotaToAwsQuota(quota *apiModels.Quota) *awsapigw.Quota {
	if quota == nil {
		return nil
	}
	period := awsapigw.ConvertPeriod(swag.StringValue(quota.Period))
	return &awsapigw.Quota{
		Period: &period,
		Limit:  quota.Limit,
		Offset: swag.Int64(quota.Offset),
	}
}

func apiStageToAwsStage(stages []*apiModels.APIStage) []*awsapigw.APIStage {
	if stages == nil {
		return nil
	}
	apiStages := make([]*awsapigw.APIStage, len(stages))
	for i, stage := range stages {
		apiStages[i] = &awsapigw.APIStage{
			APIID: stage.APIID,
			Name:  stage.Name,
		}
	}
	return apiStages
}

func usagePlanName(planName *string) *string {
	name := fmt.Sprintf("%s-%s-%s-api-usage-plan", config.ApplicationsNameShort, config.AppStage, swag.StringValue(planName))
	return swag.String(name)
}

func apiKeyName(userName *string) *string {
	name := fmt.Sprintf("%s-%s-%s-api-key", config.ApplicationsNameShort, config.AppStage, swag.StringValue(userName))
	return swag.String(name)
}
