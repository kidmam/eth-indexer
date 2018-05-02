package transaction

import (
	"os"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/maichain/eth-indexer/common"
	"github.com/maichain/eth-indexer/model"
	"github.com/maichain/mapi/base/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func makeTx(blockNum int64, blockHex, txHex string) *model.Transaction {
	return &model.Transaction{
		Hash:        common.HexToBytes(txHex),
		BlockHash:   common.HexToBytes(blockHex),
		From:        common.HexToBytes("0xB287a379e6caCa6732E50b88D23c290aA990A892"),
		Nonce:       10013,
		GasPrice:    "123456789",
		GasLimit:    45000,
		Amount:      "4840283445",
		Payload:     []byte{12, 34},
		BlockNumber: blockNum,
	}
}

var _ = Describe("Transaction Database Test", func() {
	var (
		mysql *test.MySQLContainer
		db    *gorm.DB
	)
	BeforeSuite(func() {
		var err error
		mysql, err = test.NewMySQLContainer("quay.io/amis/eth-indexer-db-migration")
		Expect(mysql).ShouldNot(BeNil())
		Expect(err).Should(Succeed())
		Expect(mysql.Start()).Should(Succeed())

		db, err = gorm.Open("mysql", mysql.URL)
		Expect(err).Should(Succeed())
		Expect(db).ShouldNot(BeNil())

		db.LogMode(os.Getenv("ENABLE_DB_LOG_IN_TEST") != "")
	})

	AfterSuite(func() {
		mysql.Stop()
	})

	BeforeEach(func() {
		db.Table(TableName).Delete(&model.Transaction{})
	})

	It("should insert", func() {
		store := NewWithDB(db)
		blockHex := "0x99bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b"

		data1 := makeTx(32100, blockHex, "0x58bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")

		By("insert new transaction")
		err := store.Insert(data1)
		Expect(err).Should(Succeed())

		By("failed to insert again")
		err = store.Insert(data1)
		Expect(err).ShouldNot(BeNil())

		data2 := makeTx(32100, blockHex, "0x68bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		By("insert another new transaction")
		err = store.Insert(data2)
		Expect(err).Should(Succeed())
	})

	It("deletes transactions from a block number", func() {
		store := NewWithDB(db)
		blockHex1 := "0x88bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b"
		blockHex2 := "0x99bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b"
		blockHex3 := "0x77bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b"
		By("insert three new transactions")
		data1 := makeTx(32100, blockHex1, "0x58bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data2 := makeTx(42100, blockHex2, "0x68bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data3 := makeTx(52100, blockHex3, "0x78bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data := []*model.Transaction{data1, data2, data3}
		for _, tx := range data {
			err := store.Insert(tx)
			Expect(err).Should(Succeed())
		}

		err := store.DeleteFromBlock(42100)
		Expect(err).Should(Succeed())

		tx, err := store.FindTransaction(data1.Hash)
		Expect(err).Should(Succeed())
		Expect(reflect.DeepEqual(*tx, *data1)).Should(BeTrue())
		tx, err = store.FindTransaction(data2.Hash)
		Expect(common.NotFoundError(err)).Should(BeTrue())
		tx, err = store.FindTransaction(data3.Hash)
		Expect(common.NotFoundError(err)).Should(BeTrue())
	})

	It("should get transaction by hash", func() {
		store := NewWithDB(db)
		blockHex1 := "0x88bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b"
		blockHex2 := "0x99bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b"
		By("insert three new transactions")
		data1 := makeTx(32100, blockHex1, "0x58bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data2 := makeTx(32100, blockHex1, "0x68bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data3 := makeTx(42100, blockHex2, "0x78bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data := []*model.Transaction{data1, data2, data3}
		for _, tx := range data {
			err := store.Insert(tx)
			Expect(err).Should(Succeed())
		}

		transaction, err := store.FindTransaction(data1.Hash)
		Expect(err).Should(Succeed())
		Expect(reflect.DeepEqual(*transaction, *data1)).Should(BeTrue())

		transaction, err = store.FindTransaction(data2.Hash)
		Expect(err).Should(Succeed())
		Expect(reflect.DeepEqual(*transaction, *data2)).Should(BeTrue())

		transaction, err = store.FindTransaction(data3.Hash)
		Expect(err).Should(Succeed())
		Expect(reflect.DeepEqual(*transaction, *data3)).Should(BeTrue())

		By("find an non-existent transaction")
		transaction, err = store.FindTransaction(data2.BlockHash)
		Expect(common.NotFoundError(err)).Should(BeTrue())
		Expect(reflect.DeepEqual(*transaction, model.Transaction{})).Should(BeTrue())
	})

	It("should get transaction by block hash", func() {
		store := NewWithDB(db)
		blockHex1 := "0x88bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b"
		blockHex2 := "0x99bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b"
		By("insert three new transactions")
		data1 := makeTx(32100, blockHex1, "0x58bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data2 := makeTx(32100, blockHex1, "0x68bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data3 := makeTx(42100, blockHex2, "0x78bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data := []*model.Transaction{data1, data2, data3}
		for _, tx := range data {
			err := store.Insert(tx)
			Expect(err).Should(Succeed())
		}

		transactions, err := store.FindTransactionsByBlockHash(data1.BlockHash)
		Expect(err).Should(Succeed())
		Expect(2).Should(Equal(len(transactions)))
		Expect(reflect.DeepEqual(*transactions[0], *data1)).Should(BeTrue())
		Expect(reflect.DeepEqual(*transactions[1], *data2)).Should(BeTrue())
	})
})

func TestTransaction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transaction Test")
}