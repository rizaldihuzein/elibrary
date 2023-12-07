package usecase

import (
	"context"
	"cosmart-library/domain"
	"cosmart-library/mocks"
	"cosmart-library/pick-up/types"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_usecase_MakeBookOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockTime := time.Now()

	type args struct {
		ctx    context.Context
		pickup domain.PickupOrder
	}
	tests := []struct {
		name    string
		u       *usecase
		args    args
		wantId  string
		wantErr bool
	}{
		{
			name: "test-1_success",
			args: args{
				ctx: context.Background(),
				pickup: domain.PickupOrder{
					ID: "id",
					Books: []domain.Book{
						{
							ID:            "someid",
							Title:         "title",
							Author:        []string{"1", "2"},
							EditionNumber: "22",
						},
					},
					PickupDate: mockTime,
				},
			},
			u: &usecase{
				repository: func() domain.PickupRepositoryInterface {
					mock := mocks.NewMockPickupRepositoryInterface(mockCtrl)
					mock.EXPECT().UpsertBookOrder(gomock.Any(), domain.PickupOrder{
						ID: "id",
						Books: []domain.Book{
							{
								ID:            "someid",
								Title:         "title",
								Author:        []string{"1", "2"},
								EditionNumber: "22",
							},
						},
						PickupDate: mockTime,
						ReturnDate: mockTime.Add(7 * 24 * time.Hour),
					}).Return(nil).Times(1)
					return mock
				}(),
			},
			wantId: "id",
		},
		{
			name: "test-2_error",
			args: args{
				ctx: context.Background(),
				pickup: domain.PickupOrder{
					ID: "id",
					Books: []domain.Book{
						{
							ID:            "someid",
							Title:         "title",
							Author:        []string{"1", "2"},
							EditionNumber: "22",
						},
					},
					PickupDate: mockTime,
				},
			},
			u: &usecase{
				repository: func() domain.PickupRepositoryInterface {
					mock := mocks.NewMockPickupRepositoryInterface(mockCtrl)
					mock.EXPECT().UpsertBookOrder(gomock.Any(), domain.PickupOrder{
						ID: "id",
						Books: []domain.Book{
							{
								ID:            "someid",
								Title:         "title",
								Author:        []string{"1", "2"},
								EditionNumber: "22",
							},
						},
						PickupDate: mockTime,
						ReturnDate: mockTime.Add(7 * 24 * time.Hour),
					}).Return(errors.New("err")).Times(1)
					return mock
				}(),
			},
			wantErr: true,
		},
		{
			name: "test-3_idempotent_update",
			args: args{
				ctx: context.WithValue(context.Background(), types.RequestIDKey, "testkey"),
				pickup: domain.PickupOrder{
					ID: "id",
					Books: []domain.Book{
						{
							ID:            "someid",
							Title:         "title",
							Author:        []string{"1", "2"},
							EditionNumber: "22",
						},
					},
					PickupDate: mockTime,
				},
			},
			u: &usecase{
				repository: func() domain.PickupRepositoryInterface {
					mock := mocks.NewMockPickupRepositoryInterface(mockCtrl)
					mock.EXPECT().GetIdempotency(gomock.Any(), "testkey").Return(mockTime.Add(1*time.Hour), nil).Times(1)
					return mock
				}(),
			},
			wantId: "id",
		},
		{
			name: "test-4_idempotent_insert",
			args: args{
				ctx: context.WithValue(context.Background(), types.RequestIDKey, "testkey"),
				pickup: domain.PickupOrder{
					Books: []domain.Book{
						{
							ID:            "someid",
							Title:         "title",
							Author:        []string{"1", "2"},
							EditionNumber: "22",
						},
					},
					PickupDate: mockTime,
				},
			},
			u: &usecase{
				repository: func() domain.PickupRepositoryInterface {
					mock := mocks.NewMockPickupRepositoryInterface(mockCtrl)
					mock.EXPECT().GetIdempotency(gomock.Any(), "testkey").Return(mockTime.Add(1*time.Hour), nil).Times(1)
					return mock
				}(),
			},
			wantId: "idempotent request",
		},
		{
			name: "test-5_idempotent_insert_err",
			args: args{
				ctx: context.WithValue(context.Background(), types.RequestIDKey, "testkey"),
				pickup: domain.PickupOrder{
					Books: []domain.Book{
						{
							ID:            "someid",
							Title:         "title",
							Author:        []string{"1", "2"},
							EditionNumber: "22",
						},
					},
					PickupDate: mockTime,
				},
			},
			u: &usecase{
				repository: func() domain.PickupRepositoryInterface {
					mock := mocks.NewMockPickupRepositoryInterface(mockCtrl)
					var ttl time.Time
					mock.EXPECT().GetIdempotency(gomock.Any(), "testkey").Return(ttl, errors.New("err")).Times(1)
					return mock
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := tt.u.MakeBookOrder(tt.args.ctx, tt.args.pickup)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.MakeBookOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("usecase.MakeBookOrder() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func Test_usecase_GetBookOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name      string
		u         *usecase
		args      args
		wantOrder domain.PickupOrder
		wantErr   bool
	}{
		{
			name: "test-1_success",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			u: &usecase{
				repository: func() domain.PickupRepositoryInterface {
					mock := mocks.NewMockPickupRepositoryInterface(mockCtrl)
					mock.EXPECT().GetBookOrder(gomock.Any(), "id").Return(domain.PickupOrder{
						ID: "id",
						Books: []domain.Book{
							{
								ID:            "someid",
								Title:         "title",
								Author:        []string{"1", "2"},
								EditionNumber: "22",
							},
						},
					}, nil).Times(1)
					return mock
				}(),
			},
			wantOrder: domain.PickupOrder{
				ID: "id",
				Books: []domain.Book{
					{
						ID:            "someid",
						Title:         "title",
						Author:        []string{"1", "2"},
						EditionNumber: "22",
					},
				},
			},
		},
		{
			name: "test-2_error",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			u: &usecase{
				repository: func() domain.PickupRepositoryInterface {
					mock := mocks.NewMockPickupRepositoryInterface(mockCtrl)
					mock.EXPECT().GetBookOrder(gomock.Any(), "id").Return(domain.PickupOrder{}, errors.New("err")).Times(1)
					return mock
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrder, err := tt.u.GetBookOrder(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetBookOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("usecase.GetBookOrder() = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}
