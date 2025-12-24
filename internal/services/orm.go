package services

type RequestBody struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
}
