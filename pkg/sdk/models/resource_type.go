package models

import (
	"github.com/opengovern/og-describer-azure/provider/configs"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"github.com/opengovern/og-util/pkg/integration"
	"golang.org/x/net/context"
)

// any types are used to load your provider configuration.
type ResourceDescriber func(context.Context, configs.AccountCredentials, enums.DescribeTriggerType, map[string]string, *StreamSender) ([]Resource, error)
type SingleResourceDescriber func(context.Context, configs.AccountCredentials, enums.DescribeTriggerType, map[string]string) (*Resource, error)

type ResourceType struct {
	IntegrationType integration.Type

	ResourceName string

	Tags map[string][]string

	ListDescriber ResourceDescriber
	GetDescriber  SingleResourceDescriber
}

func (r ResourceType) GetIntegrationType() integration.Type {
	return r.IntegrationType
}

func (r ResourceType) GetResourceName() string {
	return r.ResourceName
}

func (r ResourceType) GetTags() map[string][]string {
	return r.Tags
}