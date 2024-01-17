package main

import (
    "os"
    "strconv"
)

func main() {

    argInt, _ := strconv.Atoi(os.Args[1]) 
    for i := '1'; i <= '9'; i++ {         
        oneTo9, _ := strconv.Atoi(string(i))                                                            
        res := string(i) + " " + "x" + " " + os.Args[1] + " " + "=" + " " + strconv.Itoa(argInt*oneTo9) 
        os.Stdout.Write([]byte(res + "\n"))                                                    
    }
}