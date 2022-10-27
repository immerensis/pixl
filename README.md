Pixl
----------

Pixl is a pixel-art drawing tool written in golang using fyne ui-toolkit, this used to be final project of golang course from zerotomastery

----------

steps to run this program is :




Note :  windows users must install gcc ( [easiest way to install gcc](https://community.chocolatey.org/packages/mingw) is to use [chocolatey package-manager](https://chocolatey.org/install)) ) 



(execute the following commands in terminal)

```
git clone https://github.com/immerensis/pixl.git
```

then 

```
cd pixl
```

to get the dependencies

```
go mod tidy
```

run the program in verbose mode to trouble shoot by adding ```-v``` flag to ```go run``` 

```
go run -v ./pixl/
```

run the program with 
```go run ./pixl/```

----------

it is also possible to build an executable or binary of this program with ```go build``` command

```
go build ./pixl/
```
