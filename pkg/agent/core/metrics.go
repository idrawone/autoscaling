package core

// Definition of the Metrics type, plus reading it from vector.dev's prometheus format host metrics

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/neondatabase/autoscaling/pkg/api"
)

type Metrics struct {
	LoadAverage1Min  float32
	MemoryUsageBytes float32
}

func (m Metrics) ToAPI() api.Metrics {
	return api.Metrics{
		LoadAverage1Min:  m.LoadAverage1Min,
		LoadAverage5Min:  nil,
		MemoryUsageBytes: nil,
	}
}

// ReadMetrics generates Metrics from vector.dev's host metrics output, or returns error on failure
//
// This function could be more efficient, but realistically it doesn't matter. The size of the
// output from node_exporter/vector is so small anyways.
func ReadMetrics(nodeExporterOutput []byte, loadPrefix string) (m Metrics, err error) {
	lines := strings.Split(string(nodeExporterOutput), "\n")

	getField := func(linePrefix, dontMatch string) (float32, error) {
		var line string
		for _, l := range lines {
			if strings.HasPrefix(l, linePrefix) && (len(dontMatch) == 0 || !strings.HasPrefix(l, dontMatch)) {
				line = l
				break
			}
		}
		if line == "" {
			return 0, fmt.Errorf("No line in metrics output starting with %q", linePrefix)
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			return 0, fmt.Errorf(
				"Expected >= 2 fields in metrics output for %q. Got %v",
				linePrefix, len(fields),
			)
		}

		v, err := strconv.ParseFloat(fields[1], 32)
		if err != nil {
			return 0, fmt.Errorf(
				"Error parsing %q as float for line starting with %q: %w",
				fields[1], linePrefix, err,
			)
		}
		return float32(v), nil
	}

	m.LoadAverage1Min, err = getField(loadPrefix+"load1", loadPrefix+"load15")
	if err != nil {
		return
	}

	availableMem, err := getField(loadPrefix+"memory_available_bytes", "")
	if err != nil {
		return
	}
	totalMem, err := getField(loadPrefix+"memory_total_bytes", "")
	if err != nil {
		return
	}

	// Add an extra 100 MiB to account for kernel memory usage
	m.MemoryUsageBytes = totalMem - availableMem + 100*(1<<20)

	return
}
