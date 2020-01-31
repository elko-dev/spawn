package applicationtype

import (
	"testing"
)

const (
	appType = "SOMETYPE"
)

func TestWhenApplicationTypeIsSelectedItIsPassedToAppropriateFactory(t *testing.T) {
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// mockAppFactory := NewMockApplicationFactory(ctrl)
	// mockPrompt := NewMockPrompt(ctrl)
	// mockTempAppType := NewMockTempAppType(ctrl)
	// //set mock expectations
	// mockPrompt.EXPECT().ForType().Return(appType, nil)
	// mockAppFactory.EXPECT().Create(appType).Return(mockTempAppType, nil)
	// if "" == nil {
	// 	t.Log("Factory output was null")
	// 	t.Fail()
	// }
}
