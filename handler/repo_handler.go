package handler

import (
	"crypto/md5"
	"fmt"
	"io"
	"netdisk/model"
	"netdisk/service"
	"path"

	"github.com/gin-gonic/gin"
)

var rs = service.NewRepoService()


// 上传文件到公共资源仓库
func UploadFileHandler(c *gin.Context) {

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

	r := rs.UploadFile(repo, b)
	c.JSON(200, r)

}

// 上传到私人仓库
func SavePersonalFileHandler(c *gin.Context) {
	pr := &model.PersonalRepository{}
	if err := c.ShouldBind(pr); err != nil {
		return
	}

	r := rs.SavePersonalFile(pr)
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
	r := rs.SaveCourseFile(pr)
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