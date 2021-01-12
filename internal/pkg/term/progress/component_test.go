// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package progress

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleLineComponent_Render(t *testing.T) {
	testCases := map[string]struct {
		inText    string
		inPadding int

		wantedOut string
	}{
		"should print padded text with new line": {
			inText:    "hello world",
			inPadding: 4,

			wantedOut: "    hello world\n",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// GIVEN
			comp := &singleLineComponent{
				Text:    tc.inText,
				Padding: tc.inPadding,
			}
			buf := new(strings.Builder)

			// WHEN
			nl, err := comp.Render(buf)

			// THEN
			require.NoError(t, err)
			require.Equal(t, 1, nl, "expected only a single line to be written by a single line component")
			require.Equal(t, tc.wantedOut, buf.String())
		})
	}
}

func TestTreeComponent_Render(t *testing.T) {
	testCases := map[string]struct {
		inNode     Renderer
		inChildren []Renderer

		wantedNumLines int
		wantedOut      string
	}{
		"should render all the nodes": {
			inNode: &singleLineComponent{
				Text: "is",
			},
			inChildren: []Renderer{
				&singleLineComponent{
					Text: "this",
				},
				&singleLineComponent{
					Text: "working?",
				},
			},

			wantedNumLines: 3,
			wantedOut: `is
this
working?
`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// GIVEN
			comp := &treeComponent{
				Root:     tc.inNode,
				Children: tc.inChildren,
			}
			buf := new(strings.Builder)

			// WHEN
			nl, err := comp.Render(buf)

			// THEN
			require.NoError(t, err)
			require.Equal(t, tc.wantedNumLines, nl)
			require.Equal(t, tc.wantedOut, buf.String())
		})
	}
}
