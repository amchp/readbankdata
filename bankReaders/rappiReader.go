package bankreaders

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"
)

type RappiFileReader struct {}

func (bancolombiaFileReader RappiFileReader) ReadFile(filename string) ([]BankData, error) {
    data, err := ReadFile(filename)
    if err != nil {
        return nil, err
    }
    reader, err := parseCSV(data, ';')
    if err != nil{
        return nil, err
    }
    bankData, err := rappiProcessCSV(reader)
    if err != nil{
        return nil, err
    }
    return bankData, nil
}

func rappiProcessCSV(reader *csv.Reader) ([]BankData, error) {
    bankData := make([]BankData, 0)
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err!= nil {
            return nil, err
        }
        date, err := time.Parse("2006-01-02", record[1])
        if err != nil{
            return nil, err
        }
        cleaned_amount := strings.ReplaceAll(record[3][1:], ".", "")
        cleaned_amount = strings.ReplaceAll(cleaned_amount, ",", ".")
        amount, err := strconv.ParseFloat(cleaned_amount, 64)
        amount = -amount
        if err != nil{
            return nil, err
        }
        currentBankData := BankData{
            date: date,
            description: record[2],
            amount: int64(amount),
        }
        bankData = append(bankData, currentBankData)
    }
    return bankData, nil
}
