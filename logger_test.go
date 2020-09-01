package logger

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ignitionrobotics/web/ign-go"
	"testing"
)

type writer struct {
	Data []byte
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.Data = append(w.Data, p...)
	return len(p), nil
}

func newWriter() *writer {
	return &writer{}
}

func TestLogger_Debug(t *testing.T) {
	var w *writer
	w = newWriter()
	l := NewLogger(ign.VerbosityDebug, w)
	l.Debug(fmt.Sprintf("Value: %d", 1))
	assert.Contains(t, string(w.Data), "Value: 1")
}

func TestLogger_Info(t *testing.T) {
	var w *writer
	w = newWriter()
	l := NewLogger(ign.VerbosityDebug, w)
	l.Info(fmt.Sprintf("Value: %d", 1))
	assert.Contains(t, string(w.Data), "Value: 1")
}

func TestLogger_DebugLowerVerbosity(t *testing.T) {
	var w *writer
	w = newWriter()
	l := NewLogger(ign.VerbosityInfo, w)
	l.Debug(fmt.Sprintf("Value: %d", 1))
	assert.NotContains(t, string(w.Data), "Value: 1")
}

func TestLogger_InfoLowerVerbosity(t *testing.T) {
	var w *writer
	w = newWriter()
	l := NewLogger(ign.VerbosityWarning, w)
	l.Info(fmt.Sprintf("Value: %d", 1))
	assert.NotContains(t, string(w.Data), "Value: 1")
}
