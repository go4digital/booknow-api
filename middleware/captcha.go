package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Captcha-Token")
		if token != "" && http.MethodPost == request.Method {
			secretKey := "6LcgMgYcAAAAAAmCgeZAw7hurmIPpIM5k4xKSmYV"
			verificationUrl := fmt.Sprintf("https://www.google.com/recaptcha/api/siteverify?response=%v&secret=%s", token, secretKey)

			response, err := http.Post(verificationUrl, "application/json", bytes.NewBuffer([]byte("")))

			if err != nil {
				log.Error(err)
			}

			defer response.Body.Close()

			var res CaptchaResponse

			json.NewDecoder(response.Body).Decode(&res)

			if !res.Success && res.Score <= 0.2 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Bots not allowed!"))
				return
			}
		}
		next.ServeHTTP(w, request)
	})
}
