package functions

import (
	"testing"
)

const (
	applicationType = "apptype"
)

type mockApp struct {
}

func (mockApp mockApp) Create() error {
	return nil
}
func TestWhenCreateIsCalledNewFunctionsTypeIsReturned(t *testing.T) {
	//arrange
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// expected := mockApp{}
	// appFactory := web.NewMockAppFactory(ctrl)
	// appFactory.EXPECT().Create(applicationType).Return(expected, nil)

	// factory := NewFactory()
	// //act
	// platform, _ := factory.Create("", applicationType)
	// functions := platform.(FunctionsType)
	// //assert
	// if functions.project != expected {
	// 	t.Log("Incorrect project returned")
	// 	t.Fail()
	// }
}
