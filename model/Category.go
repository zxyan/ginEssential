package model

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt Time   `json:"deleted_at" gorm:"type:timestamp"`
}

//type Category struct {
//	gorm.Model
//	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
//}
