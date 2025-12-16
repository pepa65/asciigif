# asciigif v0.3.0

**Ascii-gifs served for terminal consumption**

* After [`ascii.live`](https://github.com/hugomd/ascii.live)
* Latest version is running on [a-di.eu](https://a-di.eu/list)
* Run locally (by default on port `8080`), run: `asciigif`
* Run on a different port: `asciigif -port 8888`
* Run with a different default framerate than 70: `asciigif -framerate 200`
* List all available framesets: `curl localhost:8080/list`
* Try in a terminal: `curl localhost:8080/badapple`
* Clientside setting framerate: `curl localhost:8080/title?framerate=300`
* License: GPLv3

## Adding framesets
* [Fork this repo](https://github.com/pepa65/asciigif/fork)
* Clone your repo to work locally: `git clone https://github.com/YOURUSERNAME/asciigif`.
* Manually:
  - Copy the frameset template [`frames/.go`](.frames/.go): `cp frames/.go frames/MYFRAMESETNAME.go`.
  - Edit the new file with your frames.
* From a .gif file:
  - Run the `gif2go` file in this repo: `bash gif2go -n YOURNAME YOURFILE.gif`
* Commit and push your changes: `git commit -a && git push`
* If you want can make a **P**ull**R**equest on github.com/pepa65/asciigif
  ([check GitHub's help page on pull requests](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)).

