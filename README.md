# go-imap-appendlimit

[![GoDoc](https://godoc.org/github.com/emersion/go-imap-appendlimit?status.svg)](https://godoc.org/github.com/emersion/go-imap-appendlimit)

The [IMAP APPENDLIMIT Extension](https://tools.ietf.org/html/rfc7889) for [go-imap](https://github.com/emersion/go-imap).

## Usage

```go
s.Enable(appendlimit.NewExtension())
```

The backend must implement `appendlimit.Backend` and `appendlimit.User`.

## License

MIT
