package processors

import (
	"github.com/VoIPGRID/opensips_exporter/opensips"
	"github.com/prometheus/client_golang/prometheus"
)

// tmProcessor exposes metrics for stateful processing of SIP transactions.
// doc: http://www.opensips.org/html/docs/modules/1.11.x/tm.html#idp5881664
// src: https://github.com/OpenSIPS/opensips/blob/1.11/modules/tm/tm.c#L283
type tmProcessor struct {
	statistics map[string]opensips.Statistic
}

var tmLabelNames = []string{}
var tmMetrics = map[string]metric{
	"received_replies":   newMetric("tm", "received_replies_total", "Total number of total replies received by TM module.", tmLabelNames, prometheus.CounterValue),
	"relayed_replies":    newMetric("tm", "relayed_replies_total", "Total number of replies received and relayed by TM module.", tmLabelNames, prometheus.CounterValue),
	"local_replies":      newMetric("tm", "local_replies_total", "Total number of replies local generated by TM module.", tmLabelNames, prometheus.CounterValue),
	"UAS_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"UAC_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"2xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"3xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"4xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"5xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"6xx_transactions":   newMetric("tm", "transactions_total", "Total number of transactions.", []string{"type"}, prometheus.CounterValue),
	"inuse_transactions": newMetric("tm", "inuse_transactions", "Number of transactions existing in memory at current time.", tmLabelNames, prometheus.GaugeValue),
}

func init() {
	for metric := range tmMetrics {
		OpensipsProcessors[metric] = tmProcessorFunc
	}
	OpensipsProcessors["tm:"] = tmProcessorFunc
}

// Describe implements prometheus.Collector.
func (p tmProcessor) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range tmMetrics {
		ch <- metric.Desc
	}
}

// Collect implements prometheus.Collector.
func (p tmProcessor) Collect(ch chan<- prometheus.Metric) {
	for _, s := range p.statistics {
		if s.Module == "tm" {
			switch s.Name {
			case "received_replies":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["received_replies"].Desc,
					tmMetrics["received_replies"].ValueType,
					s.Value,
				)
			case "relayed_replies":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["relayed_replies"].Desc,
					tmMetrics["relayed_replies"].ValueType,
					s.Value,
				)
			case "local_replies":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["local_replies"].Desc,
					tmMetrics["local_replies"].ValueType,
					s.Value,
				)
			case "UAS_transactions":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["UAS_transactions"].Desc,
					tmMetrics["UAS_transactions"].ValueType,
					s.Value,
					"UAS",
				)
			case "UAC_transactions":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["UAC_transactions"].Desc,
					tmMetrics["UAC_transactions"].ValueType,
					s.Value,
					"UAC",
				)
			case "2xx_transactions":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["2xx_transactions"].Desc,
					tmMetrics["2xx_transactions"].ValueType,
					s.Value,
					"2xx",
				)
			case "3xx_transactions":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["3xx_transactions"].Desc,
					tmMetrics["3xx_transactions"].ValueType,
					s.Value,
					"3xx",
				)
			case "4xx_transactions":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["4xx_transactions"].Desc,
					tmMetrics["4xx_transactions"].ValueType,
					s.Value,
					"4xx",
				)
			case "5xx_transactions":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["5xx_transactions"].Desc,
					tmMetrics["5xx_transactions"].ValueType,
					s.Value,
					"5xx",
				)
			case "6xx_transactions":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["6xx_transactions"].Desc,
					tmMetrics["6xx_transactions"].ValueType,
					s.Value,
					"6xx",
				)
			case "inuse_transactions":
				ch <- prometheus.MustNewConstMetric(
					tmMetrics["inuse_transactions"].Desc,
					tmMetrics["inuse_transactions"].ValueType,
					s.Value,
				)
			}
		}
	}
}

func tmProcessorFunc(s map[string]opensips.Statistic) prometheus.Collector {
	return &tmProcessor{
		statistics: s,
	}
}
