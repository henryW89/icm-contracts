// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenstakingmanager

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ConversionData is an auto generated low-level Go binding around an user-defined struct.
type ConversionData struct {
	SubnetID                     [32]byte
	ValidatorManagerBlockchainID [32]byte
	ValidatorManagerAddress      common.Address
	InitialValidators            []InitialValidator
}

// InitialValidator is an auto generated low-level Go binding around an user-defined struct.
type InitialValidator struct {
	NodeID       []byte
	BlsPublicKey []byte
	Weight       uint64
}

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// PoSValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type PoSValidatorManagerSettings struct {
	BaseSettings             ValidatorManagerSettings
	MinimumStakeAmount       *big.Int
	MaximumStakeAmount       *big.Int
	MinimumStakeDuration     uint64
	MinimumDelegationFeeBips uint16
	MaximumStakeMultiplier   uint8
	WeightToValueFactor      *big.Int
	RewardCalculator         common.Address
	UptimeBlockchainID       [32]byte
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Status         uint8
	NodeID         []byte
	StartingWeight uint64
	SentNonce      uint64
	ReceivedNonce  uint64
	Weight         uint64
	StartTime      uint64
	EndTime        uint64
}

// ValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type ValidatorManagerSettings struct {
	SubnetID               [32]byte
	ChurnPeriodSeconds     uint64
	MaximumChurnPercentage uint8
}

// ValidatorMessagesValidationPeriod is an auto generated low-level Go binding around an user-defined struct.
type ValidatorMessagesValidationPeriod struct {
	SubnetID              [32]byte
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
	Weight                uint64
}

// ValidatorRegistrationInput is an auto generated low-level Go binding around an user-defined struct.
type ValidatorRegistrationInput struct {
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}

// NativeTokenStakingManagerMetaData contains all meta data concerning the NativeTokenStakingManager contract.
var NativeTokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"InvalidBLSKeyLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"encodedConversionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedConversionID\",\"type\":\"bytes32\"}],\"name\":\"InvalidConversionID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"}],\"name\":\"InvalidDelegationFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidDelegationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidDelegatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"name\":\"InvalidMaximumChurnPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"InvalidMinStakeDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"InvalidNodeID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addressesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidPChainOwnerThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"InvalidRegistrationExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InvalidRewardRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"}],\"name\":\"InvalidStakeMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InvalidTotalWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidUptimeBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"}],\"name\":\"InvalidValidatorManagerAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidatorManagerBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"churnAmount\",\"type\":\"uint64\"}],\"name\":\"MaxChurnRateExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValidatorWeight\",\"type\":\"uint64\"}],\"name\":\"MaxWeightExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"MinStakeDurationNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PChainOwnerAddressesNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UnauthorizedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"validRegistration\",\"type\":\"bool\"}],\"name\":\"UnexpectedRegistrationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedValidationID\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorNotPoS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroWeightToValueFactor\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"CompletedValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"CompletedValidatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"CompletedValidatorWeightUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"DelegationEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"DelegatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeID\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"registrationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"InitiatedValidatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"weightUpdateMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedValidatorWeightUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeID\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"RegisteredInitialValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BIPS_CONVERSION_FACTOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_ID_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeDelegatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeValidatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorWeightUpdate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sentNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receivedNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"internalType\":\"structValidatorManagerSettings\",\"name\":\"baseSettings\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"weightToValueFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structPoSValidatorManagerSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initializeDelegatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"internalType\":\"structValidatorRegistrationInput\",\"name\":\"registrationInput\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"initializeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1TotalWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"registeredValidators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subnetID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"weightToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051615be6380380615be683398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b615a998061014d5f395ff3fe608060405260043610610254575f3560e01c806380dd672f1161013f578063ba3a4b97116100b3578063c974d1b611610078578063c974d1b6146106cd578063ce161f14146106e1578063d5f20ff61461071d578063df93d8de14610749578063fb8b11dd1461075f578063fd7ac5e71461077e575f80fd5b8063ba3a4b9714610635578063bb0b193814610654578063bc5fbfec14610668578063bee0a03f1461069b578063c599e24f146106ba575f80fd5b80639ae06447116101045780639ae0644714610597578063a3a65e48146105b6578063a9778a7a146103a8578063af2f5feb146105d5578063afb98096146105e8578063b771b3bc1461061b575f80fd5b806380dd672f146105075780638280a25a146105265780638ef34c981461053a57806393e24598146105595780639681d94014610578575f80fd5b806337b9be8f116101d657806360ad77841161019b57806360ad778414610478578063620658561461049757806363e2ca971461044f578063732214f8146104b657806376f78621146104c95780637d8d2f77146104e8575f80fd5b806337b9be8f146103d05780633a1cfff6146103ef5780635dc1f5351461040e5780635dd6a6cb1461043057806360305d621461044f575f80fd5b806320d91b7a1161021c57806320d91b7a1461030157806325e1c776146103205780632e2194d81461033f578063329c3e121461037657806335455ded146103a8575f80fd5b80630118acc4146102585780630322ed98146102795780630ba512d114610298578063151d30d1146102b75780631ec44724146102e2575b5f80fd5b348015610263575f80fd5b50610277610272366004614b41565b61079d565b005b348015610284575f80fd5b50610277610293366004614b7c565b6107ae565b3480156102a3575f80fd5b506102776102b2366004614b93565b610a41565b3480156102c2575f80fd5b506102cb600a81565b60405160ff90911681526020015b60405180910390f35b3480156102ed575f80fd5b506102776102fc366004614b41565b610b1e565b34801561030c575f80fd5b5061027761031b366004614baa565b610b2a565b34801561032b575f80fd5b5061027761033a366004614bf8565b6110de565b34801561034a575f80fd5b5061035e610359366004614b7c565b611152565b6040516001600160401b0390911681526020016102d9565b348015610381575f80fd5b506103906001600160991b0181565b6040516001600160a01b0390911681526020016102d9565b3480156103b3575f80fd5b506103bd61271081565b60405161ffff90911681526020016102d9565b3480156103db575f80fd5b506102776103ea366004614c2d565b6111a6565b3480156103fa575f80fd5b50610277610409366004614b41565b6111b9565b348015610419575f80fd5b506104226111c5565b6040519081526020016102d9565b34801561043b575f80fd5b5061027761044a366004614c2d565b6111d4565b34801561045a575f80fd5b50610463601481565b60405163ffffffff90911681526020016102d9565b348015610483575f80fd5b50610277610492366004614bf8565b6111e0565b3480156104a2575f80fd5b506104226104b1366004614c8f565b611433565b3480156104c1575f80fd5b506104225f81565b3480156104d4575f80fd5b506102776104e3366004614b41565b611453565b3480156104f3575f80fd5b50610277610502366004614c2d565b61145f565b348015610512575f80fd5b50610277610521366004614bf8565b61146b565b348015610531575f80fd5b506102cb603081565b348015610545575f80fd5b50610277610554366004614caa565b61166c565b348015610564575f80fd5b50610277610573366004614b7c565b61171d565b348015610583575f80fd5b50610422610592366004614cd8565b6117b1565b3480156105a2575f80fd5b506102776105b1366004614c2d565b61188a565b3480156105c1575f80fd5b506104226105d0366004614cd8565b611896565b6104226105e3366004614d02565b611a85565b3480156105f3575f80fd5b506104227f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0081565b348015610626575f80fd5b506103906005600160991b0181565b348015610640575f80fd5b5061027761064f366004614b7c565b611ab9565b34801561065f575f80fd5b5061035e611d12565b348015610673575f80fd5b506104227fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0081565b3480156106a6575f80fd5b506102776106b5366004614b7c565b611d34565b6104226106c8366004614b7c565b611e50565b3480156106d8575f80fd5b506102cb601481565b3480156106ec575f80fd5b506107006106fb366004614cd8565b611e7c565b604080519283526001600160401b039091166020830152016102d9565b348015610728575f80fd5b5061073c610737366004614b7c565b611fff565b6040516102d99190614dd8565b348015610754575f80fd5b5061035e6202a30081565b34801561076a575f80fd5b50610277610779366004614caa565b61214e565b348015610789575f80fd5b50610422610798366004614e8e565b6121e5565b6107a98383835f61221d565b505050565b5f6107b7612249565b5f838152600580830160205260408083208151610100810190925280549495509293909291839160ff16908111156107f1576107f1614d63565b600581111561080257610802614d63565b815260200160018201805461081690614ef9565b80601f016020809104026020016040519081016040528092919081815260200182805461084290614ef9565b801561088d5780601f106108645761010080835404028352916020019161088d565b820191905f5260205f20905b81548152906001019060200180831161087057829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b9092048116608084015260039384015480821660a0850152919091041660c0909101529091508151600581111561090557610905614d63565b14610941575f8381526005830160205260409081902054905163170cc93360e21b81526109389160ff1690600401614f2b565b60405180910390fd5b606081015160405163854a893f60e01b8152600481018590526001600160401b0390911660248201525f60448201526005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__9063854a893f906064015f60405180830381865af41580156109b8573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526109df9190810190615034565b6040518263ffffffff1660e01b81526004016109fb9190615065565b6020604051808303815f875af1158015610a17573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a3b9190615077565b50505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff1680610a8a575080546001600160401b03808416911610155b15610aa85760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b178155610ad28361226d565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050565b610a3b8383835f612281565b5f610b33612249565b600781015490915060ff1615610b5c57604051637fab81e560e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b9f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bc39190615077565b836020013514610bec576040516372b0a7e760e11b815260208401356004820152602401610938565b30610bfd606085016040860161508e565b6001600160a01b031614610c4057610c1b606084016040850161508e565b604051632f88120d60e21b81526001600160a01b039091166004820152602401610938565b5f610c4e60608501856150a9565b905090505f805b828163ffffffff161015610ed4575f610c7160608801886150a9565b8363ffffffff16818110610c8757610c876150ee565b9050602002810190610c999190615102565b610ca29061516d565b80516040519192505f916006880191610cba916151e8565b90815260200160405180910390205414610cea57805160405163a41f772f60e01b81526109389190600401615065565b805151601414610d10578051604051633e08a12560e11b81526109389190600401615065565b5f6002885f013584604051602001610d3f92919091825260e01b6001600160e01b031916602082015260240190565b60408051601f1981840301815290829052610d59916151e8565b602060405180830381855afa158015610d74573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d979190615077565b90508086600601835f0151604051610daf91906151e8565b90815260408051918290036020908101909220929092555f8381526005890190915220805460ff191660021781558251600190910190610def908261523d565b50604082810180515f84815260058a016020529290922060028101805492516001600160401b0394851667ffffffffffffffff60801b90941693909317600160c01b858516021790556003018054429093166001600160801b031990931692909217909155610e5e908561530c565b8251602001519094506bffffffffffffffffffffffff1916817f9d9c026e2cadfec89cccc2cd72705360eca1beba24774f3363f4bb33faabc7d78460400151604051610eb991906001600160401b0391909116815260200190565b60405180910390a3505080610ecd9061532c565b9050610c55565b5060038301805467ffffffffffffffff60401b1916600160401b6001600160401b0384168102919091179091556001840154606491610f17910460ff168361534e565b6001600160401b03161015610f4a57604051633e1a785160e01b81526001600160401b0382166004820152602401610938565b5f73__$fd0c147b4031eef6079b0498cbafa865f0$__634d847884610f6e876125b6565b604001516040518263ffffffff1660e01b8152600401610f8e9190615065565b602060405180830381865af4158015610fa9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fcd9190615077565b90505f73__$fd0c147b4031eef6079b0498cbafa865f0$__6387418b8e886040518263ffffffff1660e01b815260040161100791906154a4565b5f60405180830381865af4158015611021573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526110489190810190615034565b90505f60028260405161105b91906151e8565b602060405180830381855afa158015611076573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906110999190615077565b90508281146110c55760405163baaea89d60e01b81526004810182905260248101849052604401610938565b5050506007909201805460ff1916600117905550505050565b6110e7826126cc565b611107576040516330efa98b60e01b815260048101839052602401610938565b5f61111183611fff565b519050600281600581111561112857611128614d63565b14611148578060405163170cc93360e21b81526004016109389190614f2b565b610a3b83836126f5565b5f8061115c61299b565b6003015461116a9084615538565b905080158061117f57506001600160401b0381115b156111a05760405163222d164360e21b815260048101849052602401610938565b92915050565b6111b284848484612281565b5050505050565b610a3b8383835f6129bf565b5f6111ce612249565b54919050565b610a3b84848484612c06565b5f6111e961299b565b5f848152600782016020526040808220815160e0810190925280549394509192909190829060ff16600381111561122257611222614d63565b600381111561123357611233614d63565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f6112a982611fff565b90506001835160038111156112c0576112c0614d63565b146112e1578251604051633b0d540d60e21b81526109389190600401615557565b6004815160058111156112f6576112f6614d63565b0361130c5761130486612c32565b505050505050565b8260a001516001600160401b031681608001516001600160401b031610156113ac575f8061133987611e7c565b915091508184146113675760405163fee3144560e01b81526004810183905260248101859052604401610938565b8460a001516001600160401b0316816001600160401b031610156113a957604051632e19bc2d60e11b81526001600160401b0382166004820152602401610938565b50505b5f868152600785016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b810267ffffffffffffffff60401b19909216919091179091559151918252839188917f047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6910160405180910390a3505050505050565b5f61143c61299b565b600301546111a0906001600160401b038416615571565b6107a98383835f612c06565b6111b2848484846129bf565b611473612e03565b5f61147c61299b565b5f848152600782016020526040808220815160e0810190925280549394509192909190829060ff1660038111156114b5576114b5614d63565b60038111156114c6576114c6614d63565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c090910152905060038151600381111561153f5761153f614d63565b14611560578051604051633b0d540d60e21b81526109389190600401615557565b5f61156e8260400151611fff565b9050600461157f8360400151611fff565b51600581111561159157611591614d63565b141580156115b857508160c001516001600160401b031681608001516001600160401b0316105b15611646575f806115c886611e7c565b91509150818460400151146116015781846040015160405163fee3144560e01b8152600401610938929190918252602082015260400190565b806001600160401b03168460c001516001600160401b0316111561164357604051632e19bc2d60e11b81526001600160401b0382166004820152602401610938565b50505b61164f85612c32565b50505061166860015f80516020615a4483398151915255565b5050565b5f61167561299b565b90506001600160a01b0382166116a95760405163caa903f960e01b81526001600160a01b0383166004820152602401610938565b5f8381526006820160205260409020546001600160a01b031633146116ef57335b604051636e2ccd7560e11b81526001600160a01b039091166004820152602401610938565b5f928352600b01602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b5f61172661299b565b90505f61173283611fff565b519050600481600581111561174957611749614d63565b14611769578060405163170cc93360e21b81526004016109389190614f2b565b5f8381526006830160205260409020546001600160a01b0316331461178e57336116ca565b5f8381526006830160205260409020546107a9906001600160a01b031684612e4d565b5f6117ba612e03565b5f6117c361299b565b90505f806117d085612e77565b915091506117dd826126cc565b6117eb5750915061186f9050565b5f828152600684016020908152604080832054600b870190925290912080546001600160a01b031981169091556001600160a01b0391821691168061182d5750805b60048351600581111561184257611842614d63565b03611851576118518185612e4d565b611867826118628560400151611433565b61322c565b509193505050505b61188560015f80516020615a4483398151915255565b919050565b610a3b8484848461221d565b5f806118a0612249565b90505f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63021de88f6118c7876125b6565b604001516040518263ffffffff1660e01b81526004016118e79190615065565b6040805180830381865af4158015611901573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119259190615588565b915091508061194b57604051632d07135360e01b81528115156004820152602401610938565b5f8281526004840160205260409020805461196590614ef9565b90505f036119895760405163089938b360e11b815260048101839052602401610938565b60015f838152600580860160205260409091205460ff16908111156119b0576119b0614d63565b146119e3575f8281526005840160205260409081902054905163170cc93360e21b81526109389160ff1690600401614f2b565b5f82815260048401602052604081206119fb91614a8d565b5f828152600584016020908152604091829020805460ff1916600290811782556003820180546001600160401b0342811667ffffffffffffffff19909216919091179091559101549251600160c01b90930416825283917f967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06910160405180910390a2509392505050565b5f611a8e612e03565b611a9a8484843461323f565b9050611ab260015f80516020615a4483398151915255565b9392505050565b5f611ac261299b565b5f838152600782016020526040808220815160e0810190925280549394509192909190829060ff166003811115611afb57611afb614d63565b6003811115611b0c57611b0c614d63565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115611b8557611b85614d63565b14158015611ba65750600381516003811115611ba357611ba3614d63565b14155b15611bc7578051604051633b0d540d60e21b81526109389190600401615557565b5f611bd58260400151611fff565b905080606001516001600160401b03165f03611c07576040516339b894f960e21b815260048101859052602401610938565b604080830151606083015160a0840151925163854a893f60e01b81526005600160991b019363ee5b48eb9373__$fd0c147b4031eef6079b0498cbafa865f0$__9363854a893f93611c7593906004019283526001600160401b03918216602084015216604082015260600190565b5f60405180830381865af4158015611c8f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611cb69190810190615034565b6040518263ffffffff1660e01b8152600401611cd29190615065565b6020604051808303815f875af1158015611cee573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111b29190615077565b5f611d1b612249565b60030154600160401b90046001600160401b0316919050565b5f611d3d612249565b5f8381526004820160205260409020805491925090611d5b90614ef9565b90505f03611d7f5760405163089938b360e11b815260048101839052602401610938565b60015f838152600580840160205260409091205460ff1690811115611da657611da6614d63565b14611dd9575f8281526005820160205260409081902054905163170cc93360e21b81526109389160ff1690600401614f2b565b5f8281526004808301602052604091829020915163ee5b48eb60e01b81526005600160991b019263ee5b48eb92611e1092016155ab565b6020604051808303815f875af1158015611e2c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a99190615077565b5f611e59612e03565b611e6482333461346e565b905061188560015f80516020615a4483398151915255565b5f805f611e88846125b6565b90505f805f73__$fd0c147b4031eef6079b0498cbafa865f0$__6350782b0f85604001516040518263ffffffff1660e01b8152600401611ec89190615065565b606060405180830381865af4158015611ee3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f079190615635565b9250925092505f611f16612249565b5f8581526005820160205260409020600201549091506001600160401b03808516600160401b909204161015611f6a57604051632e19bc2d60e11b81526001600160401b0384166004820152602401610938565b5f8481526005820160205260409081902060020180546001600160401b038616600160801b0267ffffffffffffffff60801b199091161790555184907fc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff990611fea90869086906001600160401b0392831681529116602082015260400190565b60405180910390a25091969095509350505050565b612007614ac4565b5f612010612249565b5f84815260058083016020526040918290208251610100810190935280549394509192839160ff9091169081111561204a5761204a614d63565b600581111561205b5761205b614d63565b815260200160018201805461206f90614ef9565b80601f016020809104026020016040519081016040528092919081815260200182805461209b90614ef9565b80156120e65780601f106120bd576101008083540402835291602001916120e6565b820191905f5260205f20905b8154815290600101906020018083116120c957829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b9092048116608084015260039093015480841660a08401520490911660c0909101529392505050565b6001600160a01b0381166121805760405163caa903f960e01b81526001600160a01b0382166004820152602401610938565b5f61218961299b565b5f8481526007820160205260409020549091506001600160a01b036101009091041633146121b757336116ca565b5f928352600901602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b5f806121ef612249565b905080600601848460405161220592919061566a565b90815260200160405180910390205491505092915050565b61222984848484612281565b610a3b57604051631036cf9160e11b815260048101859052602401610938565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0090565b6122756136ae565b61227e816136f9565b50565b5f8061228b61299b565b5f878152600782016020526040808220815160e0810190925280549394509192909190829060ff1660038111156122c4576122c4614d63565b60038111156122d5576122d5614d63565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f61234b82611fff565b905060028351600381111561236257612362614d63565b14612383578251604051633b0d540d60e21b81526109389190600401615557565b60208301516001600160a01b0316331461241f575f8281526006850160205260409020546001600160a01b031633146123bc57336116ca565b5f82815260068501602052604090205460c08201516123eb91600160b01b90046001600160401b03169061530c565b6001600160401b031642101561241f5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610938565b60028151600581111561243457612434614d63565b036125565760028401546080840151612456916001600160401b03169061530c565b6001600160401b031642101561248a5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610938565b871561249c5761249a82886126f5565b505b5f8981526007850160205260409020805460ff19166003179055606083015160a08201516124d59184916124d09190615679565b613773565b505f8a8152600786016020526040812060020180546001600160401b03909316600160c01b026001600160c01b039093169290921790915561251884888c61393d565b9050828a7f366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed5760405160405180910390a3151594506125ae9350505050565b60048151600581111561256b5761256b614d63565b036125925761257b83878b61393d565b5061258589612c32565b60019450505050506125ae565b805160405163170cc93360e21b81526109389190600401614f2b565b949350505050565b60408051606080820183525f8083526020830152918101919091526040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa15801561261a573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526126419190810190615699565b915091508061266357604051636b2f19e960e01b815260040160405180910390fd5b815115612689578151604051636ba589a560e01b81526004810191909152602401610938565b60208201516001600160a01b0316156126c5576020820151604051624de75d60e31b81526001600160a01b039091166004820152602401610938565b5092915050565b5f806126d661299b565b5f938452600601602052505060409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa158015612740573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526127679190810190615699565b915091508061278957604051636b2f19e960e01b815260040160405180910390fd5b5f61279261299b565b60058101548451919250146127c0578251604051636ba589a560e01b81526004810191909152602401610938565b60208301516001600160a01b0316156127fc576020830151604051624de75d60e31b81526001600160a01b039091166004820152602401610938565b60208301516001600160a01b031615612838576020830151604051624de75d60e31b81526001600160a01b039091166004820152602401610938565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63088c246386604001516040518263ffffffff1660e01b81526004016128759190615065565b6040805180830381865af415801561288f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128b39190615729565b915091508188146128e15760405163fee3144560e01b81526004810183905260248101899052604401610938565b5f8881526006840160205260409020600101546001600160401b039081169082161115612972575f888152600684016020908152604091829020600101805467ffffffffffffffff19166001600160401b038516908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a2612990565b505f8781526006830160205260409020600101546001600160401b03165b979650505050505050565b7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0090565b5f806129c961299b565b90506129d486613b2b565b5f6129de87611fff565b90506129e9876126cc565b6129f8576001925050506125ae565b5f8781526006830160205260409020546001600160a01b03163314612a1d57336116ca565b5f87815260068301602052604090205460c0820151612a4c91600160b01b90046001600160401b03169061530c565b6001600160401b03168160e001516001600160401b03161015612a935760e081015160405163fb6ce63f60e01b81526001600160401b039091166004820152602401610938565b5f8615612aab57612aa488876126f5565b9050612ac9565b505f8781526006830160205260409020600101546001600160401b03165b600483015460408301515f916001600160a01b031690634f22429f90612aee90611433565b60c086015160e0808801516040519185901b6001600160e01b031916825260048201939093526001600160401b0391821660248201819052604482015291811660648301528516608482015260a401602060405180830381865afa158015612b58573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b7c9190615077565b90506001600160a01b038616612ba8575f8981526006850160205260409020546001600160a01b031695505b5f898152600a8501602052604081208054839290612bc790849061574c565b90915550505f898152600b909401602052604090932080546001600160a01b0387166001600160a01b0319909116179055505015159050949350505050565b612c12848484846129bf565b610a3b57604051635bff683f60e11b815260048101859052602401610938565b5f612c3b61299b565b5f838152600782016020526040808220815160e0810190925280549394509192909190829060ff166003811115612c7457612c74614d63565b6003811115612c8557612c85614d63565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c090910152810151909150612cf9613e16565b8260800151612d08919061530c565b6001600160401b0316421015612d3c5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610938565b5f848152600784016020908152604080832080546001600160a81b03191681556001810184905560020183905560098601909152902080546001600160a01b031981169091556001600160a01b031680612d97575060208201515b5f80612da4838886613e31565b91509150612dbd85602001516118628760600151611433565b6040805183815260208101839052859189917f8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993910160405180910390a350505050505050565b5f80516020615a44833981519152805460011901612e3457604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60015f80516020615a4483398151915255565b5f612e5661299b565b5f838152600a8201602052604081208054919055909150610a3b8482613ede565b5f612e80614ac4565b5f612e89612249565b90505f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63021de88f612eb0886125b6565b604001516040518263ffffffff1660e01b8152600401612ed09190615065565b6040805180830381865af4158015612eea573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f0e9190615588565b915091508015612f3557604051632d07135360e01b81528115156004820152602401610938565b5f8281526005808501602052604080832081516101008101909252805491929091839160ff90911690811115612f6d57612f6d614d63565b6005811115612f7e57612f7e614d63565b8152602001600182018054612f9290614ef9565b80601f0160208091040260200160405190810160405280929190818152602001828054612fbe90614ef9565b80156130095780601f10612fe057610100808354040283529160200191613009565b820191905f5260205f20905b815481529060010190602001808311612fec57829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b9092048116608084015260039384015480821660a0850152919091041660c0909101529091508151600581111561308157613081614d63565b141580156130a2575060018151600581111561309f5761309f614d63565b14155b156130c357805160405163170cc93360e21b81526109389190600401614f2b565b6003815160058111156130d8576130d8614d63565b036130e657600481526130eb565b600581525b83600601816020015160405161310191906151e8565b90815260408051602092819003830190205f90819055858152600587810190935220825181548493839160ff191690600190849081111561314457613144614d63565b02179055506020820151600182019061315d908261523d565b506040828101516002830180546060860151608087015160a08801516001600160401b039586166001600160801b031994851617600160401b9387168402176001600160801b0316600160801b928716929092026001600160c01b031691909117600160c01b918616919091021790925560c08601516003909501805460e09097015195841696909116959095179390911602919091179091555183907fafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c905f90a29196919550909350505050565b6116686001600160a01b03831682613f3c565b5f8061324961299b565b600281015490915061ffff600160401b90910481169086161080613272575061271061ffff8616115b1561329657604051635f12e6c360e11b815261ffff86166004820152602401610938565b60028101546001600160401b0390811690851610156132d2576040516202a06d60e11b81526001600160401b0385166004820152602401610938565b80548310806132e45750806001015483115b156133055760405163222d164360e21b815260048101849052602401610938565b825f61331082611152565b90505f6133d76133208a8061575f565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506133619250505060208c018c61575f565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506133a59250505060608d0160408e01614c8f565b6133b260608e018e6157a1565b6133bb906157b5565b6133c860808f018f6157a1565b6133d1906157b5565b87613fcf565b5f818152600686016020908152604080832080546001600160401b039c909c16600160b01b0267ffffffffffffffff60b01b1961ffff9e909e16600160a01b02336001600160b01b0319909e168e17179d909d169c909c178c556001909b01805467ffffffffffffffff19169055600b9096019095529790932080546001600160a01b031916909617909555509395945050505050565b5f8061347861299b565b90505f61348484611152565b90505f61349087611fff565b905061349b876126cc565b6134bb576040516330efa98b60e01b815260048101889052602401610938565b6002815160058111156134d0576134d0614d63565b146134f157805160405163170cc93360e21b81526109389190600401614f2b565b5f828260a00151613502919061530c565b905083600201600a9054906101000a90046001600160401b0316826040015161352b919061534e565b6001600160401b0316816001600160401b0316111561356857604051636d51fe0560e11b81526001600160401b0382166004820152602401610938565b5f806135748a84613773565b915091505f8a836040516020016135a292919091825260c01b6001600160c01b031916602082015260280190565b60408051601f1981840301815291815281516020928301205f81815260078b019093529120805491925060019160ff1916828002179055505f8181526007880160209081526040918290208054610100600160a81b0319166101006001600160a01b038f16908102919091178255600182018f9055600290910180546001600160401b038b81166001600160c01b03199092168217600160801b8a8316908102919091176001600160c01b031690935585519283528916938201939093529283019190915260608201849052908c9083907fb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a2234269060800160405180910390a49a9950505050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166136f757604051631afcd79f60e31b815260040160405180910390fd5b565b6137016136ae565b61370a816143b9565b6137126143d2565b61227e6060820135608083013561372f60c0850160a08601614c8f565b61373f60e0860160c0870161587c565b613750610100870160e08801615895565b61010087013561376861014089016101208a0161508e565b8861014001356143e2565b5f805f61377e612249565b5f868152600582016020526040902060020154909150600160c01b90046001600160401b03166137ae85826145a4565b5f6137b887614811565b5f88815260058501602052604080822060020180546001600160c01b0316600160c01b6001600160401b038c811691820292909217909255915163854a893f60e01b8152600481018c905291841660248301526044820152919250906005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__9063854a893f906064015f60405180830381865af415801561385c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526138839190810190615034565b6040518263ffffffff1660e01b815260040161389f9190615065565b6020604051808303815f875af11580156138bb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138df9190615077565b604080516001600160401b038581168252602082018490528a1681830152905191925089917f6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d39181900360600190a2909450925050505b9250929050565b5f8061394761299b565b90505f6139578660400151611fff565b90505f60038251600581111561396f5761396f614d63565b148061398d575060048251600581111561398b5761398b614d63565b145b1561399d575060e08101516139da565b6002825160058111156139b2576139b2614d63565b036139be5750426139da565b815160405163170cc93360e21b81526109389190600401614f2b565b86608001516001600160401b0316816001600160401b031611613a02575f9350505050611ab2565b600483015460608801515f916001600160a01b031690634f22429f90613a2790611433565b60c086015160808c01516040808e01515f90815260068b0160205281902060010154905160e086901b6001600160e01b031916815260048101949094526001600160401b0392831660248501529082166044840152818716606484015216608482015260a401602060405180830381865afa158015613aa8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613acc9190615077565b90506001600160a01b038716613ae457876020015196505b5f8681526008850160209081526040808320849055600990960190529390932080546001600160a01b0388166001600160a01b031990911617905550909150509392505050565b5f613b34612249565b5f838152600580830160205260408083208151610100810190925280549495509293909291839160ff1690811115613b6e57613b6e614d63565b6005811115613b7f57613b7f614d63565b8152602001600182018054613b9390614ef9565b80601f0160208091040260200160405190810160405280929190818152602001828054613bbf90614ef9565b8015613c0a5780601f10613be157610100808354040283529160200191613c0a565b820191905f5260205f20905b815481529060010190602001808311613bed57829003601f168201915b50505091835250506002828101546001600160401b038082166020850152600160401b80830482166040860152600160801b830482166060860152600160c01b9092048116608085015260039094015480851660a08501520490921660c09091015290915081516005811115613c8257613c82614d63565b14613cb5575f8381526005830160205260409081902054905163170cc93360e21b81526109389160ff1690600401614f2b565b60038152426001600160401b031660e08201525f83815260058381016020526040909120825181548493839160ff1916906001908490811115613cfa57613cfa614d63565b021790555060208201516001820190613d13908261523d565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031994851617600160401b9387168402176001600160801b0316600160801b928716929092026001600160c01b031691909117600160c01b918616919091021790925560c08501516003909401805460e090960151948416959091169490941792909116021790555f613db38482613773565b915050837fbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273828460a0015142604051613e08939291909283526001600160401b03918216602084015216604082015260600190565b60405180910390a250505050565b5f613e1f612249565b600101546001600160401b0316919050565b5f805f613e3c61299b565b5f86815260088201602052604081208054908290559192509081908015613ed0575f87815260068501602052604090205461271090613e8690600160a01b900461ffff1683615571565b613e909190615538565b91508184600a015f8981526020019081526020015f205f828254613eb4919061574c565b90915550613ec4905082826158b5565b9250613ed08984613ede565b509097909650945050505050565b6040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba906044015f604051808303815f87803b158015613f2a575f80fd5b505af1158015611304573d5f803e3d5ffd5b80471015613f5f5760405163cd78605960e01b8152306004820152602401610938565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114613fa8576040519150601f19603f3d011682016040523d82523d5f602084013e613fad565b606091505b50509050806107a957604051630a12f52160e11b815260040160405180910390fd5b5f613fd8612249565b6007015460ff16613ffc57604051637fab81e560e01b815260040160405180910390fd5b5f614005612249565b905042866001600160401b031611158061403457506140276202a3004261574c565b866001600160401b031610155b1561405d57604051635879da1360e11b81526001600160401b0387166004820152602401610938565b60038101546001600160401b039061408090600160401b9004821685831661574c565b11156140aa57604051633e1a785160e01b81526001600160401b0384166004820152602401610938565b6140b38561487a565b6140bc8461487a565b86516030146140e35786516040516326475b2f60e11b815260040161093891815260200190565b87516014146141075787604051633e08a12560e11b81526004016109389190615065565b5f801b816006018960405161411c91906151e8565b9081526020016040518091039020541461414b578760405163a41f772f60e01b81526004016109389190615065565b614155835f6145a4565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63eb97ce516040518060e00160405280865f015481526020018d81526020018c81526020018b6001600160401b031681526020018a8152602001898152602001886001600160401b03168152506040518263ffffffff1660e01b81526004016141d5919061592e565b5f60405180830381865af41580156141ef573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261421691908101906159e5565b5f82815260048601602052604090209193509150614234828261523d565b5081836006018b60405161424891906151e8565b9081526040519081900360200181209190915563ee5b48eb60e01b81525f906005600160991b019063ee5b48eb90614284908590600401615065565b6020604051808303815f875af11580156142a0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142c49190615077565b5f8481526005860160205260409020805460ff191660019081178255919250016142ee8c8261523d565b505f8381526005850160205260409020600281018054600160c01b6001600160401b038a1690810267ffffffffffffffff60801b9092161717905560030180546001600160801b03191690556143458b6020015190565b6bffffffffffffffffffffffff1916837f5881be437bdcb008bfa5f20e32d3e335ccf8ab90ef2818852a251625260af35d838c8a6040516143a2939291909283526001600160401b03918216602084015216604082015260600190565b60405180910390a350909998505050505050505050565b6143c16136ae565b6143c9614999565b61227e816149a1565b6143da6136ae565b6136f7614a85565b6143ea6136ae565b5f6143f361299b565b905061ffff8616158061440b575061271061ffff8716115b1561442f57604051635f12e6c360e11b815261ffff87166004820152602401610938565b878911156144535760405163222d164360e21b8152600481018a9052602401610938565b60ff851615806144665750600a60ff8616115b156144895760405163170db35960e31b815260ff86166004820152602401610938565b614491613e16565b6001600160401b0316876001600160401b031610156144cd576040516202a06d60e11b81526001600160401b0388166004820152602401610938565b835f036144ed5760405163a733007160e01b815260040160405180910390fd5b8161450e57604051632f6bd1db60e01b815260048101839052602401610938565b97885560018801969096556002870180546001600160401b039690961669ffffffffffffffffffff1990961695909517600160401b61ffff95909516949094029390931767ffffffffffffffff60501b191660ff92909216600160501b029190911790925560038401919091556004830180546001600160a01b0319166001600160a01b03909216919091179055600590910155565b5f6145ad612249565b90505f826001600160401b0316846001600160401b031611156145db576145d48385615679565b90506145e8565b6145e58484615679565b90505b60408051608081018252600284015480825260038501546001600160401b038082166020850152600160401b8204811694840194909452600160801b9004909216606082015242911580614655575060018401548151614651916001600160401b03169061574c565b8210155b1561467d576001600160401b038084166060830152828252604082015116602082015261469c565b828160600181815161468f919061530c565b6001600160401b03169052505b60608101516146ac90606461534e565b602082015160018601546001600160401b0392909216916146d79190600160401b900460ff1661534e565b6001600160401b0316101561471057606081015160405163dfae880160e01b81526001600160401b039091166004820152602401610938565b8581604001818151614722919061530c565b6001600160401b0316905250604081018051869190614742908390615679565b6001600160401b03169052506001840154604082015160649161477091600160401b90910460ff169061534e565b6001600160401b031610156147a9576040808201519051633e1a785160e01b81526001600160401b039091166004820152602401610938565b8051600285015560208101516003909401805460408301516060909301516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b0319909316919097161717919091169390931790925550505050565b5f8061481b612249565b5f848152600582016020526040902060020180549192509060089061484f90600160401b90046001600160401b0316615a28565b91906101000a8154816001600160401b0302191690836001600160401b031602179055915050919050565b805163ffffffff16158015614893575060208101515115155b156148c757805160208201515160405163c08a0f1d60e01b815263ffffffff90921660048301526024820152604401610938565b602081015151815163ffffffff16111561490a57805160208201515160405163c08a0f1d60e01b815263ffffffff90921660048301526024820152604401610938565b60015b81602001515181101561166857602082015161492a6001836158b5565b8151811061493a5761493a6150ee565b60200260200101516001600160a01b031682602001518281518110614961576149616150ee565b60200260200101516001600160a01b0316101561499157604051630dbc8d5f60e31b815260040160405180910390fd5b60010161490d565b6136f76136ae565b6149a96136ae565b5f6149b2612249565b82358155905060146149ca6060840160408501615895565b60ff1611806149e957506149e46060830160408401615895565b60ff16155b15614a1d576149fe6060830160408401615895565b604051634a59bbff60e11b815260ff9091166004820152602401610938565b614a2d6060830160408401615895565b60018201805460ff92909216600160401b0260ff60401b19909216919091179055614a5e6040830160208401614c8f565b600191909101805467ffffffffffffffff19166001600160401b0390921691909117905550565b612e3a6136ae565b508054614a9990614ef9565b5f825580601f10614aa8575050565b601f0160209004905f5260205f209081019061227e9190614b09565b604080516101008101909152805f81526060602082018190525f604083018190529082018190526080820181905260a0820181905260c0820181905260e09091015290565b5b80821115614b1d575f8155600101614b0a565b5090565b801515811461227e575f80fd5b803563ffffffff81168114611885575f80fd5b5f805f60608486031215614b53575f80fd5b833592506020840135614b6581614b21565b9150614b7360408501614b2e565b90509250925092565b5f60208284031215614b8c575f80fd5b5035919050565b5f6101608284031215614ba4575f80fd5b50919050565b5f8060408385031215614bbb575f80fd5b82356001600160401b03811115614bd0575f80fd5b830160808186031215614be1575f80fd5b9150614bef60208401614b2e565b90509250929050565b5f8060408385031215614c09575f80fd5b82359150614bef60208401614b2e565b6001600160a01b038116811461227e575f80fd5b5f805f8060808587031215614c40575f80fd5b843593506020850135614c5281614b21565b9250614c6060408601614b2e565b91506060850135614c7081614c19565b939692955090935050565b6001600160401b038116811461227e575f80fd5b5f60208284031215614c9f575f80fd5b8135611ab281614c7b565b5f8060408385031215614cbb575f80fd5b823591506020830135614ccd81614c19565b809150509250929050565b5f60208284031215614ce8575f80fd5b611ab282614b2e565b803561ffff81168114611885575f80fd5b5f805f60608486031215614d14575f80fd5b83356001600160401b03811115614d29575f80fd5b840160a08187031215614d3a575f80fd5b9250614d4860208501614cf1565b91506040840135614d5881614c7b565b809150509250925092565b634e487b7160e01b5f52602160045260245ffd5b60068110614d8757614d87614d63565b9052565b5f5b83811015614da5578181015183820152602001614d8d565b50505f910152565b5f8151808452614dc4816020860160208601614d8b565b601f01601f19169290920160200192915050565b60208152614dea602082018351614d77565b5f6020830151610100806040850152614e07610120850183614dad565b915060408501516001600160401b03808216606087015280606088015116608087015250506080850151614e4660a08601826001600160401b03169052565b5060a08501516001600160401b03811660c08601525060c08501516001600160401b03811660e08601525060e08501516001600160401b038116858301525090949350505050565b5f8060208385031215614e9f575f80fd5b82356001600160401b0380821115614eb5575f80fd5b818501915085601f830112614ec8575f80fd5b813581811115614ed6575f80fd5b866020828501011115614ee7575f80fd5b60209290920196919550909350505050565b600181811c90821680614f0d57607f821691505b602082108103614ba457634e487b7160e01b5f52602260045260245ffd5b602081016111a08284614d77565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715614f6f57614f6f614f39565b60405290565b604080519081016001600160401b0381118282101715614f6f57614f6f614f39565b604051601f8201601f191681016001600160401b0381118282101715614fbf57614fbf614f39565b604052919050565b5f6001600160401b03821115614fdf57614fdf614f39565b50601f01601f191660200190565b5f82601f830112614ffc575f80fd5b815161500f61500a82614fc7565b614f97565b818152846020838601011115615023575f80fd5b6125ae826020830160208701614d8b565b5f60208284031215615044575f80fd5b81516001600160401b03811115615059575f80fd5b6125ae84828501614fed565b602081525f611ab26020830184614dad565b5f60208284031215615087575f80fd5b5051919050565b5f6020828403121561509e575f80fd5b8135611ab281614c19565b5f808335601e198436030181126150be575f80fd5b8301803591506001600160401b038211156150d7575f80fd5b6020019150600581901b3603821315613936575f80fd5b634e487b7160e01b5f52603260045260245ffd5b5f8235605e19833603018112615116575f80fd5b9190910192915050565b5f82601f83011261512f575f80fd5b813561513d61500a82614fc7565b818152846020838601011115615151575f80fd5b816020850160208301375f918101602001919091529392505050565b5f6060823603121561517d575f80fd5b615185614f4d565b82356001600160401b038082111561519b575f80fd5b6151a736838701615120565b835260208501359150808211156151bc575f80fd5b506151c936828601615120565b60208301525060408301356151dd81614c7b565b604082015292915050565b5f8251615116818460208701614d8b565b601f8211156107a957805f5260205f20601f840160051c8101602085101561521e5750805b601f840160051c820191505b818110156111b2575f815560010161522a565b81516001600160401b0381111561525657615256614f39565b61526a816152648454614ef9565b846151f9565b602080601f83116001811461529d575f84156152865750858301515b5f19600386901b1c1916600185901b178555611304565b5f85815260208120601f198616915b828110156152cb578886015182559484019460019091019084016152ac565b50858210156152e857878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b038181168382160190808211156126c5576126c56152f8565b5f63ffffffff808316818103615344576153446152f8565b6001019392505050565b6001600160401b03818116838216028082169190828114615371576153716152f8565b505092915050565b5f808335601e1984360301811261538e575f80fd5b83016020810192503590506001600160401b038111156153ac575f80fd5b803603821315613936575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208086019550808560051b830101845f5b8781101561549757848303601f19018952813536889003605e1901811261541e575f80fd5b8701606061542c8280615379565b82875261543c83880182846153ba565b9250505061544c86830183615379565b8683038888015261545e8382846153ba565b92505050604080830135925061547383614c7b565b6001600160401b0392909216949091019390935297830197908301906001016153f9565b5090979650505050505050565b6020815281356020820152602082013560408201525f60408301356154c881614c19565b6001600160a01b031660608381019190915283013536849003601e190181126154ef575f80fd5b83016020810190356001600160401b0381111561550a575f80fd5b8060051b360382131561551b575f80fd5b60808085015261552f60a0850182846153e2565b95945050505050565b5f8261555257634e487b7160e01b5f52601260045260245ffd5b500490565b602081016004831061556b5761556b614d63565b91905290565b80820281158282048414176111a0576111a06152f8565b5f8060408385031215615599575f80fd5b825191506020830151614ccd81614b21565b5f60208083525f84546155bd81614ef9565b806020870152604060018084165f81146155de57600181146155fa57615627565b60ff19851660408a0152604084151560051b8a01019550615627565b895f5260205f205f5b8581101561561e5781548b8201860152908301908801615603565b8a016040019650505b509398975050505050505050565b5f805f60608486031215615647575f80fd5b83519250602084015161565981614c7b565b6040850151909250614d5881614c7b565b818382375f9101908152919050565b6001600160401b038281168282160390808211156126c5576126c56152f8565b5f80604083850312156156aa575f80fd5b82516001600160401b03808211156156c0575f80fd5b90840190606082870312156156d3575f80fd5b6156db614f4d565b8251815260208301516156ed81614c19565b6020820152604083015182811115615703575f80fd5b61570f88828601614fed565b6040830152508094505050506020830151614ccd81614b21565b5f806040838503121561573a575f80fd5b825191506020830151614ccd81614c7b565b808201808211156111a0576111a06152f8565b5f808335601e19843603018112615774575f80fd5b8301803591506001600160401b0382111561578d575f80fd5b602001915036819003821315613936575f80fd5b5f8235603e19833603018112615116575f80fd5b5f604082360312156157c5575f80fd5b6157cd614f75565b6157d683614b2e565b81526020808401356001600160401b03808211156157f2575f80fd5b9085019036601f830112615804575f80fd5b81358181111561581657615816614f39565b8060051b9150615827848301614f97565b8181529183018401918481019036841115615840575f80fd5b938501935b8385101561586a578435925061585a83614c19565b8282529385019390850190615845565b94860194909452509295945050505050565b5f6020828403121561588c575f80fd5b611ab282614cf1565b5f602082840312156158a5575f80fd5b813560ff81168114611ab2575f80fd5b818103818111156111a0576111a06152f8565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b808310156159235784516001600160a01b031682529383019360019290920191908301906158fa565b509695505050505050565b60208152815160208201525f602083015160e06040840152615954610100840182614dad565b90506040840151601f19808584030160608601526159728383614dad565b92506001600160401b03606087015116608086015260808601519150808584030160a08601526159a283836158c8565b925060a08601519150808584030160c0860152506159c082826158c8565b91505060c08401516159dd60e08501826001600160401b03169052565b509392505050565b5f80604083850312156159f6575f80fd5b8251915060208301516001600160401b03811115615a12575f80fd5b615a1e85828601614fed565b9150509250929050565b5f6001600160401b03808316818103615344576153446152f856fe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122032091ffc723a6dfec6806a0059316735961b7ecad27a6da2c86224c3a069939664736f6c63430008190033",
}

// NativeTokenStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenStakingManagerMetaData.ABI instead.
var NativeTokenStakingManagerABI = NativeTokenStakingManagerMetaData.ABI

// NativeTokenStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenStakingManagerMetaData.Bin instead.
var NativeTokenStakingManagerBin = NativeTokenStakingManagerMetaData.Bin

// DeployNativeTokenStakingManager deploys a new Ethereum contract, binding an instance of NativeTokenStakingManager to it.
func DeployNativeTokenStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *NativeTokenStakingManager, error) {
	parsed, err := NativeTokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	validatorMessagesAddr, _, _, _ := DeployValidatorMessages(auth, backend)
	NativeTokenStakingManagerBin = strings.ReplaceAll(NativeTokenStakingManagerBin, "__$fd0c147b4031eef6079b0498cbafa865f0$__", validatorMessagesAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenStakingManager{NativeTokenStakingManagerCaller: NativeTokenStakingManagerCaller{contract: contract}, NativeTokenStakingManagerTransactor: NativeTokenStakingManagerTransactor{contract: contract}, NativeTokenStakingManagerFilterer: NativeTokenStakingManagerFilterer{contract: contract}}, nil
}

// NativeTokenStakingManager is an auto generated Go binding around an Ethereum contract.
type NativeTokenStakingManager struct {
	NativeTokenStakingManagerCaller     // Read-only binding to the contract
	NativeTokenStakingManagerTransactor // Write-only binding to the contract
	NativeTokenStakingManagerFilterer   // Log filterer for contract events
}

// NativeTokenStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenStakingManagerSession struct {
	Contract     *NativeTokenStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NativeTokenStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenStakingManagerCallerSession struct {
	Contract *NativeTokenStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// NativeTokenStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenStakingManagerTransactorSession struct {
	Contract     *NativeTokenStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// NativeTokenStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenStakingManagerRaw struct {
	Contract *NativeTokenStakingManager // Generic contract binding to access the raw methods on
}

// NativeTokenStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerCallerRaw struct {
	Contract *NativeTokenStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerTransactorRaw struct {
	Contract *NativeTokenStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenStakingManager creates a new instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManager(address common.Address, backend bind.ContractBackend) (*NativeTokenStakingManager, error) {
	contract, err := bindNativeTokenStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManager{NativeTokenStakingManagerCaller: NativeTokenStakingManagerCaller{contract: contract}, NativeTokenStakingManagerTransactor: NativeTokenStakingManagerTransactor{contract: contract}, NativeTokenStakingManagerFilterer: NativeTokenStakingManagerFilterer{contract: contract}}, nil
}

// NewNativeTokenStakingManagerCaller creates a new read-only instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenStakingManagerCaller, error) {
	contract, err := bindNativeTokenStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerCaller{contract: contract}, nil
}

// NewNativeTokenStakingManagerTransactor creates a new write-only instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenStakingManagerTransactor, error) {
	contract, err := bindNativeTokenStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerTransactor{contract: contract}, nil
}

// NewNativeTokenStakingManagerFilterer creates a new log filterer instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenStakingManagerFilterer, error) {
	contract, err := bindNativeTokenStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerFilterer{contract: contract}, nil
}

// bindNativeTokenStakingManager binds a generic wrapper to an already deployed contract.
func bindNativeTokenStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) ADDRESSLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "ADDRESS_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ADDRESSLENGTH() (uint32, error) {
	return _NativeTokenStakingManager.Contract.ADDRESSLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) ADDRESSLENGTH() (uint32, error) {
	return _NativeTokenStakingManager.Contract.ADDRESSLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _NativeTokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_NativeTokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _NativeTokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_NativeTokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) BLSPUBLICKEYLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "BLS_PUBLIC_KEY_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _NativeTokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _NativeTokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMCHURNPERCENTAGELIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_CHURN_PERCENTAGE_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMREGISTRATIONEXPIRYLENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_REGISTRATION_EXPIRY_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.NATIVEMINTER(&_NativeTokenStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.NATIVEMINTER(&_NativeTokenStakingManager.CallOpts)
}

// NODEIDLENGTH is a free data retrieval call binding the contract method 0x63e2ca97.
//
// Solidity: function NODE_ID_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) NODEIDLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "NODE_ID_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NODEIDLENGTH is a free data retrieval call binding the contract method 0x63e2ca97.
//
// Solidity: function NODE_ID_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) NODEIDLENGTH() (uint32, error) {
	return _NativeTokenStakingManager.Contract.NODEIDLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// NODEIDLENGTH is a free data retrieval call binding the contract method 0x63e2ca97.
//
// Solidity: function NODE_ID_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) NODEIDLENGTH() (uint32, error) {
	return _NativeTokenStakingManager.Contract.NODEIDLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) POSVALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "POS_VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) PCHAINBLOCKCHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "P_CHAIN_BLOCKCHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_NativeTokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_NativeTokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) VALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.WARPMESSENGER(&_NativeTokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.WARPMESSENGER(&_NativeTokenStakingManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _NativeTokenStakingManager.Contract.GetValidator(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _NativeTokenStakingManager.Contract.GetValidator(&_NativeTokenStakingManager.CallOpts, validationID)
}

// L1TotalWeight is a free data retrieval call binding the contract method 0xbb0b1938.
//
// Solidity: function l1TotalWeight() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) L1TotalWeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "l1TotalWeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// L1TotalWeight is a free data retrieval call binding the contract method 0xbb0b1938.
//
// Solidity: function l1TotalWeight() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) L1TotalWeight() (uint64, error) {
	return _NativeTokenStakingManager.Contract.L1TotalWeight(&_NativeTokenStakingManager.CallOpts)
}

// L1TotalWeight is a free data retrieval call binding the contract method 0xbb0b1938.
//
// Solidity: function l1TotalWeight() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) L1TotalWeight() (uint64, error) {
	return _NativeTokenStakingManager.Contract.L1TotalWeight(&_NativeTokenStakingManager.CallOpts)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) RegisteredValidators(opts *bind.CallOpts, nodeID []byte) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "registeredValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.RegisteredValidators(&_NativeTokenStakingManager.CallOpts, nodeID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.RegisteredValidators(&_NativeTokenStakingManager.CallOpts, nodeID)
}

// SubnetID is a free data retrieval call binding the contract method 0x5dc1f535.
//
// Solidity: function subnetID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) SubnetID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "subnetID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SubnetID is a free data retrieval call binding the contract method 0x5dc1f535.
//
// Solidity: function subnetID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) SubnetID() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.SubnetID(&_NativeTokenStakingManager.CallOpts)
}

// SubnetID is a free data retrieval call binding the contract method 0x5dc1f535.
//
// Solidity: function subnetID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) SubnetID() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.SubnetID(&_NativeTokenStakingManager.CallOpts)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenStakingManager.Contract.WeightToValue(&_NativeTokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenStakingManager.Contract.WeightToValue(&_NativeTokenStakingManager.CallOpts, weight)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ChangeDelegatorRewardRecipient(opts *bind.TransactOpts, delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "changeDelegatorRewardRecipient", delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_NativeTokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_NativeTokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ChangeValidatorRewardRecipient(opts *bind.TransactOpts, validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "changeValidatorRewardRecipient", validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_NativeTokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_NativeTokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimDelegationFees(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimDelegationFees(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRemoval(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRemoval(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32, uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteValidatorWeightUpdate(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeValidatorWeightUpdate", messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32, uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteValidatorWeightUpdate(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorWeightUpdate(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32, uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteValidatorWeightUpdate(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorWeightUpdate(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitializeEndDelegation0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation0", delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndDelegation0(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndDelegation0(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitializeEndValidation0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation0", validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndValidation0(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndValidation0(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x0ba512d1.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initialize", settings)
}

// Initialize is a paid mutator transaction binding the contract method 0x0ba512d1.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) Initialize(settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.Initialize(&_NativeTokenStakingManager.TransactOpts, settings)
}

// Initialize is a paid mutator transaction binding the contract method 0x0ba512d1.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) Initialize(settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.Initialize(&_NativeTokenStakingManager.TransactOpts, settings)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeDelegatorRegistration", validationID)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeDelegatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeDelegatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeEndDelegation0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeEndDelegation0", delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndDelegation0(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndDelegation0(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeEndValidation0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeEndValidation0", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndValidation0(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndValidation0(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeValidatorRegistration", registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeValidatorSet", conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorSet(&_NativeTokenStakingManager.TransactOpts, conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorSet(&_NativeTokenStakingManager.TransactOpts, conversionData, messageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendEndValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendEndValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendRegisterValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendRegisterValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendUpdateDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendUpdateDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.SubmitUptimeProof(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.SubmitUptimeProof(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// NativeTokenStakingManagerCompletedValidatorRegistrationIterator is returned from FilterCompletedValidatorRegistration and is used to iterate over the raw logs and unpacked data for CompletedValidatorRegistration events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedValidatorRegistrationIterator struct {
	Event *NativeTokenStakingManagerCompletedValidatorRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerCompletedValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerCompletedValidatorRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerCompletedValidatorRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerCompletedValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerCompletedValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerCompletedValidatorRegistration represents a CompletedValidatorRegistration event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedValidatorRegistration struct {
	ValidationID [32]byte
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedValidatorRegistration is a free log retrieval operation binding the contract event 0x967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06.
//
// Solidity: event CompletedValidatorRegistration(bytes32 indexed validationID, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterCompletedValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerCompletedValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "CompletedValidatorRegistration", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerCompletedValidatorRegistrationIterator{contract: _NativeTokenStakingManager.contract, event: "CompletedValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchCompletedValidatorRegistration is a free log subscription operation binding the contract event 0x967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06.
//
// Solidity: event CompletedValidatorRegistration(bytes32 indexed validationID, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchCompletedValidatorRegistration(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerCompletedValidatorRegistration, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "CompletedValidatorRegistration", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerCompletedValidatorRegistration)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedValidatorRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCompletedValidatorRegistration is a log parse operation binding the contract event 0x967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06.
//
// Solidity: event CompletedValidatorRegistration(bytes32 indexed validationID, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseCompletedValidatorRegistration(log types.Log) (*NativeTokenStakingManagerCompletedValidatorRegistration, error) {
	event := new(NativeTokenStakingManagerCompletedValidatorRegistration)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerCompletedValidatorRemovalIterator is returned from FilterCompletedValidatorRemoval and is used to iterate over the raw logs and unpacked data for CompletedValidatorRemoval events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedValidatorRemovalIterator struct {
	Event *NativeTokenStakingManagerCompletedValidatorRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerCompletedValidatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerCompletedValidatorRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerCompletedValidatorRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerCompletedValidatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerCompletedValidatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerCompletedValidatorRemoval represents a CompletedValidatorRemoval event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedValidatorRemoval struct {
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedValidatorRemoval is a free log retrieval operation binding the contract event 0xafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c.
//
// Solidity: event CompletedValidatorRemoval(bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterCompletedValidatorRemoval(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerCompletedValidatorRemovalIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "CompletedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerCompletedValidatorRemovalIterator{contract: _NativeTokenStakingManager.contract, event: "CompletedValidatorRemoval", logs: logs, sub: sub}, nil
}

// WatchCompletedValidatorRemoval is a free log subscription operation binding the contract event 0xafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c.
//
// Solidity: event CompletedValidatorRemoval(bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchCompletedValidatorRemoval(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerCompletedValidatorRemoval, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "CompletedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerCompletedValidatorRemoval)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedValidatorRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCompletedValidatorRemoval is a log parse operation binding the contract event 0xafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c.
//
// Solidity: event CompletedValidatorRemoval(bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseCompletedValidatorRemoval(log types.Log) (*NativeTokenStakingManagerCompletedValidatorRemoval, error) {
	event := new(NativeTokenStakingManagerCompletedValidatorRemoval)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedValidatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerCompletedValidatorWeightUpdateIterator is returned from FilterCompletedValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for CompletedValidatorWeightUpdate events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedValidatorWeightUpdateIterator struct {
	Event *NativeTokenStakingManagerCompletedValidatorWeightUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerCompletedValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerCompletedValidatorWeightUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerCompletedValidatorWeightUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerCompletedValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerCompletedValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerCompletedValidatorWeightUpdate represents a CompletedValidatorWeightUpdate event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedValidatorWeightUpdate struct {
	ValidationID [32]byte
	Nonce        uint64
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedValidatorWeightUpdate is a free log retrieval operation binding the contract event 0xc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff9.
//
// Solidity: event CompletedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterCompletedValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerCompletedValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "CompletedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerCompletedValidatorWeightUpdateIterator{contract: _NativeTokenStakingManager.contract, event: "CompletedValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchCompletedValidatorWeightUpdate is a free log subscription operation binding the contract event 0xc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff9.
//
// Solidity: event CompletedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchCompletedValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerCompletedValidatorWeightUpdate, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "CompletedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerCompletedValidatorWeightUpdate)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedValidatorWeightUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCompletedValidatorWeightUpdate is a log parse operation binding the contract event 0xc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff9.
//
// Solidity: event CompletedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseCompletedValidatorWeightUpdate(log types.Log) (*NativeTokenStakingManagerCompletedValidatorWeightUpdate, error) {
	event := new(NativeTokenStakingManagerCompletedValidatorWeightUpdate)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegationEndedIterator struct {
	Event *NativeTokenStakingManagerDelegationEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegationEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerDelegationEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegationEnded represents a DelegationEnded event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegationEndedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegationEnded)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegationEnded is a log parse operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegationEnded(log types.Log) (*NativeTokenStakingManagerDelegationEnded, error) {
	event := new(NativeTokenStakingManagerDelegationEnded)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorAddedIterator struct {
	Event *NativeTokenStakingManagerDelegatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerDelegatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorAdded represents a DelegatorAdded event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorAdded struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatorAdded is a free log retrieval operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*NativeTokenStakingManagerDelegatorAddedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorAddedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorAdded)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorAdded is a log parse operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorAdded(log types.Log) (*NativeTokenStakingManagerDelegatorAdded, error) {
	event := new(NativeTokenStakingManagerDelegatorAdded)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRegisteredIterator struct {
	Event *NativeTokenStakingManagerDelegatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerDelegatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorRegistered represents a DelegatorRegistered event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorRegisteredIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorRegistered)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorRegistered is a log parse operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorRegistered(log types.Log) (*NativeTokenStakingManagerDelegatorRegistered, error) {
	event := new(NativeTokenStakingManagerDelegatorRegistered)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRemovalInitializedIterator struct {
	Event *NativeTokenStakingManagerDelegatorRemovalInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorRemovalInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerDelegatorRemovalInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorRemovalInitializedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorRemovalInitialized)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorRemovalInitialized is a log parse operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*NativeTokenStakingManagerDelegatorRemovalInitialized, error) {
	event := new(NativeTokenStakingManagerDelegatorRemovalInitialized)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitializedIterator struct {
	Event *NativeTokenStakingManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitialized represents a Initialized event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenStakingManagerInitializedIterator, error) {

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitializedIterator{contract: _NativeTokenStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitialized)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitialized(log types.Log) (*NativeTokenStakingManagerInitialized, error) {
	event := new(NativeTokenStakingManagerInitialized)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitiatedValidatorRegistrationIterator is returned from FilterInitiatedValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedValidatorRegistration events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedValidatorRegistrationIterator struct {
	Event *NativeTokenStakingManagerInitiatedValidatorRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerInitiatedValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitiatedValidatorRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerInitiatedValidatorRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerInitiatedValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitiatedValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitiatedValidatorRegistration represents a InitiatedValidatorRegistration event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedValidatorRegistration struct {
	ValidationID          [32]byte
	NodeID                [20]byte
	RegistrationMessageID [32]byte
	RegistrationExpiry    uint64
	Weight                uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterInitiatedValidatorRegistration is a free log retrieval operation binding the contract event 0x5881be437bdcb008bfa5f20e32d3e335ccf8ab90ef2818852a251625260af35d.
//
// Solidity: event InitiatedValidatorRegistration(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 registrationMessageID, uint64 registrationExpiry, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitiatedValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][20]byte) (*NativeTokenStakingManagerInitiatedValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "InitiatedValidatorRegistration", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitiatedValidatorRegistrationIterator{contract: _NativeTokenStakingManager.contract, event: "InitiatedValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedValidatorRegistration is a free log subscription operation binding the contract event 0x5881be437bdcb008bfa5f20e32d3e335ccf8ab90ef2818852a251625260af35d.
//
// Solidity: event InitiatedValidatorRegistration(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 registrationMessageID, uint64 registrationExpiry, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitiatedValidatorRegistration(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitiatedValidatorRegistration, validationID [][32]byte, nodeID [][20]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "InitiatedValidatorRegistration", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitiatedValidatorRegistration)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiatedValidatorRegistration is a log parse operation binding the contract event 0x5881be437bdcb008bfa5f20e32d3e335ccf8ab90ef2818852a251625260af35d.
//
// Solidity: event InitiatedValidatorRegistration(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 registrationMessageID, uint64 registrationExpiry, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitiatedValidatorRegistration(log types.Log) (*NativeTokenStakingManagerInitiatedValidatorRegistration, error) {
	event := new(NativeTokenStakingManagerInitiatedValidatorRegistration)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitiatedValidatorRemovalIterator is returned from FilterInitiatedValidatorRemoval and is used to iterate over the raw logs and unpacked data for InitiatedValidatorRemoval events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedValidatorRemovalIterator struct {
	Event *NativeTokenStakingManagerInitiatedValidatorRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerInitiatedValidatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitiatedValidatorRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerInitiatedValidatorRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerInitiatedValidatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitiatedValidatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitiatedValidatorRemoval represents a InitiatedValidatorRemoval event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedValidatorRemoval struct {
	ValidationID             [32]byte
	ValidatorWeightMessageID [32]byte
	Weight                   uint64
	EndTime                  uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterInitiatedValidatorRemoval is a free log retrieval operation binding the contract event 0xbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273.
//
// Solidity: event InitiatedValidatorRemoval(bytes32 indexed validationID, bytes32 validatorWeightMessageID, uint64 weight, uint64 endTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitiatedValidatorRemoval(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerInitiatedValidatorRemovalIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "InitiatedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitiatedValidatorRemovalIterator{contract: _NativeTokenStakingManager.contract, event: "InitiatedValidatorRemoval", logs: logs, sub: sub}, nil
}

// WatchInitiatedValidatorRemoval is a free log subscription operation binding the contract event 0xbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273.
//
// Solidity: event InitiatedValidatorRemoval(bytes32 indexed validationID, bytes32 validatorWeightMessageID, uint64 weight, uint64 endTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitiatedValidatorRemoval(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitiatedValidatorRemoval, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "InitiatedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitiatedValidatorRemoval)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiatedValidatorRemoval is a log parse operation binding the contract event 0xbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273.
//
// Solidity: event InitiatedValidatorRemoval(bytes32 indexed validationID, bytes32 validatorWeightMessageID, uint64 weight, uint64 endTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitiatedValidatorRemoval(log types.Log) (*NativeTokenStakingManagerInitiatedValidatorRemoval, error) {
	event := new(NativeTokenStakingManagerInitiatedValidatorRemoval)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitiatedValidatorWeightUpdateIterator is returned from FilterInitiatedValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for InitiatedValidatorWeightUpdate events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedValidatorWeightUpdateIterator struct {
	Event *NativeTokenStakingManagerInitiatedValidatorWeightUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerInitiatedValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitiatedValidatorWeightUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerInitiatedValidatorWeightUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerInitiatedValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitiatedValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitiatedValidatorWeightUpdate represents a InitiatedValidatorWeightUpdate event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedValidatorWeightUpdate struct {
	ValidationID          [32]byte
	Nonce                 uint64
	WeightUpdateMessageID [32]byte
	Weight                uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterInitiatedValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d3.
//
// Solidity: event InitiatedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, bytes32 weightUpdateMessageID, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitiatedValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerInitiatedValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "InitiatedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitiatedValidatorWeightUpdateIterator{contract: _NativeTokenStakingManager.contract, event: "InitiatedValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchInitiatedValidatorWeightUpdate is a free log subscription operation binding the contract event 0x6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d3.
//
// Solidity: event InitiatedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, bytes32 weightUpdateMessageID, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitiatedValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitiatedValidatorWeightUpdate, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "InitiatedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitiatedValidatorWeightUpdate)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorWeightUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiatedValidatorWeightUpdate is a log parse operation binding the contract event 0x6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d3.
//
// Solidity: event InitiatedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, bytes32 weightUpdateMessageID, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitiatedValidatorWeightUpdate(log types.Log) (*NativeTokenStakingManagerInitiatedValidatorWeightUpdate, error) {
	event := new(NativeTokenStakingManagerInitiatedValidatorWeightUpdate)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerRegisteredInitialValidatorIterator is returned from FilterRegisteredInitialValidator and is used to iterate over the raw logs and unpacked data for RegisteredInitialValidator events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerRegisteredInitialValidatorIterator struct {
	Event *NativeTokenStakingManagerRegisteredInitialValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerRegisteredInitialValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerRegisteredInitialValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerRegisteredInitialValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerRegisteredInitialValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerRegisteredInitialValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerRegisteredInitialValidator represents a RegisteredInitialValidator event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerRegisteredInitialValidator struct {
	ValidationID [32]byte
	NodeID       [20]byte
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredInitialValidator is a free log retrieval operation binding the contract event 0x9d9c026e2cadfec89cccc2cd72705360eca1beba24774f3363f4bb33faabc7d7.
//
// Solidity: event RegisteredInitialValidator(bytes32 indexed validationID, bytes20 indexed nodeID, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterRegisteredInitialValidator(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][20]byte) (*NativeTokenStakingManagerRegisteredInitialValidatorIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "RegisteredInitialValidator", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerRegisteredInitialValidatorIterator{contract: _NativeTokenStakingManager.contract, event: "RegisteredInitialValidator", logs: logs, sub: sub}, nil
}

// WatchRegisteredInitialValidator is a free log subscription operation binding the contract event 0x9d9c026e2cadfec89cccc2cd72705360eca1beba24774f3363f4bb33faabc7d7.
//
// Solidity: event RegisteredInitialValidator(bytes32 indexed validationID, bytes20 indexed nodeID, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchRegisteredInitialValidator(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerRegisteredInitialValidator, validationID [][32]byte, nodeID [][20]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "RegisteredInitialValidator", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerRegisteredInitialValidator)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "RegisteredInitialValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisteredInitialValidator is a log parse operation binding the contract event 0x9d9c026e2cadfec89cccc2cd72705360eca1beba24774f3363f4bb33faabc7d7.
//
// Solidity: event RegisteredInitialValidator(bytes32 indexed validationID, bytes20 indexed nodeID, uint64 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseRegisteredInitialValidator(log types.Log) (*NativeTokenStakingManagerRegisteredInitialValidator, error) {
	event := new(NativeTokenStakingManagerRegisteredInitialValidator)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "RegisteredInitialValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerUptimeUpdatedIterator struct {
	Event *NativeTokenStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerUptimeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerUptimeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerUptimeUpdatedIterator{contract: _NativeTokenStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerUptimeUpdated)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUptimeUpdated is a log parse operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*NativeTokenStakingManagerUptimeUpdated, error) {
	event := new(NativeTokenStakingManagerUptimeUpdated)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMessagesMetaData contains all meta data concerning the ValidatorMessages contract.
var ValidatorMessagesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidBLSPublicKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"InvalidCodecID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"actual\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expected\",\"type\":\"uint32\"}],\"name\":\"InvalidMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageType\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"}],\"name\":\"packConversionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"packL1ValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"packL1ValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"validationPeriod\",\"type\":\"tuple\"}],\"name\":\"packRegisterL1ValidatorMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conversionID\",\"type\":\"bytes32\"}],\"name\":\"packSubnetToL1ConversionMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"packValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackL1ValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackL1ValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackRegisterL1ValidatorMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackSubnetToL1ConversionMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x61217b610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b1575f3560e01c8063854a893f11610079578063854a893f146101b257806387418b8e1461020f5780639b83546514610222578063a699c13514610242578063e1d68f3014610255578063eb97ce5114610268575f80fd5b8063021de88f146100b5578063088c2463146100e25780634d8478841461011257806350782b0f146101335780637f7c427a1461016b575b5f80fd5b6100c86100c33660046118a9565b610289565b604080519283529015156020830152015b60405180910390f35b6100f56100f03660046118a9565b61044a565b604080519283526001600160401b039091166020830152016100d9565b6101256101203660046118a9565b61063b565b6040519081526020016100d9565b6101466101413660046118a9565b6107c8565b604080519384526001600160401b0392831660208501529116908201526060016100d9565b6101a56101793660046118e2565b604080515f60208201819052602282015260268082019390935281518082039093018352604601905290565b6040516100d99190611946565b6101a56101c036600461197a565b604080515f6020820152600360e01b602282015260268101949094526001600160c01b031960c093841b811660468601529190921b16604e830152805180830360360181526056909201905290565b6101a561021d3660046119eb565b610a1e565b6102356102303660046118a9565b610b60565b6040516100d99190611bb4565b6101a5610250366004611c6b565b6114ab565b6101a5610263366004611c9d565b6114ef565b61027b610276366004611d80565b611525565b6040516100d9929190611e7c565b5f8082516027146102c457825160405163cc92daa160e01b815263ffffffff9091166004820152602760248201526044015b60405180910390fd5b5f805b6002811015610313576102db816001611ea8565b6102e6906008611ebb565b61ffff168582815181106102fc576102fc611ed2565b016020015160f81c901b91909117906001016102c7565b5061ffff81161561033d5760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561039857610354816003611ea8565b61035f906008611ebb565b63ffffffff1686610371836002611ee6565b8151811061038157610381611ed2565b016020015160f81c901b9190911790600101610340565b5063ffffffff81166002146103c057604051635b60892f60e01b815260040160405180910390fd5b5f805b6020811015610415576103d781601f611ea8565b6103e2906008611ebb565b876103ee836006611ee6565b815181106103fe576103fe611ed2565b016020015160f81c901b91909117906001016103c3565b505f8660268151811061042a5761042a611ed2565b016020015191976001600160f81b03199092161515965090945050505050565b5f808251602e1461048057825160405163cc92daa160e01b815263ffffffff9091166004820152602e60248201526044016102bb565b5f805b60028110156104cf57610497816001611ea8565b6104a2906008611ebb565b61ffff168582815181106104b8576104b8611ed2565b016020015160f81c901b9190911790600101610483565b5061ffff8116156104f95760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561055457610510816003611ea8565b61051b906008611ebb565b63ffffffff168661052d836002611ee6565b8151811061053d5761053d611ed2565b016020015160f81c901b91909117906001016104fc565b5063ffffffff81161561057a57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156105cf5761059181601f611ea8565b61059c906008611ebb565b876105a8836006611ee6565b815181106105b8576105b8611ed2565b016020015160f81c901b919091179060010161057d565b505f805b600881101561062e576105e7816007611ea8565b6105f2906008611ebb565b6001600160401b031688610607836026611ee6565b8151811061061757610617611ed2565b016020015160f81c901b91909117906001016105d3565b5090969095509350505050565b5f815160261461067057815160405163cc92daa160e01b815263ffffffff9091166004820152602660248201526044016102bb565b5f805b60028110156106bf57610687816001611ea8565b610692906008611ebb565b61ffff168482815181106106a8576106a8611ed2565b016020015160f81c901b9190911790600101610673565b5061ffff8116156106e95760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561074457610700816003611ea8565b61070b906008611ebb565b63ffffffff168561071d836002611ee6565b8151811061072d5761072d611ed2565b016020015160f81c901b91909117906001016106ec565b5063ffffffff81161561076a57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156107bf5761078181601f611ea8565b61078c906008611ebb565b86610798836006611ee6565b815181106107a8576107a8611ed2565b016020015160f81c901b919091179060010161076d565b50949350505050565b5f805f83516036146107ff57835160405163cc92daa160e01b815263ffffffff9091166004820152603660248201526044016102bb565b5f805b600281101561084e57610816816001611ea8565b610821906008611ebb565b61ffff1686828151811061083757610837611ed2565b016020015160f81c901b9190911790600101610802565b5061ffff8116156108785760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b60048110156108d35761088f816003611ea8565b61089a906008611ebb565b63ffffffff16876108ac836002611ee6565b815181106108bc576108bc611ed2565b016020015160f81c901b919091179060010161087b565b5063ffffffff81166003146108fb57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156109505761091281601f611ea8565b61091d906008611ebb565b88610929836006611ee6565b8151811061093957610939611ed2565b016020015160f81c901b91909117906001016108fe565b505f805b60088110156109af57610968816007611ea8565b610973906008611ebb565b6001600160401b031689610988836026611ee6565b8151811061099857610998611ed2565b016020015160f81c901b9190911790600101610954565b505f805b6008811015610a0e576109c7816007611ea8565b6109d2906008611ebb565b6001600160401b03168a6109e783602e611ee6565b815181106109f7576109f7611ed2565b016020015160f81c901b91909117906001016109b3565b5091989097509095509350505050565b80516020808301516040808501516060868101515192515f95810186905260228101969096526042860193909352600560e21b60628601526bffffffffffffffffffffffff1990831b16606685015260e01b6001600160e01b031916607a84015291607e0160405160208183030381529060405290505f5b836060015151811015610b59578184606001518281518110610aba57610aba611ed2565b60200260200101515f01515185606001518381518110610adc57610adc611ed2565b60200260200101515f015186606001518481518110610afd57610afd611ed2565b60200260200101516020015187606001518581518110610b1f57610b1f611ed2565b602002602001015160400151604051602001610b3f959493929190611ef9565b60408051601f198184030181529190529150600101610a96565b5092915050565b610b68611712565b5f610b71611712565b5f805b6002811015610bcf57610b88816001611ea8565b610b93906008611ebb565b61ffff1686610ba863ffffffff871684611ee6565b81518110610bb857610bb8611ed2565b016020015160f81c901b9190911790600101610b74565b5061ffff811615610bf95760405163407b587360e01b815261ffff821660048201526024016102bb565b610c04600284611f72565b9250505f805b6004811015610c6957610c1e816003611ea8565b610c29906008611ebb565b63ffffffff16868563ffffffff1683610c429190611ee6565b81518110610c5257610c52611ed2565b016020015160f81c901b9190911790600101610c0a565b5063ffffffff8116600114610c9157604051635b60892f60e01b815260040160405180910390fd5b610c9c600484611f72565b9250505f805b6020811015610cf957610cb681601f611ea8565b610cc1906008611ebb565b86610cd263ffffffff871684611ee6565b81518110610ce257610ce2611ed2565b016020015160f81c901b9190911790600101610ca2565b50808252610d08602084611f72565b9250505f805b6004811015610d6d57610d22816003611ea8565b610d2d906008611ebb565b63ffffffff16868563ffffffff1683610d469190611ee6565b81518110610d5657610d56611ed2565b016020015160f81c901b9190911790600101610d0e565b50610d79600484611f72565b92505f8163ffffffff166001600160401b03811115610d9a57610d9a61176c565b6040519080825280601f01601f191660200182016040528015610dc4576020820181803683370190505b5090505f5b8263ffffffff16811015610e335786610de863ffffffff871683611ee6565b81518110610df857610df8611ed2565b602001015160f81c60f81b828281518110610e1557610e15611ed2565b60200101906001600160f81b03191690815f1a905350600101610dc9565b5060208301819052610e458285611f72565b604080516030808252606082019092529195505f92506020820181803683370190505090505f5b6030811015610ed15786610e8663ffffffff871683611ee6565b81518110610e9657610e96611ed2565b602001015160f81c60f81b828281518110610eb357610eb3611ed2565b60200101906001600160f81b03191690815f1a905350600101610e6c565b5060408301819052610ee4603085611f72565b9350505f805b6008811015610f4a57610efe816007611ea8565b610f09906008611ebb565b6001600160401b031687610f2363ffffffff881684611ee6565b81518110610f3357610f33611ed2565b016020015160f81c901b9190911790600101610eea565b506001600160401b0381166060840152610f65600885611f72565b9350505f805f5b6004811015610fcb57610f80816003611ea8565b610f8b906008611ebb565b63ffffffff16888763ffffffff1683610fa49190611ee6565b81518110610fb457610fb4611ed2565b016020015160f81c901b9190911790600101610f6c565b50610fd7600486611f72565b94505f5b600481101561103a57610fef816003611ea8565b610ffa906008611ebb565b63ffffffff16888763ffffffff16836110139190611ee6565b8151811061102357611023611ed2565b016020015160f81c901b9290921791600101610fdb565b50611046600486611f72565b94505f8263ffffffff166001600160401b038111156110675761106761176c565b604051908082528060200260200182016040528015611090578160200160208202803683370190505b5090505f5b8363ffffffff16811015611178576040805160148082528183019092525f916020820181803683370190505090505f5b601481101561112a578a6110df63ffffffff8b1683611ee6565b815181106110ef576110ef611ed2565b602001015160f81c60f81b82828151811061110c5761110c611ed2565b60200101906001600160f81b03191690815f1a9053506001016110c5565b505f601482015190508084848151811061114657611146611ed2565b6001600160a01b039092166020928302919091019091015261116960148a611f72565b98505050806001019050611095565b506040805180820190915263ffffffff9092168252602082015260808401525f80805b60048110156111fa576111af816003611ea8565b6111ba906008611ebb565b63ffffffff16898863ffffffff16836111d39190611ee6565b815181106111e3576111e3611ed2565b016020015160f81c901b919091179060010161119b565b50611206600487611f72565b95505f5b60048110156112695761121e816003611ea8565b611229906008611ebb565b63ffffffff16898863ffffffff16836112429190611ee6565b8151811061125257611252611ed2565b016020015160f81c901b929092179160010161120a565b50611275600487611f72565b95505f8263ffffffff166001600160401b038111156112965761129661176c565b6040519080825280602002602001820160405280156112bf578160200160208202803683370190505b5090505f5b8363ffffffff168110156113a7576040805160148082528183019092525f916020820181803683370190505090505f5b6014811015611359578b61130e63ffffffff8c1683611ee6565b8151811061131e5761131e611ed2565b602001015160f81c60f81b82828151811061133b5761133b611ed2565b60200101906001600160f81b03191690815f1a9053506001016112f4565b505f601482015190508084848151811061137557611375611ed2565b6001600160a01b039092166020928302919091019091015261139860148b611f72565b995050508060010190506112c4565b506040805180820190915263ffffffff9092168252602082015260a08501525f6113d18284611f72565b6113dc906014611f8f565b6113e785607a611f72565b6113f19190611f72565b90508063ffffffff1688511461142d57875160405163cc92daa160e01b815263ffffffff918216600482015290821660248201526044016102bb565b5f805b600881101561149057611444816007611ea8565b61144f906008611ebb565b6001600160401b03168a61146963ffffffff8b1684611ee6565b8151811061147957611479611ed2565b016020015160f81c901b9190911790600101611430565b506001600160401b031660c086015250929695505050505050565b6040515f6020820152600160e11b60228201526026810183905281151560f81b60468201526060906047015b60405160208183030381529060405290505b92915050565b6040515f602082018190526022820152602681018390526001600160c01b031960c083901b166046820152606090604e016114d7565b5f606082604001515160301461154e5760405163180ffa0d60e01b815260040160405180910390fd5b82516020808501518051604080880151606089015160808a01518051908701515193515f9861158f988a986001989297929690959094909390929101611fb7565b60405160208183030381529060405290505f5b84608001516020015151811015611601578185608001516020015182815181106115ce576115ce611ed2565b60200260200101516040516020016115e7929190612071565b60408051601f1981840301815291905291506001016115a2565b5060a08401518051602091820151516040516116219385939291016120a7565b60405160208183030381529060405290505f5b8460a00151602001515181101561169357818560a0015160200151828151811061166057611660611ed2565b6020026020010151604051602001611679929190612071565b60408051601f198184030181529190529150600101611634565b5060c08401516040516116aa9183916020016120e2565b60405160208183030381529060405290506002816040516116cb9190612113565b602060405180830381855afa1580156116e6573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190611709919061212e565b94909350915050565b6040805160e0810182525f808252606060208084018290528385018290528184018390528451808601865283815280820183905260808501528451808601909552918452908301529060a082019081525f60209091015290565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b03811182821017156117a2576117a261176c565b60405290565b604051606081016001600160401b03811182821017156117a2576117a261176c565b604080519081016001600160401b03811182821017156117a2576117a261176c565b60405160e081016001600160401b03811182821017156117a2576117a261176c565b604051601f8201601f191681016001600160401b03811182821017156118365761183661176c565b604052919050565b5f82601f83011261184d575f80fd5b81356001600160401b038111156118665761186661176c565b611879601f8201601f191660200161180e565b81815284602083860101111561188d575f80fd5b816020850160208301375f918101602001919091529392505050565b5f602082840312156118b9575f80fd5b81356001600160401b038111156118ce575f80fd5b6118da8482850161183e565b949350505050565b5f602082840312156118f2575f80fd5b5035919050565b5f5b838110156119135781810151838201526020016118fb565b50505f910152565b5f81518084526119328160208601602086016118f9565b601f01601f19169290920160200192915050565b602081525f611958602083018461191b565b9392505050565b80356001600160401b0381168114611975575f80fd5b919050565b5f805f6060848603121561198c575f80fd5b8335925061199c6020850161195f565b91506119aa6040850161195f565b90509250925092565b80356001600160a01b0381168114611975575f80fd5b5f6001600160401b038211156119e1576119e161176c565b5060051b60200190565b5f60208083850312156119fc575f80fd5b82356001600160401b0380821115611a12575f80fd5b9084019060808287031215611a25575f80fd5b611a2d611780565b823581528383013584820152611a45604084016119b3565b604082015260608084013583811115611a5c575f80fd5b80850194505087601f850112611a70575f80fd5b8335611a83611a7e826119c9565b61180e565b81815260059190911b8501860190868101908a831115611aa1575f80fd5b8787015b83811015611b3a57803587811115611abb575f80fd5b8801808d03601f1901861315611acf575f80fd5b611ad76117a8565b8a82013589811115611ae7575f80fd5b611af58f8d8386010161183e565b825250604082013589811115611b09575f80fd5b611b178f8d8386010161183e565b8c83015250611b2787830161195f565b6040820152845250918801918801611aa5565b506060850152509198975050505050505050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611ba95784516001600160a01b03168252938301936001929092019190830190611b80565b509695505050505050565b60208152815160208201525f602083015160e06040840152611bda61010084018261191b565b90506040840151601f1980858403016060860152611bf8838361191b565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152611c288383611b4e565b925060a08601519150808584030160c086015250611c468282611b4e565b91505060c0840151611c6360e08501826001600160401b03169052565b509392505050565b5f8060408385031215611c7c575f80fd5b8235915060208301358015158114611c92575f80fd5b809150509250929050565b5f8060408385031215611cae575f80fd5b82359150611cbe6020840161195f565b90509250929050565b5f60408284031215611cd7575f80fd5b611cdf6117ca565b9050813563ffffffff81168114611cf4575f80fd5b81526020828101356001600160401b03811115611d0f575f80fd5b8301601f81018513611d1f575f80fd5b8035611d2d611a7e826119c9565b81815260059190911b82018301908381019087831115611d4b575f80fd5b928401925b82841015611d7057611d61846119b3565b82529284019290840190611d50565b8085870152505050505092915050565b5f60208284031215611d90575f80fd5b81356001600160401b0380821115611da6575f80fd5b9083019060e08286031215611db9575f80fd5b611dc16117ec565b82358152602083013582811115611dd6575f80fd5b611de28782860161183e565b602083015250604083013582811115611df9575f80fd5b611e058782860161183e565b604083015250611e176060840161195f565b6060820152608083013582811115611e2d575f80fd5b611e3987828601611cc7565b60808301525060a083013582811115611e50575f80fd5b611e5c87828601611cc7565b60a083015250611e6e60c0840161195f565b60c082015295945050505050565b828152604060208201525f6118da604083018461191b565b634e487b7160e01b5f52601160045260245ffd5b818103818111156114e9576114e9611e94565b80820281158282048414176114e9576114e9611e94565b634e487b7160e01b5f52603260045260245ffd5b808201808211156114e9576114e9611e94565b5f8651611f0a818460208b016118f9565b60e087901b6001600160e01b0319169083019081528551611f32816004840160208a016118f9565b8551910190611f488160048401602089016118f9565b60c09490941b6001600160c01b031916600491909401908101939093525050600c01949350505050565b63ffffffff818116838216019080821115610b5957610b59611e94565b63ffffffff818116838216028082169190828114611faf57611faf611e94565b505092915050565b61ffff60f01b8a60f01b1681525f63ffffffff60e01b808b60e01b166002840152896006840152808960e01b166026840152508651611ffd81602a850160208b016118f9565b86519083019061201481602a840160208b016118f9565b60c087901b6001600160c01b031916602a9290910191820152612046603282018660e01b6001600160e01b0319169052565b61205f603682018560e01b6001600160e01b0319169052565b603a019b9a5050505050505050505050565b5f83516120828184602088016118f9565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b5f84516120b88184602089016118f9565b6001600160e01b031960e095861b8116919093019081529290931b16600482015260080192915050565b5f83516120f38184602088016118f9565b60c09390931b6001600160c01b0319169190920190815260080192915050565b5f82516121248184602087016118f9565b9190910192915050565b5f6020828403121561213e575f80fd5b505191905056fea2646970667358221220e6b22cfbb2dc3686c9e5d5c46b42e65620ca6777a1371e31c12b720ad3c60ea464736f6c63430008190033",
}

// ValidatorMessagesABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorMessagesMetaData.ABI instead.
var ValidatorMessagesABI = ValidatorMessagesMetaData.ABI

// ValidatorMessagesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorMessagesMetaData.Bin instead.
var ValidatorMessagesBin = ValidatorMessagesMetaData.Bin

// DeployValidatorMessages deploys a new Ethereum contract, binding an instance of ValidatorMessages to it.
func DeployValidatorMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorMessages, error) {
	parsed, err := ValidatorMessagesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorMessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorMessages{ValidatorMessagesCaller: ValidatorMessagesCaller{contract: contract}, ValidatorMessagesTransactor: ValidatorMessagesTransactor{contract: contract}, ValidatorMessagesFilterer: ValidatorMessagesFilterer{contract: contract}}, nil
}

// ValidatorMessages is an auto generated Go binding around an Ethereum contract.
type ValidatorMessages struct {
	ValidatorMessagesCaller     // Read-only binding to the contract
	ValidatorMessagesTransactor // Write-only binding to the contract
	ValidatorMessagesFilterer   // Log filterer for contract events
}

// ValidatorMessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorMessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorMessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorMessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorMessagesSession struct {
	Contract     *ValidatorMessages // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorMessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorMessagesCallerSession struct {
	Contract *ValidatorMessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorMessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorMessagesTransactorSession struct {
	Contract     *ValidatorMessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorMessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorMessagesRaw struct {
	Contract *ValidatorMessages // Generic contract binding to access the raw methods on
}

// ValidatorMessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorMessagesCallerRaw struct {
	Contract *ValidatorMessagesCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorMessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorMessagesTransactorRaw struct {
	Contract *ValidatorMessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorMessages creates a new instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessages(address common.Address, backend bind.ContractBackend) (*ValidatorMessages, error) {
	contract, err := bindValidatorMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessages{ValidatorMessagesCaller: ValidatorMessagesCaller{contract: contract}, ValidatorMessagesTransactor: ValidatorMessagesTransactor{contract: contract}, ValidatorMessagesFilterer: ValidatorMessagesFilterer{contract: contract}}, nil
}

// NewValidatorMessagesCaller creates a new read-only instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesCaller(address common.Address, caller bind.ContractCaller) (*ValidatorMessagesCaller, error) {
	contract, err := bindValidatorMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesCaller{contract: contract}, nil
}

// NewValidatorMessagesTransactor creates a new write-only instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorMessagesTransactor, error) {
	contract, err := bindValidatorMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesTransactor{contract: contract}, nil
}

// NewValidatorMessagesFilterer creates a new log filterer instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorMessagesFilterer, error) {
	contract, err := bindValidatorMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesFilterer{contract: contract}, nil
}

// bindValidatorMessages binds a generic wrapper to an already deployed contract.
func bindValidatorMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorMessagesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMessages *ValidatorMessagesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorMessages.Contract.ValidatorMessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMessages *ValidatorMessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.ValidatorMessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMessages *ValidatorMessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.ValidatorMessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMessages *ValidatorMessagesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorMessages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMessages *ValidatorMessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMessages *ValidatorMessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.contract.Transact(opts, method, params...)
}

// PackConversionData is a free data retrieval call binding the contract method 0x51f48008.
//
// Solidity: function packConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackConversionData(opts *bind.CallOpts, conversionData ConversionData) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packConversionData", conversionData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackConversionData is a free data retrieval call binding the contract method 0x51f48008.
//
// Solidity: function packConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackConversionData(conversionData ConversionData) ([]byte, error) {
	return _ValidatorMessages.Contract.PackConversionData(&_ValidatorMessages.CallOpts, conversionData)
}

// PackConversionData is a free data retrieval call binding the contract method 0x51f48008.
//
// Solidity: function packConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackConversionData(conversionData ConversionData) ([]byte, error) {
	return _ValidatorMessages.Contract.PackConversionData(&_ValidatorMessages.CallOpts, conversionData)
}

// PackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xa699c135.
//
// Solidity: function packL1ValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackL1ValidatorRegistrationMessage(opts *bind.CallOpts, validationID [32]byte, valid bool) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packL1ValidatorRegistrationMessage", validationID, valid)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xa699c135.
//
// Solidity: function packL1ValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackL1ValidatorRegistrationMessage(validationID [32]byte, valid bool) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, validationID, valid)
}

// PackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xa699c135.
//
// Solidity: function packL1ValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackL1ValidatorRegistrationMessage(validationID [32]byte, valid bool) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, validationID, valid)
}

// PackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x854a893f.
//
// Solidity: function packL1ValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackL1ValidatorWeightMessage(opts *bind.CallOpts, validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packL1ValidatorWeightMessage", validationID, nonce, weight)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x854a893f.
//
// Solidity: function packL1ValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackL1ValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, validationID, nonce, weight)
}

// PackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x854a893f.
//
// Solidity: function packL1ValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackL1ValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, validationID, nonce, weight)
}

// PackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0xe0d5478f.
//
// Solidity: function packRegisterL1ValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackRegisterL1ValidatorMessage(opts *bind.CallOpts, validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packRegisterL1ValidatorMessage", validationPeriod)

	if err != nil {
		return *new([32]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// PackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0xe0d5478f.
//
// Solidity: function packRegisterL1ValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackRegisterL1ValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _ValidatorMessages.Contract.PackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, validationPeriod)
}

// PackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0xe0d5478f.
//
// Solidity: function packRegisterL1ValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackRegisterL1ValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _ValidatorMessages.Contract.PackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, validationPeriod)
}

// PackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x7f7c427a.
//
// Solidity: function packSubnetToL1ConversionMessage(bytes32 conversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackSubnetToL1ConversionMessage(opts *bind.CallOpts, conversionID [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packSubnetToL1ConversionMessage", conversionID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x7f7c427a.
//
// Solidity: function packSubnetToL1ConversionMessage(bytes32 conversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackSubnetToL1ConversionMessage(conversionID [32]byte) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, conversionID)
}

// PackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x7f7c427a.
//
// Solidity: function packSubnetToL1ConversionMessage(bytes32 conversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackSubnetToL1ConversionMessage(conversionID [32]byte) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, conversionID)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackValidationUptimeMessage(opts *bind.CallOpts, validationID [32]byte, uptime uint64) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packValidationUptimeMessage", validationID, uptime)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackValidationUptimeMessage(&_ValidatorMessages.CallOpts, validationID, uptime)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackValidationUptimeMessage(&_ValidatorMessages.CallOpts, validationID, uptime)
}

// UnpackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x021de88f.
//
// Solidity: function unpackL1ValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackL1ValidatorRegistrationMessage(opts *bind.CallOpts, input []byte) ([32]byte, bool, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackL1ValidatorRegistrationMessage", input)

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// UnpackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x021de88f.
//
// Solidity: function unpackL1ValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackL1ValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x021de88f.
//
// Solidity: function unpackL1ValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackL1ValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x50782b0f.
//
// Solidity: function unpackL1ValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackL1ValidatorWeightMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, uint64, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackL1ValidatorWeightMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// UnpackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x50782b0f.
//
// Solidity: function unpackL1ValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackL1ValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x50782b0f.
//
// Solidity: function unpackL1ValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackL1ValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0x9b835465.
//
// Solidity: function unpackRegisterL1ValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackRegisterL1ValidatorMessage(opts *bind.CallOpts, input []byte) (ValidatorMessagesValidationPeriod, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackRegisterL1ValidatorMessage", input)

	if err != nil {
		return *new(ValidatorMessagesValidationPeriod), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorMessagesValidationPeriod)).(*ValidatorMessagesValidationPeriod)

	return out0, err

}

// UnpackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0x9b835465.
//
// Solidity: function unpackRegisterL1ValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesSession) UnpackRegisterL1ValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _ValidatorMessages.Contract.UnpackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0x9b835465.
//
// Solidity: function unpackRegisterL1ValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackRegisterL1ValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _ValidatorMessages.Contract.UnpackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x4d847884.
//
// Solidity: function unpackSubnetToL1ConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackSubnetToL1ConversionMessage(opts *bind.CallOpts, input []byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackSubnetToL1ConversionMessage", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnpackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x4d847884.
//
// Solidity: function unpackSubnetToL1ConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackSubnetToL1ConversionMessage(input []byte) ([32]byte, error) {
	return _ValidatorMessages.Contract.UnpackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x4d847884.
//
// Solidity: function unpackSubnetToL1ConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackSubnetToL1ConversionMessage(input []byte) ([32]byte, error) {
	return _ValidatorMessages.Contract.UnpackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackValidationUptimeMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackValidationUptimeMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _ValidatorMessages.Contract.UnpackValidationUptimeMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _ValidatorMessages.Contract.UnpackValidationUptimeMessage(&_ValidatorMessages.CallOpts, input)
}
