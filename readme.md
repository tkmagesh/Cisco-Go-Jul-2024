# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com
- https://www.linkedin.com/in/tkmagesh

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:40 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hour)
- Tea Break     : 03:00 PM (20 mins)
- Wind up       : 04:30 PM 

## Methodolgy
- No powerpoint presentations
- Discuss & Code
- No dedicated time for Q & A (the floor is open at all times)

## Software Requirements
- Go Tools (https://go.dev/dl)
    - Verification
        ```
        $ go version
        ```
- Any Editor (Visual Studio Code)

## Repository
- https://github.com/tkmagesh/cisco-go-jul-2024

## Why Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers (public/private/protected etc)
    - No reference types (everything is a value in go)
    - No pointer arithmatic 
    - No classes (only structs)
    - No inheritance (only compostion)
    - No execptions (only errors)
    - No try..catch..finally construct
    - No implicit type conversions
- Managed Concurrency
    - Concurrency Programming
        - Application with more than one execution path
        - Typically achieved by creating multiple OS threads
        - OS thread management is taken care by the OS thread scheduler
        ![image Thread based concurrency](./images/Traditional_concurrency.png)
    - Builtin Scheduler
    ![image Managed concurrency](./images/managed_concurrency.png)

    - lightweight goroutines (~4KB)
    - Concurrency features are built in the language
        - go keyword, channel (chan) data type, channel operator ( <- ), range , select-case
        - Standard Library APIs
            - "sync" package
            - "sync/atomic" packages

- Close to hardware
    - Compile to machine code
    - Less startup time
    - One build per deployment platform 
    - Builtin tooling support for cross-compilation

## Program Structure

### Compile a go program
```
go build [file_name.go]
```
```
go build -o [binary_name] [file_name.go]
```

### Compile & Execute
```
go run [file_name.go]
```

## Data Types



