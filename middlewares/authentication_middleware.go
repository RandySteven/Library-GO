package middlewares

import (
	"context"
	"github.com/RandySteven/Library-GO/enums"
	jwt2 "github.com/RandySteven/Library-GO/pkg/jwt"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 || auth == "" {
			utils.ResponseHandler(w, http.StatusUnauthorized, `Invalid get token from auth`, nil, nil, nil)
			return
		}
		tokenStr := auth[len("Bearer "):]
		if tokenStr == "" {
			utils.ResponseHandler(w, http.StatusUnauthorized, `Invalid token failed to split from bearer`, nil, nil, nil)
			return
		}
		claims := &jwt2.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(j *jwt.Token) (interface{}, error) {
			return jwt2.JwtKey, nil
		})
		if err != nil || !token.Valid {
			utils.ResponseHandler(w, http.StatusUnauthorized, `Invalid token`, nil, nil, err)
			return
		}
		ctx := context.WithValue(r.Context(), enums.UserID, claims.UserID)
		ctx2 := context.WithValue(ctx, enums.RoleID, claims.RoleID)
		r = r.WithContext(ctx2)
		next.ServeHTTP(w, r)
	})
}
