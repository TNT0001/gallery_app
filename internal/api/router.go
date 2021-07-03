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
	"tung.gallery/pkg/utils"
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
	// Set router html template render
	multiRender := utils.LoadDynamicTemplate("view/tmpl", "01")
	r.Engine.HTMLRender = multiRender

	// Setup database
	db := models.NewDB()
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.Users{}, &entity.Galleries{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// User Handler
	userRepo := repo.NewUserRepo(db)
	newUserService := services.NewUserService(userRepo)
	userHandler := uh.NewUserHandler(newUserService)

	// Gallery Handler
	galleryRepo := repo.NewGalleryRepo(db)
	newGalleryService := services.NewGalleryService(galleryRepo)
	galleryHandler := gh.NewGalleryHandler(newGalleryService)

	// Home, Contact, Faq pages router
	r.Engine.Use(middleware.AuthorizeJWT(userRepo))
	r.Engine.GET("/", handlers.Hello)
	r.Engine.GET("/contact", handlers.Contact)
	r.Engine.GET("/faq", handlers.Faq)

	// User router
	userAPi := r.Engine.Group("/user")
	{
		userAPi.GET("/signup", userHandler.GetSignUpPage)
		userAPi.POST("/signup", userHandler.SignUp)
		userAPi.GET("/login", userHandler.GetLoginPage)
		userAPi.POST("/login", userHandler.Login)
		userAPi.GET("/logout", userHandler.LogOut)
	}

	// Gallery router
	galleryApi := r.Engine.Group("/gallery")
	galleryApi.Use(middleware.LoginOnly())
	{
		galleryApi.GET("/", galleryHandler.ShowALlGalleries)
		galleryApi.GET("/new", galleryHandler.GetNewGalleryPage)
		galleryApi.POST("/new", galleryHandler.NewGallery)
		galleryApi.GET("/:id", galleryHandler.GetGalleryPage)
		galleryApi.GET("/:id/edit", galleryHandler.GetEditPage)
		galleryApi.POST("/:id/update", galleryHandler.EditGallery)
		galleryApi.POST("/:id/delete", galleryHandler.Delete)
	}
}