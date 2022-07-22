package router

import "github.com/gin-gonic/gin"

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")

	return g
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

