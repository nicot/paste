A program to host small texts. Useful for Unixy sharing of scripts.

To build/run:
- `go install github.com/nicot/paste/pasteServer`
- `pasteServer`

Usage:
- `curl localhost:8080 -d 'my data'`
- `cat my_file | curl localhost:8080 -d @-`
