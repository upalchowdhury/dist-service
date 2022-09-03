package discovery

import (
	"github.com/hashicorp/serf/serf"
	"go.uber.org/zap"
)

type Membership struct {
	serf.Config
	handler Handler
	serf    *serf.Serf
	events  chan serf.Event
	logger  *zap.Logger
}
