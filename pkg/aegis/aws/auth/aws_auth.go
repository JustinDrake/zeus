package aws_aegis_auth

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/rs/zerolog/log"
)

type AuthAWS struct {
	AccountNumber string `json:"accountNumber"`
	Region        string `json:"region"`
	AccessKey     string `json:"accessKey"`
	SecretKey     string `json:"secretKey"`
}

func (a *AuthAWS) CreateConfig(ctx context.Context) (aws.Config, error) {
	creds := credentials.NewStaticCredentialsProvider(a.AccessKey, a.SecretKey, "")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(creds), config.WithRegion(a.Region))
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("AuthAWS: CreateConfig error loading config")
		return cfg, err
	}
	cfg.Region = a.Region
	return cfg, err
}

func (a *AuthAWS) GetCredentials(ctx context.Context) aws.Credentials {
	creds := aws.Credentials{
		AccessKeyID:     a.AccessKey,
		SecretAccessKey: a.SecretKey,
	}
	return creds
}