package usecase

import (
	"context"
	"cosmart-library/domain"
	"cosmart-library/mocks"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_usecase_GetBookList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		ctx     context.Context
		request domain.GetBookListRequest
	}
	tests := []struct {
		name         string
		u            *usecase
		args         args
		wantResponse domain.GetBookListResponse
		wantErr      bool
	}{
		{
			name: "test-1_success",
			u: &usecase{
				repository: func() domain.LibraryRepositoryInterface {
					mock := mocks.NewMockLibraryRepositoryInterface(mockCtrl)
					mock.EXPECT().FetchRawBooksBySubject(gomock.Any(), "test").Return(
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
			args: args{
				ctx: context.TODO(),
				request: domain.GetBookListRequest{
					Subject: "test",
				},
			},
			wantResponse: domain.GetBookListResponse{
				Books: []domain.Book{
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
				Page: 1,
			},
		},
		{
			name: "test-2_default",
			u: &usecase{
				repository: func() domain.LibraryRepositoryInterface {
					mock := mocks.NewMockLibraryRepositoryInterface(mockCtrl)
					mock.EXPECT().FetchRawBooksBySubject(gomock.Any(), "fiction").Return(
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
			args: args{
				ctx: context.TODO(),
				request: domain.GetBookListRequest{
					Subject: "",
				},
			},
			wantResponse: domain.GetBookListResponse{
				Books: []domain.Book{
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
				Page: 1,
			},
		},
		{
			name: "test-3_error",
			u: &usecase{
				repository: func() domain.LibraryRepositoryInterface {
					mock := mocks.NewMockLibraryRepositoryInterface(mockCtrl)
					mock.EXPECT().FetchRawBooksBySubject(gomock.Any(), "fiction").Return(nil, errors.New("err")).Times(1)
					return mock
				}(),
			},
			args: args{
				ctx: context.TODO(),
				request: domain.GetBookListRequest{
					Subject: "",
				},
			},
			wantResponse: domain.GetBookListResponse{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.u.GetBookList(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetBookList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("usecase.GetBookList() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
