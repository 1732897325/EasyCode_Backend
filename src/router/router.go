package router

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/controller"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/utils/remotejwtauth"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Controller *controller.Controller
}

func NewRouter(controller *controller.Controller) *Router {
	return &Router{
		Controller: controller,
	}
}

func (r *Router) RegisterRouters(engine *gin.Engine) {
	// config
	engine.UseRawPath = true

	// init route
	routerGroup := engine.Group("/api/v1")

	// builderRouter := routerGroup.Group("/teams/:teamID/builder")
	appRouter := routerGroup.Group("/teams/:teamID/apps")
	// appsRouter := routerGroup.Group("/apps")
	// publicAppRouter := routerGroup.Group("/teams/byIdentifier/:teamIdentifier/publicApps")
	// resourceRouter := routerGroup.Group("/teams/:teamID/resources")
	// actionRouter := routerGroup.Group("/teams/:teamID/apps/:appID/actions")
	// publicActionRouter := routerGroup.Group("/teams/byIdentifier/:teamIdentifier/apps/:appID/publicActions")
	// internalActionRouter := routerGroup.Group("/teams/:teamID/apps/:appID/internalActions")
	// roomRouter := routerGroup.Group("/teams/:teamID/room")
	// statusRouter := routerGroup.Group("/status")
	// oauth2Router := routerGroup.Group("/oauth2")

	// // register auth
	// builderRouter.Use(remotejwtauth.RemoteJWTAuth())
	appRouter.Use(remotejwtauth.RemoteJWTAuth())
	// appsRouter.Use(remotejwtauth.RemoteJWTAuth())
	// roomRouter.Use(remotejwtauth.RemoteJWTAuth())
	// actionRouter.Use(remotejwtauth.RemoteJWTAuth())
	// internalActionRouter.Use(remotejwtauth.RemoteJWTAuth())
	// resourceRouter.Use(remotejwtauth.RemoteJWTAuth())

	// // builder routers
	// builderRouter.GET("/desc", r.Controller.GetTeamBuilderDesc)

	// // app routers
	appRouter.POST("", r.Controller.CreateApp)
	// appRouter.DELETE(":appID", r.Controller.DeleteApp)
	// appRouter.PATCH(":appID/config", r.Controller.ConfigApp)
	// appRouter.GET("", r.Controller.GetAllApps)
	// appRouter.GET(":appID/versions/:version", r.Controller.GetFullApp)
	// appRouter.POST(":appID/duplication", r.Controller.DuplicateApp)
	// appRouter.POST(":appID/deploy", r.Controller.ReleaseApp)
	// appRouter.POST(":appID/takeSnapshot", r.Controller.TakeSnapshot)
	// appRouter.GET(":appID/snapshotList/limit/:pageLimit/page/:page", r.Controller.GetSnapshotList)
	// appRouter.GET(":appID/snapshot/:snapshotID", r.Controller.GetSnapshot)
	// appRouter.POST(":appID/recoverSnapshot/:snapshotID", r.Controller.RecoverSnapshot)
	// appRouter.GET("/list", r.Controller.GetAllAppByPage)
	// appRouter.GET("/list/like", r.Controller.SearchAppByKeywordsByPageUsingURIParam)
	// appRouter.GET("/list/limit/:limit/page/:page/sortBy/:sortBy/like/keywords/:keywords", r.Controller.SearchAppByKeywordsByPage)

	// // apps router
	// appsRouter.POST(":appID/forkTo/teams/:toTeamID", r.Controller.ForkMarketplaceApp)

	// // room routers
	// roomRouter.GET("/websocketConnection/dashboard", r.Controller.GetDashboardRoomConnectionAddress)
	// roomRouter.GET("/websocketConnection/app/:appID", r.Controller.GetAppRoomConnectionAddress)
	// roomRouter.GET("/binaryWebsocketConnection/app/:appID", r.Controller.GetAppRoomBinaryConnectionAddress)

	// // action routers
	// actionRouter.POST("", r.Controller.CreateAction)
	// actionRouter.GET("/:actionID", r.Controller.GetAction)
	// actionRouter.PUT("/:actionID", r.Controller.UpdateAction)
	// actionRouter.DELETE("/:actionID", r.Controller.DeleteAction)
	// actionRouter.POST("/:actionID/run", r.Controller.RunAction)

	// // internal action routers
	// internalActionRouter.POST("/generateSQL", r.Controller.GenerateSQL)

	// // resource routers
	// resourceRouter.GET("", r.Controller.GetAllResources)
	// resourceRouter.POST("", r.Controller.CreateResource)
	// resourceRouter.GET("/:resourceID", r.Controller.GetResource)
	// resourceRouter.PUT("/:resourceID", r.Controller.UpdateResource)
	// resourceRouter.DELETE("/:resourceID", r.Controller.DeleteResource)
	// resourceRouter.POST("/testConnection", r.Controller.TestConnection)
	// resourceRouter.GET("/:resourceID/meta", r.Controller.GetMetaInfo)
	// resourceRouter.POST("/:resourceID/token", r.Controller.CreateGoogleOAuthToken)
	// resourceRouter.GET("/:resourceID/oauth2", r.Controller.GetGoogleSheetsOAuth2Token)
	// resourceRouter.POST("/:resourceID/refresh", r.Controller.RefreshGoogleSheetsOAuth)

	// // public app routers
	// publicAppRouter.GET(":appID/versions/:version", r.Controller.GetFullPublicApp)
	// publicAppRouter.GET(":appID/isPublic", r.Controller.IsPublicApp)

	// // public action router
	// publicActionRouter.POST("/:actionID/run", r.Controller.RunPublicAction)

	// // oauth2 router
	// oauth2Router.GET("/authorize", r.Controller.GoogleOAuth2Exchange)

	// // status router
	// statusRouter.GET("", r.Controller.GetStatus)

	//database-manager  routers
	databaseManagerGroup := engine.Group("/database-manager")

	//database-manager GET /apps  router
	databaseManagerGroup.GET("/apps", r.Controller.GetAppsDataHandler)

	databaseManagerGroup.POST("/apps/:aid/tables", r.Controller.CreateTableHandler)

	databaseManagerGroup.PUT("/apps/:aid/tables/:tid", r.Controller.RenameTableHandler)

	databaseManagerGroup.DELETE("/apps/:aid/tables/:tid", r.Controller.DeleteTableHandler)

	databaseManagerGroup.GET("/apps/:aid/tables/:tid", r.Controller.GetTableData)

	databaseManagerGroup.POST("/apps", r.Controller.CreateAppData)

	//app routers
	appGroup := engine.Group("")

	appGroup.POST("/apps", r.Controller.CreateAppV2)

}
