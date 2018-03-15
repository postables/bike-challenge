pragma solidity 0.4.21;

import "blah/Modules/Administration.sol";
import "blah/Math/SafeMath.sol";
import "blah/Interfaces/ERC20Interface.sol";

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
	bool private initialized;

	// although this isn't "upgradable" it saves gas at deployment time by using a private variable, with further savings from `constant`
	ERC20Interface private constant ercI = ERC20Interface(address(0));

	mapping (address => uint256) public contributions;

	event Sale(uint256 reward);

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

	function launch()
		public
		onlyOwner
		isUninitialized
		returns (bool)
	{
		// make sure we can sell at least one bike token
		require(ercI.balanceOf(address(this)) > btWei);
		remainingTokens = ercI.balanceOf(address(this));
		initialized = true;
		enabled = true;
		// event placeholder
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
		// event placeholder
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
		// event placeholder
		return true;
	}

	// lets someone calculate what they reward would be
    function contributionCalculator(uint256 _eth) public view returns (uint256) {
    	return ((_eth.mul(10)).div(btWei)).mul(10**17);
    }

	function contribute()
		internal
		returns (bool)
	{
		require(msg.value > 111100);
		require(remainingTokens > 0);
		uint256 reward = ((msg.value.mul(10)).div(btWei)).mul(10**17);
		assert(reward > 0);
		ercI.transfer(msg.sender, reward);
		// event placeholder
		return true;
	}
}