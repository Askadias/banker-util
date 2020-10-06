pragma solidity >= 0.4.21 < 0.7.0;

import "@openzeppelin/contracts-ethereum-package/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-ethereum-package/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/upgrades/contracts/Initializable.sol";

contract MultiSend is Initializable {

    using SafeERC20 for IERC20;
    using SafeMath for uint256;

    address public _owner;

    event Transfer(address recipient, uint amount);
    event Refund(uint amount);

    modifier onlyOwner() {
        require(msg.sender == _owner);
        _;
    }

    function initialize(address _admin) public initializer {
        _owner = _admin;
    }

    function bulkSendEth(address payable[] memory addresses, uint256[] memory amounts) public payable returns (bool success) {
        require(addresses.length <= 300, "number of recipients is larger than 300");
        require(addresses.length == amounts.length, "parameters not match");

        uint total = 0;
        for (uint8 i = 0; i < amounts.length; i++) {
            total = total.add(amounts[i]);
        }

        require(msg.value >= total, "not enough ETH balance");

        //transfer to each address
        for (uint8 j = 0; j < addresses.length; j++) {
            addresses[j].transfer(amounts[j]);
            emit Transfer(addresses[j], amounts[j]);
        }

        //return change to the sender
        if (msg.value > total) {
            uint change = msg.value.sub(total);
            msg.sender.transfer(change);
            emit Refund(change);
        }
        return true;
    }

    function bulkSendToken(IERC20 token, address[] memory addresses, uint256[] memory amounts) public payable returns (bool success) {
        require(addresses.length <= 300, "number of recipients is larger than 300");
        require(addresses.length == amounts.length, "parameters not match");
        uint total = 0;
        for (uint8 i = 0; i < amounts.length; i++) {
            total = total.add(amounts[i]);
        }

        address payable from = msg.sender;
        require(token.allowance(from, address(this)) >= total, "not enough USDT balance");

        // transfer token to addresses
        for (uint8 j = 0; j < addresses.length; j++) {
            address payable destination = address(uint160(addresses[j])); // cast the address to payable
            token.safeTransferFrom(from, destination, amounts[j]);
            emit Transfer(destination, amounts[j]);
        }
        // transfer accidentally sent ETH back to sender
        if (msg.value > 0) {
            from.transfer(msg.value);
            emit Refund(msg.value);
        }
        return true;
    }
}
