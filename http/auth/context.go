package auth

type contextKey string

func (c contextKey) String() string {
	return "auth context key " + string(c)
}
