pragma solidity 0.4.21;

import "Modules/Administration.sol";
import "Math/SafeMath.sol";
import "Interfaces/ERC20Interface.sol";

/*
	Used to buy BT tokens for fixed rate eth
*/
contract BTETH is Administration {
	using SafeMath for uint256;
	// 31415.9
	uint256 public constant btWei = 31831000000000;

	uint256 public remainingTokens;
	uint256 public numBikeTokensSold;
	uint256 public numContributions;
	uint256 public ethRaised;

	bool public enabled;
	bool public initialized;

	ERC20Interface private ercI;

	mapping (address => uint256) public contributions;

	event TokensSold(address indexed _purchaser, uint256 _contribution);
	event ContractLaunchedAndEnabled();
	event SalesDisabled();
	event SalesEnabled();

	modifier salesDisabled() {
		require(!enabled);
		_;
	}

	modifier salesEnabled() {
		require(enabled);
		_;
	}

	modifier isUninitialized() {
		require(!initialized);
		_;
	}

	modifier isInitialized() {
		require(initialized);
		_;
	}

	function () payable isInitialized salesEnabled {
		require(contribute());
	}

	function launch(
		address _tokenAddress)
		public
		onlyOwner
		isUninitialized
		returns (bool)
	{
		ercI = ERC20Interface(_tokenAddress);
		// make sure we can sell at least one bike token
		require(ercI.balanceOf(address(this)) > btWei);
		remainingTokens = ercI.balanceOf(address(this));
		initialized = true;
		enabled = true;
		emit ContractLaunchedAndEnabled();
		return true;
	}

	function enableSales()
		public
		onlyOwner
		salesDisabled
		isInitialized
		returns (bool)
	{
		enabled = true;
		emit SalesEnabled();
		return true;
	}

	function disableSales()
		public
		onlyOwner
		salesEnabled
		isInitialized
		returns (bool)
	{
		enabled = false;
		emit SalesDisabled();
		return true;
	}

	function contribute()
		internal
		returns (bool)
	{
		require(msg.value > 0);
		require(remainingTokens > 0);
		uint256 reward = ((msg.value.mul(10)).div(btWei)).mul(10**17);
		assert(reward > 0);
		contributions[msg.sender] = contributions[msg.sender].add(reward);
		remainingTokens = remainingTokens.sub(reward);
		numBikeTokensSold = numBikeTokensSold.add(reward);
		numContributions = numContributions.add(1);
		// if there are no more tokens left, disable sales
		if (remainingTokens == 0) {
			enabled = false;
			emit SalesDisabled();
		}
		emit TokensSold(msg.sender, reward);
		require(ercI.transfer(msg.sender, reward));
		return true;
	}

	// lets someone calculate what they reward would be
    function contributionCalculator(uint256 _eth) public pure returns (uint256) {
    	return ((_eth.mul(10)).div(btWei)).mul(10**17);
    }

    function checkIfLaunched() public view returns (bool) {
    	if (initialized == true && enabled == true) {
    		return true;
    	} else {
    		return false;
    	}
    }

}