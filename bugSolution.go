func main() {

        fmt.Println("Before the panic")

        defer func() {
                if r := recover(); r != nil {
                        fmt.Println("Recovered from panic:", r)
                        // Handle the panic here.  For example, log the error, retry the operation, or gracefully exit.
                }
        }()

        panic("This is a panic!")

        //This will not be executed even with recover
        fmt.Println("After the panic")

}        

