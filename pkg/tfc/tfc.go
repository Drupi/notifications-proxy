package tfc

// Struct definition is here: https://www.terraform.io/docs/cloud/api/notification-configurations.html#notification-payload
type TerraformWebhook struct {
	Version          int                             `json:"payload_version"`
	ConfigID         string                          `json:"notification_configuration_id"`
	RunURL           string                          `json:"run_url"`
	RunID            string                          `json:"run_id"`
	RunMessage       string                          `json:"run_message"`
	RunCreatedAt     string                          `json:"run_created_at"`
	RunCreatedBy     string                          `json:"run_created_by"`
	WorkspaceID      string                          `json:"workspace_id"`
	WorkspaceName    string                          `json:"workspace_name"`
	OrganizationName string                          `json:"organization_name"`
	Notifications    []TerraformWebhookNotifications `json:"notifications"`
}

type TerraformWebhookNotifications struct {
	Message      string `json:"message"`
	Trigger      string `json:"trigger"`
	RunStatus    string `json:"run_status"`
	RunUpdatedAt string `json:"run_updated_at"`
	RunUpdatedBy string `json:"run_updated_by"`
}
