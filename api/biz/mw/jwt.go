package mw

import (
	"context"
	"douyin_backend_microService/api/biz/rpc"
	"douyin_backend_microService/pkg/constants"
	"douyin_backend_microService/pkg/errno"
	"douyin_backend_microService/user/kitex_gen/userdemo"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(constants.SecretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   constants.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			userid, _ := claims[constants.IdentityKey].(json.Number).Int64()
			return &userdemo.User{
				Id: userid,
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req userdemo.CheckUserRequest

			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Name) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{
				Name:     req.Name,
				Password: req.Password,
			})
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			claims := jwt.ExtractClaims(ctx, c)
			userId := int64(claims[constants.IdentityKey].(float64))
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": errno.SuccessCode,
				"token":       token,
				"user_id":     userId,
				"expire":      expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{

				"code":    errno.AuthorizationFailedErr.ErrCode,
				"message": message,
			})
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
		ParseOptions: []jwtv4.ParserOption{jwtv4.WithJSONNumber()},
	})
}
