package ova_rule_api

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-rule-api/internal/flusher"
	"github.com/ozonva/ova-rule-api/internal/mocks"
	"github.com/ozonva/ova-rule-api/internal/models"
	"github.com/ozonva/ova-rule-api/internal/saver"
	"time"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

var _ = Describe("Service", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		api      desc.APIServer
		ctx      context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		ctx = context.Background()
		flusher_ := flusher.NewFlusher(10, mockRepo)
		saver_ := saver.NewSaver(100, flusher_, 100*time.Millisecond)
		saver_.Init()

		api = NewAPIServer(mockRepo, saver_)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Позитивные кейсы.", func() {
		It("CreateRule", func() {
			someRule := models.Rule{ID: 1, Name: "awesome", UserID: 777}
			mockRepo.EXPECT().AddRules([]models.Rule{someRule}).Times(1)

			_, err := api.CreateRule(ctx, &desc.CreateRuleRequest{
				Id:     1,
				Name:   "awesome",
				UserId: 777,
			})
			// Дадим время для сохранения правила через очередь.
			time.Sleep(200 * time.Millisecond)

			Expect(err).ShouldNot(HaveOccurred())
		})

		It("DescribeRule", func() {
			someRule := models.Rule{ID: 1, Name: "awesome", UserID: 777}
			mockRepo.EXPECT().DescribeRule(uint64(1)).Return(&someRule, nil).Times(1)

			res, err := api.DescribeRule(ctx, &desc.DescribeRuleRequest{
				Id: 1,
			})
			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.Result.Id).To(Equal(someRule.ID))
			Expect(res.Result.Name).To(Equal(someRule.Name))
			Expect(res.Result.UserId).To(Equal(someRule.UserID))
		})

		It("ListRules", func() {
			someRule := models.Rule{ID: 1, Name: "awesome", UserID: 777}
			mockRepo.EXPECT().ListRules(uint64(10), uint64(0)).Return([]models.Rule{someRule}, nil).Times(1)

			res, err := api.ListRules(ctx, &desc.ListRulesRequest{
				Limit:  uint64(10),
				Offset: uint64(0),
			})
			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.Result[0].Id).To(Equal(someRule.ID))
			Expect(res.Result[0].Name).To(Equal(someRule.Name))
			Expect(res.Result[0].UserId).To(Equal(someRule.UserID))
		})

		It("RemoveRule", func() {
			mockRepo.EXPECT().RemoveRule(uint64(777)).Return(nil).Times(1)

			_, err := api.RemoveRule(ctx, &desc.RemoveRuleRequest{
				Id: uint64(777),
			})
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Негативные кейсы. Ошибки валидации.", func() {
		It("CreateRule", func() {
			_, err := api.CreateRule(ctx, &desc.CreateRuleRequest{
				Id:     0,
				Name:   "",
				UserId: 0,
			})

			Expect(err).Should(HaveOccurred())
		})

		It("DescribeRule", func() {
			_, err := api.DescribeRule(ctx, &desc.DescribeRuleRequest{
				Id: 0,
			})
			Expect(err).Should(HaveOccurred())
		})

		It("ListRules", func() {
			_, err := api.ListRules(ctx, &desc.ListRulesRequest{
				Limit:  uint64(0),
				Offset: uint64(0),
			})
			Expect(err).Should(HaveOccurred())
		})

		It("RemoveRule", func() {
			_, err := api.RemoveRule(ctx, &desc.RemoveRuleRequest{
				Id: uint64(0),
			})
			Expect(err).Should(HaveOccurred())
		})
	})
})
