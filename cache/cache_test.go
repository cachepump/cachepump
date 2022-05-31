package cache

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	type args struct {
		key   string
		value []byte
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantData map[string][]byte
	}{
		{
			name:    "Set first value",
			args:    args{key: "key1", value: []byte("value1")},
			wantErr: false,
			wantData: map[string][]byte{
				"key1": []byte("value1"),
			},
		},
		{
			name:    "Set second value",
			args:    args{key: "key2", value: []byte("value2")},
			wantErr: false,
			wantData: map[string][]byte{
				"key1": []byte("value1"),
				"key2": []byte("value2"),
			},
		},
		{
			name:    "Replace first value",
			args:    args{key: "key1", value: []byte("new_value")},
			wantErr: false,
			wantData: map[string][]byte{
				"key1": []byte("new_value"),
				"key2": []byte("value2"),
			},
		},
	}
	data = make(map[string][]byte)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(data, tt.wantData) {
				t.Errorf("Internal data store %v, wantData %v", data, tt.wantData)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		args      args
		wantValue []byte
		wantErr   bool
	}{
		{
			name: "Receive of exist value by key",
			args: args{
				key: "key1",
			},
			wantValue: []byte("value1"),
			wantErr:   false,
		},
		{
			name: "Receive of not exist value by key",
			args: args{
				key: "key2",
			},
			wantValue: []byte{},
			wantErr:   true,
		},
	}
	data = map[string][]byte{"key1": []byte("value1")}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestDel(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name     string
		args     args
		wantData map[string][]byte
	}{
		{
			name: "Delete of second key",
			args: args{
				key: "key2",
			},
			wantData: map[string][]byte{
				"key1": []byte("value1"),
				"key3": []byte("value3"),
			},
		},
	}
	data = map[string][]byte{
		"key1": []byte("value1"),
		"key2": []byte("value2"),
		"key3": []byte("value3"),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Del(tt.args.key)
			if !reflect.DeepEqual(data, tt.wantData) {
				t.Errorf("Internal data store %v, wantData %v", data, tt.wantData)
			}
		})
	}
}
