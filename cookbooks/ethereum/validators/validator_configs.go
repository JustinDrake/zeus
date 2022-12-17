package validator_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

// TODO, will use matrix class later. Cluster class is a good near-term substitute
// extends on the beacon cluster test cases for now
var (
	ValidatorClusterClassName = "ethereumEphemeralValidatorCluster"
	execSkeletonBase          = "gethHercules"
	consensusSkeletonBase     = "lighthouseHercules"

	consensusValidatorClientComponentBaseName = "consensusValidatorClients"
	validatorSkeletonBaseName                 = "lighthouseHerculesValidatorClient"
)

var cd = zeus_req_types.ClusterTopologyDeployRequest{
	ClusterName: ValidatorClusterClassName,
	BaseOptions: []string{execSkeletonBase, consensusSkeletonBase, validatorSkeletonBaseName},
	CloudCtxNs:  ValidatorCloudCtxNs,
}

var DeployConsensusValidatorClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: ValidatorCloudCtxNs,
}

var ValidatorCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "ephemeral.staking", // set with your own namespace
	Env:           "production",
}

// chart workload metadata
var validatorsChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      validatorSkeletonBaseName,
	ChartName:         validatorSkeletonBaseName,
	ChartDescription:  validatorSkeletonBaseName,
	Version:           fmt.Sprintf("validatorBase-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  validatorSkeletonBaseName,
	ComponentBaseName: consensusValidatorClientComponentBaseName,
	ClusterBaseName:   ValidatorClusterClassName,
	Tag:               "latest",
}

var validatorsChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/validators/infra/validators",
	DirOut:      "./ethereum/validators/infra/processed_validators",
	FnIn:        validatorSkeletonBaseName, // filename for your gzip workload
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}
