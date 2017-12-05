/*
  https://github.com/go-sql-driver/mysql/blob/master/connection.go
*/
package /* mysql */ main

// a copy of context.Context for Go 1.7 and earlier
type mysqlContext interface {
	Done() <-chan struct{}
	Err() error

	// defined in context.Context, but not used in this driver:
	// Deadline() (deadline time.Time, ok bool)
	// Value(key interface{}) interface{}
}

// Closes the network connection and unsets internal variables. Do not call this
// function after successfully authentication, call Close instead. This function
// is called before auth or on auth failure because MySQL will have already
// closed the network connection.
func (mc *mysqlConn) cleanup() {
	if !mc.closed.TrySet(true) {
		return
	}

	// Makes cleanup idempotent
	close(mc.closech)
	if mc.netConn == nil {
		return
	}
	if err := mc.netConn.Close(); err != nil {
		errLog.Print(err)
	}
}

/*
  https://github.com/go-sql-driver/mysql/blob/master/connection_go18.go
*/
func (mc *mysqlConn) startWatcher() {
	watcher := make(chan mysqlContext, 1)
	mc.watcher = watcher
	finished := make(chan struct{})
	mc.finished = finished
	go func() {
		for {
			var ctx mysqlContext
			select {
			case ctx = <-watcher:
			case <-mc.closech:
				return
			}

			select {
			case <-ctx.Done():
				mc.cancel(ctx.Err())
			case <-finished:
			case <-mc.closech:
				return
			}
		}
	}()
}

// finish is called when the query has canceled.
func (mc *mysqlConn) cancel(err error) {
	mc.canceled.Set(err)
	mc.cleanup()
}
