func (s *queryManager) LoadDataAndErrorRate(
	queryContext *params.QueryContext,
	allContexts []*params.QueryContext,
	stores slimstore.SlimStores,
) (slimstore.SlimStores, error) {
	startTime := time.Now()
	metricsTags := lib.GetLatencyMetricTags(queryContext.Query)
	metricsTags["plan_name"] = queryContext.Name
	metricsTags["is_cumulative"] = strconv.FormatBool(len(queryContext.CarouselNames) > 1)
	metricsTags["parallel_special_calls"] = strconv.FormatBool(queryContext.Query.GetMakeParallelSiaCalls())
	defer func() {
		s.ServiceDeps.Metrics.Tagged(metricsTags).Histogram("storeindex.histogram.somequerymanager.loaddataanderrorrate",
			observabilityhelpers.DurationHistogramBuckets,
		).RecordDuration(time.Since(startTime))
	}()

	var collectedStores slimstore.SlimStores
	var err error
	var group safegroup.Group
	uuidErrorRateMap := make(map[string]float64)
	group.Go(
		func() error {
			collectedStores, err = s.LoadDataInBuckets(queryContext, allContexts, stores)
			return err
		},
	)
	group.Go(
		func() error {
			uuidErrorRateMap, err = xhelper.LoadOAData(queryContext, s.ServiceDeps, s.DiscoveryDocstoreClient, stores)
			return err
		},
	)
	err = group.Wait()
	collectedStores = xhelper.PopulateWeightedOADefectRate(queryContext, stores, uuidErrorRateMap)

	return collectedStores, err
}
