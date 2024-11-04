package provider

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

type AccountConfig struct {
	TenantID        string `json:"azure_tenant_id"`
	ClientID        string `json:"azure_client_id"`
	ClientSecret    string `json:"azure_client_password"`
	CertificatePath string `json:"certificatePath"`
	CertificatePass string `json:"certificatePass"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}

func NewAuthorizerFromConfig(cfg AccountConfig) (autorest.Authorizer, error) {
	settings, err := GetSettingsFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	return settings.GetAuthorizer()
}

func GetSettingsFromConfig(cfg AccountConfig) (s auth.EnvironmentSettings, err error) {
	s = auth.EnvironmentSettings{
		Values: map[string]string{},
	}

	if cfg.TenantID != "" {
		s.Values[auth.TenantID] = cfg.TenantID
	}
	if cfg.ClientID != "" {
		s.Values[auth.ClientID] = cfg.ClientID
	}
	if cfg.ClientSecret != "" {
		s.Values[auth.ClientSecret] = cfg.ClientSecret
	}
	if cfg.CertificatePath != "" {
		s.Values[auth.CertificatePath] = cfg.CertificatePath
	}
	if cfg.CertificatePass != "" {
		s.Values[auth.CertificatePassword] = cfg.CertificatePass
	}
	if cfg.Username != "" {
		s.Values[auth.Username] = cfg.Username
	}
	if cfg.Password != "" {
		s.Values[auth.Password] = cfg.Password
	}

	if v := s.Values[auth.EnvironmentName]; v == "" {
		s.Environment = azure.PublicCloud
	} else {
		s.Environment, err = azure.EnvironmentFromName(v)
	}
	if s.Values[auth.Resource] == "" {
		s.Values[auth.Resource] = s.Environment.ResourceManagerEndpoint
	}
	return
}

func AccountConfigFromMap(m map[string]any) (AccountConfig, error) {
	mj, err := json.Marshal(m)
	if err != nil {
		return AccountConfig{}, err
	}

	var c AccountConfig
	err = json.Unmarshal(mj, &c)
	if err != nil {
		return AccountConfig{}, err
	}

	return c, nil
}
