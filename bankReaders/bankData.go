package bankreaders

import "time"

type BankData struct {
    date time.Time
    description string
    amount int64
}
