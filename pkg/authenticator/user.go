package authenticator

type User struct {
	Provider  string
	Email     string
	Name      string
	FirstName string
	LastName  string
	NickName  string
	UserID    string
	AvatarURL string
}
