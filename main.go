package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/amchp/readbankdata/bankreaders"
)


type BankReader interface {
    ReadFile(filename string) ([]bankreaders.BankData, error)
}

type BankReaderService struct {
	bankReader BankReader
}

type BankReaderEnum int

const (
    Bancolombia BankReaderEnum = iota
    Rappi
)

func main() {
    if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <filename> <reader>")
		os.Exit(1)
	}

	argsMap := make(map[string]string)
	argsMap["filename"] = os.Args[1]
	argsMap["reader"] = os.Args[2]

    bankReaderType, err := ParseBankReader(argsMap["reader"])
    if err != nil{
        log.Fatal(err)
        return
    }
    data, err := ReadBankData(argsMap["filename"], bankReaderType)
    if err != nil{
        log.Fatal(err)
        return
    }
    for _, dt := range data{
        fmt.Println(dt)
    }
}

func ReadBankData(filename string, bankReaderType BankReaderEnum) ([]bankreaders.BankData, error){
    bankReaderService, err := GetBankReaderService(bankReaderType)
    if err != nil{
        return nil, err
    }
    data, err := bankReaderService.bankReader.ReadFile(filename)
    if err != nil{
        return nil, err
    }
    return data, nil
}

var bankReaderMap = map[string]BankReaderEnum{
    "bancolombia":    Bancolombia,
    "rappi":    Rappi,
}

func ParseBankReader(bankReaderString string) (BankReaderEnum, error) {
    bankReaderString = strings.ToLower(bankReaderString)
    if day, ok := bankReaderMap[bankReaderString]; ok {
        return day, nil
    }
    return 0, fmt.Errorf("invalid bank reader: %s", bankReaderString)
}

func GetBankReaderService(bankReaderType BankReaderEnum) (BankReaderService, error){
    switch bankReaderType{
        case Bancolombia:
            return BankReaderService{
                bankReader: bankreaders.BancolombiaFileReader{},
            }, nil
        case Rappi:
            return BankReaderService{
                bankReader: bankreaders.RappiFileReader{},
            }, nil
        default:
            return BankReaderService{}, fmt.Errorf("invalid bank reader: %s", bankReaderType) 
    }
}
