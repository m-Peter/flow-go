package templates_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/sdk/emulator/constants"
	"github.com/dapperlabs/flow-go/sdk/templates"
)

func TestCreateAccount(t *testing.T) {
	publicKey := []byte{4, 136, 178, 30, 0, 0, 0, 0, 0, 0, 0, 0, 0, 111, 117, 56, 107, 245, 122, 184, 40, 127, 172, 19, 175, 225, 131, 184, 22, 122, 23, 90, 172, 214, 144, 150, 92, 69, 119, 218, 11, 191, 120, 226, 74, 2, 217, 156, 75, 44, 44, 121, 152, 143, 47, 180, 169, 205, 18, 77, 47, 135, 146, 34, 34, 157, 69, 149, 177, 141, 80, 99, 66, 186, 33, 25, 73, 179, 224, 166, 205, 172}

	accountKey := flow.AccountKey{
		PublicKey: publicKey,
		Weight:    constants.AccountKeyWeightThreshold,
	}

	// create account with no code
	scriptA := templates.CreateAccount([]flow.AccountKey{accountKey}, []byte{})

	expectedScriptA := []byte(`
		fun main() {
			let publicKeys: [[Int]] = [[4,136,178,30,0,0,0,0,0,0,0,0,0,111,117,56,107,245,122,184,40,127,172,19,175,225,131,184,22,122,23,90,172,214,144,150,92,69,119,218,11,191,120,226,74,2,217,156,75,44,44,121,152,143,47,180,169,205,18,77,47,135,146,34,34,157,69,149,177,141,80,99,66,186,33,25,73,179,224,166,205,172]]
			let keyWeights: [Int] = [1000]
			let code: [Int]? = []
			createAccount(publicKeys, keyWeights, code)
		}
	`)

	assert.Equal(t, expectedScriptA, scriptA)

	// create account with code
	scriptB := templates.CreateAccount([]flow.AccountKey{accountKey}, []byte("fun main() {}"))

	expectedScriptB := []byte(`
		fun main() {
			let publicKeys: [[Int]] = [[4,136,178,30,0,0,0,0,0,0,0,0,0,111,117,56,107,245,122,184,40,127,172,19,175,225,131,184,22,122,23,90,172,214,144,150,92,69,119,218,11,191,120,226,74,2,217,156,75,44,44,121,152,143,47,180,169,205,18,77,47,135,146,34,34,157,69,149,177,141,80,99,66,186,33,25,73,179,224,166,205,172]]
			let keyWeights: [Int] = [1000]
			let code: [Int]? = [102,117,110,32,109,97,105,110,40,41,32,123,125]
			createAccount(publicKeys, keyWeights, code)
		}
	`)

	assert.Equal(t, expectedScriptB, scriptB)
}

func TestUpdateAccountCode(t *testing.T) {
	script := templates.UpdateAccountCode([]byte("fun main() {}"))

	expectedScript := []byte(`
		fun main(account: Account) {
			let code = [102,117,110,32,109,97,105,110,40,41,32,123,125]
			updateAccountCode(account.address, code)
		}
	`)

	assert.Equal(t, expectedScript, script)
}
