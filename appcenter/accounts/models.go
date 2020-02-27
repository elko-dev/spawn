package accounts

// DistributionGroupArg to create group
type DistributionGroupArg struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

// DistributionGroupResponse response
type DistributionGroupResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Origin      string `json:"origin"`
	IsPublic    bool   `json:"is_public"`
}

// AppsForDistributionArg arguments
type AppsForDistributionArg struct {
	Apps []Apps `json:"apps"`
}

// Apps struct
type Apps struct {
	Name string `json:"name"`
}

// AddMemberArgs for distribution groups
type AddMemberArgs struct {
	UserEmails []string `json:"user_emails"`
}
