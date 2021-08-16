// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve cid value in a variable
 */
contract CIDStorage {
    string private cid_ = "";

    /**
     * @dev We aim to store the CID value for once only
     * @param cid value to store
     */
    function store(string memory cid) public returns (string memory) {
        if (bytes(cid_).length == 0 && bytes(cid).length != 0) {
            cid_ = cid;
            return string(abi.encodePacked("Stored new CID: ", cid));
        }
        return "Nothing's Stored!";
    }

    /**
     * @dev Return value
     * @return value of 'cid_'
     */
    function retrieve() public view returns (string memory) {
        return cid_;
    }
}
