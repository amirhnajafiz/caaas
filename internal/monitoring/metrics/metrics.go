package metrics

// Metrics are being used to monitor the service status.
type Metrics struct {
	requests    map[string]int     // label for each endpoint
	failedCalls map[string]int     // label for each endpoint
	latency     map[string]float64 // label for each endpoint
}

func (m *Metrics) AddRequest(endpoint string) {}

func (m *Metrics) AddFailedCall(endpoint string) {}

func (m *Metrics) ObserveLatency(endpoint string) {}
