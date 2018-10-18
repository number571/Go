package main

/*
    [WRITE_BLOCK]
        ./main -w from to summ
    
    [READ_BLOCKS]
        ./main -r
*/

import (
    "crypto/sha256"
    "encoding/json"
    "encoding/hex"
    "io/ioutil"
    "strings"
    "strconv"
    "fmt"
    "os"
)

const (
    BLOCKCHAIN_DIR = "Blockchain/"
    EXTENSION = ".json"
    BUFF = 512
)

type Block struct {
    From string
    To   string
    Summ string
    Hash string
}

func main () {
    checkArgs(os.Args)
    os.Mkdir(BLOCKCHAIN_DIR, 0777)

    switch (os.Args[1]) {
        case "-r", "--read":
            readBlocks(getLastBlock(BLOCKCHAIN_DIR))
        case "-w", "--write":
            writeBlock(getLastBlock(BLOCKCHAIN_DIR), os.Args[2:])
    }
}

func checkArgs (args []string) {
    if len(args) < 2 {
        getError("len(args) < 2")
    }
    switch (args[1]) {
        case "-r", "--read":
            return
        case "-w", "--write":
            if len(args[1:]) != 4 {
                getError("len(args) for mode '-w' != 3")
            }
        default:
            getError("argument is not r/w")    
    }
}

func readBlocks (last_block int) {
    var (
        lines_of_content []string
        list_of_hashes []string
        data_json Block
    )

    if last_block > 1 {
        fmt.Println("Checked from blocks:")
        
        for i := 2; i <= last_block; i++ {
            json.Unmarshal([]byte(getFileContent(getFilename(i))), &data_json)

            list_of_hashes = append(
                list_of_hashes,
                getHash(getFileContent(getFilename(i - 1))),
            )

            if (data_json.Hash == list_of_hashes[i - 2]) {
                fmt.Printf("[Block: %d] -> Readable\n", i - 1)
            } else {
                fmt.Printf("[Block: %d] -> Corrupted\n", i - 1)
            }

        }
    }

    list_of_hashes = append(
        list_of_hashes,
        getHash(getFileContent(getFilename(last_block))),
    )

    lines_of_content = strings.Split(getFileContent(getFilename(0)), "\n")

    fmt.Println("\nChecked from zero-block:")

    for i := 0; i < last_block; i++ {
        if list_of_hashes[i] == lines_of_content[i] {
            fmt.Printf("[Block: %d] -> Readable\n", i + 1)
        } else {
            fmt.Printf("[Block: %d] -> Corrupted\n", i + 1)
        }
    }
}

func writeBlock (last_block int, args []string) {
    var (
        new_block string = getFilename(last_block + 1)
        hash string = ""
    )

    file, err := os.Create(new_block)
    checkError(err)
    defer file.Close()

    if last_block > 0 {
        hash = getHash(getFileContent(getFilename(last_block)))
    }

    x, _ := json.Marshal(Block{args[0], args[1], args[2], hash});
    file.WriteString(string(x))

    zeroBlock(getHash(getFileContent(new_block)))
}

func zeroBlock (hash string) {
    var zero string = getFilename(0)

    if fileIsNotExist(zero) {
        createFile(zero)
    }

    file, err := os.OpenFile(zero, os.O_WRONLY|os.O_APPEND, 0777)
    checkError(err)
    defer file.Close()

    file.WriteString(hash + "\n")
}

func fileIsNotExist (filename string) bool {
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        return true
    }
    return false
}

func createFile (filename string) {
    file, err := os.Create(filename)
    checkError(err)
    file.Close()
}

func getHash (message string) string {
    hash := sha256.New() 
    hash.Write([]byte(message))
    return hex.EncodeToString(hash.Sum(nil))
}

func getFilename (base int) string {
    return BLOCKCHAIN_DIR + strconv.Itoa(base) + EXTENSION
}

func getFileContent (filename string) string {
    file, err := os.Open(filename)
    checkError(err)
    defer file.Close()

    var buffer []byte = make([]byte, BUFF)
    var message string = ""
    
    for {
        length, err := file.Read(buffer)
        if err != nil || length == 0 { break }
        message += string(buffer[:length])
    }

    return message
}

func getLastBlock (dir string) int {
    files, err := ioutil.ReadDir(dir)
    checkError(err)

    var max int = 0
    for _, file := range files {
        max_block, _ := strconv.Atoi(getFilebase(file.Name()))
        if max_block > max {
            max = max_block
        }   
    }
    return max
}

func getFilebase (file string) string {
    return strings.Split(file, ".")[0]
}

func checkError (err error) {
    if err != nil {
        getError(err.Error())
    }
}

func getError (err string) {
    fmt.Println("Error:", err)
    os.Exit(1)
}
