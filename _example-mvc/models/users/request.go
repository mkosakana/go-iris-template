package users

type Request struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       uint32 `json:"age"`
}
