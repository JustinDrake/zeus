package sui_actions

import (
	"fmt"

	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

func (t *SuiActionsCookbookTestSuite) TestGetCheckpoint() {
	cloudCtxNs := zeus_common_types.CloudCtxNs{
		CloudProvider: "aws",
		Region:        "us-west-1",
		Context:       "zeus-us-west-1",
		Namespace:     "sui-03e7d0b6",
	}
	resp, err := t.su.GetCheckpoint(ctx, cloudCtxNs, "1000")
	t.NoError(err)
	t.Equal(1, len(resp))
	for _, r := range resp {
		fmt.Println(r)
	}
}

func (t *SuiActionsCookbookTestSuite) TestGetLatestCheckpointSeqNum() {
	cloudCtxNs := zeus_common_types.CloudCtxNs{
		CloudProvider: "aws",
		Region:        "us-west-1",
		Context:       "zeus-us-west-1",
		Namespace:     "sui-03e7d0b6",
	}
	resp, err := t.su.GetLatestCheckpointSeqNumber(ctx, cloudCtxNs)
	t.NoError(err)
	t.Equal(1, len(resp))
	for _, r := range resp {
		fmt.Println(r)
	}
}
