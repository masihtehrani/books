package dtosignup

type Request struct {
	FullName  string `json:"full_name"`
	Pseudonym string `json:"pseudonym"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Response struct {
	Success bool `json:"success"`
}
