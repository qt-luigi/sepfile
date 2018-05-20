# sepfile

sepfile separates a target file by a keyword line.

Each separated file names are added a dot and a sequence number after the target file name.

## Installation

When you have installed the Go, Please execute following `go get` command:

```sh
go get -u github.com/qt-luigi/sepfile
```

## Usage

```sh
$ sepfile
sepfile separates a target file by a keyword line.

Each separated file names are added a dot and a sequence number after the target file name.

Usage:

	sepfile <srcfile> <kwdline> [<outpath>]

Each arguments are:

	<srcfile>
		a target file.
	<kwdline>
		a keyword line for separating.
		if an error occurred then surround it by '"'.
	[<outpath>]
		an output file path.
		default is ".".

```

Check to exist the target file.

```sh
$ ls
target.html
```

See the target file.

```sh
$ cat target.html
<html>
<head>
  <title>sepfile test</title>
</head>
<body>
  foo
<br>
  bar
  <br>
  baz
</body>
</html>
```

When you separate the target.html by "<br>" line of between "foo" to "bar", Execute following `sepfile` command:

```sh
$ sepfile target.html "<br>"
```

If you want to create the separating files in another directory, Please specify output path at the third argument.

```sh
$ sepfile target.html "<br>" ../
```

Create the separating files in current directory when argument length is two.

```sh
$ ls
target.html   target.html.1 target.html.2
```

See the separating files.

```sh
$ cat target.html.1 
<html>
<head>
  <title>sepfile test</title>
</head>
<body>
  a
```

```sh
$ cat target.html.2
  b
  <br>
  c
</body>
</html>
```

You specified "<br>" line, not "  <br>" line. A keyword line must equals completely.

## License

MIT

## Author

Ryuji Iwata

## Note

This tool is mainly using by myself. :-)

