package main

import "context"

type metricsService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricsService{
		next: next,
	}
}

func (s *metricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {

	return s.next.FetchPrice(ctx, ticker)
}
