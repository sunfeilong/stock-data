package errors

type StockDataError struct {
    Msg string
}

func (s StockDataError) Error() string {
    return s.Msg
}

func (s StockDataError) String() string {
    return s.Msg
}
