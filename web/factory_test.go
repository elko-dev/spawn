package web

import (
	reflect "reflect"
	"testing"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/constants"
	gomock "github.com/golang/mock/gomock"
)

func TestWhenServerIsNodeJsWebTypeContainsNodeServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPrompt := NewMockPrompt(ctrl)
	mockNodeJsFactory := NewMockAppFactory(ctrl)
	mockReactFactory := NewMockAppFactory(ctrl)

	mockPrompt.EXPECT().getClientType().Return(constants.ReactClientLanguageType, nil)
	mockPrompt.EXPECT().getServerType().Return(constants.NodeServerType, nil)
	mockNodeJsFactory.EXPECT().Create().Return(applications.NodeJs{})
	mockReactFactory.EXPECT().Create().Return(applications.React{})
	factory := Factory{mockNodeJsFactory, mockReactFactory, mockPrompt}

	webType := factory.Create()
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
	mockReactFactory := NewMockAppFactory(ctrl)

	mockPrompt.EXPECT().getClientType().Return(constants.ReactClientLanguageType, nil)
	mockPrompt.EXPECT().getServerType().Return(constants.NodeServerType, nil)
	mockNodeJsFactory.EXPECT().Create().Return(applications.NodeJs{})
	mockReactFactory.EXPECT().Create().Return(applications.React{})
	factory := Factory{mockNodeJsFactory, mockReactFactory, mockPrompt}

	webType := factory.Create()
	client := webType.Client

	if !isReactType(client) {
		t.Log("Incorrect client type returned, got ", reflect.TypeOf(client))
		t.Fail()
	}

}

func isNodeJsType(t interface{}) bool {
	switch t.(type) {

	case applications.NodeJs:
		return true
	default:
		return false
	}
}

func isReactType(t interface{}) bool {
	switch t.(type) {

	case applications.React:
		return true
	default:
		return false
	}

}
