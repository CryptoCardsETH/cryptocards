pragma solidity ^0.4.17;
import "../CryptoCardsCore.sol";

contract Marketplace {

	struct Listing {
		uint64 creationTime;	// The timestamp from the block when the battlegroup was created
		uint256 card;			// CardID for listing
		uint256 price;			// Price to purchase listing
	}

	// Array of all existing Listings
	Listing[] listings;

	// New Listing Event: Emitted every time a new Listing is created
	event NewListing(uint256 cardID, uint256 price);

	// Purashed Listing Event: Emitted every time a new Listing/Card is purshased
	event PurchasedListing(address oldOwner, address newOwner, uint256 cardID, uint256 price);

	CryptoCardsCore public cryptoCardsContract; // Pointer to Card Contract
	function Marketplace() public {
		CryptoCardsCore candidateContract = CryptoCardsCore(msg.sender);
		//require(candidateContract.supportsInterface(InterfaceSignature_ERC721));
		cryptoCardsContract = candidateContract;
	}

	// Return total count of all listings
	function countListings() public view returns (uint) {
		return listings.length;
	}

	function priceOf(uint256 _listingID) external view returns (uint256 listing) {
		require (_listingID < countListings());
		price = listings[_listingID].price;
	}

	function createListing(uint256 _card) external returns (uint256) {
		uint256 listingPrice = 10000;

		BattleGroup memory _listing = BattleGroup({
			creationTime: uint64(now),
			card: _card,
			price: listingPrice
		});
		uint256 newID = listings.push(_listing) - 1;

		// Make sure we never overflow the listing max (4 billion listings)
		require(newID == uint256(uint32(newID)));

		// Emit NewBattleGroup Event
		NewListing(
			newID,
			listingPrice
		);

		return newID;
	}

	function purchaseListing(uint256 _listingID) external returns (bool) {
		require (_listingID < countListings());
		Listing memory _listing = listings[_listingID];

		// Emit NewBattleGroup Event
		PurchasedListing(
			0,
			0,
			_listing.card,
			_listing.price
		);

		return true;
	}
}