package routes

import (
	"finance-api/controller"
	"finance-api/middlewares"
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

	r.POST("/register", controller.CreateUser)
	r.POST("/login", controller.LoginUser)
	// r.GET("/status")

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.POST("/category", controller.CreateCategory)
		authorized.GET("/user", controller.GetUser)
		authorized.GET("/category", controller.GetCatergory)
		authorized.GET("/categories", controller.GetAllCategories)
		authorized.GET("/transaction/:idUser", controller.GetTransactions)
		authorized.POST("/transaction", controller.CreateTransaction)
		authorized.POST("/moneyType", controller.CreateMoneyType)
		authorized.GET("/moneyType", controller.GetAllMoneyType)
	}

	return r
}
