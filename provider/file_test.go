package provider

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cachepump/cachepump/cache"
)

const (
	testFileName = "../data_for_test_file_provider.txt"
	testFileData = "123 test 123"
)

func TestFile_IsEmpty(t *testing.T) {
	type fields struct {
		Path string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Empty file provider",
			fields: fields{Path: ""},
			want:   true,
		},
		{
			name:   "No empty file provider",
			fields: fields{Path: "../test.txt"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := File{
				Path: tt.fields.Path,
			}
			if got := f.IsEmpty(); got != tt.want {
				t.Errorf("File.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Pump(t *testing.T) {

	// Initialisation of environments.
	if err := initFileProvider(); err != nil {
		t.Errorf("Test initialisation error = %v", err)
	}
	defer cleaneFileProvider()

	type fields struct {
		Path string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantCacheValue []byte
		wantCacheErr   bool
	}{
		{
			name:           "Target file is not exist",
			fields:         fields{Path: "./no_file.txt"},
			args:           args{name: "no_file"},
			wantCacheValue: []byte{},
			wantCacheErr:   true,
		},
		{
			name:           "Target file with test data",
			fields:         fields{Path: testFileName},
			args:           args{name: testFileName},
			wantCacheValue: []byte(testFileData),
			wantCacheErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := File{
				Path: tt.fields.Path,
			}
			fn := f.Pump(tt.args.name)
			fn()
			value, err := cache.Get(tt.args.name)
			if (err != nil) != tt.wantCacheErr {
				t.Errorf("Job function generated by File.Pump() returns error = %v", err)
			}
			if string(value) != string(tt.wantCacheValue) {
				t.Errorf("Job function generated by File.Pump() returns value = %s, want = %s", value, tt.wantCacheValue)
			}
		})
	}
}

func initFileProvider() (err error) {
	return ioutil.WriteFile(testFileName, []byte(testFileData), fs.ModePerm)
}

func cleaneFileProvider() {
	os.Remove(testFileName)
}
