package config

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

func TestConfig_NewConfig(t *testing.T) {
	type fields map[string]string
	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name: "New Config",
			fields: fields{
				"PINKIE_HTTP_PORT":   "80",
				"PINKIE_DB_HOST":     "localhost",
				"PINKIE_DB_PORT":     "13306",
				"PINKIE_DB_USER":     "user",
				"PINKIE_DB_PASSWORD": "password",
				"PINKIE_DB_NAME":     "name",
			},
			want: &Config{
				HTTPPort:   80,
				DBHost:     "localhost",
				DBPort:     13306,
				DBUser:     "user",
				DBPassword: "password",
				DBName:     "name",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, val := range tt.fields {
				t.Setenv(key, val)
			}

			actual, err := NewConfig()

			require.NoError(t, err)
			assert.Equal(t, tt.want, actual)
		})
	}
}
