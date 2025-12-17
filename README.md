# asciigif v0.4.4

**Ascii-gifs served for terminal consumption**

* License: GPLv3
* Repo: https://github.com/pepa65/asciigif
* After [`ascii.live`](https://github.com/hugomd/ascii.live)
* Latest version is running on [a-di.eu](https://a-di.eu/list)
* Run locally (by default on port `8080`), run: `asciigif`
* Run on a different port: `asciigif --port 8888`
* Run with a different default framerate than 70: `asciigif --framerate 200`
* Show version: `asciigif --version`
* List all available framesets: `curl localhost:8080/list`
* Try in a terminal: `curl localhost:8080/badapple`
* Clientside setting framerate: `curl localhost:8080/title?framerate=300`

## Install
### Download standalone binary
```
wget -q 4e4.in/asciigif
mv asciigif ~/bin/  # Assuming ~/bin is in $PATH
```

### Build locally (needs Golang install)
```
git clone https://github.com/pepa65/asciigif
cd asciigif
CGO_ENABLED=0 go install -ldflags="-s -w"  # Flags for smaller binary
upx --best --lzma asciigif  # Compress the binary for smaller size
mv asciigif ~/bin/  # Assuming ~/bin is in PATH

# Build for various architectures:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o asciigif
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o asciigif_pi
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o asciigif_bsd
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o asciigif_osx
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o asciigif.exe
```

## Adding framesets
* [Fork this repo](https://github.com/pepa65/asciigif/fork)
* Clone your repo to work locally: `git clone https://github.com/YOURUSERNAME/asciigif`.
* Manually:
  - Copy the frameset template [`frames/.go`](.frames/.go): `cp frames/.go frames/NAME.go`.
  - Edit the new file with your frames.
* From a .gif file:
  - Run the `gif2go` file in this repo: `bash gif2go -n NAME YOURFILE.gif`
* Commit and push your changes: `git commit -a && git push`
* If you want can make a **P**ull**R**equest on github.com/pepa65/asciigif
  ([check GitHub's help page on pull requests](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)).

