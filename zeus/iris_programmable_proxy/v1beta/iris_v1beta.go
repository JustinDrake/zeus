package iris_programmable_proxy_v1_beta

import (
	"context"

	hestia_client "github.com/zeus-fyi/zeus/pkg/hestia/client"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	"github.com/zeus-fyi/zeus/zeus/iris_programmable_proxy"
)

type IrisV1Beta struct {
	iris_programmable_proxy.Iris
}

func NewIrisV1BetaClient(irisBase iris_programmable_proxy.Iris) IrisV1Beta {
	return IrisV1Beta{
		irisBase,
	}
}

func (i *IrisV1Beta) CreateProcedure(ctx context.Context, procedure IrisRoutingProcedure) error {
	hc := hestia_client.NewHestia(i.GetHestiaRoute(), i.Token)
	rr := hestia_req_types.IrisRoutingProcedureRequest{
		ProcedureName: procedure.Name,
	}
	err := hc.CreateIrisRoutingProcedure(ctx, rr)
	if err != nil {
		return err
	}
	return nil
}
