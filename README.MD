# Overview

**hn** stands for `HackerNews`, it is a command line tool to query hacker news effectively as per user preference.

# Usage

```
Usage: hn <days> <comma seperated search terms>
Example:
	hn 5 golang
	hn 3 devops,java
	hn 2 "open source"
	hn 7 "cloud native,microservices"
```

# Install

```bash
curl -sSfL https://raw.githubusercontent.com/nurali-techie/hn/main/scripts/install.sh | sh
```

Note, the script requires root or sudo privileges to move 'hn' tool binary to `/usr/local/bin`.


# References

- HackerNews Search API - https://hn.algolia.com/api
- Install script - https://github.com/ddosify/ddosify
