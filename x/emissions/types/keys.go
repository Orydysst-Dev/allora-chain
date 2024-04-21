package types

import "cosmossdk.io/collections"

const ModuleName = "emissions"
const AlloraStakingAccountName = "allorastaking"
const AlloraRequestsAccountName = "allorarequests"
const AlloraRewardsAccountName = "allorarewards"

var (
	ParamsKey                          = collections.NewPrefix(0)
	TotalStakeKey                      = collections.NewPrefix(1)
	TopicStakeKey                      = collections.NewPrefix(2)
	LastRewardsUpdateKey               = collections.NewPrefix(3)
	RewardsKey                         = collections.NewPrefix(4)
	NextTopicIdKey                     = collections.NewPrefix(5)
	TopicsKey                          = collections.NewPrefix(6)
	TopicWorkersKey                    = collections.NewPrefix(7)
	TopicReputersKey                   = collections.NewPrefix(8)
	DelegatorStakeKey                  = collections.NewPrefix(9)
	DelegatedStakePlacementKey         = collections.NewPrefix(10)
	TargetStakeKey                     = collections.NewPrefix(11)
	InferencesKey                      = collections.NewPrefix(12)
	ForecastsKey                       = collections.NewPrefix(13)
	WorkerNodesKey                     = collections.NewPrefix(14)
	ReputerNodesKey                    = collections.NewPrefix(15)
	LatestInferencesTsKey              = collections.NewPrefix(16)
	MempoolKey                         = collections.NewPrefix(17)
	RequestUnmetDemandKey              = collections.NewPrefix(18)
	TopicUnmetDemandKey                = collections.NewPrefix(19)
	AllInferencesKey                   = collections.NewPrefix(20)
	AllForecastsKey                    = collections.NewPrefix(21)
	AllLossBundlesKey                  = collections.NewPrefix(22)
	StakeRemovalQueueKey               = collections.NewPrefix(23)
	StakeByReputerAndTopicId           = collections.NewPrefix(24)
	DelegatedStakeRemovalQueueKey      = collections.NewPrefix(25)
	AllTopicStakeSumKey                = collections.NewPrefix(26)
	WhitelistAdminsKey                 = collections.NewPrefix(27)
	TopicCreationWhitelistKey          = collections.NewPrefix(28)
	ReputerWhitelistKey                = collections.NewPrefix(29)
	ChurnReadyTopicsKey                = collections.NewPrefix(30)
	NetworkLossBundlesKey              = collections.NewPrefix(31)
	NetworkRegretsKey                  = collections.NewPrefix(32)
	StakeByReputerAndTopicIdKey        = collections.NewPrefix(33)
	ReputerScoresKey                   = collections.NewPrefix(34)
	InferenceScoresKey                 = collections.NewPrefix(35)
	ForecastScoresKey                  = collections.NewPrefix(36)
	ReputerListeningCoefficientKey     = collections.NewPrefix(37)
	InfererNetworkRegretsKey           = collections.NewPrefix(38)
	ForecasterNetworkRegretsKey        = collections.NewPrefix(39)
	OneInForecasterNetworkRegretsKey   = collections.NewPrefix(40)
	UnfulfilledWorkerNoncesKey         = collections.NewPrefix(41)
	UnfulfilledReputerNoncesKey        = collections.NewPrefix(42)
	AverageWorkerRewardKey             = collections.NewPrefix(43)
	FeeRevenueEpochKey                 = collections.NewPrefix(44)
	TopicFeeRevenueKey                 = collections.NewPrefix(45)
	PreviousTopicWeightKey             = collections.NewPrefix(46)
	PreviousReputerRewardFractionKey   = collections.NewPrefix(47)
	PreviousInferenceRewardFractionKey = collections.NewPrefix(48)
	PreviousForecastRewardFractionKey  = collections.NewPrefix(49)
	LatestInfererScoresByWorkerKey     = collections.NewPrefix(50)
	LatestForecasterScoresByWorkerKey  = collections.NewPrefix(51)
	LatestReputerScoresByReputerKey    = collections.NewPrefix(52)
	ActiveTopicsKey                    = collections.NewPrefix(53)
)
