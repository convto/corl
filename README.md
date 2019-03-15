# corl
This is cmd for url convert. It supports decode and encode.

corl means **co**nvert U**RL**.

## installation
```
$ go get -u github.com/srttk/corl
```

## usage
### -h
Print help.

### -d $URL
Decode `$URL`, eg.
```
$ coul -d https%3A%2F%2Fgithub.com%2Fsrttk
https://github.com/srttk
```

### -e $URL
Encode `$URL`, eg.
```
$ coul -e https://github.com/srttk
https%3A%2F%2Fgithub.com%2Fsrttk
```

## LICENSE
MIT
