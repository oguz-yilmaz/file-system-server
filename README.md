## File System Server

Welcome to the File System Server Daemon! This project is a simple file system
server written in Golang. It allows you to create, edit, and delete files and
directories by sending JSON data. Communication with the server is done using
the JSON-RPC protocol.

### Features

- **Create Files**: Create new files with specified content, type, and permissions.
- **Edit Files**: Modify existing files.
- **Delete Files**: Remove files from the file system.
- **Directory Management**: Handle files within specified directories.

### Getting Started

#### Installation

- Download the binary from the [releases
page](https://github.com/oguz-yilmaz/file-system-server/releases)

Or

- Clone the repository and build the binary with the following command:

```bash
git clone https://github.com/oguz-yilmaz/file-system-server.git
cd file-system-server
go build -o fss cmd/fss/main.go
```

#### Running the Server

- Run the server with the following command:

```bash
./fss -c tcp -a localhost:8080 -d /tmp
```

- The server will start listening on `localhost:8080` and set the root
directory to `/tmp`.

#### Usage

The server listens for JSON-RPC requests. Here are some example requests you
can use to interact with the server.

##### Create File

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "createFile",
  "params": {
    "name": "example.txt",
    "dir": ".",
    "root": "/tmp",
    "content": "this is an example file.",
    "file-type": "txt",
    "permissions": 438,
    "overwrite": true
  }
}
```

### JSON-RPC Methods

#### createFile: Creates a new file.

- **name**: Name of the file. If you only want to create empty directories, you
can set the name to an empty string or omit entirely.
- **dir**: Directory where the file will be created (optional).
- **content**: Content of the file.
- **file-type**: Type of the file (optional).
- **root**: Root directory of the file system. Only used if `dir` is relative
(optional).
- **permissions**: Permissions for the file (optional).
- **overwrite**: Boolean to indicate if existing files should be overwritten
(optional).

#### editFile: Edits an existing file.

- (similar parameters as createFile)

#### deleteFile: Deletes a specified file.

- **name**: Name of the file.
- **dir**: Directory where the file is located (optional).

### Example JSON-RPC Client

You can use tools like curl or any JSON-RPC client library in your preferred
language to send requests to the server.

#### Using curl

```bash

curl -X POST -H "Content-Type: application/json" -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "createFile",
  "params": {
    "name": "example.txt",
    "dir": ".",
    "content": "this is an example file.",
    "file-type": "txt",
    "permissions": 438,
    "overwrite": true
  }
}' localhost:8080
```

#### Using stdin and stdout

You can also interact with the server using stdin and stdout. Here is an
example of creating a file using stdin:

```bash
echo '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "createFile",
  "params": {
    "name": "example.txt",
    "dir": ".",
    "content": "this is an example file.",
    "file-type": "txt",
    "permissions": 438,
    "overwrite": true
  }
}' | ./fss -c stdin
```

or in a node.js script:

```javascript
const { spawn } = require('child_process');

const fss = spawn('./fss', ['-c', 'stdin']);

fss.stdin.write('{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "createFile",
  "params": {
    "name": "example.txt",
    "dir": ".",
    "content": "this is an example file.",
    "file-type": "txt",
    "permissions": 438,
    "overwrite": true
  }
}');
```

### Contributing

Feel free to fork the repository and submit pull requests. For major changes,
please open an issue first to discuss what you would like to change.

### License

This project is licensed under the MIT License.

