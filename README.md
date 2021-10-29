[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

# gwntools
simple pwntools-like framework for go to interact with local and remote processes

&nbsp;
&nbsp;

## INSTALL
`go get -u github.com/r0lh/gwntools`

&nbsp;
&nbsp;
## USAGE EXAMPLES
&nbsp;
#### START LOCAL BINARY WITH COMMAND-LINE ARGUMENT
```golang
p, err := gwntools.Local(exec.Command("./my_local_binary", "--this-is-an-argument"))
if err != nil {
 log.Fatal(err)
}
```
&nbsp;
#### CONNECT TO REMOTE TCP SERVICE
```golang
p, err := gwntool.Remote("127.0.0.1:4444")
if err != nil {
  log.Fatal(err)
}
```
&nbsp;
#### WRITE TO LOCAL OR REMOTE PROCESS
```golang
err := p.Write([]byte("AAAAAAPAYLOADAAAAAA\x41\x42\x43\x44\n")
if err != nil {
  log.Fatal(err)
}
```
&nbsp;
#### READ LINE FROM LOCAL OR REMOTE PROCESS
```golang
r, err := p.ReadLine()
if err != nil {
  log.Fatal(err)
}

fmt.Printf("Got response from process: %s", string(r))
```
&nbsp;
#### INTERACT WITH LOCAL OR REMOTE PROCESS
```golang
err := p.Interactive()
if err != nil {
  log.Fatal(err)
}
```
