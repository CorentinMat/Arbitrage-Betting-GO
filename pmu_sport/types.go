package PMU

// Final Data
type CleanBet struct {
	Title string
	Odd   float64
}
type CleanMatch struct {
	Title string
	Bet1  CleanBet
	Bet2  CleanBet
	Draw  CleanBet
}
type PMUData []CleanMatch
