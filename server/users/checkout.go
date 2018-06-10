package users

import (
	"errors"
	"net/http"
)

type Check map[string]interface{}

func (user *User) Checkout(valor string) error {
	//API KEY - gmxCheckoutGtwy
	ch := Check{
		"restApi":                          "true",
		"recorrencia.modalidadeVenda":      "1",
		"empresa.hashEmpresa":              "gmxCheckoutGtwy",
		"recorrencia.descricaoFatura":      "Doação",
		"recorrencia.idrecorrenciaEmpresa": "test",
		"recorrencia.consumidor.nome":      user.Name,
		"recorrencia.consumidor.email":     user.Email,
		"recorrencia.valor":                valor,
		"recorrencia.parcelas":             "1",
		"cartaoCredito.portador":           user.Name,
		"recorrencia.consumidor.cpf":       user.Cpf,
		"cartaoCredito.numero":             "4012001037141114",
		"cartaoCredito.mesValidade":        "01",
		"cartaoCredito.anoValidade":        "2020",
		"cartaoCredito.codSeguranca":       "123",
		"cartaoCredito.bandeira":           "visa",
		"recorrencia.produto":              "Doação",
	}
	url := builderCheckoutUrl(ch)
	req, err := http.Get(url)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	if req.StatusCode != http.StatusOK {
		return errors.New("falha ao efetuar checkout")
	}
	return nil
}

func builderCheckoutUrl(data Check) string {
	req, _ := http.NewRequest("GET", "https://www.gmxcheckout.com.br/txn/post", nil)
	q := req.URL.Query()
	for k, v := range data {
		q.Add(k, v.(string))
	}
	req.URL.RawQuery = q.Encode()
	return req.URL.String()
}
