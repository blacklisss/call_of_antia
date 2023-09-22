package routergin

import (
	"antia/internal/infrastructure/api/handlers"
	"antia/internal/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterGin struct {
	*gin.Engine
	hs     *handlers.Handlers
	config *util.Config
}

func NewRouterGin(config *util.Config, hs *handlers.Handlers) (*RouterGin, error) {
	r := gin.Default()
	ret := &RouterGin{
		hs: hs,
	}

	ret.config = config
	ret.setupRouter(r)

	ret.Engine = r
	return ret, nil
}

func (router *RouterGin) setupRouter(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(config))
	r.LoadHTMLGlob("templates/*")

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.GET("/:id", router.GetByUserID)
	r.POST("/addruneforteam", router.AddRuneForTeam)
	r.POST("/deleterelationbyid/:id", router.DeleteRelationByID)

	/*r.POST("/users", router.CreateUser)
	r.POST("/users/login", router.loginUser)
	r.GET("/email/verify/:token", router.VerifyUserEmail)
	r.POST("/tokens/renew_access", router.renewAccessToken)
	r.POST("/search", router.SearchOrganizations)
	r.GET("/organization/:id", router.GetOrganization)
	r.GET("/organizations/verify/:id", router.VerifyOrganization)*/

	/*authRoutes := r.Group("/").Use(authMiddleware(router.tokenMaker))
	authRoutes.POST("/organizations/create", router.CreateOrganization)
	authRoutes.GET("/organizations/:id", router.GetOrganizationByOwner)
	authRoutes.POST("/organizations/:id", router.UpdateOrganization)
	authRoutes.DELETE("/organizations/delete/:id", router.DeleteOrganization)
	authRoutes.GET("/organizations/my", router.ListMyOrganization)
	// authRoutes.GET("/organizations/verify/:id", router.VerifyOrganization)
	authRoutes.POST("/organizations/upload", router.UploadRegistration)
	authRoutes.POST("/organizations/contact", router.CreateOrganizationContact)
	authRoutes.GET("/organizations/contact/:id", router.GetOrganizationContact)
	authRoutes.GET("/organizations/contact/list/:id", router.ListOrganizationContact)
	authRoutes.DELETE("/organizations/contact/:id", router.DeleteOrganizationContact)*/

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
