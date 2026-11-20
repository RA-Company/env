package env

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// maxDuration is the longest time ParseDuration accepts: far more than any setting needs, and far from an overflow.
const maxDuration = 100 * 365 * 24 * time.Hour

// GetEnvDuration returns the value of the environment variable named by key parsed as time.Duration
// (see ParseDuration), or def if the variable is not set, is empty, or cannot be parsed.
func GetEnvDuration(key string, def time.Duration) time.Duration {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}

	d, err := ParseDuration(v)
	if err != nil {
		return def
	}

	return d
}

// ParseDuration parses a time in a human format: parts with the units d (days), h (hours), m (minutes)
// and s (seconds), in any case, divided by spaces or written together: "7d", "2h 30m", "2h30m", "1d 2h 30m 15s".
// A number without a unit is seconds ("90" is 90 seconds). The numbers are whole.
//
// A time that is empty, is zero or negative, has an unknown unit or a unit twice ("1h 2h"),
// is longer than 100 years or is not a time at all is an error.
func ParseDuration(s string) (time.Duration, error) {
	str := strings.ToLower(strings.TrimSpace(s))
	if str == "" {
		return 0, errors.New("the time is empty")
	}

	var (
		total time.Duration
		seen  = map[byte]bool{}
	)
	for rest := str; rest != ""; {
		i := 0
		for i < len(rest) && rest[i] >= '0' && rest[i] <= '9' {
			i++
		}
		if i == 0 {
			return 0, fmt.Errorf("%q: a number is expected at %q", s, rest)
		}
		digits := rest[:i]
		rest = rest[i:]

		unit := byte('s') // no unit: seconds, but only for the whole value (see below)
		if rest != "" && rest[0] != ' ' {
			unit = rest[0]
			rest = rest[1:]
		} else if len(seen) > 0 || strings.TrimSpace(rest) != "" {
			return 0, fmt.Errorf("%q: the number %s has no unit", s, digits)
		}

		var factor time.Duration
		switch unit {
		case 'd':
			factor = 24 * time.Hour
		case 'h':
			factor = time.Hour
		case 'm':
			factor = time.Minute
		case 's':
			factor = time.Second
		default:
			return 0, fmt.Errorf("%q: unknown unit %q", s, string(unit))
		}
		if seen[unit] {
			return 0, fmt.Errorf("%q: the unit %q is repeated", s, string(unit))
		}
		seen[unit] = true

		var n time.Duration
		for _, c := range []byte(digits) {
			n = n*10 + time.Duration(c-'0')
			if n > maxDuration/factor {
				return 0, fmt.Errorf("%q: the time is too long", s)
			}
		}
		total += n * factor
		if total > maxDuration {
			return 0, fmt.Errorf("%q: the time is too long", s)
		}
		rest = strings.TrimLeft(rest, " ")
	}

	if total <= 0 {
		return 0, fmt.Errorf("%q: the time must be positive", s)
	}

	return total, nil
}
