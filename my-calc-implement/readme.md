
#Using my own package


Implement the package : [go_calc_package](github.com/davidramos-om/go_calc_package)

1. Start a new project
  ```
    package main

    import "github.com/davidramos-om/go_calc_package"

    func main() {
      go_calc_package.ShowMenu()
    }

  ```
2. Run
```
    go mod init your-module-name
    go mod init implement-go_calc_package
```
3. This step has created a new go.mod file, open it
4. Add next line follow "module implement-go_calc_package"
  ```require github.com/davidramos-om/go_calc_package v1.0.1```

5. Run  
    ```
    go run main.go
    ```
5. A new file "go.sum" have been added

6. Run your application
```
  go run main.go 
```