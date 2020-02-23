package mobile

import (
	"testing"

	"github.com/elko-dev/spawn/applications"
	"github.com/golang/mock/gomock"
)

func TestMobileTypeCallsClientOnlyWhenIncludeBackendIsFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := applications.NewMockProject(ctrl)
	mockServer := applications.NewMockProject(ctrl)

	mockClient.EXPECT().Create().Return(nil)
	mockServer.EXPECT().Create().MaxTimes(0)

	mobileType := NewMobileType(mockClient, mockServer, false)

	err := mobileType.Create()

	if err != nil {
		t.Log("Error found when none existed ", err)
		t.Fail()
	}

}

func TestMobileTypeCallsClientAndServerWhenIncludeBackendIsFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := applications.NewMockProject(ctrl)
	mockServer := applications.NewMockProject(ctrl)

	mockClient.EXPECT().Create().Return(nil)
	mockServer.EXPECT().Create().Return(nil)

	mobileType := NewMobileType(mockClient, mockServer, true)

	err := mobileType.Create()

	if err != nil {
		t.Log("Error found when none existed ", err)
		t.Fail()
	}

}
