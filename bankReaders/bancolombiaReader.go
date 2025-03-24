package bankreaders

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"
)

type BancolombiaFileReader struct {}

func (bancolombiaFileReader BancolombiaFileReader) ReadFile(filename string) ([]BankData, error){
    data, err := ReadFile(filename)
    if err != nil {
        return nil, err
    }
    reader, err := parseCSV(data, '\t')
    if err != nil{
        return nil, err
    }
    bankData, err := bancolombiaProcessTSV(reader)
    if err != nil{
        return nil, err
    }
    return bankData, nil
}

func bancolombiaProcessTSV(reader *csv.Reader) ([]BankData, error) {
    bankData := make([]BankData, 0)
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err!= nil {
            return nil, err
        }
        date, err := time.Parse("2006/01/02", record[0])
        if err != nil{
            return nil, err
        }
        cleaned_amount := strings.ReplaceAll(record[5], ",", "")
        amount, err := strconv.ParseFloat(cleaned_amount, 64)
        if err != nil{
            return nil, err
        }
        currentBankData := BankData{
            date: date,
            description :record[3],
            amount: int64(amount),
        }
        bankData = append(bankData, currentBankData)
    }
    return bankData, nil
}
