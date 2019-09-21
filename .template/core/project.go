package core

import (
	// "log"
	platforms "github.com/YaleOpenLab/openx/platforms"
)

// Project defines the project investment structure in opensolar
type Project struct {
	// The project is split into two parts - parts which are used in the smart contract and parts which are not
	// we define them as critparams and noncritparams

	// start crit params
	Index                int                // an Index to keep track of how many projects exist
	TotalValue           float64            // the total money that we need from investors
	Lock                 bool               // lock investment in order to wait for recipient's confirmation
	LockPwd              string             // the recipient's seedpwd. Will be set to null as soon as we use it.
	Chain                string             // the chain on which the project desires to be.
	OneTimeUnlock        string             // a one time unlock password where the recipient can store his seedpwd in (will be decrypted after investment)
	AmountOwed           float64            // the amoutn owed to investors as a cumulative sum. Used in case of a breach
	Reputation           float64            // the positive reputation associated with a given project
	Votes                float64            // the number of votes towards a proposed contract by investors
	OwnershipShift       float64            // the percentage of the project that the recipient now owns
	StageData            []string           // the data associated with stage migrations
	StageChecklist       []map[string]bool  // the checklist that has to be completed before moving on to the next stage
	InvestorMap          map[string]float64 // publicKey: percentage donation
	SeedInvestorMap      map[string]float64 // the list of all seed investors who've invested in the project
	WaterfallMap         map[string]float64 // publickey:amount map in order to pay multiple accounts. A bit ugly, but should work fine. Make map before using
	RecipientIndex       int                // The index of the project's recipient
	InvestorIndices      []int              // The various investors who have invested in the project
	SeedInvestorIndices  []int              // Investors who took part before the contract was at stage 3
	RecipientIndices     []int              // the indices of the recipient family (offtakers, beneficiaries, etc)
	DateInitiated        string             // date the project was created on the platform
	DateFunded           string             // date that the project completed the stage 4-5 migration
	DateLastPaid         int64              // int64 ie unix time since we need comparisons on this one
	AuctionType          string             // the type of the auction in question. Default is blind auction unless explicitly mentioned
	InvestmentType       string             // the type of investment - equity crowdfunding, municipal bond, normal crowdfunding, etc defined in models
	PaybackPeriod        int                // the frequency in number of weeks that the recipient has to pay the platform
	Stage                int                // the stage at which the contract is at, float due to potential support of 0.5 state changes in the future
	InvestorAssetCode    string             // the code of the asset given to investors on investment in the project
	DebtAssetCode        string             // the code of the asset given to recipients on receiving a project
	PaybackAssetCode     string             // the code of the asset given to recipients on receiving a project
	SeedAssetCode        string             // the code of the asset given to seed investors on seed investment in the project
	SeedInvestmentFactor float64            // the factor that a seed investor's investment is multiplied by in case he does invest at the seed stage
	SeedInvestmentCap    float64            // the max amount that a seed investor can put in a project when it is in its seed stages
	EscrowPubkey         string             // the publickey of the escrow we setup after project investment
	EscrowLock           bool               // used to lock the escrow in case someting goes wrong
	MoneyRaised          float64            // total money that has been raised until now
	SeedMoneyRaised      float64            // total seed money that has been raised until now
	EstimatedAcquisition int                // the year in which the recipient is expected to repay the initial investment amount by
	BalLeft              float64            // denotes the balance left to pay by the party, percentage raised is not stored in the database since that can be calculated
	AdminFlagged         bool               // denotes if someone reports the project as flagged
	FlaggedBy            int                // the index of the admin who flagged the project
	UserFlaggedBy        []int              // the indices of the users who flagged the project
	Reports              int                // the number of reports against htis particular project
	BrokerUrl            string             // the url of the MQTT broker that is associated with the project
	Metadata             string             // other metadata which does not have an explicit name can be stored here. Used to derive assetIDs
	InterestRate         float64            // the rate of return for investors
	OriginatorFee        float64            // fee paid to the originator included in the total value of the project
}

// Stage is the evolution of the erstwhile static stage integer construction
type Stage struct {
	Number          int
	FriendlyName    string   // the informal name that one can use while referring to the stage (nice for UI as well)
	Name            string   // this is a more formal name to give to the given stage
	Activities      []string // the activities that are covered in this particular stage and need to be fulfilled in order to move to the next stage.
	StateTrigger    []string // trigger state change from n to n+1
	BreachCondition []string // define breach conditions for a particular stage
}

const (
	// InvestorWeight is the percentage weight of the project's total reputation assigned to the investor
	InvestorWeight = 0.7

	// RecipientWeight is the percentage weight of the project's total reputation assigned to the recipient
	RecipientWeight = 0.3

	// NormalThreshold is the normal payback interval of 1 payback period. Regular notifications are sent regardless of whether the user has paid back towards the project.
	NormalThreshold = 1

	// AlertThreshold is the threshold above which the user gets a nice email requesting a quick payback whenever possible
	AlertThreshold = 2

	// SternAlertThreshold is the threshold above when the user gets a warning that services will be disconnected if the user doesn't payback soon.
	SternAlertThreshold = 4

	// DisconnectionThreshold is the threshold above which the user gets a notification telling that services have been disconnected.
	DisconnectionThreshold = 6
)

// InitializePlatform imports handlers from the main platform struct that are necessary for starting the platform
func InitializePlatform() error {
	return platforms.InitializePlatform()
}

// RefillPlatform checks whether the platform has any xlm and if its balance
// is less than 21 XLM, it proceeds to ask friendbot for more test xlm
func RefillPlatform(publicKey string) error {
	return platforms.RefillPlatform(publicKey)
}
