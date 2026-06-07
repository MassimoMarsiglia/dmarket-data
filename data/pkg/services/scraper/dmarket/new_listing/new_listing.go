package newlisting

import (
	"context"
	"errors"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/dmarket"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
	"go.uber.org/zap"
)

var ErrRateLimitSurpassed = errors.New("rate limit has been surpassed (429)")

var ErrBadStatusCode = errors.New("bad status code")

func (s *Service) getNewListings(ctx context.Context, c *dmarket.ClientWithResponses) ([]models.Item, error) {
	resp, err := c.GetMarketItemsWithResponse(ctx, &dmarket.GetMarketItemsParams{
		GameId:   dmarket.GameIDA8db,
		Currency: "USD",
		Limit:    utils.Ptr(20),
		OrderBy:  utils.Ptr("updated"),
		OrderDir: utils.Ptr("desc"),
	})
	if err != nil {
		s.logger.Error("Failed to fetch new listings...", zap.Error(err))
		return nil, err
	}

	statusCode := resp.StatusCode()
	if statusCode == 429 {
		s.logger.Warn("rate limit surpassed")
		return nil, ErrRateLimitSurpassed
	}
	if statusCode != 200 && statusCode != 204 {
		s.logger.Warn("Bad Status code", zap.Int("code", statusCode))
		return nil, fmt.Errorf("%s %w", statusCode, ErrBadStatusCode)
	}
	items, err := resp.JSON200.Items(s.filters...)
	if err != nil {
		s.logger.Error(
			"failed getting listings",
			zap.Any("resp", resp.JSON200),
			zap.Error(err),
		)
		return nil, err
	}
	return items, nil
}
