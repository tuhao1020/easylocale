# easylocale
a simple locale detector for current system

## Usage
```
go get github.com/tuhao1020/easylocale
```

```go
locale, err = easylocale.CurrentLocale()
if err != nil {
  fmt.Println(err.Error())
} else {
  fmt.Println(locale)
}
```
