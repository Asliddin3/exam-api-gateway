package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/Asliddin3/exam-api-gateway/config"
	cb "github.com/Asliddin3/exam-api-gateway/genproto/customer"
	pb "github.com/Asliddin3/exam-api-gateway/genproto/post"
	rb "github.com/Asliddin3/exam-api-gateway/genproto/review"
)

type IServiceManager interface {
	CustomerService() cb.CustomerServiceClient
	PostService() pb.PostServiceClient
	ReviewService() rb.ReviewServiceClient
}

type ServiceManager struct {
	customerService cb.CustomerServiceClient
	postService     pb.PostServiceClient
	reviewService   rb.ReviewServiceClient
}

func (s *ServiceManager) PostService() pb.PostServiceClient {
	return s.postService
}

func (s *ServiceManager) CustomerService() cb.CustomerServiceClient {
	return s.customerService
}

func (s *ServiceManager) ReviewService() rb.ReviewServiceClient {
	return s.reviewService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connCustomer, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CustomerServiceHost, conf.CustomerServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connReview, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ReviewServiceHost, conf.ReviewServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	serviceManager := &ServiceManager{
		customerService: cb.NewCustomerServiceClient(connCustomer),
		postService:     pb.NewPostServiceClient(connPost),
		reviewService:   rb.NewReviewServiceClient(connReview),
	}

	return serviceManager, nil
}
