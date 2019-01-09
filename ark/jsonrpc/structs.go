package jsonrpc

import "time"

type ErrorMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Error   struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

type BlocksInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		ID       string `json:"id"`
		Version  int    `json:"version"`
		Height   int    `json:"height"`
		Previous string `json:"previous"`
		Forged   struct {
			Reward int `json:"reward"`
			Fee    int `json:"fee"`
			Total  int `json:"total"`
		} `json:"forged"`
		Payload struct {
			Hash   string `json:"hash"`
			Length int    `json:"length"`
		} `json:"payload"`
		Generator struct {
			Username  string `json:"username"`
			Address   string `json:"address"`
			PublicKey string `json:"publicKey"`
		} `json:"generator"`
		Signature    string `json:"signature"`
		Transactions int    `json:"transactions"`
		Timestamp    struct {
			Epoch int       `json:"epoch"`
			Unix  int       `json:"unix"`
			Human time.Time `json:"human"`
		} `json:"timestamp"`
	} `json:"result"`
}

type BlocksLatest struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		ID       string `json:"id"`
		Version  int    `json:"version"`
		Height   int    `json:"height"`
		Previous string `json:"previous"`
		Forged   struct {
			Reward int `json:"reward"`
			Fee    int `json:"fee"`
			Total  int `json:"total"`
		} `json:"forged"`
		Payload struct {
			Hash   string `json:"hash"`
			Length int    `json:"length"`
		} `json:"payload"`
		Generator struct {
			Username  string `json:"username"`
			Address   string `json:"address"`
			PublicKey string `json:"publicKey"`
		} `json:"generator"`
		Signature    string `json:"signature"`
		Transactions int    `json:"transactions"`
		Timestamp    struct {
			Epoch int       `json:"epoch"`
			Unix  int       `json:"unix"`
			Human time.Time `json:"human"`
		} `json:"timestamp"`
	} `json:"result"`
}

type BlocksTransactions struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Count int `json:"count"`
		Data  []struct {
			ID            string `json:"id"`
			BlockID       string `json:"blockId"`
			Type          int    `json:"type"`
			Amount        int64  `json:"amount"`
			Fee           int    `json:"fee"`
			Sender        string `json:"sender"`
			Recipient     string `json:"recipient"`
			Signature     string `json:"signature"`
			Confirmations int    `json:"confirmations"`
			Timestamp     struct {
				Epoch int       `json:"epoch"`
				Unix  int       `json:"unix"`
				Human time.Time `json:"human"`
			} `json:"timestamp"`
		} `json:"data"`
	} `json:"result"`
}

type TransactionsInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		ID            string `json:"id"`
		BlockID       string `json:"blockId"`
		Type          int    `json:"type"`
		Amount        int    `json:"amount"`
		Fee           int    `json:"fee"`
		Sender        string `json:"sender"`
		Recipient     string `json:"recipient"`
		Signature     string `json:"signature"`
		Confirmations int    `json:"confirmations"`
		Timestamp     struct {
			Epoch int       `json:"epoch"`
			Unix  int       `json:"unix"`
			Human time.Time `json:"human"`
		} `json:"timestamp"`
	} `json:"result"`
}

type TransactionsBroadcast struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		ID              string `json:"id"`
		Signature       string `json:"signature"`
		Timestamp       int    `json:"timestamp"`
		Type            int    `json:"type"`
		Fee             int    `json:"fee"`
		SenderPublicKey string `json:"senderPublicKey"`
		Amount          int    `json:"amount"`
		RecipientID     string `json:"recipientId"`
	} `json:"result"`
}

type TransactionsCreate struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		ID              string `json:"id"`
		Signature       string `json:"signature"`
		Timestamp       int    `json:"timestamp"`
		Type            int    `json:"type"`
		Fee             int    `json:"fee"`
		SenderPublicKey string `json:"senderPublicKey"`
		Amount          int    `json:"amount"`
		RecipientID     string `json:"recipientId"`
	} `json:"result"`
}

type TransactionsBIP38Create struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		ID              string `json:"id"`
		Signature       string `json:"signature"`
		Timestamp       int    `json:"timestamp"`
		Type            int    `json:"type"`
		Fee             int    `json:"fee"`
		SenderPublicKey string `json:"senderPublicKey"`
		Amount          int    `json:"amount"`
		RecipientID     string `json:"recipientId"`
	} `json:"result"`
}

type WalletsInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Address         string      `json:"address"`
		PublicKey       string      `json:"publicKey"`
		SecondPublicKey interface{} `json:"secondPublicKey"`
		Balance         int64       `json:"balance"`
		IsDelegate      bool        `json:"isDelegate"`
	} `json:"result"`
}

type WalletsTransactions struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Count int `json:"count"`
		Data  []struct {
			ID            string `json:"id"`
			BlockID       string `json:"blockId"`
			Type          int    `json:"type"`
			Amount        int64  `json:"amount"`
			Fee           int    `json:"fee"`
			Sender        string `json:"sender"`
			Recipient     string `json:"recipient"`
			Signature     string `json:"signature"`
			Confirmations int    `json:"confirmations"`
			Timestamp     struct {
				Epoch int       `json:"epoch"`
				Unix  int       `json:"unix"`
				Human time.Time `json:"human"`
			} `json:"timestamp"`
		} `json:"data"`
	} `json:"result"`
}

type WalletsCreate struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		PublicKey string `json:"publicKey"`
		Address   string `json:"address"`
	} `json:"result"`
}

type WalletsBIP38Info struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Wif string `json:"wif"`
	} `json:"result"`
}

type WalletsBIP38Create struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		PublicKey string `json:"publicKey"`
		Address   string `json:"address"`
		Wif       string `json:"wif"`
	} `json:"result"`
}
