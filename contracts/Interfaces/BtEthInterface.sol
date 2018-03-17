pragma solidity 0.4.21;

interface BtEthInterface {
	function launch(address _token) external returns (bool);
	function checkIfLaunched() external view returns (bool);
	function enabled() external view returns (bool);
	function enableSales() external returns (bool);
	function disableSales() external returns (bool);
}