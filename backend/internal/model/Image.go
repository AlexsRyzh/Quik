package model

type Image struct {
	ID          uint   `gorm:"primaryKey"`
	FileName    string `gorm:"filename"`
	FileFormat  string `gorm:"column:file_format"`
	TypeConnect string `gorm:"column:type_connect"`

	IDConnect uint `gorm:"column:id_connect"`
}

func (Image) TableName() string {
	return "images"
}
