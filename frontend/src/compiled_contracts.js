//DO NOT EDIT; abigen.js generates this
 const contracts = {
  "BattleGroups.sol:BattleGroups": {
    "bytecode": "6060604052341561000f57600080fd5b6103f68061001e6000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637bf13d82811461005b578063bfeea70814610080578063fa74efc114610093575b600080fd5b341561006657600080fd5b61006e6100c1565b60405190815260200160405180910390f35b341561008b57600080fd5b61006e6100c8565b341561009e57600080fd5b61006e73ffffffffffffffffffffffffffffffffffffffff6004351660246100cd565b6000545b90565b600581565b60006100d7610285565b60006060604051908101604052804267ffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001856005806020026040519081016040529190828260a0808284375050509190925250506000805491935060019180830161014b83826102ab565b600092835260209092208591600602018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815473ffffffffffffffffffffffffffffffffffffffff9190911668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff90911617815560408201516101e190600183019060056102dc565b50505003905063ffffffff811681146101f957600080fd5b7ff7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d8582846040015160405173ffffffffffffffffffffffffffffffffffffffff8416815260208101839052604081018260a080838360005b83811015610269578082015183820152602001610251565b50505050905001935050505060405180910390a1949350505050565b60e060405190810160409081526000808352602083015281016102a661031a565b905290565b8154818355818115116102d7576006028160060283600052602060002091820191016102d79190610341565b505050565b826005810192821561030a579160200282015b8281111561030a5782518255916020019190600101906102ef565b5061031692915061038d565b5090565b60a06040519081016040526005815b60008152602001906001900390816103295790505090565b6100c591905b808211156103165780547fffffffff00000000000000000000000000000000000000000000000000000000168155600061038460018301826103a7565b50600601610347565b6100c591905b808211156103165760008155600101610393565b5060008155600101600081556001016000815560010160008155600101600090555600a165627a7a723058204d4ef0fe8f00a82bc3fec8789e09c3b1219aaa675f1d921d7395c887b6b8019e0029",
    "interface": [
      {
        "constant": true,
        "inputs": [],
        "name": "countBattleGroups",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "MAX_CARDS_PER_GROUP",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "_owner",
            "type": "address"
          },
          {
            "name": "_cards",
            "type": "uint256[5]"
          }
        ],
        "name": "createBattleGroup",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "battleGroupID",
            "type": "uint256"
          },
          {
            "indexed": false,
            "name": "cards",
            "type": "uint256[5]"
          }
        ],
        "name": "NewBattleGroup",
        "type": "event"
      }
    ]
  },
  "Battles.sol:Battles": {
    "bytecode": "6060604052341561000f57600080fd5b6102838061001e6000396000f30060606040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166363c785aa8114610050578063df73c49f14610075575b600080fd5b341561005b57600080fd5b61006361008e565b60405190815260200160405180910390f35b341561008057600080fd5b610063600435602435610095565b6000545b90565b600061009f6101b9565b60006080604051908101604052804267ffffffffffffffff168152602001868152602001858152602001600063ffffffff1681525091506001600080548060010182816100ec91906101e0565b600092835260209092208591600402018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518160010155604082015181600201556060820151600391909101805463ffffffff191663ffffffff928316179055929091039250508116811461016257600080fd5b7fea1f5e30f8edb20fc59251726f685effaf003f76c4a416c179f62bda57477116818360200151846040015160405180848152602001838152602001828152602001935050505060405180910390a1949350505050565b60806040519081016040908152600080835260208301819052908201819052606082015290565b81548183558181151161020c5760040281600402836000526020600020918201910161020c9190610211565b505050565b61009291905b8082111561025357805467ffffffffffffffff19168155600060018201819055600282015560038101805463ffffffff19169055600401610217565b50905600a165627a7a7230582082fc20ed8f28b4c68a74dbbd343647c5c61c880e671ef9ea10f2f52137a5ad6a0029",
    "interface": [
      {
        "constant": true,
        "inputs": [],
        "name": "countBattles",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "_op1",
            "type": "uint256"
          },
          {
            "name": "_op2",
            "type": "uint256"
          }
        ],
        "name": "createBattle",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "battleID",
            "type": "uint256"
          },
          {
            "indexed": false,
            "name": "op1BattleGroup",
            "type": "uint256"
          },
          {
            "indexed": false,
            "name": "op2BattleGroup",
            "type": "uint256"
          }
        ],
        "name": "NewBattle",
        "type": "event"
      }
    ]
  },
  "CardBase.sol:CardBase": {
    "bytecode": "6060604052341561000f57600080fd5b6101528061001e6000396000f30060606040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312c2debf811461005057806378d0df5014610091575b600080fd5b341561005b57600080fd5b61007f73ffffffffffffffffffffffffffffffffffffffff600435166024356100d0565b60405190815260200160405180910390f35b341561009c57600080fd5b6100a76004356100fe565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6002602052816000526040600020818154811015156100eb57fe5b6000918252602090912001549150829050565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820623ac2f79f6010e6782c99c443f1048b8a91c81bc6ac33ccee75d657a9dafa050029",
    "interface": [
      {
        "constant": true,
        "inputs": [
          {
            "name": "",
            "type": "address"
          },
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "name": "cardsHeldByOwner",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "name": "cardIndexToOwner",
        "outputs": [
          {
            "name": "",
            "type": "address"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "from",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "to",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "tokenID",
            "type": "uint256"
          }
        ],
        "name": "Transfer",
        "type": "event"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "cardID",
            "type": "uint256"
          },
          {
            "indexed": false,
            "name": "creationBattleID",
            "type": "uint128"
          },
          {
            "indexed": false,
            "name": "attributes",
            "type": "uint256"
          }
        ],
        "name": "NewCard",
        "type": "event"
      }
    ]
  },
  "CardBase.sol:CardOwnership": {
    "bytecode": "6060604052341561000f57600080fd5b61092d8061001e6000396000f3006060604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100a857806312c2debf1461013257806318160ddd146101665780635de038b1146101795780636352211e1461019b57806370a08231146101cd57806378d0df50146101ec5780638462151c1461020257806395d89b4114610274578063a9059cbb14610287575b600080fd5b34156100b357600080fd5b6100bb6102ab565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100f75780820151838201526020016100df565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561013d57600080fd5b610154600160a060020a03600435166024356102e2565b60405190815260200160405180910390f35b341561017157600080fd5b610154610310565b341561018457600080fd5b610154600160a060020a0360043516602435610317565b34156101a657600080fd5b6101b160043561034f565b604051600160a060020a03909116815260200160405180910390f35b34156101d857600080fd5b610154600160a060020a0360043516610378565b34156101f757600080fd5b6101b1600435610393565b341561020d57600080fd5b610221600160a060020a03600435166103ae565b60405160208082528190810183818151815260200191508051906020019060200280838360005b83811015610260578082015183820152602001610248565b505050509050019250505060405180910390f35b341561027f57600080fd5b6100bb61048f565b341561029257600080fd5b6102a9600160a060020a03600435166024356104c6565b005b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b6002602052816000526040600020818154811015156102fd57fe5b6000918252602090912001549150829050565b6000545b90565b60008030600160a060020a031684600160a060020a03161415151561033b57600080fd5b61034760008486610520565b949350505050565b600081815260016020526040902054600160a060020a031680151561037357600080fd5b919050565b600160a060020a031660009081526003602052604090205490565b600160205260009081526040902054600160a060020a031681565b6103b661085e565b60006103c061085e565b60008060006103ce87610378565b94508415156103fe5760006040518059106103e65750595b90808252806020026020018201604052509550610485565b8460405180591061040c5750595b90808252806020026020018201604052509350610427610310565b925060009150600190505b82811161048157600081815260016020526040902054600160a060020a0388811691161415610479578084838151811061046857fe5b602090810290910101526001909101905b600101610432565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b600160a060020a03821615156104db57600080fd5b30600160a060020a031682600160a060020a0316141515156104fc57600080fd5b6105063382610775565b151561051157600080fd5b61051c338383610795565b5050565b600061052a610870565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600080548060010182816105bf91906108a5565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff811681146106ec57600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a161076c60008583610795565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905583161561080957600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b8154818355818115116108d1576003028160030283600052602060002091820191016108d191906108d6565b505050565b61031491905b808211156108fd5760008082556001820181905560028201556003016108dc565b50905600a165627a7a7230582075a449288ab459fa5787da589725e6e8bcd2cd84e5aa23c994fcbd32b76051ed0029",
    "interface": [
      {
        "constant": true,
        "inputs": [],
        "name": "name",
        "outputs": [
          {
            "name": "",
            "type": "string"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "",
            "type": "address"
          },
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "name": "cardsHeldByOwner",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "totalSupply",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "_owner",
            "type": "address"
          },
          {
            "name": "_attributes",
            "type": "uint256"
          }
        ],
        "name": "createCard",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "_tokenId",
            "type": "uint256"
          }
        ],
        "name": "ownerOf",
        "outputs": [
          {
            "name": "owner",
            "type": "address"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "_owner",
            "type": "address"
          }
        ],
        "name": "balanceOf",
        "outputs": [
          {
            "name": "count",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "name": "cardIndexToOwner",
        "outputs": [
          {
            "name": "",
            "type": "address"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "_owner",
            "type": "address"
          }
        ],
        "name": "tokensOfOwner",
        "outputs": [
          {
            "name": "ownerTokens",
            "type": "uint256[]"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "symbol",
        "outputs": [
          {
            "name": "",
            "type": "string"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "_to",
            "type": "address"
          },
          {
            "name": "_cardId",
            "type": "uint256"
          }
        ],
        "name": "transfer",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "from",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "to",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "tokenID",
            "type": "uint256"
          }
        ],
        "name": "Transfer",
        "type": "event"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "cardID",
            "type": "uint256"
          },
          {
            "indexed": false,
            "name": "creationBattleID",
            "type": "uint128"
          },
          {
            "indexed": false,
            "name": "attributes",
            "type": "uint256"
          }
        ],
        "name": "NewCard",
        "type": "event"
      }
    ]
  },
  "GenerateCard.sol:GenerateCard": {
    "bytecode": "606060405261000c610052565b604051809103906000f080151561002257600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055341561004d57600080fd5b610062565b60405161094b806102ba83390190565b610249806100716000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631b6f68b0811461005b578063abbaf85014610097578063d8c27fae146100c5575b600080fd5b341561006657600080fd5b61006e6100ed565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100a257600080fd5b6100c373ffffffffffffffffffffffffffffffffffffffff60043516610109565b005b34156100d057600080fd5b6100db600435610145565b60405190815260200160405180910390f35b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006001815b838211610211575060005460019073ffffffffffffffffffffffffffffffffffffffff16635de038b130846040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815273ffffffffffffffffffffffffffffffffffffffff90921660048301526024820152604401602060405180830381600087803b15156101df57600080fd5b5af115156101ec57600080fd5b505050604051805150508015156102065760009250610216565b60019091019061014b565b600192505b50509190505600a165627a7a72305820ab1b657cf3c024d1ab755247c02e46579c14ceeefe849be9f8f9fd13b3a0090100296060604052341561000f57600080fd5b61092d8061001e6000396000f3006060604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100a857806312c2debf1461013257806318160ddd146101665780635de038b1146101795780636352211e1461019b57806370a08231146101cd57806378d0df50146101ec5780638462151c1461020257806395d89b4114610274578063a9059cbb14610287575b600080fd5b34156100b357600080fd5b6100bb6102ab565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100f75780820151838201526020016100df565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561013d57600080fd5b610154600160a060020a03600435166024356102e2565b60405190815260200160405180910390f35b341561017157600080fd5b610154610310565b341561018457600080fd5b610154600160a060020a0360043516602435610317565b34156101a657600080fd5b6101b160043561034f565b604051600160a060020a03909116815260200160405180910390f35b34156101d857600080fd5b610154600160a060020a0360043516610378565b34156101f757600080fd5b6101b1600435610393565b341561020d57600080fd5b610221600160a060020a03600435166103ae565b60405160208082528190810183818151815260200191508051906020019060200280838360005b83811015610260578082015183820152602001610248565b505050509050019250505060405180910390f35b341561027f57600080fd5b6100bb61048f565b341561029257600080fd5b6102a9600160a060020a03600435166024356104c6565b005b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b6002602052816000526040600020818154811015156102fd57fe5b6000918252602090912001549150829050565b6000545b90565b60008030600160a060020a031684600160a060020a03161415151561033b57600080fd5b61034760008486610520565b949350505050565b600081815260016020526040902054600160a060020a031680151561037357600080fd5b919050565b600160a060020a031660009081526003602052604090205490565b600160205260009081526040902054600160a060020a031681565b6103b661085e565b60006103c061085e565b60008060006103ce87610378565b94508415156103fe5760006040518059106103e65750595b90808252806020026020018201604052509550610485565b8460405180591061040c5750595b90808252806020026020018201604052509350610427610310565b925060009150600190505b82811161048157600081815260016020526040902054600160a060020a0388811691161415610479578084838151811061046857fe5b602090810290910101526001909101905b600101610432565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b600160a060020a03821615156104db57600080fd5b30600160a060020a031682600160a060020a0316141515156104fc57600080fd5b6105063382610775565b151561051157600080fd5b61051c338383610795565b5050565b600061052a610870565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600080548060010182816105bf91906108a5565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff811681146106ec57600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a161076c60008583610795565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905583161561080957600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b8154818355818115116108d1576003028160030283600052602060002091820191016108d191906108d6565b505050565b61031491905b808211156108fd5760008082556001820181905560028201556003016108dc565b50905600a165627a7a723058200a84a1287fd762a72fd81884e2c3eb43fc1e91e8a4740617b265ca608ccee4f80029",
    "interface": [
      {
        "constant": true,
        "inputs": [],
        "name": "cardOwnership",
        "outputs": [
          {
            "name": "",
            "type": "address"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "_cardOwnership",
            "type": "address"
          }
        ],
        "name": "setCardOwnership",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "numberOfCards",
            "type": "uint256"
          }
        ],
        "name": "initialGeneration",
        "outputs": [
          {
            "name": "isSuccess",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      }
    ]
  },
  "Migrations.sol:Migrations": {
    "bytecode": "6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556102318061003b6000396000f3006060604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630900f0108114610066578063445df0ac146100945780638da5cb5b146100b9578063fdacd576146100f5575b600080fd5b341561007157600080fd5b61009273ffffffffffffffffffffffffffffffffffffffff6004351661010b565b005b341561009f57600080fd5b6100a76101b6565b60405190815260200160405180910390f35b34156100c457600080fd5b6100cc6101bc565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561010057600080fd5b6100926004356101d8565b600080543373ffffffffffffffffffffffffffffffffffffffff908116911614156101b2578190508073ffffffffffffffffffffffffffffffffffffffff1663fdacd5766001546040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401600060405180830381600087803b15156101a157600080fd5b5af115156101ae57600080fd5b5050505b5050565b60015481565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff908116911614156102025760018190555b505600a165627a7a723058208cb9a73d77bfa557836f71b587a2b4f7f40cdc7e2edf943852ce554af598b6530029",
    "interface": [
      {
        "constant": false,
        "inputs": [
          {
            "name": "new_address",
            "type": "address"
          }
        ],
        "name": "upgrade",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "last_completed_migration",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "owner",
        "outputs": [
          {
            "name": "",
            "type": "address"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "completed",
            "type": "uint256"
          }
        ],
        "name": "setCompleted",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      },
      {
        "inputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "constructor"
      }
    ]
  },
  "contracts/CardBase.sol:CardBase": {
    "bytecode": "6060604052341561000f57600080fd5b6101528061001e6000396000f30060606040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312c2debf811461005057806378d0df5014610091575b600080fd5b341561005b57600080fd5b61007f73ffffffffffffffffffffffffffffffffffffffff600435166024356100d0565b60405190815260200160405180910390f35b341561009c57600080fd5b6100a76004356100fe565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6002602052816000526040600020818154811015156100eb57fe5b6000918252602090912001549150829050565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820e032a4fe550083dc27f1351cb1cabeeeed8e2b645943d647d9ad37475d6cda0d0029",
    "interface": [
      {
        "constant": true,
        "inputs": [
          {
            "name": "",
            "type": "address"
          },
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "name": "cardsHeldByOwner",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "name": "cardIndexToOwner",
        "outputs": [
          {
            "name": "",
            "type": "address"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "from",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "to",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "tokenID",
            "type": "uint256"
          }
        ],
        "name": "Transfer",
        "type": "event"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "cardID",
            "type": "uint256"
          },
          {
            "indexed": false,
            "name": "creationBattleID",
            "type": "uint128"
          },
          {
            "indexed": false,
            "name": "attributes",
            "type": "uint256"
          }
        ],
        "name": "NewCard",
        "type": "event"
      }
    ]
  },
  "contracts/CardBase.sol:CardOwnership": {
    "bytecode": "6060604052341561000f57600080fd5b61092d8061001e6000396000f3006060604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100a857806312c2debf1461013257806318160ddd146101665780635de038b1146101795780636352211e1461019b57806370a08231146101cd57806378d0df50146101ec5780638462151c1461020257806395d89b4114610274578063a9059cbb14610287575b600080fd5b34156100b357600080fd5b6100bb6102ab565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100f75780820151838201526020016100df565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561013d57600080fd5b610154600160a060020a03600435166024356102e2565b60405190815260200160405180910390f35b341561017157600080fd5b610154610310565b341561018457600080fd5b610154600160a060020a0360043516602435610317565b34156101a657600080fd5b6101b160043561034f565b604051600160a060020a03909116815260200160405180910390f35b34156101d857600080fd5b610154600160a060020a0360043516610378565b34156101f757600080fd5b6101b1600435610393565b341561020d57600080fd5b610221600160a060020a03600435166103ae565b60405160208082528190810183818151815260200191508051906020019060200280838360005b83811015610260578082015183820152602001610248565b505050509050019250505060405180910390f35b341561027f57600080fd5b6100bb61048f565b341561029257600080fd5b6102a9600160a060020a03600435166024356104c6565b005b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b6002602052816000526040600020818154811015156102fd57fe5b6000918252602090912001549150829050565b6000545b90565b60008030600160a060020a031684600160a060020a03161415151561033b57600080fd5b61034760008486610520565b949350505050565b600081815260016020526040902054600160a060020a031680151561037357600080fd5b919050565b600160a060020a031660009081526003602052604090205490565b600160205260009081526040902054600160a060020a031681565b6103b661085e565b60006103c061085e565b60008060006103ce87610378565b94508415156103fe5760006040518059106103e65750595b90808252806020026020018201604052509550610485565b8460405180591061040c5750595b90808252806020026020018201604052509350610427610310565b925060009150600190505b82811161048157600081815260016020526040902054600160a060020a0388811691161415610479578084838151811061046857fe5b602090810290910101526001909101905b600101610432565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b600160a060020a03821615156104db57600080fd5b30600160a060020a031682600160a060020a0316141515156104fc57600080fd5b6105063382610775565b151561051157600080fd5b61051c338383610795565b5050565b600061052a610870565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600080548060010182816105bf91906108a5565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff811681146106ec57600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a161076c60008583610795565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905583161561080957600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b8154818355818115116108d1576003028160030283600052602060002091820191016108d191906108d6565b505050565b61031491905b808211156108fd5760008082556001820181905560028201556003016108dc565b50905600a165627a7a723058200a84a1287fd762a72fd81884e2c3eb43fc1e91e8a4740617b265ca608ccee4f80029",
    "interface": [
      {
        "constant": true,
        "inputs": [],
        "name": "name",
        "outputs": [
          {
            "name": "",
            "type": "string"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "",
            "type": "address"
          },
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "name": "cardsHeldByOwner",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "totalSupply",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "_owner",
            "type": "address"
          },
          {
            "name": "_attributes",
            "type": "uint256"
          }
        ],
        "name": "createCard",
        "outputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "_tokenId",
            "type": "uint256"
          }
        ],
        "name": "ownerOf",
        "outputs": [
          {
            "name": "owner",
            "type": "address"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "_owner",
            "type": "address"
          }
        ],
        "name": "balanceOf",
        "outputs": [
          {
            "name": "count",
            "type": "uint256"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "",
            "type": "uint256"
          }
        ],
        "name": "cardIndexToOwner",
        "outputs": [
          {
            "name": "",
            "type": "address"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [
          {
            "name": "_owner",
            "type": "address"
          }
        ],
        "name": "tokensOfOwner",
        "outputs": [
          {
            "name": "ownerTokens",
            "type": "uint256[]"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "symbol",
        "outputs": [
          {
            "name": "",
            "type": "string"
          }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "_to",
            "type": "address"
          },
          {
            "name": "_cardId",
            "type": "uint256"
          }
        ],
        "name": "transfer",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "from",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "to",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "tokenID",
            "type": "uint256"
          }
        ],
        "name": "Transfer",
        "type": "event"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "cardID",
            "type": "uint256"
          },
          {
            "indexed": false,
            "name": "creationBattleID",
            "type": "uint128"
          },
          {
            "indexed": false,
            "name": "attributes",
            "type": "uint256"
          }
        ],
        "name": "NewCard",
        "type": "event"
      }
    ]
  }
}; export default contracts;