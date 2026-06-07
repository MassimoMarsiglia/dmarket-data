package sales

// type Module struct {
// 	itemDir string
// 	accDir  string
// 	enabled bool
// 	delay   time.Duration
// 	srvc    *sales.Service
// 	logger  *zap.Logger
// 	nc      *nats.Conn
// }

// func New(logger *zap.Logger, cmd *cobra.Command, cfg config.SalesCfg) (*Module, error) {
// 	sales := &Module{
// 		logger:  logger,
// 		accDir:  cfg.AccDir,
// 		itemDir: cfg.ItemDir,
// 		enabled: true,
// 		delay:   cfg.Delay,
// 		nc:      cfg.Conn,
// 	}
// 	return sales, nil
// }

// func (m *Module) Service() *sales.Service {
// 	return m.srvc
// }

// // InitFlags implements [Module].
// func InitFlags(cmd *cobra.Command) {
// 	cmd.Flags().Duration("dmarket.sales.delay", 100*time.Millisecond, "new listing refresh rate")
// }

// // Name implements [Module].
// func (m *Module) Name() string {
// 	return "dmarket.sales"
// }

// // Run implements [Module].
// func (m *Module) Run(ctx context.Context) error {
// 	if !m.enabled {
// 		return nil
// 	}

// 	salesSvc, err := sales.New(sales.ServiceCfg{
// 		Conn:    m.nc,
// 		ItemDir: m.itemDir,
// 		Delay:   m.delay,
// 		Context: context.Background(),
// 		Logger:  m.logger,
// 		AccDir:  &m.accDir,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	m.srvc = salesSvc
// 	return nil
// }

// var _ (modules.Module) = (*Module)(nil)
