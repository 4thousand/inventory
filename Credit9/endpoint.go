package credit9

import (
	"context"
	"fmt"
)

func SentMessageOTP(s Service) interface{} {
	type request struct {
		Phone string `json:"phone"`
	}
	return func(ctx context.Context, req *request) (interface{}, error) {
		resp, err := s.SentMessageOTP(req.Phone)
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func ValidateMessageOTP(s Service) interface{} {
	type request struct {
		Secret string `json:"secret"`
		Phone  string `json:"phone"`
	}
	return func(ctx context.Context, req *request) (interface{}, error) {
		resp, err := s.ValidateMessageOTP(req.Secret, req.Phone)
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
