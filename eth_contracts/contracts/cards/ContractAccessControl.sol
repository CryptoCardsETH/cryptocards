contract ContractAccessControl {

    ////////////////////////// Ownership //////////////////////////

    address public ownerAddress;        // Address of Contract Owner

    modifier onlyOwner() {
        require(msg.sender == ownerAddress);
        _;
    }

    function setOwner(address _newOwner) external onlyOwner {
        require(_newOwner != address(0));

        ownerAddress = _newOwner;
    }

    ////////////////////////// Pausing //////////////////////////

    bool public paused = false;         // Used to Pause Application

    modifier whenNotPaused() {
        require(!paused);
        _;
    }

    modifier whenPaused {
        require(paused);
        _;
    }

    function pause() external onlyOwner whenNotPaused {
        paused = true;
    }

    function unpause() public onlyOwner whenPaused {
        paused = false;
    }

    ////////////////////////// Contract Upgrades //////////////////////////

    address public newContractAddress;          // Address of new Contract clients should switch to

    event ContractUpgrade(address newContract); // Event emitted after Contract Upgrade

    function setNewAddress(address _newAddress) external onlyOwner whenPaused {
        newContractAddress = _newAddress;
        ContractUpgrade(_newAddress); // Emit event announcing update to new contract
    }
}