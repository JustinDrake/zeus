package zk8s_templates

import (
	"context"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func BuildContainerDriver(ctx context.Context, name string, container Container) (v1.Container, error) {
	pp := "IfNotPresent"
	if len(container.ImagePullPolicy) > 0 {
		pp = container.ImagePullPolicy
	}
	c := v1.Container{
		Name:            name,
		Image:           container.DockerImage.ImageName,
		ImagePullPolicy: v1.PullPolicy(pp),
		Command:         strings.Split(container.DockerImage.Cmd, ","),
		Args:            strings.Split(container.DockerImage.Args, ","),
		Ports:           []v1.ContainerPort{},
		EnvFrom:         nil,
		Env:             nil,
		Resources: v1.ResourceRequirements{
			Limits:   make(map[v1.ResourceName]resource.Quantity),
			Requests: make(map[v1.ResourceName]resource.Quantity),
		},
		VolumeMounts: []v1.VolumeMount{},
	}

	for _, p := range container.DockerImage.Ports {
		if len(p.Name) <= 0 || len(p.Number) <= 0 {
			continue
		}
		// Use strconv.ParseInt to convert the string to int64
		numberInt64, err := strconv.ParseInt(p.Number, 10, 32)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("failed to parse port number")
			return c, err
		}
		c.Ports = append(c.Ports, v1.ContainerPort{
			Name:          p.Name,
			ContainerPort: int32(numberInt64),
			Protocol:      v1.Protocol(p.Protocol),
		})
	}

	for _, v := range container.DockerImage.VolumeMounts {
		if len(v.Name) > 0 && len(v.MountPath) > 0 {
			c.VolumeMounts = append(c.VolumeMounts, v1.VolumeMount{
				Name:      v.Name,
				MountPath: v.MountPath,
			})
		}
	}
	if len(container.DockerImage.ResourceRequirements.CPU) > 0 {
		c.Resources.Requests["cpu"] = resource.MustParse(container.DockerImage.ResourceRequirements.CPU)
		c.Resources.Limits["cpu"] = resource.MustParse(container.DockerImage.ResourceRequirements.CPU)
	}
	if len(container.DockerImage.ResourceRequirements.Memory) > 0 {
		c.Resources.Requests["memory"] = resource.MustParse(container.DockerImage.ResourceRequirements.Memory)
		c.Resources.Limits["memory"] = resource.MustParse(container.DockerImage.ResourceRequirements.Memory)
	}
	return c, nil
}
