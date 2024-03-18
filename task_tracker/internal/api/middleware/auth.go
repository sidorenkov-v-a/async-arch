package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"

	"async-arch/task_tracker/internal/pkg/domain"
	"async-arch/task_tracker/pkg/env"
)

type authMiddleware struct {
	env       env.JWT
	usersRepo domain.UserRepository
}

func NewAuthMiddleware(env env.JWT, usersRepo domain.UserRepository) *authMiddleware {
	return &authMiddleware{
		env:       env,
		usersRepo: usersRepo,
	}
}

func (a *authMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("X-Bearer")

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(a.env.Secret), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := a.usersRepo.GetByEmail(r.Context(), fmt.Sprint(claims["email"]))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		r.WithContext(context.WithValue(r.Context(), "userID", user.ID))
		r.Header.Add("userID", user.ID.String())

		next.ServeHTTP(w, r)
	})
}
