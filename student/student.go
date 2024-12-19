package student

type student struct {
	name string
}

func NewUser(name string) *student {
	return &student{name: name}
}

func (u *student) Name() string {
	return u.name
}
