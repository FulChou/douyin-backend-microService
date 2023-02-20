package pack

import (
	"douyin_backend_microService/user/dal/db"
	"douyin_backend_microService/user/kitex_gen/userdemo"
)

func ConcertUser(user *db.User) (res *userdemo.User) {
	if user == nil {
		return nil
	}

	return &userdemo.User{
		Id:            int64(user.ID),
		Name:          user.UserName,
		FollowCount:   int64(user.FollowCount),
		FollowerCount: int64(user.FollowerCount),
	}

}

func ConcertUsers(users []*db.User) (res []*userdemo.User) {
	if users == nil || len(users) == 0 {
		return nil
	}

	res = make([]*userdemo.User, 0)
	for _, user := range users {
		if u := ConcertUser(user); u != nil {
			res = append(res, u)
		}
	}

	return res
}
