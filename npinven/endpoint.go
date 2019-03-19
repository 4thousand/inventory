package NPINven

import (
	"context"
	"fmt"
)

func GenDocNo(s Service) interface{} {
	type request struct {
		Type   string `json:"type"`
		Search string `json:"search"`
		Branch string `json:"branch"`
	}
	return func(ctx context.Context, req *request) (interface{}, error) {
		fmt.Println(1)
		resp, err := s.GenDocNoInven(req.Type, req.Search, req.Branch)
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
