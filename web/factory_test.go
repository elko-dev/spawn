package web

import (
	reflect "reflect"
	"testing"

	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/nodejs"
	gomock "github.com/golang/mock/gomock"
)

const (
	applicationType = "SOME_TYPE"
)

func TestWhenServerIsNodeJsWebTypeContainsNodeServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPrompt := NewMockPrompt(ctrl)
	mockNodeJsFactory := NewMockAppFactory(ctrl)
	// mockReactFactory := NewMockAppFactory(ctrl)

	mockPrompt.EXPECT().ForClientType(applicationType).Return(constants.ReactClientLanguageType, nil)
	mockPrompt.EXPECT().ForServerType().Return(constants.NodeServerType, nil)
	//TODO: Hacked to call twice until react implement
	mockNodeJsFactory.EXPECT().Create().Return(nodejs.NewNode(nil, nil, "", ""), nil)
	mockNodeJsFactory.EXPECT().Create().Return(nodejs.NewNode(nil, nil, "", ""), nil)
	// mockReactFactory.EXPECT().Create().Return(nil)
	factory := Factory{mockNodeJsFactory, mockPrompt}

	webType := factory.Create(applicationType)
	client := webType.Server

	if !isNodeJsType(client) {
		t.Log("Incorrect client type returned, got ", reflect.TypeOf(client))
		t.Fail()
	}

}

func TestWhenClientIsReactWebTypeContainsReactClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPrompt := NewMockPrompt(ctrl)
	mockNodeJsFactory := NewMockAppFactory(ctrl)
	// mockReactFactory := NewMockAppFactory(ctrl)

	mockPrompt.EXPECT().ForClientType(applicationType).Return(constants.ReactClientLanguageType, nil)
	mockPrompt.EXPECT().ForServerType().Return(constants.NodeServerType, nil)
	mockNodeJsFactory.EXPECT().Create().Return(nodejs.NewNode(nil, nil, "", ""), nil)
	mockNodeJsFactory.EXPECT().Create().Return(nodejs.NewNode(nil, nil, "", ""), nil)
	// mockReactFactory.EXPECT().Create().Return(nil)
	factory := Factory{mockNodeJsFactory, mockPrompt}

	webType := factory.Create(applicationType)
	client := webType.Client

	if !isReactType(client) {
		t.Log("Incorrect client type returned, got ", reflect.TypeOf(client))
		t.Fail()
	}

}

func isNodeJsType(t interface{}) bool {
	switch t.(type) {

	case nodejs.Node:
		return true
	default:
		return false
	}
}

func isReactType(t interface{}) bool {
	switch t.(type) {

	case nodejs.Node:
		return true
	default:
		return false
	}

}
