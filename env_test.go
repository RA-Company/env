package env

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnvStr(t *testing.T) {
	t.Run("set value returned", func(t *testing.T) {
		t.Setenv("TEST_STR", "hello")
		require.Equal(t, "hello", GetEnvStr("TEST_STR", "default"))
	})

	t.Run("unset returns default", func(t *testing.T) {
		require.Equal(t, "default", GetEnvStr("TEST_STR_UNSET_XYZ", "default"))
	})

	t.Run("empty var returns default", func(t *testing.T) {
		t.Setenv("TEST_STR_EMPTY", "")
		require.Equal(t, "default", GetEnvStr("TEST_STR_EMPTY", "default"))
	})

	t.Run("empty default and unset returns empty", func(t *testing.T) {
		require.Equal(t, "", GetEnvStr("TEST_STR_UNSET_XYZ", ""))
	})

	t.Run("whitespace value returned as-is", func(t *testing.T) {
		t.Setenv("TEST_STR_SPACE", "   ")
		require.Equal(t, "   ", GetEnvStr("TEST_STR_SPACE", "default"))
	})
}

func TestGetEnvInt(t *testing.T) {
	t.Run("valid integer returned", func(t *testing.T) {
		t.Setenv("TEST_INT", "42")
		require.Equal(t, 42, GetEnvInt("TEST_INT", 0))
	})

	t.Run("unset returns default", func(t *testing.T) {
		require.Equal(t, 99, GetEnvInt("TEST_INT_UNSET_XYZ", 99))
	})

	t.Run("invalid string returns default", func(t *testing.T) {
		t.Setenv("TEST_INT_BAD", "not_a_number")
		require.Equal(t, 5, GetEnvInt("TEST_INT_BAD", 5))
	})

	t.Run("float string returns default", func(t *testing.T) {
		t.Setenv("TEST_INT_FLOAT", "3.14")
		require.Equal(t, 0, GetEnvInt("TEST_INT_FLOAT", 0))
	})

	t.Run("negative integer returned", func(t *testing.T) {
		t.Setenv("TEST_INT_NEG", "-7")
		require.Equal(t, -7, GetEnvInt("TEST_INT_NEG", 0))
	})

	t.Run("zero value returned", func(t *testing.T) {
		t.Setenv("TEST_INT_ZERO", "0")
		require.Equal(t, 0, GetEnvInt("TEST_INT_ZERO", 99))
	})
}

func TestGetEnvBool(t *testing.T) {
	trueCases := []string{"1", "t", "T", "TRUE", "true", "True"}
	for _, v := range trueCases {
		v := v
		t.Run("true value: "+v, func(t *testing.T) {
			t.Setenv("TEST_BOOL_V", v)
			require.True(t, GetEnvBool("TEST_BOOL_V", false))
		})
	}

	falseCases := []string{"0", "f", "F", "FALSE", "false", "False"}
	for _, v := range falseCases {
		v := v
		t.Run("false value: "+v, func(t *testing.T) {
			t.Setenv("TEST_BOOL_V", v)
			require.False(t, GetEnvBool("TEST_BOOL_V", true))
		})
	}

	t.Run("unset returns default true", func(t *testing.T) {
		require.True(t, GetEnvBool("TEST_BOOL_UNSET_XYZ", true))
	})

	t.Run("unset returns default false", func(t *testing.T) {
		require.False(t, GetEnvBool("TEST_BOOL_UNSET_XYZ", false))
	})

	t.Run("invalid string returns default", func(t *testing.T) {
		t.Setenv("TEST_BOOL_BAD", "yes")
		require.True(t, GetEnvBool("TEST_BOOL_BAD", true))
	})
}

func TestGetEnvSalt(t *testing.T) {
	t.Run("result is always 32 chars when input non-empty", func(t *testing.T) {
		t.Setenv("TEST_SALT", "abc")
		got := GetEnvSalt("TEST_SALT", "")
		require.Equal(t, 32, len(got))
	})

	t.Run("short value padded by repetition", func(t *testing.T) {
		t.Setenv("TEST_SALT", "ab")
		got := GetEnvSalt("TEST_SALT", "")
		require.Equal(t, 32, len(got))
		require.Equal(t, "abababababababababababababababababab"[:32], got)
	})

	t.Run("exactly 32 chars unchanged", func(t *testing.T) {
		exact := "abcdefghijklmnopqrstuvwxyz012345"
		require.Equal(t, 32, len(exact))
		t.Setenv("TEST_SALT", exact)
		got := GetEnvSalt("TEST_SALT", "")
		require.Equal(t, exact, got)
	})

	t.Run("longer than 32 chars truncated to 32", func(t *testing.T) {
		long := "abcdefghijklmnopqrstuvwxyz0123456789"
		t.Setenv("TEST_SALT", long)
		got := GetEnvSalt("TEST_SALT", "")
		require.Equal(t, 32, len(got))
		require.Equal(t, long[:32], got)
	})

	t.Run("single char repeated 32 times", func(t *testing.T) {
		t.Setenv("TEST_SALT", "x")
		got := GetEnvSalt("TEST_SALT", "")
		require.Equal(t, 32, len(got))
		require.Equal(t, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", got)
	})

	t.Run("unset uses default", func(t *testing.T) {
		got := GetEnvSalt("TEST_SALT_UNSET_XYZ", "key")
		require.Equal(t, 32, len(got))
	})

	t.Run("empty env and empty default returns empty", func(t *testing.T) {
		require.Equal(t, "", GetEnvSalt("TEST_SALT_UNSET_XYZ", ""))
	})
}

func TestGetEnvUrl(t *testing.T) {
	t.Run("valid http url gets trailing slash", func(t *testing.T) {
		t.Setenv("TEST_URL", "http://example.com")
		require.Equal(t, "http://example.com/", GetEnvUrl("TEST_URL", ""))
	})

	t.Run("valid https url gets trailing slash", func(t *testing.T) {
		t.Setenv("TEST_URL", "https://example.com")
		require.Equal(t, "https://example.com/", GetEnvUrl("TEST_URL", "https://default.url"))
	})

	t.Run("already has trailing slash unchanged", func(t *testing.T) {
		t.Setenv("TEST_URL", "https://example.com/")
		require.Equal(t, "https://example.com/", GetEnvUrl("TEST_URL", ""))
	})

	t.Run("valid wss scheme accepted", func(t *testing.T) {
		t.Setenv("TEST_URL", "wss://ws.example.com")
		require.Equal(t, "wss://ws.example.com/", GetEnvUrl("TEST_URL", ""))
	})

	t.Run("url with path preserved", func(t *testing.T) {
		t.Setenv("TEST_URL", "https://example.com/api/v1")
		require.Equal(t, "https://example.com/api/v1/", GetEnvUrl("TEST_URL", ""))
	})

	t.Run("url with port accepted", func(t *testing.T) {
		t.Setenv("TEST_URL", "https://example.com:8443")
		require.Equal(t, "https://example.com:8443/", GetEnvUrl("TEST_URL", ""))
	})

	t.Run("invalid scheme returns default", func(t *testing.T) {
		t.Setenv("TEST_URL", "ftp://example.com")
		require.Equal(t, "https://default.url/", GetEnvUrl("TEST_URL", "https://default.url"))
	})

	t.Run("non-url string returns default", func(t *testing.T) {
		t.Setenv("TEST_URL", "not-a-url")
		require.Equal(t, "https://default.url/", GetEnvUrl("TEST_URL", "https://default.url"))
	})

	t.Run("unset returns default with slash", func(t *testing.T) {
		require.Equal(t, "https://default.url/", GetEnvUrl("TEST_URL_UNSET_XYZ", "https://default.url"))
	})

	t.Run("empty env and empty default returns empty", func(t *testing.T) {
		require.Equal(t, "", GetEnvUrl("TEST_URL_UNSET_XYZ", ""))
	})
}

func TestGetEnvFloat(t *testing.T) {
	t.Run("valid float returned", func(t *testing.T) {
		t.Setenv("TEST_FLOAT", "3.14")
		require.InDelta(t, 3.14, GetEnvFloat("TEST_FLOAT", 0), 1e-9)
	})

	t.Run("unset returns default", func(t *testing.T) {
		require.Equal(t, 1.5, GetEnvFloat("TEST_FLOAT_UNSET_XYZ", 1.5))
	})

	t.Run("invalid string returns default", func(t *testing.T) {
		t.Setenv("TEST_FLOAT_BAD", "abc")
		require.Equal(t, 2.0, GetEnvFloat("TEST_FLOAT_BAD", 2.0))
	})

	t.Run("negative float returned", func(t *testing.T) {
		t.Setenv("TEST_FLOAT_NEG", "-0.5")
		require.InDelta(t, -0.5, GetEnvFloat("TEST_FLOAT_NEG", 0), 1e-9)
	})

	t.Run("integer string parsed as float", func(t *testing.T) {
		t.Setenv("TEST_FLOAT_INT", "10")
		require.Equal(t, 10.0, GetEnvFloat("TEST_FLOAT_INT", 0))
	})

	t.Run("zero value returned", func(t *testing.T) {
		t.Setenv("TEST_FLOAT_ZERO", "0")
		require.Equal(t, 0.0, GetEnvFloat("TEST_FLOAT_ZERO", 99.9))
	})
}
