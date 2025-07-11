package gobuildinfo

import (
	"bytes"
	"cmp"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	defaultInfo := &Info{
		version:   "dev",
		commit:    "none",
		date:      "unknown",
		treeState: "none",
		project:   Project{},
	}

	tests := []struct {
		name     string
		options  []Option
		expected *Info
	}{
		{
			name:     "Default Info values",
			options:  nil,
			expected: defaultInfo,
		},
		{
			name: "Custom Info version",
			options: []Option{
				WithVersion("v1.0.0"),
			},
			expected: func() *Info {
				info := *defaultInfo
				info.version = "v1.0.0"
				return &info
			}(),
		},
		{
			name: "Custom Info commit and date",
			options: []Option{
				WithCommit("abc123"),
				WithDate("2023-01-01"),
			},
			expected: func() *Info {
				info := *defaultInfo
				info.commit = "abc123"
				info.date = "2023-01-01"
				return &info
			}(),
		},
		{
			name: "All custom Info values",
			options: []Option{
				WithVersion("v2.0.0"),
				WithCommit("def456"),
				WithDate("2023-02-01"),
				WithTreeState("clean"),
			},
			expected: func() *Info {
				info := *defaultInfo
				info.version = "v2.0.0"
				info.commit = "def456"
				info.date = "2023-02-01"
				info.treeState = "clean"
				return &info
			}(),
		},
		{
			name: "Metadata with custom values",
			options: []Option{
				WithProject(Project{
					Name:        "CustomApp",
					Description: "Custom Description",
					URL:         "https://customapp.com",
					ASCIILogo:   "Custom ASCII Art",
				}),
			},
			expected: func() *Info {
				info := *defaultInfo
				info.project = Project{
					Name:        "CustomApp",
					Description: "Custom Description",
					URL:         "https://customapp.com",
					ASCIILogo:   "Custom ASCII Art",
				}
				return &info
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.options...)
			if got == nil {
				t.Errorf("New() = nil, want %v", tt.expected)
				return
			}
			assertEqual(t, tt.expected.version, got.version)
			assertEqual(t, tt.expected.version, got.version)
			assertEqual(t, tt.expected.commit, got.commit)
			assertEqual(t, tt.expected.date, got.date)
			assertEqual(t, tt.expected.treeState, got.treeState)

			assertEqual(t, tt.expected.project.Name, got.project.Name)
			assertEqual(t, tt.expected.project.Description, got.project.Description)
			assertEqual(t, tt.expected.project.URL, got.project.URL)
			assertEqual(t, tt.expected.project.ASCIILogo, got.project.ASCIILogo)

			assertNotEmpty(t, got.runtime.Goos)
			assertNotEmpty(t, got.runtime.Goarch)
			assertNotEmpty(t, got.runtime.Compiler)
			assertNotEmpty(t, got.runtime.GoVersion)
			assertNotEmpty(t, got.runtime.ModuleSum)
		})
	}
}

func TestInfo_String_Runtime(t *testing.T) {
	info := New()
	info.runtime.GoVersion = "go1.24.2"
	expected := `Version:    dev
Commit:     none
Date:       unknown
TreeState:  none
GoVersion:  go1.24.2
Compiler:   gc
Platform:   linux/amd64
ModuleSum:  none
`
	var buf bytes.Buffer
	fmt.Fprint(&buf, info.String())
	assertEqual(t, expected, buf.String())
}

func assertNotEmpty[T cmp.Ordered](t *testing.T, value T) {
	t.Helper()
	var zero T
	if cmp.Compare(value, zero) == 0 {
		t.Errorf("expected non-empty value, got %v", value)
	}
}

func assertEqual[T cmp.Ordered](t *testing.T, expected, actual T) {
	t.Helper()
	if cmp.Compare(expected, actual) != 0 {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
