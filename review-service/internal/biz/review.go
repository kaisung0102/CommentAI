package biz

import (
	"context"
	"review-service/internal/data/model"
	// v1 "review-service/api/review/v1"
	// "github.com/go-kratos/kratos/v2/errors"
	// "github.com/go-kratos/kratos/v2/log"
)

// // Review is a Review model.
// type Review struct {

// }

// ReviewRepo is a Review repo.
type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	// Update(context.Context, *Review) (*Review, error)
	// FindByID(context.Context, int64) (*Review, error)
	// ListByHello(context.Context, string) ([]*Review, error)
	// ListAll(context.Context) ([]*Review, error)
}

// ReviewUsecase is a Review usecase.
type ReviewUsecase struct {
	repo ReviewRepo
}

// NewReviewUsecase new a Review usecase.
func NewReviewUsecase(repo ReviewRepo) *ReviewUsecase {
	return &ReviewUsecase{repo: repo}
}

func (uc *ReviewUsecase) CreateReview(ctx context.Context, r *model.ReviewInfo) (*model.ReviewInfo, error) {
	// 业务逻辑
	return uc.repo.SaveReview(ctx, r)
}
