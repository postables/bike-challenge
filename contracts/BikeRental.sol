pragma solidity 0.4.21;

import "blah/Math/SafeMath.sol";
import "blah/Modules/Administration.sol";
import "Interfaces/ERC20Interface.sol";

contract BikeRental is Administration {

	using SafeMath for uint256;

	uint256 public bikeCount;

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
		uint256 daysRented;
		uint256 returnDate;
		bool	rented;
	}

	mapping (uint256 => BikeStruct) public bikes;
	mapping (address => RentalStruct) private rentals;

	// this allows us to easily search for the event using the `id` topic
	event BikeAdded(uint256 _id, uint256 _costPerDay);
	event BikeRented(address _renter, uint256 _id, uint256 _daysRented);

	modifier bikeAvailable(uint256 _bikeId) {
		require(bikes[_bikeId].state == RentalStates(0) && bikes[_bikeId].exists == true);
		_;
	}

	modifier bikeNotAdded(uint256 _bikeId) {
		require(bikes[_bikeId].state == RentalStates(0) && bikes[_bikeId].exists == false);
		_;
	}

	modifier isRentingBikeId(address _renter, uint256 _bikeId) {
		require(rentals[_renter].rented == true && rentals[_renter].bikeId == _bikeId);
		_;
	}

	modifier notRenting(address _renter) {
		require(!rentals[_renter].rented);
		_;
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
		uint256 returnDate = _daysToRent.mul(1 days);
		returnDate = now.add(returnDate);
		bikes[_bikeId].state = RentalStates(1);
		rentals[msg.sender] = RentalStruct(msg.sender, _bikeId, _deposit, _daysToRent, returnDate, true);
		emit BikeRented(msg.sender, _bikeId, _daysToRent);
		require(ercI.transferFrom(msg.sender, address(this), _deposit));
		return true;
	}

	function returnBike(
		uint256 _bikeId)
		public
		view
		isRentingBikeId(msg.sender, _bikeId)
		returns (bool)
	{
		if (now <= rentals[msg.sender].returnDate) {
			// ontime return
			return true;
		} else {
			// late return
			return true;
		}
	}

}