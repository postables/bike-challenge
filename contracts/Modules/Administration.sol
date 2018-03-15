pragma solidity 0.4.21;

contract Administration {

    address public owner;
    address public admin;

    event AdminSet(address _admin);
    event NewOwner(address _owner);

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier permissioned() {
        require(msg.sender == owner || msg.sender == admin);
        _;
    }
    
    /* solhint-disable func-visibility */
    function Administration() {
        owner = msg.sender;
        admin = msg.sender;
    }

    function setAdmin(
        address _admin)
        public
        onlyOwner
        returns (bool)
    {
        admin = _admin;
        emit AdminSet(_admin);
        return true;
    }

    function changeOwner(
        address _owner)
        public
        onlyOwner
        returns (bool)
    {
        owner = _owner;
        emit NewOwner(_owner);
        return true;
    }
}