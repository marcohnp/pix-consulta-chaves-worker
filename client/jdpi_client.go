package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pix-consulta-chaves-worker/model"
)

func ConsultarChavePix(chave string) (*model.Chave, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", chave))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro HTTP: %s", resp.Status)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &model.Chave{
		Chave:                     chave,
		Nome:                      fmt.Sprintf("%v", data["name"]),
		NomeFantasia:              fmt.Sprintf("%v", data["username"]),
		CpfCnpj:                   "12345678900",
		TpChave:                   1,
		TpConta:                   1,
		Ispb:                      "12345678",
		NrAgencia:                 "0001",
		NrConta:                   "123456-7",
		DtHrAberturaConta:         "2020-01-01T00:00:00Z",
		TpPessoa:                  1,
		DtHrCriacaoChave:          "2020-01-02T00:00:00Z",
		DtHrIncioPosseChave:       "2020-01-03T00:00:00Z",
		DtHrAberturaReivindicacao: "",
		EndToEndId:                "E2E123456789",
	}, nil
}
