// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorStatus, ConversionData, PChainOwner} from "../ACP99Manager.sol";

/**
 * @dev Describes the current churn period
 */
struct ValidatorChurnPeriod {
    uint256 startTime;
    uint64 initialWeight;
    uint64 totalWeight;
    uint64 churnAmount;
}

/**
 * @notice Validator Manager settings, used to initialize the Validator Manager
 * @notice The subnetID is the ID of the L1 that the Validator Manager is managing
 * @notice The churnPeriodSeconds is the duration of the churn period in seconds
 * @notice The maximumChurnPercentage is the maximum percentage of the total weight that can be added or removed in a single churn period
 */
struct ValidatorManagerSettings {
    bytes32 subnetID;
    uint64 churnPeriodSeconds;
    uint8 maximumChurnPercentage;
}

/**
 * @dev Specifies a validator to register.
 */
struct ValidatorRegistrationInput {
    bytes nodeID;
    bytes blsPublicKey;
    uint64 registrationExpiry;
    PChainOwner remainingBalanceOwner;
    PChainOwner disableOwner;
}

/**
 * @notice Interface for Validator Manager contracts that implement Subnet-only Validator management.
 */
interface IValidatorManager {

    /**
     * @notice Event emitted when validator weight is updated.
     * @param validationID The ID of the validation period being updated
     * @param nonce The message nonce used to update the validator weight
     * @param weight The updated validator weight that is sent to the P-Chain
     * @param setWeightMessageID The ID of the ICM message that updates the validator's weight on the P-Chain
     */
    event ValidatorWeightUpdate(
        bytes32 indexed validationID,
        uint64 indexed nonce,
        uint64 weight,
        bytes32 setWeightMessageID
    );

    /**
     * @notice Resubmits a validator registration message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being registered.
     */
    function resendRegisterValidatorMessage(bytes32 validationID) external;

    /**
     * @notice Resubmits a validator end message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being ended.
     */
    function resendEndValidatorMessage(bytes32 validationID) external;
}
