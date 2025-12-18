package service

import (
	"context"
	"fmt"

	pb "review-service/api/review/v1"
	"review-service/internal/biz"
	"review-service/internal/data/model"
)

type ReviewService struct {
	pb.UnimplementedReviewServer
	uc *biz.ReviewUsecase
}

// NewReviewService new a ReviewService.
func NewReviewService(uc *biz.ReviewUsecase) *ReviewService {
	return &ReviewService{uc: uc}
}

// CreateReview implements pb.ReviewServer.
func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewReply, error) {
	// 请求转换为数据模型 model.ReviewInfo
	var anonymous int32
	if req.Anonymous {
		anonymous = 1
	} else {
		anonymous = 0
	}
	fmt.Println(req)
	review, err := s.uc.CreateReview(ctx, &model.ReviewInfo{
		UserID:       req.UserID,
		OrderID:      req.OrderID,
		StoreID:      req.StoreID,
		Score:        req.Score,
		ServiceScore: req.ServiceScore,
		ExpressScore: req.ExpressScore,
		Content:      req.Content,
		PicInfo:      req.PicInfo,
		VideoInfo:    req.VideoInfo,
		Anonymous:    anonymous,
		Status:       0,
	})
	fmt.Println(review)
	if err != nil {
		return nil, err
	}
	return &pb.CreateReviewReply{ReviewID: review.ReviewID}, nil
}

// ReplyReview 回复评价
func (s *ReviewService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	fmt.Printf("[service] ReplyReview req:%#v\n", req)
	// 调用biz层
	reply, err := s.uc.CreateReply(ctx, &biz.ReplyParam{
		ReviewID:  req.GetReviewID(),
		StoreID:   req.GetStoreID(),
		Content:   req.GetContent(),
		PicInfo:   req.GetPicInfo(),
		VideoInfo: req.GetVideoInfo(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReplyReviewReply{ReplyID: reply.ReplyID}, nil
}
