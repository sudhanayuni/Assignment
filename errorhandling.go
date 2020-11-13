package main
import "fmt"
import "os"
import "errors"

//function accepts a filename and tries to open it.
func fileopen(name string) (string, error) {
    f, er := os.Open(name)

    //er will be nil if the file exists else it returns an error object  
    if er != nil {
        //created a new error object and returns it  
        return "", errors.New("Custom error message: File name is wrong")
    }else{
    	return f.Name(),nil
    }
}

func main() {  
    //receives custom error or nil after trying to open the file
    filename, error := fileopen("invalid.txt")
    if error != nil {
        fmt.Println(error)
    }else{
    	fmt.Println("file opened", filename)
    }  
}