package shirokuro_test

import (
	"github.com/hajimehoshi/shirokuro"
	"testing"
)

func bitsToUint8(bits ...uint8) uint8 {
	u := uint8(0)
	for i := 0; i < len(bits); i++ {
		u <<= 1
		u += bits[i]
	}
	return u
}

func TestNextLine(t *testing.T) {
	tests := []struct {
		want     shirokuro.Line
		before   shirokuro.Line
		nextMove uint8
	}{
		{
			shirokuro.Line{
				Player:   0,
				Opponent: 0,
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 0, 0, 0, 0, 0, 0, 0),
				Opponent: bitsToUint8(0, 0, 0, 0, 0, 0, 0, 0),
			},
			0,
		},
		{
			shirokuro.Line{
				Player:   bitsToUint8(0, 1, 1, 0, 0, 0, 0, 1),
				Opponent: bitsToUint8(0, 0, 0, 1, 1, 0, 1, 0),
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 0, 1, 0, 0, 0, 0, 1),
				Opponent: bitsToUint8(0, 1, 0, 1, 1, 0, 1, 0),
			},
			0,
		},
		{
			shirokuro.Line{
				Player:   0,
				Opponent: 0,
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 0, 1, 0, 0, 0, 0, 1),
				Opponent: bitsToUint8(0, 1, 0, 1, 1, 0, 1, 0),
			},
			1,
		},
		{
			shirokuro.Line{
				Player:   0,
				Opponent: 0,
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 0, 1, 0, 0, 0, 0, 1),
				Opponent: bitsToUint8(0, 1, 0, 1, 1, 0, 1, 0),
			},
			2,
		},
		{
			shirokuro.Line{
				Player:   bitsToUint8(0, 0, 1, 1, 1, 0, 1, 1),
				Opponent: bitsToUint8(0, 1, 0, 0, 0, 0, 0, 0),
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 0, 1, 0, 0, 0, 0, 1),
				Opponent: bitsToUint8(0, 1, 0, 1, 1, 0, 1, 0),
			},
			5,
		},
		{
			shirokuro.Line{
				Player:   0,
				Opponent: 0,
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 0, 1, 0, 0, 0, 0, 1),
				Opponent: bitsToUint8(0, 1, 0, 1, 1, 0, 1, 0),
			},
			7,
		},
		{
			shirokuro.Line{
				Player:   bitsToUint8(0, 1, 1, 0, 1, 0, 0, 1),
				Opponent: bitsToUint8(0, 0, 0, 1, 0, 0, 1, 0),
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 1, 1, 0, 1, 0, 0, 1),
				Opponent: bitsToUint8(0, 0, 0, 1, 0, 0, 1, 0),
			},
			0,
		},
		{
			shirokuro.Line{
				Player:   bitsToUint8(0, 1, 1, 0, 1, 0, 1, 1),
				Opponent: bitsToUint8(0, 0, 0, 1, 0, 0, 0, 0),
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 1, 1, 0, 1, 0, 0, 1),
				Opponent: bitsToUint8(0, 0, 0, 1, 0, 0, 1, 0),
			},
			5,
		},
		{
			shirokuro.Line{
				Player:   bitsToUint8(0, 1, 1, 0, 1, 1, 1, 1),
				Opponent: bitsToUint8(0, 0, 0, 0, 0, 0, 0, 0),
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 1, 0, 0, 0, 0, 0, 1),
				Opponent: bitsToUint8(0, 0, 1, 0, 1, 1, 1, 0),
			},
			3,
		},
		{
			shirokuro.Line{
				Player:   bitsToUint8(0, 1, 1, 1, 1, 1, 1, 1),
				Opponent: bitsToUint8(0, 0, 0, 0, 0, 0, 0, 0),
			},
			shirokuro.Line{
				Player:   bitsToUint8(0, 0, 0, 0, 0, 0, 0, 1),
				Opponent: bitsToUint8(0, 1, 1, 1, 1, 1, 1, 0),
			},
			0,
		},
		{
			shirokuro.Line{
				Player:   bitsToUint8(1, 1, 1, 1, 0, 1, 1, 1),
				Opponent: bitsToUint8(0, 0, 0, 0, 0, 0, 0, 0),
			},
			shirokuro.Line{
				Player:   bitsToUint8(1, 0, 0, 0, 0, 0, 0, 1),
				Opponent: bitsToUint8(0, 1, 1, 1, 0, 1, 1, 0),
			},
			4,
		},
		{
			shirokuro.Line{
				Player:   bitsToUint8(1, 1, 1, 1, 1, 1, 1, 0),
				Opponent: bitsToUint8(0, 0, 0, 0, 0, 0, 0, 0),
			},
			shirokuro.Line{
				Player:   bitsToUint8(1, 0, 0, 0, 0, 0, 0, 0),
				Opponent: bitsToUint8(0, 1, 1, 1, 1, 1, 1, 0),
			},
			7,
		},
	}
	for _, test := range tests {
		actual := shirokuro.NextLine(test.before, test.nextMove)
		if actual != test.want {
			t.Errorf("NextLine(%+v, %d) = %+v; want %+v", test.before, test.nextMove, actual, test.want)
		}
	}
}
