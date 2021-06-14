package dtodeletebook

type Request struct {
	UserID string
	BookID string
}

type Response struct {
	Success bool
}
