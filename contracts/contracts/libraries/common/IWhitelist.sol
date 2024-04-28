// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

interface IWhitelist {
    /// @notice Emitted when account whitelist status changed.
    /// @param _account The address of account whose status is changed.
    /// @param _status The current whitelist status.
    event WhitelistStatusChanged(address indexed _account, bool _status);

    /// @notice Check whether the sender is allowed to do something.
    /// @param _sender The address of sender.
    function isSenderAllowed(address _sender) external view returns (bool);
}
