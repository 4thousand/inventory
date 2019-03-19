package NPINven

import "errors"

var (
	StatusNotFound           = errors.New("sql: no rows in result set")
	ErrMethodNotAllowed      = errors.New("Customer: Data not found")
	ErrForbidden             = errors.New("Customer: Data not found")
	ArCodeNull               = errors.New("Arcode is null")
	NotHaveItem              = errors.New("Docno is not have item")
	NotHavePayMoney          = errors.New("Docno not set money to another type payment")
	NotHaveSumOfItem         = errors.New("SumofItemAmount = 0")
	ItemNotHaveQty           = errors.New("Item not have qty")
	ItemNotHaveUnit          = errors.New("Item not have unitcode")
	MoneyOverTotalAmount     = errors.New("Rec money is over totalamount")
	MoneyLessThanTotalAmount = errors.New("Rec money is less than totalamount")
	PosNotHaveDate           = errors.New("Docno not have pos data")
	PosNotHaveChqData        = errors.New("Docno not have chq data")
	PosNotHaveCreditCardData = errors.New("Docno not have credit card data")
)
