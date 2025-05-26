package mapper

import "github.com/Auxesia23/url_shortener/internal/models"

type AnalyticInput struct {
	ShortenedUrl string `json:"shortened_url"`
	IpAddress string `json:"ip_address"`
	Country string `json:"country"`
	UserAgent string `json:"user_agent"`
}

type ClickStat struct {
	Name string `json:"name"`
	Count int `json:"count"`
}

type DailyClickStat struct{
	Date string `json:"date"`
	Count int `json:"count"`
}

type AnalyticResponse struct {
	TotalClicks int64 `json:"total_click"`
	ClicksPerDay []DailyClickStat `json:"clicks_per_day"`
	ClicksPerCountry []ClickStat `json:"clicks_per_country"`
	ClicksPerUserAgent []ClickStat `json:"clicks_per_user_agent"`
	
}

func ParseAnalyticInput(input AnalyticInput)models.Analytic{
	return models.Analytic{
		ShortenedUrl: input.ShortenedUrl,
		IpAddress: input.IpAddress,
		Country: input.Country,
		UserAgent: input.UserAgent,
	}
}

func ParseAnalyticResponse(total int64, daily []DailyClickStat,perCountry,perUserAgent []ClickStat)AnalyticResponse{
	return AnalyticResponse{
		TotalClicks: total,
		ClicksPerDay: daily,
		ClicksPerCountry: perCountry,
		ClicksPerUserAgent: perUserAgent,
	}
}