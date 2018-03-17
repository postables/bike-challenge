pragma solidity 0.4.21;

interface AdministrationInterface {
	function changeOwner(address  _newOwner) external returns (bool);
}