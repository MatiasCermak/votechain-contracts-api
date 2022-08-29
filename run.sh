solc --overwrite --abi contracts/implementation/Vote.sol -o ./build/abi
solc --overwrite --bin contracts/implementation/Vote.sol -o ./build/bin

abigen --bin=build/bin/Vote.bin --abi=build/abi/Vote.abi --pkg=api --out=./contracts/api/Vote.go && echo "Abigen run successful, created Vote.go in directory \"./contracts/api/Vote.go\""