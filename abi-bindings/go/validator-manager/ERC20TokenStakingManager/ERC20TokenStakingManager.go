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

// ERC20TokenStakingManagerMetaData contains all meta data concerning the ERC20TokenStakingManager contract.
var ERC20TokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"InvalidBLSKeyLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"encodedConversionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedConversionID\",\"type\":\"bytes32\"}],\"name\":\"InvalidConversionID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"}],\"name\":\"InvalidDelegationFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidDelegationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidDelegatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"name\":\"InvalidMaximumChurnPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"InvalidMinStakeDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"InvalidNodeID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addressesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidPChainOwnerThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"InvalidRegistrationExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InvalidRewardRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"}],\"name\":\"InvalidStakeMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"InvalidTokenAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InvalidTotalWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidUptimeBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"}],\"name\":\"InvalidValidatorManagerAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidatorManagerBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"churnAmount\",\"type\":\"uint64\"}],\"name\":\"MaxChurnRateExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValidatorWeight\",\"type\":\"uint64\"}],\"name\":\"MaxWeightExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"MinStakeDurationNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PChainOwnerAddressesNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UnauthorizedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"validRegistration\",\"type\":\"bool\"}],\"name\":\"UnexpectedRegistrationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedValidationID\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorNotPoS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroWeightToValueFactor\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"CompletedDelegatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"CompletedDelegatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"CompletedValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"CompletedValidatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"CompletedValidatorWeightUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"InitiatedDelegatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InitiatedDelegatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeID\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"registrationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"InitiatedValidatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"weightUpdateMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedValidatorWeightUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeID\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"RegisteredInitialValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BIPS_CONVERSION_FACTOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC20_STAKING_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_ID_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeDelegatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeValidatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorWeightUpdate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20\",\"outputs\":[{\"internalType\":\"contractIERC20Mintable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitiateDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"forceInitiateDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"forceInitiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sentNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receivedNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"internalType\":\"structValidatorManagerSettings\",\"name\":\"baseSettings\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"weightToValueFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structPoSValidatorManagerSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delegationAmount\",\"type\":\"uint256\"}],\"name\":\"initiateDelegatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initiateDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initiateDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"initiateValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1TotalWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"registeredValidators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subnetID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"weightToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051615e13380380615e1383398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b615cc68061014d5f395ff3fe608060405234801561000f575f80fd5b506004361061026b575f3560e01c80638af5499e1161014b578063bb0b1938116100bf578063ce161f1411610084578063ce161f141461056b578063d5f20ff61461059b578063df93d8de146105bb578063e4a63c40146105c5578063fb8b11dd146105d9578063fd7ac5e7146105ec575f80fd5b8063bb0b19381461050e578063bc5fbfec14610516578063bee0a03f1461053d578063c974d1b614610550578063cc71bbba14610558575f80fd5b8063a3a65e4811610110578063a3a65e48146104a0578063a9778a7a14610374578063aac80c39146104b3578063afb98096146104c6578063b2c1712e146104ed578063b771b3bc14610500575f80fd5b80638af5499e146104415780638ef34c981461045457806393e24598146104675780639681d9401461047a5780639e478eea1461048d575f80fd5b806335455ded116101e257806362065856116101a757806362065856146103db57806363e2ca97146103ab578063732214f8146103ee578063739b4eff146103f5578063785e9e86146104085780638280a25a14610439575f80fd5b806335455ded146103745780634b396bcc146103905780635dc1f535146103a357806360305d62146103ab57806360ad7784146103c8575f80fd5b806320d91b7a1161023357806320d91b7a146102ea578063245dafcb146102fd57806325e1c7761461031057806327bf60cd146103235780632aa56638146103365780632e2194d814610349575f80fd5b80630322ed981461026f5780631340964514610284578063151d30d11461029757806316679564146102b65780631af1872b146102c9575b5f80fd5b61028261027d366004614c91565b6105ff565b005b610282610292366004614cbb565b610892565b61029f600a81565b60405160ff90911681526020015b60405180910390f35b6102826102c4366004614cf2565b610a93565b6102dc6102d7366004614f40565b610a9f565b6040519081526020016102ad565b6102826102f836600461501f565b610adc565b61028261030b366004614c91565b611090565b61028261031e366004614cbb565b6112f0565b610282610331366004614cf2565b611364565b610282610344366004614cf2565b611370565b61035c610357366004614c91565b611381565b6040516001600160401b0390911681526020016102ad565b61037d61271081565b60405161ffff90911681526020016102ad565b61028261039e366004615064565b6113d5565b6102dc6113e1565b6103b3601481565b60405163ffffffff90911681526020016102ad565b6102826103d6366004614cbb565b6113f0565b6102dc6103e93660046150b2565b611643565b6102dc5f81565b6102dc6104033660046150cd565b611663565b5f80516020615c51833981519152546001600160a01b03165b6040516001600160a01b0390911681526020016102ad565b61029f603081565b61028261044f366004615064565b61168f565b6102826104623660046150ed565b61169b565b610282610475366004614c91565b61174c565b6102dc61048836600461511b565b6117e0565b61028261049b366004615134565b6118b9565b6102dc6104ae36600461511b565b611998565b6102826104c1366004615064565b611b87565b6102dc7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0081565b6102826104fb366004614cf2565b611b93565b6104216005600160991b0181565b61035c611b9f565b6102dc7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0081565b61028261054b366004614c91565b611bc1565b61029f601481565b610282610566366004615064565b611cdd565b61057e61057936600461511b565b611ce9565b604080519283526001600160401b039091166020830152016102ad565b6105ae6105a9366004614c91565b611e6c565b6040516102ad91906151dc565b61035c6202a30081565b6102dc5f80516020615c5183398151915281565b6102826105e73660046150ed565b611fbb565b6102dc6105fa366004615292565b612052565b5f61060861208a565b5f838152600580830160205260408083208151610100810190925280549495509293909291839160ff169081111561064257610642615167565b600581111561065357610653615167565b8152602001600182018054610667906152fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610693906152fd565b80156106de5780601f106106b5576101008083540402835291602001916106de565b820191905f5260205f20905b8154815290600101906020018083116106c157829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b9092048116608084015260039384015480821660a0850152919091041660c0909101529091508151600581111561075657610756615167565b14610792575f8381526005830160205260409081902054905163170cc93360e21b81526107899160ff1690600401615335565b60405180910390fd5b606081015160405163854a893f60e01b8152600481018590526001600160401b0390911660248201525f60448201526005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__9063854a893f906064015f60405180830381865af4158015610809573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108309190810190615385565b6040518263ffffffff1660e01b815260040161084c91906153b6565b6020604051808303815f875af1158015610868573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061088c91906153c8565b50505050565b61089a6120ae565b5f6108a36120e5565b5f848152600782016020526040808220815160e0810190925280549394509192909190829060ff1660038111156108dc576108dc615167565b60038111156108ed576108ed615167565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c090910152905060038151600381111561096657610966615167565b14610987578051604051633b0d540d60e21b815261078991906004016153df565b5f6109958260400151611e6c565b905060046109a68360400151611e6c565b5160058111156109b8576109b8615167565b141580156109df57508160c001516001600160401b031681608001516001600160401b0316105b15610a6d575f806109ef86611ce9565b9150915081846040015114610a285781846040015160405163fee3144560e01b8152600401610789929190918252602082015260400190565b806001600160401b03168460c001516001600160401b03161115610a6a57604051632e19bc2d60e11b81526001600160401b0382166004820152602401610789565b50505b610a7685612109565b505050610a8f60015f80516020615c7183398151915255565b5050565b61088c8383835f6122ed565b5f610aa86120ae565b610ab88989898989898989612534565b9050610ad060015f80516020615c7183398151915255565b98975050505050505050565b5f610ae561208a565b600781015490915060ff1615610b0e57604051637fab81e560e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b51573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b7591906153c8565b836020013514610b9e576040516372b0a7e760e11b815260208401356004820152602401610789565b30610baf60608501604086016153f9565b6001600160a01b031614610bf257610bcd60608401604085016153f9565b604051632f88120d60e21b81526001600160a01b039091166004820152602401610789565b5f610c006060850185615414565b905090505f805b828163ffffffff161015610e86575f610c236060880188615414565b8363ffffffff16818110610c3957610c39615459565b9050602002810190610c4b919061546d565b610c549061548b565b80516040519192505f916006880191610c6c91615506565b90815260200160405180910390205414610c9c57805160405163a41f772f60e01b815261078991906004016153b6565b805151601414610cc2578051604051633e08a12560e11b815261078991906004016153b6565b5f6002885f013584604051602001610cf192919091825260e01b6001600160e01b031916602082015260240190565b60408051601f1981840301815290829052610d0b91615506565b602060405180830381855afa158015610d26573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d4991906153c8565b90508086600601835f0151604051610d619190615506565b90815260408051918290036020908101909220929092555f8381526005890190915220805460ff191660021781558251600190910190610da1908261555b565b50604082810180515f84815260058a016020529290922060028101805492516001600160401b0394851667ffffffffffffffff60801b90941693909317600160c01b858516021790556003018054429093166001600160801b031990931692909217909155610e10908561562a565b8251602001519094506bffffffffffffffffffffffff1916817f9d9c026e2cadfec89cccc2cd72705360eca1beba24774f3363f4bb33faabc7d78460400151604051610e6b91906001600160401b0391909116815260200190565b60405180910390a3505080610e7f9061564a565b9050610c07565b5060038301805467ffffffffffffffff60401b1916600160401b6001600160401b0384168102919091179091556001840154606491610ec9910460ff168361566c565b6001600160401b03161015610efc57604051633e1a785160e01b81526001600160401b0382166004820152602401610789565b5f73__$fd0c147b4031eef6079b0498cbafa865f0$__634d847884610f20876126bc565b604001516040518263ffffffff1660e01b8152600401610f4091906153b6565b602060405180830381865af4158015610f5b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f7f91906153c8565b90505f73__$fd0c147b4031eef6079b0498cbafa865f0$__6387418b8e886040518263ffffffff1660e01b8152600401610fb991906157c2565b5f60405180830381865af4158015610fd3573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610ffa9190810190615385565b90505f60028260405161100d9190615506565b602060405180830381855afa158015611028573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061104b91906153c8565b90508281146110775760405163baaea89d60e01b81526004810182905260248101849052604401610789565b5050506007909201805460ff1916600117905550505050565b5f6110996120e5565b5f838152600782016020526040808220815160e0810190925280549394509192909190829060ff1660038111156110d2576110d2615167565b60038111156110e3576110e3615167565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101529091508151600381111561115c5761115c615167565b1415801561117d575060038151600381111561117a5761117a615167565b14155b1561119e578051604051633b0d540d60e21b815261078991906004016153df565b5f6111ac8260400151611e6c565b905080606001516001600160401b03165f036111de576040516339b894f960e21b815260048101859052602401610789565b604080830151606083015160a0840151925163854a893f60e01b81526005600160991b019363ee5b48eb9373__$fd0c147b4031eef6079b0498cbafa865f0$__9363854a893f9361124c93906004019283526001600160401b03918216602084015216604082015260600190565b5f60405180830381865af4158015611266573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261128d9190810190615385565b6040518263ffffffff1660e01b81526004016112a991906153b6565b6020604051808303815f875af11580156112c5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112e991906153c8565b5050505050565b6112f9826127d2565b611319576040516330efa98b60e01b815260048101839052602401610789565b5f61132383611e6c565b519050600281600581111561133a5761133a615167565b1461135a578060405163170cc93360e21b81526004016107899190615335565b61088c83836127fb565b61088c8383835f612aa1565b61137c8383835f612dce565b505050565b5f8061138b6120e5565b600301546113999084615856565b90508015806113ae57506001600160401b0381115b156113cf5760405163222d164360e21b815260048101849052602401610789565b92915050565b6112e9848484846122ed565b5f6113ea61208a565b54919050565b5f6113f96120e5565b5f848152600782016020526040808220815160e0810190925280549394509192909190829060ff16600381111561143257611432615167565b600381111561144357611443615167565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f6114b982611e6c565b90506001835160038111156114d0576114d0615167565b146114f1578251604051633b0d540d60e21b815261078991906004016153df565b60048151600581111561150657611506615167565b0361151c5761151486612109565b505050505050565b8260a001516001600160401b031681608001516001600160401b031610156115bc575f8061154987611ce9565b915091508184146115775760405163fee3144560e01b81526004810183905260248101859052604401610789565b8460a001516001600160401b0316816001600160401b031610156115b957604051632e19bc2d60e11b81526001600160401b0382166004820152602401610789565b50505b5f868152600785016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b810267ffffffffffffffff60401b19909216919091179091559151918252839188917f3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa910160405180910390a3505050505050565b5f61164c6120e5565b600301546113cf906001600160401b038416615875565b5f61166c6120ae565b611677833384612dfa565b90506113cf60015f80516020615c7183398151915255565b61088c84848484612dce565b5f6116a46120e5565b90506001600160a01b0382166116d85760405163caa903f960e01b81526001600160a01b0383166004820152602401610789565b5f8381526006820160205260409020546001600160a01b0316331461171e57335b604051636e2ccd7560e11b81526001600160a01b039091166004820152602401610789565b5f928352600b01602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b5f6117556120e5565b90505f61176183611e6c565b519050600481600581111561177857611778615167565b14611798578060405163170cc93360e21b81526004016107899190615335565b5f8381526006830160205260409020546001600160a01b031633146117bd57336116f9565b5f83815260068301602052604090205461137c906001600160a01b03168461303f565b5f6117e96120ae565b5f6117f26120e5565b90505f806117ff85613069565b9150915061180c826127d2565b61181a5750915061189e9050565b5f828152600684016020908152604080832054600b870190925290912080546001600160a01b031981169091556001600160a01b0391821691168061185c5750805b60048351600581111561187157611871615167565b0361188057611880818561303f565b611896826118918560400151611643565b61341e565b509193505050505b6118b460015f80516020615c7183398151915255565b919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff1680611902575080546001600160401b03808416911610155b156119205760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b17815561194b8484613441565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050565b5f806119a261208a565b90505f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63021de88f6119c9876126bc565b604001516040518263ffffffff1660e01b81526004016119e991906153b6565b6040805180830381865af4158015611a03573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a27919061588c565b9150915080611a4d57604051632d07135360e01b81528115156004820152602401610789565b5f82815260048401602052604090208054611a67906152fd565b90505f03611a8b5760405163089938b360e11b815260048101839052602401610789565b60015f838152600580860160205260409091205460ff1690811115611ab257611ab2615167565b14611ae5575f8281526005840160205260409081902054905163170cc93360e21b81526107899160ff1690600401615335565b5f8281526004840160205260408120611afd91614bfd565b5f828152600584016020908152604091829020805460ff1916600290811782556003820180546001600160401b0342811667ffffffffffffffff19909216919091179091559101549251600160c01b90930416825283917f967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06910160405180910390a2509392505050565b6112e984848484612aa1565b61137c8383835f61345b565b5f611ba861208a565b60030154600160401b90046001600160401b0316919050565b5f611bca61208a565b5f8381526004820160205260409020805491925090611be8906152fd565b90505f03611c0c5760405163089938b360e11b815260048101839052602401610789565b60015f838152600580840160205260409091205460ff1690811115611c3357611c33615167565b14611c66575f8281526005820160205260409081902054905163170cc93360e21b81526107899160ff1690600401615335565b5f8281526004808301602052604091829020915163ee5b48eb60e01b81526005600160991b019263ee5b48eb92611c9d92016158af565b6020604051808303815f875af1158015611cb9573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061137c91906153c8565b61088c8484848461345b565b5f805f611cf5846126bc565b90505f805f73__$fd0c147b4031eef6079b0498cbafa865f0$__6350782b0f85604001516040518263ffffffff1660e01b8152600401611d3591906153b6565b606060405180830381865af4158015611d50573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d749190615939565b9250925092505f611d8361208a565b5f8581526005820160205260409020600201549091506001600160401b03808516600160401b909204161015611dd757604051632e19bc2d60e11b81526001600160401b0384166004820152602401610789565b5f8481526005820160205260409081902060020180546001600160401b038616600160801b0267ffffffffffffffff60801b199091161790555184907fc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff990611e5790869086906001600160401b0392831681529116602082015260400190565b60405180910390a25091969095509350505050565b611e74614c34565b5f611e7d61208a565b5f84815260058083016020526040918290208251610100810190935280549394509192839160ff90911690811115611eb757611eb7615167565b6005811115611ec857611ec8615167565b8152602001600182018054611edc906152fd565b80601f0160208091040260200160405190810160405280929190818152602001828054611f08906152fd565b8015611f535780601f10611f2a57610100808354040283529160200191611f53565b820191905f5260205f20905b815481529060010190602001808311611f3657829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b9092048116608084015260039093015480841660a08401520490911660c0909101529392505050565b6001600160a01b038116611fed5760405163caa903f960e01b81526001600160a01b0382166004820152602401610789565b5f611ff66120e5565b5f8481526007820160205260409020549091506001600160a01b0361010090910416331461202457336116f9565b5f928352600901602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b5f8061205c61208a565b9050806006018484604051612072929190615979565b90815260200160405180910390205491505092915050565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0090565b5f80516020615c718339815191528054600119016120df57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0090565b5f6121126120e5565b5f838152600782016020526040808220815160e0810190925280549394509192909190829060ff16600381111561214b5761214b615167565b600381111561215c5761215c615167565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091506121d0613487565b82608001516121df919061562a565b6001600160401b03164210156122135760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610789565b5f848152600784016020908152604080832080546001600160a81b03191681556001810184905560020183905560098601909152902080546001600160a01b031981169091556001600160a01b03168061226e575060208201515b5f8061227b8388866134a2565b9150915061229485602001516118918760600151611643565b6040805183815260208101839052859189917f5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e910160405180910390a350505050505050565b60015f80516020615c7183398151915255565b5f806122f76120e5565b90506123028661354f565b5f61230c87611e6c565b9050612317876127d2565b6123265760019250505061252c565b5f8781526006830160205260409020546001600160a01b0316331461234b57336116f9565b5f87815260068301602052604090205460c082015161237a91600160b01b90046001600160401b03169061562a565b6001600160401b03168160e001516001600160401b031610156123c15760e081015160405163fb6ce63f60e01b81526001600160401b039091166004820152602401610789565b5f86156123d9576123d288876127fb565b90506123f7565b505f8781526006830160205260409020600101546001600160401b03165b600483015460408301515f916001600160a01b031690634f22429f9061241c90611643565b60c086015160e0808801516040519185901b6001600160e01b031916825260048201939093526001600160401b0391821660248201819052604482015291811660648301528516608482015260a401602060405180830381865afa158015612486573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124aa91906153c8565b90506001600160a01b0386166124d6575f8981526006850160205260409020546001600160a01b031695505b5f898152600a85016020526040812080548392906124f5908490615988565b90915550505f898152600b94909401602052604090932080546001600160a01b0319166001600160a01b0387161790555050151590505b949350505050565b5f8061253e6120e5565b600281015490915061ffff600160401b90910481169086161080612567575061271061ffff8616115b1561258b57604051635f12e6c360e11b815261ffff86166004820152602401610789565b60028101546001600160401b0390811690851610156125c7576040516202a06d60e11b81526001600160401b0385166004820152602401610789565b80548310806125d95750806001015483115b156125fa5760405163222d164360e21b815260048101849052602401610789565b5f6126048461383a565b90505f61261082611381565b90505f6126218d8d8d8d8d8761385d565b5f818152600686016020908152604080832080546001600160401b039c909c16600160b01b0267ffffffffffffffff60b01b1961ffff9e909e16600160a01b02336001600160b01b0319909e168e17179d909d169c909c178c556001909b01805467ffffffffffffffff19169055600b9096019095529790932080546001600160a01b03191690961790955550939998505050505050505050565b60408051606080820183525f8083526020830152918101919091526040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa158015612720573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612747919081019061599b565b915091508061276957604051636b2f19e960e01b815260040160405180910390fd5b81511561278f578151604051636ba589a560e01b81526004810191909152602401610789565b60208201516001600160a01b0316156127cb576020820151604051624de75d60e31b81526001600160a01b039091166004820152602401610789565b5092915050565b5f806127dc6120e5565b5f938452600601602052505060409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa158015612846573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261286d919081019061599b565b915091508061288f57604051636b2f19e960e01b815260040160405180910390fd5b5f6128986120e5565b60058101548451919250146128c6578251604051636ba589a560e01b81526004810191909152602401610789565b60208301516001600160a01b031615612902576020830151604051624de75d60e31b81526001600160a01b039091166004820152602401610789565b60208301516001600160a01b03161561293e576020830151604051624de75d60e31b81526001600160a01b039091166004820152602401610789565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63088c246386604001516040518263ffffffff1660e01b815260040161297b91906153b6565b6040805180830381865af4158015612995573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129b99190615a2b565b915091508188146129e75760405163fee3144560e01b81526004810183905260248101899052604401610789565b5f8881526006840160205260409020600101546001600160401b039081169082161115612a78575f888152600684016020908152604091829020600101805467ffffffffffffffff19166001600160401b038516908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a2612a96565b505f8781526006830160205260409020600101546001600160401b03165b979650505050505050565b5f80612aab6120e5565b5f878152600782016020526040808220815160e0810190925280549394509192909190829060ff166003811115612ae457612ae4615167565b6003811115612af557612af5615167565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f612b6b82611e6c565b9050600283516003811115612b8257612b82615167565b14612ba3578251604051633b0d540d60e21b815261078991906004016153df565b60208301516001600160a01b03163314612c3f575f8281526006850160205260409020546001600160a01b03163314612bdc57336116f9565b5f82815260068501602052604090205460c0820151612c0b91600160b01b90046001600160401b03169061562a565b6001600160401b0316421015612c3f5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610789565b600281516005811115612c5457612c54615167565b03612d765760028401546080840151612c76916001600160401b03169061562a565b6001600160401b0316421015612caa5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610789565b8715612cbc57612cba82886127fb565b505b5f8981526007850160205260409020805460ff19166003179055606083015160a0820151612cf5918491612cf09190615a4e565b613c47565b505f8a8152600786016020526040812060020180546001600160401b03909316600160c01b026001600160c01b0390931692909217909155612d3884888c613e11565b9050828a7f5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf60405160405180910390a31515945061252c9350505050565b600481516005811115612d8b57612d8b615167565b03612db257612d9b83878b613e11565b50612da589612109565b600194505050505061252c565b805160405163170cc93360e21b81526107899190600401615335565b612dda84848484612aa1565b61088c57604051631036cf9160e11b815260048101859052602401610789565b5f80612e046120e5565b90505f612e136103578561383a565b90505f612e1f87611e6c565b9050612e2a876127d2565b612e4a576040516330efa98b60e01b815260048101889052602401610789565b600281516005811115612e5f57612e5f615167565b14612e8057805160405163170cc93360e21b81526107899190600401615335565b5f828260a00151612e91919061562a565b905083600201600a9054906101000a90046001600160401b03168260400151612eba919061566c565b6001600160401b0316816001600160401b03161115612ef757604051636d51fe0560e11b81526001600160401b0382166004820152602401610789565b5f80612f038a84613c47565b915091505f8a83604051602001612f3192919091825260c01b6001600160c01b031916602082015260280190565b60408051601f1981840301815291815281516020928301205f81815260078b019093529120805491925060019160ff1916828002179055505f8181526007880160209081526040918290208054610100600160a81b0319166101006001600160a01b038f16908102919091178255600182018f9055600290910180546001600160401b038b81166001600160c01b03199092168217600160801b8a8316908102919091176001600160c01b031690935585519283528916938201939093529283019190915260608201849052908c9083907fdf91f7709a30fda3fc5fc5dc97cb5d5b05e67e193dccaaef3cb332d23fda83d19060800160405180910390a496505050505050505b9392505050565b5f6130486120e5565b5f838152600a820160205260408120805491905590915061088c8482613fff565b5f613072614c34565b5f61307b61208a565b90505f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63021de88f6130a2886126bc565b604001516040518263ffffffff1660e01b81526004016130c291906153b6565b6040805180830381865af41580156130dc573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613100919061588c565b91509150801561312757604051632d07135360e01b81528115156004820152602401610789565b5f8281526005808501602052604080832081516101008101909252805491929091839160ff9091169081111561315f5761315f615167565b600581111561317057613170615167565b8152602001600182018054613184906152fd565b80601f01602080910402602001604051908101604052809291908181526020018280546131b0906152fd565b80156131fb5780601f106131d2576101008083540402835291602001916131fb565b820191905f5260205f20905b8154815290600101906020018083116131de57829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b9092048116608084015260039384015480821660a0850152919091041660c0909101529091508151600581111561327357613273615167565b14158015613294575060018151600581111561329157613291615167565b14155b156132b557805160405163170cc93360e21b81526107899190600401615335565b6003815160058111156132ca576132ca615167565b036132d857600481526132dd565b600581525b8360060181602001516040516132f39190615506565b90815260408051602092819003830190205f90819055858152600587810190935220825181548493839160ff191690600190849081111561333657613336615167565b02179055506020820151600182019061334f908261555b565b506040828101516002830180546060860151608087015160a08801516001600160401b039586166001600160801b031994851617600160401b9387168402176001600160801b0316600160801b928716929092026001600160c01b031691909117600160c01b918616919091021790925560c08601516003909501805460e09097015195841696909116959095179390911602919091179091555183907fafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c905f90a29196919550909350505050565b5f80516020615c5183398151915254610a8f906001600160a01b03168383614075565b6134496140d4565b6134528261411f565b610a8f8161419c565b613467848484846122ed565b61088c57604051635bff683f60e11b815260048101859052602401610789565b5f61349061208a565b600101546001600160401b0316919050565b5f805f6134ad6120e5565b5f86815260088201602052604081208054908290559192509081908015613541575f878152600685016020526040902054612710906134f790600160a01b900461ffff1683615875565b6135019190615856565b91508184600a015f8981526020019081526020015f205f8282546135259190615988565b9091555061353590508282615a6e565b92506135418984613fff565b509097909650945050505050565b5f61355861208a565b5f838152600580830160205260408083208151610100810190925280549495509293909291839160ff169081111561359257613592615167565b60058111156135a3576135a3615167565b81526020016001820180546135b7906152fd565b80601f01602080910402602001604051908101604052809291908181526020018280546135e3906152fd565b801561362e5780601f106136055761010080835404028352916020019161362e565b820191905f5260205f20905b81548152906001019060200180831161361157829003601f168201915b50505091835250506002828101546001600160401b038082166020850152600160401b80830482166040860152600160801b830482166060860152600160c01b9092048116608085015260039094015480851660a08501520490921660c090910152909150815160058111156136a6576136a6615167565b146136d9575f8381526005830160205260409081902054905163170cc93360e21b81526107899160ff1690600401615335565b60038152426001600160401b031660e08201525f83815260058381016020526040909120825181548493839160ff191690600190849081111561371e5761371e615167565b021790555060208201516001820190613737908261555b565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031994851617600160401b9387168402176001600160801b0316600160801b928716929092026001600160c01b031691909117600160c01b918616919091021790925560c08501516003909401805460e090960151948416959091169490941792909116021790555f6137d78482613c47565b915050837fbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273828460a001514260405161382c939291909283526001600160401b03918216602084015216604082015260600190565b60405180910390a250505050565b5f6113cf825f80516020615c51833981519152546001600160a01b031690614204565b5f61386661208a565b6007015460ff1661388a57604051637fab81e560e01b815260040160405180910390fd5b5f61389361208a565b905042866001600160401b03161115806138c257506138b56202a30042615988565b866001600160401b031610155b156138eb57604051635879da1360e11b81526001600160401b0387166004820152602401610789565b60038101546001600160401b039061390e90600160401b90048216858316615988565b111561393857604051633e1a785160e01b81526001600160401b0384166004820152602401610789565b61394185614210565b61394a84614210565b86516030146139715786516040516326475b2f60e11b815260040161078991815260200190565b87516014146139955787604051633e08a12560e11b815260040161078991906153b6565b5f801b81600601896040516139aa9190615506565b908152602001604051809103902054146139d9578760405163a41f772f60e01b815260040161078991906153b6565b6139e3835f61432f565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63eb97ce516040518060e00160405280865f015481526020018d81526020018c81526020018b6001600160401b031681526020018a8152602001898152602001886001600160401b03168152506040518263ffffffff1660e01b8152600401613a639190615ae7565b5f60405180830381865af4158015613a7d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613aa49190810190615b9e565b5f82815260048601602052604090209193509150613ac2828261555b565b5081836006018b604051613ad69190615506565b9081526040519081900360200181209190915563ee5b48eb60e01b81525f906005600160991b019063ee5b48eb90613b129085906004016153b6565b6020604051808303815f875af1158015613b2e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b5291906153c8565b5f8481526005860160205260409020805460ff19166001908117825591925001613b7c8c8261555b565b505f8381526005850160205260409020600281018054600160c01b6001600160401b038a1690810267ffffffffffffffff60801b9092161717905560030180546001600160801b0319169055613bd38b6020015190565b6bffffffffffffffffffffffff1916837f5881be437bdcb008bfa5f20e32d3e335ccf8ab90ef2818852a251625260af35d838c8a604051613c30939291909283526001600160401b03918216602084015216604082015260600190565b60405180910390a350909998505050505050505050565b5f805f613c5261208a565b5f868152600582016020526040902060020154909150600160c01b90046001600160401b0316613c82858261432f565b5f613c8c8761459c565b5f88815260058501602052604080822060020180546001600160c01b0316600160c01b6001600160401b038c811691820292909217909255915163854a893f60e01b8152600481018c905291841660248301526044820152919250906005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__9063854a893f906064015f60405180830381865af4158015613d30573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613d579190810190615385565b6040518263ffffffff1660e01b8152600401613d7391906153b6565b6020604051808303815f875af1158015613d8f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613db391906153c8565b604080516001600160401b038581168252602082018490528a1681830152905191925089917f6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d39181900360600190a2909450925050505b9250929050565b5f80613e1b6120e5565b90505f613e2b8660400151611e6c565b90505f600382516005811115613e4357613e43615167565b1480613e615750600482516005811115613e5f57613e5f615167565b145b15613e71575060e0810151613eae565b600282516005811115613e8657613e86615167565b03613e92575042613eae565b815160405163170cc93360e21b81526107899190600401615335565b86608001516001600160401b0316816001600160401b031611613ed6575f9350505050613038565b600483015460608801515f916001600160a01b031690634f22429f90613efb90611643565b60c086015160808c01516040808e01515f90815260068b0160205281902060010154905160e086901b6001600160e01b031916815260048101949094526001600160401b0392831660248501529082166044840152818716606484015216608482015260a401602060405180830381865afa158015613f7c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613fa091906153c8565b90506001600160a01b038716613fb857876020015196505b5f8681526008850160209081526040808320849055600990960190529390932080546001600160a01b0388166001600160a01b031990911617905550909150509392505050565b5f5f80516020615c5183398151915280546040516340c10f1960e01b81526001600160a01b038681166004830152602482018690529293509116906340c10f19906044015f604051808303815f87803b15801561405a575f80fd5b505af115801561406c573d5f803e3d5ffd5b50505050505050565b6040516001600160a01b0383811660248301526044820183905261137c91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614605565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661411d57604051631afcd79f60e31b815260040160405180910390fd5b565b6141276140d4565b61413081614666565b61413861467f565b6141996060820135608083013561415560c0850160a086016150b2565b61416560e0860160c08701615be1565b614176610100870160e08801615bfa565b61010087013561418e61014089016101208a016153f9565b88610140013561468f565b50565b6141a46140d4565b5f80516020615c518339815191526001600160a01b0382166141e457604051637330680360e01b81526001600160a01b0383166004820152602401610789565b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f613038833384614851565b805163ffffffff16158015614229575060208101515115155b1561425d57805160208201515160405163c08a0f1d60e01b815263ffffffff90921660048301526024820152604401610789565b602081015151815163ffffffff1611156142a057805160208201515160405163c08a0f1d60e01b815263ffffffff90921660048301526024820152604401610789565b60015b816020015151811015610a8f5760208201516142c0600183615a6e565b815181106142d0576142d0615459565b60200260200101516001600160a01b0316826020015182815181106142f7576142f7615459565b60200260200101516001600160a01b0316101561432757604051630dbc8d5f60e31b815260040160405180910390fd5b6001016142a3565b5f61433861208a565b90505f826001600160401b0316846001600160401b031611156143665761435f8385615a4e565b9050614373565b6143708484615a4e565b90505b60408051608081018252600284015480825260038501546001600160401b038082166020850152600160401b8204811694840194909452600160801b90049092166060820152429115806143e05750600184015481516143dc916001600160401b031690615988565b8210155b15614408576001600160401b0380841660608301528282526040820151166020820152614427565b828160600181815161441a919061562a565b6001600160401b03169052505b606081015161443790606461566c565b602082015160018601546001600160401b0392909216916144629190600160401b900460ff1661566c565b6001600160401b0316101561449b57606081015160405163dfae880160e01b81526001600160401b039091166004820152602401610789565b85816040018181516144ad919061562a565b6001600160401b03169052506040810180518691906144cd908390615a4e565b6001600160401b0316905250600184015460408201516064916144fb91600160401b90910460ff169061566c565b6001600160401b03161015614534576040808201519051633e1a785160e01b81526001600160401b039091166004820152602401610789565b8051600285015560208101516003909401805460408301516060909301516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b0319909316919097161717919091169390931790925550505050565b5f806145a661208a565b5f84815260058201602052604090206002018054919250906008906145da90600160401b90046001600160401b0316615c1a565b91906101000a8154816001600160401b0302191690836001600160401b031602179055915050919050565b5f6146196001600160a01b038416836149b4565b905080515f1415801561463d57508080602001905181019061463b9190615c35565b155b1561137c57604051635274afe760e01b81526001600160a01b0384166004820152602401610789565b61466e6140d4565b6146766149c1565b614199816149c9565b6146876140d4565b61411d614aad565b6146976140d4565b5f6146a06120e5565b905061ffff861615806146b8575061271061ffff8716115b156146dc57604051635f12e6c360e11b815261ffff87166004820152602401610789565b878911156147005760405163222d164360e21b8152600481018a9052602401610789565b60ff851615806147135750600a60ff8616115b156147365760405163170db35960e31b815260ff86166004820152602401610789565b61473e613487565b6001600160401b0316876001600160401b0316101561477a576040516202a06d60e11b81526001600160401b0388166004820152602401610789565b835f0361479a5760405163a733007160e01b815260040160405180910390fd5b816147bb57604051632f6bd1db60e01b815260048101839052602401610789565b97885560018801969096556002870180546001600160401b039690961669ffffffffffffffffffff1990961695909517600160401b61ffff95909516949094029390931767ffffffffffffffff60501b191660ff92909216600160501b029190911790925560038401919091556004830180546001600160a01b0319166001600160a01b03909216919091179055600590910155565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015614897573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906148bb91906153c8565b90506148d26001600160a01b038616853086614ab5565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015614916573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061493a91906153c8565b90508181116149a05760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610789565b6149aa8282615a6e565b9695505050505050565b606061303883835f614aee565b61411d6140d4565b6149d16140d4565b5f6149da61208a565b82358155905060146149f26060840160408501615bfa565b60ff161180614a115750614a0c6060830160408401615bfa565b60ff16155b15614a4557614a266060830160408401615bfa565b604051634a59bbff60e11b815260ff9091166004820152602401610789565b614a556060830160408401615bfa565b60018201805460ff92909216600160401b0260ff60401b19909216919091179055614a8660408301602084016150b2565b600191909101805467ffffffffffffffff19166001600160401b0390921691909117905550565b6122da6140d4565b6040516001600160a01b03848116602483015283811660448301526064820183905261088c9186918216906323b872dd906084016140a2565b606081471015614b135760405163cd78605960e01b8152306004820152602401610789565b5f80856001600160a01b03168486604051614b2e9190615506565b5f6040518083038185875af1925050503d805f8114614b68576040519150601f19603f3d011682016040523d82523d5f602084013e614b6d565b606091505b50915091506149aa868383606082614b8d57614b8882614bd4565b613038565b8151158015614ba457506001600160a01b0384163b155b15614bcd57604051639996b31560e01b81526001600160a01b0385166004820152602401610789565b5080613038565b805115614be45780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b508054614c09906152fd565b5f825580601f10614c18575050565b601f0160209004905f5260205f20908101906141999190614c79565b604080516101008101909152805f81526060602082018190525f604083018190529082018190526080820181905260a0820181905260c0820181905260e09091015290565b5b80821115614c8d575f8155600101614c7a565b5090565b5f60208284031215614ca1575f80fd5b5035919050565b803563ffffffff811681146118b4575f80fd5b5f8060408385031215614ccc575f80fd5b82359150614cdc60208401614ca8565b90509250929050565b8015158114614199575f80fd5b5f805f60608486031215614d04575f80fd5b833592506020840135614d1681614ce5565b9150614d2460408501614ca8565b90509250925092565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715614d6357614d63614d2d565b60405290565b604051606081016001600160401b0381118282101715614d6357614d63614d2d565b604051601f8201601f191681016001600160401b0381118282101715614db357614db3614d2d565b604052919050565b5f6001600160401b03821115614dd357614dd3614d2d565b50601f01601f191660200190565b5f82601f830112614df0575f80fd5b8135614e03614dfe82614dbb565b614d8b565b818152846020838601011115614e17575f80fd5b816020850160208301375f918101602001919091529392505050565b6001600160401b0381168114614199575f80fd5b80356118b481614e33565b6001600160a01b0381168114614199575f80fd5b5f60408284031215614e76575f80fd5b614e7e614d41565b9050614e8982614ca8565b81526020808301356001600160401b0380821115614ea5575f80fd5b818501915085601f830112614eb8575f80fd5b813581811115614eca57614eca614d2d565b8060051b9150614edb848301614d8b565b8181529183018401918481019088841115614ef4575f80fd5b938501935b83851015614f1e5784359250614f0e83614e52565b8282529385019390850190614ef9565b808688015250505050505092915050565b803561ffff811681146118b4575f80fd5b5f805f805f805f80610100898b031215614f58575f80fd5b88356001600160401b0380821115614f6e575f80fd5b614f7a8c838d01614de1565b995060208b0135915080821115614f8f575f80fd5b614f9b8c838d01614de1565b9850614fa960408c01614e47565b975060608b0135915080821115614fbe575f80fd5b614fca8c838d01614e66565b965060808b0135915080821115614fdf575f80fd5b50614fec8b828c01614e66565b945050614ffb60a08a01614f2f565b925061500960c08a01614e47565b915060e089013590509295985092959890939650565b5f8060408385031215615030575f80fd5b82356001600160401b03811115615045575f80fd5b830160808186031215615056575f80fd5b9150614cdc60208401614ca8565b5f805f8060808587031215615077575f80fd5b84359350602085013561508981614ce5565b925061509760408601614ca8565b915060608501356150a781614e52565b939692955090935050565b5f602082840312156150c2575f80fd5b813561303881614e33565b5f80604083850312156150de575f80fd5b50508035926020909101359150565b5f80604083850312156150fe575f80fd5b82359150602083013561511081614e52565b809150509250929050565b5f6020828403121561512b575f80fd5b61303882614ca8565b5f80828403610180811215615147575f80fd5b61016080821215615156575f80fd5b849350830135905061511081614e52565b634e487b7160e01b5f52602160045260245ffd5b6006811061518b5761518b615167565b9052565b5f5b838110156151a9578181015183820152602001615191565b50505f910152565b5f81518084526151c881602086016020860161518f565b601f01601f19169290920160200192915050565b602081526151ee60208201835161517b565b5f602083015161010080604085015261520b6101208501836151b1565b915060408501516001600160401b0380821660608701528060608801511660808701525050608085015161524a60a08601826001600160401b03169052565b5060a08501516001600160401b03811660c08601525060c08501516001600160401b03811660e08601525060e08501516001600160401b038116858301525090949350505050565b5f80602083850312156152a3575f80fd5b82356001600160401b03808211156152b9575f80fd5b818501915085601f8301126152cc575f80fd5b8135818111156152da575f80fd5b8660208285010111156152eb575f80fd5b60209290920196919550909350505050565b600181811c9082168061531157607f821691505b60208210810361532f57634e487b7160e01b5f52602260045260245ffd5b50919050565b602081016113cf828461517b565b5f82601f830112615352575f80fd5b8151615360614dfe82614dbb565b818152846020838601011115615374575f80fd5b61252c82602083016020870161518f565b5f60208284031215615395575f80fd5b81516001600160401b038111156153aa575f80fd5b61252c84828501615343565b602081525f61303860208301846151b1565b5f602082840312156153d8575f80fd5b5051919050565b60208101600483106153f3576153f3615167565b91905290565b5f60208284031215615409575f80fd5b813561303881614e52565b5f808335601e19843603018112615429575f80fd5b8301803591506001600160401b03821115615442575f80fd5b6020019150600581901b3603821315613e0a575f80fd5b634e487b7160e01b5f52603260045260245ffd5b5f8235605e19833603018112615481575f80fd5b9190910192915050565b5f6060823603121561549b575f80fd5b6154a3614d69565b82356001600160401b03808211156154b9575f80fd5b6154c536838701614de1565b835260208501359150808211156154da575f80fd5b506154e736828601614de1565b60208301525060408301356154fb81614e33565b604082015292915050565b5f825161548181846020870161518f565b601f82111561137c57805f5260205f20601f840160051c8101602085101561553c5750805b601f840160051c820191505b818110156112e9575f8155600101615548565b81516001600160401b0381111561557457615574614d2d565b6155888161558284546152fd565b84615517565b602080601f8311600181146155bb575f84156155a45750858301515b5f19600386901b1c1916600185901b178555611514565b5f85815260208120601f198616915b828110156155e9578886015182559484019460019091019084016155ca565b508582101561560657878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b038181168382160190808211156127cb576127cb615616565b5f63ffffffff80831681810361566257615662615616565b6001019392505050565b6001600160401b0381811683821602808216919082811461568f5761568f615616565b505092915050565b5f808335601e198436030181126156ac575f80fd5b83016020810192503590506001600160401b038111156156ca575f80fd5b803603821315613e0a575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208086019550808560051b830101845f5b878110156157b557848303601f19018952813536889003605e1901811261573c575f80fd5b8701606061574a8280615697565b82875261575a83880182846156d8565b9250505061576a86830183615697565b8683038888015261577c8382846156d8565b92505050604080830135925061579183614e33565b6001600160401b039290921694909101939093529783019790830190600101615717565b5090979650505050505050565b6020815281356020820152602082013560408201525f60408301356157e681614e52565b6001600160a01b031660608381019190915283013536849003601e1901811261580d575f80fd5b83016020810190356001600160401b03811115615828575f80fd5b8060051b3603821315615839575f80fd5b60808085015261584d60a085018284615700565b95945050505050565b5f8261587057634e487b7160e01b5f52601260045260245ffd5b500490565b80820281158282048414176113cf576113cf615616565b5f806040838503121561589d575f80fd5b82519150602083015161511081614ce5565b5f60208083525f84546158c1816152fd565b806020870152604060018084165f81146158e257600181146158fe5761592b565b60ff19851660408a0152604084151560051b8a0101955061592b565b895f5260205f205f5b858110156159225781548b8201860152908301908801615907565b8a016040019650505b509398975050505050505050565b5f805f6060848603121561594b575f80fd5b83519250602084015161595d81614e33565b604085015190925061596e81614e33565b809150509250925092565b818382375f9101908152919050565b808201808211156113cf576113cf615616565b5f80604083850312156159ac575f80fd5b82516001600160401b03808211156159c2575f80fd5b90840190606082870312156159d5575f80fd5b6159dd614d69565b8251815260208301516159ef81614e52565b6020820152604083015182811115615a05575f80fd5b615a1188828601615343565b604083015250809450505050602083015161511081614ce5565b5f8060408385031215615a3c575f80fd5b82519150602083015161511081614e33565b6001600160401b038281168282160390808211156127cb576127cb615616565b818103818111156113cf576113cf615616565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015615adc5784516001600160a01b03168252938301936001929092019190830190615ab3565b509695505050505050565b60208152815160208201525f602083015160e06040840152615b0d6101008401826151b1565b90506040840151601f1980858403016060860152615b2b83836151b1565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152615b5b8383615a81565b925060a08601519150808584030160c086015250615b798282615a81565b91505060c0840151615b9660e08501826001600160401b03169052565b509392505050565b5f8060408385031215615baf575f80fd5b8251915060208301516001600160401b03811115615bcb575f80fd5b615bd785828601615343565b9150509250929050565b5f60208284031215615bf1575f80fd5b61303882614f2f565b5f60208284031215615c0a575f80fd5b813560ff81168114613038575f80fd5b5f6001600160401b0380831681810361566257615662615616565b5f60208284031215615c45575f80fd5b815161303881614ce556fe6e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220134c522d7ebb5407597523b03383f5d7ff6b691106136ade49f5c163b07b692164736f6c63430008190033",
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

// NODEIDLENGTH is a free data retrieval call binding the contract method 0x63e2ca97.
//
// Solidity: function NODE_ID_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) NODEIDLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "NODE_ID_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NODEIDLENGTH is a free data retrieval call binding the contract method 0x63e2ca97.
//
// Solidity: function NODE_ID_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) NODEIDLENGTH() (uint32, error) {
	return _ERC20TokenStakingManager.Contract.NODEIDLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// NODEIDLENGTH is a free data retrieval call binding the contract method 0x63e2ca97.
//
// Solidity: function NODE_ID_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) NODEIDLENGTH() (uint32, error) {
	return _ERC20TokenStakingManager.Contract.NODEIDLENGTH(&_ERC20TokenStakingManager.CallOpts)
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
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64,uint64))
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
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ERC20TokenStakingManager.Contract.GetValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ERC20TokenStakingManager.Contract.GetValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// L1TotalWeight is a free data retrieval call binding the contract method 0xbb0b1938.
//
// Solidity: function l1TotalWeight() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) L1TotalWeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "l1TotalWeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// L1TotalWeight is a free data retrieval call binding the contract method 0xbb0b1938.
//
// Solidity: function l1TotalWeight() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) L1TotalWeight() (uint64, error) {
	return _ERC20TokenStakingManager.Contract.L1TotalWeight(&_ERC20TokenStakingManager.CallOpts)
}

// L1TotalWeight is a free data retrieval call binding the contract method 0xbb0b1938.
//
// Solidity: function l1TotalWeight() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) L1TotalWeight() (uint64, error) {
	return _ERC20TokenStakingManager.Contract.L1TotalWeight(&_ERC20TokenStakingManager.CallOpts)
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

// SubnetID is a free data retrieval call binding the contract method 0x5dc1f535.
//
// Solidity: function subnetID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) SubnetID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "subnetID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SubnetID is a free data retrieval call binding the contract method 0x5dc1f535.
//
// Solidity: function subnetID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) SubnetID() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.SubnetID(&_ERC20TokenStakingManager.CallOpts)
}

// SubnetID is a free data retrieval call binding the contract method 0x5dc1f535.
//
// Solidity: function subnetID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) SubnetID() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.SubnetID(&_ERC20TokenStakingManager.CallOpts)
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

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeDelegatorRemoval", delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32, uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteValidatorWeightUpdate(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeValidatorWeightUpdate", messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32, uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteValidatorWeightUpdate(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorWeightUpdate(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32, uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteValidatorWeightUpdate(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorWeightUpdate(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval0 is a paid mutator transaction binding the contract method 0xaac80c39.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitiateDelegatorRemoval0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitiateDelegatorRemoval0", delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitiateDelegatorRemoval0 is a paid mutator transaction binding the contract method 0xaac80c39.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitiateDelegatorRemoval0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateDelegatorRemoval0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitiateDelegatorRemoval0 is a paid mutator transaction binding the contract method 0xaac80c39.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitiateDelegatorRemoval0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateDelegatorRemoval0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval0 is a paid mutator transaction binding the contract method 0x4b396bcc.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitiateValidatorRemoval0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitiateValidatorRemoval0", validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitiateValidatorRemoval0 is a paid mutator transaction binding the contract method 0x4b396bcc.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitiateValidatorRemoval0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateValidatorRemoval0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// ForceInitiateValidatorRemoval0 is a paid mutator transaction binding the contract method 0x4b396bcc.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitiateValidatorRemoval0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateValidatorRemoval0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
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

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0x739b4eff.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateDelegatorRegistration", validationID, delegationAmount)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0x739b4eff.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0x739b4eff.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval0 is a paid mutator transaction binding the contract method 0x8af5499e.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateDelegatorRemoval0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateDelegatorRemoval0", delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitiateDelegatorRemoval0 is a paid mutator transaction binding the contract method 0x8af5499e.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateDelegatorRemoval0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRemoval0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitiateDelegatorRemoval0 is a paid mutator transaction binding the contract method 0x8af5499e.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateDelegatorRemoval0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRemoval0(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1af1872b.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, uint64 registrationExpiry, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, registrationExpiry uint64, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, registrationExpiry, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1af1872b.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, uint64 registrationExpiry, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, registrationExpiry uint64, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, nodeID, blsPublicKey, registrationExpiry, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1af1872b.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, uint64 registrationExpiry, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, registrationExpiry uint64, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, nodeID, blsPublicKey, registrationExpiry, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval0 is a paid mutator transaction binding the contract method 0xcc71bbba.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateValidatorRemoval0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateValidatorRemoval0", validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitiateValidatorRemoval0 is a paid mutator transaction binding the contract method 0xcc71bbba.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateValidatorRemoval0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRemoval0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
}

// InitiateValidatorRemoval0 is a paid mutator transaction binding the contract method 0xcc71bbba.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateValidatorRemoval0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRemoval0(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex, rewardRecipient)
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

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendUpdateDelegator(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendUpdateDelegator", delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegator(&_ERC20TokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegator(&_ERC20TokenStakingManager.TransactOpts, delegationID)
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

// ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator is returned from FilterCompletedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRegistration events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator struct {
	Event *ERC20TokenStakingManagerCompletedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerCompletedDelegatorRegistration)
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
		it.Event = new(ERC20TokenStakingManagerCompletedDelegatorRegistration)
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
func (it *ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerCompletedDelegatorRegistration represents a CompletedDelegatorRegistration event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedDelegatorRegistration struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRegistration is a free log retrieval operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterCompletedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator{contract: _ERC20TokenStakingManager.contract, event: "CompletedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRegistration is a free log subscription operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchCompletedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerCompletedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerCompletedDelegatorRegistration)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
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

// ParseCompletedDelegatorRegistration is a log parse operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseCompletedDelegatorRegistration(log types.Log) (*ERC20TokenStakingManagerCompletedDelegatorRegistration, error) {
	event := new(ERC20TokenStakingManagerCompletedDelegatorRegistration)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerCompletedDelegatorRemovalIterator is returned from FilterCompletedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRemoval events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedDelegatorRemovalIterator struct {
	Event *ERC20TokenStakingManagerCompletedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerCompletedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerCompletedDelegatorRemoval)
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
		it.Event = new(ERC20TokenStakingManagerCompletedDelegatorRemoval)
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
func (it *ERC20TokenStakingManagerCompletedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerCompletedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerCompletedDelegatorRemoval represents a CompletedDelegatorRemoval event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterCompletedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerCompletedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCompletedDelegatorRemovalIterator{contract: _ERC20TokenStakingManager.contract, event: "CompletedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRemoval is a free log subscription operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchCompletedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerCompletedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerCompletedDelegatorRemoval)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
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

// ParseCompletedDelegatorRemoval is a log parse operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseCompletedDelegatorRemoval(log types.Log) (*ERC20TokenStakingManagerCompletedDelegatorRemoval, error) {
	event := new(ERC20TokenStakingManagerCompletedDelegatorRemoval)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerCompletedValidatorRegistrationIterator is returned from FilterCompletedValidatorRegistration and is used to iterate over the raw logs and unpacked data for CompletedValidatorRegistration events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedValidatorRegistrationIterator struct {
	Event *ERC20TokenStakingManagerCompletedValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerCompletedValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerCompletedValidatorRegistration)
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
		it.Event = new(ERC20TokenStakingManagerCompletedValidatorRegistration)
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
func (it *ERC20TokenStakingManagerCompletedValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerCompletedValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerCompletedValidatorRegistration represents a CompletedValidatorRegistration event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedValidatorRegistration struct {
	ValidationID [32]byte
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedValidatorRegistration is a free log retrieval operation binding the contract event 0x967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06.
//
// Solidity: event CompletedValidatorRegistration(bytes32 indexed validationID, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterCompletedValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerCompletedValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "CompletedValidatorRegistration", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCompletedValidatorRegistrationIterator{contract: _ERC20TokenStakingManager.contract, event: "CompletedValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchCompletedValidatorRegistration is a free log subscription operation binding the contract event 0x967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06.
//
// Solidity: event CompletedValidatorRegistration(bytes32 indexed validationID, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchCompletedValidatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerCompletedValidatorRegistration, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "CompletedValidatorRegistration", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerCompletedValidatorRegistration)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedValidatorRegistration", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseCompletedValidatorRegistration(log types.Log) (*ERC20TokenStakingManagerCompletedValidatorRegistration, error) {
	event := new(ERC20TokenStakingManagerCompletedValidatorRegistration)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerCompletedValidatorRemovalIterator is returned from FilterCompletedValidatorRemoval and is used to iterate over the raw logs and unpacked data for CompletedValidatorRemoval events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedValidatorRemovalIterator struct {
	Event *ERC20TokenStakingManagerCompletedValidatorRemoval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerCompletedValidatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerCompletedValidatorRemoval)
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
		it.Event = new(ERC20TokenStakingManagerCompletedValidatorRemoval)
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
func (it *ERC20TokenStakingManagerCompletedValidatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerCompletedValidatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerCompletedValidatorRemoval represents a CompletedValidatorRemoval event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedValidatorRemoval struct {
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedValidatorRemoval is a free log retrieval operation binding the contract event 0xafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c.
//
// Solidity: event CompletedValidatorRemoval(bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterCompletedValidatorRemoval(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerCompletedValidatorRemovalIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "CompletedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCompletedValidatorRemovalIterator{contract: _ERC20TokenStakingManager.contract, event: "CompletedValidatorRemoval", logs: logs, sub: sub}, nil
}

// WatchCompletedValidatorRemoval is a free log subscription operation binding the contract event 0xafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c.
//
// Solidity: event CompletedValidatorRemoval(bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchCompletedValidatorRemoval(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerCompletedValidatorRemoval, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "CompletedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerCompletedValidatorRemoval)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedValidatorRemoval", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseCompletedValidatorRemoval(log types.Log) (*ERC20TokenStakingManagerCompletedValidatorRemoval, error) {
	event := new(ERC20TokenStakingManagerCompletedValidatorRemoval)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedValidatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerCompletedValidatorWeightUpdateIterator is returned from FilterCompletedValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for CompletedValidatorWeightUpdate events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedValidatorWeightUpdateIterator struct {
	Event *ERC20TokenStakingManagerCompletedValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerCompletedValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerCompletedValidatorWeightUpdate)
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
		it.Event = new(ERC20TokenStakingManagerCompletedValidatorWeightUpdate)
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
func (it *ERC20TokenStakingManagerCompletedValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerCompletedValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerCompletedValidatorWeightUpdate represents a CompletedValidatorWeightUpdate event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedValidatorWeightUpdate struct {
	ValidationID [32]byte
	Nonce        uint64
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedValidatorWeightUpdate is a free log retrieval operation binding the contract event 0xc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff9.
//
// Solidity: event CompletedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterCompletedValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerCompletedValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "CompletedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCompletedValidatorWeightUpdateIterator{contract: _ERC20TokenStakingManager.contract, event: "CompletedValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchCompletedValidatorWeightUpdate is a free log subscription operation binding the contract event 0xc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff9.
//
// Solidity: event CompletedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchCompletedValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerCompletedValidatorWeightUpdate, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "CompletedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerCompletedValidatorWeightUpdate)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedValidatorWeightUpdate", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseCompletedValidatorWeightUpdate(log types.Log) (*ERC20TokenStakingManagerCompletedValidatorWeightUpdate, error) {
	event := new(ERC20TokenStakingManagerCompletedValidatorWeightUpdate)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedValidatorWeightUpdate", log); err != nil {
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

// ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator is returned from FilterInitiatedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRegistration events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator struct {
	Event *ERC20TokenStakingManagerInitiatedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitiatedDelegatorRegistration)
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
		it.Event = new(ERC20TokenStakingManagerInitiatedDelegatorRegistration)
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
func (it *ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitiatedDelegatorRegistration represents a InitiatedDelegatorRegistration event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedDelegatorRegistration struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitiatedDelegatorRegistration is a free log retrieval operation binding the contract event 0xdf91f7709a30fda3fc5fc5dc97cb5d5b05e67e193dccaaef3cb332d23fda83d1.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitiatedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator{contract: _ERC20TokenStakingManager.contract, event: "InitiatedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRegistration is a free log subscription operation binding the contract event 0xdf91f7709a30fda3fc5fc5dc97cb5d5b05e67e193dccaaef3cb332d23fda83d1.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitiatedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitiatedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitiatedDelegatorRegistration)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
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

// ParseInitiatedDelegatorRegistration is a log parse operation binding the contract event 0xdf91f7709a30fda3fc5fc5dc97cb5d5b05e67e193dccaaef3cb332d23fda83d1.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitiatedDelegatorRegistration(log types.Log) (*ERC20TokenStakingManagerInitiatedDelegatorRegistration, error) {
	event := new(ERC20TokenStakingManagerInitiatedDelegatorRegistration)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator is returned from FilterInitiatedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRemoval events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator struct {
	Event *ERC20TokenStakingManagerInitiatedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitiatedDelegatorRemoval)
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
		it.Event = new(ERC20TokenStakingManagerInitiatedDelegatorRemoval)
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
func (it *ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitiatedDelegatorRemoval represents a InitiatedDelegatorRemoval event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiatedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitiatedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator{contract: _ERC20TokenStakingManager.contract, event: "InitiatedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRemoval is a free log subscription operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitiatedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitiatedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitiatedDelegatorRemoval)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
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

// ParseInitiatedDelegatorRemoval is a log parse operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitiatedDelegatorRemoval(log types.Log) (*ERC20TokenStakingManagerInitiatedDelegatorRemoval, error) {
	event := new(ERC20TokenStakingManagerInitiatedDelegatorRemoval)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitiatedValidatorRegistrationIterator is returned from FilterInitiatedValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedValidatorRegistration events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedValidatorRegistrationIterator struct {
	Event *ERC20TokenStakingManagerInitiatedValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitiatedValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitiatedValidatorRegistration)
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
		it.Event = new(ERC20TokenStakingManagerInitiatedValidatorRegistration)
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
func (it *ERC20TokenStakingManagerInitiatedValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitiatedValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitiatedValidatorRegistration represents a InitiatedValidatorRegistration event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedValidatorRegistration struct {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitiatedValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][20]byte) (*ERC20TokenStakingManagerInitiatedValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitiatedValidatorRegistration", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitiatedValidatorRegistrationIterator{contract: _ERC20TokenStakingManager.contract, event: "InitiatedValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedValidatorRegistration is a free log subscription operation binding the contract event 0x5881be437bdcb008bfa5f20e32d3e335ccf8ab90ef2818852a251625260af35d.
//
// Solidity: event InitiatedValidatorRegistration(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 registrationMessageID, uint64 registrationExpiry, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitiatedValidatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitiatedValidatorRegistration, validationID [][32]byte, nodeID [][20]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitiatedValidatorRegistration", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitiatedValidatorRegistration)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorRegistration", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitiatedValidatorRegistration(log types.Log) (*ERC20TokenStakingManagerInitiatedValidatorRegistration, error) {
	event := new(ERC20TokenStakingManagerInitiatedValidatorRegistration)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitiatedValidatorRemovalIterator is returned from FilterInitiatedValidatorRemoval and is used to iterate over the raw logs and unpacked data for InitiatedValidatorRemoval events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedValidatorRemovalIterator struct {
	Event *ERC20TokenStakingManagerInitiatedValidatorRemoval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitiatedValidatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitiatedValidatorRemoval)
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
		it.Event = new(ERC20TokenStakingManagerInitiatedValidatorRemoval)
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
func (it *ERC20TokenStakingManagerInitiatedValidatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitiatedValidatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitiatedValidatorRemoval represents a InitiatedValidatorRemoval event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedValidatorRemoval struct {
	ValidationID             [32]byte
	ValidatorWeightMessageID [32]byte
	Weight                   uint64
	EndTime                  uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterInitiatedValidatorRemoval is a free log retrieval operation binding the contract event 0xbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273.
//
// Solidity: event InitiatedValidatorRemoval(bytes32 indexed validationID, bytes32 validatorWeightMessageID, uint64 weight, uint64 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitiatedValidatorRemoval(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerInitiatedValidatorRemovalIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitiatedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitiatedValidatorRemovalIterator{contract: _ERC20TokenStakingManager.contract, event: "InitiatedValidatorRemoval", logs: logs, sub: sub}, nil
}

// WatchInitiatedValidatorRemoval is a free log subscription operation binding the contract event 0xbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273.
//
// Solidity: event InitiatedValidatorRemoval(bytes32 indexed validationID, bytes32 validatorWeightMessageID, uint64 weight, uint64 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitiatedValidatorRemoval(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitiatedValidatorRemoval, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitiatedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitiatedValidatorRemoval)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorRemoval", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitiatedValidatorRemoval(log types.Log) (*ERC20TokenStakingManagerInitiatedValidatorRemoval, error) {
	event := new(ERC20TokenStakingManagerInitiatedValidatorRemoval)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitiatedValidatorWeightUpdateIterator is returned from FilterInitiatedValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for InitiatedValidatorWeightUpdate events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedValidatorWeightUpdateIterator struct {
	Event *ERC20TokenStakingManagerInitiatedValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitiatedValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitiatedValidatorWeightUpdate)
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
		it.Event = new(ERC20TokenStakingManagerInitiatedValidatorWeightUpdate)
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
func (it *ERC20TokenStakingManagerInitiatedValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitiatedValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitiatedValidatorWeightUpdate represents a InitiatedValidatorWeightUpdate event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedValidatorWeightUpdate struct {
	ValidationID          [32]byte
	Nonce                 uint64
	WeightUpdateMessageID [32]byte
	Weight                uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterInitiatedValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d3.
//
// Solidity: event InitiatedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, bytes32 weightUpdateMessageID, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitiatedValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerInitiatedValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitiatedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitiatedValidatorWeightUpdateIterator{contract: _ERC20TokenStakingManager.contract, event: "InitiatedValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchInitiatedValidatorWeightUpdate is a free log subscription operation binding the contract event 0x6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d3.
//
// Solidity: event InitiatedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, bytes32 weightUpdateMessageID, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitiatedValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitiatedValidatorWeightUpdate, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitiatedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitiatedValidatorWeightUpdate)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorWeightUpdate", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitiatedValidatorWeightUpdate(log types.Log) (*ERC20TokenStakingManagerInitiatedValidatorWeightUpdate, error) {
	event := new(ERC20TokenStakingManagerInitiatedValidatorWeightUpdate)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerRegisteredInitialValidatorIterator is returned from FilterRegisteredInitialValidator and is used to iterate over the raw logs and unpacked data for RegisteredInitialValidator events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerRegisteredInitialValidatorIterator struct {
	Event *ERC20TokenStakingManagerRegisteredInitialValidator // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerRegisteredInitialValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerRegisteredInitialValidator)
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
		it.Event = new(ERC20TokenStakingManagerRegisteredInitialValidator)
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
func (it *ERC20TokenStakingManagerRegisteredInitialValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerRegisteredInitialValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerRegisteredInitialValidator represents a RegisteredInitialValidator event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerRegisteredInitialValidator struct {
	ValidationID [32]byte
	NodeID       [20]byte
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredInitialValidator is a free log retrieval operation binding the contract event 0x9d9c026e2cadfec89cccc2cd72705360eca1beba24774f3363f4bb33faabc7d7.
//
// Solidity: event RegisteredInitialValidator(bytes32 indexed validationID, bytes20 indexed nodeID, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterRegisteredInitialValidator(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][20]byte) (*ERC20TokenStakingManagerRegisteredInitialValidatorIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "RegisteredInitialValidator", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerRegisteredInitialValidatorIterator{contract: _ERC20TokenStakingManager.contract, event: "RegisteredInitialValidator", logs: logs, sub: sub}, nil
}

// WatchRegisteredInitialValidator is a free log subscription operation binding the contract event 0x9d9c026e2cadfec89cccc2cd72705360eca1beba24774f3363f4bb33faabc7d7.
//
// Solidity: event RegisteredInitialValidator(bytes32 indexed validationID, bytes20 indexed nodeID, uint64 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchRegisteredInitialValidator(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerRegisteredInitialValidator, validationID [][32]byte, nodeID [][20]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "RegisteredInitialValidator", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerRegisteredInitialValidator)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "RegisteredInitialValidator", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseRegisteredInitialValidator(log types.Log) (*ERC20TokenStakingManagerRegisteredInitialValidator, error) {
	event := new(ERC20TokenStakingManagerRegisteredInitialValidator)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "RegisteredInitialValidator", log); err != nil {
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
