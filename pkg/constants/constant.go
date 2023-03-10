// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package constants

const (
	VideoTableName              = "video"
	UserTableName               = "user"
	FollowTableName             = "follow"
	CommentTableName            = "comment"
	FriendTableName             = "friend"
	SecretKey                   = "secret key"
	IdentityKey                 = "id"
	StatusCode                  = "status_code"
	StatusMsg                   = "status_msg"
	User                        = "user"
	UserID                      = "user_id"
	Token                       = "token"
	VideoList                   = "video_list"
	NextTime                    = "next_time"
	MinioBucketName             = "dousheng"
	MinioEndPoint               = "127.0.0.1:9000"
	MinioAccessID               = "admin"
	MinioAccessKey              = "12345678"
	Location                    = "GuangZhou"
	ApiServiceName              = "demoapi"
	VideoServiceName            = "videodemo"
	UserServiceName             = "userdemo"
	RelationServiceName         = "relationdemo"
	MessageServiceName          = "messagedemo"
	MySQLDefaultDSN             = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                 = "127.0.0.1:2379"
	CPURateLimit        float64 = 80.0
	DefaultLimit                = 10
	ExportEndpoint              = "localhost:4317"
	TCP                         = "tcp"
	UserServiceAddr             = ":2000"
	VideoServiceAddr            = ":10000"
	RelationServiceAddr         = ":3000"
	MessageServiceAddr          = ":4000"
)
