package types_test

import (
	"testing"

	"github.com/michimani/text-detector/types"
	"github.com/stretchr/testify/assert"
)

func Test_DetectedText_IsContainedIn(t *testing.T) {
	cases := []struct {
		name   string
		dt     *types.DetectedText
		target types.Position
		expect bool
	}{
		{
			name: "true",
			dt: &types.DetectedText{
				Text: "test text",
				Position: types.Position{
					LT: types.Coordinate{X: 10, Y: 20},
					RT: types.Coordinate{X: 30, Y: 20},
					RB: types.Coordinate{X: 30, Y: 5},
					LB: types.Coordinate{X: 10, Y: 5},
				},
			},
			target: types.Position{
				LT: types.Coordinate{X: 9, Y: 19},
				RT: types.Coordinate{X: 31, Y: 19},
				RB: types.Coordinate{X: 31, Y: 6},
				LB: types.Coordinate{X: 9, Y: 6},
			},
			expect: true,
		},
		{
			name: "false",
			dt: &types.DetectedText{
				Text: "test text",
				Position: types.Position{
					LT: types.Coordinate{X: 10, Y: 20},
					RT: types.Coordinate{X: 30, Y: 20},
					RB: types.Coordinate{X: 30, Y: 5},
					LB: types.Coordinate{X: 10, Y: 5},
				},
			},
			target: types.Position{
				LT: types.Coordinate{X: 9, Y: 19},
				RT: types.Coordinate{X: 31, Y: 19},
				RB: types.Coordinate{X: 31, Y: 4},
				LB: types.Coordinate{X: 9, Y: 4},
			},
			expect: false,
		},
		{
			name: "false: DetectedText is nil",
			dt:   nil,
			target: types.Position{
				LT: types.Coordinate{X: 9, Y: 19},
				RT: types.Coordinate{X: 31, Y: 19},
				RB: types.Coordinate{X: 31, Y: 6},
				LB: types.Coordinate{X: 9, Y: 6},
			},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			b := c.dt.IsContainedIn(c.target)
			assert.Equal(tt, c.expect, b)
		})
	}
}
