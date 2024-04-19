package andres_ruiz_golang

import (
	"testing"
)

func TestCreateBasicBidConfig(t *testing.T) {
	b := NewBid(50, 80, 10)

	True(t, b.Initial == 50, "initial value must be 50")
	True(t, b.Max == 80, "max value must be 80")
	True(t, b.Increment == 10, "increment value must be 10")
	True(t, b.ActualMax == 80, "actual max must be 80, got %d", b.ActualMax)
}

func TestBidEqual(t *testing.T) {
	b1 := NewBid(50, 50, 1)
	b2 := NewBid(51, 51, 1)

	True(t, BidEqual(b1, b1), "bid must be equal to itself")
	True(t, BidEqual(b2, b2), "bid must be equal to itself")
	True(t, !BidEqual(b1, b2), "bids must be different")
}

var winningBidTests = []struct {
	name           string
	bids           []*Bid
	expectedWinner *Bid
}{
	{
		"simple - only 1 bid",
		[]*Bid{
			NewBid(51, 51, 1),
		},
		NewBid(51, 51, 1),
	},
	{
		"simple - first bid wins",
		[]*Bid{
			NewBid(51, 51, 1),
			NewBid(50, 50, 1),
		},
		NewBid(51, 51, 1),
	},
	{
		"simple - second bid wins",
		[]*Bid{
			NewBid(50, 50, 1),
			NewBid(51, 51, 1),
		},
		NewBid(51, 51, 1),
	},
	{
		"simple - third bid wins",
		[]*Bid{
			NewBid(50, 50, 1),
			NewBid(51, 51, 1),
			NewBid(52, 52, 1),
		},
		NewBid(52, 52, 1),
	},
	{
		"inc by one - third bid wins",
		[]*Bid{
			NewBid(1, 50, 1),
			NewBid(1, 51, 1),
			NewBid(1, 52, 1),
		},
		NewBid(1, 52, 1),
	},
	{
		"start at 1 and 2 - second bid wins ",
		[]*Bid{
			NewBid(1, 50, 2),
			NewBid(2, 50, 2),
		},
		NewBid(2, 50, 2),
	},
	{
		"assignment test case 1",
		[]*Bid{
			NewBid(5000, 8000, 300),
			NewBid(6000, 8200, 200),
			NewBid(5500, 8500, 500),
		},
		NewBid(5500, 8500, 500),
	},
	{
		"assignment test case 2",
		[]*Bid{
			NewBid(70000, 72500, 200),
			NewBid(59900, 72500, 1500),
			NewBid(62500, 72500, 800),
		},
		NewBid(70000, 72500, 200),
	},
	{
		"assignment test case 3",
		[]*Bid{
			NewBid(250000, 300000, 50000),
			NewBid(280000, 310000, 20100),
			NewBid(250100, 320000, 24700),
		},
		NewBid(280000, 310000, 20100),
	},
}

func TestWinningBid(t *testing.T) {
	for _, tt := range winningBidTests {
		t.Run(tt.name, func(t *testing.T) {
			actualWinner := GetWinningBid(tt.bids)

			True(t, BidEqual(tt.expectedWinner, actualWinner), "expected (%+v) and actual (%+v) winner dont match", tt.expectedWinner, actualWinner)
		})

	}
}
