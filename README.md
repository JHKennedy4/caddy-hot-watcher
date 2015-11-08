## Caddy Hot Watcher

A simple filesystem watcher for use with [caddy-hot-loader](https://github.com/JHKennedy4/caddy-hot-loader)

### Installation
If $GOPATH/bin is in your $PATH:
`go get github.com/jhkennedy4/caddy-hot-watcher`

Or:
Download the appropriate build, rename it, and add it to your path

### Usage
Add the following line to your Caddyfile:
```
websocket /hot-watcher caddy-hot-watcher
```
