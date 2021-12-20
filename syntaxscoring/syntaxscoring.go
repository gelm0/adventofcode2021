package syntaxscoring

func ScoreSyntax(line string) (int, []rune) {
	var stack []rune
	var scoreMap = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	for _, val := range line {
		if val == '{' || val == '[' || val == '(' || val == '<' {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 {
				return 0, nil
			}
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if current == '{' && val != '}' {
				return scoreMap[val], nil
			}
			if current == '(' && val != ')' {
				return scoreMap[val], nil
			}
			if current == '[' && val != ']' {
				return scoreMap[val], nil
			}
			if current == '<' && val != '>' {
				return scoreMap[val], nil
			}

		}
	}
	return 0, stack
}

func ScoreAutoComplete(toComplete []rune) int {
	score := 0
	for i := len(toComplete) - 1; i >= 0; i-- {
		val := toComplete[i]
		score *= 5
		if val == '{' {
			score += 3
		}
		if val == '[' {
			score += 2
		}
		if val == '(' {
			score += 1
		}
		if val == '<' {
			score += 4
		}
	}
	return score
}
