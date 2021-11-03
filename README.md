[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

# gwntools
simple pwntools-like framework for go to interact with local and remote processes

&nbsp;
&nbsp;

## INSTALL
`go get -u github.com/r0lh/gwntools`

&nbsp;
&nbsp;
## GETTING STARTED
&nbsp;

#### START LOCAL PROCESS
```golang
package main

import github.com/r0lh/gwntools

func main() {

    # start local binary with command-line argument
    p := gwntools.Local(exec.Command("./my_local_binary", "--this-is-an-argument"))
     
    # read line from process
    r := p.ReadLine()
    
    fmt.Printf("banner from process: %s", string(r))
 
    # write to process
    p.Write([]byte("AAAAAAPAYLOADAAAAAA\x41\x42\x43\x44\n")
    
    # interact with process
    p.Interactive()
    
}
```
&nbsp;
#### CONNECT TO REMOTE TCP SERVICE
```golang
package main

import github.com/r0lh/gwntools

func main() {
    p := gwntool.Remote("127.0.0.1:4444")
        
    # read line from remote service
    r := p.ReadLine()
    
    fmt.Printf("banner from server: %s", string(r))
    
    # write to remote service
    p.Write([]byte("AAAAAAPAYLOADAAAAAA\x41\x42\x43\x44\n")
        
}
```
