package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/go4digital/booknow-api/logger"
)

type CaptchaResponse struct {
	Success      bool    `json:"success"`
	Score        float64 `json:"score"`
	Challenge_ts string  `json:"challenge_ts"`
	Hostname     string  `json:"hostname"`
	Action       string  `json:"action"`
}

func VerifyCaptcha(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header.Get("Form-Submit") != "" {
			token := request.Header.Get("Captcha-Token")
			if token != "" && http.MethodPost == request.Method {
				captchaBaseUrl := os.Getenv("GOOGLE_CAPTCHA_VERIFICATION_URL")
				capptchaSecretKey := os.Getenv("GOOGLE_CAPTCHA_SECRET_KEY")
				urlWithParam := fmt.Sprintf("%v?secret=%v&response=%v", captchaBaseUrl, capptchaSecretKey, token)

				data, err := http.Post(urlWithParam, "application/json", bytes.NewBuffer([]byte("")))

				if err != nil {
					log.Error(err)
				}

				defer data.Body.Close()

				var res CaptchaResponse

				json.NewDecoder(data.Body).Decode(&res)

				if !res.Success && res.Score <= 0.2 {
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte("Bots not allowed!"))
					return
				}
			}
		}
		next.ServeHTTP(response, request)
	})
}
