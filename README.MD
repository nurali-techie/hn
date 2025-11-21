# Overview
**hn** stands for `HackerNews`, it is a command line tool to query hacker news effectively as per user preference.

# Install

## Linux

```
wget -O hn https://github.com/nurali-techie/hn/releases/download/v0.2.0/hn-linux-amd64
chmod +x hn
mv hn ~/.local/bin/
```

Use `hn` from terminal.

## Mac

```
wget -O hn https://github.com/nurali-techie/hn/releases/download/v0.2.0/hn-darwin-amd64
chmod +x hn
```

Use `./hn` from terminal.

## Windows

```
wget -O hn.exe https://github.com/nurali-techie/hn/releases/download/v0.2.0/hn-windows-amd64.exe
```

Add folder path of `hn.exe` into your PATH variable under system environment variable; after that use `hn` from terminal.

# Usage

## Help
```
$ hn
Hello from HackerNews (link)

Usage-1: hn <days> <comma seperated search topics>
	hn 3 golang		// search golang topic for last 3 days (with 10+ points)
	hn 1 ai,llm		// search both ai and llm topics for last 1 days (with 10+ points)
	hn 2 "open source"	// use dobule-quotes for search topic having multiple words

Usage-2: hn <days> <comma seperated search topics> -a
	hn 3 golang -a		// search golang topic for last 3 days (with 1+ points)

Usage-3: hn <points>
	hn 200			// search any news from last 2 days with 200 to 300 points
	hn 500			// search any news from last 2 days with 500+ points
```

## Example

![hn example screenshot](assets/example.png)

## News Workflow

You could have below workflow for your daily news.

- run `hn` tool with topics of your interest.
- use (link) to open in brave browser.
- use Leo AI to summarize the page.

Check below video.

<video src="assets/hn_brave.mp4" width="960" height="720" controls></video>

# Contribution

Please help to test Mac and Windows installation.

# References
- HackerNews Search API - https://hn.algolia.com/api
