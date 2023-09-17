package model





type Repository struct {
	Rid int  `gorm:"primary_key;column:rid" json:"rid"`
	Hash string `gorm:"column:hash" json:"hash"`
	Filename string `gorm:"column:filename" json:"filename"`
	Ext string `gorm:"column:ext" json:"ext"`
	Size int `gorm:"column:size" json:"size"`
	Path string `gorm:"column:path" json:"path"`
}

func (r Repository) TableName() string {
	return "repository"
}

func NewRepository(rid int, hash, filename, ext, path string, size int) *Repository {
	return &Repository{
		Rid: rid,
		Hash: hash,
		Filename: filename,
		Ext: ext,
		Path: path,
		Size: size,
	}
}