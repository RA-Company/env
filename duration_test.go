package env

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseDuration(t *testing.T) {
	valid := map[string]time.Duration{
		"7d":            7 * 24 * time.Hour,
		"2h 30m":        2*time.Hour + 30*time.Minute,
		"2h30m":         2*time.Hour + 30*time.Minute,
		"1d 2h 30m 15s": 24*time.Hour + 2*time.Hour + 30*time.Minute + 15*time.Second,
		"90":            90 * time.Second,
		"  5M  ":        5 * time.Minute,
		"1D":            24 * time.Hour,
	}
	for in, want := range valid {
		got, err := ParseDuration(in)
		require.NoError(t, err, in)
		require.Equal(t, want, got, in)
	}

	invalid := []string{
		"", "   ", "0", "0s", "abc", "-5s", "1.5h", "5x", "1h 2h", "1h 30",
		"h", "1h abc", "99999d", "99999999999999999999s",
	}
	for _, in := range invalid {
		_, err := ParseDuration(in)
		require.Error(t, err, in)
	}
}

func TestGetEnvDuration(t *testing.T) {
	t.Run("valid value returned", func(t *testing.T) {
		t.Setenv("TEST_DUR", "2h 30m")
		require.Equal(t, 2*time.Hour+30*time.Minute, GetEnvDuration("TEST_DUR", time.Second))
	})

	t.Run("number is seconds", func(t *testing.T) {
		t.Setenv("TEST_DUR", "90")
		require.Equal(t, 90*time.Second, GetEnvDuration("TEST_DUR", time.Second))
	})

	t.Run("unset returns default", func(t *testing.T) {
		require.Equal(t, time.Minute, GetEnvDuration("TEST_DUR_UNSET_XYZ", time.Minute))
	})

	t.Run("empty returns default", func(t *testing.T) {
		t.Setenv("TEST_DUR", "")
		require.Equal(t, time.Minute, GetEnvDuration("TEST_DUR", time.Minute))
	})

	t.Run("invalid returns default", func(t *testing.T) {
		t.Setenv("TEST_DUR", "abc")
		require.Equal(t, time.Minute, GetEnvDuration("TEST_DUR", time.Minute))
	})

	t.Run("zero returns default", func(t *testing.T) {
		t.Setenv("TEST_DUR", "0s")
		require.Equal(t, time.Minute, GetEnvDuration("TEST_DUR", time.Minute))
	})
}
