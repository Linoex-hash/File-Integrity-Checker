# File Integrity Checker 

## Description

This is a file integrity checker, written in Golang. This will compute the SHA256 hash a single file or files within a folder, then store those hashes in a user specified sqlite3 database file. This program can then use the database to see if those files were modified, potentially unknowingly. This helps ensure file integrity. 

## How to Build

In the repository directory, run `go build` to build the binary for this executable

## Usage

```
$ ./file-integrity-checker -h

Usage of ./file-integrity-checker:
  -c    Checks the path for hash mismatch (simplified output, no color)
  -check
        Checks the path for hash mismatch
  -hash-db-loc string
        The database location for the SHA 256 hashes (default "./hashes.db")
  -init
        Initializes the database and stores SHA 256 hashes retrieved from "path"
  -update
        Updates the hash database for the path
  path
        The file/directory to retrieve the hashes of. If this is a directory, 
        it will be recursively scanned, and all contained files will be added to the database as separate entries.
```

## Examples

```
# Get the file hashes from /var/syslog

$ ./file-integrity-checker -hash-db-loc ./log_hash.db -init /var/log

Identified file /var/log/README and successfully stored its hash: 8c89ac70ed581dc651aa075af6f654a8ec3c5820f1e203219d073d13c4cb6831
Identified file /var/log/alternatives.log and successfully stored its hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
Identified file /var/log/alternatives.log.1 and successfully stored its hash: 4bd7b260e9967e8c614a10d082ca4a7b64fc4fa437c5cbe3a658aec0b5361645
Identified file /var/log/alternatives.log.2.gz and successfully stored its hash: 24995b32bf36a19acb8b071337f4f1fff24bbd353b56d03e6a8edfe20799625e
Identified file /var/log/alternatives.log.3.gz and successfully stored its hash: 4dc1fe563824bc062cbfd3e53eece2c6982159ecbdcfd6a3e3877f2c57f39973
Identified file /var/log/alternatives.log.4.gz and successfully stored its hash: 7dc29b7fb4680adb5f86f1cbc752dd5042ea9cfd0b242e79209feec9e329e125
...
```

```
# Get check the hashes from /var/syslog

$ ./file-integrity-checker -hash-db-loc ./log_hash.db -check /var/log

Identified file /var/log/README and successfully stored its hash: 8c89ac70ed581dc651aa075af6f654a8ec3c5820f1e203219d073d13c4cb6831
Identified file /var/log/alternatives.log and successfully stored its hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
Identified file /var/log/alternatives.log.1 and successfully stored its hash: 4bd7b260e9967e8c614a10d082ca4a7b64fc4fa437c5cbe3a658aec0b5361645
Identified file /var/log/alternatives.log.2.gz and successfully stored its hash: 24995b32bf36a19acb8b071337f4f1fff24bbd353b56d03e6a8edfe20799625e
Identified file /var/log/alternatives.log.3.gz and successfully stored its hash: 4dc1fe563824bc062cbfd3e53eece2c6982159ecbdcfd6a3e3877f2c57f39973
...
```

```
# Update the hashes from /var/syslog

$ ./file-integrity-checker -hash-db-loc ./log_hash.db -update /var/log

Hash for file [/var/log/README] updated successfully.
Hash for file [/var/log/alternatives.log] updated successfully.
Hash for file [/var/log/alternatives.log.1] updated successfully.
Hash for file [/var/log/alternatives.log.2.gz] updated successfully.
Hash for file [/var/log/alternatives.log.3.gz] updated successfully.
Hash for file [/var/log/alternatives.log.4.gz] updated successfully.
Hash for file [/var/log/apt/eipp.log.xz] updated successfully.
...
```