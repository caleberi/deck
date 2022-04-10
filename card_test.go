package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Three, Suit: Club})
	fmt.Println(Card{Rank: Nine, Suit: Spade})
	fmt.Println(Card{Rank: King, Suit: Club})
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Hearts
	// Three of Clubs
	// Nine of Spades
	// King of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in new deck")
	}
}

func TestDeck(t *testing.T) {
	cards := Deck(2)(New())
	if len(cards) != 13*4*2 {
		t.Error("Wrong number of cards in new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	card := New(DefaultSort)
	expect := Card{Rank: Ace, Suit: Spade}
	if card[0] != expect {
		t.Errorf("Expected Ace of Hearts but Recieved : %s", card[0])
	}
}

func TestSort(t *testing.T) {
	card := New(Sort(Less))
	expect := Card{Rank: Ace, Suit: Spade}
	if card[0] != expect {
		t.Errorf("Expected Ace of Hearts but Recieved : %s", card[0])
	}
}

func TestAddJoker(t *testing.T) {
	jokerNos := 3
	cards := (New(Sort(Less), AddJoker(jokerNos)))
	idx := sort.Search(len(cards), func(i int) bool {
		return cards[i].Suit == Joker
	})
	if idx == -1 {
		t.Error("Did not find a card with Joker suit")
	}
	cnt := 0
	for _, c := range cards {
		if c.Suit == Joker {
			cnt++
		}
	}
	if cnt != 3 {
		t.Errorf("Expected %d Jokers , Recieved %d", jokerNos, cnt)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Errorf("Expected all twos and threes to be filtered out")
		}
	}

}

func TestShuffle(t *testing.T) {
	_ = rand.New(rand.NewSource(0))
	original := New()
	first := original[40]
	second := original[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s , received %s", first, cards[0])
	}

	if cards[1] != second {
		t.Errorf("Expected the first card to be %s , received %s", second, cards[1])
	}
}
