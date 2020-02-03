package functions

import (
	"testing"

	"github.com/elko-dev/spawn/web"
	"github.com/golang/mock/gomock"
)

type mockApp struct {
}

func (mockApp mockApp) Create() error {
	return nil
}
func TestWhenCreateIsCalledNewFunctionsTypeIsReturned(t *testing.T) {
	//arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := mockApp{}
	appFactory := web.NewMockAppFactory(ctrl)
	appFactory.EXPECT().Create().Return(expected, nil)

	factory := NewFactory(appFactory)
	//act
	platform, _ := factory.Create("", "")
	functions := platform.(FunctionsType)
	//assert
	if functions.project != expected {
		t.Log("Incorrect project returned")
		t.Fail()
	}
}
