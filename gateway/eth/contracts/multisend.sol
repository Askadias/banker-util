pragma solidity >= 0.4.21 < 0.7.0;

import "openzeppelin-contracts/blob/v2.4.0/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/blob/v2.4.0/contracts/token/ERC20/SafeERC20.sol";

contract MultiSend {

    using SafeERC20 for IERC20;
    using SafeMath for uint256;

    address public owner;

    constructor() public payable{
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function bulkSendEth(address payable[] memory addresses, uint256[] memory amounts) public payable returns (bool success){
        uint total = 0;
        for (uint8 i = 0; i < amounts.length; i++) {
            total = total.add(amounts[i]);
        }

        require(msg.value >= total);

        //transfer to each address
        for (uint8 j = 0; j < addresses.length; j++) {
            addresses[j].transfer(amounts[j]);
        }

        //return change to the sender
        if (msg.value > total) {
            uint change = msg.value.sub(total);
            msg.sender.transfer(change);
        }
        return true;
    }

    function bulkSendToken(address asset, address payable[] memory addresses, uint256[] memory amounts) public payable returns (bool success){
        IERC20 token = IERC20(asset);
        uint total = 0;
        for (uint8 i = 0; i < amounts.length; i++) {
            total = total.add(amounts[i]);
        }
        token.safeApprove(address(this), uint256(-1));

        // transfer token to addresses
        for (uint8 j = 0; j < addresses.length; j++) {
            token.safeTransferFrom(msg.sender, addresses[j], amounts[j]);
        }
        // transfer change back to the sender
        if (msg.value > 0) {
            msg.sender.transfer(msg.value);
        }
        return true;
    }
}
