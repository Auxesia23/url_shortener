package handler

type ErrorMessage struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	AccesToken string `json:"acces_token"`
}

type GoogleUrlResponse struct {
	Url string `json:"url"`
}
