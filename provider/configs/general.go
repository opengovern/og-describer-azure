package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "azure"                                    // example: aws, azure
	IntegrationName      = integration.Type("AZURE_SUBSCRIPTION")     // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-azure" // example: github.com/opengovern/og-describer-aws
)
