package restapi

import (
	"context"
	"cosmart-library/domain"
	"reflect"
	"testing"
)

func Test_api_FetchBookListBySubject(t *testing.T) {
	type args struct {
		ctx     context.Context
		subject string
	}
	tests := []struct {
		name      string
		a         *api
		args      args
		wantBooks []domain.Book
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBooks, err := tt.a.FetchBookListBySubject(tt.args.ctx, tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("api.FetchBookListBySubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBooks, tt.wantBooks) {
				t.Errorf("api.FetchBookListBySubject() = %v, want %v", gotBooks, tt.wantBooks)
			}
		})
	}
}
