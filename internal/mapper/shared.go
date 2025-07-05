package mapper

type UrlAnalyticResponse struct {
	Url UrlResponse `json:"url"`
	Analytic AnalyticResponse `json:"analytic"`
}
