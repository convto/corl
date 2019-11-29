# corl
This is cmd for url convert. It supports decode and encode.

corl means **co**nvert U**RL**.

## installation
```
$ go get -u github.com/convto/corl
```

## usage
### -h
Print help.

### -d $URL
Decode `$URL`, eg.
```
$ coul -d https%3A%2F%2Fgithub.com%2Fconvto
https://github.com/convto
```

### -e $URL
Encode `$URL`, eg.
```
$ coul -e https://github.com/srttk
https%3A%2F%2Fgithub.com%2Fconvto
```

## LICENSE
MIT
