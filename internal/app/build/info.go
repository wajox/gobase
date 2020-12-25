package build

import (
	"runtime"
)

// Info represents all available build & runtime environment information.
type Info struct {
	Version    string `json:"version,omitempty"`
	CommitHash string `json:"commit_hash,omitempty"`
	BuildDate  string `json:"build_date,omitempty"`
	GoVersion  string `json:"go_version,omitempty"`
	Os         string `json:"os,omitempty"`
	Arch       string `json:"arch,omitempty"`
	Compiler   string `json:"compiler,omitempty"`
}

// NewInfo returns all available build information.
func NewInfo() *Info {
	return &Info{
		Version:    Version,
		CommitHash: CommitHash,
		BuildDate:  BuildDate,
		GoVersion:  runtime.Version(),
		Os:         runtime.GOOS,
		Arch:       runtime.GOARCH,
		Compiler:   runtime.Compiler,
	}
}
