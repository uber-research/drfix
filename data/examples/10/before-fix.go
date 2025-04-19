func BuildStoreAdapterWithUpdatedArrangement(
	fetchActorAccountJob *accountjobs.FetchActorAccountJob,
	someDeliveryResult *datamodels.SomeDeliveryResponse,
) func() (winterfell.StoreAdapter, bool) {
	return func() (wrapper winterfell.StoreAdapter, ok bool) {
		actorAccount := fetchActorAccountJob.RestaurantAccount

		if isValid := isValidActorAcct(actorAccount); !isValid {
			return nil, false
		}

		actorUUID := string(*actorAccount.ActorInternalAcct.UUID)
		if someDeliveryResult != nil &&
			someDeliveryResult.DeliverabilityMap != nil {
			actorDelivery := someDeliveryResult.DeliverabilityMap[generic.UUID(actorUUID)]
			if actorDelivery != nil {
				actorAccount.ActorInternalAcct.ArrangementType = actorDelivery.ArrangementType
			}
		}
		return &winterfell.ActorAccountWrapper{
			Base: *actorAccount.ActorInternalAcct,
		}, true
	}
}
