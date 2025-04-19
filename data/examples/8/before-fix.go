func (m *Monitor) bigDataMonitor(
	ctx context.Context, baseFields []zap.Field,
	tuning *config.TuningConfig, queryTemplate, zone, env, measureType string) (stats.Keys, error) {
	defer m.monitorTimer(map[string]string{
		"part":             "largekeys",
		"controller-zone":  zone,
		"controller-env":   env,
		"statstype":        measureType,
		"data-source":      someDataSource,
	}).Stop()
	upperBound := tuning.Tolerance[1]
	kstats := make(stats.Keys)
	calculateFrom := time.Now().Add(-4 * time.Minute)
	keysList := getKeysProm()
	results := make([]promql.Matrix, len(keysList))

	respChan := make(chan queryRespProm, len(keysList))
	keysChan := make(chan string, len(keysList))

	for w := 1; w <= _concurrentQueryWorkers; w++ {
		go func(keysChan chan string, respChan chan queryRespProm) {
			for keys := range keysChan {
				var tooLargeQuery string
				var err error
				promTuningWindow := convertToPromTimeFormat(tuning.Window)
				tooLargeQuery = fmt.Sprintf(queryTemplate, m.clusterName, keys, promTuningWindow, upperBound)
				end := time.Now()
				start := end.Add(-time.Duration(10) * time.Second)
				interval := time.Duration(30) * time.Second
				series, err := m.mc.RangeQuerySeries(ctx, tooLargeQuery, start, end, interval, "manager_bigdata")
				respChan <- queryRespProm{
					err:    err,
					result: series,
				}
				if err != nil {
					for key := range keys {
						keyTags := map[string]string{
							"part":                  "largekeys",
							"key":                   fmt.Sprint(key),
							"cluster":               m.clusterName,
							"controller-serviceName": m.serviceName,
							"statstype":             measureType,
						}
						m.monitorStats.queryFailure.Tagged(keyTags).Inc(1)
					}
					continue
				}
				baseFields = append(baseFields,
					zap.String("manager-service", m.clusterName),
					zap.String("manager-zone", zone),
					zap.String("manager-env", env),
					zap.String("result", series.String()),
					zap.String("query", tooLargeQuery),
					zap.String("keys", keys),
					zap.Time("start", start),
					zap.Time("end", end),
					zap.Float64("intervalSec", interval.Seconds()))
				m.logger.
					Info("Manager: Query Success.", baseFields...,
					)
			}
		}(keysChan, respChan)
	}

	for _, keys := range keysList {
		keysChan <- keys
	}
	close(keysChan)

	var errs error
	for i := 0; i < len(keysList); i++ {
		qresp := <-respChan
		if qresp.err != nil {
			errs = multierr.Append(errs, qresp.err)
		} else {
			results = append(results, qresp.result)
		}
	}

	if errs != nil {
		return nil, errs
	}

	for _, result := range results {
		for _, r := range result {
			key := r.Metric.Get("key")
			part := r.Metric.Get(measureType)
			filteredTotal := 0.0
			filteredCnt := 0
			for _, point := range r.Points {
				if point.T >= calculateFrom.Unix() && point.V != 0 {
					filteredCnt++
					filteredTotal += point.V
				}
			}
			if filteredCnt == 0 {
				kstats[key] = append(kstats[key], stats.Part{Num: part, Value: 0})
			} else {
				kstats[key] = append(kstats[key], stats.Part{Num: part, Value: filteredTotal / float64(filteredCnt)})
			}
		}
	}
	return kstats, nil
}
