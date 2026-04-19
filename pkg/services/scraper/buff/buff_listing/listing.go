package bufflisting

import (
	"context"
	"errors"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/buff"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
	buff_utils "github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/buff"
	"go.uber.org/zap"
)

var ErrRateLimitSurpassed = errors.New("rate limit has been surpassed (429)")

var ErrBadStatusCode = errors.New("bad status code")

func (s *Service) getListings(ctx context.Context, c *buff.ClientWithResponses, param buff_utils.GetListingParams) ([]models.Item, error) {

	p := &buff.GetApiMarketGoodsSellOrderParams{
		GoodsId:               param.Item.Buff163GoodsId,
		Game:                  "csgo",
		PageNum:               utils.Ptr(1),
		SortBy:                utils.Ptr("default"),
		AllowTradableCooldown: utils.Ptr(1),
	}

	resp, err := c.GetApiMarketGoodsSellOrderWithResponse(ctx, p)
	if err != nil {
		s.logger.Error(
			"Failed to fetch buff new listings...",
			zap.Any("params", param),
			zap.Error(err),
		)
		return nil, err
	}

	statusCode := resp.StatusCode()
	if statusCode == 429 {
		s.logger.Warn("rate limit surpassed")
		return nil, ErrRateLimitSurpassed
	}

	if statusCode == 403 {
		s.logger.Error("code 403", zap.Any("params", p))
	}

	if statusCode != 200 && statusCode != 204 {
		s.logger.Warn("Bad Status code", zap.Int("code", statusCode))
		return nil, fmt.Errorf("%s %w", statusCode, ErrBadStatusCode)
	}
	items, err := resp.JSON200.Items(param.Item, param.MarketHashName, s.filters...)
	if err != nil {
		s.logger.Error(
			"failed getting listings",
			// zap.Any("resp", resp.JSON200),
			zap.Error(err),
		)
		return nil, err
	}
	return items, nil
}
