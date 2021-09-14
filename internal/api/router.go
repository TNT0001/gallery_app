package api

import (
	"log"
	gallerieshandler "tung.gallery/internal/handlers/galleries_handler"
	"tung.gallery/internal/middleware"
	"tung.gallery/internal/repo/gallery_repo"
	"tung.gallery/internal/repo/user_repo/friend"
	"tung.gallery/internal/repo/user_repo/user"
	"tung.gallery/internal/services/galleries"
	"tung.gallery/internal/services/users"

	"github.com/gin-gonic/gin"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/handlers"
	uh "tung.gallery/internal/handlers/users_handler"
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
	// Set router html template render
	//multiRender := utils.LoadDynamicTemplate("view/tmpl", "01")
	//r.Engine.HTMLRender = multiRender

	// Setup database
	db := models.NewDB()
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.Users{}, &entity.Galleries{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// User Handler
	userRepo := user.NewUserRepo(db)
	friendRepo := friend.NewFriendRepo(db)
	newUserService := users.NewUserService(userRepo, friendRepo)
	userHandler := uh.NewUserHandler(newUserService)

	// Gallery Handler
	galleryRepo := gallery_repo.NewGalleryRepo(db)
	newGalleryService := galleries.NewGalleryService(galleryRepo)
	galleryHandler := gallerieshandler.NewGalleryHandler(newGalleryService)

	// Home, Contact, Faq pages router
	//r.Engine.Use(middleware.AuthorizeJWT(userRepo))
	r.Engine.GET("/", handlers.Hello)
	r.Engine.GET("/contact", handlers.Contact)
	r.Engine.GET("/faq", handlers.Faq)

	r.Engine.Static("/assets/images", "./assets/images")
	// User router
	userAPI := r.Engine.Group("/user")
	{
		//userAPI.GET("/signup", userHandler.GetSignUpPage)
		userAPI.POST("/signup", userHandler.SignUp)
		//userAPI.GET("/login", userHandler.GetLoginPage)
		userAPI.POST("/login", userHandler.Login)
		//userAPI.GET("/logout", userHandler.LogOut)
		userAPI.GET("/:id", userHandler.GetUserInfo)

		userAPI.GET("/friendlist/:id", userHandler.GetUserFriendList)
	}

	userAPI.Use(middleware.AuthorizeJWT(userRepo))
	{
		userAPI.DELETE("/delete", userHandler.Delete)

		userAPI.PUT("/update", userHandler.Update)

		userAPI.POST("/add_friend", userHandler.AddFriend)
	}

	//Gallery router
	galleryApi := r.Engine.Group("/gallery")
	galleryApi.Use(middleware.AuthorizeJWT(userRepo))
	{
		galleryApi.GET("/show_all", galleryHandler.GetALlGalleryByUserID)
		galleryApi.GET("/show", galleryHandler.GetGalleryByID)
		galleryApi.POST("/new", galleryHandler.CreateGallery)
		galleryApi.POST("/:id/update", galleryHandler.UpdateGallery)
		galleryApi.DELETE("/:id/delete", galleryHandler.Delete)
	}
}
