package entity

type Student struct {
	StudentID           string `json:"studentID" gorm:"primary_key"`
	UserID              string `json:"userID"`
	Birthday            string `json:"birthday"`
	Name                string `json:"name"`
	Gender              string `json:"gender"`
	University          string `json:"university"`
	Major               string `json:"major"`
	EducationBackground string `json:"educationBackground"`
	GraduateTime        string `json:"graduateTime"`
}

func NewStudent(studentID string, userID string) *Student {
	return &Student{
		UserID:              userID,
		StudentID:           studentID,
		Birthday:            "保密",
		Name:                "保密",
		Gender:              "保密",
		University:          "保密",
		Major:               "保密",
		EducationBackground: "保密",
		GraduateTime:        "保密",
	}
}
