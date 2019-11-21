package version

import (
	"fmt"
	"runtime"
)

var (
	GitTag       string
	GitCommit    string
	GitTreeState string
	BuildTime    string
	Os           string
	Arch         string
)

// Info contains version information.
type Info struct {
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildTime    string `json:"buildTime"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	return info.GitTag
}

func Get() Info {
	return Info{
		GitTag:       GitTag,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildTime:    BuildTime,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", Os, Arch),
	}
}
