pragma solidity >=0.5.0 <0.7.0;

import "token.sol";
import "math.sol";

contract Multisend {
    using SafeMath for uint256;
    address payable public owner;

    event Transfer(address recipient, uint amount);
    event Refund(uint refund);

    constructor() public payable {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function bulkSendEth(address payable[] memory recipients, uint256[] memory amounts) public payable {
        require(recipients.length <= 300, "number of recipients is larger than 300");
        require(recipients.length == amounts.length, "parameters not match");
        uint totalAmount = 0;
        for (uint8 i = 0; i < recipients.length; i++) {
            totalAmount = totalAmount.add(amounts[i]);
        }
        require(msg.value >= totalAmount, "not enough token");
        for (uint8 i = 0; i < recipients.length; i++) {
            recipients[i].transfer(amounts[i]);
            emit Transfer(recipients[i], amounts[i]);
        }
        if (msg.value > totalAmount) {
            uint refund = msg.value.sub(totalAmount);
            msg.sender.transfer(refund);
            emit Refund(refund);
        }
    }

    function bulkSendToken(Token tokenAddr, address payable[] memory recipients, uint256[] memory amounts) public payable {
        require(recipients.length <= 300, "number of recipients is larger than 300");
        require(recipients.length == amounts.length, "parameters not match");
        uint totalAmount = 0;
        for (uint8 i = 0; i < recipients.length; i++) {
            totalAmount = totalAmount.add(amounts[i]);
        }

        address multisendContractAddress = address(this);
        // check if user has enough balance
        require(totalAmount <= tokenAddr.allowance(msg.sender, multisendContractAddress), "not enough token balance");

        // transfer token to addresses
        for (uint8 j = 0; j < recipients.length; j++) {
            tokenAddr.transferFrom(msg.sender, recipients[j], amounts[j]);
            emit Transfer(recipients[j], amounts[j]);
        }
    }

    function destroy(address payable _to) public onlyOwner {
        selfdestruct(_to);
    }
}