package env

import (
	"net/url"
	"os"
	"slices"
	"strconv"
	"strings"
)

// GetEnvStr returns the value of the environment variable named by key,
// or def if the variable is not set or is empty.
func GetEnvStr(key string, def string) string {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	return v
}

// GetEnvInt returns the value of the environment variable named by key parsed as int,
// or def if the variable is not set, is empty, or cannot be parsed.
func GetEnvInt(key string, def int) int {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	vInt, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return vInt
}

// GetEnvBool returns the value of the environment variable named by key parsed as bool,
// or def if the variable is not set, is empty, or cannot be parsed.
// Accepted values: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
func GetEnvBool(key string, def bool) bool {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	vBool, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}
	return vBool
}

// GetEnvSalt returns a 32-byte key derived from an environment variable,
// suitable for use as an AES-256 symmetric key or similar fixed-size secret.
//
// The raw value is stretched or truncated to exactly 32 bytes:
//   - shorter than 32 bytes: the value is repeated until it reaches at least
//     32 bytes, then truncated to exactly 32.
//   - longer than 32 bytes: truncated to the first 32 bytes.
//   - exactly 32 bytes: returned unchanged.
//
// If the variable is not set or is empty, def is used as the source material.
// If both are empty, an empty string is returned.
//
// Note: this function operates on bytes, not Unicode code points.
// Multibyte UTF-8 characters may be split at the 32-byte boundary.
// Use ASCII values for predictable results.
func GetEnvSalt(key, def string) string {
	salt := GetEnvStr(key, def)
	if salt == "" {
		return ""
	}

	if len(salt) < 32 {
		salt = strings.Repeat(salt, (32/len(salt))+1)
	}
	if len(salt) > 32 {
		salt = salt[:32]
	}

	return salt
}

// GetEnvUrl returns the value of the environment variable named by key as a URL with a trailing slash.
// The URL must have an http, https, or wss scheme and a non-empty hostname; otherwise def is used.
// If the variable is not set or is empty, def is used. If def is also empty, an empty string is returned.
func GetEnvUrl(key string, def string) string {
	s := GetEnvStr(key, "")
	if s == "" {
		s = def
	}

	u, err := url.Parse(s)
	if err != nil {
		s = def
	} else {
		if !slices.Contains([]string{"http", "https", "wss"}, u.Scheme) {
			s = def
		}

		if u.Hostname() == "" {
			s = def
		}
	}

	if s == "" {
		return s
	}

	if s[len(s)-1] != '/' {
		s += "/"
	}

	return s
}

// GetEnvFloat returns the value of the environment variable named by key parsed as float64,
// or def if the variable is not set, is empty, or cannot be parsed.
func GetEnvFloat(key string, def float64) float64 {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	vFloat, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return def
	}
	return vFloat
}
