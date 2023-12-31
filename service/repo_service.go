package service

import (
	"fmt"
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


func (rs *RepoService) UploadSource(repo *model.Repository, content []byte) tool.Res {
	repo.Filename = strings.Trim(repo.Filename, " ")
	if tool.WordsInspect(repo.Filename) {
		// 判断文件是否已存储, 秒传
		r := rs.repoDao.RepoInfo(repo)
		if r.Filename == "" {
	
			// 上传到COS
			cosPath := tool.Conf.COS.InnerPath + repo.Hash + repo.Ext
			// 小文件
			if repo.Size <= tool.Conf.COS.ChunkSize * 1024 * 1024 {
				if err := tool.COSUpload(content, cosPath); err != nil {
					return tool.GetBadResult(fmt.Errorf("%w: COS upload failed", err).Error())
				} 
			} else {
				// 分片
				if err := tool.COSMultipartUpload(content, cosPath) ; err != nil{
					return tool.GetBadResult("COS multipart upload failed")
				} 
			}
			repo.Path = cosPath
			
			// 本地记录
			if rs.repoDao.AddRepo(repo) {
				return tool.GetGoodResult(repo)
			} else {
				return tool.GetBadResult("upload failed")
			}
		} else {
			return tool.GetGoodResult(r)
		}
	} else {
		return tool.GetBadResult("illegal words")
	}
	
}


func (rs *RepoService) GetRepoDetails(repo *model.Repository) tool.Res {
	return tool.GetGoodResult(rs.repoDao.RepoDetail(repo))
}

func (rs *RepoService) DownloadSource(repo *model.Repository) tool.Res {
	r := rs.repoDao.RepoDetail(repo)
	// fmt.Println(repo, r)
	if r.Hash != "" && r.Path != "" {
		if file, err := tool.COSDownload(r.Path); err != nil {
			return tool.GetBadResult("COS download failed")
		} else {
			return tool.GetGoodResult(file)
		}
	} else {
		return tool.GetBadResult("source do not exist")
	}
}

func (rs *RepoService) SavePersonalSource(pr *model.PersonalRepository) tool.Res {
	pr.SrcName = strings.Trim(pr.SrcName, " ")
	if tool.WordsInspect(pr.SrcName) {
		if rs.repoDao.PersonalRepoInfo(pr).SrcName == "" {
			if rs.repoDao.AddPersonalRepo(pr) {
				return tool.GetGoodResult(pr)
			} else {
				return tool.GetBadResult("save failed")
			}
		} else {
			return tool.GetGoodResult("src exists")
		}
	} else {
		return tool.GetBadResult("illegal words")
	}

}

func (rs *RepoService) SaveCourseSource(pr *model.CourseRepository) tool.Res {
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
			return tool.GetGoodResult("src exists")
		}
	} else {
		return tool.GetBadResult("illegal words")
	}
}

// !下级目录
func (rs *RepoService) GetSubRepoList(info model.Repo) tool.Res {

	switch rp := info.(type) {
	case *model.PersonalRepository:
		if rp.Uid == 0 {
			list := rs.repoDao.TotalRepoList()
			return tool.GetGoodResult(list)
		} else {
			list := rs.repoDao.PersonalRepoList(rp.Uid, rp.PRid)
			return tool.GetGoodResult(list)
		}
	case *model.CourseRepository:
		rp.Cid = 1
		list := rs.repoDao.CourseRepoList(rp.Cid, rp.CRid)
		return tool.GetGoodResult(list)
	default: 
		return tool.GetBadResult("unkown type")
	}
}

// !上级目录
func (rs *RepoService) GetParentRepoList(info model.Repo) tool.Res {
	switch rp := info.(type) {
	case *model.PersonalRepository:
		pr := rs.repoDao.PersonalRepoDetail(rp.ParentId)
		list := rs.repoDao.PersonalRepoList(pr.Uid, pr.ParentId)
		return tool.GetGoodResult(list)
	case *model.CourseRepository:
		cr := rs.repoDao.CourseRepoDetail(rp.ParentId)
		list := rs.repoDao.CourseRepoList(cr.Cid, cr.ParentId)
		return tool.GetGoodResult(list)
	default:
		return tool.GetBadResult("unkown type")
	}
}

func (rs *RepoService) DeletePersonalSource(rp *model.PersonalRepository) tool.Res {
	if rs.repoDao.DeletePersonalRepo(rp) {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("delete failed")
	}
}

func (rs *RepoService) DeleteCourseSource(rp *model.CourseRepository) tool.Res {
	if rs.repoDao.DeleteCourseRepo(rp) {
		return tool.GetGoodResult(nil)
	} else {
		return tool.GetBadResult("delete failed")
	}
}

func (rs *RepoService) MovePersonalSource(rp *model.PersonalRepository) tool.Res {
	if rs.repoDao.UpdatePersonalRepoParentID(rp) {
		return tool.GetGoodResult(rp)
	} else {
		return tool.GetBadResult("move failed")
	}
}
	
func (rs *RepoService) MoveCourseSource(rp *model.CourseRepository) tool.Res {
	if rs.repoDao.UpdateCourseRepoParentID(rp) {
		return tool.GetGoodResult(rp)
	} else {
		return tool.GetBadResult("move failed")
	}
}