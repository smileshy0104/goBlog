package routes

import (
	"github.com/gin-gonic/gin"
	"goBlog/lib/global"
)

func InitSysRouter(r *gin.Engine) {
	sysGp := r.Group("/api/sys")
	{
		sysGp.GET("/test", func(c *gin.Context) {
			names := []string{}
			global.GVA_DB.Table("blog_user").Select("user_name").Scan(&names)
			c.JSON(200, gin.H{
				"msg": names,
			})
		})
		sysGp.GET("/test", func(c *gin.Context) {
			names := []string{}
			global.GVA_DB.Table("blog_user").Select("user_name").Scan(&names)
			c.JSON(200, gin.H{
				"msg": names,
			})
		})
	}
}
