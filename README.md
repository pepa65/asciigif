# asciigif v0.2.0

**Ascii-gifs served for terminal consumption**

* After [`ascii.live`](https://github.com/hugomd/ascii.live)
* Run locally (by default on port `8080`), run: `asciigif`
* Run on a different port: `asciigif -port 8888`
* Run with a different default framerate than 70: `asciigif -framerate 200`
* List all available framesets: `curl localhost:8080/list`
* Try in a terminal: `curl localhost:8080/badapple`
* Clientside setting framerate: `curl localhost:8080/title?framerate=300`
* License: GPLv3

## Adding frames
* [Fork this repo](https://github.com/pepa65/asciigif/fork)
* Clone your repo to work locally: `git clone https://github.com/USERNAME/REPONAME`
* Copy the frame template [`frames/.go`](.frames/.go): `cp frames/.go frames/MYFRAMESETNAME.go`
* Edit the new file:
```Golang
package frames

var MyFrameset = DefaultFrameType(_MyFrameset)

var _MyFrameset = []string{
	`Frame1 Line1
Frame1 Line2
Frame1 Line3`,

	`Frame2 Line1
Frame2 Line2
Frame2 Line3`,

	`Frame3 Line1
Frame3 Line2
Frame3 Line2`,
}
```
* At the bottom of [`frames/frames.go`](frames/frames.go), add your `	"label": "MyFrameset",` in `FrameMap`.
* Commit and push your changes: `git commit -a && git push`
* Make a **P**ull**R**equest on github.com ([check GitHub's help page on the topic](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)).

