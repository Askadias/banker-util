pragma solidity >=0.5.0 <0.7.0;

import "math.sol";
import "token.sol";

contract Transaction {
    using SafeMath for uint256;

    event Transfer(address recipient, uint amount);
    event Refund(uint refund);

    address payable public owner;
    uint8 public feePerc;
    uint8 public internalFeePerc;
    uint8 public rewardPerc;
    address payable[] partners;
    uint8[] partnersPerc;
    uint8[] affiliatesPerc;

    constructor() public payable{
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function bulkSendEth(address payable target, address payable[] memory affiliates) public payable {
        require(affiliates.length <= affiliatesPerc.length, "too many affiliates");

        uint256 amount = msg.value;
        uint256 commission = amount.mul(feePerc).div(100);

        require(commission < tx.gasprice, "not enough gas to cover commission cost");

        uint256 internalFee = commission.mul(internalFeePerc).div(100);
        uint256 reward = commission.mul(rewardPerc).div(100);

        uint256 targetAmount = amount.sub(commission);
        target.transfer(targetAmount);
        emit Transfer(target, targetAmount);

        uint256 rest = amount.sub(internalFee);

        for (uint8 i = 0; i < affiliates.length; i++) {
            uint256 rewardAmount = reward.mul(affiliatesPerc[i]).div(100);
            rest = rest.sub(rewardAmount);
            affiliates[i].transfer(rewardAmount);
            emit Transfer(affiliates[i], rewardAmount);
        }
        for (uint8 j = 0; j < partnersPerc.length; j++) {
            uint256 dividendAmount = rest.mul(partnersPerc[j]).div(100);
            partners[j].transfer(dividendAmount);
            emit Transfer(partners[j], dividendAmount);
        }
    }

    function bulkSendToken(Token tokenAddr, uint256 amount, address payable target, address payable[] memory affiliates) public payable {
        require(affiliates.length <= affiliatesPerc.length, "too many affiliates");
        // check if user has enough balance
        require(amount <= tokenAddr.allowance(msg.sender, address(this)), "not enough balance");

        uint256 commission = amount.mul(feePerc).div(100);

        uint256 internalFee = commission.mul(internalFeePerc).div(100);
        uint256 reward = commission.mul(rewardPerc).div(100);

        uint256 targetAmount = amount.sub(commission);
        target.transfer(targetAmount);
        emit Transfer(target, targetAmount);

        uint256 rest = amount.sub(internalFee);

        for (uint8 i = 0; i < affiliates.length; i++) {
            uint256 rewardAmount = reward.mul(affiliatesPerc[i]).div(100);
            rest = rest.sub(rewardAmount);
            tokenAddr.transferFrom(msg.sender, affiliates[i], rewardAmount);
            emit Transfer(affiliates[i], rewardAmount);
        }
        for (uint8 j = 0; j < partnersPerc.length; j++) {
            uint256 dividendAmount = rest.mul(partnersPerc[j]).div(100);
            tokenAddr.transferFrom(msg.sender, partners[j], dividendAmount);
            emit Transfer(partners[j], dividendAmount);
        }
    }

    function getBalance(address payable addr) public view returns (uint value){
        return addr.balance;
    }

    function deposit() payable public returns (bool){
        return true;
    }

    function withdrawEther(address payable addr, uint amount) public onlyOwner returns (bool success){
        addr.transfer(amount);
        return true;
    }

    function withdrawToken(Token tokenAddr, address _to, uint _amount) public onlyOwner returns (bool success){
        tokenAddr.transfer(_to, _amount);
        return true;
    }

    function setFeePerc(uint8 _feePerc) public onlyOwner returns (bool success){
        require(_feePerc < 100, "fee percentage should be less than 100");
        feePerc = _feePerc;
        return true;
    }

    function setInternalFeePerc(uint8 _internalFeePerc) public onlyOwner returns (bool success){
        require(_internalFeePerc < 100, "fee percentage should be less than 100");
        internalFeePerc = _internalFeePerc;
        return true;
    }

    function setRewardPerc(uint8 _rewardPerc) public onlyOwner returns (bool success){
        require(_rewardPerc < 100, "fee percentage should be less than 100");
        rewardPerc = _rewardPerc;
        return true;
    }

    function setAffiliatesPerc(uint8[] memory _affiliatesPerc) public onlyOwner returns (bool success) {
        uint8 totalPerc;
        for (uint8 i = 0; i < _affiliatesPerc.length; i++) {
            totalPerc = totalPerc + _affiliatesPerc[i];
        }
        require(totalPerc == 100, "total percentage should be 100");
        affiliatesPerc = _affiliatesPerc;
        return true;
    }

    function setPartners(address payable[] memory _partners, uint8[] memory _partnersPerc) public onlyOwner returns (bool success){
        require(_partners.length == _partnersPerc.length, "parameters not match");
        uint8 totalPerc;
        for (uint8 i = 0; i < _partnersPerc.length; i++) {
            totalPerc = totalPerc + _partnersPerc[i];
        }
        require(totalPerc == 100, "total percentage should be 100");
        partners = _partners;
        partnersPerc = _partnersPerc;
        return true;
    }

    function destroy(address payable _to) public onlyOwner {
        selfdestruct(_to);
    }
}