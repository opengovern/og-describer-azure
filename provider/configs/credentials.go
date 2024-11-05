package configs

type AccountCredentials struct {
	TenantID        string `json:"azure_tenant_id"`
	ClientID        string `json:"azure_client_id"`
	ClientSecret    string `json:"azure_client_password"`
	CertificatePath string `json:"certificatePath"`
	CertificatePass string `json:"certificatePass"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}
