package env

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Setup env variables for testing "TEST_STR" = "Is String"
func TestGetEnvStr(t *testing.T) {
	type args struct {
		str string
		def string
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		{"Valid", args{"TEST_STR", "string"}, "Is String"},
		{"Default", args{"TEST_STR1", "string"}, "string"},
		{"Empty", args{"TEST_STR_EMPTY", ""}, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetEnvStr(tc.args.str, tc.args.def)
			if got != tc.want {
				require.Equal(t, tc.want, got, "getEnvStr(%q, %q) = %q; want: %q", tc.args.str, tc.args.def, got, tc.want)
			}
		})
	}
}

// Setup env variables for testing "TEST_INT" = "10"
func TestGetEnvInt(t *testing.T) {
	type args struct {
		str string
		def int
	}

	testCases := []struct {
		name string
		args args
		want int
	}{
		{"Valid", args{"TEST_INT", 20}, 10},
		{"Default", args{"TEST_INT1", 20}, 20},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetEnvInt(tc.args.str, tc.args.def)
			if got != tc.want {
				require.Equal(t, tc.want, got, "GetEnvInt() = %d, want %d", got, tc.want)
			}
		})
	}
}

// Setup env variables for testing "TEST_BOOL" = false
func TestGetEnvBool(t *testing.T) {
	type args struct {
		str string
		def bool
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{"Valid", args{"TEST_BOOL", true}, false},
		{"Default", args{"TEST_BOOL1", true}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetEnvBool(tc.args.str, tc.args.def)
			if got != tc.want {
				require.Equal(t, tc.want, got, "GetEnvBool() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetEnvSalt(t *testing.T) {
	type args struct {
		str string
		def string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Valid", args{"TEST_SALT", "default_salt"}, "Is StringIs StringIs StringIs String"},
		{"Default", args{"TEST_SALT1", "default_salt"}, "default_saltdefault_saltdefault_saltdefault_salt"},
		{"Long", args{"TEST_SALT_LONG", "This is very long salt string"}, "This is very long salt str"},
		{"Empty", args{"TEST_SALT_EMPTY", ""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvSalt(tt.args.str, tt.args.def); got != tt.want {
				require.Equal(t, 32, len(got), "GetEnvSalt() length = %d; want: 32", len(got))
				require.Equal(t, tt.want[:26], got[:26], "GetEnvSalt() = %v, want %v", got[:26], tt.want[:26])
			}
		})
	}
}

func TestGetEnvUrl(t *testing.T) {
	type args struct {
		key       string
		byDefault string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Valid", args{"TEST_URL", "https://default.url"}, "https://example.com/"},
		{"Default", args{"TEST_URL1", "https://default.url"}, "https://default.url/"},
		{"Invalid", args{"TEST_URL_INVALID", "https://default.url"}, "https://default.url/"},
		{"Empty", args{"TEST_URL_EMPTY", ""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvUrl(tt.args.key, tt.args.byDefault); got != tt.want {
				require.Equal(t, tt.want, got, "GetEnvUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvFloat(t *testing.T) {
	type args struct {
		key       string
		byDefault float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Valid", args{"TEST_FLOAT", 1.0}, 3.14},
		{"Default", args{"TEST_FLOAT1", 1.0}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvFloat(tt.args.key, tt.args.byDefault); got != tt.want {
				require.Equal(t, tt.want, got, "GetEnvFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
