/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/
 
 package main
  
 import (
     "io"
     "net/http"
     "fmt"
 )
  
 func hello(w http.ResponseWriter, r *http.Request) {
     io.WriteString(w, "Hello cloud environments from Go!")
 }
  
 func main() {
     portNumber := "9000"
     http.HandleFunc("/", hello)
     fmt.Println("Server listening on port ", portNumber)
     http.ListenAndServe(":" + portNumber, nil)
    }