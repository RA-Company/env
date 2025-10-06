# Simple Environment loader module

Load different values from environment settings.

# Functions list

```func GetEnvStr(key string, byDefault string) string```

<details>
<summary>Function description</summary>

GetEnvStr retrieves an environment variable as a string, returning a default value if not set.

Parameters:
  - key: The name of the environment variable to retrieve.
  - byDefault: The default value to return if the environment variable is not set or is empty.
</details>

```func GetEnvInt(key string, byDefault int) int```

<details>
<summary>Function description</summary>

GetEnvInt retrieves an environment variable as an int, returning a default value if not set.

Parameters:
  - key: The name of the environment variable to retrieve.
  - byDefault: The default value to return if the environment variable is not set or is empty.
</details>

```func GetEnvBool(key string, byDefault bool) bool```

<details>
<summary>Function description</summary>

GetEnvBool retrieves an environment variable as a boolean, returning a default value if not set or if parsing fails.

Parameters:
  - key: The name of the environment variable to retrieve.
  - byDefault: The default value to return if the environment variable is not set or if parsing fails.
</details>

```func GetEnvSalt(key string, byDefault string) string```

<details>
<summary>Function description</summary>

GetEnvSalt retrieves an environment variable as a salt string, ensuring it is exactly 32 characters long.
If the environment variable is shorter than 32 characters, it repeats the value until it reaches 32 characters.
If the environment variable is longer than 32 characters, it truncates the value to 32 characters.
If the environment variable is not set, it returns a default salt value, which is also adjusted to be 32 characters long.

Parameters:
  - key: The name of the environment variable to retrieve.
  - byDefault: The default value to use.
</details>

```func GetEnvUrl(key string, byDefault string) string```

<details>
<summary>Function description</summary>

GetEnvUrl retrieves an environment variable as a URL, ensuring it ends with a slash.
If the environment variable is not set, it returns a default URL with a trailing slash.
If the environment variable is set but does not contain a valid URL, it returns the default URL.

Parameters:
  - key: The name of the environment variable to retrieve.
  - byDefault: The default URL to return if the environment variable is not set or is invalid.
</details>

```func GetEnvFloat(key string, byDefault float64) float64```

<details>
<summary>Function description</summary>

GetEnvFloat retrieves an environment variable as a float64, returning a default value if not set or if parsing fails.

Parameters:
  - key: The name of the environment variable to retrieve.
  - byDefault: The default value to return if the environment variable is not set or if parsing fails.
</details>

## Used libraries
* https://github.com/stretchr/testify - Module for tests (MIT License)

# Staying up to date
To update library to the latest version, use go get -u github.com/ra-company/env.

# Supported go versions
We currently support the most recent major Go versions from 1.25.0 onward.

# License
This project is licensed under the terms of the GPL-3.0 license.