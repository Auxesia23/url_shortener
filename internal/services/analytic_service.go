package service

import (
	"context"
	"net"

	"github.com/Auxesia23/url_shortener/internal/mapper"
	repository "github.com/Auxesia23/url_shortener/internal/repositories"
	"github.com/ipinfo/go/v2/ipinfo"
	"golang.org/x/sync/errgroup"
)

type AnalyticService interface{
	Save(ctx context.Context, url,ip,agent string)
	Get(ctx context.Context, url string)(mapper.AnalyticResponse, error)
}

type analyticService struct{
	analyticRepo repository.AnalyticRepository
	ipInfo *ipinfo.Client
}

func NewAnalyticService(analyticRepo repository.AnalyticRepository,ipInfo *ipinfo.Client)AnalyticService{
	return &analyticService{
		analyticRepo: analyticRepo,
		ipInfo: ipInfo,
	}
}

func(service *analyticService)Save(ctx context.Context, url,ip,agent string){

	var country string
	info, err := service.ipInfo.GetIPInfo(net.ParseIP(ip))
	if err != nil {
		country = "Unknown"
		return
	}
	
	if info.CountryName == ""{
		country = "Unknown"
	}else{
		country = info.CountryName
	}
	
	input := mapper.AnalyticInput{
		ShortenedUrl: url,
		IpAddress: ip,
		UserAgent: agent,
		Country: country,
	}
	
	analytic := mapper.ParseAnalyticInput(input)
	
	_ = service.analyticRepo.Create(ctx, analytic)
	return
}

func (service *analyticService) Get(ctx context.Context, url string) (mapper.AnalyticResponse, error) {
	var totalClicks int64
	var dailyClicks []mapper.DailyClickStat
	var clicksPerCountry []mapper.ClickStat
	var clicksPerUserAgent []mapper.ClickStat

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error 
		totalClicks, err = service.analyticRepo.GetTotalClicks(gCtx, url)

		return err
	})

	g.Go(func() error {
		var err error
		dailyClicks, err = service.analyticRepo.GetClicksPerDay(gCtx, url)
		return err
	})

	g.Go(func() error {
		var err error
		clicksPerCountry, err = service.analyticRepo.GetClicksPerCountry(gCtx, url)
		return err
	})

	g.Go(func() error {
		var err error
		clicksPerUserAgent, err = service.analyticRepo.GetClicksPerUserAgent(gCtx, url)
		return err
	})

	err := g.Wait()
	if err != nil {
		return mapper.AnalyticResponse{}, err
	}
	response := mapper.ParseAnalyticResponse(totalClicks, dailyClicks, clicksPerCountry, clicksPerUserAgent)

	return response, nil
}