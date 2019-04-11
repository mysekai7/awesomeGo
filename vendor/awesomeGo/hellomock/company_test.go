package hellomock

import (
	"testing"
	"github.com/golang/mock/gomock"
	"awesomeGo/hellomock/mock"
)

func TestCompany_Meetig(t *testing.T) {
	person := NewPerson("前台小妹")
	company := NewCompany(person)
	t.Log(company.Meetig("花花"))

}

func TestCompany_Meetig2(t *testing.T) {
	ctl := gomock.NewController(t)
	mock_talker := mock_hellomock.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("花花")).Return("这是自定义的返回值，可以是任意类型")

	company := NewCompany(mock_talker)
	t.Log(company.Meetig("花花"))
	//t.Log(company.Meetig("明明"))
}
