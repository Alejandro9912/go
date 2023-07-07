package users

import (
	"fmt"
	"time"

	"github.com/Alejandro9912/go/models"
)

func AltaUser() {
	u := new(models.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
