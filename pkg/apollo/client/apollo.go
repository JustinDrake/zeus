package apollo_client

import (
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

type Apollo struct {
	zeus_client.ZeusClient
}

func NewApollo(baseURL, bearer string) Apollo {
	z := Apollo{}
	z.Resty = resty_base.GetBaseRestyClient(baseURL, bearer)
	return z
}

const ApolloEndpoint = "https://apollo.eth.zeus.fyi"

func NewDefaultApolloClient(bearer string) Apollo {
	return NewApollo(ApolloEndpoint, bearer)
}

const ApolloLocalEndpoint = "http://localhost:9000"

func NewLocalApolloClient(bearer string) Apollo {
	return NewApollo(ApolloLocalEndpoint, bearer)
}
