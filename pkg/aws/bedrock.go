package aws_client

import (
	"context"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type BedrockClient struct {
	brc *bedrockruntime.Client
}

func NewBedrockClient(configYml *configs.Config) (*bedrockruntime.Client, error) {
	awsCfg := configYml.Config.AWS

	bedrockCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(awsCfg.Region))
	if err != nil {
		return nil, err
	}
	return bedrockruntime.NewFromConfig(bedrockCfg), nil
}

func (b *BedrockClient) Client() *bedrockruntime.Client {
	return b.brc
}
