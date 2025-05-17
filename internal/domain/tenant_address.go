package domain

// TenantAddress represents a structured address.
type TenantAddress struct {
	Street     string `json:"street" gorm:"column:street"`
	City       string `json:"city" gorm:"column:city"`
	State      string `json:"state" gorm:"column:state"`
	PostalCode string `json:"postal_code" gorm:"column:postal_code"`
	Country    string `json:"country" gorm:"column:country"`
}
