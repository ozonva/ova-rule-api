package flusher

import (
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mocks "github.com/ozonva/ova-rule-api/internal/mocks"
	"github.com/ozonva/ova-rule-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl        *gomock.Controller
		mockRepo    *mocks.MockRepo
		testFlusher Flusher
		rules       []models.Rule
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("сохранение правил батчами", func() {
		Context("когда валидный chunkSize", func() {
			BeforeEach(func() {
				testFlusher = NewFlusher(2, mockRepo)
				rules = createRules(5)
			})

			Context("когда нет ошибок при сохранении", func() {
				It("должен сохранить все правила и вернуть nil", func() {
					mockRepo.EXPECT().AddRules(rules[:2]).Return(nil).Times(1)
					mockRepo.EXPECT().AddRules(rules[2:4]).Return(nil).Times(1)
					mockRepo.EXPECT().AddRules(rules[4:]).Return(nil).Times(1)

					result := testFlusher.Flush(rules)

					Expect(result).To(BeNil())
				})
			})

			Context("когда часть батчей не сохранилась", func() {
				It("должен вернуть список несохраненных правил", func() {
					mockRepo.EXPECT().AddRules(rules[:2]).Return(nil).Times(1)
					mockRepo.EXPECT().AddRules(rules[2:4]).Return(errors.New("ошибка сохранения")).Times(1)
					mockRepo.EXPECT().AddRules(rules[4:]).Return(nil).Times(1)

					result := testFlusher.Flush(rules)

					Expect(result).To(Equal(rules[2:4]))
				})
			})

			Context("когда пустой список правил", func() {
				It("должен вернуть пустой список", func() {
					rules = createRules(0)

					result := testFlusher.Flush(rules)

					Expect(result).To(BeEmpty())
				})
			})
		})

		Context("когда невалидный chunkSize", func() {
			BeforeEach(func() {
				testFlusher = NewFlusher(-1, mockRepo)
				rules = createRules(5)
			})

			It("должен вернуть все изначальные правила", func() {
				result := testFlusher.Flush(rules)

				Expect(result).To(Equal(rules))
			})
		})
	})
})

func createRules(number int) []models.Rule {
	result := make([]models.Rule, number)

	for i := 0; i < number; i++ {
		result[i] = models.Rule{
			ID:     int64(i),
			Name:   fmt.Sprintf("some rule %d", i),
			UserID: int64(i),
		}
	}

	return result
}
