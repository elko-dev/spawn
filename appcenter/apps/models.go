package apps

// CreateAppArgs struct
type CreateAppArgs struct {
	Description *string `json:"description,omitempty"`
	ReleaseType *string `json:"release_type,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Name        *string `json:"name,omitempty"`
	OS          *string `json:"os,omitempty"`
	Platform    *string `json:"platform,omitempty"`
}

// App Struct
type App struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	DisplayName string `json:"display_name"`
	ReleaseType string `json:"release_type"`
	IconURL     string `json:"icon_url"`
	IconSource  string `json:"icon_source"`
	Name        string `json:"name"`
	Os          string `json:"os"`
	Owner       struct {
		ID          string `json:"id"`
		AvatarURL   string `json:"avatar_url"`
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"owner"`
	AppSecret         string `json:"app_secret"`
	AzureSubscription struct {
		SubscriptionID      string `json:"subscription_id"`
		TenantID            string `json:"tenant_id"`
		SubscriptionName    string `json:"subscription_name"`
		IsBilling           bool   `json:"is_billing"`
		IsBillable          bool   `json:"is_billable"`
		IsMicrosoftInternal bool   `json:"is_microsoft_internal"`
	} `json:"azure_subscription"`
	Platform          string   `json:"platform"`
	Origin            string   `json:"origin"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_at"`
	MemberPermissions []string `json:"member_permissions"`
}
