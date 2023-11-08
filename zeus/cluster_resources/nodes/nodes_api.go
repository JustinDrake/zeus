package nodes

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

const (
	searchNodesEndpoint = "/v1/search/nodes"
)

type NodeSearchParams struct {
	CloudProvider  string         `json:"cloudProvider,omitempty"`
	CloudRegion    string         `json:"region,omitempty"`
	DiskType       string         `json:"diskType,omitempty"`
	ResourceMinMax ResourceMinMax `json:"resourceMinMax,omitempty"`
}

type ResourceMinMax struct {
	Max ResourceAggregate `json:"max"`
	Min ResourceAggregate `json:"min"`
}

type ResourceAggregate struct {
	Price       float64 `json:"price"`
	MemRequests string  `json:"memRequests"`
	CpuRequests string  `json:"cpuRequests"`
}

func GetNodes(ctx context.Context, z zeus_client.ZeusClient, searchParams NodeSearchParams) (any, error) {
	z.PrintReqJson(searchParams)
	var respBody []Node
	respJson := zeus_resp_types.DeployStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(&respBody).
		Post(searchNodesEndpoint)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: GetNodes")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

func DeployNodes() {
	// TODO
}

func ScheduleNodes() {
	// TODO
}

func CreateNodeGroups() {
	// TODO
}
