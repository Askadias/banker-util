pragma solidity >=0.5.0 <0.7.0;

contract Token {

    function transfer(address _to, uint256 _value) public returns (bool success) {}

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {}

    function allowance(address  _owner, address _spender) public view returns (uint256 remaining) {}
}
