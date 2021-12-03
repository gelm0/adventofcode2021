package dive

type Direction string

const (
	Forward Direction = "forward"
	Up      Direction = "up"
	Down    Direction = "down"
)

type Position struct {
	Direction Direction
	Distance  int
}

func Dive(positions []Position) int {
	depth := 0
	horizontal := 0
	for _, pos := range positions {
		switch pos.Direction {
		case Forward:
			horizontal += pos.Distance
		case Up:
			depth -= pos.Distance
		case Down:
			depth += pos.Distance

		}
	}
	return depth * horizontal
}

func DiveAim(positions []Position) int {
	depth := 0
	horizontal := 0
	aim := 0
	for _, pos := range positions {
		switch pos.Direction {
		case Forward:
			horizontal += pos.Distance
			if aim != 0 {
				depth += pos.Distance * aim
			}
		case Up:
			aim -= pos.Distance
		case Down:
			aim += pos.Distance

		}
	}
	return depth * horizontal
}
