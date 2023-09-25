package handler

import (
	"net/http"
	"netdisk/tool"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func RouteInit(r *gin.Engine, indexPath string) {

	// token 验签
	// r.Use(TokenVerify())

	// 加载静态文件
	r.Static("/dist", "./dist")
	// r.Static("/css", "./dist/css")
	// r.Static("/js", "./dist/js")
	r.Static("/static/img", "./dist/static/img")
	// r.Static("/fonts", "./dist/fonts")


	// 首页跳转路由
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = indexPath
		r.HandleContext(c)
	})
	r.GET("/login", func(c *gin.Context) {
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
	r.POST("/register", UserRegisterHandler)
	r.POST("/login", UserLoginHandler)
	r.POST("/loginByEmail", UserLoginByEmailHandler)
	r.POST("/upwd", UpdatePasswordHandler)
	r.POST("/delUser", DelUserHandler)
	r.POST("/regEmail", RegisterSendCodeEmail)
	r.POST("/logEmail", LoginSendCodeEmail)
	r.POST("/upwdEmail", UpdatePasswordSendCodeEmail)

	// 课程
	r.POST("/addCourse", AddCourseHandler)
	r.POST("/joinCourse", JoinCourseHandler)
	

	// 公共资源仓库
	r.POST("/upload", UploadSourceHandler)
	r.POST("/detail", DetailSourceHandler)
	r.POST("/download", DownloadSourceHandler)
	// 私人
	r.POST("/personalSave", SavePersonalFileHandler)
	r.POST("/subPersonalList", SubPersonalRepoListHandler)
	r.POST("/parentPersonalList", ParentPersonalRepoListHandler)
	r.POST("/personalDel", DeletePersonalSourceHandler)
	r.POST("/personalMove", MovePersonalSourceHandler)
	// 课程
	r.POST("/courseSave", SaveCourseFileHandler)
	r.POST("/subCourseList", SubCourseRepoListHandler)
	r.POST("/parentCourseList", ParentCourseRepoListHandler)
	r.POST("/courseDel", DeleteCourseSourceHandler)
	r.POST("/courseMove", MoveCourseSourceHandler)
}

// token验签
func TokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {


		if ctx.FullPath() == "/register" || ctx.FullPath() == "/login" || ctx.FullPath() == "/regEmail" || ctx.FullPath() == "/logEmail" || ctx.FullPath() == "/upwdEmail" || ctx.FullPath() == "/loginByEmail" {
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


 
 