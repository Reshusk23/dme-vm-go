package arwenmandos

// ReshuskProtectedKeyPrefix prefixes all Reshusk reserved storage. Only the protocol can write to keys starting with this.
const ReshuskProtectedKeyPrefix = "RESHUSK"

// ReshuskRewardKey is the storage key where the protocol writes when sending out rewards.
const ReshuskRewardKey = ReshuskProtectedKeyPrefix + "reward"
