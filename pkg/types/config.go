package types

import "time"

// Config holds the global configuration
type Config struct {
	DryRun      bool
	Verbose     bool
	Silent      bool
	Force       bool
	MaxAge      time.Duration
	Paths       []string
	SkipTypes   []string
	Concurrency int
}

// CleanOptions holds cleaning specific options
type CleanOptions struct {
	MaxAge    time.Duration
	SkipTypes []string
}

// ScanOptions holds scanning specific options
type ScanOptions struct {
	Paths []string
	Depth int
}

// MonitorOptions holds monitoring specific options
type MonitorOptions struct {
	Interval  time.Duration
	Threshold int64
}
