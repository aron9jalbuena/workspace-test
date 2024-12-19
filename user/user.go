package user

type user struct {
	name string
}

func NewUser(name string) *user {
	return &user{name: name}
}

func (u *user) Name() string {
	return u.name
}
