package api

import "github.com/andrepinto/goway-api/domain"

type ApiOptions struct {
	AnalyticRespository domain.IAnalyticRepository
}

type ApiResource struct {
	AnalyticRespository domain.IAnalyticRepository
}

func NewApiResource(options *ApiOptions) *ApiResource{
	repo := options.AnalyticRespository
	if(repo==nil){
		panic("Repository is required")
	}

	return &ApiResource{
		repo,
	}
}
