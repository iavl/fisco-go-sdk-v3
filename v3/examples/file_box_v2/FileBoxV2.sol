// SPDX-License-Identifier: MIT
// solhint-disable comprehensive-interface,not-rely-on-time,gas-custom-errors
pragma solidity 0.8.11;

contract FileBoxV2 {
    struct User {
        address user;
        bool deleted;
        string userType; // super admin, admin, member, guest
        uint256 createdTimestamp;
        uint256 updatedTimestamp;
    }

    struct File {
        bytes32 fileHash;
        address owner;
        bool deleted;
        string storageSpaceType; // organization, group, person
        address storageSpaceAddress;
        string extra;
        uint256 uploadTimestamp;
        uint256 updatedTimestamp;
    }

    mapping(address => User) internal _users;
    mapping(address => mapping(bytes32 => File)) internal _fileInfo;

    // events
    event UserAdded(address indexed user, string userType, address admin);
    event UserDeleted(address indexed user, address admin);

    event FileUploaded(
        address indexed user,
        bytes32 indexed fileHash,
        string storageSpaceType,
        address storageSpaceAddress,
        string extra,
        uint256 timestamp
    );
    event FileShared(
        address indexed user,
        bytes32 indexed fileHash,
        address indexed to,
        string shareType,
        string extra,
        uint256 timestamp
    );
    event FileDeleted(address indexed user, bytes32 indexed fileHash, uint256 timestamp);
    event FileRecovered(address indexed user, bytes32 indexed fileHash, uint256 timestamp);

    modifier validUser(address user) {
        require(_users[user].createdTimestamp > 0 && !_users[user].deleted, "User not found");
        _;
    }

    /// @notice add a new user
    /// @param user user address
    /// @param userType user type
    /// @param admin admin address
    function addUser(address user, string memory userType, address admin) public {
        require(_users[user].createdTimestamp == 0 || _users[user].deleted, "User already exists");

        _users[user] = User({
            user: user,
            deleted: false,
            userType: userType,
            createdTimestamp: block.timestamp,
            updatedTimestamp: block.timestamp
        });

        emit UserAdded(user, userType, admin);
    }

    /// @notice delete a user
    /// @param user user address
    /// @param admin admin address
    function deleteUser(address user, address admin) public validUser(user) {
        _users[user].deleted = true;
        _users[user].updatedTimestamp = block.timestamp;

        emit UserDeleted(user, admin);
    }

    /// @notice upload a file
    /// @param user user address
    /// @param fileHash file hash
    /// @param signature user signature
    /// @param storageSpaceType storage space type: organization, group, person
    /// @param storageSpaceAddress storage space address, can be organization, group, or user
    /// address(null)
    function uploadFile(
        address user,
        bytes32 fileHash,
        bytes memory signature,
        string memory storageSpaceType,
        address storageSpaceAddress,
        string memory extra
    ) public validUser(user) {
        signature;

        _fileInfo[user][fileHash] = File({
            fileHash: fileHash,
            owner: user,
            deleted: false,
            storageSpaceType: storageSpaceType,
            storageSpaceAddress: storageSpaceAddress,
            extra: extra,
            uploadTimestamp: block.timestamp,
            updatedTimestamp: block.timestamp
        });

        emit FileUploaded({
            user: user,
            fileHash: fileHash,
            storageSpaceType: storageSpaceType,
            storageSpaceAddress: storageSpaceAddress,
            extra: extra,
            timestamp: block.timestamp
        });
    }

    function shareFile(
        address user,
        bytes32 fileHash,
        address to,
        string memory shareType,
        string memory extra
    ) public validUser(user) {
        File memory file = _fileInfo[user][fileHash];
        require(file.uploadTimestamp > 0, "File not found");
        require(!file.deleted, "File is deleted");

        emit FileShared({
            user: user,
            fileHash: fileHash,
            to: to,
            shareType: shareType,
            extra: extra,
            timestamp: block.timestamp
        });
    }

    function deleteFile(address user, bytes32 fileHash) public {
        File memory file = _fileInfo[user][fileHash];
        require(file.uploadTimestamp > 0, "File not found");
        require(!file.deleted, "File is deleted");

        _fileInfo[user][fileHash].deleted = true;
        _fileInfo[user][fileHash].updatedTimestamp = block.timestamp;

        emit FileDeleted(user, fileHash, block.timestamp);
    }

    function recoverFile(address user, bytes32 fileHash) public validUser(user) {
        File memory file = _fileInfo[user][fileHash];
        require(file.deleted, "File is not deleted");

        _fileInfo[user][fileHash].deleted = false;
        _fileInfo[user][fileHash].updatedTimestamp = block.timestamp;

        emit FileRecovered(user, fileHash, block.timestamp);
    }

    function getFile(address user, bytes32 fileHash) public view returns (File memory) {
        return _fileInfo[user][fileHash];
    }

    function getUser(address user_) public view returns (User memory) {
        return _users[user_];
    }
}
