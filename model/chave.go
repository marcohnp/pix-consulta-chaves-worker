package model

type Chave struct {
	Chave                     string `json:"chave"`
	TpChave                   int    `json:"tpChave"`
	Ispb                      string `json:"ispb"`
	NrAgencia                 string `json:"nrAgencia"`
	TpConta                   int    `json:"tpConta"`
	NrConta                   string `json:"nrConta"`
	DtHrAberturaConta         string `json:"dtHrAberturaConta"`
	TpPessoa                  int    `json:"tpPessoa"`
	CpfCnpj                   string `json:"cpfCnpj"`
	Nome                      string `json:"nome"`
	NomeFantasia              string `json:"nomeFantasia"`
	DtHrCriacaoChave          string `json:"dtHrCriacaoChave"`
	DtHrIncioPosseChave       string `json:"dtHrIncioPosseChave"`
	DtHrAberturaReivindicacao string `json:"dtHrAberturaReivindicacao"`
	EndToEndId                string `json:"endToEndId"`
}
