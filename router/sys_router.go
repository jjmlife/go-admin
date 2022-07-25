package router

import "github.com/gin-gonic/gin"

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")

	pingRouter(g)
	
	return g
}

func pingRouter(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 0,
				"msg": "pong",
			})
		})
	}
}

func sysBaseRouter(r *gin.RouterGroup) {

}

func sysDBRouter(r *gin.RouterGroup) {

}

func sysStaticFileRouter(r *gin.RouterGroup) {

}

func sysSwaggerRouter() {

}

func sysNoCheckRoleRouter() {

}

func sysCheckRoleRouter() {

}

