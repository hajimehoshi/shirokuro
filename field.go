package shirokuro

import (
	"fmt"
	"io"
)

type Stone int

const (
	StoneNone Stone = iota
	StoneBlack
	StoneWhite
)

type BitBoard uint64

func (b BitBoard) HorizontalLine(row uint8) uint8 {
	return uint8((b >> ((7 - row) * 8)))
}

func (b *BitBoard) SetHorizontalLine(line, row uint8) {
	var masks = [...]BitBoard{
		0x00ffffffffffffff,
		0xff00ffffffffffff,
		0xffff00ffffffffff,
		0xffffff00ffffffff,
		0xffffffff00ffffff,
		0xffffffffff00ffff,
		0xffffffffffff00ff,
		0xffffffffffffff00,
	}
	*b &= masks[row]
	*b |= BitBoard(uint64(line) << ((7 - row) * 8))
}

func (b BitBoard) VerticalLine(col uint8) uint8 {
	b <<= col
	b &= 0x8080808080808080
	b *= 0x0002040810204081
	return uint8(b >> 56)
}

func (b *BitBoard) SetVerticalLine(line, col uint8) {
	var masks = [...]BitBoard{
		0x7f7f7f7f7f7f7f7f,
		0xbfbfbfbfbfbfbfbf,
		0xdfdfdfdfdfdfdfdf,
		0xefefefefefefefef,
		0xf7f7f7f7f7f7f7f7,
		0xfbfbfbfbfbfbfbfb,
		0xfdfdfdfdfdfdfdfd,
		0xfefefefefefefefe,
	}

	l := BitBoard(line)
	l *= 0x0101010101010101
	l &= 0x8040201008040201
	l *= 0x0101010101010101
	l &= 0x0101010101010101
	l <<= 7 - col
	*b &= masks[col]
	*b |= l
}

func (b BitBoard) DiagonalDownLine(i, j uint8) uint8 {
	var masks = map[int]BitBoard{
		-5: 0x0000000000804020,
		-4: 0x0000000080402010,
		-3: 0x0000008040201008,
		-2: 0x0000804020100804,
		-1: 0x0080402010080402,
		0:  0x8040201008040201,
		1:  0x4020100804020100,
		2:  0x2010080402010000,
		3:  0x1008040201000000,
		4:  0x0804020100000000,
		5:  0x0402010000000000,
	}

	b &= masks[int(i)-int(j)]
	b *= 0x0101010101010101
	b >>= 56
	return uint8(b)
}

func (b BitBoard) SetDiagonalDownLine(line, i, j uint8) {
	// TODO
}

type Field struct {
	Black BitBoard
	White BitBoard
}

type Diagonal int

const (
	DiagonalUp Diagonal = iota
	DiagonalDown
)

type Hash [2]uint64

func NewField() *Field {
	return &Field{
		0x0000000810000000,
		0x0000001008000000,
	}
}

func NewFieldFromStones(stones [8][8]Stone) *Field {
	black := BitBoard(0)
	white := BitBoard(0)
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			black <<= 1
			white <<= 1
			switch stones[j][i] {
			case StoneBlack:
				black |= 1
			case StoneWhite:
				white |= 1
			}
		}
	}
	return &Field{
		Black: black,
		White: white,
	}
}

func (f *Field) Stone(i, j uint8) Stone {
	blackRow := uint8((f.Black >> ((7 - j) * 8)) & 0xff)
	whiteRow := uint8((f.White >> ((7 - j) * 8)) & 0xff)
	if (blackRow >> (7 - i) & 0x1) != 0 {
		return StoneBlack
	} else if (whiteRow >> (7 - i) & 0x1) != 0 {
		return StoneWhite
	}
	return StoneNone
}

func (f *Field) PrettyPrint(w io.Writer) {
	fmt.Fprint(w, "  a b c d e f g h\n")
	for j := uint8(0); j < 8; j++ {
		fmt.Fprintf(w, "%d", j+1)
		for i := uint8(0); i < 8; i++ {
			switch f.Stone(i, j) {
			case StoneBlack:
				io.WriteString(w, " X")
			case StoneWhite:
				io.WriteString(w, " O")
			case StoneNone:
				io.WriteString(w, " .")
			}
		}
		io.WriteString(w, "\n")
	}
}

func (f *Field) Hash() Hash {
	// TODO: normalize
	return [2]uint64{uint64(f.Black), uint64(f.White)}
}

func (f *Field) MakeMove(i, j uint8, stone Stone) bool {
	pl := (*BitBoard)(nil)
	op := (*BitBoard)(nil)
	switch stone {
	case StoneBlack:
		pl = &f.Black
		op = &f.White
	case StoneWhite:
		pl = &f.White
		op = &f.Black
	}

	var emptyLine Line

	// Horizontal
	hoLine := Line{
		Player:   pl.HorizontalLine(j),
		Opponent: op.HorizontalLine(j),
	}
	hoLineAfter := NextLine(hoLine, i)
	if hoLineAfter == emptyLine {
		return false
	}

	// Vertical
	veLine := Line{
		Player:   pl.VerticalLine(i),
		Opponent: op.VerticalLine(i),
	}
	veLineAfter := NextLine(veLine, j)
	if veLineAfter == emptyLine {
		return false
	}

	// TODO: 斜め

	pl.SetHorizontalLine(hoLineAfter.Player, j)
	op.SetHorizontalLine(hoLineAfter.Opponent, j)

	pl.SetVerticalLine(veLineAfter.Player, i)
	op.SetVerticalLine(veLineAfter.Opponent, i)

	b := BitBoard(1 << ((7-j)*8 + (7 - i)))
	*pl |= b

	return true
}







