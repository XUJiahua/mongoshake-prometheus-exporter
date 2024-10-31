package probe

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xujiahua/mongoshake-prometheus-exporter/pkg/metrics"
	"github.com/xujiahua/mongoshake-prometheus-exporter/pkg/mongoshake"
)

type MongoShakeClient interface {
	GetRepl() (*mongoshake.Repl, error)
	GetAlias() string
	GetBaseURL() string
}

type Prober struct {
	labels           []string
	mongoShakeClient MongoShakeClient
	interval         time.Duration
	logger           *logrus.Entry
}

func NewProber(mongoShakeClient MongoShakeClient, intervalSeconds int) *Prober {
	interval := 5 * time.Second
	if intervalSeconds > 0 {
		interval = time.Duration(intervalSeconds) * time.Second
	}
	labels := []string{mongoShakeClient.GetAlias(), mongoShakeClient.GetBaseURL()}
	logger := logrus.WithFields(logrus.Fields{
		"alias": mongoShakeClient.GetAlias(),
	})
	return &Prober{labels: labels, mongoShakeClient: mongoShakeClient, interval: interval, logger: logger}
}

func (p Prober) Probe() {
	p.logger.Debugf("start to probe with interval %f seconds", p.interval.Seconds())

	p.probe()

	ticker := time.NewTicker(p.interval)
	defer ticker.Stop()
	for range ticker.C {
		p.probe()
	}
}

func (p Prober) probe() {
	p.logger.Debugf("start to probe")
	repl, err := p.mongoShakeClient.GetRepl()
	if err != nil {
		p.logger.Warnf("get repl error: %v", err)
		metrics.Up.WithLabelValues(p.labels...).Set(0)
		return
	}

	metrics.Up.WithLabelValues(p.labels...).Set(1)
	metrics.LogsGet.WithLabelValues(p.labels...).Set(float64(repl.LogsGet))
	metrics.LogsRepl.WithLabelValues(p.labels...).Set(float64(repl.LogsRepl))
	metrics.LogsSuccess.WithLabelValues(p.labels...).Set(float64(repl.LogsSuccess))
	metrics.TPS.WithLabelValues(p.labels...).Set(float64(repl.Tps))

	latencyS := repl.Now.Unix - repl.LsnAck.Unix
	metrics.ReplicationLatency.WithLabelValues(p.labels...).Set(float64(latencyS))
}
