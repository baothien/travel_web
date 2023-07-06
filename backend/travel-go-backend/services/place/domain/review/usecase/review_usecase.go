package usecase

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	placeRepo "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/repository"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/review/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/review/repository"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
	"gitlab.com/virtual-travel/travel-go-backend/utils/randutil"
	"go.uber.org/zap"
)

type ReviewUseCase struct {
	ReviewRepository repository.ReviewRepository
	PlaceRepository  placeRepo.PlaceRepository
}

func NewReviewUseCase(reviewRepository repository.ReviewRepository, placeRepository placeRepo.PlaceRepository) ReviewUseCase {
	return ReviewUseCase{
		ReviewRepository: reviewRepository,
		PlaceRepository:  placeRepository,
	}
}

func (r ReviewUseCase) CreateReview(ctx context.Context, review entity.Review) error {
	review.ID = randutil.UUIDRand()
	for i := range review.ReviewImg {
		review.ReviewImg[i].ID = randutil.UUIDRand()
		review.ReviewImg[i].ReviewID = review.ID
	}
	review.ChildReview = nil
	err := r.ReviewRepository.CreateReview(ctx, review)
	if err != nil {
		logger.Error("Error create review", zap.Error(err))
		return errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return nil
}

func (r ReviewUseCase) CreateChildReview(ctx context.Context, childReview entity.ChildReview) error {
	childReview.ID = randutil.UUIDRand()
	for i := range childReview.ReviewImg {
		childReview.ReviewImg[i].ID = randutil.UUIDRand()
		childReview.ReviewImg[i].ReviewID = childReview.ID
	}
	err := r.ReviewRepository.CreateChildReview(ctx, childReview)
	if err != nil {
		logger.Error("Error create review", zap.Error(err))
		return errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	go func() {

		notifyServiceUrl := "https://travel-api.huytx.com/stag/notify-service"
		reviewInfo, _ := r.ReviewRepository.GetReviewByID(ctx, childReview.ParentID)

		if *reviewInfo.UserID != *childReview.UserID {
			// get info user
			client := resty.New()
			respUser, errUser := client.R().
				SetAuthToken(ctx.Value("JWT_TOKEN").(string)).
				Get(fmt.Sprintf("https://travel-api.huytx.com/stag/user-service/user/detail/%s", *childReview.UserID))

			if errUser != nil {
				logger.Error("Error get info user", zap.Error(errUser))
			}

			esultGjson := gjson.Get(string(respUser.Body()), "error_info")
			if esultGjson.String() != "" {
				logger.Error("Error push notify", zap.String("err", esultGjson.String()))

			}

			resultGjson := gjson.Get(string(respUser.Body()), "data")
			if resultGjson.String() != "" {

			}

			resp, errNotify := client.R().
				SetAuthToken(ctx.Value("JWT_TOKEN").(string)).
				SetBody(map[string]string{
					"child_id":       childReview.ID,
					"destination_id": *reviewInfo.PlaceID,
					"from_user_id":   *childReview.UserID,
					"title":          fmt.Sprintf("%s đã phản hồi đánh giá của bạn: %s", resultGjson.Get("full_name"), childReview.Description),
					"body":           "",
					"to_user_id":     *reviewInfo.UserID,
					"type":           "REPLY_COMMENT",
				}).
				Post(fmt.Sprintf("%s/notify/create", notifyServiceUrl))

			fmt.Println(*reviewInfo.UserID)
			fmt.Println(string(resp.Body()))

			if errNotify != nil {
				fmt.Println(errNotify)
			}

			if resp.StatusCode() != 200 {

			}
		}

		// notify review for child user
		if reviewInfo.ChildReview != nil {
			userIDs := []string{}

			for _, item := range *reviewInfo.ChildReview {
				if *reviewInfo.UserID != *item.UserID {
					userIDs = append(userIDs, *item.UserID)
				}
			}

			userNoti := randutil.UniqueArrStr(userIDs)
			for i := range userNoti {
				if *reviewInfo.UserID != userNoti[i] && *childReview.UserID != userNoti[i] {
					// get info user
					client := resty.New()
					respUser, errUser := client.R().
						SetAuthToken(ctx.Value("JWT_TOKEN").(string)).
						Get(fmt.Sprintf("https://travel-api.huytx.com/stag/user-service/user/detail/%s", *childReview.UserID))

					if errUser != nil {
						logger.Error("Error get info user", zap.Error(errUser))
					}

					esultGjson := gjson.Get(string(respUser.Body()), "error_info")
					if esultGjson.String() != "" {
						logger.Error("Error push notify", zap.String("err", esultGjson.String()))

					}

					resultGjson := gjson.Get(string(respUser.Body()), "data")
					if resultGjson.String() != "" {

					}

					resp, errNotify := client.R().
						SetAuthToken(ctx.Value("JWT_TOKEN").(string)).
						SetBody(map[string]string{
							"child_id":       childReview.ID,
							"destination_id": *reviewInfo.PlaceID,
							"from_user_id":   *childReview.UserID,
							"title":          fmt.Sprintf("%s đã phản hồi đánh giá của bạn: %s", resultGjson.Get("full_name"), childReview.Description),
							"body":           "",
							"to_user_id":     userNoti[i],
							"type":           "REPLY_COMMENT",
						}).
						Post(fmt.Sprintf("%s/notify/create", notifyServiceUrl))

					fmt.Println(*reviewInfo.UserID)
					fmt.Println(string(resp.Body()))

					if errNotify != nil {
						fmt.Println(errNotify)
					}

					if resp.StatusCode() != 200 {

					}
				}

			}

		}

	}()

	return nil
}

func (r ReviewUseCase) GetReviewByPlaceID(ctx context.Context, req request.ListReviewPlaceReq) (api_response.Pagination, error) {
	paging, err := r.ReviewRepository.GetReviewByPlaceID(ctx, req)
	if err != nil {
		return paging, errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	return paging, nil
}

func (r ReviewUseCase) GetChildReviewByReviewID(ctx context.Context, req request.ListChildReviewReq) (api_response.Pagination, error) {
	paging, err := r.ReviewRepository.GetChildReviewByReviewID(ctx, req)
	if err != nil {
		return paging, errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	return paging, nil
}
