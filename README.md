# dirInfo
CLI in Go for hashing of files in a directory.

Will go through either CWD or specified directory and al sub-directories outputing the results to a json file.


## Commands

**--duplicate, -d** Looks for duplicates. Uses md5 hashing if no hashes are selected.

**--path *value*, -p *value*** Root directory for the process. Defaults to CWD if none specified.

**--info, -i** Get the system info for the file.

**--noArray, --na** Doesn't output data as JSON Array. (Useful for mongoimport with files larger than 16MB)

**--note *value*, -n *value*** Useful for putting a note (Example: comparing two servers, working and notWorking)

**--output *value*, -o *value*** Output file(Example -o foo will create foo.json)

**--talkative,-t** Verbose output

**--help, -h** CLI help

**--version, -v** Version info


### Hashing Flags
  **--sha1, -s** SHA-1 hashing of the file. *Default*

  **--sha256, --s2** SHA-256 hasing of the file.

  **--md5, -m** MD5 hashing of the file.


## Examples

### Regular Execution

Running `go run *.go --s2 -s -m -note 'for readme.md'` in the project directory will get the following output.

```bash
File Read Completed in 55.741498ms
Data written to: iMac-results.json
```

```json
[
   {
      "directory":"~/git-repo/src/dirInfo/",
      "filepath":"~/git-repo/src/dirInfo/LICENSE",
      "name":"LICENSE",
      "size":1076,
      "mode":420,
      "machineName":"iMac",
      "md5":"e101e9a69504fcc422e363f49cb48827",
      "sha1":"152e742349f2eef53bd53ef9ee7cb3c08572a29d",
      "sha256":"8f1f67e58e674ca7b9beebb86736eb391d5dafb31e53231106c54e4e239ae105",
      "lastModified":"2016-08-29T21:44:52-07:00",
      "note":"for readme.md"
   },
   {
      "directory":"~/git-repo/src/dirInfo/",
      "filepath":"~/git-repo/src/dirInfo/README.md",
      "name":"README.md",
      "size":1658,
      "mode":420,
      "machineName":"iMac",
      "md5":"c75d7ab4e02ae1936ea67248e8883496",
      "sha1":"68f291662caa826fa2c61fac6ae114000deaddef",
      "sha256":"42840517bce7bea4433583ea27e5ee323469535102a4ce760038fb101bf2dbdb",
      "extension":".md",
      "lastModified":"2016-08-29T22:00:22-07:00",
      "note":"for readme.md"
   },
   {
      "directory":"~/git-repo/src/dirInfo/",
      "filepath":"~/git-repo/src/dirInfo/main.go",
      "name":"main.go",
      "size":5213,
      "mode":420,
      "machineName":"iMac",
      "md5":"7001ab02606cb481c2bfb840abafdae3",
      "sha1":"fb87ef900453755e5d7dc219e40b08e4f715c230",
      "sha256":"b44ac4fa3d4176c61f3004c6ba3ddf547209dc2415c4af5c2eba1dd3f6c679a8",
      "extension":".go",
      "lastModified":"2016-08-29T21:53:05-07:00",
      "note":"for readme.md"
   }
]
```

### No Array

Running `go run *.go --s2 -s -m -note 'for readme.md' --na` in the project directory will get the following output.

```bash
File Read Completed in 55.741498ms
Data written to: iMac-results.json
```

```json
{"directory":"~/git-repo/src/dirInfo/","filepath":"~/git-repo/src/dirInfo/LICENSE","name":"LICENSE","size":1076,"mode":420,"machineName":"iMac","md5":"e101e9a69504fcc422e363f49cb48827","sha1":"152e742349f2eef53bd53ef9ee7cb3c08572a29d","sha256":"8f1f67e58e674ca7b9beebb86736eb391d5dafb31e53231106c54e4e239ae105","lastModified":"2016-08-29T21:44:52-07:00","note":"for readme.md"}
{"directory":"~/git-repo/src/dirInfo/","filepath":"~/git-repo/src/dirInfo/README.md","name":"README.md","size":2830,"mode":420,"machineName":"iMac","md5":"d48289da348053021066042ad7a2fbc4","sha1":"49482c8ce9352a7b0eb9923790941e7951cd9632","sha256":"0cdfe499b98c2d4686751989e6f5fafba83d9bacff541bc0098bf130041b0fc3","extension":".md","lastModified":"2016-09-25T16:38:26-07:00","note":"for readme.md"}
{"directory":"~/git-repo/src/dirInfo/","filepath":"~/git-repo/src/dirInfo/dirInfo.go","name":"dirInfo.go","size":6631,"mode":420,"machineName":"iMac","md5":"fa9faf00f702e37c27ed6c5eccef32e1","sha1":"c5ad726490b070a54a143e7b76d154b3772a8503","sha256":"2b766597ffcec440fd93d445ba55e22d58e68703d656a1f2784909cbaa95e7e2","extension":".go","lastModified":"2016-09-25T16:29:01-07:00","note":"for readme.md"}
```

