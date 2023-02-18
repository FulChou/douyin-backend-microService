package dal

import (
	"douyin_backend_microService/user/dal/db"
)

func Init() {
	db.Init() // mysql init
}
