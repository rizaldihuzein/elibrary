package memorycache

import (
	"context"
	"cosmart-library/domain"
	mocks_repository "cosmart-library/mocks/pickup/repository"
	"cosmart-library/pkg/memcache"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_memorycache_UpsertBookOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		ctx    context.Context
		pickup domain.PickupOrder
	}
	tests := []struct {
		name    string
		m       *memorycache
		args    args
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
				},
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Upsert(gomock.Any(), "id", domain.PickupOrder{
						ID: "id",
						Books: []domain.Book{
							{
								ID:            "someid",
								Title:         "title",
								Author:        []string{"1", "2"},
								EditionNumber: "22",
							},
						},
					}).Return(nil).Times(1)
					return mock
				}(),
			},
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
				},
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Upsert(gomock.Any(), "id", domain.PickupOrder{
						ID: "id",
						Books: []domain.Book{
							{
								ID:            "someid",
								Title:         "title",
								Author:        []string{"1", "2"},
								EditionNumber: "22",
							},
						},
					}).Return(errors.New("err")).Times(1)
					return mock
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.UpsertBookOrder(tt.args.ctx, tt.args.pickup); (err != nil) != tt.wantErr {
				t.Errorf("memorycache.UpsertBookOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_memorycache_GetBookOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name       string
		m          *memorycache
		args       args
		wantPickup domain.PickupOrder
		wantErr    bool
	}{
		{
			name: "test-1_success",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Get(gomock.Any(), "id").Return(domain.PickupOrder{
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
			wantPickup: domain.PickupOrder{
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
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Get(gomock.Any(), "id").Return(nil, errors.New("err")).Times(1)
					return mock
				}(),
			},
			wantErr: true,
		},
		{
			name: "test-3_error_nil",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Get(gomock.Any(), "id").Return(nil, memcache.ErrNil).Times(1)
					return mock
				}(),
			},
			wantPickup: domain.PickupOrder{},
		},
		{
			name: "test-4_nil",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Get(gomock.Any(), "id").Return(nil, nil).Times(1)
					return mock
				}(),
			},
			wantPickup: domain.PickupOrder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPickup, err := tt.m.GetBookOrder(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("memorycache.GetBookOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPickup, tt.wantPickup) {
				t.Errorf("memorycache.GetBookOrder() = %v, want %v", gotPickup, tt.wantPickup)
			}
		})
	}
}

func Test_memorycache_StoreIdempotency(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockTime := time.Now()

	type args struct {
		ctx context.Context
		key string
		ttl time.Time
	}
	tests := []struct {
		name    string
		m       *memorycache
		args    args
		wantErr bool
	}{
		{
			name: "test-1_success",
			args: args{
				ctx: context.Background(),
				key: "key",
				ttl: mockTime,
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Upsert(gomock.Any(), "id_key", mockTime).Return(nil).Times(1)
					return mock
				}(),
			},
		},
		{
			name: "test-2_error",
			args: args{
				ctx: context.Background(),
				key: "key",
				ttl: mockTime,
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Upsert(gomock.Any(), "id_key", mockTime).Return(errors.New("err")).Times(1)
					return mock
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.StoreIdempotency(tt.args.ctx, tt.args.key, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("memorycache.StoreIdempotency() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_memorycache_GetIdempotency(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockTime := time.Now()

	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		m       *memorycache
		args    args
		wantTtl time.Time
		wantErr bool
	}{
		{
			name: "test-1_success",
			args: args{
				ctx: context.Background(),
				key: "id",
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Get(gomock.Any(), "id_id").Return(mockTime, nil).Times(1)
					return mock
				}(),
			},
			wantTtl: mockTime,
		},
		{
			name: "test-2_error",
			args: args{
				ctx: context.Background(),
				key: "id",
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Get(gomock.Any(), "id_id").Return(nil, errors.New("err")).Times(1)
					return mock
				}(),
			},
			wantErr: true,
		},
		{
			name: "test-3_error_nil",
			args: args{
				ctx: context.Background(),
				key: "id",
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Get(gomock.Any(), "id_id").Return(nil, memcache.ErrNil).Times(1)
					return mock
				}(),
			},
		},
		{
			name: "test-4_nil",
			args: args{
				ctx: context.Background(),
				key: "id",
			},
			m: &memorycache{
				persistentDriver: func() DriverInterface {
					mock := mocks_repository.NewMockDriverInterface(mockCtrl)
					mock.EXPECT().Get(gomock.Any(), "id_id").Return(nil, nil).Times(1)
					return mock
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTtl, err := tt.m.GetIdempotency(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("memorycache.GetIdempotency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTtl, tt.wantTtl) {
				t.Errorf("memorycache.GetIdempotency() = %v, want %v", gotTtl, tt.wantTtl)
			}
		})
	}
}
