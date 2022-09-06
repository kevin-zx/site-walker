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
		{
			name: "javascript",
			args: args{
				href: "javascript:alert(1)",
			},
			want: false,
		},
		{
			name: "mailto",
			args: args{
				href: "mailto:",
			},
			want: false,
		},
		{
			name: "https",
			args: args{
				href: "https://www.baidu.com",
			},
			want: true,
		},
		{
			name: "http",
			args: args{
				href: "http://www.baidu.com",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidHref(tt.args.href); got != tt.want {
				t.Errorf("IsValidHref() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCleanHref(t *testing.T) {
	type args struct {
		href string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "suffix double slash",
			args: args{
				href: "https://www.baidu.com//",
			},
			want: "https://www.baidu.com/",
		},
		{
			name: "prefix double slash",
			args: args{
				href: "//www.baidu.com/",
			},
			want: "//www.baidu.com/",
		},
		{
			name: "middle double slash",
			args: args{
				href: "https://www.baidu.com//aa/",
			},
			want: "https://www.baidu.com/aa/",
		},
		{
			name: "tail utf8 space",
			args: args{
				href: "https://www.baidu.com/aa/%20",
			},
			want: "https://www.baidu.com/aa/",
		},
		{
			name: "unicode space",
			args: args{
				href: "https://www.baidu.com/aa/　",
			},
			want: "https://www.baidu.com/aa/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanHref(tt.args.href); got != tt.want {
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
		{
			name: "relative path",
			args: args{
				href:    "/aa/bb",
				currURL: &url.URL{Scheme: "https", Host: "www.baidu.com"},
			},
			want: &url.URL{
				Scheme: "https",
				Host:   "www.baidu.com",
				Path:   "/aa/bb",
			},
			wantErr: false,
		},
		{
			name: "absolute path",
			args: args{
				href:    "https://www.baidu.com/aa/bb",
				currURL: &url.URL{Scheme: "https", Host: "www.baidu.com"},
			},
			want: &url.URL{
				Scheme: "https",
				Host:   "www.baidu.com",
				Path:   "/aa/bb",
			},
			wantErr: false,
		},
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
