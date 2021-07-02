package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/handlers"
	gh "tung.gallery/internal/handlers/galleries_handler"
	uh "tung.gallery/internal/handlers/users_handler"
	"tung.gallery/internal/middleware"
	"tung.gallery/internal/repo"
	"tung.gallery/internal/services"
	"tung.gallery/pkg/models"
)

type router struct {
	Engine *gin.Engine
	Log    *log.Logger
}

func NewRouter() *router {
	return &router{
		Engine: gin.Default(),
		Log:    &log.Logger{},
	}
}

// func DynamicLayoutRender(r *router, nameTemplate string) {
// 	template := template.ParseFiles()
// }

func Initialize(r *router) {
	r.Engine.LoadHTMLGlob("view/tmpl/*/*.html")
	db := models.NewDB()
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.Users{}, &entity.Galleries{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := repo.NewUserRepo(db)
	newUserService := services.NewUserService(userRepo)
	userHandler := uh.NewUserHandler(newUserService)

	galleryRepo := repo.NewGalleryRepo(db)
	newGalleryService := services.NewGalleryService(galleryRepo)
	galleryHandler := gh.NewGalleryHandler(newGalleryService)

	r.Engine.GET("/", handlers.Hello)
	r.Engine.GET("/contact", handlers.Contact)
	r.Engine.GET("/faq", handlers.Faq)

	userAPi := r.Engine.Group("/user")
	{
		userAPi.GET("/signup", userHandler.SignUp)
		userAPi.POST("/signup", userHandler.Create)
		userAPi.GET("/login", userHandler.LoginPage)
		userAPi.POST("/login", userHandler.Login)
		userAPi.Use(middleware.AuthorizeJWT(userRepo))
	}
	galleryApi := r.Engine.Group("/gallery")
	galleryApi.Use(middleware.AuthorizeJWT(userRepo))
	{
		galleryApi.GET("/new", galleryHandler.New)
		galleryApi.POST("/new", galleryHandler.Create)
		galleryApi.GET("/:id", galleryHandler.Show)
		galleryApi.GET("/:id/edit", galleryHandler.Edit)
	}

}
