package router

import (
	"net/http"
	"poe-debugserver/src/model"
	"reflect"
	"testing"
)

func TestGetCharacterHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name       string
		args       args
		wantResult []model.Character
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetCharacterHandler(tt.args.w, tt.args.r); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetCharacterHandler() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
