package models

// The Model
// The omitempty options means that should the field be falsey, then omit it from the response
type Token struct {
    Name          string
    Ticker        string
    Description   string
}