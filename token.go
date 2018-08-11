package auth

//stores a short time password (default 1 week) so the user doesn't have
//to type it's password every time.
type Token struct {
	Token  string
	Ip     string
	Expire int64
}
