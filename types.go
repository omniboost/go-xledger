package xledger

type QLQueryPaginated[T any] struct {
	PageInfo struct {
		HasNextPage bool
	}
	Edges []struct {
		Cursor string
		Node   T
	}
}

type QLQuery[T any] struct {
	Edges []struct {
		Node T
	}
}

type Account struct {
	DBId                  int      `graphql:"dbId"`
	Code                  string   `graphql:"code"`
	Description           string   `graphql:"description"`
	OwnerDBId             int      `graphql:"ownerDbId"`
	ModifiedAt            DateTime `graphql:"modifiedAt"`
	BalanceSheet          bool     `graphql:"balanceSheet"`
	Asset                 bool     `graphql:"asset"`
	Cost                  bool     `graphql:"cost"`
	Credit                bool     `graphql:"credit"`
	Equity                bool     `graphql:"equity"`
	Finance               bool     `graphql:"finance"`
	Liabilities           bool     `graphql:"liabilities"`
	Posting               bool     `graphql:"posting"`
	Revenue               bool     `graphql:"revenue"`
	ChartOfAccountDBId    int      `graphql:"chartOfAccountDbId"`
	AccountGroupDBId      int      `graphql:"accountGroupDbId"`
	Group                 bool     `graphql:"group"`
	DescriptionTranslated string   `graphql:"descriptionTranslated"`
	YearEnd               bool     `graphql:"yearEnd"`
	Vat                   bool     `graphql:"vat"`
	CostOfGoodsSold       bool     `graphql:"costOfGoodsSold"`
	CostWagesSalaries     bool     `graphql:"costWagesSalaries"`
	CostDepreciation      bool     `graphql:"costDepreciation"`
	CostSalesAndMarketing bool     `graphql:"costSalesAndMarketing"`
	CostOtherOperating    bool     `graphql:"costOtherOperating"`
	//SysAccount            interface{} `graphql:"sysAccount"`
	//AccountGroup          interface{} `graphql:"accountGroup"`
	//Owner                 interface{} `graphql:"owner"`
	//AccountIndicator      interface{} `graphql:"accountIndicator"`
	//ChartOfAccount        interface{} `graphql:"chartOfAccount"`
	//LedgerType            interface{} `graphql:"ledgerType"`
}

type Customer struct {
	IsIntraGroup          bool     `graphql:"isIntraGroup"`
	Email                 string   `graphql:"email"`
	DBId                  int      `graphql:"dbId"`
	ExtIdentifier         string   `graphql:"extIdentifier"`
	CreatedAt             DateTime `graphql:"createdAt"`
	ModifiedAt            DateTime `graphql:"modifiedAt"`
	Phone                 string   `graphql:"phone"`
	InvoiceNumber         string   `graphql:"invoiceNumber"`
	YourReference         string   `graphql:"yourReference"`
	FromDate              string   `graphql:"fromDate"`
	ToDate                string   `graphql:"toDate"`
	Code                  string   `graphql:"code"`
	Description           string   `graphql:"description"`
	OwnerDBId             int      `graphql:"ownerDbId"`
	CodeTranslated        string   `graphql:"codeTranslated"`
	DescriptionTranslated string   `graphql:"descriptionTranslated"`
	ID                    string   `graphql:"id"`
	Number                int      `graphql:"number"`
	PaymentNotification   bool     `graphql:"paymentNotification"`
	CreditLimit           string   `graphql:"creditLimit"`
	Contract              string   `graphql:"contract"`
	Notes                 string   `graphql:"notes"`
	Phone2                string   `graphql:"phone2"`
	TaxNo                 string   `graphql:"taxNo"`
	PayDocText            string   `graphql:"payDocText"`
	BankAccount           string   `graphql:"bankAccount"`
	CompanyDBId           int      `graphql:"companyDbId"`
	SubledgerGroupDBId    int      `graphql:"subledgerGroupDbId"`
	//HasFlexiField         bool     `graphql:"hasFlexiField"`
	//FlexiFieldFlag1       bool     `graphql:"flexiFieldFlag1"`
	//FlexiFieldFlag2       bool     `graphql:"flexiFieldFlag2"`
	//FlexiFieldFlag3       bool     `graphql:"flexiFieldFlag3"`
	//FlexiFieldFlag4       bool     `graphql:"flexiFieldFlag4"`
	//FlexiFieldDate1       Date     `graphql:"flexiFieldDate1"`
	//FlexiFieldDate2       Date     `graphql:"flexiFieldDate2"`
	//FlexiFieldDate3       Date     `graphql:"flexiFieldDate3"`
	//FlexiFieldDate4       Date     `graphql:"flexiFieldDate4"`
	//FlexiFieldValue1      float64  `graphql:"flexiFieldValue1"`
	//FlexiFieldValue2      float64  `graphql:"flexiFieldValue2"`
	//FlexiFieldValue3      float64  `graphql:"flexiFieldValue3"`
	//FlexiFieldValue4      float64  `graphql:"flexiFieldValue4"`
	//FlexiFieldNumber1     int      `graphql:"flexiFieldNumber1"`
	//FlexiFieldNumber2     int      `graphql:"flexiFieldNumber2"`
	//FlexiFieldText1       string   `graphql:"flexiFieldText1"`
	//FlexiFieldText2       string   `graphql:"flexiFieldText2"`
	//FlexiFieldText3       string   `graphql:"flexiFieldText3"`
	//FlexiFieldText4       string   `graphql:"flexiFieldText4"`
	//FlexiFieldText5       string   `graphql:"flexiFieldText5"`
	//FlexiFieldText6       string   `graphql:"flexiFieldText6"`
	//FlexiFieldText7       string   `graphql:"flexiFieldText7"`
	//FlexiFieldText8       string   `graphql:"flexiFieldText8"`
	//FlexiFieldCode1DBId   int      `graphql:"flexiFieldCode1DbId"`
	//FlexiFieldCode2DBId   int      `graphql:"flexiFieldCode2DbId"`
	//FlexiFieldCode3DBId   int      `graphql:"flexiFieldCode3DbId"`
	//FlexiFieldCode4DBId   int      `graphql:"flexiFieldCode4DbId"`
	//FlexiFieldCode5DBId   int      `graphql:"flexiFieldCode5DbId"`
	//FlexiFieldCode6DBId   int      `graphql:"flexiFieldCode6DbId"`
	//FlexiFieldCode7DBId   int      `graphql:"flexiFieldCode7DbId"`
	//FlexiFieldCode8DBId   int      `graphql:"flexiFieldCode8DbId"`
	//PriceList              interface{} `graphql:"priceList"`
	//PriceGroup             interface{} `graphql:"priceGroup"`
	//ReportSetup            interface{} `graphql:"reportSetup"`
	//Changes                interface{} `graphql:"_changes"`
	//PayDocCode             interface{} `graphql:"payDocCode"`
	//SubledgerGroup         interface{} `graphql:"subledgerGroup"`
	//FlexiFieldsItem        interface{} `graphql:"flexiFieldsItem"`
	//PhoneType2             interface{} `graphql:"phoneType2"`
	//PhoneType              interface{} `graphql:"phoneType"`
	//Language               interface{} `graphql:"language"`
	//InvoiceLayout          interface{} `graphql:"invoiceLayout"`
	//InvoiceDeliveryMethod  interface{} `graphql:"invoiceDeliveryMethod"`
	//Currency               interface{} `graphql:"currency"`
	//Xgl                    interface{} `graphql:"xgl"`
	//ledgerType             interface{} `graphql:"ledgerType"`
	//PaymentCharge          interface{} `graphql:"paymentCharge"`
	//CultureInfo            interface{} `graphql:"cultureInfo"`
	//ApTaxReporting         interface{} `graphql:"apTaxReporting"`
	//Contact                interface{} `graphql:"contact"`
	//OurRef                 interface{} `graphql:"ourRef"`
	//OurSales               interface{} `graphql:"ourSales"`
	//ShipAddress            interface{} `graphql:"shipAddress"`
	//BillAddress            interface{} `graphql:"billAddress"`
	//GlObject3              interface{} `graphql:"glObject3"`
	//GlObject2              interface{} `graphql:"glObject2"`
	//GlObject1              interface{} `graphql:"glObject1"`
	//GlObject4              interface{} `graphql:"glObject4"`
	//GlObject5              interface{} `graphql:"glObject5"`
	Address Address `graphql:"address"`
	//Company                interface{} `graphql:"company"`
	//Owner                  interface{} `graphql:"owner"`
	//Collection             interface{} `graphql:"collection"`
	//PayTerms               interface{} `graphql:"payTerms"`
	//PayMethod              interface{} `graphql:"payMethod"`
	//TaxRule                interface{} `graphql:"taxRule"`
	//AccountsPayableAccount interface{} `graphql:"accountsPayableAccount"`
	//Bank                   interface{} `graphql:"bank"`
	//BankRecord             interface{} `graphql:"bankRecord"`
}

type Address struct {
	Country       Country `graphql:"country"`
	StreetAddress string  `graphql:"streetAddress"`
	ZipCode       string  `graphql:"zipCode"`
	Place         string  `graphql:"place"`
}

type Country struct {
	DBId        int    `graphql:"dbId"`
	Code        string `graphql:"code"`
	Description string `graphql:"description"`
}

type Currency struct {
	DBId        int    `graphql:"dbId"`
	Description string `graphql:"description"`
	Code        string `graphql:"code"`
}

type SubLedger struct {
	DBId        int    `graphql:"dbId"`
	Description string `graphql:"description"`
}

type TaxCode struct {
	DBId      int    `graphql:"dbId"`
	OwnerDbId int    `graphql:"ownerDbId"`
	Text      string `graphql:"text"`
	Code      string `graphql:"code"`
}

type Company struct {
	DBId          int      `graphql:"dbId"`
	Description   string   `graphql:"description"`
	Code          string   `graphql:"code"`
	Country       string   `graphql:"country"`
	Address       Address  `graphql:"address"`
	TaxNumber     string   `graphql:"taxNumber"`
	CompanyName   string   `graphql:"companyName"`
	CompanyNumber string   `graphql:"companyNumber"`
	Email         string   `graphql:"email"`
	ModifiedAt    DateTime `graphql:"modifiedAt"`
	Name          string   `graphql:"name"`
	Notes         string   `graphql:"notes"`
	Phone         string   `graphql:"phone"`
	Phone2        string   `graphql:"phone2"`
	Phone3        string   `graphql:"phone3"`
	TaxNo         string   `graphql:"taxNo"`
}

type ObjectKind struct {
	DBId string `graphql:"dbId"`
	Name string `graphql:"name"`
}

type ObjectValue struct {
	DBId        int        `graphql:"dbId"`
	ObjectKind  ObjectKind `graphql:"objectKind"`
	Code        string     `graphql:"code"`
	Description string     `graphql:"description"`
}

type CompaniesInput struct {
	Description   string                `json:"description,omitempty"`
	Code          string                `json:"code,omitempty"`
	Country       CompaniesInputCountry `json:"country,omitempty"`
	Address       CompaniesInputAddress `json:"address"`
	TaxNumber     string                `json:"taxNumber,omitempty"`
	CompanyNumber string                `json:"companyNumber,omitempty"`
	Email         string                `json:"email,omitempty"`
	Phone         string                `json:"phone,omitempty"`
	Phone2        string                `json:"phone2,omitempty"`
}

type CompaniesInputCountry struct {
	Code string `json:"code"`
}
type CompaniesInputAddress struct {
	Country       string `json:"country"`
	StreetAddress string `json:"streetAddress,omitempty"`
	ZipCode       string `json:"zipCode,omitempty"`
	Place         string `json:"place,omitempty"`
}

func (CompaniesInput) GetGraphQLType() string {
	return "AddCompaniesInputNode"
}
