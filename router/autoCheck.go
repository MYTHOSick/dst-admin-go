package router

import (
	"dst-admin-go/api"
	"github.com/gin-gonic/gin"
)

func initAutoCheck(router *gin.RouterGroup) {

	autoCheckApi := api.AutoCheckApi{}
	autoCheck := router.Group("/api/auto/check")
	{
		autoCheck.GET("/status", autoCheckApi.GetAutoCheckStatus)
		autoCheck.GET("/run", autoCheckApi.EnableAutoCheckRun)
		autoCheck.GET("/version", autoCheckApi.EnableAutoCheckUpdateVersion)
		autoCheck.GET("/mod", autoCheckApi.EnableAutoCheckGameMod)
	}

}
