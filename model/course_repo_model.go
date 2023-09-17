package model


type CourseRepository struct {
	CRid int  `gorm:"primary_key;column:crid" json:"crid"`
	Cid int `gorm:"column:cid" json:"cid"`
	ParentId int `gorm:"column:parent_id" json:"parent_id"`
	Rid int  `gorm:"column:rid" json:"rid"`
	IsDir int `gorm:"column:isdir" json:"isdir"`
	SrcName string `gorm:"column:src_name" json:"src_name"`
}

func (r CourseRepository) TableName() string {
	return "course_repo"
}

func (c *CourseRepository) f()  {
}

func NewCourseRepository(crid, cid, parentId, rid, isdir int, srcName string) *CourseRepository {
	return &CourseRepository{
		CRid: crid,
		Cid: cid,
		ParentId: parentId,
		Rid: rid,
		IsDir: isdir,
		SrcName: srcName,
	}
}