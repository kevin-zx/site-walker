package seo

import (
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestExtractSEOTextInfo(t *testing.T) {
	type args struct {
		html *goquery.Selection
	}
	tests := []struct {
		name string
		args args
		want SEOText
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractSEOTextInfo(tt.args.html); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractSEOTextInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
