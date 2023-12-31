package main

import (
	"log"
	"netdisk/handler"
	"netdisk/tool"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)




func main() {
	

	e := gin.Default()
	
	// 跨域
	e.Use(cors.Default())

	handler.RouteInit(e, "/dist")

	// fmt.Println(tool.Conf)

	err := e.Run(tool.Conf.Server.String())
	if err != nil {
		log.Fatal(err)
	}
}


// 解决跨域问题
// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 	   method := c.Request.Method
// 	   origin := c.Request.Header.Get("Origin")
// 	   if origin != "" {
// 		  c.Header("Access-Control-Allow-Origin", "*")  
// 		  c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
// 		  c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
// 		  c.Header("Access-Control-Allow-Credentials", "true")
// 		  c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") 
// 	   }
// 	   if method == "OPTIONS" {
// 		   /* 添加 */
// 		  c.Header("Access-Control-Allow-Origin", "*")  
// 		  c.Header("Access-Control-Allow-Methods", "OPTIONS")
// 		  c.Header("Access-Control-Allow-Headers", "*")
// 		   /* 添加 */
// 		  c.AbortWithStatus(http.StatusNoContent)
// 	   }
// 	   c.Next()
// 	}
//  }