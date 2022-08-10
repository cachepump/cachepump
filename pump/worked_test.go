package pump

import (
	"reflect"
	"sort"
	"testing"

	"github.com/robfig/cron/v3"
)

func Test_getWorkindSourceNames(t *testing.T) {

	// Initialisation of a storege.
	storageCronIDs = map[string]cron.EntryID{
		"A": 1, "B": 2, "C": 3,
	}

	tests := []struct {
		name      string
		wantNames []string
	}{
		{
			name:      "Receive all names",
			wantNames: []string{"A", "B", "C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNames := getWorkindSourceNames()
			sort.Strings(gotNames)
			sort.Strings(tt.wantNames)
			if !reflect.DeepEqual(gotNames, tt.wantNames) {
				t.Errorf("getWorkindSourceNames() = %v, want %v", gotNames, tt.wantNames)
			}
		})
	}
}

func Test_getWorkedID(t *testing.T) {

	// Initialisation of a storege.
	storageCronIDs = map[string]cron.EntryID{
		"A": 1,
	}

	type args struct {
		name string
	}
	tests := []struct {
		name   string
		args   args
		wantId cron.EntryID
		wantOk bool
	}{
		{
			name:   "Receiving of an invalid key",
			args:   args{name: "B"},
			wantId: 0,
			wantOk: false,
		},
		{
			name:   "Receiving of a valid key",
			args:   args{name: "A"},
			wantId: 1,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotOk := getWorkedID(tt.args.name)
			if !reflect.DeepEqual(gotId, tt.wantId) {
				t.Errorf("getWorkedID() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotOk != tt.wantOk {
				t.Errorf("getWorkedID() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_setWorkedID(t *testing.T) {

	// Initialisation of a storege.
	storageCronIDs = map[string]cron.EntryID{}

	type args struct {
		name string
		id   cron.EntryID
	}
	tests := []struct {
		name        string
		args        args
		wantStorage StorageCronID
	}{
		{
			name:        "Set new key",
			args:        args{name: "NewKey", id: 321},
			wantStorage: map[string]cron.EntryID{"NewKey": 321},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setWorkedID(tt.args.name, tt.args.id)
		})
		if !reflect.DeepEqual(storageCronIDs, tt.wantStorage) {
			t.Errorf("After using setWorkedID() storageCronIDs = %v, wantStorage %v", storageCronIDs, tt.wantStorage)
		}
	}
}

func Test_delWorkindSource(t *testing.T) {

	// Initialisation of a storege.
	storageCronIDs = map[string]cron.EntryID{
		"badKey": 123,
	}

	type args struct {
		name string
	}
	tests := []struct {
		name        string
		args        args
		wantStorage StorageCronID
	}{
		{
			name:        "Removing a bad key",
			args:        args{name: "badKey"},
			wantStorage: map[string]cron.EntryID{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delWorkindSource(tt.args.name)
		})
		if !reflect.DeepEqual(storageCronIDs, tt.wantStorage) {
			t.Errorf("After using delWorkindSource() storageCronIDs = %v, wantStorage %v", storageCronIDs, tt.wantStorage)
		}
	}
}
