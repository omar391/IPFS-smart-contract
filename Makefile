# install solc and abigen cmd tools first
all:install_solc install_abigen compile install_ganche
install_abigen:
	git clone --depth 1 -b "v1.10.7" https://github.com/ethereum/go-ethereum.git build/go-ethereum
	$(MAKE) -C build/go-ethereum
	$(MAKE) -C build/go-ethereum devtools
	rm -rf build/go-ethereum

install_solc:
#	its highly system dependant, so do it seperately
#	Lik: https://github.com/ethereum/solidity/releases

install_ganche:
#	its highly system dependant, so do it seperately
#	Link: https://github.com/trufflesuite/ganache/releases

compile:
	solc --overwrite --optimize --abi ./contracts/cid_storage.sol -o build
	solc --overwrite --optimize --bin ./contracts/cid_storage.sol -o build
	abigen --abi=./build/CIDStorage.abi --bin=./build/CIDStorage.bin --pkg=api --out=./api/cid_storage.go