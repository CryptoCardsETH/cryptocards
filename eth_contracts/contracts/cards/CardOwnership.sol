pragma solidity ^0.4.17;
import "./CardBase.sol";

contract CardOwnership is CardBase {
	// Name and symbol of the ERC token
	string public constant name = "CryptoCards";
	string public constant symbol = "CCB";

	// Require that a game card exists
	function requireCardExists(uint256 cardID) public view {
		require (cardID < totalSupply());
	}

	// Return if address _claimant currently holds card _cardId
	function _owns(address _claimant, uint256 _cardID) internal view returns (bool) {
		return cardIndexToOwner[_cardID] == _claimant;
	}

	// Return total count of all existing tokens (cards)
	function totalSupply() public view returns (uint) {
		return cards.length;
	}

	// Return current token balance of address _owner
	function balanceOf(address _owner) public view returns (uint256 count) {
		return ownershipTokenCount[_owner];
	}

	function transfer(address _to, uint256 _cardId) external whenNotPaused {
		// Prevent transfer to 0x00 address
		require(_to != address(0));

		// Prevent transfers to CryptoCards contracts
		require(_to != address(this));

		// Prevent transfers if card _cardId is not owned by the sender
		require(_owns(msg.sender, _cardId));

		// Reassign ownership and emit transfer event
		_transfer(msg.sender, _to, _cardId);
	}

	function createCard(address _owner, uint256 _attributes) public whenNotPaused returns (uint256) {
		// Prevent ownership by CryptoCards contracts
		require(_owner != address(this));

		// Create the card
		uint256 cardID = _createCard(0, _attributes, _owner);

		return cardID;
	}

	function ownerOf(uint256 _tokenId) external view returns (address owner) {
		owner = cardIndexToOwner[_tokenId];
		require (owner != address(0));
	}

	// NOTE: Do not call from within smartcontracts - too expensive to loop through all cards
	function tokensOfOwner(address _owner) external view returns(uint256[] ownerTokens) {
		uint256 tokenCount = balanceOf(_owner);

		if (tokenCount == 0) { // No tokens, return an empty array
			return new uint256[](0);
		} else {
			uint256[] memory result = new uint256[](tokenCount);
			uint256 totalTokens = totalSupply();
			uint256 resultIndex = 0;

			uint256 cardID;
			for (cardID = 1; cardID <= totalTokens; cardID++) {
				if (cardIndexToOwner[cardID] == _owner) {
					result[resultIndex] = cardID;
					resultIndex++;
				}
			}

			return result;
		}
	}
}
