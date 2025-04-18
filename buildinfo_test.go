package gobuildinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	defaultInfo := &Info{
		Version:   "dev",
		Commit:    "none",
		Date:      "unknown",
		TreeState: "none",
		Project:   Project{},
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
				info.Version = "v1.0.0"
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
				info.Commit = "abc123"
				info.Date = "2023-01-01"
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
				info.Version = "v2.0.0"
				info.Commit = "def456"
				info.Date = "2023-02-01"
				info.TreeState = "clean"
				return &info
			}(),
		},
		{
			name: "Metadata with custom values",
			options: []Option{
				WithProject("CustomApp", "Custom Description", "https://customapp.com"),
				WithASCIILogo("Custom ASCII Art"),
			},
			expected: func() *Info {
				info := *defaultInfo
				info.Project = Project{
					Name:      "CustomApp",
					Desc:      "Custom Description",
					URL:       "https://customapp.com",
					ASCIILogo: "Custom ASCII Art",
				}
				return &info
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.options...)
			assert.NotNil(t, got)
			assert.EqualExportedValues(t, tt.expected, got)
		})
	}
}
