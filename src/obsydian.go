package main

import (
    "io"
    "log"
    "net/http"
	"os"
	"fmt"
)

func main() {
	var userinput string
	fmt.Print("site:")
	fmt.Scan(&userinput)
    response, err := http.Get(userinput)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    outFile, err := os.Create("output.html")
    if err != nil {
        log.Fatal(err)
    }
    defer outFile.Close()
    _, err = io.Copy(outFile, response.Body)
    if err != nil {
        log.Fatal(err)
	}
    fmt.Print(`    
                                                           

                                                                                                                                                               
         
     
     
     
     
     
                    ,╓▄▄▄,
                    ╒▀    ▌ ▀█▀w╓
                    █    █    ▀▄  ╙M▄
                   ▐'   ▐       ╙N   ▀,
                   ▌    ▌         '▀,  N
                  ▐    █            █▀Nç▀▄
                  ▌   ▐▄            █   ▀▀█
                 █ ,▄▀█ ▀▀N▄,       ▌     ▄█▌ç
                 █▀   █      └▀N▄, j▌ ▄Æ▀"  █▐
                 ▌└▀▀N█mmAARMP▀▀▀▀▀▀▀'      ▐"µ
                 █    ▐µ           ▐═       ▐L▌
                 ▐     █           ▐═        ▌█
                  ▌    ╘▌          ▐░        █▐
                  ▐     █          ▐░        ▐U▌
                   ▄     ▌         ▐U       ,▄▀
                     ╙M▄ ▐         ▐U   ╓m╜└
                         ╙▀ⁿ*++∞∞w▄▐▌ª╙                                                                                                                                                                                                                           
                                                                                                                                                                                                      
                                                                                                          
    ███████████████████████████████████████████████████████████████
    █████████████████████obsydian has finished█████████████████████
    ███████████████████████████████████████████████████████████████
                                      
                                      
                                      `)
}
