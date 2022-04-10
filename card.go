//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var suits = [...]Suit{
	Spade, Diamond, Club, Heart,
}

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Shuffle(cards []Card) []Card {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func Less(c []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(c[i]) < absRank(c[j])
	}
}

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Rank: rank, Suit: suit})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func absRank(c Card) int {
	return int(c.Suit) * int(maxRank) * int(c.Rank)
}

func AddJoker(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

func Filter(filterFunc func(card Card) bool) func([]Card) []Card {
	return func(c []Card) []Card {
		var ret []Card
		for i := 0; i < len(c); i++ {
			if !filterFunc(c[i]) {
				ret = append(ret, c[i])
			}
		}
		return ret
	}
}

func Deck(n int) func(c []Card) []Card {
	return func(c []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, c...)
		}
		return ret
	}
}
