# dirInfo
CLI in Go for hashing of files in a directory.

Will go through either CWD or specified directory and al sub-directories outputing the results to a json file.


## Commands
**--path *value*, -p *value*** Root directory for the process. Defaults to CWD if none specified.

**--info, -i** Get the system info for the file.

**--note *value*, -n *value*** Useful for putting a note (Example: comparing two servers, working and notWorking)

**--output *value*, -o *value*** Output file(Example -o foo will create foo.json)

**--talkative,-t** Verbose output

**--help, -h** CLI help

**--version, -v** Version info


### Hashing Flags
  **--sha1, -s** SHA-1 hashing of the file.

  **--sha256, --s2** SHA-256 hasing of the file.

  **--md5, -m** MD5 hashing of the file.


## Example

Running `go run *.go --s2 -s -m -note 'for readme.md'` in the project directory will get the following output.

```
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
