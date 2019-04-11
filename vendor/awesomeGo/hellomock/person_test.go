package hellomock

import "testing"

func TestPerson_SayHello(t *testing.T) {
	person := NewPerson("小白")
	t.Log(person.SayHello("小黑"))
}
