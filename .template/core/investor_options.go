package core

// ChangeVotingBalance changes the voting balance of a user
func (a *Investor) ChangeVotingBalance(votes float64) error {
	// this function is caled when we want to refund the user with the votes once
	// an order has been finalized.
	a.VotingBalance += votes
	if a.VotingBalance < 0 {
		a.VotingBalance = 0 // to ensure no one has negative votes or something
	}
	return a.Save()
}
