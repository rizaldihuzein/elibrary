package usecase

import (
	"cosmart-library/domain"
	"cosmart-library/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockInf := mocks.NewMockLibraryRepositoryInterface(mockCtrl)

	type args struct {
		repository domain.LibraryRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want domain.LibraryUsecaseInterface
	}{
		{
			name: "test-1_success",
			args: args{
				repository: mockInf,
			},
			want: &usecase{
				repository: mockInf,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
