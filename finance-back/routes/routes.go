package routes

import (
	"finance-api/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://tusitio.com"}, // Aquí defines los dominios permitidos
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Habilita si tu API necesita compartir cookies o autenticación
	}))

	r.POST("/category", controller.CreateCategory)
	r.POST("/register", controller.CreateUser)
	r.GET("/user", controller.GetUser)
	r.POST("/login", controller.LoginUser)
	r.GET("/category", controller.GetCatergory)
	r.GET("/categories", controller.GetAllCategories)
	r.GET("/transaction/:idUser", controller.GetTransactions)
	r.POST("/transaction", controller.CreateTransaction)
	r.GET("/status")
	r.POST("/moneyType", controller.CreateMoneyType)
	r.GET("/moneyType", controller.GetAllMoneyType)
	// authorized:= r.Group("/")
	// authorized.Use(middlewares.AuthMiddleware())
	// {
	// }
	return r
}
