pragma solidity 0.4.21;

import "blah/Math/SafeMath.sol";
import "blah/Modules/Administration.sol";
import "Interfaces/ERC20Interface.sol";
/*


	When a bike is returned on time, the deposit is returns to the renter, and the cost of the rent is sent to the owner
	Other than this, there is only one other method for funds to be withdrawn from the contract, which is when the bike is not returned.

	when returning a bike, instead of going straight to available, mark as returned (to do)
*/
contract BikeRental is Administration {

	using SafeMath for uint256;

	uint256 public bikeCount;
	bool	public dev = true;
	RentalStates private defaultState = RentalStates.available;
	ERC20Interface private ercI = ERC20Interface(address(0));

	enum RentalStates { available, rented, returned, awol }

	struct BikeStruct {
		uint256 id;
		uint256 costPerDay;
		RentalStates state;
		bool	exists;
	}

	struct RentalStruct {
		address renter;
		uint256 bikeId;
		uint256 deposit;
		uint256 cost;
		uint256 daysRented;
		uint256 returnDate;
		bool	rented;
	}

	mapping (uint256 => BikeStruct) public bikes;
	mapping (address => RentalStruct) public rentals;

	event BikeAdded(uint256 _id, uint256 _costPerDay);
	event BikeRented(address _renter, uint256 _id, uint256 _daysRented);
	event BikeReturned(address _renter, uint256 _bikeId, bool _late);
	event BikeAwol(address _renter, uint256 _bikeId);
	event ErcInterfaceSet();

	modifier bikeAvailable(uint256 _bikeId) {
		require(bikes[_bikeId].state == RentalStates(0) && bikes[_bikeId].exists == true);
		_;
	}

	modifier bikeNotAdded(uint256 _bikeId) {
		require(bikes[_bikeId].state == RentalStates(0) && bikes[_bikeId].exists == false);
		_;
	}

	// this prevents someone from running the return function if the bike went awol and they failed to show up
	modifier isRentingBikeId(address _renter, uint256 _bikeId) {
		require(rentals[_renter].rented == true && rentals[_renter].bikeId == _bikeId && bikes[_bikeId].state == RentalStates(1));
		_;
	}

	modifier notRenting(address _renter) {
		require(!rentals[_renter].rented);
		_;
	}

	function setErcInterface(
		address _bikeTokenContractAddress)
		public
		onlyOwner
		returns (bool)
	{
		ercI = ERC20Interface(_bikeTokenContractAddress);
		emit ErcInterfaceSet();
		return true;
	}

	function addBike(
		uint256 _bikeId,
		uint256 _costPerDay)
		public
		bikeNotAdded(_bikeId)
		returns (bool)
	{
		bikes[_bikeId] = BikeStruct(_bikeId, _costPerDay, defaultState, true);
		bikeCount = bikeCount.add(1);
		emit BikeAdded(_bikeId, _costPerDay);
		return true;
	}

	function rentBike(
		uint256 _bikeId,
		uint256 _daysToRent,
		uint256 _deposit)
		public
		bikeAvailable(_bikeId)
		notRenting(msg.sender)
		returns (bool)
	{
		// they must deposit 3 times the cost of the bike, if returned within time they 2/3 of their money back
		require(_deposit == (bikes[_bikeId].costPerDay.mul(_daysToRent)).mul(3));
		uint256 cost = bikes[_bikeId].costPerDay.mul(_daysToRent);
		uint256 returnDate = _daysToRent.mul(1 days);
		returnDate = now.add(returnDate);
		bikes[_bikeId].state = RentalStates(1);
		rentals[msg.sender] = RentalStruct(msg.sender, _bikeId, _deposit, cost, _daysToRent, returnDate, true);
		emit BikeRented(msg.sender, _bikeId, _daysToRent);
		require(ercI.transferFrom(msg.sender, address(this), _deposit));
		require(ercI.transferFrom(msg.sender, owner, cost));
		return true;
	}

	function returnBike(
		uint256 _bikeId)
		public
		isRentingBikeId(msg.sender, _bikeId)
		returns (bool)
	{
		if (now <= rentals[msg.sender].returnDate) {
			require(onTimeReturn(_bikeId));
			return true;
		} else {
			require(lateReturn(_bikeId));
			return true;
		}
	}

	function forceLateReturn(
		address _renter,
		uint256 _bikeId)
		public
		onlyOwner
		returns (bool)
	{
		require(dev);
		uint256 remainingFunds = rentals[_renter].deposit.sub(rentals[_renter].cost);
		/*
			Take advantage of storage refund mechanic, and refund them with a bit of gas
		*/
		delete rentals[_renter];
		bikes[_bikeId].state = defaultState;
		emit BikeReturned(_renter, _bikeId, true);
		require(ercI.transfer(owner, remainingFunds));
		return true;
	}

	function forceBikeAwol(
		address _renter,
		uint256 _bikeId)
		public
		onlyOwner
		isRentingBikeId(_renter, _bikeId)
		returns (bool)
	{
		require(dev);
		uint256 remainingFunds = rentals[_renter].deposit.sub(rentals[_renter].cost);
		/*
			storage refund mechanics
		*/
		delete rentals[_renter];
		bikes[_bikeId].state = RentalStates.awol;
		emit BikeAwol(_renter, _bikeId);
		require(ercI.transfer(msg.sender, remainingFunds));
		return true;
	}

	function bikeAwol(
		address _renter,
		uint256 _bikeId)
		public
		onlyOwner
		isRentingBikeId(_renter, _bikeId)
		returns (bool)
	{
		require(now >= rentals[_renter].returnDate.add(30 days));
		uint256 remainingFunds = rentals[_renter].deposit.sub(rentals[_renter].cost);
		/*
			storage refund mechanics
		*/
		delete rentals[_renter];
		bikes[_bikeId].state = RentalStates.awol;
		emit BikeAwol(_renter, _bikeId);
		require(ercI.transfer(msg.sender, remainingFunds));
		return true;
	}

	function onTimeReturn(
		uint256 _bikeId)
		internal
		returns (bool)
	{
		uint256 fundsReturned = rentals[msg.sender].deposit.sub(rentals[msg.sender].cost);
		/*
			Here we can take advantage of the storage refund mechanic, and refund the user with a little bit of gas
		*/
		delete rentals[msg.sender];
		bikes[_bikeId].state = defaultState;
		emit BikeReturned(msg.sender, _bikeId, false);
		require(ercI.transfer(msg.sender, fundsReturned));
		return true;
	}

	function lateReturn(
		uint256 _bikeId)
		internal
		returns (bool)
	{
		uint256 remainingFunds = rentals[msg.sender].deposit.sub(rentals[msg.sender].cost);
		/*
			Take advantage of storage refund mechanic, and refund them with a bit of gas
		*/
		delete rentals[msg.sender];
		bikes[_bikeId].state = defaultState;
		emit BikeReturned(msg.sender, _bikeId, true);
		require(ercI.transfer(owner, remainingFunds));
		return true;
	}

}

/*

	struct RentalStruct {
		address renter;
		uint256 bikeId;
		uint256 deposit;
		uint256 cost;
		uint256 daysRented;
		uint256 returnDate;
		bool	rented;
	}
*/