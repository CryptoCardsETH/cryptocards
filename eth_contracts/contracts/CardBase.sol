pragma solidity ^0.4.17;

contract CardBase {
	struct GameCard {
		// The timestamp from the block when the card was created
		uint64 creationTime;

		// Time after which the card is no longer battle locked (0 if not currently in battle)
		uint64 battleCooldownEnd;

		// ID of battle which created card
		uint128 creationBattleID;

		// Count of battles card has participated in
		uint128 battleCount;

		// ID of battle which card is most recently in
		uint128 currentBattleID;

		// Card Attributes
		uint256 attributes;
	}

	// Array of all existing GameCards
	GameCard[] cards;

	// Mapping of cardIDs to owner address
	mapping (uint256 => address) public cardIndexToOwner;

	// Mapping of owner address to number of tokends held
	mapping (address => uint256) ownershipTokenCount;

	function balanceOf(address _owner) public view returns (uint256 count) {
		return ownershipTokenCount[_owner];
	}

	// Transfer Event: Emitted every time a card is transfered to a new address
	event Transfer(address from, address to, uint256 tokenID);

	// Process Transfer: Reassign ownership and emit transfer event
	function _transfer(address _from, address _to, uint256 _cardID) internal
	{
		// Increase new owner's token count
		ownershipTokenCount[_to]++;

		// Transfer ownership
		cardIndexToOwner[_cardID] = _to;

		// Handle old ownership (newly created cards are sent from address 0x00)
		if (_from != address(0)) {
			ownershipTokenCount[_from]--;
		}

		// Emit transfer event
		Transfer(_from, _to, _cardID);
	}

	// New Card Event: Emitted every time a new card is created
	event NewCard(address owner, uint256 cardID, uint128 creationBattleID, uint256 attributes);

	function _createCard(uint128 _battleID, uint256 _attributes, address _owner) internal returns (uint)
	{
		GameCard memory _card = GameCard({
			creationTime: uint64(now),
			battleCooldownEnd: 0,
			creationBattleID: _battleID,
			battleCount: 0,
			currentBattleID: 0,
			attributes: _attributes
		});
		uint256 newCardID = cards.push(_card) - 1;

		// Make sure we never overflow the card max (4 billion cards)
		require(newCardID == uint256(uint32(newCardID)));

		// Emit NewCard Event
		NewCard(
			_owner,
			newCardID,
			_card.creationBattleID,
			_card.attributes
		);

		// Assign ownership and emit the Transfer event - per ERC732 draft
		_transfer(0, _owner, newCardID);

		return newCardID;
	}
}

contract CardOwnership is CardBase {
	// Name and symbol of the ERC token
	string public constant name = "CryptoCards";
	string public constant symbol = "CCB";

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

	function createCard(address _owner) external returns (uint) {
		// Prevent ownership by CryptoCards contracts
		require(_owner != address(this));

		// Create the card
		uint cardID = _createCard(0, 0, _owner);

		return cardID;
	}

	function ownerOf(uint256 _tokenId) external view returns (address owner)
	{
		owner = cardIndexToOwner[_tokenId];
		require(owner != address(0));
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
