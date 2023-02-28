// Package version 版本
package version

import (
	"bytes"
	"runtime"
	"text/template"
)

var (
	// Version is version
	Version = "v0.1.2"
)

// versionOptions include version
type versionOptions struct {
	Version   string
	GoVersion string
	Os        string
	Arch      string
}

var versionTemplate = ` Version:      {{.Version}}
 Go version:   {{.GoVersion}}
 OS/Arch:      {{.Os}}/{{.Arch}}
 `

// Get 获取版本信息
func Get() string {
	var doc bytes.Buffer
	vo := versionOptions{
		Version:   Version,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
	tmpl, _ := template.New("version").Parse(versionTemplate)
	_ = tmpl.Execute(&doc, vo)
	return doc.String()
}
