package repository

import (
	"context"
	"cosmart-library/domain"
	mocks_repository "cosmart-library/mocks/library/repository"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_repository_FetchRawBooksBySubject(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		ctx     context.Context
		subject string
	}
	tests := []struct {
		name      string
		r         *repository
		args      args
		wantBooks []domain.Book
		wantErr   bool
	}{
		{
			name: "test-1_success",
			args: args{
				ctx:     context.TODO(),
				subject: "test",
			},
			r: &repository{
				api: func() APIInterface {
					mock := mocks_repository.NewMockAPIInterface(mockCtrl)
					mock.EXPECT().FetchBookListBySubject(gomock.Any(), "test").Return(
						[]domain.Book{
							{
								ID:            "12",
								Subject:       "fiction",
								Title:         "book title",
								Author:        []string{"author1", "author2"},
								EditionNumber: "1123",
							},
							{
								ID:            "13",
								Subject:       "fiction",
								Title:         "book title 2",
								Author:        []string{"author3"},
								EditionNumber: "1124",
							},
						}, nil,
					).Times(1)
					return mock
				}(),
			},
			wantBooks: []domain.Book{
				{
					ID:            "12",
					Subject:       "fiction",
					Title:         "book title",
					Author:        []string{"author1", "author2"},
					EditionNumber: "1123",
				},
				{
					ID:            "13",
					Subject:       "fiction",
					Title:         "book title 2",
					Author:        []string{"author3"},
					EditionNumber: "1124",
				},
			},
		},
		{
			name: "test-2_error",
			args: args{
				ctx:     context.TODO(),
				subject: "test",
			},
			r: &repository{
				api: func() APIInterface {
					mock := mocks_repository.NewMockAPIInterface(mockCtrl)
					mock.EXPECT().FetchBookListBySubject(gomock.Any(), "test").Return(
						nil, errors.New("err"),
					).Times(1)
					return mock
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBooks, err := tt.r.FetchRawBooksBySubject(tt.args.ctx, tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.FetchRawBooksBySubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBooks, tt.wantBooks) {
				t.Errorf("repository.FetchRawBooksBySubject() = %v, want %v", gotBooks, tt.wantBooks)
			}
		})
	}
}
