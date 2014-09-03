package shirokuro

type Player uint8

const (
	PlayerNone Player = iota
	PlayerPlayer
	PlayerOpponent
)

type Line struct {
	Player   uint8
	Opponent uint8
}

func (l *Line) Index() int {
	index := 0
	for i := 0; i < 8; i++ {
		index *= 3
		if l.Player&uint8(1<<uint(i)) != 0 {
			index += int(PlayerPlayer)
		} else if l.Opponent&uint8(1<<uint(i)) != 0 {
			index += int(PlayerOpponent)
		}
	}
	return index
}

var mobilityTable [6561][8]Line

func init() {
	for i := 0; i < len(mobilityTable); i++ {
		index := i
		pattern := [8]Player{}
		for col := 0; col < 8; col++ {
			pattern[col] = Player(index % 3)
			index /= 3
		}
		for col := 0; col < 8; col++ {
			if pattern[col] != PlayerNone {
				// Can't put the stone here
				continue
			}

			var after Line
			for col2 := 0; col2 < 8; col2++ {
				switch pattern[col2] {
				case PlayerPlayer:
					after.Player |= uint8(1 << uint(7-col2))
				case PlayerOpponent:
					after.Opponent |= uint8(1 << uint(7-col2))
				}
			}

		LeftSide:
			for num := 2; num <= 7; num++ {
				if col < num {
					break
				}
				if pattern[col-num] != PlayerPlayer {
					continue
				}
				for col2 := col - 1; col-num+1 <= col2; col2-- {
					if pattern[col2] != PlayerOpponent {
						continue LeftSide
					}
				}
				for col2 := col - 1; col-num+1 <= col2; col2-- {
					after.Player |= uint8(1 << uint(7-col2))
					after.Opponent &^= uint8(1 << uint(7-col2))
				}
			}

		RightSide:
			for num := 2; num <= 7; num++ {
				if 8 <= col+num {
					break
				}
				if pattern[col+num] != PlayerPlayer {
					continue
				}
				for col2 := col + 1; col2 <= col+num-1; col2++ {
					if pattern[col2] != PlayerOpponent {
						continue RightSide
					}
				}
				for col2 := col + 1; col2 <= col+num-1; col2++ {
					after.Player |= uint8(1 << uint(7-col2))
					after.Opponent &^= uint8(1 << uint(7-col2))
				}
			}

			mobilityTable[i][col] = after
		}
	}
}

func NextLine(before Line, nextMove uint8) Line {
	return mobilityTable[before.Index()][nextMove]
}
