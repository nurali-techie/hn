# Overview
**hn** stands for `HackerNews`, it is a command line tool to query hacker news effectively as per user preference.

# Install
Step1 - git clone
```
git clone https://github.com/nurali-techie/hn.git
```

Step2 - install using go
```
cd hn
go install
```

# Usage
```
Usage: hn <days> <comma seperated search terms>
Example:
	hn 5 golang
	hn 3 devops,java
	hn 2 "open source"
	hn 7 "cloud native,microservices"
```

Note, the script requires root or sudo privileges to move 'hn' tool binary to `/usr/local/bin`.


# References
- HackerNews Search API - https://hn.algolia.com/api
- Install script - https://github.com/ddosify/ddosify
