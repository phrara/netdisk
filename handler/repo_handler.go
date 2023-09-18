package handler

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"netdisk/model"
	"netdisk/service"
	"path"

	"github.com/gin-gonic/gin"
)

var rs = service.NewRepoService()


// 上传文件到公共资源仓库
func UploadSourceHandler(c *gin.Context) {

	f, err := c.FormFile("file")
	if err != nil {
		return
	}

	filename := f.Filename
	ext := path.Ext(filename)
	size := int(f.Size)
	f2, _ := f.Open()
	b, _ := io.ReadAll(f2)
	// fmt.Println(b)
	f2.Close()
	hash := fmt.Sprintf("%x", md5.Sum(b))
	repo := model.NewRepository(0, hash, filename, ext, "", size)

	r := rs.UploadSource(repo, b)
	c.JSON(200, r)

}


// 资源详情
func DetailSourceHandler(c *gin.Context) {
	repo := &model.Repository{}
	if err := c.ShouldBind(repo); err != nil {
		return
	}
	r := rs.GetRepoDetails(repo)
	c.JSON(200, r)
}

// 资源下载
func DownloadSourceHandler(c *gin.Context) {
	repo := &model.Repository{}
	if err := c.ShouldBind(repo); err != nil {
		return
	}
	r := rs.DownloadSource(repo)

	if r["msg"] == "ok" {
		file := r["data"].([]byte)
		// 设置响应头
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", repo.Filename))
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Cache-Control", "no-cache")
		c.Header("Pragma", "no-cache")
		if _, err := io.Copy(c.Writer, bytes.NewReader(file)); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Copy file error: %s", err.Error()))
		}
	} else {
		c.JSON(http.StatusInternalServerError, r)
	}

}


// 上传到私人仓库
func SavePersonalFileHandler(c *gin.Context) {
	pr := &model.PersonalRepository{}
	if err := c.ShouldBind(pr); err != nil {
		return
	}

	r := rs.SavePersonalSource(pr)
	c.JSON(200, r)

}

// 私人仓库列表
func PersonalRepoListHandler(c *gin.Context) {
	info := &model.PersonalRepository{}
	if err := c.ShouldBind(info); err != nil {
		return
	}
	r := rs.GetRepoList(info)
	c.JSON(200, r)
}

// 上传到课程仓库
func SaveCourseFileHandler(c *gin.Context) {
	pr := &model.CourseRepository{}
	if err := c.ShouldBind(pr); err != nil {
		return
	}
	r := rs.SaveCourseSource(pr)
	c.JSON(200, r)
}

// 课程仓库列表
func CourseRepoListHandler(c *gin.Context) {
	info := &model.CourseRepository{}
	if err := c.ShouldBind(info); err != nil {
		return
	}
	r := rs.GetRepoList(info)
	c.JSON(200, r)
}

// 删除课程资源
func DeleteCourseSourceHandler(c *gin.Context) {
	rp := &model.CourseRepository{}
	if err := c.ShouldBind(rp); err != nil {
		return
	}
	r := rs.DeleteCourseSource(rp)
	c.JSON(200, r)
}

// 移动课程资源
func MoveCourseSourceHandler(c *gin.Context) {
	rp := &model.CourseRepository{}
	if err := c.ShouldBind(rp); err != nil {
		return
	}
	r := rs.MoveCourseSource(rp)
	c.JSON(200, r)
}