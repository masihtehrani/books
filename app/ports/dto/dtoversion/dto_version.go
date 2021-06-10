package dtoversion

type Response struct {
	VersionTag    string `json:"tag"`
	VersionCommit string `json:"commit"`
	VersionDate   string `json:"date"`
	ServiceName   string `json:"service"`
}
