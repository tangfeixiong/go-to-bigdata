Testing went into failed as _wal_ file system is _vboxsf_, but succeeded after changing _raft.go_
```
		waldir:      fmt.Sprintf("/tmp/raftexample-%d", id),
		snapdir:     fmt.Sprintf("/tmp/raftexample-%d-snap", id),
```
