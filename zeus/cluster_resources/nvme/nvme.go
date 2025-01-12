package zeus_nvme

import (
	aws_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/aws"
	do_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/do"
	gcp_nvme "github.com/zeus-fyi/zeus/zeus/cluster_resources/nvme/gcp"
)

// TODO: add local nvme vs block storage

func ConfigureCloudProviderStorageClass(cp string) string {
	switch cp {
	case "aws":
		return aws_nvme.AwsStorageClass
	case "gcp":
		return gcp_nvme.GcpStorageClass
	case "do":
		return do_nvme.DoStorageClass
	case "ovh":
		return ""
	// TODO when nvme is available in public cloud (OvhCloud US)
	default:
		return ""
	}
}

type DiskConfigDriverOpts struct {
	PvcName       string `json:"pvcName"`
	CloudProvider string `json:"cloudProvider"`
	Size          string `json:"diskSize"`
}
