package authenticator

type UserProfile struct {
	Id   string `json:"sub"`
	Name string `json:"name"`
}
