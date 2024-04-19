package andres_ruiz_golang

type Bid struct {
	Initial   int
	Max       int
	Increment int
	ActualMax int
}

func NewBid(initial, max, increment int) *Bid {
	return &Bid{
		Initial:   initial,
		Max:       max,
		Increment: increment,
		ActualMax: getActualMax(initial, max, increment),
	}
}

func getActualMax(initial, max, increment int) int {
	diff := max - initial
	totalIncrements := diff / increment

	return initial + (totalIncrements * increment)
}

func BidEqual(b1 *Bid, b2 *Bid) bool {
	return b1.Initial == b2.Initial && b1.Max == b2.Max && b1.Increment == b2.Increment
}

func GetWinningBid(bids []*Bid) *Bid {
	maxBid := bids[0]

	for _, b := range bids[1:] {
		if b.ActualMax > maxBid.ActualMax {
			maxBid = b
		}
	}

	return maxBid
}
