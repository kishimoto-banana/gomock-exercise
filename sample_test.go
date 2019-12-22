package sample

import (
	"testing"

	mock "./mock_sample"
	"github.com/golang/mock/gomock"
)

func TestSample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSample := mock.NewMockSample(ctrl)
	mockSample.EXPECT().Method("foo").Return(1)

	t.Log("result: ", sampleMock.Method("foo"))
}
