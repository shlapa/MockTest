package main

import (
	"checksmt/ttt"
	repo_mocks "checksmt/ttt/mocks"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestHandler_GetUnit1(t *testing.T) {
	expectErr := false
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedRepo := repo_mocks.NewMockIRepo(ctrl)

	mockedRepo.EXPECT().Get("someTestUnit").Return("any", nil)

	handler := NewHandler(&Dependency{
		Repo: mockedRepo,
	})

	str, err := handler.GetUnit("someTestUnit")
	if err != nil {
		if expectErr {
			t.Log(err)
		} else {
			t.Fatal(err)
		}
	}

	t.Log(str)
}

func TestHandler_GetUnit3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedRepo := repo_mocks.NewMockIRepo(ctrl)

	mockedRepo.EXPECT().GetTime().Return(time.Now().String(), nil)
	handler := NewHandler(&Dependency{
		Repo: mockedRepo,
	})
	str, err := handler.GetTimeVarUnit()
	if err != nil {
		t.Log(err)
	}
	t.Log(str)
}

type fields struct {
	rrr ttt.IRepo
}

func TestHandler_GetUnit4(t *testing.T) {

	tests := []struct {
		name        string
		string_test string
		times       time.Time
	}{
		{
			name:        "Test1",
			string_test: "hello world",
		},
		{
			name:        "Test2",
			string_test: "hi",
		},
		{
			name:        "Test3",
			string_test: "holla",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockedRepo := repo_mocks.NewMockIRepo(ctrl)
			handler := NewHandler(&Dependency{
				Repo: mockedRepo,
			})
			mockedRepo.EXPECT().Get(tt.string_test).Return(tt.string_test, nil)
			str, err := handler.GetUnit(tt.string_test)
			if err != nil {
				t.Error(err)
			}
			if len(str) == 0 {
				t.Error("len==0")
			}
		})
	}
}
func TestHandler_GetUnit5(t *testing.T) {

	tests := []struct {
		name string
		s
	}{
		{
			name: "Test1",
		},
		{
			name: "Test2",
		},
		{
			name: "Test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockedRepo := repo_mocks.NewMockIRepo(ctrl)
			handler := NewHandler(&Dependency{
				Repo: mockedRepo,
			})
			mockedRepo.EXPECT().GetTime().Return(time.Now().String(), nil)
			str, err := handler.GetTimeVarUnit()
			if err != nil {
				t.Error(err)
			}
			if len(str) == 0 {
				t.Error("len==0")
			}
		})
	}
}
