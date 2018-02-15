pragma solidity ^0.4.17;

contract CardBase {
	struct GameCard {
		// The timestamp from the block when the card was created
		uint64 creationTime;

		// Time after which the card is no longer battle locked (0 if not currently in battle)
		uint64 battleCooldownEnd;

		// ID of battle which created card
		uint32 creationBattleID;

		// ID of battle which card is most recently in
		uint32 currentBattleID;

		// Card Attributes
		uint256 attributes;
	}

	// Array of all existing GameCards
	GameCard[] cards;

	// Mapping of cardIDs to owner address
	mapping (uint256 => address) public cardIndexToOwner;

	// Mapping of owner address to number of cards held
	mapping (address => uint256) ownershipCardCount;

	// Transfer Event: Emitted every time a card is transfered to a new address
	event Transfer(address from, address to, uint256 tokenId);

	// Process Transfer: Reassign ownership and emit transfer event
	function _transfer(address _from, address _to, uint256 _cardId) internal {
		// Increase new owner's token count
		ownershipCardCount[_to]++;

		// Transfer ownership
		cardIndexToOwner[_cardId] = _to;

		// Handle old ownership (newly created cards are sent from address 0x00)
		if (_from != address(0)) {
			ownershipCardCount[_from]--;
		}

		// Emit transfer event
		Transfer(_from, _to, _tokenId);
	}

	// New Card Event: Emitted every time a new card is created
	event NewCard(address owner, uint256 cardID, uint256 battleID, uint256 attributes);
}

contract CardOwnership is CardBase {
	// Name and symbol of the ERC token
	string public constant name = "CryptoCards";
	string public constant symbol = "CCB";

	// Return if address _claimant currently holds card _cardId
	function _owns(address _claimant, uint256 _cardId) internal view returns (bool) {
		return cardIndexToOwner[_cardId] == _claimant;
	}

	// Return current token balance of address _owner
	function balanceOf(address _owner) public view returns (uint256 count) {
		return ownershipTokenCount[_owner];
	}

	function transfer(address _to, uint256 _cardId) external {
		// Prevent transfer to 0x00 address
		require(_to != address(0));

		// Prevent transfers to CryptoCards contracts
		require(_to != address(this));

		// Prevent transfers if card _cardId is not owned by the sender
		require(_owns(msg.sender, _cardId));

		// Reassign ownership and emit transfer event
		_transfer(msg.sender, _to, _cardId);
	}
}