### pbcp

#### Install

```
go get -u github.com/pbcp/pbcp/...
```

#### Usage

Use `pbc` to copy from standard input, `pbp` to paste to standard output. `pbp` also takes in an optional argument, the index of the clip to paste (for example, `pbp 2` pastes the second most recent clip).

#### Configuration

Your unique user ID is stored in `$HOME/.config/pbcp/id`. Make sure this ID is the same across all your machines.
