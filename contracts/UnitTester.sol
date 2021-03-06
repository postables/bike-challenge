pragma solidity 0.4.21;

import "Interfaces/ERC20Interface.sol";
import "Interfaces/BtEthInterface.sol";
import "Interfaces/BikeRentalInterface.sol";
import "Math/SafeMath.sol";
import "Modules/Administration.sol";

contract UnitTester is Administration {

	using SafeMath for uint256;

	uint8 public contractCreationNonce = 1;
	address public tokenContract;
	address public saleContract;
	address public rentalContract;

	ERC20Interface private ercI;
	BikeRentalInterface private brI;
	BtEthInterface private btI;

	event TestPassed(string _testName);
	event ContractCreated(address _contractAddress, string _contractName);

	// ok
	function setErcI(address _token)  public onlyOwner returns (bool) {
		ercI = ERC20Interface(_token);
		tokenContract = _token;
		return true;
	}

	// ok
	function setBrI(address _brAddress) public onlyOwner returns (bool) {
		brI = BikeRentalInterface(_brAddress);
		rentalContract = _brAddress;
		return true;
	}

	// ok
	function setBtI(address _btAddress) public onlyOwner returns (bool) {
		btI = BtEthInterface(_btAddress);
		saleContract = _btAddress;
		return true;
	}

	// ok
	function testTransferTokens(address _recipient, uint256 _amount) public returns (bool) {
		require(ercI.transfer(_recipient, _amount));
		emit TestPassed("token transfer test");
		assert(ercI.balanceOf(_recipient) == _amount);
	}

	// ok  [ requires sending to a  tokenless address ]
	function testTransferFrom(address _owner, address _recipient, uint256 _amount) public returns (bool) {
		require(ercI.transferFrom(_owner, _recipient, _amount));
		emit TestPassed("token transfer from test");
		assert(ercI.balanceOf(_recipient) == _amount);
	}

	// ok [ only usable once per destination address]
	function testTokenApproval(address _spender, uint256 _amount) public returns (bool) {
		require(ercI.approve(_spender, _amount));
		emit TestPassed("token approval test");
		assert(ercI.allowance(address(this), _spender) == _amount);
	}

	// ok [ requires that you have sent tokens to the sale contract first ]
	function testLaunchBtEth() public returns (bool) {
		require(btI.launch(tokenContract));
		emit TestPassed("BT-ETH initialization test");
		assert(btI.checkIfLaunched() == true);
	}

	// ok
	function testDisableSales() public returns (bool) {
		require(btI.disableSales());
		emit TestPassed("BT-ETH sales disable test");
		assert(btI.enabled() == false);
	}

	// ok
	function testEnableSales() public returns (bool) {
		require(btI.enableSales());
		emit TestPassed("BTH-ETH sales enable test");
		assert(btI.enabled() == true);
	}

	// ok
	function testAddBike(uint256 _id, uint256 _costPerDay) public returns (bool) {
		require(brI.addBike(_id, _costPerDay));
		emit TestPassed("bike add test");
		assert(brI.bikeExists(_id));
	}

	// ok
	function testSetRentalTokenInterface() public returns (bool) {
		require(brI.setErcInterface(tokenContract));
		emit TestPassed("erc interface set for bike rental");
		assert(true);
	}

	// ok
	function testRentBike(uint256 _id, uint256 _daysToRent, uint256 _deposit) public returns (bool) {
		require(brI.rentBike(_id, _daysToRent, _deposit));
		emit TestPassed("bike rent test");
		assert(brI.checkIfRentingBikeId(address(this), _id));
	}

	// ok
	function testReturnBike(uint256 _bikeId) public returns (bool) {
		require(brI.returnBike(_bikeId));
		emit TestPassed("bike return test");
		assert(brI.checkIfBikeIsReturned(_bikeId));
	}
}