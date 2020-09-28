
# Using my own module


Implement the package : [go_calc_package](https://github.com/davidramos-om/go_calc_package)

1. Start a new project

2. Run
```
    go mod init your-module-name
    go mod init implement-go_calc_package
```
3. Previous step has created a new **go.mod** file, open it

4. Add a line followed by : *"module implement-go_calc_package"*
  ```
    require github.com/davidramos-om/go_calc_package v1.0.1
  ```

5. Run  
    ```
    go run main.go
    ```

5. A new file **go.sum** have been added

6. Add a new main file
  
  ```
    package main

    import "github.com/davidramos-om/go_calc_package"

    func main() {
      // Using a provided func by the module
      go_calc_package.ShowMenu()
    }

  ```

7. Run your application
```
  go run main.go 
```