// Code generated by hertz generator.

package api

import (
	"context"
	api "douyin_backend_microService/api/biz/model/api"
	"douyin_backend_microService/api/biz/mw"
	"douyin_backend_microService/api/biz/rpc"
	"douyin_backend_microService/pkg/constants"
	"douyin_backend_microService/pkg/errno"
	"douyin_backend_microService/relation/kitex_gen/relationdemo"
	"douyin_backend_microService/user/kitex_gen/userdemo"
	"douyin_backend_microService/video/kitex_gen/videodemo"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/jwt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// / RegisterUser .
// @router /douyin/user/register [POST]
func RegisterUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	fmt.Printf(req.Username)
	if err != nil {
		SendErrResponse(c, err)
		return
	}

	if len(req.Username) == 0 || len(req.Password) == 0 {
		SendErrResponse(c, errno.ParamErr)
		return
	}

	userId, err := rpc.CreateUser(ctx, &userdemo.CreateUserRequest{
		Name:     req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendErrResponse(c, err)
		return
	}

	resp := new(api.DouyinUserRegisterResponse)
	resp.UserId = userId
	resp.StatusCode = 0

	c.JSON(consts.StatusOK, resp)
}

// LoginUser .
// @router /douyin/user/login [POST]
func LoginUser(ctx context.Context, c *app.RequestContext) {

	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// Feed .
// @router /douyin/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {

	var err error
	var req api.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))

	feed, nexttime, err := rpc.Feed(ctx, &videodemo.FeedRequest{
		LatestTime: req.LatestTime,
		UserID:     userID,
	})

	c.JSON(consts.StatusOK, map[string]interface{}{
		constants.StatusCode: 0, constants.VideoList: feed, constants.NextTime: nexttime,
	})
}

// PublishAction .
// @router /douyin/publish/action [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinPublishActionResponse)

	if len(req.Token) == 0 || len(req.Title) == 0 {
		SendErrResponse(c, errno.ParamErr)
	}
	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	fmt.Println("get userID, ", userId)
	//TODO video byte to url
	data, err := c.FormFile("data")
	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	fileName := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%s_%s", strconv.FormatInt(userId, 16), req.Title, fileName)
	savePath := filepath.Join("../../../static/", finalName)

	// save to local
	if err := c.SaveUploadedFile(data, savePath); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	// save to minio
	bucketName := "dousheng"
	if err := rpc.FileUploader(ctx, bucketName, finalName, savePath); err != nil {
		fmt.Println("save to minio failed")
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	// get URL from minio
	url, err := rpc.GetFileUrl(bucketName, finalName, 0)
	if err != nil {
		log.Printf("get url failed")
	} else {
		log.Printf("User uploaded a file, %s", url)
	}

	if err := rpc.Publish(ctx, &videodemo.PublishRequest{Playurl: url.String(), Title: req.Title, UserId: userId}); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "upload successful."
	c.JSON(consts.StatusOK, &resp)

}

// PulishVideoList .
// @router /douyin/publish/action [GET]
func PulishVideoList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	publishList, err := rpc.GetPublishList(ctx, &videodemo.PublishListRequest{UserId: req.UserId})

	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
	}

	c.JSON(consts.StatusOK, map[string]interface{}{
		constants.StatusCode: 0, constants.VideoList: publishList,
	})
}

// RelationAction .
// @router /douyin/relation/action [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))

	err = rpc.Relation(ctx, &relationdemo.RelationActionRequest{
		UserId:     userId,
		ToUserId:   req.ToUserId,
		ActionType: int64(req.ActionType),
	})
	resp := new(api.DouyinRelationActionResponse)

	if err != nil {
		klog.Errorf("rpc relation action error")
		SendErrResponse(c, err)
	}
	resp.StatusCode = 0

	c.JSON(consts.StatusOK, resp)
}

// GetRelationFollowList .
// @router /douyin/relation/follow/list [GET]
func GetRelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))

	followList, err := rpc.GetFollowList(ctx, &relationdemo.RelationFollowListRequest{UserId: userId})
	if err != nil {
		klog.Errorf("rpc get follow error, %s", err.Error())
		SendErrResponse(c, err)
		return
	}

	resp := new(api.DouyinRelationFollowListResponse)
	resp.UserList = followList
	resp.StatusCode = 0
	c.JSON(consts.StatusOK, resp)
}

// GetRelationFollowerList .
// @router /douyin/relation/follower/list [GET]
func GetRelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))

	followList, err := rpc.GetFollowerList(ctx, &relationdemo.RelationFollowerListRequest{UserId: userId})
	if err != nil {
		klog.Errorf("rpc get follow error, %s", err.Error())
		SendErrResponse(c, err)
		return
	}

	resp := new(api.DouyinRelationFollowerListResponse)
	resp.UserList = followList
	resp.StatusCode = 0
	c.JSON(consts.StatusOK, resp)

}
