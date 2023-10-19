package auth

import (
	"context"
	_ "embed"
	"log"
	"net/url"
	"sync"

	"github.com/nicus101/KittyPlantMonitor-server/pkg/common/inmem"
)

//go:embed testdata/inmem_user.json
var inMemUsers []byte

var inMemLoadOnce sync.Once

func inMemLoad() {
	inMemLoadOnce.Do(func() {
		inmem.LoadFromJsonBytes("users", inMemUsers)
	})
}

func inMemLogIn(_ context.Context, login, password string) (User, error) {
	inMemLoad()

	users := inmem.ListFromTable("users", url.Values{
		"login":    []string{login},
		"password": []string{password},
	})

	switch l := len(users); l {
	case 0:
		return User{}, ErrUserNotFound
	default:
		log.Println("More than one valid user in mock data!", users)
		fallthrough
	case 1:
		// intentionaly empty
	}

	user := users[0]
	return User{
		Id:      int(user["id"].(float64)),
		Name:    user["login"].(string),
		isAdmin: user["admin"].(bool),
	}, nil
}
