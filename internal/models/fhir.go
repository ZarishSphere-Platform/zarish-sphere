package models

// Practitioner represents a healthcare provider
type Practitioner struct {
	BaseModel

	Identifier string `gorm:"size:100;uniqueIndex;not null" json:"identifier"`
	Active     bool   `gorm:"default:true" json:"active"`

	FirstName string `gorm:"size:100" json:"first_name"`
	LastName  string `gorm:"size:100" json:"last_name"`
	Gender    string `gorm:"size:20" json:"gender"`

	Phone string `gorm:"size:50" json:"phone"`
	Email string `gorm:"size:255" json:"email"`

	Qualification string `gorm:"size:255" json:"qualification"`
	Specialty     string `gorm:"size:255" json:"specialty"`
}

// Organization represents a healthcare organization
type Organization struct {
	BaseModel

	Identifier string `gorm:"size:100;uniqueIndex;not null" json:"identifier"`
	Active     bool   `gorm:"default:true" json:"active"`
	Name       string `gorm:"size:255;not null" json:"name"`
	Type       string `gorm:"size:100" json:"type"` // hospital, clinic, pharmacy, etc.

	Phone   string `gorm:"size:50" json:"phone"`
	Email   string `gorm:"size:255" json:"email"`
	Address string `gorm:"type:text" json:"address"`
}

// Location represents a physical location
type Location struct {
	BaseModel

	Name   string `gorm:"size:255;not null" json:"name"`
	Status string `gorm:"size:50;default:'active'" json:"status"`
	Mode   string `gorm:"size:50" json:"mode"`  // instance, kind
	Type   string `gorm:"size:100" json:"type"` // ward, room, building

	Address string `gorm:"type:text" json:"address"`

	OrganizationID *uint `gorm:"index" json:"organization_id,omitempty"`
}

// TableName overrides
func (Practitioner) TableName() string {
	return "practitioners"
}

func (Organization) TableName() string {
	return "organizations"
}

func (Location) TableName() string {
	return "locations"
}
