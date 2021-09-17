package reactsservice

import (
	"context"
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/reactdto"
)

func (s *reactService) GetReactByID(ctx context.Context, reactID int64) (*reactdto.GetReactByIDResponse, error) {
	react, err := s.Repo.GetReactByID(ctx, reactID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get react info")
	}

	reactMap, err := s.GetReactMap(ctx)
	if err != nil {
		return nil, err
	}

	return &reactdto.GetReactByIDResponse{
		Type:    reactMap[react.TypeID],
		UserID:  react.UserID,
		ImageID: react.ImageID,
	}, nil
}

func (s *reactService) GetReactByUserID(ctx context.Context, userID int64) (*reactdto.GetListReactByUserIDResponse, error) {
	reactIDs, err := s.Repo.ListReactByUserID(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get react info by user id")
	}

	reactMap, err := s.GetReactMap(ctx)
	if err != nil {
		return nil, err
	}

	reacts := make([]*reactdto.GetReactByIDResponse, 0)
	for _, r := range reactIDs {
		reacts = append(reacts, &reactdto.GetReactByIDResponse{
			Type:    reactMap[r.TypeID],
			ImageID: r.ImageID,
		})
	}

	return &reactdto.GetListReactByUserIDResponse{
		UserID: userID,
		Reacts: reacts,
	}, nil
}

func (s *reactService) GetReactByImageID(ctx context.Context, imageID int64) (*reactdto.GetListReactByImagesIDResponse, error) {
	reactIDs, err := s.Repo.ListReactByImageID(ctx, imageID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get react info by user id")
	}

	reactMap, err := s.GetReactMap(ctx)
	if err != nil {
		return nil, err
	}

	reacts := make([]*reactdto.GetReactByIDResponse, 0)
	for _, r := range reactIDs {
		reacts = append(reacts, &reactdto.GetReactByIDResponse{
			Type:   reactMap[r.TypeID],
			UserID: r.UserID,
		})
	}

	return &reactdto.GetListReactByImagesIDResponse{
		ImageID: imageID,
		Reacts:  reacts,
	}, nil
}

func (s *reactService) GetReactCountByImageID(ctx context.Context, imageID int64) (*reactdto.GetReactTotalByImageIDResponse, error) {
	total, err := s.Repo.GetReactCountTotalByImageID(ctx, imageID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get react info")
	}

	return &reactdto.GetReactTotalByImageIDResponse{
		Total:   total,
		ImageID: imageID,
	}, nil
}

func (s *reactService) GetReactCountByEachTypeAndImageID(ctx context.Context, imageID int64) (
	*reactdto.GetReactTotalByEachTypeAndImageIDResponse, error) {
	totalEachType, err := s.Repo.GetReactCountTotalByImageIDEachType(ctx, imageID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get react info")
	}

	reactMap, err := s.GetReactMap(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*reactdto.TotalReactCountByType, 0)

	for _, t := range totalEachType {
		res = append(res, &reactdto.TotalReactCountByType{
			Type:  reactMap[t.TypeID],
			Total: t.Total,
		})
	}

	return &reactdto.GetReactTotalByEachTypeAndImageIDResponse{TotalForEachType: res}, nil
}

func (s *reactService) GetReactMap(ctx context.Context) (map[int64]string, error) {
	reactType, err := s.Repo.GetTypeReact(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get react info")
	}

	reactMap := make(map[int64]string)
	for _, r := range reactType {
		reactMap[int64(r.ID)] = r.Type
	}

	return reactMap, nil
}
