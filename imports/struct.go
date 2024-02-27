package imports

import "time"

type StreamSchedule []struct {
	Data struct {
		CurrentUser any `json:"currentUser"`
		User        struct {
			ID              string `json:"id"`
			PrimaryColorHex any    `json:"primaryColorHex"`
			LastBroadcast   struct {
				ID        string    `json:"id"`
				StartedAt time.Time `json:"startedAt"`
				Typename  string    `json:"__typename"`
			} `json:"lastBroadcast"`
			BroadcastSettings struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				Typename string `json:"__typename"`
			} `json:"broadcastSettings"`
			Stream any `json:"stream"`
			Videos struct {
				Edges    []any  `json:"edges"`
				Typename string `json:"__typename"`
			} `json:"videos"`
			Channel struct {
				ID       string `json:"id"`
				Schedule any    `json:"schedule"`
				Typename string `json:"__typename"`
			} `json:"channel"`
			Typename string `json:"__typename"`
		} `json:"user"`
	} `json:"data"`
	Extensions struct {
		DurationMilliseconds int    `json:"durationMilliseconds"`
		OperationName        string `json:"operationName"`
		RequestID            string `json:"requestID"`
	} `json:"extensions"`
	Data0 struct {
		User struct {
			ID        string `json:"id"`
			Followers struct {
				TotalCount int    `json:"totalCount"`
				Typename   string `json:"__typename"`
			} `json:"followers"`
			IsPartner       bool   `json:"isPartner"`
			PrimaryColorHex any    `json:"primaryColorHex"`
			Typename        string `json:"__typename"`
		} `json:"user"`
	} `json:"data0"`
}

type UserInfo []struct {
	Data struct {
		TargetUser struct {
			ID          string `json:"id"`
			Login       string `json:"login"`
			DisplayName string `json:"displayName"`
			Typename    string `json:"__typename"`
		} `json:"targetUser"`
		CurrentUser any `json:"currentUser"`
		RequestInfo struct {
			CountryCode string `json:"countryCode"`
			Typename    string `json:"__typename"`
		} `json:"requestInfo"`
		ReportWizard struct {
			Reasons struct {
				ID                   string `json:"id"`
				CountryCode          string `json:"countryCode"`
				DisclosureText       any    `json:"disclosureText"`
				ToSAndCountryReasons []struct {
					ID                               string `json:"id"`
					IsApplicableToCountryRegulations bool   `json:"isApplicableToCountryRegulations"`
					Text                             string `json:"text"`
					Description                      string `json:"description"`
					DeadEndType                      string `json:"deadEndType"`
					DetailedReasons                  any    `json:"detailedReasons"`
					Typename                         string `json:"__typename"`
				} `json:"toSAndCountryReasons"`
				Typename string `json:"__typename"`
			} `json:"reasons"`
			ReportableContent []struct {
				ID                string `json:"id"`
				Type              string `json:"type"`
				ApplicableReasons []struct {
					ID           string `json:"id"`
					Visibility   string `json:"visibility"`
					ReportReason struct {
						ID       string `json:"id"`
						Typename string `json:"__typename"`
					} `json:"reportReason"`
					Typename string `json:"__typename"`
				} `json:"applicableReasons"`
				DeadEndType string `json:"deadEndType"`
				Typename    string `json:"__typename"`
			} `json:"reportableContent"`
			Typename string `json:"__typename"`
		} `json:"reportWizard"`
	} `json:"data"`
	Extensions struct {
		DurationMilliseconds int    `json:"durationMilliseconds"`
		OperationName        string `json:"operationName"`
		RequestID            string `json:"requestID"`
	} `json:"extensions"`
}
