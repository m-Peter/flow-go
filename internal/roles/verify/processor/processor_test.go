package processor

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/dapperlabs/bamboo-node/internal/pkg/types"
	. "github.com/dapperlabs/bamboo-node/internal/roles/verify/processor/test_mocks"
)

type test struct {
	title      string
	m          Mock
	expectFunc func(Mock, *testing.T)
}

func Test(t *testing.T) {

	tests := []test{
		test{
			title: "Happy Path",
			m:     &MockEffectsHappyPath{},
			expectFunc: func(m Mock, t *testing.T) {
				RegisterTestingT(t)
				Expect(m.CallCountIsValidExecutionReceipt()).To(Equal(1))
				Expect(m.CallCountHasMinStake()).To(Equal(1))
				Expect(m.CallCountIsSealedWithDifferentReceipt()).To(Equal(1))
				Expect(m.CallCountSend()).To(Equal(1))
				Expect(m.CallCountSlashExpiredReceipt()).To(Equal(0))
				Expect(m.CallCountSlashInvalidReceipt()).To(Equal(0))
				Expect(m.CallCountHandleError()).To(Equal(0))
			},
		},
		test{
			title: "No min stake should fail early",
			m:     NewMockEffectsNoMinStake(&MockEffectsHappyPath{}),
			expectFunc: func(m Mock, t *testing.T) {
				RegisterTestingT(t)
				Expect(m.CallCountIsValidExecutionReceipt()).To(Equal(0))
				Expect(m.CallCountHasMinStake()).To(Equal(1))
				Expect(m.CallCountIsSealedWithDifferentReceipt()).To(Equal(0))
				Expect(m.CallCountSend()).To(Equal(0))
				Expect(m.CallCountSlashExpiredReceipt()).To(Equal(0))
				Expect(m.CallCountSlashInvalidReceipt()).To(Equal(0))
				Expect(m.CallCountHandleError()).To(Equal(1))
			},
		},
		test{
			title: "sealed with different Receipt",
			m:     NewMockEffectsSealWithDifferentReceipt(&MockEffectsHappyPath{}),
			expectFunc: func(m Mock, t *testing.T) {
				RegisterTestingT(t)
				Expect(m.CallCountIsValidExecutionReceipt()).To(Equal(0))
				Expect(m.CallCountHasMinStake()).To(Equal(1))
				Expect(m.CallCountIsSealedWithDifferentReceipt()).To(Equal(1))
				Expect(m.CallCountSend()).To(Equal(0))
				Expect(m.CallCountSlashExpiredReceipt()).To(Equal(1))
				Expect(m.CallCountSlashInvalidReceipt()).To(Equal(0))
				Expect(m.CallCountHandleError()).To(Equal(0))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			RegisterTestingT(t)

			c := &ReceiptProcessorConfig{
				QueueBuffer: 100,
				CacheBuffer: 100,
			}
			p := NewReceiptProcessor(test.m, c)

			receipt := &types.ExecutionReceipt{}
			done := make(chan bool, 1)
			p.Submit(receipt, done)

			select {
			case _ = <-done:
				test.expectFunc(test.m, t)
			case <-time.After(1 * time.Second):
				t.Error("waited for receipt to be processed for me than 1 sec")
			}

		})
	}
}
