package shop

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRepo_SetPM(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		var shopID int64 = 64

		// new gomock
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// create mock for repo interface
		mockKYC := NewMockKYC(ctrl)
		mockShopScore := NewMockShopScore(ctrl)
		mockOrder := NewMockOrder(ctrl)
		mockShopClass := NewMockShopClass(ctrl)

		mockKYC.EXPECT().CheckKYC(ctx, shopID).Return(true, nil)
		mockShopScore.EXPECT().GetShopScore(ctx, shopID).Return(85, nil)
		mockOrder.EXPECT().GetNetIncome(ctx, shopID).Return(int64(400000), nil)
		mockOrder.EXPECT().GetTotalFinishOrder(ctx, shopID).Return(int64(5), nil)
		mockShopClass.EXPECT().SetMembership(ctx, shopID, 2).Return(nil)

		uc := NewUsecase(mockKYC, mockShopScore, mockOrder, mockShopClass)
		newStatus, err := uc.SetPM(context.Background(), 64)
		assert.Nil(t, err)
		assert.Equal(t, 2, newStatus)
	})

}
