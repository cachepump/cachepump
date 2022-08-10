package pump

import (
	"reflect"
	"testing"

	"github.com/cachepump/cachepump/provider"
)

func TestSource_getProvider(t *testing.T) {
	type fields struct {
		Rule      string
		StaticSrc provider.Static
		HttpSrc   provider.Http
		FileSrc   provider.File
	}
	tests := []struct {
		name    string
		fields  fields
		wantPrv provider.Provider
	}{
		{
			name: "All empty providers",
			fields: fields{
				Rule:      "* * * * * *",
				StaticSrc: provider.Static{},
				HttpSrc:   provider.Http{},
				FileSrc:   provider.File{},
			},
			wantPrv: provider.EmptyProvider{},
		},
		{
			name: "First not empty provider",
			fields: fields{
				Rule:      "* * * * * *",
				StaticSrc: provider.Static{Value: "no_empty"},
				HttpSrc:   provider.Http{},
				FileSrc:   provider.File{},
			},
			wantPrv: provider.Static{Value: "no_empty"},
		},
		{
			name: "Second not empty provider",
			fields: fields{
				Rule:      "* * * * * *",
				StaticSrc: provider.Static{},
				HttpSrc:   provider.Http{Endpoint: "0.0.0.0"},
				FileSrc:   provider.File{},
			},
			wantPrv: provider.Http{Endpoint: "0.0.0.0"},
		},
		{
			name: "Third not empty provider",
			fields: fields{
				Rule:      "* * * * * *",
				StaticSrc: provider.Static{},
				HttpSrc:   provider.Http{},
				FileSrc:   provider.File{Path: "no_empty.txt"},
			},
			wantPrv: provider.File{Path: "no_empty.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := Source{
				Rule:      tt.fields.Rule,
				StaticSrc: tt.fields.StaticSrc,
				HttpSrc:   tt.fields.HttpSrc,
				FileSrc:   tt.fields.FileSrc,
			}
			if gotPrv := src.getProvider(); !reflect.DeepEqual(gotPrv, tt.wantPrv) {
				t.Errorf("Source.getProvider() = %v, want %v", gotPrv, tt.wantPrv)
			}
		})
	}
}
