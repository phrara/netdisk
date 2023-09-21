package dao

import "netdisk/model"

type RepoDao struct {
}



 
/*
					!公共资源仓库              
*/

func (rd *RepoDao) AddRepo(repo *model.Repository) bool {
	res := DBMgr.Create(repo)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func (rd *RepoDao) RepoInfo(repo *model.Repository) *model.Repository {
	r := &model.Repository{}
	DBMgr.Where("hash = ?", repo.Hash).First(r)
	return r
}

func (rd *RepoDao) RepoDetail(repo *model.Repository) *model.Repository {
	DBMgr.Where("rid = ?", repo.Rid).First(repo)
	return repo
}

func (rd *RepoDao) TotalRepoList() []*model.PersonalRepository {
	rpList := make([]*model.PersonalRepository, 10)
	DBMgr.Find(&rpList)
	return rpList
}


//             		!私有仓库

func (rd *RepoDao) PersonalRepoInfo(rp *model.PersonalRepository) *model.PersonalRepository {
	r := &model.PersonalRepository{}
	DBMgr.Where("uid = ? and isdir = ? and src_name = ? and parent_id = ?", rp.Uid, rp.IsDir, rp.SrcName, rp.ParentId).First(r)
	return r
}

func (rd *RepoDao) AddPersonalRepo(rp *model.PersonalRepository) bool {
	res := DBMgr.Create(rp)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func (rd *RepoDao) PersonalRepoDetail(prid int) *model.PersonalRepository {
	r := &model.PersonalRepository{}
	DBMgr.Where("prid = ?", prid).First(r)
	return r
}

func (rd *RepoDao) PersonalRepoList(uid, parentID int) []*model.PersonalRepository {
	rpList := make([]*model.PersonalRepository, 10)
	DBMgr.Where("uid = ? and parent_id = ?", uid, parentID).Find(&rpList)
	return rpList
}

func (rd *RepoDao) DeletePersonalRepo(rp *model.PersonalRepository) bool {
	res := DBMgr.Where("prid = ?", rp.PRid).Delete(rp)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}


func (rd *RepoDao) UpdatePersonalRepoParentID(rp *model.PersonalRepository) bool {
	res := DBMgr.Model(rp).Where("prid = ?", rp.PRid).Update("parent_id", rp.ParentId)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

//       			!课程共享仓库

func (rd *RepoDao) CourseRepoInfo(rp *model.CourseRepository) *model.CourseRepository {
	r := &model.CourseRepository{}
	DBMgr.Where("cid = ? and isdir = ? and src_name = ? and parent_id = ?", rp.Cid, rp.IsDir, rp.SrcName, rp.ParentId).First(r)
	return r
}

func (rd *RepoDao) AddCourseRepo(rp *model.CourseRepository) bool {
	res := DBMgr.Create(rp)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func (rd *RepoDao) CourseRepoDetail(crid int) *model.CourseRepository {
	r := &model.CourseRepository{}
	DBMgr.Where("crid = ?", crid).First(r)
	return r
}

func (rd *RepoDao) CourseRepoList(cid, parentID int) []*model.CourseRepository {
	rpList := make([]*model.CourseRepository, 10)
	DBMgr.Where("cid = ? and parent_id = ?", cid, parentID).Find(&rpList)
	return rpList
}

func (rd *RepoDao) DeleteCourseRepo(rp *model.CourseRepository) bool {
	res := DBMgr.Where("crid = ?", rp.CRid).Delete(rp)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
} 

func (rd *RepoDao) UpdateCourseRepoParentID(rp *model.CourseRepository) bool {
	res := DBMgr.Model(rp).Where("crid = ?", rp.CRid).Update("parent_id", rp.ParentId)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}