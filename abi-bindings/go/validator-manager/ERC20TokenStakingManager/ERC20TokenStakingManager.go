// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokenstakingmanager

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
	MessageNonce   uint64
	Weight         uint64
	StartedAt      uint64
	EndedAt        uint64
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

// ERC20TokenStakingManagerMetaData contains all meta data concerning the ERC20TokenStakingManager contract.
var ERC20TokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"InvalidBLSKeyLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"encodedConversionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedConversionID\",\"type\":\"bytes32\"}],\"name\":\"InvalidConversionID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"}],\"name\":\"InvalidDelegationFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidDelegationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidDelegatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"name\":\"InvalidMaximumChurnPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"InvalidMinStakeDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"InvalidNodeID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addressesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidPChainOwnerThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"InvalidRegistrationExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InvalidRewardRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"}],\"name\":\"InvalidStakeMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"InvalidTokenAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InvalidTotalWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidUptimeBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"}],\"name\":\"InvalidValidatorManagerAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidatorManagerBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"churnAmount\",\"type\":\"uint64\"}],\"name\":\"MaxChurnRateExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValidatorWeight\",\"type\":\"uint64\"}],\"name\":\"MaxWeightExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"MinStakeDurationNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PChainOwnerAddressesNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UnauthorizedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"validRegistration\",\"type\":\"bool\"}],\"name\":\"UnexpectedRegistrationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorNotPoS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroWeightToValueFactor\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"DelegationEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"DelegatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"InitialValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"ValidationPeriodCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidationPeriodEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidationPeriodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorWeightUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BIPS_CONVERSION_FACTOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC20_STAKING_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeDelegatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeValidatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20\",\"outputs\":[{\"internalType\":\"contractIERC20Mintable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"messageNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endedAt\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"internalType\":\"structValidatorManagerSettings\",\"name\":\"baseSettings\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"weightToValueFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structPoSValidatorManagerSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delegationAmount\",\"type\":\"uint256\"}],\"name\":\"initializeDelegatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"internalType\":\"structValidatorRegistrationInput\",\"name\":\"registrationInput\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"initializeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"registeredValidators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"weightToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051615f68380380615f6883398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b615e1b8061014d5f395ff3fe608060405234801561000f575f80fd5b506004361061023f575f3560e01c806380dd672f11610135578063b771b3bc116100b4578063d5f20ff611610079578063d5f20ff61461051c578063df93d8de1461053c578063e4a63c4014610546578063fb8b11dd1461055a578063fd7ac5e71461056d575f80fd5b8063b771b3bc146104cc578063ba3a4b97146104da578063bc5fbfec146104ed578063bee0a03f14610501578063c974d1b614610514575f80fd5b80639e1bc4ef116100fa5780639e1bc4ef1461046c5780639e478eea1461047f578063a3a65e4814610492578063a9778a7a146102ee578063afb98096146104a5575f80fd5b806380dd672f146104185780638280a25a1461042b5780638ef34c981461043357806393e24598146104465780639ae0644714610459575f80fd5b8063467ef06f116101c1578063620658561161018657806362065856146103a7578063732214f8146103ba57806376f78621146103c1578063785e9e86146103d45780637d8d2f7714610405575f80fd5b8063467ef06f146103305780634bee0040146103435780635dd6a6cb1461036457806360305d621461037757806360ad778414610394575f80fd5b806325e1c7761161020757806325e1c776146102b05780632e2194d8146102c357806335455ded146102ee57806337b9be8f1461030a5780633a1cfff61461031d575f80fd5b80630118acc4146102435780630322ed9814610258578063151d30d11461026b5780631ec447241461028a57806320d91b7a1461029d575b5f80fd5b610256610251366004614d07565b610580565b005b610256610266366004614d42565b610591565b610273600a81565b60405160ff90911681526020015b60405180910390f35b610256610298366004614d07565b610821565b6102566102ab366004614d59565b61082d565b6102566102be366004614da7565b610dcc565b6102d66102d1366004614d42565b610e40565b6040516001600160401b039091168152602001610281565b6102f761271081565b60405161ffff9091168152602001610281565b610256610318366004614ddc565b610e94565b61025661032b366004614d07565b610ea7565b61025661033e366004614e2a565b610eb3565b610356610351366004614e68565b610f85565b604051908152602001610281565b610256610372366004614ddc565b610fba565b61037f601481565b60405163ffffffff9091168152602001610281565b6102566103a2366004614da7565b610fc6565b6103566103b5366004614ecf565b61128e565b6103565f81565b6102566103cf366004614d07565b6112ae565b5f80516020615d66833981519152546001600160a01b03165b6040516001600160a01b039091168152602001610281565b610256610413366004614ddc565b6112ba565b610256610426366004614da7565b6112c6565b610273603081565b610256610441366004614eea565b611500565b610256610454366004614d42565b6115b1565b610256610467366004614ddc565b611645565b61035661047a366004614f18565b611651565b61025661048d366004614f38565b61167d565b6102566104a0366004614e2a565b61175c565b6103567f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0081565b6103ed6005600160991b0181565b6102566104e8366004614d42565b611952565b6103565f80516020615d8683398151915281565b61025661050f366004614d42565b611bab565b610273601481565b61052f61052a366004614d42565b611ce7565b6040516102819190614fe0565b6102d66202a30081565b6103565f80516020615d6683398151915281565b610256610568366004614eea565b611e36565b61035661057b366004615060565b611ecd565b61058c8383835f611f28565b505050565b5f8181525f80516020615da68339815191526020526040808220815160e0810190925280545f80516020615d8683398151915293929190829060ff1660058111156105de576105de614f6b565b60058111156105ef576105ef614f6b565b8152602001600182018054610603906150cb565b80601f016020809104026020016040519081016040528092919081815260200182805461062f906150cb565b801561067a5780601f106106515761010080835404028352916020019161067a565b820191905f5260205f20905b81548152906001019060200180831161065d57829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a090910152909150815160058111156106e5576106e5614f6b565b14610721575f8381526005830160205260409081902054905163170cc93360e21b81526107189160ff1690600401615103565b60405180910390fd5b606081015160405163854a893f60e01b8152600481018590526001600160401b0390911660248201525f60448201526005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__9063854a893f906064015f60405180830381865af4158015610798573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526107bf919081019061520c565b6040518263ffffffff1660e01b81526004016107db919061523d565b6020604051808303815f875af11580156107f7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061081b919061524f565b50505050565b61081b8383835f611f54565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb07545f80516020615d868339815191529060ff161561087f57604051637fab81e560e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108c2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e6919061524f565b83602001351461090f576040516372b0a7e760e11b815260208401356004820152602401610718565b306109206060850160408601615266565b6001600160a01b0316146109635761093e6060840160408501615266565b604051632f88120d60e21b81526001600160a01b039091166004820152602401610718565b5f6109716060850185615281565b905090505f805b828163ffffffff161015610bc2575f6109946060880188615281565b8363ffffffff168181106109aa576109aa6152c6565b90506020028101906109bc91906152da565b6109c590615345565b80516040519192505f9160068801916109dd916153c0565b90815260200160405180910390205414610a0d57805160405163a41f772f60e01b8152610718919060040161523d565b5f6002885f013584604051602001610a3c92919091825260e01b6001600160e01b031916602082015260240190565b60408051601f1981840301815290829052610a56916153c0565b602060405180830381855afa158015610a71573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610a94919061524f565b90508086600601835f0151604051610aac91906153c0565b90815260408051918290036020908101909220929092555f8381526005890190915220805460ff191660021781558251600190910190610aec9082615415565b50604082810180515f84815260058a016020529290922060028101805492516001600160401b039485166001600160c01b031990941693909317600160801b85851602176001600160c01b0316600160c01b429590951694909402939093179092556003909101805467ffffffffffffffff19169055610b6c90856154e4565b9350807fd80a750132c7a4fb67e226bdfecab0d8609bbb4e48cdad62cee218f25b758fff8360400151845f0151604051610ba7929190615504565b60405180910390a2505080610bbb90615525565b9050610978565b5060038301805467ffffffffffffffff60401b1916600160401b6001600160401b0384168102919091179091556001840154606491610c05910460ff1683615547565b6001600160401b03161015610c3857604051633e1a785160e01b81526001600160401b0382166004820152602401610718565b5f73__$fd0c147b4031eef6079b0498cbafa865f0$__634d847884610c5c87612281565b604001516040518263ffffffff1660e01b8152600401610c7c919061523d565b602060405180830381865af4158015610c97573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cbb919061524f565b90505f73__$fd0c147b4031eef6079b0498cbafa865f0$__6387418b8e886040518263ffffffff1660e01b8152600401610cf5919061569d565b5f60405180830381865af4158015610d0f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d36919081019061520c565b90505f600282604051610d4991906153c0565b602060405180830381855afa158015610d64573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d87919061524f565b9050828114610db35760405163baaea89d60e01b81526004810182905260248101849052604401610718565b5050506007909201805460ff1916600117905550505050565b610dd582612397565b610df5576040516330efa98b60e01b815260048101839052602401610718565b5f610dff83611ce7565b5190506002816005811115610e1657610e16614f6b565b14610e36578060405163170cc93360e21b81526004016107189190615103565b61081b83836123c0565b5f80610e4a61265f565b60030154610e589084615731565b9050801580610e6d57506001600160401b0381115b15610e8e5760405163222d164360e21b815260048101849052602401610718565b92915050565b610ea084848484611f54565b5050505050565b61081b8383835f612683565b610ebb6128c0565b5f610ec461265f565b90505f80610ed1846128f7565b91509150610ede82612397565b610eea57505050610f6c565b5f828152600684016020908152604080832054600b870190925290912080546001600160a01b031981169091556001600160a01b03918216911680610f2c5750805b600483516005811115610f4157610f41614f6b565b03610f5057610f508185612caf565b610f6682610f61856040015161128e565b612cd9565b50505050505b610f8260015f80516020615dc683398151915255565b50565b5f610f8e6128c0565b610f9a85858585612d0f565b9050610fb260015f80516020615dc683398151915255565b949350505050565b61081b84848484612e8f565b5f610fcf61265f565b5f848152600782016020526040808220815160e0810190925280549394509192909190829060ff16600381111561100857611008614f6b565b600381111561101957611019614f6b565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f61108f82611ce7565b90506001835160038111156110a6576110a6614f6b565b146110c7578251604051633b0d540d60e21b81526107189190600401615750565b6004815160058111156110dc576110dc614f6b565b036110f2576110ea86612ebb565b505050505050565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__6350782b0f61111789612281565b604001516040518263ffffffff1660e01b8152600401611137919061523d565b606060405180830381865af4158015611152573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611176919061576a565b50915091508184146111a357846040015160405163089938b360e11b815260040161071891815260200190565b806001600160401b031683606001516001600160401b031610806111dc5750806001600160401b03168560a001516001600160401b0316115b1561120557604051632e19bc2d60e11b81526001600160401b0382166004820152602401610718565b5f888152600787016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b810267ffffffffffffffff60401b1990921691909117909155915191825285918a917f047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6910160405180910390a35050505050505050565b5f61129761265f565b60030154610e8e906001600160401b0384166157aa565b61058c8383835f612e8f565b610ea084848484612683565b6112ce6128c0565b5f6112d761265f565b5f848152600782016020526040808220815160e0810190925280549394509192909190829060ff16600381111561131057611310614f6b565b600381111561132157611321614f6b565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c090910152905060038151600381111561139a5761139a614f6b565b146113bb578051604051633b0d540d60e21b81526107189190600401615750565b60046113ca8260400151611ce7565b5160058111156113dc576113dc614f6b565b146114db575f6113eb84612281565b90505f8073__$fd0c147b4031eef6079b0498cbafa865f0$__6350782b0f84604001516040518263ffffffff1660e01b815260040161142a919061523d565b606060405180830381865af4158015611445573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611469919061576a565b5091509150818460400151146114955760405163089938b360e11b815260048101839052602401610718565b806001600160401b03168460c001516001600160401b031611156114d757604051632e19bc2d60e11b81526001600160401b0382166004820152602401610718565b5050505b6114e484612ebb565b50506114fc60015f80516020615dc683398151915255565b5050565b5f61150961265f565b90506001600160a01b03821661153d5760405163caa903f960e01b81526001600160a01b0383166004820152602401610718565b5f8381526006820160205260409020546001600160a01b0316331461158357335b604051636e2ccd7560e11b81526001600160a01b039091166004820152602401610718565b5f928352600b01602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b5f6115ba61265f565b90505f6115c683611ce7565b51905060048160058111156115dd576115dd614f6b565b146115fd578060405163170cc93360e21b81526004016107189190615103565b5f8381526006830160205260409020546001600160a01b03163314611622573361155e565b5f83815260068301602052604090205461058c906001600160a01b031684612caf565b61081b84848484611f28565b5f61165a6128c0565b6116658333846130b5565b9050610e8e60015f80516020615dc683398151915255565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff16806116c6575080546001600160401b03808416911610155b156116e45760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b17815561170f84846132fa565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050565b5f80516020615d868339815191525f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63021de88f61178f86612281565b604001516040518263ffffffff1660e01b81526004016117af919061523d565b6040805180830381865af41580156117c9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117ed91906157c1565b915091508061181357604051632d07135360e01b81528115156004820152602401610718565b5f8281526004840160205260409020805461182d906150cb565b90505f036118515760405163089938b360e11b815260048101839052602401610718565b60015f838152600580860160205260409091205460ff169081111561187857611878614f6b565b146118ab575f8281526005840160205260409081902054905163170cc93360e21b81526107189160ff1690600401615103565b5f82815260048401602052604081206118c391614c56565b5f828152600584016020908152604091829020805460ff1916600290811782550180546001600160401b0342818116600160c01b026001600160c01b0390931692909217928390558451600160801b9093041682529181019190915283917f8629ec2bfd8d3b792ba269096bb679e08f20ba2caec0785ef663cf94788e349b910160405180910390a250505050565b5f61195b61265f565b5f838152600782016020526040808220815160e0810190925280549394509192909190829060ff16600381111561199457611994614f6b565b60038111156119a5576119a5614f6b565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115611a1e57611a1e614f6b565b14158015611a3f5750600381516003811115611a3c57611a3c614f6b565b14155b15611a60578051604051633b0d540d60e21b81526107189190600401615750565b5f611a6e8260400151611ce7565b905080606001516001600160401b03165f03611aa0576040516339b894f960e21b815260048101859052602401610718565b60408083015160608301516080840151925163854a893f60e01b81526005600160991b019363ee5b48eb9373__$fd0c147b4031eef6079b0498cbafa865f0$__9363854a893f93611b0e93906004019283526001600160401b03918216602084015216604082015260600190565b5f60405180830381865af4158015611b28573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611b4f919081019061520c565b6040518263ffffffff1660e01b8152600401611b6b919061523d565b6020604051808303815f875af1158015611b87573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ea0919061524f565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb046020526040902080545f80516020615d868339815191529190611bf2906150cb565b90505f03611c165760405163089938b360e11b815260048101839052602401610718565b60015f838152600580840160205260409091205460ff1690811115611c3d57611c3d614f6b565b14611c70575f8281526005820160205260409081902054905163170cc93360e21b81526107189160ff1690600401615103565b5f8281526004808301602052604091829020915163ee5b48eb60e01b81526005600160991b019263ee5b48eb92611ca792016157e4565b6020604051808303815f875af1158015611cc3573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058c919061524f565b611cef614c8d565b5f8281525f80516020615da6833981519152602052604090819020815160e0810190925280545f80516020615d86833981519152929190829060ff166005811115611d3c57611d3c614f6b565b6005811115611d4d57611d4d614f6b565b8152602001600182018054611d61906150cb565b80601f0160208091040260200160405190810160405280929190818152602001828054611d8d906150cb565b8015611dd85780601f10611daf57610100808354040283529160200191611dd8565b820191905f5260205f20905b815481529060010190602001808311611dbb57829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b9091048116608083015260039092015490911660a0909101529392505050565b6001600160a01b038116611e685760405163caa903f960e01b81526001600160a01b0382166004820152602401610718565b5f611e7161265f565b5f8481526007820160205260409020549091506001600160a01b03610100909104163314611e9f573361155e565b5f928352600901602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6040515f905f80516020615d86833981519152907fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0690611f10908690869061586e565b90815260200160405180910390205491505092915050565b611f3484848484611f54565b61081b57604051631036cf9160e11b815260048101859052602401610718565b5f80611f5e61265f565b5f878152600782016020526040808220815160e0810190925280549394509192909190829060ff166003811115611f9757611f97614f6b565b6003811115611fa857611fa8614f6b565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f61201e82611ce7565b905060028351600381111561203557612035614f6b565b14612056578251604051633b0d540d60e21b81526107189190600401615750565b60208301516001600160a01b031633146120f2575f8281526006850160205260409020546001600160a01b0316331461208f573361155e565b5f82815260068501602052604090205460a08201516120be91600160b01b90046001600160401b0316906154e4565b6001600160401b03164210156120f25760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610718565b60028151600581111561210757612107614f6b565b036122295760028401546080840151612129916001600160401b0316906154e4565b6001600160401b031642101561215d5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610718565b871561216f5761216d82886123c0565b505b5f8981526007850160205260409020805460ff19166003179055606083015160808201516121a89184916121a3919061587d565b613314565b505f8a8152600786016020526040812060020180546001600160401b03909316600160c01b026001600160c01b03909316929092179091556121eb84888c6134eb565b9050828a7f366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed5760405160405180910390a315159450610fb29350505050565b60048151600581111561223e5761223e614f6b565b036122655761224e83878b6134eb565b5061225889612ebb565b6001945050505050610fb2565b805160405163170cc93360e21b81526107189190600401615103565b60408051606080820183525f8083526020830152918101919091526040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa1580156122e5573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261230c919081019061589d565b915091508061232e57604051636b2f19e960e01b815260040160405180910390fd5b815115612354578151604051636ba589a560e01b81526004810191909152602401610718565b60208201516001600160a01b031615612390576020820151604051624de75d60e31b81526001600160a01b039091166004820152602401610718565b5092915050565b5f806123a161265f565b5f938452600601602052505060409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa15801561240b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612432919081019061589d565b915091508061245457604051636b2f19e960e01b815260040160405180910390fd5b5f61245d61265f565b600581015484519192501461248b578251604051636ba589a560e01b81526004810191909152602401610718565b60208301516001600160a01b0316156124c7576020830151604051624de75d60e31b81526001600160a01b039091166004820152602401610718565b60208301516001600160a01b031615612503576020830151604051624de75d60e31b81526001600160a01b039091166004820152602401610718565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63088c246386604001516040518263ffffffff1660e01b8152600401612540919061523d565b6040805180830381865af415801561255a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061257e919061592d565b915091508188146125a55760405163089938b360e11b815260048101899052602401610718565b5f8881526006840160205260409020600101546001600160401b039081169082161115612636575f888152600684016020908152604091829020600101805467ffffffffffffffff19166001600160401b038516908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a2612654565b505f8781526006830160205260409020600101546001600160401b03165b979650505050505050565b7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0090565b5f8061268d61265f565b90505f612699876136d9565b90506126a487612397565b6126b357600192505050610fb2565b5f8781526006830160205260409020546001600160a01b031633146126d8573361155e565b5f87815260068301602052604090205460a082015161270791600160b01b90046001600160401b0316906154e4565b6001600160401b03168160c001516001600160401b0316101561274e5760c081015160405163fb6ce63f60e01b81526001600160401b039091166004820152602401610718565b5f86156127665761275f88876123c0565b9050612784565b505f8781526006830160205260409020600101546001600160401b03165b600483015460408301515f916001600160a01b031690634f22429f906127a99061128e565b60a086015160c087015160405160e085901b6001600160e01b031916815260048101939093526001600160401b03918216602484018190526044840152811660648301528516608482015260a401602060405180830381865afa158015612812573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612836919061524f565b90506001600160a01b038616612862575f8981526006850160205260409020546001600160a01b031695505b5f898152600a8501602052604081208054839290612881908490615950565b90915550505f898152600b909401602052604090932080546001600160a01b0387166001600160a01b0319909116179055505015159050949350505050565b5f80516020615dc68339815191528054600119016128f157604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f612900614c8d565b5f80516020615d868339815191525f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63021de88f61293388612281565b604001516040518263ffffffff1660e01b8152600401612953919061523d565b6040805180830381865af415801561296d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061299191906157c1565b9150915080156129b857604051632d07135360e01b81528115156004820152602401610718565b5f82815260058085016020526040808320815160e08101909252805491929091839160ff909116908111156129ef576129ef614f6b565b6005811115612a0057612a00614f6b565b8152602001600182018054612a14906150cb565b80601f0160208091040260200160405190810160405280929190818152602001828054612a40906150cb565b8015612a8b5780601f10612a6257610100808354040283529160200191612a8b565b820191905f5260205f20905b815481529060010190602001808311612a6e57829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a09091015290915081516005811115612af657612af6614f6b565b14158015612b175750600181516005811115612b1457612b14614f6b565b14155b15612b3857805160405163170cc93360e21b81526107189190600401615103565b600381516005811115612b4d57612b4d614f6b565b03612b5b5760048152612b60565b600581525b836006018160200151604051612b7691906153c0565b90815260408051602092819003830190205f90819055858152600587810190935220825181548493839160ff1916906001908490811115612bb957612bb9614f6b565b021790555060208201516001820190612bd29082615415565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff19169190921617905580516005811115612c7857612c78614f6b565b60405184907f1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16905f90a39196919550909350505050565b5f612cb861265f565b5f838152600a820160205260408120805491905590915061081b84826139be565b5f80516020615d66833981519152546114fc906001600160a01b03168383613a34565b60015f80516020615dc683398151915255565b5f80612d1961265f565b600281015490915061ffff600160401b90910481169086161080612d42575061271061ffff8616115b15612d6657604051635f12e6c360e11b815261ffff86166004820152602401610718565b60028101546001600160401b039081169085161015612da2576040516202a06d60e11b81526001600160401b0385166004820152602401610718565b8054831080612db45750806001015483115b15612dd55760405163222d164360e21b815260048101849052602401610718565b5f612ddf84613a93565b90505f612deb82610e40565b90505f612df88983613ab6565b5f818152600686016020908152604080832080546001600160401b039c909c16600160b01b0267ffffffffffffffff60b01b1961ffff9e909e16600160a01b02336001600160b01b0319909e168e17179d909d169c909c178c556001909b01805467ffffffffffffffff19169055600b9096019095529790932080546001600160a01b031916909617909555509395945050505050565b612e9b84848484612683565b61081b57604051635bff683f60e11b815260048101859052602401610718565b5f612ec461265f565b5f838152600782016020526040808220815160e0810190925280549394509192909190829060ff166003811115612efd57612efd614f6b565b6003811115612f0e57612f0e614f6b565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c090910152810151909150612fab7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb01546001600160401b031690565b8260800151612fba91906154e4565b6001600160401b0316421015612fee5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610718565b5f848152600784016020908152604080832080546001600160a81b03191681556001810184905560020183905560098601909152902080546001600160a01b031981169091556001600160a01b031680613049575060208201515b5f80613056838886614004565b9150915061306f8560200151610f61876060015161128e565b6040805183815260208101839052859189917f8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993910160405180910390a350505050505050565b5f806130bf61265f565b90505f6130ce6102d185613a93565b90505f6130da87611ce7565b90506130e587612397565b613105576040516330efa98b60e01b815260048101889052602401610718565b60028151600581111561311a5761311a614f6b565b1461313b57805160405163170cc93360e21b81526107189190600401615103565b5f82826080015161314c91906154e4565b905083600201600a9054906101000a90046001600160401b031682604001516131759190615547565b6001600160401b0316816001600160401b031611156131b257604051636d51fe0560e11b81526001600160401b0382166004820152602401610718565b5f806131be8a84613314565b915091505f8a836040516020016131ec92919091825260c01b6001600160c01b031916602082015260280190565b60408051601f1981840301815291815281516020928301205f81815260078b019093529120805491925060019160ff1916828002179055505f8181526007880160209081526040918290208054610100600160a81b0319166101006001600160a01b038f16908102919091178255600182018f9055600290910180546001600160401b038b81166001600160c01b03199092168217600160801b8a8316908102919091176001600160c01b031690935585519283528916938201939093529283019190915260608201849052908c9083907fb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a2234269060800160405180910390a496505050505050505b9392505050565b6133026140b1565b61330b826140fc565b6114fc81614176565b5f8281525f80516020615da6833981519152602052604081206002015481905f80516020615d8683398151915290600160801b90046001600160401b031661335c85826141de565b5f61336687614448565b5f888152600585016020526040808220600201805467ffffffffffffffff60801b1916600160801b6001600160401b038c811691820292909217909255915163854a893f60e01b8152600481018c905291841660248301526044820152919250906005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__9063854a893f906064015f60405180830381865af415801561340f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613436919081019061520c565b6040518263ffffffff1660e01b8152600401613452919061523d565b6020604051808303815f875af115801561346e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613492919061524f565b604080516001600160401b038a811682526020820184905282519394508516928b927f07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df928290030190a3909450925050505b9250929050565b5f806134f561265f565b90505f6135058660400151611ce7565b90505f60038251600581111561351d5761351d614f6b565b148061353b575060048251600581111561353957613539614f6b565b145b1561354b575060c0810151613588565b60028251600581111561356057613560614f6b565b0361356c575042613588565b815160405163170cc93360e21b81526107189190600401615103565b86608001516001600160401b0316816001600160401b0316116135b0575f93505050506132f3565b600483015460608801515f916001600160a01b031690634f22429f906135d59061128e565b60a086015160808c01516040808e01515f90815260068b0160205281902060010154905160e086901b6001600160e01b031916815260048101949094526001600160401b0392831660248501529082166044840152818716606484015216608482015260a401602060405180830381865afa158015613656573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061367a919061524f565b90506001600160a01b03871661369257876020015196505b5f8681526008850160209081526040808320849055600990960190529390932080546001600160a01b0388166001600160a01b031990911617905550909150509392505050565b6136e1614c8d565b5f8281525f80516020615da68339815191526020526040808220815160e0810190925280545f80516020615d8683398151915293929190829060ff16600581111561372e5761372e614f6b565b600581111561373f5761373f614f6b565b8152602001600182018054613753906150cb565b80601f016020809104026020016040519081016040528092919081815260200182805461377f906150cb565b80156137ca5780601f106137a1576101008083540402835291602001916137ca565b820191905f5260205f20905b8154815290600101906020018083116137ad57829003601f168201915b50505091835250506002828101546001600160401b038082166020850152600160401b820481166040850152600160801b820481166060850152600160c01b9091048116608084015260039093015490921660a0909101529091508151600581111561383857613838614f6b565b1461386b575f8481526005830160205260409081902054905163170cc93360e21b81526107189160ff1690600401615103565b60038152426001600160401b031660c08201525f84815260058381016020526040909120825181548493839160ff19169060019084908111156138b0576138b0614f6b565b0217905550602082015160018201906138c99082615415565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff1916919092161790555f6139678582613314565b6080840151604080516001600160401b03909216825242602083015291935083925087917ffbfc4c00cddda774e9bce93712e29d12887b46526858a1afb0937cce8c30fa42910160405180910390a3509392505050565b5f5f80516020615d6683398151915280546040516340c10f1960e01b81526001600160a01b038681166004830152602482018690529293509116906340c10f19906044015f604051808303815f87803b158015613a19575f80fd5b505af1158015613a2b573d5f803e3d5ffd5b50505050505050565b6040516001600160a01b0383811660248301526044820183905261058c91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506144bd565b5f610e8e825f80516020615d66833981519152546001600160a01b03169061451e565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb07545f9060ff16613afa57604051637fab81e560e01b815260040160405180910390fd5b5f80516020615d8683398151915242613b196060860160408701614ecf565b6001600160401b0316111580613b535750613b376202a30042615950565b613b476060860160408701614ecf565b6001600160401b031610155b15613b8d57613b686060850160408601614ecf565b604051635879da1360e11b81526001600160401b039091166004820152602401610718565b60038101546001600160401b0390613bb090600160401b90048216858316615950565b1115613bda57604051633e1a785160e01b81526001600160401b0384166004820152602401610718565b613bef613bea6060860186615963565b61452a565b613bff613bea6080860186615963565b6030613c0e6020860186615977565b905014613c4057613c226020850185615977565b6040516326475b2f60e11b8152610718925060040190815260200190565b613c4a8480615977565b90505f03613c7757613c5c8480615977565b604051633e08a12560e11b81526004016107189291906159b9565b5f60068201613c868680615977565b604051613c9492919061586e565b90815260200160405180910390205414613ccd57613cb28480615977565b60405163a41f772f60e01b81526004016107189291906159b9565b613cd7835f6141de565b6040805160e08101909152815481525f90819073__$fd0c147b4031eef6079b0498cbafa865f0$__9063eb97ce519060208101613d148a80615977565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505090825250602090810190613d5c908b018b615977565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505090825250602001613da560608b0160408c01614ecf565b6001600160401b03168152602001613dc060608b018b615963565b613dc9906159cc565b8152602001613ddb60808b018b615963565b613de4906159cc565b8152602001886001600160401b03168152506040518263ffffffff1660e01b8152600401613e129190615af9565b5f60405180830381865af4158015613e2c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613e539190810190615bb0565b5f82815260048601602052604090209193509150613e718282615415565b508160068401613e818880615977565b604051613e8f92919061586e565b9081526040519081900360200181209190915563ee5b48eb60e01b81525f906005600160991b019063ee5b48eb90613ecb90859060040161523d565b6020604051808303815f875af1158015613ee7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f0b919061524f565b5f8481526005860160205260409020805460ff191660011790559050613f318780615977565b5f858152600587016020526040902060010191613f4f919083615bf3565b505f83815260058501602052604090206002810180546001600160c01b0319166001600160401b038916908117600160801b91909102176001600160c01b03169055600301805467ffffffffffffffff1916905580837fe093cab27f3fc60439e008f3a2073310d3ac9ac9b7d59e53a06b2588f7498bd688613fd18b80615977565b613fe160608e0160408f01614ecf565b604051613ff19493929190615cac565b60405180910390a3509095945050505050565b5f805f61400f61265f565b5f868152600882016020526040812080549082905591925090819080156140a3575f8781526006850160205260409020546127109061405990600160a01b900461ffff16836157aa565b6140639190615731565b91508184600a015f8981526020019081526020015f205f8282546140879190615950565b9091555061409790508282615ce3565b92506140a389846139be565b509097909650945050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166140fa57604051631afcd79f60e31b815260040160405180910390fd5b565b6141046140b1565b61410d81614693565b6141156146ac565b610f826060820135608083013561413260c0850160a08601614ecf565b61414260e0860160c08701615cf6565b614153610100870160e08801615d0f565b61010087013561416b61014089016101208a01615266565b8861014001356146bc565b61417e6140b1565b5f80516020615d668339815191526001600160a01b0382166141be57604051637330680360e01b81526001600160a01b0383166004820152602401610718565b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f80516020615d868339815191525f6001600160401b0380841690851611156142125761420b838561587d565b905061421f565b61421c848461587d565b90505b60408051608081018252600284015480825260038501546001600160401b038082166020850152600160401b8204811694840194909452600160801b900490921660608201524291158061428c575060018401548151614288916001600160401b031690615950565b8210155b156142b4576001600160401b03808416606083015282825260408201511660208201526142d3565b82816060018181516142c691906154e4565b6001600160401b03169052505b60608101516142e3906064615547565b602082015160018601546001600160401b03929092169161430e9190600160401b900460ff16615547565b6001600160401b0316101561434757606081015160405163dfae880160e01b81526001600160401b039091166004820152602401610718565b858160400181815161435991906154e4565b6001600160401b031690525060408101805186919061437990839061587d565b6001600160401b0316905250600184015460408201516064916143a791600160401b90910460ff1690615547565b6001600160401b031610156143e0576040808201519051633e1a785160e01b81526001600160401b039091166004820152602401610718565b8051600285015560208101516003909401805460408301516060909301516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b0319909316919097161717919091169390931790925550505050565b5f8181525f80516020615da68339815191526020526040812060020180545f80516020615d86833981519152919060089061449290600160401b90046001600160401b0316615d2f565b91906101000a8154816001600160401b0302191690836001600160401b031602179055915050919050565b5f6144d16001600160a01b038416836148a1565b905080515f141580156144f55750808060200190518101906144f39190615d4a565b155b1561058c57604051635274afe760e01b81526001600160a01b0384166004820152602401610718565b5f6132f38333846148ae565b6145376020820182614e2a565b63ffffffff1615801561455757506145526020820182615281565b151590505b1561459e576145696020820182614e2a565b6145766020830183615281565b60405163c08a0f1d60e01b815263ffffffff9093166004840152602483015250604401610718565b6145ab6020820182615281565b90506145ba6020830183614e2a565b63ffffffff1611156145d3576145696020820182614e2a565b60015b6145e36020830183615281565b90508110156114fc576145f96020830183615281565b614604600184615ce3565b818110614613576146136152c6565b90506020020160208101906146289190615266565b6001600160a01b031661463e6020840184615281565b8381811061464e5761464e6152c6565b90506020020160208101906146639190615266565b6001600160a01b0316101561468b57604051630dbc8d5f60e31b815260040160405180910390fd5b6001016145d6565b61469b6140b1565b6146a3614a11565b610f8281614a19565b6146b46140b1565b6140fa614b01565b6146c46140b1565b5f6146cd61265f565b905061ffff861615806146e5575061271061ffff8716115b1561470957604051635f12e6c360e11b815261ffff87166004820152602401610718565b8789111561472d5760405163222d164360e21b8152600481018a9052602401610718565b60ff851615806147405750600a60ff8616115b156147635760405163170db35960e31b815260ff86166004820152602401610718565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb01546001600160401b03166001600160401b0316876001600160401b031610156147ca576040516202a06d60e11b81526001600160401b0388166004820152602401610718565b835f036147ea5760405163a733007160e01b815260040160405180910390fd5b8161480b57604051632f6bd1db60e01b815260048101839052602401610718565b97885560018801969096556002870180546001600160401b039690961669ffffffffffffffffffff1990961695909517600160401b61ffff95909516949094029390931767ffffffffffffffff60501b191660ff92909216600160501b029190911790925560038401919091556004830180546001600160a01b0319166001600160a01b03909216919091179055600590910155565b60606132f383835f614b09565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa1580156148f4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614918919061524f565b905061492f6001600160a01b038616853086614b98565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015614973573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614997919061524f565b90508181116149fd5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610718565b614a078282615ce3565b9695505050505050565b6140fa6140b1565b614a216140b1565b80355f80516020615d868339815191529081556014614a466060840160408501615d0f565b60ff161180614a655750614a606060830160408401615d0f565b60ff16155b15614a9957614a7a6060830160408401615d0f565b604051634a59bbff60e11b815260ff9091166004820152602401610718565b614aa96060830160408401615d0f565b60018201805460ff92909216600160401b0260ff60401b19909216919091179055614ada6040830160208401614ecf565b600191909101805467ffffffffffffffff19166001600160401b0390921691909117905550565b612cfc6140b1565b606081471015614b2e5760405163cd78605960e01b8152306004820152602401610718565b5f80856001600160a01b03168486604051614b4991906153c0565b5f6040518083038185875af1925050503d805f8114614b83576040519150601f19603f3d011682016040523d82523d5f602084013e614b88565b606091505b5091509150614a07868383614bd1565b6040516001600160a01b03848116602483015283811660448301526064820183905261081b9186918216906323b872dd90608401613a61565b606082614be657614be182614c2d565b6132f3565b8151158015614bfd57506001600160a01b0384163b155b15614c2657604051639996b31560e01b81526001600160a01b0385166004820152602401610718565b50806132f3565b805115614c3d5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b508054614c62906150cb565b5f825580601f10614c71575050565b601f0160209004905f5260205f2090810190610f829190614cca565b6040805160e08101909152805f81526060602082018190525f604083018190529082018190526080820181905260a0820181905260c09091015290565b5b80821115614cde575f8155600101614ccb565b5090565b8015158114610f82575f80fd5b803563ffffffff81168114614d02575f80fd5b919050565b5f805f60608486031215614d19575f80fd5b833592506020840135614d2b81614ce2565b9150614d3960408501614cef565b90509250925092565b5f60208284031215614d52575f80fd5b5035919050565b5f8060408385031215614d6a575f80fd5b82356001600160401b03811115614d7f575f80fd5b830160808186031215614d90575f80fd5b9150614d9e60208401614cef565b90509250929050565b5f8060408385031215614db8575f80fd5b82359150614d9e60208401614cef565b6001600160a01b0381168114610f82575f80fd5b5f805f8060808587031215614def575f80fd5b843593506020850135614e0181614ce2565b9250614e0f60408601614cef565b91506060850135614e1f81614dc8565b939692955090935050565b5f60208284031215614e3a575f80fd5b6132f382614cef565b803561ffff81168114614d02575f80fd5b6001600160401b0381168114610f82575f80fd5b5f805f8060808587031215614e7b575f80fd5b84356001600160401b03811115614e90575f80fd5b850160a08188031215614ea1575f80fd5b9350614eaf60208601614e43565b92506040850135614ebf81614e54565b9396929550929360600135925050565b5f60208284031215614edf575f80fd5b81356132f381614e54565b5f8060408385031215614efb575f80fd5b823591506020830135614f0d81614dc8565b809150509250929050565b5f8060408385031215614f29575f80fd5b50508035926020909101359150565b5f80828403610180811215614f4b575f80fd5b61016080821215614f5a575f80fd5b8493508301359050614f0d81614dc8565b634e487b7160e01b5f52602160045260245ffd5b60068110614f8f57614f8f614f6b565b9052565b5f5b83811015614fad578181015183820152602001614f95565b50505f910152565b5f8151808452614fcc816020860160208601614f93565b601f01601f19169290920160200192915050565b60208152614ff2602082018351614f7f565b5f602083015160e0604084015261500d610100840182614fb5565b905060408401516001600160401b0380821660608601528060608701511660808601528060808701511660a08601528060a08701511660c08601528060c08701511660e086015250508091505092915050565b5f8060208385031215615071575f80fd5b82356001600160401b0380821115615087575f80fd5b818501915085601f83011261509a575f80fd5b8135818111156150a8575f80fd5b8660208285010111156150b9575f80fd5b60209290920196919550909350505050565b600181811c908216806150df57607f821691505b6020821081036150fd57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208101610e8e8284614f7f565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561514757615147615111565b60405290565b604080519081016001600160401b038111828210171561514757615147615111565b604051601f8201601f191681016001600160401b038111828210171561519757615197615111565b604052919050565b5f6001600160401b038211156151b7576151b7615111565b50601f01601f191660200190565b5f82601f8301126151d4575f80fd5b81516151e76151e28261519f565b61516f565b8181528460208386010111156151fb575f80fd5b610fb2826020830160208701614f93565b5f6020828403121561521c575f80fd5b81516001600160401b03811115615231575f80fd5b610fb2848285016151c5565b602081525f6132f36020830184614fb5565b5f6020828403121561525f575f80fd5b5051919050565b5f60208284031215615276575f80fd5b81356132f381614dc8565b5f808335601e19843603018112615296575f80fd5b8301803591506001600160401b038211156152af575f80fd5b6020019150600581901b36038213156134e4575f80fd5b634e487b7160e01b5f52603260045260245ffd5b5f8235605e198336030181126152ee575f80fd5b9190910192915050565b5f82601f830112615307575f80fd5b81356153156151e28261519f565b818152846020838601011115615329575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60608236031215615355575f80fd5b61535d615125565b82356001600160401b0380821115615373575f80fd5b61537f368387016152f8565b83526020850135915080821115615394575f80fd5b506153a1368286016152f8565b60208301525060408301356153b581614e54565b604082015292915050565b5f82516152ee818460208701614f93565b601f82111561058c57805f5260205f20601f840160051c810160208510156153f65750805b601f840160051c820191505b81811015610ea0575f8155600101615402565b81516001600160401b0381111561542e5761542e615111565b6154428161543c84546150cb565b846153d1565b602080601f831160018114615475575f841561545e5750858301515b5f19600386901b1c1916600185901b1785556110ea565b5f85815260208120601f198616915b828110156154a357888601518255948401946001909101908401615484565b50858210156154c057878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b03818116838216019080821115612390576123906154d0565b6001600160401b0383168152604060208201525f610fb26040830184614fb5565b5f63ffffffff80831681810361553d5761553d6154d0565b6001019392505050565b6001600160401b0381811683821602808216919082811461556a5761556a6154d0565b505092915050565b5f808335601e19843603018112615587575f80fd5b83016020810192503590506001600160401b038111156155a5575f80fd5b8036038213156134e4575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208086019550808560051b830101845f5b8781101561569057848303601f19018952813536889003605e19018112615617575f80fd5b870160606156258280615572565b82875261563583880182846155b3565b9250505061564586830183615572565b868303888801526156578382846155b3565b92505050604080830135925061566c83614e54565b6001600160401b0392909216949091019390935297830197908301906001016155f2565b5090979650505050505050565b6020815281356020820152602082013560408201525f60408301356156c181614dc8565b6001600160a01b031660608381019190915283013536849003601e190181126156e8575f80fd5b83016020810190356001600160401b03811115615703575f80fd5b8060051b3603821315615714575f80fd5b60808085015261572860a0850182846155db565b95945050505050565b5f8261574b57634e487b7160e01b5f52601260045260245ffd5b500490565b602081016004831061576457615764614f6b565b91905290565b5f805f6060848603121561577c575f80fd5b83519250602084015161578e81614e54565b604085015190925061579f81614e54565b809150509250925092565b8082028115828204841417610e8e57610e8e6154d0565b5f80604083850312156157d2575f80fd5b825191506020830151614f0d81614ce2565b5f60208083525f84546157f6816150cb565b806020870152604060018084165f8114615817576001811461583357615860565b60ff19851660408a0152604084151560051b8a01019550615860565b895f5260205f205f5b858110156158575781548b820186015290830190880161583c565b8a016040019650505b509398975050505050505050565b818382375f9101908152919050565b6001600160401b03828116828216039080821115612390576123906154d0565b5f80604083850312156158ae575f80fd5b82516001600160401b03808211156158c4575f80fd5b90840190606082870312156158d7575f80fd5b6158df615125565b8251815260208301516158f181614dc8565b6020820152604083015182811115615907575f80fd5b615913888286016151c5565b6040830152508094505050506020830151614f0d81614ce2565b5f806040838503121561593e575f80fd5b825191506020830151614f0d81614e54565b80820180821115610e8e57610e8e6154d0565b5f8235603e198336030181126152ee575f80fd5b5f808335601e1984360301811261598c575f80fd5b8301803591506001600160401b038211156159a5575f80fd5b6020019150368190038213156134e4575f80fd5b602081525f610fb26020830184866155b3565b5f604082360312156159dc575f80fd5b6159e461514d565b6159ed83614cef565b81526020808401356001600160401b0380821115615a09575f80fd5b9085019036601f830112615a1b575f80fd5b813581811115615a2d57615a2d615111565b8060051b9150615a3e84830161516f565b8181529183018401918481019036841115615a57575f80fd5b938501935b83851015615a815784359250615a7183614dc8565b8282529385019390850190615a5c565b94860194909452509295945050505050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015615aee5784516001600160a01b03168252938301936001929092019190830190615ac5565b509695505050505050565b60208152815160208201525f602083015160e06040840152615b1f610100840182614fb5565b90506040840151601f1980858403016060860152615b3d8383614fb5565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152615b6d8383615a93565b925060a08601519150808584030160c086015250615b8b8282615a93565b91505060c0840151615ba860e08501826001600160401b03169052565b509392505050565b5f8060408385031215615bc1575f80fd5b8251915060208301516001600160401b03811115615bdd575f80fd5b615be9858286016151c5565b9150509250929050565b6001600160401b03831115615c0a57615c0a615111565b615c1e83615c1883546150cb565b836153d1565b5f601f841160018114615c4f575f8515615c385750838201355b5f19600387901b1c1916600186901b178355610ea0565b5f83815260208120601f198716915b82811015615c7e5786850135825560209485019460019092019101615c5e565b5086821015615c9a575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f6001600160401b03808716835260606020840152615ccf6060840186886155b3565b915080841660408401525095945050505050565b81810381811115610e8e57610e8e6154d0565b5f60208284031215615d06575f80fd5b6132f382614e43565b5f60208284031215615d1f575f80fd5b813560ff811681146132f3575f80fd5b5f6001600160401b0380831681810361553d5761553d6154d0565b5f60208284031215615d5a575f80fd5b81516132f381614ce256fe6e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb059b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122091fd9781d213c817aed3e146319d27f6fbf2535f84575753f151f98535f5793564736f6c63430008190033",
}

// ERC20TokenStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenStakingManagerMetaData.ABI instead.
var ERC20TokenStakingManagerABI = ERC20TokenStakingManagerMetaData.ABI

// ERC20TokenStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenStakingManagerMetaData.Bin instead.
var ERC20TokenStakingManagerBin = ERC20TokenStakingManagerMetaData.Bin

// DeployERC20TokenStakingManager deploys a new Ethereum contract, binding an instance of ERC20TokenStakingManager to it.
func DeployERC20TokenStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *ERC20TokenStakingManager, error) {
	parsed, err := ERC20TokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	validatorMessagesAddr, _, _, _ := DeployValidatorMessages(auth, backend)
	ERC20TokenStakingManagerBin = strings.ReplaceAll(ERC20TokenStakingManagerBin, "__$fd0c147b4031eef6079b0498cbafa865f0$__", validatorMessagesAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenStakingManager{ERC20TokenStakingManagerCaller: ERC20TokenStakingManagerCaller{contract: contract}, ERC20TokenStakingManagerTransactor: ERC20TokenStakingManagerTransactor{contract: contract}, ERC20TokenStakingManagerFilterer: ERC20TokenStakingManagerFilterer{contract: contract}}, nil
}

// ERC20TokenStakingManager is an auto generated Go binding around an Ethereum contract.
type ERC20TokenStakingManager struct {
	ERC20TokenStakingManagerCaller     // Read-only binding to the contract
	ERC20TokenStakingManagerTransactor // Write-only binding to the contract
	ERC20TokenStakingManagerFilterer   // Log filterer for contract events
}

// ERC20TokenStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenStakingManagerSession struct {
	Contract     *ERC20TokenStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20TokenStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenStakingManagerCallerSession struct {
	Contract *ERC20TokenStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ERC20TokenStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenStakingManagerTransactorSession struct {
	Contract     *ERC20TokenStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ERC20TokenStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenStakingManagerRaw struct {
	Contract *ERC20TokenStakingManager // Generic contract binding to access the raw methods on
}

// ERC20TokenStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerCallerRaw struct {
	Contract *ERC20TokenStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerTransactorRaw struct {
	Contract *ERC20TokenStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenStakingManager creates a new instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManager(address common.Address, backend bind.ContractBackend) (*ERC20TokenStakingManager, error) {
	contract, err := bindERC20TokenStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManager{ERC20TokenStakingManagerCaller: ERC20TokenStakingManagerCaller{contract: contract}, ERC20TokenStakingManagerTransactor: ERC20TokenStakingManagerTransactor{contract: contract}, ERC20TokenStakingManagerFilterer: ERC20TokenStakingManagerFilterer{contract: contract}}, nil
}

// NewERC20TokenStakingManagerCaller creates a new read-only instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenStakingManagerCaller, error) {
	contract, err := bindERC20TokenStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCaller{contract: contract}, nil
}

// NewERC20TokenStakingManagerTransactor creates a new write-only instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenStakingManagerTransactor, error) {
	contract, err := bindERC20TokenStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerTransactor{contract: contract}, nil
}

// NewERC20TokenStakingManagerFilterer creates a new log filterer instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenStakingManagerFilterer, error) {
	contract, err := bindERC20TokenStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerFilterer{contract: contract}, nil
}

// bindERC20TokenStakingManager binds a generic wrapper to an already deployed contract.
func bindERC20TokenStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ADDRESSLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "ADDRESS_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ADDRESSLENGTH() (uint32, error) {
	return _ERC20TokenStakingManager.Contract.ADDRESSLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ADDRESSLENGTH() (uint32, error) {
	return _ERC20TokenStakingManager.Contract.ADDRESSLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20TokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20TokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) BLSPUBLICKEYLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "BLS_PUBLIC_KEY_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ERC20STAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "ERC20_STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ERC20STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.ERC20STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ERC20STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.ERC20STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMCHURNPERCENTAGELIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_CHURN_PERCENTAGE_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMREGISTRATIONEXPIRYLENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_REGISTRATION_EXPIRY_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) POSVALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "POS_VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) PCHAINBLOCKCHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "P_CHAIN_BLOCKCHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_ERC20TokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_ERC20TokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) VALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.WARPMESSENGER(&_ERC20TokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.WARPMESSENGER(&_ERC20TokenStakingManager.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) Erc20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "erc20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) Erc20() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.Erc20(&_ERC20TokenStakingManager.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) Erc20() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.Erc20(&_ERC20TokenStakingManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ERC20TokenStakingManager.Contract.GetValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ERC20TokenStakingManager.Contract.GetValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) RegisteredValidators(opts *bind.CallOpts, nodeID []byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "registeredValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.RegisteredValidators(&_ERC20TokenStakingManager.CallOpts, nodeID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.RegisteredValidators(&_ERC20TokenStakingManager.CallOpts, nodeID)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20TokenStakingManager.Contract.WeightToValue(&_ERC20TokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20TokenStakingManager.Contract.WeightToValue(&_ERC20TokenStakingManager.CallOpts, weight)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ChangeDelegatorRewardRecipient(opts *bind.TransactOpts, delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "changeDelegatorRewardRecipient", delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ChangeValidatorRewardRecipient(opts *bind.TransactOpts, validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "changeValidatorRewardRecipient", validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimDelegationFees(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimDelegationFees(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndValidation(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndValidation(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndDelegation0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation0", delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndValidation0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation0", validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e478eea.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initialize", settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e478eea.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) Initialize(settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.Initialize(&_ERC20TokenStakingManager.TransactOpts, settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e478eea.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) Initialize(settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.Initialize(&_ERC20TokenStakingManager.TransactOpts, settings, token)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0x9e1bc4ef.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeDelegatorRegistration", validationID, delegationAmount)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0x9e1bc4ef.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0x9e1bc4ef.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndDelegation0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndDelegation0", delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndValidation0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndValidation0", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x4bee0040.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeValidatorRegistration", registrationInput, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x4bee0040.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x4bee0040.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeValidatorSet", conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorSet(&_ERC20TokenStakingManager.TransactOpts, conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorSet(&_ERC20TokenStakingManager.TransactOpts, conversionData, messageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendEndValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendEndValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendRegisterValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendRegisterValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.SubmitUptimeProof(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.SubmitUptimeProof(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// ERC20TokenStakingManagerDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegationEndedIterator struct {
	Event *ERC20TokenStakingManagerDelegationEnded // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegationEnded)
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
		it.Event = new(ERC20TokenStakingManagerDelegationEnded)
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
func (it *ERC20TokenStakingManagerDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegationEnded represents a DelegationEnded event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegationEndedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegationEnded)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegationEnded(log types.Log) (*ERC20TokenStakingManagerDelegationEnded, error) {
	event := new(ERC20TokenStakingManagerDelegationEnded)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorAddedIterator struct {
	Event *ERC20TokenStakingManagerDelegatorAdded // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorAdded)
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
		it.Event = new(ERC20TokenStakingManagerDelegatorAdded)
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
func (it *ERC20TokenStakingManagerDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorAdded represents a DelegatorAdded event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorAdded struct {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*ERC20TokenStakingManagerDelegatorAddedIterator, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorAddedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorAdded)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorAdded(log types.Log) (*ERC20TokenStakingManagerDelegatorAdded, error) {
	event := new(ERC20TokenStakingManagerDelegatorAdded)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRegisteredIterator struct {
	Event *ERC20TokenStakingManagerDelegatorRegistered // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorRegistered)
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
		it.Event = new(ERC20TokenStakingManagerDelegatorRegistered)
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
func (it *ERC20TokenStakingManagerDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorRegistered represents a DelegatorRegistered event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorRegisteredIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorRegistered)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorRegistered(log types.Log) (*ERC20TokenStakingManagerDelegatorRegistered, error) {
	event := new(ERC20TokenStakingManagerDelegatorRegistered)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRemovalInitializedIterator struct {
	Event *ERC20TokenStakingManagerDelegatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
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
		it.Event = new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
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
func (it *ERC20TokenStakingManagerDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorRemovalInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*ERC20TokenStakingManagerDelegatorRemovalInitialized, error) {
	event := new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialValidatorCreatedIterator struct {
	Event *ERC20TokenStakingManagerInitialValidatorCreated // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitialValidatorCreated)
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
		it.Event = new(ERC20TokenStakingManagerInitialValidatorCreated)
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
func (it *ERC20TokenStakingManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	Weight       uint64
	NodeID       []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0xd80a750132c7a4fb67e226bdfecab0d8609bbb4e48cdad62cee218f25b758fff.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, uint64 weight, bytes nodeID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitialValidatorCreatedIterator{contract: _ERC20TokenStakingManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0xd80a750132c7a4fb67e226bdfecab0d8609bbb4e48cdad62cee218f25b758fff.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, uint64 weight, bytes nodeID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitialValidatorCreated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitialValidatorCreated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
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

// ParseInitialValidatorCreated is a log parse operation binding the contract event 0xd80a750132c7a4fb67e226bdfecab0d8609bbb4e48cdad62cee218f25b758fff.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, uint64 weight, bytes nodeID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*ERC20TokenStakingManagerInitialValidatorCreated, error) {
	event := new(ERC20TokenStakingManagerInitialValidatorCreated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitializedIterator struct {
	Event *ERC20TokenStakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitialized)
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
		it.Event = new(ERC20TokenStakingManagerInitialized)
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
func (it *ERC20TokenStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitialized represents a Initialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20TokenStakingManagerInitializedIterator, error) {

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitialized(log types.Log) (*ERC20TokenStakingManagerInitialized, error) {
	event := new(ERC20TokenStakingManagerInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerUptimeUpdatedIterator struct {
	Event *ERC20TokenStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerUptimeUpdated)
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
		it.Event = new(ERC20TokenStakingManagerUptimeUpdated)
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
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerUptimeUpdatedIterator{contract: _ERC20TokenStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerUptimeUpdated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*ERC20TokenStakingManagerUptimeUpdated, error) {
	event := new(ERC20TokenStakingManagerUptimeUpdated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodCreatedIterator struct {
	Event *ERC20TokenStakingManagerValidationPeriodCreated // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationPeriodCreated)
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
		it.Event = new(ERC20TokenStakingManagerValidationPeriodCreated)
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
func (it *ERC20TokenStakingManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	RegisterValidationMessageID [32]byte
	Weight                      uint64
	NodeID                      []byte
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0xe093cab27f3fc60439e008f3a2073310d3ac9ac9b7d59e53a06b2588f7498bd6.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes32 indexed registerValidationMessageID, uint64 weight, bytes nodeID, uint64 registrationExpiry)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, registerValidationMessageID [][32]byte) (*ERC20TokenStakingManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationPeriodCreatedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xe093cab27f3fc60439e008f3a2073310d3ac9ac9b7d59e53a06b2588f7498bd6.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes32 indexed registerValidationMessageID, uint64 weight, bytes nodeID, uint64 registrationExpiry)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationPeriodCreated, validationID [][32]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationPeriodCreated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
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

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0xe093cab27f3fc60439e008f3a2073310d3ac9ac9b7d59e53a06b2588f7498bd6.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes32 indexed registerValidationMessageID, uint64 weight, bytes nodeID, uint64 registrationExpiry)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*ERC20TokenStakingManagerValidationPeriodCreated, error) {
	event := new(ERC20TokenStakingManagerValidationPeriodCreated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodEndedIterator struct {
	Event *ERC20TokenStakingManagerValidationPeriodEnded // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationPeriodEnded)
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
		it.Event = new(ERC20TokenStakingManagerValidationPeriodEnded)
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
func (it *ERC20TokenStakingManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*ERC20TokenStakingManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationPeriodEndedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationPeriodEnded)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
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

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*ERC20TokenStakingManagerValidationPeriodEnded, error) {
	event := new(ERC20TokenStakingManagerValidationPeriodEnded)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodRegisteredIterator struct {
	Event *ERC20TokenStakingManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationPeriodRegistered)
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
		it.Event = new(ERC20TokenStakingManagerValidationPeriodRegistered)
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
func (it *ERC20TokenStakingManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       uint64
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0x8629ec2bfd8d3b792ba269096bb679e08f20ba2caec0785ef663cf94788e349b.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint64 weight, uint256 timestamp)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationPeriodRegisteredIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0x8629ec2bfd8d3b792ba269096bb679e08f20ba2caec0785ef663cf94788e349b.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint64 weight, uint256 timestamp)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationPeriodRegistered)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
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

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0x8629ec2bfd8d3b792ba269096bb679e08f20ba2caec0785ef663cf94788e349b.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint64 weight, uint256 timestamp)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*ERC20TokenStakingManagerValidationPeriodRegistered, error) {
	event := new(ERC20TokenStakingManagerValidationPeriodRegistered)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRemovalInitializedIterator struct {
	Event *ERC20TokenStakingManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidatorRemovalInitialized)
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
		it.Event = new(ERC20TokenStakingManagerValidatorRemovalInitialized)
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
func (it *ERC20TokenStakingManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             uint64
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0xfbfc4c00cddda774e9bce93712e29d12887b46526858a1afb0937cce8c30fa42.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint64 weight, uint256 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*ERC20TokenStakingManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidatorRemovalInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0xfbfc4c00cddda774e9bce93712e29d12887b46526858a1afb0937cce8c30fa42.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint64 weight, uint256 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidatorRemovalInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
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

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0xfbfc4c00cddda774e9bce93712e29d12887b46526858a1afb0937cce8c30fa42.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint64 weight, uint256 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*ERC20TokenStakingManagerValidatorRemovalInitialized, error) {
	event := new(ERC20TokenStakingManagerValidatorRemovalInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorWeightUpdateIterator struct {
	Event *ERC20TokenStakingManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidatorWeightUpdate)
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
		it.Event = new(ERC20TokenStakingManagerValidatorWeightUpdate)
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
func (it *ERC20TokenStakingManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	Weight             uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 weight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*ERC20TokenStakingManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidatorWeightUpdateIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 weight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidatorWeightUpdate)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
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

// ParseValidatorWeightUpdate is a log parse operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 weight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*ERC20TokenStakingManagerValidatorWeightUpdate, error) {
	event := new(ERC20TokenStakingManagerValidatorWeightUpdate)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMessagesMetaData contains all meta data concerning the ValidatorMessages contract.
var ValidatorMessagesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidBLSPublicKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"InvalidCodecID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"actual\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expected\",\"type\":\"uint32\"}],\"name\":\"InvalidMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageType\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"}],\"name\":\"packConversionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"packL1ValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"packL1ValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"validationPeriod\",\"type\":\"tuple\"}],\"name\":\"packRegisterL1ValidatorMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conversionID\",\"type\":\"bytes32\"}],\"name\":\"packSubnetToL1ConversionMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"packValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackL1ValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackL1ValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackRegisterL1ValidatorMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackSubnetToL1ConversionMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x61217b610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b1575f3560e01c8063854a893f11610079578063854a893f146101b257806387418b8e1461020f5780639b83546514610222578063a699c13514610242578063e1d68f3014610255578063eb97ce5114610268575f80fd5b8063021de88f146100b5578063088c2463146100e25780634d8478841461011257806350782b0f146101335780637f7c427a1461016b575b5f80fd5b6100c86100c33660046118a9565b610289565b604080519283529015156020830152015b60405180910390f35b6100f56100f03660046118a9565b61044a565b604080519283526001600160401b039091166020830152016100d9565b6101256101203660046118a9565b61063b565b6040519081526020016100d9565b6101466101413660046118a9565b6107c8565b604080519384526001600160401b0392831660208501529116908201526060016100d9565b6101a56101793660046118e2565b604080515f60208201819052602282015260268082019390935281518082039093018352604601905290565b6040516100d99190611946565b6101a56101c036600461197a565b604080515f6020820152600360e01b602282015260268101949094526001600160c01b031960c093841b811660468601529190921b16604e830152805180830360360181526056909201905290565b6101a561021d3660046119eb565b610a1e565b6102356102303660046118a9565b610b60565b6040516100d99190611bb4565b6101a5610250366004611c6b565b6114ab565b6101a5610263366004611c9d565b6114ef565b61027b610276366004611d80565b611525565b6040516100d9929190611e7c565b5f8082516027146102c457825160405163cc92daa160e01b815263ffffffff9091166004820152602760248201526044015b60405180910390fd5b5f805b6002811015610313576102db816001611ea8565b6102e6906008611ebb565b61ffff168582815181106102fc576102fc611ed2565b016020015160f81c901b91909117906001016102c7565b5061ffff81161561033d5760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561039857610354816003611ea8565b61035f906008611ebb565b63ffffffff1686610371836002611ee6565b8151811061038157610381611ed2565b016020015160f81c901b9190911790600101610340565b5063ffffffff81166002146103c057604051635b60892f60e01b815260040160405180910390fd5b5f805b6020811015610415576103d781601f611ea8565b6103e2906008611ebb565b876103ee836006611ee6565b815181106103fe576103fe611ed2565b016020015160f81c901b91909117906001016103c3565b505f8660268151811061042a5761042a611ed2565b016020015191976001600160f81b03199092161515965090945050505050565b5f808251602e1461048057825160405163cc92daa160e01b815263ffffffff9091166004820152602e60248201526044016102bb565b5f805b60028110156104cf57610497816001611ea8565b6104a2906008611ebb565b61ffff168582815181106104b8576104b8611ed2565b016020015160f81c901b9190911790600101610483565b5061ffff8116156104f95760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561055457610510816003611ea8565b61051b906008611ebb565b63ffffffff168661052d836002611ee6565b8151811061053d5761053d611ed2565b016020015160f81c901b91909117906001016104fc565b5063ffffffff81161561057a57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156105cf5761059181601f611ea8565b61059c906008611ebb565b876105a8836006611ee6565b815181106105b8576105b8611ed2565b016020015160f81c901b919091179060010161057d565b505f805b600881101561062e576105e7816007611ea8565b6105f2906008611ebb565b6001600160401b031688610607836026611ee6565b8151811061061757610617611ed2565b016020015160f81c901b91909117906001016105d3565b5090969095509350505050565b5f815160261461067057815160405163cc92daa160e01b815263ffffffff9091166004820152602660248201526044016102bb565b5f805b60028110156106bf57610687816001611ea8565b610692906008611ebb565b61ffff168482815181106106a8576106a8611ed2565b016020015160f81c901b9190911790600101610673565b5061ffff8116156106e95760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561074457610700816003611ea8565b61070b906008611ebb565b63ffffffff168561071d836002611ee6565b8151811061072d5761072d611ed2565b016020015160f81c901b91909117906001016106ec565b5063ffffffff81161561076a57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156107bf5761078181601f611ea8565b61078c906008611ebb565b86610798836006611ee6565b815181106107a8576107a8611ed2565b016020015160f81c901b919091179060010161076d565b50949350505050565b5f805f83516036146107ff57835160405163cc92daa160e01b815263ffffffff9091166004820152603660248201526044016102bb565b5f805b600281101561084e57610816816001611ea8565b610821906008611ebb565b61ffff1686828151811061083757610837611ed2565b016020015160f81c901b9190911790600101610802565b5061ffff8116156108785760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b60048110156108d35761088f816003611ea8565b61089a906008611ebb565b63ffffffff16876108ac836002611ee6565b815181106108bc576108bc611ed2565b016020015160f81c901b919091179060010161087b565b5063ffffffff81166003146108fb57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156109505761091281601f611ea8565b61091d906008611ebb565b88610929836006611ee6565b8151811061093957610939611ed2565b016020015160f81c901b91909117906001016108fe565b505f805b60088110156109af57610968816007611ea8565b610973906008611ebb565b6001600160401b031689610988836026611ee6565b8151811061099857610998611ed2565b016020015160f81c901b9190911790600101610954565b505f805b6008811015610a0e576109c7816007611ea8565b6109d2906008611ebb565b6001600160401b03168a6109e783602e611ee6565b815181106109f7576109f7611ed2565b016020015160f81c901b91909117906001016109b3565b5091989097509095509350505050565b80516020808301516040808501516060868101515192515f95810186905260228101969096526042860193909352600560e21b60628601526bffffffffffffffffffffffff1990831b16606685015260e01b6001600160e01b031916607a84015291607e0160405160208183030381529060405290505f5b836060015151811015610b59578184606001518281518110610aba57610aba611ed2565b60200260200101515f01515185606001518381518110610adc57610adc611ed2565b60200260200101515f015186606001518481518110610afd57610afd611ed2565b60200260200101516020015187606001518581518110610b1f57610b1f611ed2565b602002602001015160400151604051602001610b3f959493929190611ef9565b60408051601f198184030181529190529150600101610a96565b5092915050565b610b68611712565b5f610b71611712565b5f805b6002811015610bcf57610b88816001611ea8565b610b93906008611ebb565b61ffff1686610ba863ffffffff871684611ee6565b81518110610bb857610bb8611ed2565b016020015160f81c901b9190911790600101610b74565b5061ffff811615610bf95760405163407b587360e01b815261ffff821660048201526024016102bb565b610c04600284611f72565b9250505f805b6004811015610c6957610c1e816003611ea8565b610c29906008611ebb565b63ffffffff16868563ffffffff1683610c429190611ee6565b81518110610c5257610c52611ed2565b016020015160f81c901b9190911790600101610c0a565b5063ffffffff8116600114610c9157604051635b60892f60e01b815260040160405180910390fd5b610c9c600484611f72565b9250505f805b6020811015610cf957610cb681601f611ea8565b610cc1906008611ebb565b86610cd263ffffffff871684611ee6565b81518110610ce257610ce2611ed2565b016020015160f81c901b9190911790600101610ca2565b50808252610d08602084611f72565b9250505f805b6004811015610d6d57610d22816003611ea8565b610d2d906008611ebb565b63ffffffff16868563ffffffff1683610d469190611ee6565b81518110610d5657610d56611ed2565b016020015160f81c901b9190911790600101610d0e565b50610d79600484611f72565b92505f8163ffffffff166001600160401b03811115610d9a57610d9a61176c565b6040519080825280601f01601f191660200182016040528015610dc4576020820181803683370190505b5090505f5b8263ffffffff16811015610e335786610de863ffffffff871683611ee6565b81518110610df857610df8611ed2565b602001015160f81c60f81b828281518110610e1557610e15611ed2565b60200101906001600160f81b03191690815f1a905350600101610dc9565b5060208301819052610e458285611f72565b604080516030808252606082019092529195505f92506020820181803683370190505090505f5b6030811015610ed15786610e8663ffffffff871683611ee6565b81518110610e9657610e96611ed2565b602001015160f81c60f81b828281518110610eb357610eb3611ed2565b60200101906001600160f81b03191690815f1a905350600101610e6c565b5060408301819052610ee4603085611f72565b9350505f805b6008811015610f4a57610efe816007611ea8565b610f09906008611ebb565b6001600160401b031687610f2363ffffffff881684611ee6565b81518110610f3357610f33611ed2565b016020015160f81c901b9190911790600101610eea565b506001600160401b0381166060840152610f65600885611f72565b9350505f805f5b6004811015610fcb57610f80816003611ea8565b610f8b906008611ebb565b63ffffffff16888763ffffffff1683610fa49190611ee6565b81518110610fb457610fb4611ed2565b016020015160f81c901b9190911790600101610f6c565b50610fd7600486611f72565b94505f5b600481101561103a57610fef816003611ea8565b610ffa906008611ebb565b63ffffffff16888763ffffffff16836110139190611ee6565b8151811061102357611023611ed2565b016020015160f81c901b9290921791600101610fdb565b50611046600486611f72565b94505f8263ffffffff166001600160401b038111156110675761106761176c565b604051908082528060200260200182016040528015611090578160200160208202803683370190505b5090505f5b8363ffffffff16811015611178576040805160148082528183019092525f916020820181803683370190505090505f5b601481101561112a578a6110df63ffffffff8b1683611ee6565b815181106110ef576110ef611ed2565b602001015160f81c60f81b82828151811061110c5761110c611ed2565b60200101906001600160f81b03191690815f1a9053506001016110c5565b505f601482015190508084848151811061114657611146611ed2565b6001600160a01b039092166020928302919091019091015261116960148a611f72565b98505050806001019050611095565b506040805180820190915263ffffffff9092168252602082015260808401525f80805b60048110156111fa576111af816003611ea8565b6111ba906008611ebb565b63ffffffff16898863ffffffff16836111d39190611ee6565b815181106111e3576111e3611ed2565b016020015160f81c901b919091179060010161119b565b50611206600487611f72565b95505f5b60048110156112695761121e816003611ea8565b611229906008611ebb565b63ffffffff16898863ffffffff16836112429190611ee6565b8151811061125257611252611ed2565b016020015160f81c901b929092179160010161120a565b50611275600487611f72565b95505f8263ffffffff166001600160401b038111156112965761129661176c565b6040519080825280602002602001820160405280156112bf578160200160208202803683370190505b5090505f5b8363ffffffff168110156113a7576040805160148082528183019092525f916020820181803683370190505090505f5b6014811015611359578b61130e63ffffffff8c1683611ee6565b8151811061131e5761131e611ed2565b602001015160f81c60f81b82828151811061133b5761133b611ed2565b60200101906001600160f81b03191690815f1a9053506001016112f4565b505f601482015190508084848151811061137557611375611ed2565b6001600160a01b039092166020928302919091019091015261139860148b611f72565b995050508060010190506112c4565b506040805180820190915263ffffffff9092168252602082015260a08501525f6113d18284611f72565b6113dc906014611f8f565b6113e785607a611f72565b6113f19190611f72565b90508063ffffffff1688511461142d57875160405163cc92daa160e01b815263ffffffff918216600482015290821660248201526044016102bb565b5f805b600881101561149057611444816007611ea8565b61144f906008611ebb565b6001600160401b03168a61146963ffffffff8b1684611ee6565b8151811061147957611479611ed2565b016020015160f81c901b9190911790600101611430565b506001600160401b031660c086015250929695505050505050565b6040515f6020820152600160e11b60228201526026810183905281151560f81b60468201526060906047015b60405160208183030381529060405290505b92915050565b6040515f602082018190526022820152602681018390526001600160c01b031960c083901b166046820152606090604e016114d7565b5f606082604001515160301461154e5760405163180ffa0d60e01b815260040160405180910390fd5b82516020808501518051604080880151606089015160808a01518051908701515193515f9861158f988a986001989297929690959094909390929101611fb7565b60405160208183030381529060405290505f5b84608001516020015151811015611601578185608001516020015182815181106115ce576115ce611ed2565b60200260200101516040516020016115e7929190612071565b60408051601f1981840301815291905291506001016115a2565b5060a08401518051602091820151516040516116219385939291016120a7565b60405160208183030381529060405290505f5b8460a00151602001515181101561169357818560a0015160200151828151811061166057611660611ed2565b6020026020010151604051602001611679929190612071565b60408051601f198184030181529190529150600101611634565b5060c08401516040516116aa9183916020016120e2565b60405160208183030381529060405290506002816040516116cb9190612113565b602060405180830381855afa1580156116e6573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190611709919061212e565b94909350915050565b6040805160e0810182525f808252606060208084018290528385018290528184018390528451808601865283815280820183905260808501528451808601909552918452908301529060a082019081525f60209091015290565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b03811182821017156117a2576117a261176c565b60405290565b604051606081016001600160401b03811182821017156117a2576117a261176c565b604080519081016001600160401b03811182821017156117a2576117a261176c565b60405160e081016001600160401b03811182821017156117a2576117a261176c565b604051601f8201601f191681016001600160401b03811182821017156118365761183661176c565b604052919050565b5f82601f83011261184d575f80fd5b81356001600160401b038111156118665761186661176c565b611879601f8201601f191660200161180e565b81815284602083860101111561188d575f80fd5b816020850160208301375f918101602001919091529392505050565b5f602082840312156118b9575f80fd5b81356001600160401b038111156118ce575f80fd5b6118da8482850161183e565b949350505050565b5f602082840312156118f2575f80fd5b5035919050565b5f5b838110156119135781810151838201526020016118fb565b50505f910152565b5f81518084526119328160208601602086016118f9565b601f01601f19169290920160200192915050565b602081525f611958602083018461191b565b9392505050565b80356001600160401b0381168114611975575f80fd5b919050565b5f805f6060848603121561198c575f80fd5b8335925061199c6020850161195f565b91506119aa6040850161195f565b90509250925092565b80356001600160a01b0381168114611975575f80fd5b5f6001600160401b038211156119e1576119e161176c565b5060051b60200190565b5f60208083850312156119fc575f80fd5b82356001600160401b0380821115611a12575f80fd5b9084019060808287031215611a25575f80fd5b611a2d611780565b823581528383013584820152611a45604084016119b3565b604082015260608084013583811115611a5c575f80fd5b80850194505087601f850112611a70575f80fd5b8335611a83611a7e826119c9565b61180e565b81815260059190911b8501860190868101908a831115611aa1575f80fd5b8787015b83811015611b3a57803587811115611abb575f80fd5b8801808d03601f1901861315611acf575f80fd5b611ad76117a8565b8a82013589811115611ae7575f80fd5b611af58f8d8386010161183e565b825250604082013589811115611b09575f80fd5b611b178f8d8386010161183e565b8c83015250611b2787830161195f565b6040820152845250918801918801611aa5565b506060850152509198975050505050505050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611ba95784516001600160a01b03168252938301936001929092019190830190611b80565b509695505050505050565b60208152815160208201525f602083015160e06040840152611bda61010084018261191b565b90506040840151601f1980858403016060860152611bf8838361191b565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152611c288383611b4e565b925060a08601519150808584030160c086015250611c468282611b4e565b91505060c0840151611c6360e08501826001600160401b03169052565b509392505050565b5f8060408385031215611c7c575f80fd5b8235915060208301358015158114611c92575f80fd5b809150509250929050565b5f8060408385031215611cae575f80fd5b82359150611cbe6020840161195f565b90509250929050565b5f60408284031215611cd7575f80fd5b611cdf6117ca565b9050813563ffffffff81168114611cf4575f80fd5b81526020828101356001600160401b03811115611d0f575f80fd5b8301601f81018513611d1f575f80fd5b8035611d2d611a7e826119c9565b81815260059190911b82018301908381019087831115611d4b575f80fd5b928401925b82841015611d7057611d61846119b3565b82529284019290840190611d50565b8085870152505050505092915050565b5f60208284031215611d90575f80fd5b81356001600160401b0380821115611da6575f80fd5b9083019060e08286031215611db9575f80fd5b611dc16117ec565b82358152602083013582811115611dd6575f80fd5b611de28782860161183e565b602083015250604083013582811115611df9575f80fd5b611e058782860161183e565b604083015250611e176060840161195f565b6060820152608083013582811115611e2d575f80fd5b611e3987828601611cc7565b60808301525060a083013582811115611e50575f80fd5b611e5c87828601611cc7565b60a083015250611e6e60c0840161195f565b60c082015295945050505050565b828152604060208201525f6118da604083018461191b565b634e487b7160e01b5f52601160045260245ffd5b818103818111156114e9576114e9611e94565b80820281158282048414176114e9576114e9611e94565b634e487b7160e01b5f52603260045260245ffd5b808201808211156114e9576114e9611e94565b5f8651611f0a818460208b016118f9565b60e087901b6001600160e01b0319169083019081528551611f32816004840160208a016118f9565b8551910190611f488160048401602089016118f9565b60c09490941b6001600160c01b031916600491909401908101939093525050600c01949350505050565b63ffffffff818116838216019080821115610b5957610b59611e94565b63ffffffff818116838216028082169190828114611faf57611faf611e94565b505092915050565b61ffff60f01b8a60f01b1681525f63ffffffff60e01b808b60e01b166002840152896006840152808960e01b166026840152508651611ffd81602a850160208b016118f9565b86519083019061201481602a840160208b016118f9565b60c087901b6001600160c01b031916602a9290910191820152612046603282018660e01b6001600160e01b0319169052565b61205f603682018560e01b6001600160e01b0319169052565b603a019b9a5050505050505050505050565b5f83516120828184602088016118f9565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b5f84516120b88184602089016118f9565b6001600160e01b031960e095861b8116919093019081529290931b16600482015260080192915050565b5f83516120f38184602088016118f9565b60c09390931b6001600160c01b0319169190920190815260080192915050565b5f82516121248184602087016118f9565b9190910192915050565b5f6020828403121561213e575f80fd5b505191905056fea26469706673582212207a0b6ee4bb75144dd6b8635b3883de8d9364507951a0e3d076ca36181dac9a1764736f6c63430008190033",
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
