/*
------------------------------------------------------------------------------------------------------------------------
####### aff5635 - License MIT
------------------------------------------------------------------------------------------------------------------------
*/

package sdk

import "log"

type (
	Logger interface {
		New(id, name string) Logger
		Trace(msg string, kv ...any)
		Debug(msg string, kv ...any)
		Info(msg string, kv ...any)
		Notice(msg string, kv ...any)
		Warning(err error, msg string, kv ...any)
		Error(err error, msg string, kv ...any)
		NewStdLogger(level, prefix string, flag int) *log.Logger
	}
)

/*
########################################################################################################################
*/
