package entity

type Admin struct {
	AdminID string `json:"adminID" gorm:"primary_key"`
	UserID  string `json:"userID"`
}

func NewAdmin(adminID string, userID string) *Admin {
	return &Admin{
		UserID:  userID,
		AdminID: adminID,
	}
}
