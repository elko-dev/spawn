package functions

import (
	"testing"

	"github.com/elko-dev/spawn/platform"
	"github.com/elko-dev/spawn/web"
	"github.com/golang/mock/gomock"
)

type mockApp struct {
}

func (mockApp mockApp) Create(application platform.Application) error {
	return nil
}
func TestWhenCreateIsCalledNewFunctionsTypeIsReturned(t *testing.T) {
	//arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := mockApp{}
	appFactory := web.NewMockAppFactory(ctrl)
	appFactory.EXPECT().Create().Return(expected)

	factory := NewFactory(appFactory)
	//act
	actual := factory.Create().function
	//assert
	if actual != expected {
		t.Log("Incorrect function returned")
		t.Fail()
	}
}
