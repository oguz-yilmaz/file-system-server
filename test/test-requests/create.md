### Request 1 
```json
{
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
}
```


### Request 1 
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "createFile",
  "params": {
    "name": "another_file.txt",
    "dir": "./",
    "content": "This is another example file.",
    "overwrite": false
  }
}
```



### Request 1 
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "createFile",
  "params": {
    "name": "example2.txt",
    "content": "This is yet another example file."
  }
}
```

