package user

type Group struct{}

func (g *Group) User() *SUser {
	return &insUser
}
