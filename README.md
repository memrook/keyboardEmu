# Emulating keyboard events

> in this version the application simply presses F5 every 5 seconds

On Windows, you should build like this:

```
env GO111MODULE=on go build -ldflags "-H=windowsgui"
```
