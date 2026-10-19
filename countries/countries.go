package countries

type Country struct {
	ISO3        string
	ISO2        string
	Name        string
	IsISOMember bool
}

type Group struct {
	Code    string
	Name    string
	Members []Country
}
