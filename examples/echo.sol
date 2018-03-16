pragma solidity ^0.4.21;

contract PingPong {

	event Pong(address sender, bytes4 funcHash, bytes callData, bytes32 blockhash);

	function Ping() public {
		emit Pong(msg.sender, msg.sig, msg.data, block.blockhash(block.number));
	}

	function () public {
		emit Pong(msg.sender, msg.sig, msg.data, block.blockhash(block.number));
	}
}
