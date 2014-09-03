package shirokuro_test

import (
	"github.com/hajimehoshi/shirokuro"
	"os"
	"testing"
)

// 00100100
// 00111111
// 01101010
// 10001000
// 10000101
// 10100011
// 00001000
// 11010011
const bitBoard = shirokuro.BitBoard(0x243f6a8885a308d3)

func Example() {
	f := shirokuro.NewField()
	f.PrettyPrint(os.Stdout)
}

func TestInit(t *testing.T) {
	f := shirokuro.NewField()
	tests := []struct {
		want shirokuro.Stone
		i    uint8
		j    uint8
	}{
		{shirokuro.StoneNone, 0, 0},
		{shirokuro.StoneNone, 7, 7},
		{shirokuro.StoneWhite, 3, 3},
		{shirokuro.StoneBlack, 4, 3},
		{shirokuro.StoneBlack, 3, 4},
		{shirokuro.StoneWhite, 4, 4},
	}

	for _, test := range tests {
		actual := f.Stone(test.i, test.j)
		if actual != test.want {
			t.Errorf("(%d, %d) = %v; want %v", test.i, test.j, actual, test.want)
		}
	}
}

func TestVerticalLine(t *testing.T) {
	tests := []struct {
		want uint8
		i    uint8
	}{
		{0x1d, 0}, {0x21, 1}, {0xe4, 2},
	}
	for _, test := range tests {
		actual := bitBoard.VerticalLine(test.i)
		if actual != test.want {
			t.Errorf("bitBoard.VerticalLine(%d) = %d; want %d", test.i, actual, test.want)
		}
	}
}

func TestDiagonalDownLine(t *testing.T) {
	tests := []struct {
		want uint8
		i    uint8
		j    uint8
	}{
		{0x06, 5, 0}, {0x06, 6, 1}, {0x06, 7, 2},
		{0x06, 4, 0}, {0x06, 5, 1}, {0x06, 6, 2}, {0x06, 7, 3},
		{0x09, 3, 0}, {0x09, 4, 1}, {0x09, 5, 2}, {0x09, 6, 3}, {0x09, 7, 4},
		{0x39, 2, 0}, {0x39, 3, 1}, {0x39, 4, 2}, {0x39, 5, 3}, {0x39, 6, 4}, {0x39, 7, 5},
		{0x2e, 1, 0}, {0x2e, 2, 1}, {0x2e, 3, 2}, {0x2e, 4, 3}, {0x2e, 5, 4}, {0x2e, 6, 5}, {0x2e, 7, 6},
		{0x21, 0, 0}, {0x21, 1, 1}, {0x21, 2, 2}, {0x21, 3, 3}, {0x21, 4, 4}, {0x21, 5, 5}, {0x21, 6, 6}, {0x21, 7, 7},
		{0x42, 0, 1}, {0x42, 1, 2}, {0x42, 2, 3}, {0x42, 3, 4}, {0x42, 4, 5}, {0x42, 5, 6}, {0x42, 6, 7},
		{0x08, 0, 2}, {0x08, 1, 3}, {0x08, 2, 4}, {0x08, 3, 5}, {0x08, 4, 6}, {0x08, 5, 7},
		{0xa0, 0, 3}, {0xa0, 1, 4}, {0xa0, 2, 5}, {0xa0, 3, 6}, {0xa0, 4, 7},
		{0x90, 0, 4}, {0x90, 1, 5}, {0x90, 2, 6}, {0x90, 3, 7},
		{0x80, 0, 5}, {0x80, 1, 6}, {0x80, 2, 7},
	}

	for _, test := range tests {
		actual := bitBoard.DiagonalDownLine(test.i, test.j)
		if actual != test.want {
			t.Errorf("bitBoard.DiagonalDownLine(%d, %d) = %d; want %d", test.i, test.j, actual, test.want)
		}
	}
}

func TestMakeMove(t *testing.T) {
	const (
		n = shirokuro.StoneNone
		b = shirokuro.StoneBlack
		w = shirokuro.StoneWhite
	)
	tests := []struct {
		want  shirokuro.Field
		field shirokuro.Field
		i     uint8
		j     uint8
		stone shirokuro.Stone
	}{
		{
			*shirokuro.NewFieldFromStones([...][8]shirokuro.Stone{
				{n, n, n, b, n, n, n, n},
				{n, n, n, b, n, n, n, n},
				{n, n, n, b, n, n, n, n},
				{n, b, b, b, b, b, b, b},
				{n, n, n, b, n, n, n, n},
				{n, n, n, b, n, n, n, n},
				{n, n, n, b, n, n, n, n},
				{n, n, n, b, n, n, n, n},
			}),
			*shirokuro.NewFieldFromStones([...][8]shirokuro.Stone{
				{n, n, n, b, n, n, n, n},
				{n, n, n, w, n, n, n, n},
				{n, n, n, w, n, n, n, n},
				{n, b, w, n, w, w, w, b},
				{n, n, n, w, n, n, n, n},
				{n, n, n, w, n, n, n, n},
				{n, n, n, w, n, n, n, n},
				{n, n, n, b, n, n, n, n},
			}),
			3, 3, b,
		},
		{
			*shirokuro.NewFieldFromStones([...][8]shirokuro.Stone{
				{n, n, n, n, b, n, n, n},
				{n, n, n, n, b, n, n, n},
				{n, n, n, n, b, n, n, n},
				{n, b, b, b, b, b, b, b},
				{n, n, n, n, b, n, n, n},
				{n, n, n, n, b, n, n, n},
				{n, n, n, n, b, n, n, n},
				{n, n, n, n, b, n, n, n},
			}),
			*shirokuro.NewFieldFromStones([...][8]shirokuro.Stone{
				{n, n, n, n, b, n, n, n},
				{n, n, n, n, w, n, n, n},
				{n, n, n, n, w, n, n, n},
				{n, b, w, w, n, w, w, b},
				{n, n, n, n, w, n, n, n},
				{n, n, n, n, w, n, n, n},
				{n, n, n, n, w, n, n, n},
				{n, n, n, n, b, n, n, n},
			}),
			4, 3, b,
		},
	}

	for _, test := range tests {
		result := test.field.MakeMove(test.i, test.j, test.stone)
		if !result {
			t.Errorf("MakeMove(%d, %d, %d) failed", test.i, test.j, test.stone)
			continue
		}
		if test.field != test.want {
			t.Errorf("MakeMove(%d, %d, %d) = %+v; want %+v", test.i, test.j, test.stone, test.field, test.want)
		}
	}
}
