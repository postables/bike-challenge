pragma solidity 0.4.21;

interface BikeRentalInterface {
	function addBike(uint256 _bikeId, uint256 _costPerDay) external returns (bool);
	function rentBike(uint256 _bikeId, uint256 _daysToRent, uint256 _deposit) external returns (bool);
	function returnBike(uint256 _bikeId) external returns (bool);
	function bikeExists(uint256 _id) external view returns (bool);
	function checkIfRentingBikeId(address _renter, uint256 _bikeId) external view returns (bool);
	function checkIfBikeIsReturned(uint256 _id) external view returns (bool);
	function setErcInterface(address _token) external returns (bool);
}