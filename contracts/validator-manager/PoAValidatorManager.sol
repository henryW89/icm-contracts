// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManager} from "./ValidatorManager.sol";
import {IPoAValidatorManager, PChainOwner} from "./interfaces/IPoAValidatorManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

/**
 * @dev Implementation of the {IPoAValidatorManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract PoAValidatorManager is IPoAValidatorManager, OwnableUpgradeable {
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.PoAValidatorManager
    struct PoAValidatorManagerStorage {
        ValidatorManager _manager;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.PoAValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant POA_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x81773fca73a14ca21edf1cadc6ec0b26d6a44966f6e97607e90422658d423500;

    // solhint-disable ordering
    function _getPoAValidatorManagerStorage()
        private
        pure
        returns (PoAValidatorManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := POA_VALIDATOR_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    function initialize(ValidatorManager manager, address initialOwner) external initializer {
        __PoAValidatorManager_init(manager, initialOwner);
    }

    // solhint-disable func-name-mixedcase, ordering
    function __PoAValidatorManager_init(
        ValidatorManager manager,
        address initialOwner
    ) internal onlyInitializing {
        __Ownable_init(initialOwner);
        __PoAValidatorManager_init_unchained(manager);
    }

    // solhint-disable-next-line no-empty-blocks
    function __PoAValidatorManager_init_unchained(ValidatorManager manager)
        internal
        onlyInitializing
    {
        PoAValidatorManagerStorage storage $ = _getPoAValidatorManagerStorage();
        $._manager = manager;
    }

    // solhint-enable func-name-mixedcase

    /**
     * @notice See {IPoAValidatorManager-initiateValidatorRegistration}.
     */
    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        uint64 registrationExpiry,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) external onlyOwner returns (bytes32 validationID) {
        return _getPoAValidatorManagerStorage()._manager.initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            registrationExpiry: registrationExpiry,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            weight: weight
        });
    }

    // solhint-enable ordering
    /**
     * @notice See {IPoAValidatorManager-initiateValidatorRemoval}.
     */
    function initiateValidatorRemoval(bytes32 validationID) external override onlyOwner {
        _getPoAValidatorManagerStorage()._manager.initiateValidatorRemoval(validationID);
    }

    /**
     * @notice Completes validator removal by forwarding to the validator manager.
     */
    function completeValidatorRemoval(uint32 messageIndex) public returns (bytes32) {
        return _getPoAValidatorManagerStorage()._manager.completeValidatorRemoval(messageIndex);
    }

    function completeValidatorRegistration(uint32 messageIndex) public returns (bytes32) {
        return _getPoAValidatorManagerStorage()._manager.completeValidatorRegistration(messageIndex);
    }
}
