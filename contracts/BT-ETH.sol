pragma solidity 0.4.21;

import "blah/Modules/Administration.sol";
import "blah/Math/SafeMath.sol";
import "blah/Interfaces/ERC20Interface.sol";

/*
	Used to buy BT tokens for fixed rate eth
*/
contract BTETH is Administration {

	// 31415.9
	uint256 public constant btEth = 31415900000000000000000;

	uint256 public remainingTokens;
	uint256 public numBikeTokensSold;
	uint256 public numContributions;
	uint256 public ethRaised;

	bool public enabled;
	bool private initialized;

	// although this isn't "upgradable" it saves gas at deployment time by using a private variable, with further savings from `constant`
	ERC20Interface private constant ercI = ERC20Interface(address(0));


	mapping (address => uint256) public contributions;

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
		contribute();
	}

	function launch()
		public
		onlyOwner
		isUninitialized
		returns (bool)
	{
		// make sure we can sell at least one bike token
		require(ercI.balanceOf(address(this)) > btEth);
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

	function contribute() internal pure {}
}