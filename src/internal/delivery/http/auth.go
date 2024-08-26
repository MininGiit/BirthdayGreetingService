package http

import (
	"encoding/json"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"context"
	"log"
	//"fmt"
)

var jwtKey = []byte("secret_key")

type Credentailas struct {
	Login string `json:"username"`
	Password string `json:"password"`
}

type MyClaims struct {
    Login string `json:"username"`
    jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var user Credentailas
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error JSON decode")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Login != "user" || user.Password != "password" {
        w.WriteHeader(http.StatusUnauthorized)
		log.Println("Error Auth")
        return
    }

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &MyClaims{
        Login: user.Login,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   tokenString,
        Expires: expirationTime,
    })
}

func JwtMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
        if err != nil {
            if err == http.ErrNoCookie {
                w.WriteHeader(http.StatusUnauthorized) // Нет куки с токеном
				log.Println("cookie empty")
                return
            }
            w.WriteHeader(http.StatusBadRequest) // Ошибка при извлечении куки
			log.Println("Error read cookie")
            return
        }
		tokenStr := cookie.Value
        claims := &MyClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

		if err != nil {
            if err == jwt.ErrSignatureInvalid {
                w.WriteHeader(http.StatusUnauthorized)
                return
            }
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        if !token.Valid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

		if claims.ExpiresAt < time.Now().Unix() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "username", claims.Login)
		next.ServeHTTP(w, r.WithContext(ctx))
    })
}
