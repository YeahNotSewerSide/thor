package logdb

import (
	"time"

	"github.com/vechain/thor/v2/metrics"
)

var (
	bucket = []int64{
		0,
		time.Nanosecond.Nanoseconds() * 250,  // 250ns
		time.Nanosecond.Nanoseconds() * 500,  // 500ns
		time.Microsecond.Nanoseconds(),       // 1ms
		time.Microsecond.Nanoseconds() * 5,   // 5ms
		time.Microsecond.Nanoseconds() * 10,  // 10ms
		time.Microsecond.Nanoseconds() * 25,  // 25ms
		time.Microsecond.Nanoseconds() * 50,  // 50ms
		time.Microsecond.Nanoseconds() * 100, // 100ms
		time.Microsecond.Nanoseconds() * 250, // 250ms
		time.Microsecond.Nanoseconds() * 500, // 500ms
		time.Second.Nanoseconds(),            // 1s
	}
	metricEventsQuery            = metrics.LazyLoadHistogram("logdb_duration_read_events", bucket)
	metricProcessedEventsCount   = metrics.LazyLoadGauge("logdb_count_event")
	metricTransfersQuery         = metrics.LazyLoadHistogram("logdb_duration_read_transfers", bucket)
	metricProcesesdTransferCount = metrics.LazyLoadGauge("logdb_count_transfer")
	metricProcessBlockDuration   = metrics.LazyLoadHistogram("logdb_duration_process_block", bucket)
)
