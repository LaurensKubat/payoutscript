package jsonrpc

type Blocks Service

func (b Blocks) info(blockID string) (interface{}, error) {
	r, err := b.client.prepBody("blocks.info", blockID)
	if err != nil {
		return nil, err
	}
	res, err := b.client.send(contenttype, r)
	return res, err
}

func (b Blocks) latest() (interface{}, error) {
	r, err := b.client.prepBody("blocks.latest")
	if err != nil {
		return nil, err
	}
	res, err := b.client.send(contenttype, r)
	return res, err
}

func (b Blocks) transactions(blockID string, offset string) (interface{}, error) {
	r, err := b.client.prepBody("blocks.transactions", blockID, offset)
	if err != nil {
		return nil, err
	}
	res, err := b.client.send(contenttype, r)
	return res, err
}
