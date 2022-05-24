package utils

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func (p *Person) Fullname() string {
	return p.Firstname + " " + p.Lastname
}
