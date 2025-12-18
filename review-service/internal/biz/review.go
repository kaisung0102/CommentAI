package biz

import (
	"context"
	v1 "review-service/api/review/v1"
	"review-service/internal/data/model"
	"review-service/third_party/uniqueid"
)

// // Review is a Review model.
// type Review struct {

// }

// ReviewRepo is a Review repo.
type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReviewByOrderID(context.Context, int64) ([]*model.ReviewInfo, error)
	GetReviewByReviewID(context.Context, int64) (*model.ReviewInfo, error)
	SaveReply(context.Context, *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error)
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

// CreateReview 创建评价 c端
func (uc *ReviewUsecase) CreateReview(ctx context.Context, r *model.ReviewInfo) (*model.ReviewInfo, error) {
	//1. 参数业务逻辑校验,同一个订单只能评价一次
	review, err := uc.repo.GetReviewByOrderID(ctx, r.OrderID)
	if err != nil {
		return nil, v1.ErrorDbError("database error: %v", err)
	}
	if len(review) > 0 {
		return nil, v1.ErrorOrderReviewed("order %d already reviewed", r.OrderID)
	}
	// 2. 生成review_ID
	r.ReviewID = uniqueid.GetID()
	// 3. 查询订单和商品快照

	// 4. 保存评价
	return uc.repo.SaveReview(ctx, r)
}

// CreateReply 创建评价回复 b端
func (uc *ReviewUsecase) CreateReply(ctx context.Context, r *ReplyParam) (*model.ReviewReplyInfo, error) {
	// 1.只能回复一次
	review, err := uc.repo.GetReviewByReviewID(ctx, r.ReviewID)
	if err != nil {
		return nil, v1.ErrorDbError("database error: %v", err)
	}
	if review.HasReply == 1 {
		return nil, v1.ErrorReplyExist("review %d already replied", r.ReviewID)
	}
	// 2.防止水平越权
	if review.StoreID != r.StoreID {
		return nil, v1.ErrorReplyNotAllowed("store %d not allowed to reply review %d", r.StoreID, r.ReviewID)
	}

	// 3. 保存回复
	return uc.repo.SaveReply(ctx, &model.ReviewReplyInfo{
		ReplyID:   uniqueid.GetID(),
		ReviewID:  r.ReviewID,
		StoreID:   r.StoreID,
		Content:   r.Content,
		PicInfo:   r.PicInfo,
		VideoInfo: r.VideoInfo,
	})
}
