package data

import (
	"context"

	"review-service/internal/biz"
	"review-service/internal/data/model"
	"review-service/internal/data/query"

	"github.com/go-kratos/kratos/v2/log"
)

// 实现了 biz.ReviewRepo 接口
type reviewRepo struct {
	data *Data
	log  *log.Helper
}

// NewReviewRepo .
func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	err := r.data.query.ReviewInfo.WithContext(ctx).Save(review)
	return review, err
}

func (r *reviewRepo) GetReviewByOrderID(ctx context.Context, orderID int64) ([]*model.ReviewInfo, error) {
	return r.data.query.ReviewInfo.
		WithContext(ctx).
		Where(r.data.query.ReviewInfo.OrderID.Eq(orderID)).
		Find()
}

func (r *reviewRepo) GetReviewByReviewID(ctx context.Context, reviewID int64) (*model.ReviewInfo, error) {
	return r.data.query.ReviewInfo.
		WithContext(ctx).
		Where(r.data.query.ReviewInfo.ReviewID.Eq(reviewID)).
		First()
}

func (r *reviewRepo) SaveReply(ctx context.Context, reply *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error) {
	// 引入事务，同时修改评价表的HasReply字段+保存回复
	err := r.data.query.Transaction(func(tx *query.Query) error {
		// 修改评价表的HasReply字段
		_, err := tx.ReviewInfo.WithContext(ctx).
			Where(tx.ReviewInfo.ReviewID.Eq(reply.ReviewID)).
			Update(tx.ReviewInfo.HasReply, 1)
		if err != nil {
			return err
		}
		// 保存回复
		if err := tx.ReviewReplyInfo.WithContext(ctx).Save(reply); err != nil {
			return err
		}
		return nil
	})
	return reply, err
}
