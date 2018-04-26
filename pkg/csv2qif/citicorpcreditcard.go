package CitiCorpCreditCard

import qif "github.com/mdhender/qif/qif"

// Name CSV columns to make the code more readable
const (
	// AccountNumber is the redacted CCard number
	accountNumber = iota
	// AccountName is the CCard holder name
	accountName
	// TransactionDate is the date the transaction took place
	transactionDate
	// PostDate is the date the transaction cleared
	postDate
	// ReferenceNumber is a transaction ID
	referenceNumber
	// TransactionDetail usually contains vendor name and location
	transactionDetail
	// BillingAmount is the transaction amount in the CCard local currency
	billingAmount
	// SourceCurrency is a three letter currency ID
	sourceCurrency
	// SourceAmount is the transaction amount in SourceCurrency
	sourceAmount
	// CustomerRef so far has been an empty field
	customerRef
	// EmployeeID is a 9 digit number. For US employees it's their SSN
	employeeID
)

func Csv2qif(r io.Reader)

qif.go()