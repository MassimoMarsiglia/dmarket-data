package order_book

import (
	"context"
	"errors"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"go.uber.org/zap"
)

var ErrRateLimitSurpassed = errors.New("rate limit has been surpassed (429)")

var ErrBadStatusCode = errors.New("bad status code")

func (s *Service) getOrderBook(ctx context.Context, c *client.ClientWithResponses, params *client.GetOrderBookParams) ([]models.BuyOrder, error) {
	resp, err := c.GetOrderBookWithResponse(ctx, params)
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
	orders, err := resp.JSON200.GetOrders(params.Title, s.filters...)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
