# coul
This is cmd for url convert. It supports decode and encode.

coul means **co**nvert **U**R**L**.

## installation
```
$ go get -u github.com/srttk/coul
```

## usage
### -h
Print help.

### -d $URL
url decode `$URL`, eg.
```
$ coul -d https%3A%2F%2Fgithub.com%2Fsrttk
https://github.com/srttk
```

### -e $URL
url encode `$URL`, eg.
```
$ coul -e https://github.com/srttk
https%3A%2F%2Fgithub.com%2Fsrttk
```

## LISENCE
MIT
