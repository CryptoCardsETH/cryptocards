pragma solidity ^0.4.17;
import "./ContractAccessControl.sol";

contract CardBase is ContractAccessControl {
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
	mapping (address => uint256[]) public cardsHeldByOwner;

	// Mapping of owner address to number of tokens held
	mapping (address => uint256) ownershipTokenCount;

	// Transfer Event: Emitted every time a card is transfered to a new address
	event Transfer(address from, address to, uint256 tokenID);

	// Process Transfer: Reassign ownership and emit transfer event
	function _transfer(address _from, address _to, uint256 _cardID) internal {
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

	function _createCard(uint128 _battleID, uint256 _attributes, address _owner) internal returns (uint256) {
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
