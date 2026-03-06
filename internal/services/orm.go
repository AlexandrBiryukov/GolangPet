package services

type Tasks struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
}
