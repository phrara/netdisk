package handler

import (
	"net/http"
	"netdisk/tool"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func RouteInit(r *gin.Engine, indexPath string) {

	// token 验签
	r.Use(TokenVerify())

	// 加载静态文件
	r.Static("/dist", "../dist")
	r.Static("/assets", "../dist/assets")


	// 首页跳转路由
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = indexPath
		r.HandleContext(c)
	})

	gets(r)
	posts(r)
}

// GET 请求路由
func gets(r *gin.Engine) {
	


}

// POST 请求路由
func posts(r *gin.Engine) {
	


	// 用户管理
	r.POST("/register", AddUserHandler)
	r.POST("/login", UserLoginHandler)
	r.POST("/upwd", UpdatePasswordHandler)
	r.POST("/delUser", DelUserHandler)

	// 课程
	r.POST("/addCourse", AddCourseHandler)
	r.POST("/joinCourse", JoinCourseHandler)
	

	// 公共资源仓库
	r.POST("/upload", UploadSourceHandler)
	r.POST("/detail", DetailSourceHandler)
	r.POST("download", DownloadSourceHandler)
	// 私人
	r.POST("/personalSave", SavePersonalFileHandler)
	r.POST("/personalList", PersonalRepoListHandler)
	// 课程
	r.POST("/courseSave", SaveCourseFileHandler)
	r.POST("/courseList", CourseRepoListHandler)
	r.POST("/courseDel", DeleteCourseSourceHandler)
	r.POST("/courseMove", MoveCourseSourceHandler)
}

// token验签
func TokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {


		if ctx.FullPath() == "/register" || ctx.FullPath() == "/login" {
			ctx.Next()
			return
		}

		tokenString := ctx.GetHeader("Authorization")
		if _, err := tool.ParseToken(tokenString); err != nil {
			ctx.JSON(http.StatusUnauthorized, tool.GetBadResult("invalid or expired token"))
            ctx.Abort()
			return
		}
		ctx.Next()
	}
}


 
 