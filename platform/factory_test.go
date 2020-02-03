package platform

import (
	"testing"
)

const (
	projectName = "nane"
	applType    = "type"
)

func TestWhenHerokuIsSelectedHerokuPlatformIsReturned(t *testing.T) {
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// mockPrompt := NewMockPrompt(ctrl)
	// mockHerokuPlatform := NewMockPlatformFactory(ctrl)
	// mockFunctionsPlatform := NewMockPlatformFactory(ctrl)

	// mockPrompt.EXPECT().forPlatformType().Return(constants.HerokuPlatform, nil)
	// mockHerokuPlatform.EXPECT().Create(projectName, applType).Return(herokuplatform.Heroku{}, nil)
	// mockFunctionsPlatform.EXPECT().Create(projectName, applType).MaxTimes(0)
	// factory := NewFactory(mockPrompt, mockHerokuPlatform, mockFunctionsPlatform)

	// platform, _ := factory.Create(projectName, applType)

	// if !isHerokuType(platform) {
	// 	t.Log("Expected heroku type returned")
	// 	t.Fail()
	// 	return
	// }
}

func TestWhenFunctionsIsSelectedFunctionsPlatformIsReturned(t *testing.T) {
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// mockPrompt := NewMockPrompt(ctrl)
	// mockHerokuPlatform := NewMockPlatformFactory(ctrl)
	// mockFunctionsPlatform := NewMockPlatformFactory(ctrl)

	// mockPrompt.EXPECT().forPlatformType().Return(constants.AzureFunctions, nil)
	// mockFunctionsPlatform.EXPECT().Create(projectName, applType).Return(functions.FunctionsType{}, nil)
	// mockHerokuPlatform.EXPECT().Create(projectName, applType).MaxTimes(0)

	// factory := NewFactory(mockPrompt, mockHerokuPlatform, mockFunctionsPlatform)

	// platform, _ := factory.Create(projectName, applType)

	// if !isFunctionsType(platform) {
	// 	t.Log("Expected function type returned")
	// 	t.Fail()
	// 	return
	// }
}

// func isHerokuType(t interface{}) bool {
// 	switch t.(type) {

// 	case herokuplatform.Heroku:
// 		return true
// 	default:
// 		return false
// 	}

// }

// func isFunctionsType(t interface{}) bool {
// 	switch t.(type) {

// 	case functions.FunctionsType:
// 		return true
// 	default:
// 		return false
// 	}

// }
