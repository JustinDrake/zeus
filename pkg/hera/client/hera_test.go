package hera_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	test_base "github.com/zeus-fyi/zeus/test"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type HeraClientTestSuite struct {
	test_suites.BaseTestSuite
	HeraTestClient HeraClient
}

var ctx = context.Background()

func (t *HeraClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	//t.HeraTestClient = NewLocalHeraClient(tc.Bearer)
	t.HeraTestClient = NewDefaultHeraClient(tc.Bearer)
	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()

	// generates outputs to /test/outputs dir
	// uses inputs from /test/mocks dir
}

func TestHeraClientTestSuite(t *testing.T) {
	suite.Run(t, new(HeraClientTestSuite))
}