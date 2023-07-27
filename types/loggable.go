package types

import "strings"

///////////////////////////////////////////////////

type Loggable interface {
	LogStack() []string
	Log(s string)
	LogsToString() string
}
type LogData struct {
	logs []string
}

func (l *LogData) LogStack() []string {
	return l.logs
}
func (l *LogData) LogsToString() string {
	return strings.Join(l.logs, "\n")
}

func (r *LogData) Log(line string) {
	r.logs = append(r.logs, line)
}
