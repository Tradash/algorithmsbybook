package main

import (
	"fmt"
)

// Планирование турнира, с47

type champ struct {
	participant02 *participant
	participant01 *participant
	day           int
}

type game struct {
	day   int
	enemy int
}

type participant struct {
	code  int
	games []game
}

func createStruct(p []int) []participant {
	ps := make([]participant, 0, 0)
	for _, m := range p {
		ps = append(ps, participant{code: m})
	}
	return ps
}

func findFreePlayers(ps *[]participant, maxMatch int, day int, enemy int) int {

	for i, p := range *ps {

		if p.code == enemy {
			continue
		}

		if len(p.games) >= maxMatch {
			continue
		}

		if len(p.games) == 0 {
			return i
		}

		isFree := true
		for _, d := range p.games {
			if d.day == day || d.enemy == enemy {
				isFree = false
			}
		}
		if isFree {
			return i
		}

	}
	return -1
}

func printCh(ch *[]champ) {
	for i, d := range *ch {
		fmt.Printf("Матч № %d, День %d, Код игрока №1 %d, Код игрока №2 %d\n", i+1, d.day, d.participant01.code, d.participant02.code)
	}
}

func createStanding(ps []participant, maxMatch int, maxDays int) []champ {
	ch := make([]champ, 0, 0)
	for d := 1; d < maxDays; d++ {

		for m := 0; m < maxMatch; m++ {

			i := findFreePlayers(&ps, maxMatch, d, -1)
			if i == -1 {
				return ch
			}
			j := findFreePlayers(&ps, maxMatch, d, ps[i].code)
			if i == -1 || j == -1 {
				panic("Не возможно создать турнирную таблицу")
			}
			ps[i].games = append(ps[i].games, game{d, ps[j].code})
			ps[j].games = append(ps[j].games, game{d, ps[i].code})

			ch = append(ch, champ{
				participant02: &ps[j],
				participant01: &ps[i],
				day:           d,
			})

		}
	}

	return ch
}

func main() {
	participantCode := []int{1, 2, 3, 4, 5, 6, 7, 8}
	maxMatch := 4
	maxDays := len(participantCode) * maxMatch / 2

	ps := createStruct(participantCode)
	fmt.Println(maxDays, ps)

	ch := createStanding(ps, maxMatch, maxDays)
	printCh(&ch)
}
