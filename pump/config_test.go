package pump

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cachepump/cachepump/provider"
)

func Test_uploadConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name       string
		args       args
		wantConfig Config
		wantErr    bool
	}{
		{
			name:       "File with configuration is not exist",
			args:       args{file: "../no_config.yml"},
			wantConfig: Config{},
			wantErr:    true,
		},
		{
			name: "Valid file with configuration",
			args: args{file: "../config.yml"},
			wantConfig: Config{
				Version: "1.0",
				Sources: map[string]Source{
					"static_key": {
						Rule:      "* * * * * *",
						StaticSrc: provider.Static{Value: "test_value"},
					},
					"count_202103": {
						Rule: "0 */2 * * * *",
						HttpSrc: provider.Http{
							Endpoint: "http://0.0.0.0:8123",
							Method:   "POST",
							Auth:     provider.Auth{User: "admin", Password: "adminadmin"},
							Body:     "SELECT date, count() FROM DB.Raw_Data PREWHERE toYYYYMM(date) = 202103 GROUP BY date ORDER BY date\n",
						},
					},
					"file_go.sum": {
						Rule: "0 */2 * * * *",
						FileSrc: provider.File{
							Path: "go.sum",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := uploadConfig(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("uploadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			if cfg := getConfig(); !reflect.DeepEqual(cfg, tt.wantConfig) {
				fmt.Println(cfg.Sources["ga_count_202103:"].HttpSrc.Body) //todo
				t.Errorf("getConfig() = %v, wantConfig %v", cfg, tt.wantConfig)
			}
		})
	}
}
