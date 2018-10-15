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
    check_args(os.Args)
    os.Mkdir(BLOCKCHAIN_DIR, 0777)

    switch (os.Args[1]) {
        case "-r", "--read":
            read_blocks(get_last_block(BLOCKCHAIN_DIR))
        case "-w", "--write":
            write_block(get_last_block(BLOCKCHAIN_DIR), os.Args[2:])
    }
}

func check_args (args []string) {
    if len(args) < 2 {
        get_error("len(args) < 2")
    }
    switch (args[1]) {
        case "-r", "--read":
            return
        case "-w", "--write":
            if len(args[1:]) != 4 {
                get_error("len(args) for mode '-w' != 3")
            }
        default:
            get_error("argument is not r/w")    
    }
}

func read_blocks (last_block int) {
    var (
        lines_of_content []string
        list_of_hashes []string
        data_json Block
    )

    if last_block > 1 {
        fmt.Println("Checked from blocks:")
        
        for i := 2; i <= last_block; i++ {
            json.Unmarshal([]byte(get_file_content(get_filename(i))), &data_json)

            list_of_hashes = append(
                list_of_hashes,
                get_hash(get_file_content(get_filename(i - 1))),
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
        get_hash(get_file_content(get_filename(last_block))),
    )

    lines_of_content = strings.Split(get_file_content(get_filename(0)), "\n")

    fmt.Println("\nChecked from zero-block:")

    for i := 0; i < last_block; i++ {
        if list_of_hashes[i] == lines_of_content[i] {
            fmt.Printf("[Block: %d] -> Readable\n", i + 1)
        } else {
            fmt.Printf("[Block: %d] -> Corrupted\n", i + 1)
        }
    }
}

func write_block (last_block int, args []string) {
    var (
        new_block string = get_filename(last_block + 1)
        hash string = ""
    )

    file, err := os.Create(new_block)
    check_error(err)
    defer file.Close()

    if last_block > 0 {
        hash = get_hash(get_file_content(get_filename(last_block)))
    }

    x, _ := json.Marshal(Block{args[0], args[1], args[2], hash});
    file.WriteString(string(x))

    zero_block(get_hash(get_file_content(new_block)))
}

func zero_block (hash string) {
    var zero_block_name string = get_filename(0)

    if file_is_not_exist(zero_block_name) {
        create_file(zero_block_name)
    }

    file, err := os.OpenFile(zero_block_name, os.O_WRONLY|os.O_APPEND, 0777)
    check_error(err)
    defer file.Close()

    file.WriteString(hash + "\n")
}

func file_is_not_exist (filename string) bool {
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        return true
    }
    return false
}

func create_file (filename string) {
    file, err := os.Create(filename)
    check_error(err)
    file.Close()
}

func get_hash (message string) string {
    hash := sha256.New() 
    hash.Write([]byte(message))
    return hex.EncodeToString(hash.Sum(nil))
}

func get_filename (base int) string {
    return BLOCKCHAIN_DIR + strconv.Itoa(base) + EXTENSION
}

func get_file_content (filename string) string {
    file, err := os.Open(filename)
    check_error(err)
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

func get_last_block (dir string) int {
    files, err := ioutil.ReadDir(dir)
    check_error(err)

    var max int = 0
    for _, file := range files {
        max_block, _ := strconv.Atoi(get_filebase(file.Name()))
        if max_block > max {
            max = max_block
        }   
    }
    return max
}

func get_filebase (file string) string {
    return strings.Split(file, ".")[0]
}

func check_error (err error) {
    if err != nil {
        get_error(err.Error())
    }
}

func get_error (err string) {
    fmt.Println("Error:", err)
    os.Exit(1)
}
