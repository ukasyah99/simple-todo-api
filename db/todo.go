package db

type Todo struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `form:"title" json:"title"`
}
