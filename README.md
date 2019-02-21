# Go Split Join (GoSJ) 
GoSJ is a command line file splitting and joining utility, written in Go.

## Features
- Split large files into small files of your desired filesize, real fast.
- Use the utility straight from your command line.
- Utility generates a text file with file hashes making it easy for users to verify that the end file isn't corrupted.

## Usage
1. Clone the repository:
```bash
git clone https://github.com/ImRaviTejaG/go-split-join.git
```

2. To run directly, build the project as follows:
```bash
go build main.go
```

followed by:
```bash
main -opt=split -file=test.zip -size=50
```

3. Else, you can run the file using:
```bash
go run main.go -opt=split -file=test.zip -size=50
```

4. For more help, use the `-h` flag
```bash
go run main.go -h
```

which would output something like:
```
Usage of C:\******\main.exe:
  -file string
        The option: SPLIT / JOIN (default "file.txt")
  -opt string
        The option: SPLIT / JOIN (default "SPLIT")
  -parts int
        Number of parts, mandatory for JOINing files
  -size int
        Split size, mandatory for SPLITing files (in MB)
exit status 2
```

__NOTE__: Please take special care of the `GOROOT` and `GOPATH` variables. In case you get errors, simply create a new folder named `gosj` in the `GOPATH` and copy `main.go` file & the `sj` folder into the `gosj` directory. This should get the utility working fine.

---
#### Find the Python implementation of the file splitting and joining utility here: https://github.com/ImRaviTejaG/pysj.
#### _In case you face trouble, please feel free to open an issue._
