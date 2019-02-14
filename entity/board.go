package entity

// Board defines the board entity and its attributes.
type Board struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"descData,omitempty"`
	IsClosed       bool   `json:"closed"`
	OrganizationID string `json:"idOrganization"`
	IsPinned       bool   `json:"pinned"`
	URL            string `json:"url"`
	ShortURL       string `json:"shortUrl"`
	// Prefs
	// LabelNames
	IsStarred bool `json:"starred"`
	// Limits
	// Memberships
}

// prefs
// object

// Short for "preferences", these are the settings for the board

// labelNames
// object

// Object containing color keys and the label names given for one label of each color on the board. To get a full list of labels on the board see /boards/{id}/labels/.

// limits
// object

// An object containing information on the limits that exist for the board. Read more about at Limits.

// memberships
// array
