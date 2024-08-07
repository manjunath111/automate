package relaxting

import (
	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"

	reportingapi "github.com/chef/automate/api/interservice/compliance/reporting"
	"github.com/chef/automate/components/compliance-service/reporting"
	"github.com/chef/automate/lib/stringutils"
)

//todo - this looks very close to the report depth version of it.. double check and if so lets' harmonize them
func (depth *ProfileDepth) getProfileMinsFromNodesAggs(filters map[string][]string) map[string]elastic.Aggregation {
	aggs := make(map[string]elastic.Aggregation)

	termsQuery := elastic.NewTermsAggregation().
		Field("profiles.profile").Size(reporting.ESize)
	termsQuery.SubAggregation("failures", elastic.NewSumAggregation().
		Field("profiles.controls_sums.failed.total"))
	termsQuery.SubAggregation("passed", elastic.NewSumAggregation().
		Field("profiles.controls_sums.passed.total"))
	termsQuery.SubAggregation("skipped", elastic.NewSumAggregation().
		Field("profiles.controls_sums.skipped.total"))
	termsQuery.SubAggregation("waived", elastic.NewSumAggregation().
		Field("profiles.controls_sums.waived.total"))
	termsQuery.SubAggregation("status", elastic.NewTermsAggregation().
		Field("profiles.status"))

	aggs["totals"] = termsQuery

	return depth.wrap(aggs)
}

//todo - this looks very close to the report depth version of it.. double check and if so lets' harmonize them
func (depth *ProfileDepth) getProfileMinsFromNodesResults(
	filters map[string][]string,
	searchResult *elastic.SearchResult,
	statusFilters []string) ([]reporting.ProfileMin, *reportingapi.ProfileCounts, error) {
	myName := "ProfileDepth::getProfileMinsFromNodesResults"

	profileMins := make([]reporting.ProfileMin, 0)
	statusMap := make(map[string]int, 4)

	if aggRoot, found := depth.unwrap(&searchResult.Aggregations); found {
		totalsAgg, _ := aggRoot.Terms("totals")
		if totalsAgg != nil {
			for _, bucket := range totalsAgg.Buckets {
				profileName, profileId := rightSplit(bucket.Key.(string), "|")

				if profilesFilterArray, found := filters["profile_id"]; found {
					if !stringutils.SliceContains(profilesFilterArray, profileId) {
						continue
					}
				}

				// Using the status of the profile, introduced with inspec 3.0 to overwrite the status calculations from totals
				profileStatusHash := make(map[string]int64, 0)
				statuses, _ := bucket.Aggregations.Terms("status")
				if statuses.Buckets != nil {
					for _, statusBucket := range statuses.Buckets {
						status := statusBucket.Key.(string)
						profileStatusHash[status] = statusBucket.DocCount
					}
				}

				var profileStatus string
				if profileStatusHash["skipped"] > 0 && profileStatusHash["loaded"] == 0 && profileStatusHash[""] == 0 {
					profileStatus = "skipped"
					logrus.Debugf("%s profile_name=%q, root status=skipped", myName, profileName)
				} else if profileStatusHash["failed"] > 0 && profileStatusHash["loaded"] == 0 && profileStatusHash[""] == 0 {
					profileStatus = "failed"
					logrus.Debugf("%s profile_name=%q, root status=failed", myName, profileName)
				} else {
					sumFailures, _ := bucket.Aggregations.Sum("failures")
					sumPassed, _ := bucket.Aggregations.Sum("passed")
					sumSkipped, _ := bucket.Aggregations.Sum("skipped")
					sumWaived, _ := bucket.Aggregations.Sum("waived")
					profileStatus = computeStatus(int32(*sumFailures.Value), int32(*sumPassed.Value), int32(*sumSkipped.Value), int32(*sumWaived.Value))

					logrus.Debugf(
						"%s profile_name=%s, status=%s (sumFailures=%d, sumPassed=%d, sumSkipped=%d, sumWaived=%d)",
						myName, profileName, profileStatus, int32(*sumFailures.Value), int32(*sumPassed.Value),
						int32(*sumSkipped.Value), int32(*sumWaived.Value))
				}

				//let's keep track of the counts even if they're not in the filter so that we may know that they're there for UI chicklets
				statusMap[profileStatus]++

				if len(statusFilters) > 0 && !stringutils.SliceContains(statusFilters, profileStatus) {
					continue
				}

				summary := reporting.ProfileMin{
					Name:   profileName,
					ID:     profileId,
					Status: profileStatus,
				}
				profileMins = append(profileMins, summary)
			}
		}
	}
	logrus.Debugf("%s Done with statusMap=%+v", myName, statusMap)
	counts := &reportingapi.ProfileCounts{
		Total:   int32(statusMap["failed"] + statusMap["passed"] + statusMap["skipped"] + statusMap["waived"]),
		Failed:  int32(statusMap["failed"]),
		Passed:  int32(statusMap["passed"]),
		Skipped: int32(statusMap["skipped"]),
		Waived:  int32(statusMap["waived"]),
	}
	return profileMins, counts, nil
}
