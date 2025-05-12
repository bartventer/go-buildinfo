// Package buildinfo provides information about the build.
//
// This implementation was inspired by the [go-version] module.
//
// [go-version]: https://github.com/caarlos0/go-version
package gobuildinfo

import (
	"cmp"
	_ "embed"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"text/tabwriter"
)

type Project struct {
	Name      string
	Desc      string
	URL       string
	ASCIILogo string
}

type runtimeEnv struct {
	Goos      string
	Goarch    string
	GoVersion string
	Compiler  string
	ModuleSum string
}

func (r *runtimeEnv) Platform() string {
	return fmt.Sprintf("%s/%s", r.Goos, r.Goarch)
}

type Info struct {
	Version   string
	Commit    string
	Date      string
	TreeState string
	Project   Project
	runtime   *runtimeEnv
}

type Options struct {
	*Info
	DisableRuntime bool
}

type Option func(*Options)

func WithVersion(version string) Option {
	return func(opts *Options) {
		opts.Version = version
	}
}

func WithCommit(commit string) Option {
	return func(opts *Options) {
		opts.Commit = commit
	}
}

func WithDate(date string) Option {
	return func(opts *Options) {
		opts.Date = date
	}
}

func WithTreeState(treeState string) Option {
	return func(opts *Options) {
		opts.TreeState = treeState
	}
}

func WithProject(name, desc, url string) Option {
	return func(opts *Options) {
		opts.Project.Name = name
		opts.Project.Desc = desc
		opts.Project.URL = url
	}
}

func WithASCIILogo(logo string) Option {
	return func(opts *Options) {
		opts.Project.ASCIILogo = logo
	}
}

func WithDisableRuntime() Option {
	return func(opts *Options) {
		opts.DisableRuntime = true
	}
}

// New creates a new Info instance with the provided options.
func New(opts ...Option) *Info {
	options := &Options{Info: new(Info)}
	for _, opt := range opts {
		opt(options)
	}

	info := &Info{
		Version:   cmp.Or(options.Version, "dev"),
		Commit:    cmp.Or(options.Commit, "none"),
		Date:      cmp.Or(options.Date, "unknown"),
		TreeState: cmp.Or(options.TreeState, "none"),
		Project:   options.Project,
	}
	if !options.DisableRuntime {
		bi, _ := debug.ReadBuildInfo()
		info.runtime = &runtimeEnv{
			Goos:      runtime.GOOS,
			Goarch:    runtime.GOARCH,
			Compiler:  runtime.Compiler,
			GoVersion: bi.GoVersion,
			ModuleSum: cmp.Or(bi.Main.Sum, "none"),
		}
	}

	return info
}

func (i *Info) String() string {
	var sb strings.Builder
	if i.Project.Name != "" {
		if i.Project.ASCIILogo != "" {
			sb.WriteString(i.Project.ASCIILogo)
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("%s: %s\n", i.Project.Name, i.Project.Desc))
		sb.WriteString(fmt.Sprintf("%s\n", i.Project.URL))
		sb.WriteString("\n")
	}

	w := tabwriter.NewWriter(&sb, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintf(w, "Version\t%s\n", i.Version)
	_, _ = fmt.Fprintf(w, "Commit\t%s\n", i.Commit)
	_, _ = fmt.Fprintf(w, "Date\t%s\n", i.Date)
	_, _ = fmt.Fprintf(w, "TreeState\t%s\n", i.TreeState)

	if i.runtime != nil {
		_, _ = fmt.Fprintf(w, "GoVersion\t%s\n", i.runtime.GoVersion)
		_, _ = fmt.Fprintf(w, "Compiler\t%s\n", i.runtime.Compiler)
		_, _ = fmt.Fprintf(w, "Platform\t%s\n", i.runtime.Platform())
		_, _ = fmt.Fprintf(w, "ModuleSum\t%s\n", i.runtime.ModuleSum)
	}
	_ = w.Flush()

	return sb.String()
}
