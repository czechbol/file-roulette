# file_roulette.sh
Give it a folder and it will pick a random file inside. Give it folder and index and it will return a file with that index. 

I have no idea what am I even doing with my time.


## Usage
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
