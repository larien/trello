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
	Preferences    `json:"prefs"`
	LabelNames     `json:"labelNames"`
	IsStarred      bool `json:"starred"`
	// TODO - Limits  -> an object containing information on the limits that exist for the board
	Memberships []Membership `json:"memberships"`
}

// BackgroundImage contains the board background's width, height and URL.
type BackgroundImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// Preferences defines the board preferences.
type Preferences struct {
	Background            string            `json:"background"`
	BackgroundBrightness  string            `json:"backgroundBrightness"`
	BackgroundColor       string            `json:"backgroundColor"`
	BackgroundImage       string            `json:"backgroundImage"`
	BackgroundImageScaled []BackgroundImage `json:"backgroundImageScaled"`
	BackgroundTile        bool              `json:"backgroundTile"`
	CalendarFeedEnabled   bool              `json:"calendarFeedEnabled"`
	CanBeOrg              bool              `json:"canBeOrg"`
	CanBePublic           bool              `json:"canBePublic"`
	CanBePrivate          bool              `json:"canBePrivate"`
	CanInvite             bool              `json:"canInvite"`
	CardAging             string            `json:"cardAging"`
	CardCovers            bool              `json:"cardCovers"`
	Comments              string            `json:"comments"`
	Invitations           string            `json:"invitations"`
	PermissionLevel       string            `json:"permissionLevel"`
	SelfJoin              bool              `json:"selfjoin"`
	Voting                string            `json:"voting"`
}

// LabelNames contains all card labels for each color available.
type LabelNames struct {
	Black  string `json:"black,omitempty"`
	Blue   string `json:"blue,omitempty"`
	Green  string `json:"green,omitempty"`
	Lime   string `json:"lime,omitempty"`
	Orange string `json:"orange,omitempty"`
	Pink   string `json:"pink,omitempty"`
	Purple string `json:"purple,omitempty"`
	Red    string `json:"red,omitempty"`
	Sky    string `json:"sky,omitempty"`
	Yellow string `json:"yellow,omitempty"`
}

// Membership contains the member's ID
type Membership struct {
	ID string `json:"id"`
}
