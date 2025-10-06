package env

import (
	"net/url"
	"os"
	"slices"
	"strconv"
	"strings"
)

// GetEnvStr retrieves an environment variable as a string, returning a default value if not set.
//
// Parameters:
//   - key: The name of the environment variable to retrieve.
//   - byDefault: The default value to return if the environment variable is not set or is empty.
func GetEnvStr(key string, byDefault string) string {
	v := os.Getenv(key)
	if v == "" {
		return byDefault
	}
	return v
}

// GetEnvInt retrieves an environment variable as an int, returning a default value if not set.
//
// Parameters:
//   - key: The name of the environment variable to retrieve.
//   - byDefault: The default value to return if the environment variable is not set or is empty.
func GetEnvInt(key string, byDefault int) int {
	s := GetEnvStr(key, "")
	v, err := strconv.Atoi(s)
	if err != nil {
		return byDefault
	}
	return v
}

// GetEnvBool retrieves an environment variable as a boolean, returning a default value if not set or if parsing fails.
//
// Parameters:
//   - key: The name of the environment variable to retrieve.
//   - byDefault: The default value to return if the environment variable is not set or if parsing fails.
func GetEnvBool(key string, byDefault bool) bool {
	s := GetEnvStr(key, "")
	v, err := strconv.ParseBool(s)
	if err != nil {
		return byDefault
	}
	return v
}

// GetEnvSalt retrieves an environment variable as a salt string, ensuring it is exactly 32 characters long.
// If the environment variable is shorter than 32 characters, it repeats the value until it reaches 32 characters.
// If the environment variable is longer than 32 characters, it truncates the value to 32 characters.
// If the environment variable is not set, it returns a default salt value, which is also adjusted to be 32 characters long.
//
// Parameters:
//   - key: The name of the environment variable to retrieve.
//   - byDefault: The default value to use.
func GetEnvSalt(key, byDefault string) string {
	Salt := GetEnvStr(key, byDefault)
	if Salt == "" {
		Salt = byDefault
	}

	if Salt == "" {
		return ""
	}

	count := float64(32) / float64(len(Salt))
	if count > 1 {
		Salt = strings.Repeat(Salt, int(count+1))
	}
	if len(Salt) > 32 {
		Salt = Salt[:32]
	}

	return Salt
}

// GetEnvUrl retrieves an environment variable as a URL, ensuring it ends with a slash.
// If the environment variable is not set, it returns a default URL with a trailing slash.
// If the environment variable is set but does not contain a valid URL, it returns the default URL.
//
// Parameters:
//   - key: The name of the environment variable to retrieve.
//   - byDefault: The default URL to return if the environment variable is not set or is invalid.
func GetEnvUrl(key string, byDefault string) string {
	s := GetEnvStr(key, "")
	if s == "" {
		s = byDefault
	}

	u, err := url.Parse(s)
	if err != nil {
		s = byDefault
	} else {
		if !slices.Contains([]string{"http", "https", "wss"}, u.Scheme) {
			s = byDefault
		}

		if u.Host == "" {
			s = byDefault
		}

		if u.Hostname() == "" {
			s = byDefault
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

// GetEnvFloat retrieves an environment variable as a float64, returning a default value if not set or if parsing fails.
//
// Parameters:
//   - key: The name of the environment variable to retrieve.
//   - byDefault: The default value to return if the environment variable is not set or if parsing fails.
func GetEnvFloat(key string, byDefault float64) float64 {
	s := GetEnvStr(key, "")
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return byDefault
	}
	return v
}
