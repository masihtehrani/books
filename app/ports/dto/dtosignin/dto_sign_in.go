package dtosignin

import "time"

type Request struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Response struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}
