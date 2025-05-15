// Package buildinfo provides information about the build.
//
// This implementation was inspired by the [go-version] module.
//
// [go-version]: https://github.com/caarlos0/go-version
package gobuildinfo

import (
	"cmp"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"text/tabwriter"
)

// Project implements the [Option] interface to set the project information.
type Project struct {
	Name        string
	Description string
	URL         string
	ASCIILogo   string
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
	version   string
	commit    string
	date      string
	treeState string
	project   Project

	disableRuntime bool
	runtime        *runtimeEnv
}

type Options Info

type Option interface {
	apply(*Options)
}

type optionFunc func(*Options)

func (f optionFunc) apply(opts *Options) {
	f(opts)
}

func WithVersion(version string) Option {
	return optionFunc(func(opts *Options) {
		opts.version = version
	})
}

func WithCommit(commit string) Option {
	return optionFunc(func(opts *Options) {
		opts.commit = commit
	})
}

func WithDate(date string) Option {
	return optionFunc(func(opts *Options) {
		opts.date = date
	})
}

func WithTreeState(treeState string) Option {
	return optionFunc(func(opts *Options) {
		opts.treeState = treeState
	})
}

func WithProject(project Project) Option {
	return optionFunc(func(opts *Options) {
		opts.project = project
	})
}

func WithDisableRuntime() Option {
	return optionFunc(func(opts *Options) {
		opts.disableRuntime = true
	})
}

// New creates a new Info instance with the provided options.
func New(opts ...Option) *Info {
	// options := &Options{Info: new(Info)}
	options := new(Options)
	for _, opt := range opts {
		opt.apply(options)
	}

	info := &Info{
		version:   cmp.Or(options.version, "dev"),
		commit:    cmp.Or(options.commit, "none"),
		date:      cmp.Or(options.date, "unknown"),
		treeState: cmp.Or(options.treeState, "none"),
		project:   options.project,
	}
	if !options.disableRuntime {
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

func (i Info) String() string {
	var sb strings.Builder
	if i.project.Name != "" {
		if i.project.ASCIILogo != "" {
			sb.WriteString(i.project.ASCIILogo)
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("%s: %s\n", i.project.Name, i.project.Description))
		sb.WriteString(fmt.Sprintf("%s\n", i.project.URL))
		sb.WriteString("\n")
	}

	w := tabwriter.NewWriter(&sb, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintf(w, "Version:\t%s\n", i.version)
	_, _ = fmt.Fprintf(w, "Commit:\t%s\n", i.commit)
	_, _ = fmt.Fprintf(w, "Date:\t%s\n", i.date)
	_, _ = fmt.Fprintf(w, "TreeState:\t%s\n", i.treeState)

	if i.runtime != nil {
		_, _ = fmt.Fprintf(w, "GoVersion:\t%s\n", i.runtime.GoVersion)
		_, _ = fmt.Fprintf(w, "Compiler:\t%s\n", i.runtime.Compiler)
		_, _ = fmt.Fprintf(w, "Platform:\t%s\n", i.runtime.Platform())
		_, _ = fmt.Fprintf(w, "ModuleSum:\t%s\n", i.runtime.ModuleSum)
	}
	_ = w.Flush()

	return sb.String()
}
