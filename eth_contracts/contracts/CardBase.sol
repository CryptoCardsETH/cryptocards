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

		// Card name
		bytes32 name;
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

	// Generate name of card
	function _name() internal returns (bytes32) {

		bytes32[] firstname = {
			Dave,
			Jason,
			Harris,
			Ben,
			Lena,
			Nicky,
			Shelton,
			Salty,
			Ricky,
			Nate,
			Arnold,
			Dick,
			Richard,
			Matthew,
			Noah,
			Connor,
			Wayne,
			Chronos
			Walton,
			Olympus,
			Jarman,
			Oliver,
			Tenenan,
			Quasim,
			Fear,
			Saxton,
			Arland,
			Maxwell,
			Kemp,
			Patamon,
			Sigmund,
			Rory,
			Mason,
			Sylvester,
			Ketan,
			Flunkey,
			Kaniel,
			Charm,
			Jeff,
			Nils,
			Jude,
			Jacob,
			Fane,
			Ferran,
			Benedict,
			Zephan,
			Harith,
			Adan,
			Lawrence,
			Merle,
			Jarvis,
			Willie,
			Hubert,
			Donny,
			Rehoboam,
			Tory,
			Brandon,
			Carmen,
			Courtney,
			Rey,
			Kevin,
			Ruben,
			Sung,
			Donnie,
			Ronnie,
			Micah,
			Hans,
			Calvin,
			Malcom,
			Bradley,
			Rolland,
			Hollis,
			Nathanial,
			Jerrell,
			Wallace,
			Columbus,
			Alfonzo,
			Isreal,
			Guadalupe,
			Kim,
			Tomas,
			Donn,
			Michale,
			Martin,
			Booker,
			Bryon,
			Sylvester,
			Cameron,
			Nathan,
			Sam,
			Isaac,
			Merrill,
			Ove,
			Vid,
			Pumba,
			Aitan,
			Ross,
			John,
			Thorfinn,
			Holt,
			Sol,
			Del,
			Coleman,
			Jonathon,
			Winston,
			Austin,
			Carlos
		};

		bytes32[] adjectives = {
			repulsive,
			resonant,
			ripe,
			roasted,
			robust,
			rotten,
			rough,
			round,
			sad,
			salty,
			scary,
			scattered,
			scrawny,
			screeching,
			selfish,
			shaggy,
			shaky,
			shallow,
			sharp,
			shivering,
			short,
			shrill,
			silent,
			silky,
			silly,
			skinny,
			slimy,
			slippery,
			slow,
			small,
			smart,
			smiling,
			smooth,
			soft,
			solid,
			sore,
			sour,
			spicy,
			splendid,
			spotty,
			square,
			squealing,
			stale,
			steady,
			steep,
			sticky,
			stingy,
			straight,
			strange,
			striped,
			strong,
			stupendous,
			defeated,
			defiant,
			delicious,
			delightful,
			depressed,
			determined,
			dirty,
			disgusted,
			disturbed,
			dizzy,
			dry,
			dull,
			dusty,
			eager,
			early,
			elated,
			embarrassed,
			empty,
			encouraging,
			energetic,
			enthusiastic,
			envious,
			evil,
			excited,
			exuberant,
			faint,
			fair,
			faithful,
			fantastic,
			fast,
			fat,
			few,
			fierce,
			filthy,
			fine,
			flaky,
			flat,
			fluffy,
			foolish,
			forlorn,
			frail,
			frantic,
			fresh,
			friendly,
			frightened,
			funny,
			fuzzy,
			gentle,
			giant,
			gigantic,
			good
		};

		bytes32[] lastname = {
			Architect,
			Cello,
			Comfort,
			Lady,
			Rowboat,
			Puddle,
			Shirt,
			Syrup,
			Bag,
			Spy,
			Lettuce,
			Kite,
			Football,
			Wonder,
			Stranger,
			Torso,
			Vineyard,
			Giant,
			Log,
			Bomb,
			Oval,
			Wall,
			Roadway,
			Sheriff,
			Deputy,
			Programmer,
			Student,
			Assassin,
			Vehicle,
			Artist,
			Father,
			Father-in-law,
			Female,
			Fighter,
			Goldfish,
			Gorilla,
			Toon,
			Gobbler,
			Tickler,
			Guitar,
			Hacksaw,
			Hammer,
			Hamburger,
			Guy,
			Item,
			Jackal,
			Judge,
			Jumpsuit,
			Ladybug,
			Knife,
			Leg,
			Lion,
			Lipstick,
			Lobster,
			Produce,
			Substance,
			Scale,
			Letter,
			Oil,
			Meeting,
			Size,
			Death,
			Servant,
			Driving,
			Hope,
			Room,
			Invention,
			Shake,
			Nation,
			Relation,
			Love,
			Tin,
			River,
			Summer,
			Weight,
			Increase,
			Crack,
			Flower,
			Base,
			Offer,
			Mountain,
			Committee,
			Ray,
			Blood,
			Rain,
			Name,
			Balance,
			Example,
			Development,
			Secretary,
			Self,
			Mass,
			Damage,
			Reward,
			Growth,
			Act,
			Effect,
			Rhythm,
			Work,
			Verse,
			Copper,
			Plant,
			Sign,
			Slip,
			Cork,
			View,
			Reason,
			Space,
			Side,
			Hate,
			Run,
			Disease,
			Edge,
			Iron,
			Measure,
			Middle,
			Grip,
			Law,
			Fact,
			End,
			Note,
			Regret,
			Roll,
			Attention,
			Design,
			Attack,
			Country,
			Talk,
			Kick,
			Insurance,
			Exchange,
			Steel,
			Road
		};

		// "weak seeding" pseudorandom from Apple's FreeBSD 'do_rand' func
		// cardID (seed) cannot be 0, so use cardID + 1
		uint32 hashval = 16807 * ((cardID+1) % 127773) - (2836 * ((cardID+1)) / 127773));
		bytes32 cardName = firtname[hashVal % firstname.length] + adjectives[hashVal % adjectives.length] + lastname[hashVal % lastname.length];

		return cardName;

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
			attributes: _attributes,
			name: _name
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

	function createCard(address _owner, uint256 _attributes) external returns (uint) {
		// Prevent ownership by CryptoCards contracts
		require(_owner != address(this));

		// Create the card
		uint cardID = _createCard(0, _attributes, _owner);

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
