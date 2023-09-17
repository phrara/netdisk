package model

type Repo interface {
	f()
}

type PersonalRepository struct {
	PRid int  `gorm:"primary_key;column:prid" json:"prid"`
	Uid int `gorm:"column:uid" json:"uid"`
	ParentId int `gorm:"column:parent_id" json:"parent_id"`
	Rid int  `gorm:"column:rid" json:"rid"`
	IsDir int `gorm:"column:isdir" json:"isdir"`
	SrcName string `gorm:"column:src_name" json:"src_name"`
}

func (r PersonalRepository) TableName() string {
	return "personal_repo"
}

func (r *PersonalRepository) f()  {
}

func NewPersonalRepository(prid, uid, parentId, rid, isdir int, srcName string) *PersonalRepository {
	return &PersonalRepository{
		PRid: prid,
		Uid: uid,
		ParentId: parentId,
		Rid: rid,
		IsDir: isdir,
		SrcName: srcName,
	}
}