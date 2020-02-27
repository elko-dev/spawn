package organization

// CreateOrganizationArgs struct for creating organization
type CreateOrganizationArgs struct {
	DisplayName *string `json:"display_name,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// Organization struct
type Organization struct {
	ID          *string `json:"id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Name        *string `json:"name,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	Origin      *string `json:"origin,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}
