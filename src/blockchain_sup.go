package src

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

type BnCoreSup struct {
	gen.Supervisor
}

func (ds *BnCoreSup) Init(args ...etf.Term) (gen.SupervisorSpec, error) {
	spec := gen.SupervisorSpec{
		Name: "bnAppSup",
		Children: []gen.SupervisorChildSpec{
			gen.SupervisorChildSpec{
				Name:  "bnServer01",
				Child: &BnCoreGenServ{},
			},
			gen.SupervisorChildSpec{
				Name:  "bnServer02",
				Child: &BnCoreGenServ{},
				Args:  []etf.Term{12345},
			},
			gen.SupervisorChildSpec{
				Name:  "bnServer03",
				Child: &BnCoreGenServ{},
				Args:  []etf.Term{"abc", 67890},
			},
		},
		Strategy: gen.SupervisorStrategy{
			Type: gen.SupervisorStrategyOneForAll,
			// Type:      gen.SupervisorStrategyRestForOne,
			// Type:      gen.SupervisorStrategyOneForOne,
			Intensity: 2,
			Period:    5,
			Restart:   gen.SupervisorStrategyRestartTemporary,
			// Restart: gen.SupervisorStrategyRestartTransient,
			// Restart: gen.SupervisorStrategyRestartPermanent,
		},
	}
	return spec, nil
}
