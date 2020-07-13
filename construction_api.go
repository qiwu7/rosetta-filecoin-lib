/*******************************************************************************
*   (c) 2020 Zondax GmbH
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
********************************************************************************/
package rosetta_filecoin_lib

type ConstructionClient interface {
	// DeriveFromPublicKey defines the function to derive the address from an public key (secp256k1)
	// @return
	//   - derivedAddress [string]
	//   - error when deriving address from the public key
	DeriveFromPublicKey(publicKey []byte) (string, error)

	// Sign defines the function to sign an arbitrary message with the private key (secp256k1)
	// @return (secp256k1)
	//   - signature [string] the signature after the message is signed with the private key
	//   - error when signing a message
	Sign(message []byte, privateKey []byte) ([]byte, error)

	// Verify defines the function to verify the signature of an arbitrary message with the public key (secp256k1)
	// @return
	//   - error if invalid signature
	Verify(message []byte, publicKey []byte, signature []byte) error

	// ConstructPayment creates transaction for a normal send
	// @return
	//   - unsignedTx [string]
	//   - error while constructing the normal send transaction
	ConstructPayment(request *PaymentRequest) (string, error)

	// ConstructMultisigPayment creates transaction for a multisig send
	// @return
	//   - unsignedTx [string]
	//   - error while constructing the multisig send transaction
	ConstructMultisigPayment(request *MultisigPaymentRequest) (string, error)

	// ConstructSwapAuthorizedParty creates transaction for a multisig SwapAuthorizedParty call
	// @return
	//   - unsignedTx [string]
	//   - error while constructing the multisig SwapAuthorizedParty call
	ConstructSwapAuthorizedParty(request *MultisigPaymentRequest) (string, error)

	// SignTx signs an unsignedTx using the private key (secp256k1) and return a signedTx that can be submitted to the node
	// @return
	//   - signedTx [string] the signed transaction
	//   - error when signing a transaction
	SignTx(unsignedTransaction string, privateKey []byte) (string, error)

	// ParseTx defines the function to parse a transaction
	// @return
	//   - message [string] the parsed transaction (message), this will either be a Message or a SignedMessage
	//   - error when parsing a transaction
	ParseTx(request *ParseTxRequest) (interface{}, error)

	// Hash defines the function to calculate a tx hash
	// @return
	//   - txHash [string] transaction hash
	//   - error when calculating the tx hash
	Hash(signedTx string) (string, error)
}

// Modify this as needed to add in new fields
type TxMetadata struct {
	Nonce               uint64 `json:"nonce"`
	GasPrice            string `json:"gasPrice,omitempty"`
	GasLimit            string `json:"gasLimit,omitempty"`
	ChainId             string `json:"chainId"`
	Method              uint64 `json:"method,omitempty"`
	Params              []byte `json:"params,omitempty"`
}

// PaymentRequest defines the input to ConstructPayment
type PaymentRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity uint64 `json:"quantity"`
	Metadata TxMetadata `json:"metadata"`
}


// MultisigPaymentRequest defines the input to ConstructMultisigPayment
type MultisigPaymentRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity uint64 `json:"quantity"`
	Metadata TxMetadata `json:"metadata"`
}

// SwapAuthorizedPartyRequest defines the input to ConstructSwapAuthorizedParty
type SwapAuthorizedPartyRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Metadata TxMetadata `json:"metadata"`
}

// ParseTxRequest defines the input to ParseTx
type ParseTxRequest struct {
	UnsignedTransaction string `json:"unsigned_tx"`
	SignedTransaction   string `json:"signed_tx"`
}

// PaymentRequest defines the input to ConstructPayment
type ParseTxResponse struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity uint64 `json:"quantity,omitempty"`
	Method   uint64 `json:"method,omitempty"`
}
