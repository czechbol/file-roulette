# file-roulette

Give it a folder and it will pick a random file inside. Give it folder and index and it will return a file with that index. 

I have no idea what am I even doing with my time.


## Golang program usage

```console
foo@bar:~$ file-roulette  -h 
file-roulette - Recursively selecting random files in a given folder

Give it a folder and it will pick a random file inside. 
Give it folder and index and it will return a file with that index.

I have no idea what am I even doing with my time.

Usage:
  file-roulette [flags]

Flags:
  -h, --help            help for file-roulette
  -i, --index int       Do not perform download test. Zero to disable
  -v, --verbose count   Logging verbosity. Specify multiple times for higher verbosity
      --version         version for file-roulette
```

```console
foo@bar:~$ file-roulette foo/
Total number of files in provided folder: 68
Choosing a random file:
foo/bar/image.jpg
```

```console
foo@bar:~$ file-roulette -i 5 foo/
Total number of files in provided folder: 68
Selected file
foo/bar/image7.jpg
```


## Bash script usage

```console
foo@bar:~$ sh roulette.sh -s foo/
Total number of files in provided folder: 79
I'm going to choose a random file now because you didn't specify an index. (-i number)

Your roulette file iiiiiiiiiiis:
foo/bar.jpg
```

```console
foo@bar:~$ sh roulette.sh -s foo/ -i 5
Your roulette file iiiiiiiiiiis:
foo/bar/image.jpg
```

```console
foo@bar:~$ sh roulette.sh
I need some input you know? (-s source_folder/ and -i number)
```
