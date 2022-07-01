package apigw

import (
	"context"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

// CreateAPIKey API Key を発行します
func (ag *APIGateway) CreateAPIKey(name, description *string) (*apigateway.ApiKey, error) {
	input := &apigateway.CreateApiKeyInput{
		Name:        name,
		Description: description,
		Enabled:     awssdk.Bool(true),
	}

	out, err := ag.client.CreateApiKeyWithContext(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateAPIKey API Key を更新します
func (ag *APIGateway) UpdateAPIKey(id *string, from *string, op *Operation, path *Path, value *string) error {
	var mOp, mPath *string
	if op != nil {
		mOp = awssdk.String(op.String())
	}
	if path != nil {
		mPath = awssdk.String(path.String())
	}

	input := &apigateway.UpdateApiKeyInput{
		ApiKey: id,
		PatchOperations: []*apigateway.PatchOperation{
			{
				From:  from,
				Op:    mOp,
				Path:  mPath,
				Value: value,
			},
		},
	}

	_, err := ag.client.UpdateApiKeyWithContext(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}

// EnableAPIKey API Key を有効にします
func (ag *APIGateway) EnableAPIKey(id *string) error {
	op := OperationReplace
	path := PathEnabled
	return ag.UpdateAPIKey(id, nil, &op, &path, awssdk.String(Enable.String()))
}

// DisableAPIKey API Key を無効にします
func (ag *APIGateway) DisableAPIKey(id *string) error {
	op := OperationReplace
	path := PathEnabled
	return ag.UpdateAPIKey(id, nil, &op, &path, awssdk.String(Disable.String()))
}

// DeleteAPIKey API Key を削除します
func (ag *APIGateway) DeleteAPIKey(id *string) error {
	input := &apigateway.DeleteApiKeyInput{
		ApiKey: id,
	}

	_, err := ag.client.DeleteApiKeyWithContext(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}

// CreateUsagePlan Usage Plan を作成します
func (ag *APIGateway) CreateUsagePlan(name *string, throttle *Throttle, quota *Quota, stages []*APIStage) (*apigateway.UsagePlan, error) {
	var throttleSettings *apigateway.ThrottleSettings
	if throttle != nil {
		throttleSettings = &apigateway.ThrottleSettings{
			RateLimit:  throttle.RateLimit,
			BurstLimit: throttle.BurstLimit,
		}
	}

	var quotaSettings *apigateway.QuotaSettings
	if quota != nil {
		var period *string
		if quota.Period != nil {
			period = awssdk.String(quota.Period.String())
		}
		quotaSettings = &apigateway.QuotaSettings{
			Period: period,
			Limit:  quota.Limit,
			Offset: quota.Offset,
		}
	}

	apiStages := make([]*apigateway.ApiStage, len(stages))
	for i, stage := range stages {
		apiStages[i] = &apigateway.ApiStage{
			ApiId: stage.APIID,
			Stage: stage.Name,
			//Throttle: nil, // TODO: 確認
		}
	}

	input := &apigateway.CreateUsagePlanInput{
		Name:      name,
		Throttle:  throttleSettings,
		Quota:     quotaSettings,
		ApiStages: apiStages,
	}

	out, err := ag.client.CreateUsagePlanWithContext(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteUsagePlan Usage Plan を削除します
func (ag *APIGateway) DeleteUsagePlan(id *string) error {
	input := &apigateway.DeleteUsagePlanInput{
		UsagePlanId: id,
	}

	_, err := ag.client.DeleteUsagePlanWithContext(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}

// CreateUsagePlanKey Usage Plan Key を作成します
func (ag *APIGateway) CreateUsagePlanKey(apiKeyID, usagePlanID *string) (*apigateway.UsagePlanKey, error) {
	input := &apigateway.CreateUsagePlanKeyInput{
		KeyId:       apiKeyID,
		KeyType:     awssdk.String("API_KEY"),
		UsagePlanId: usagePlanID,
	}

	out, err := ag.client.CreateUsagePlanKeyWithContext(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteUsagePlanKey Usage Plan Key を削除します
func (ag *APIGateway) DeleteUsagePlanKey(apiKeyID, usagePlanID *string) error {
	input := &apigateway.DeleteUsagePlanKeyInput{
		KeyId:       apiKeyID,
		UsagePlanId: usagePlanID,
	}

	_, err := ag.client.DeleteUsagePlanKeyWithContext(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}
