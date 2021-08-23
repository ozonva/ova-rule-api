package saver

import (
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-rule-api/internal/mocks"
	"github.com/ozonva/ova-rule-api/internal/models"
)

const (
	shortTimeout    = 100 * time.Millisecond
	middleTimeout   = 200 * time.Millisecond
	longTimeout     = 300 * time.Millisecond
	veryLongTimeout = time.Minute
)

var _ = Describe("Saver", func() {
	var (
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher
		testSaver   Saver
	)

	Describe("сохранение правил через saver методом Save", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockFlusher = mocks.NewMockFlusher(ctrl)
		})

		AfterEach(func() {
			ctrl.Finish()
		})

		Context("в буффере есть свободное место", func() {
			It("правило добавили в буффер, метод Flush не вызывается", func() {
				testSaver = NewSaver(3, mockFlusher, veryLongTimeout)
				testSaver.Init()
				someRule := models.Rule{ID: 1, Name: "Some rule"}
				mockFlusher.EXPECT().Flush(gomock.Any()).Times(0)

				err := testSaver.Save(someRule)

				time.Sleep(shortTimeout)

				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("в буффере нет свободного место", func() {
			It("при сохранении правила возникает ошибка", func() {
				testSaver = NewSaver(3, mockFlusher, veryLongTimeout)
				testSaver.Init()
				mockFlusher = mocks.NewMockFlusher(ctrl)
				for i := 0; i < 3; i++ {
					err := testSaver.Save(models.Rule{
						ID: int64(i),
					})
					Expect(err).ShouldNot(HaveOccurred())
				}

				someRule := models.Rule{ID: 1, Name: "Some rule"}
				err := testSaver.Save(someRule)

				Expect(err).Should(HaveOccurred())
			})
		})

		Context("случился таймаут для сброса данных в хранилище", func() {
			It("вызвали метод Flush у flusher", func() {
				testSaver = NewSaver(3, mockFlusher, middleTimeout)
				testSaver.Init()
				someRule := models.Rule{ID: 1, Name: "Some rule"}
				mockFlusher.EXPECT().Flush([]models.Rule{someRule}).Times(1)

				err := testSaver.Save(someRule)
				Expect(err).ShouldNot(HaveOccurred())

				time.Sleep(longTimeout)
			})
		})

		Context("закрыли канал методов Close", func() {
			It("вызвали метод Flush у flusher", func() {
				testSaver = NewSaver(3, mockFlusher, middleTimeout)
				testSaver.Init()
				someRule := models.Rule{ID: 1, Name: "Some rule"}
				mockFlusher.EXPECT().Flush([]models.Rule{someRule}).Times(1)

				err := testSaver.Save(someRule)
				Expect(err).ShouldNot(HaveOccurred())

				testSaver.Close()
				time.Sleep(shortTimeout)
			})
		})
	})
})
