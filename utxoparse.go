package main

// local packages
import "bitcoin/btcleveldb" // chainstate leveldb decoding functions
import "bitcoin/keys"   // bitcoin addresses
import "bitcoin/bech32" // segwit bitcoin addresses

import "github.com/syndtr/goleveldb/leveldb" // go get github.com/syndtr/goleveldb/leveldb
import "github.com/syndtr/goleveldb/leveldb/opt" // set no compression when opening leveldb
import "flag"         // command line arguments
import "fmt"
import "os"           // open file for writing
import "os/exec"      // execute shell command (check bitcoin isn't running)
import "os/signal"    // catch interrupt signals CTRL-C to close db connection safely
import "syscall"      // catch kill commands too
import "bufio"        // bulk writing to file
import "encoding/hex" // convert byte slice to hexadecimal
import "strings"      // parsing flags from command line
import "runtime"      // Check OS type for file-handler limitations

func main() {


        int equality = 2
        long value = 1000000000
        args := os.Args
        if(args[0] == "greater"){equality = 0}
        else if(args[0] == "lesser"){equality = 1}
        else if(args[0]) {equality = 2}

        value = Int, err := strconv.Atoi(args[1])

        defaultfolder := fmt.Sprintf("%s/.bitcoin/chainstate/", os.Getenv("HOME")) // %s = string
        defaultfile := "utxodump.csv"
        file := flag.String("o", defaultfile, "Name of file to dump utxo list to.") // output file
        f,err := os.open(file)
        r := f.NewReader(file)
        for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
                switch equality {
                case 0:
                        if(record[3] > value){
                                fmt.Printf(record[4])
                         }
                case 1:
                        if(record[3] < value){
                                fmt.Printf(record[4])
                         }             
                default:
                        if(record[3] == value){
                                fmt.Printf(record[4])
                         }
                }
                 
	}

}
