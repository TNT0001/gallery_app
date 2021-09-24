package api

import (
	"log"
	"tung.gallery/internal/handlers/comment"
	gallerieshandler "tung.gallery/internal/handlers/galleries_handler"
	"tung.gallery/internal/handlers/image"
	"tung.gallery/internal/handlers/react"
	"tung.gallery/internal/middleware"
	"tung.gallery/internal/repo"
	"tung.gallery/internal/services/commentservice"
	"tung.gallery/internal/services/galleryservice"
	"tung.gallery/internal/services/imageservice"
	"tung.gallery/internal/services/reactsservice"
	"tung.gallery/internal/services/userservice"

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

	ds := repo.NewRepo(db)

	// User Handler
	newUserService := userservice.NewUserService(ds)
	userHandler := uh.NewUserHandler(newUserService)

	// Gallery Handler
	newGalleryService := galleryservice.NewGalleryService(ds)
	galleryHandler := gallerieshandler.NewGalleryHandler(newGalleryService)

	// ImageHandler
	imageHanler := image.NewImageHandler(imageservice.NewImageService(ds))

	// ReactHandler
	reactHandler := react.NewReactHandler(reactsservice.NewReactService(ds))

	// CommentHandler
	commentHandler := comment.NewCommentHandler(commentservice.NewCommentService(ds))

	// middleware
	authorizeJWT := middleware.AuthorizeJWT(ds)

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

	userAPI.Use(authorizeJWT)
	{
		userAPI.DELETE("/delete", userHandler.Delete)

		userAPI.PUT("/update", userHandler.Update)

		userAPI.POST("/add_friend", userHandler.AddFriend)
	}

	//Gallery router
	galleryAPI := r.Engine.Group("/gallery")
	galleryAPI.Use(authorizeJWT)
	{
		galleryAPI.GET("/show_all", galleryHandler.GetALlGalleryByUserID)
		galleryAPI.GET("/show", galleryHandler.GetGalleryByID)
		galleryAPI.POST("/new", galleryHandler.CreateGallery)
		galleryAPI.POST("/:id/update", galleryHandler.UpdateGallery)
		galleryAPI.DELETE("/:id/delete", galleryHandler.Delete)
	}

	imageAPI := r.Engine.Group("/image")
	imageAPI.Use(authorizeJWT)
	{
		imageAPI.POST("/create", imageHanler.CreateImage)
		imageAPI.GET("/:id", imageHanler.GetImageByID)
		imageAPI.GET("/user/:id", imageHanler.GetImageByUserID)
		imageAPI.GET("/gallery/:id", imageHanler.GetImageByGalleryID)
	}

	commentAPI := r.Engine.Group("/comment")
	commentAPI.Use(authorizeJWT)
	{
		commentAPI.POST("/create", commentHandler.CreateComment)
		commentAPI.PUT("/edit/", commentHandler.EditComment)
		commentAPI.GET("/:id", commentHandler.GetComment)
		commentAPI.GET("/user/:id", commentHandler.GetCommentListByUserID)
		commentAPI.GET("image/:id", commentHandler.GetCommentListByImageID)
	}

	reactAPI := r.Engine.Group("/react")
	reactAPI.Use(authorizeJWT)
	reactRoute(reactAPI, reactHandler)
}

func reactRoute(r *gin.RouterGroup, h *react.ReactHandler) {
	{
		r.POST("/create", h.CreateReact)
		r.GET("/:id", h.GetReact)
		r.GET("/user/:id", h.GetReactByUserID)
		r.GET("/image/:id", h.GetReactByImageID)
		r.GET("/total_by_image/:id", h.GetReactCountByImageID)
		r.GET("/total_each_type_by_image/:id", h.GetReactCountEachTypeByImageID)
		r.DELETE("/delete/:id", h.DeleteReact)
		r.GET("/react_type", h.GetReactType)
	}
}
