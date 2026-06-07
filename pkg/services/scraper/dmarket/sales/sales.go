package sales

import (
	"context"
	"errors"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/dmarket"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"go.uber.org/zap"
)

var ErrRateLimitSurpassed = errors.New("rate limit has been surpassed (429)")

var ErrBadStatusCode = errors.New("bad status code")

func (s *Service) getSales(ctx context.Context, c *dmarket.ClientWithResponses, params *dmarket.AggregatorGetLastSalesParams) ([]models.Sale, error) {
	resp, err := c.AggregatorGetLastSalesWithResponse(ctx, params)
	if err != nil {
		s.logger.Error("Failed to fetch last sales...", zap.Error(err))
		return nil, err
	}

	statusCode := resp.StatusCode()
	if statusCode == 429 {
		s.logger.Warn("rate limit surpassed")
		return nil, ErrRateLimitSurpassed
	}
	if statusCode != 200 && statusCode != 204 {
		s.logger.Warn("Bad Status code", zap.Int("code", statusCode))
		return nil, fmt.Errorf("%v %w", statusCode, ErrBadStatusCode)
	}
	orders, err := resp.JSON200.GetSales(params.Title, s.filters...)
	if err != nil {
		s.logger.Error(
			"failed getting orders",
			zap.Any("resp", resp.JSON200),
			zap.Error(err),
		)
		return nil, err
	}
	return orders, nil
}
