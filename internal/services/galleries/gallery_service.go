package galleries

import (
	"mime/multipart"
	"tung.gallery/internal/dt/dto/gallery_dto"
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
)

//var imageExt = []string{".jpg", ".jpeg", ".png"}

type GalleriesServiceInterface interface {
	CreateGallery(*entity.Users, *gallery_dto.GalleryCreateRequest) (*gallery_dto.GalleryCreateResponse, error)
	ShowGallery(id uint) (*gallery_dto.ShowGalleryResponse, error)
	Update(*entity.Users, *gallery_dto.GalleryUpdateRequest) (*gallery_dto.GalleryUpdateResponse, error)
	Delete(id uint) (*gallery_dto.GalleryDeleteResponse, error)
	GetAllGalleriesByUserID(id uint) (*user_dto.UserGetAllGalleriesResponse, error)
	UploadImage([]*multipart.FileHeader) bool
}

//type galleryService struct {
//	Repo repo.GalleryRepositoryInterface
//}
//
//func NewGalleryService(r repo.GalleryRepositoryInterface) GalleriesServiceInterface {
//	return &galleryService{
//		Repo: r,
//	}
//}
//
//func (g *galleryService) CreateGallery(user *entity.Users, req *dto.GalleryCreateRequest) (*dto.GalleryCreateResponse, error) {
//	gallery := entity.Galleries{
//		UserID: user.ID,
//		Title:  req.Title,
//	}
//
//	res, err := g.Repo.CreateGallery(&gallery)
//	if err != nil {
//		return nil, err
//	}
//
//	return &dto.GalleryCreateResponse{
//		Login:        true,
//		Title:        res.Title,
//		ID:           res.ID,
//		UserID:       res.UserID,
//	}, nil
//}
//
//func (g *galleryService) ShowGallery(id uint) (*dto.ShowGalleryResponse, error) {
//	gallery, err := g.Repo.ByID(id)
//	if err != nil {
//		return nil, err
//	}
//
//	basePath := filepath.Join("assets/images", strconv.Itoa(int(id)))
//	err = os.MkdirAll(basePath, 0777)
//	if err != nil {
//		return nil, err
//	}
//
//	img, err := filepath.Glob(basePath + "/*.*")
//	if err != nil {
//		return nil, err
//	}
//
//	images := make([]string, 0)
//	if len(img) > 0 {
//		for _, image := range img {
//			images = append(images, "/"+image)
//		}
//	}
//
//	return &dto.ShowGalleryResponse{
//		Title:        gallery.Title,
//		ID:           gallery.ID,
//		UserID:       gallery.UserID,
//		ImageUrls:    images,
//	}, nil
//}
//
//func (g *galleryService) Update(user *entity.Users, req *dto.GalleryUpdateRequest) (*dto.GalleryUpdateResponse, error) {
//	gallery, err := g.Repo.ByID(req.ID)
//	if err != nil {
//		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrShowGalleryFail.Error())
//		return dto.GalleryUpdateResponse{
//			Title:        gallery.Title,
//			ID:           gallery.ID,
//			BaseResponse: baseResponse}, err
//	}
//
//	gallery.Title = req.Title
//
//	err = g.Repo.Update(gallery)
//	if err != nil {
//		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrorUpdateGalleryFail.Error())
//		return dto.GalleryUpdateResponse{
//			Title:        gallery.Title,
//			ID:           gallery.ID,
//			BaseResponse: baseResponse}, err
//	}
//
//	baseResponse := utils.BaseResponse(false, AlertLvlSuccess, "update successfully")
//	return dto.GalleryUpdateResponse{
//		Title:        gallery.Title,
//		ID:           gallery.ID,
//		BaseResponse: baseResponse}, nil
//}
//
//func (g *galleryService) Delete(id uint) (*dto.GalleryDeleteResponse, error) {
//	err := g.Repo.Delete(id)
//	if err != nil {
//		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrorDeleteGalleryFail.Error())
//		return dto.GalleryDeleteResponse{BaseResponse: baseResponse}, err
//	}
//
//	baseResponse := utils.BaseResponse(false, AlertLvlSuccess, "success delete gallery")
//	return dto.GalleryDeleteResponse{BaseResponse: baseResponse}, nil
//}
//
//func (g *galleryService) GetAllGalleriesByUserID(id uint) (*dto.UserGetAllGalleriesResponse, error) {
//	galleries := make([]dto.Gallrery, 0)
//
//	allGalleries, err := g.Repo.ByUserID(id)
//	if err != nil {
//		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrShowGalleryFail.Error())
//		return dto.UserGetAllGalleriesResponse{BaseResponse: baseResponse}, err
//	}
//
//	for _, g := range allGalleries {
//		gallery := dto.Gallrery{
//			Title:  g.Title,
//			ID:     g.ID,
//			UserID: g.UserID,
//		}
//		galleries = append(galleries, gallery)
//	}
//
//	baseResponse := utils.BaseResponse(false, AlertLvlSuccess, "success get all gallery")
//	return dto.UserGetAllGalleriesResponse{
//		Galleries:    galleries,
//		BaseResponse: baseResponse,
//	}, nil
//}
//
//func (g *galleryService) UploadImage(files []*multipart.FileHeader) bool {
//	for _, file := range files {
//		fmt.Println(file.Filename)
//		validExt := utils.CheckingExt(filepath.Ext(file.Filename), imageExt)
//		if !validExt {
//			return false
//		}
//	}
//	return true
//}
