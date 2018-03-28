abigen --sol=../../eth_contracts/contracts/CardBase.sol --pkg=contracts --out=../../ethereum_proxy/src/GoRpc/contracts/cardbase.go
echo "Created cardbase.go"
abigen --sol=../../eth_contracts/contracts/Battles.sol --pkg=contracts --out=../../ethereum_proxy/src/GoRpc/contracts/battles.go
echo "Created battles.go"
abigen --sol=../../eth_contracts/contracts/BattleGroups.sol --pkg=contracts --out=../../ethereum_proxy/src/GoRpc/contracts/battlegroups.go
echo "Generate battlegroups.go"
abigen --sol=../../eth_contracts/contracts/GenerateCard.sol --pkg=contracts --out=../../ethereum_proxy/src/GoRpc/contracts/generatecard.go
echo "Generate generatecard.go"

