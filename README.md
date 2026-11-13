# Simple Environment Loader

Read typed values from environment variables with fallback defaults.

## Functions

### `GetEnvStr(key, def string) string`

Returns the value of the environment variable named by key,
or def if the variable is not set or is empty.

### `GetEnvInt(key string, def int) int`

Returns the value of the environment variable named by key parsed as int,
or def if the variable is not set, is empty, or cannot be parsed.

### `GetEnvBool(key string, def bool) bool`

Returns the value of the environment variable named by key parsed as bool,
or def if the variable is not set, is empty, or cannot be parsed.
Accepted values: `1`, `t`, `T`, `TRUE`, `true`, `True`, `0`, `f`, `F`, `FALSE`, `false`, `False`.

### `GetEnvSalt(key, def string) string`

Returns a 32-byte key derived from the environment variable named by key,
suitable for use as an AES-256 symmetric key or similar fixed-size secret.

The raw value is stretched or truncated to exactly 32 bytes:
- shorter than 32 bytes: repeated until at least 32 bytes, then truncated to 32.
- longer than 32 bytes: truncated to the first 32 bytes.
- exactly 32 bytes: returned unchanged.

If the variable is not set or is empty, def is used as the source material.
If both are empty, an empty string is returned.

> **Note:** operates on bytes, not Unicode code points. Use ASCII values for predictable results.

### `GetEnvUrl(key string, def string) string`

Returns the value of the environment variable named by key as a URL with a trailing slash.
The URL must have an `http`, `https`, or `wss` scheme and a non-empty hostname; otherwise def is used.
If the variable is not set or is empty, def is used. If def is also empty, an empty string is returned.

### `GetEnvFloat(key string, def float64) float64`

Returns the value of the environment variable named by key parsed as float64,
or def if the variable is not set, is empty, or cannot be parsed.

## Staying up to date

```bash
go get -u github.com/ra-company/env
```

## Supported Go versions

Go 1.26 and later.

## License

MIT — see [LICENSE](LICENSE).

## Dependencies

**Tests only**

| Library | Purpose | License |
|---------|---------|---------|
| [github.com/stretchr/testify](https://github.com/stretchr/testify) | Test assertions | MIT |
