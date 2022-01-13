package entity

type Advice struct {
	AdviceID     string `json:"adviceID" gorm:"primary_key"`
	UserID       string `json:"userID"`
	Device       string `json:"device"`
	AdviceDetail string `json:"adviceDetail"`
	Time         string `json:"time"`
	Identity     string `json:"identity"`
}
