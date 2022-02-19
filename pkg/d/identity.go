package d

type Identity struct {
	Account  *Account   `json:"account"`
	Profiles []*Profile `json:"profiles"`
}
