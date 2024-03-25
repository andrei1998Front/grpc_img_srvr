package storage

import (
	"reflect"
	"testing"

	"github.com/andrei1998Front/grpc_img_srvr/internal/domain/models"
)

func Test_prepareListImages(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.ImgInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prepareListImages(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("prepareListImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareListImages() = %v, want %v", got, tt.want)
			}
		})
	}
}
