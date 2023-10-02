package interceptor

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"

	"connectrpc.com/connect"
	auth0 "github.com/auth0-community/go-auth0"
	// "github.com/go-jose/go-jose"
	// "github.com/go-jose/go-jose/v3/jwt"
)

// リクエストのJWTを検証する。成功時にはUserIDをコンテキストにセットする
func NewAuthInterceptor(auth0_url string, audience string) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {

			// TODO: contextから取得しないで良いか？
			// token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
			token := req.Header().Get("Authorization")
			if token == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("error: invalid token"))
			}
			token = strings.TrimPrefix(token, "Bearer")
			token = strings.TrimSpace(token)
			fmt.Println("token: ", token)

			parsedToken, err := jwt.ParseSigned(token)
			if err != nil {
				// return nil, grpc.Errorf(codes.Unauthenticated, "Cannot parse token because of", err)
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: "https://" + auth0_url + "/.well-known/jwks.json"}, nil)
			configuration := auth0.NewConfiguration(client, []string{audience}, "https://"+auth0_url+"/", jose.RS256)
			validator := auth0.NewValidator(configuration, nil)

			if err := validator.ValidateToken(parsedToken); err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			claims := make(map[string]interface{})
			validator.Claims(parsedToken, &claims)
			fmt.Println("claims: ", claims)

			// TODO: contextにユーザーIDをセットする？

			// return context.WithValue(ctx, "userId", claims["sub"]), nil
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
