package urltool

import (
	"net/url"
	"reflect"
	"testing"
)

func TestIsValidHref(t *testing.T) {
	type args struct {
		href string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidHref(tt.args.href); got != tt.want {
				t.Errorf("IsValidHref() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearHref(t *testing.T) {
	type args struct {
		href string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearHref(tt.args.href); got != tt.want {
				t.Errorf("ClearHref() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertHref2URL(t *testing.T) {
	type args struct {
		href    string
		currURL *url.URL
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertHref2URL(tt.args.href, tt.args.currURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertHref2URL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertHref2URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
