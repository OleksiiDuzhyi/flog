// +build !windows

package main

import (
	"errors"
	"os"
	"path/filepath"
	"syscall"

	"github.com/smira/go-statsd"
)

// Run checks overwrite flag and generates logs with given options
func Run(option *Option) error {
	client := statsd.NewClient(option.StatsdHost + ":" + option.StatsdPort,
		statsd.MaxPacketSize(1400),
		statsd.MetricPrefix(option.StatsdPrefix))

	logDir := filepath.Dir(option.Output)
	oldMask := syscall.Umask(0000)
	if err := os.MkdirAll(logDir, 0766); err != nil {
		return err
	}
	syscall.Umask(oldMask)
	if _, err := os.Stat(option.Output); err == nil && !option.Overwrite {
		return errors.New(option.Output + " already exists. You can overwrite with -w option")
	}

	generated := Generate(option, client)
	return generated
}
