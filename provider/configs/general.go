package configs

import integration "github.com/opengovern/og-util/pkg/integration"

const (
	Provider        = "azure"
	Cloud           = "AzureCloud"
	UpperProvider   = "Azure"
	IntegrationName = integration.Type("AZURE_SUBSCRIPTION")     // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL = "github.com/opengovern/og-describer-azure" // example: github.com/opengovern/og-describer-aws
)
