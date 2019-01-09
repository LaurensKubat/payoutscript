package jsonrpc

type transactions Service

type transactionID Service

func (t transactions) Create(amount string, recipientID string, passphrase string) (interface{}, error) {
	r, err := t.client.prepBody("transactions.create", amount, recipientID, passphrase)
	if err != nil {
		return nil, err
	}
	res, err := t.client.send(contenttype, r)
	return res, err
}

func (t transactions) broadcast(id string) (interface{}, error) {
	r, err := t.client.prepBody("transactions.broadcast", id)
	if err != nil {
		return nil, err
	}
	res, err := t.client.send(contenttype, r)
	return res, err
}

func (t transactions) info(id string) (interface{}, error) {
	r, err := t.client.prepBody("transactions.info", id)
	if err != nil {
		return nil, err
	}
	res, err := t.client.send(contenttype, r)
	return res, err
}

func (t transactions) bip38create(recipientID string, amount string, bip38 string, userID string) {

}
