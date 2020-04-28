# addSome

Simple go script to add found subdomains (from Amass, bruteforced with Massdns, etc) to your Postgres database one uses for Findomain (credits to https://github.com/Edu4rdSHL/findomain for that great tool). 

Setup: 

```
- Update the const object in the Go file to your specific setup. The constant object is necessary to be able to connect to your Postgres instance. This object is line 12 to 18 in addSome.go.
- Install the golang Postgres package with go get github.com/lib/pq
```


To use it either use ```go run addSome.go [filename]``` or use ```go build addSome.go``` to be able to use it as ```./addSome [filename]```.

This tool is made to read a text file and thus expects a file which has one domain per line.

For any issues, please open one. Always welcome to open a pull request, if something can be done better! 
