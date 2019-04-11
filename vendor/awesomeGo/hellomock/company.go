package hellomock

type Company struct {
	Usher Talker
}

func NewCompany(t Talker) *Company {
	return &Company{
		Usher:t,
	}
}

func (c *Company)Meetig(guestName string) string  {
	return c.Usher.SayHello(guestName)
}
