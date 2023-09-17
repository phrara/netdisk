package service

import (
	"netdisk/dao"
	"netdisk/model"
	"netdisk/tool"
	"strings"
)

type RepoService struct {
	repoDao *dao.RepoDao
}

func NewRepoService() *RepoService {
	return &RepoService{
		repoDao: new(dao.RepoDao),
	}
}


func (rs *RepoService) UploadFile(repo *model.Repository, content []byte) tool.Res {
	repo.Filename = strings.Trim(repo.Filename, " ")
	if tool.WordsInspect(repo.Filename) {
		// 判断文件是否已存储
		if rs.repoDao.RepoInfo(repo).Filename == "" {
	
			// 上传到COS
			cosPath := "repository/" + repo.Hash + repo.Ext
			if path, err := tool.COSUpload(content, cosPath); err != nil {
				return tool.GetBadResult("COS upload failed")
			} else {
				repo.Path = path
			}
	
			// 本地记录
			if rs.repoDao.AddRepo(repo) {
				return tool.GetGoodResult(repo)
			} else {
				return tool.GetBadResult("upload failed")
			}
		} else {
			return tool.GetBadResult("file exists")	
		}
	} else {
		return tool.GetBadResult("illegal words")
	}
	
}


func (rs *RepoService) SavePersonalFile(pr *model.PersonalRepository) tool.Res {
	pr.SrcName = strings.Trim(pr.SrcName, " ")
	if tool.WordsInspect(pr.SrcName) {
		if rs.repoDao.PersonalRepoInfo(pr).SrcName == "" {
			if rs.repoDao.AddPersonalRepo(pr) {
				return tool.GetGoodResult(pr)
			} else {
				return tool.GetBadResult("save failed")
			}
		} else {
			return tool.GetBadResult("src exists")
		}
	} else {
		return tool.GetBadResult("illegal words")
	}

}


func (rs *RepoService) SaveCourseFile(pr *model.CourseRepository) tool.Res {
	pr.SrcName = strings.Trim(pr.SrcName, " ")
	pr.Cid = 1
	if tool.WordsInspect(pr.SrcName) {
		if rs.repoDao.CourseRepoInfo(pr).SrcName == "" {
			if rs.repoDao.AddCourseRepo(pr) {
				return tool.GetGoodResult(pr)
			} else {
				return tool.GetBadResult("save failed")
			}
		} else {
			return tool.GetBadResult("src exists")
		}
	} else {
		return tool.GetBadResult("illegal words")
	}
}

func (rs *RepoService) GetRepoList(info model.Repo) tool.Res {

	switch rp := info.(type) {
	case *model.PersonalRepository:
		if rp.Uid == 0 {
			list := rs.repoDao.TotalRepoList()
			return tool.GetGoodResult(list)
		} else {
			list := rs.repoDao.PersonalRepoList(rp)
			return tool.GetGoodResult(list)
		}
	case *model.CourseRepository:
		rp.Cid = 1
		list := rs.repoDao.CourseRepoList(rp)
		return tool.GetGoodResult(list)
	default: 
		return tool.GetBadResult("unkown type")
	}

	
}
	