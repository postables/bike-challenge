pragma solidity 0.4.21;

import "blah/Math/SafeMath.sol";
import "blah/Modules/Administration.sol";
import "Interfaces/ERC20Interface.sol";

contract BikeRental is Administration {

	using SafeMath for uint256;

	uint256 public bikeCount;

	RentalStates private defaultState = RentalStates.available;

	enum RentalStates { available, rented, returned, awol }

	struct BikeStruct {
		uint256 id;
		uint256 cost;
		RentalStates state;
	}

	struct RentalStruct {
		address renter;
		uint256 bikeId;
		uint256 daysRented;
		uint256 deposit;
	}

	mapping (uint256 => BikeStruct) public bikes;
	mapping (address => RentalStruct) private rentals;

	modifier bikeAvailable(uint256 _bikeId) {
		require(bikes[_bikeId].state == RentalStates(0));
		_;
	}

}