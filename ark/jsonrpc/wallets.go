package jsonrpc

type wallets Service

func (w wallets) create(passphrase string) (interface{}, error) {
	r, err := w.client.prepBody("wallets.create", passphrase)
	if err != nil {
		return nil, err
	}
	res, err := w.client.send(contenttype, r)
	return res, err
}

func (w wallets) info(address string) (interface{}, error) {
	r, err := w.client.prepBody("wallets.info", address)
	if err != nil {
		return nil, err
	}
	res, err := w.client.send(contenttype, r)
	return res, err
}

func (w wallets) transactions(address string, offset string) (interface{}, error) {
	r, err := w.client.prepBody("wallets.transactions", address, offset)
	if err != nil {
		return nil, err
	}
	res, err := w.client.send(contenttype, r)
	return res, err
}
