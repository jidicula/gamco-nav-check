package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/jidicula/go-gamco"
)

func dateSetup(priceDate string, inceptionDate string, lastMonthEnd string, lastQtrEnd string) (map[string]time.Time, error) {
	dates := make(map[string]time.Time)
	dateFormat := "01/02/2006"
	var err error
	dates["priceDate"], err = time.Parse(time.RFC3339, priceDate)
	if err != nil {
		return dates, err
	}
	dates["inceptionDate"], err = time.Parse(time.RFC3339, inceptionDate)
	if err != nil {
		return dates, err
	}
	dates["lastMonthEnd"], err = time.Parse(dateFormat, lastMonthEnd)
	if err != nil {
		return dates, err
	}
	dates["lastQtrEnd"], err = time.Parse(dateFormat, lastQtrEnd)
	if err != nil {
		return dates, err
	}

	return dates, err
}

var datesMap, _ = dateSetup("2021-04-01T00:00:00.000Z", "1999-07-09T00:00:00.000Z", "03/31/2021", "03/31/2021")
var fl = []gamco.Fund{
	gamco.Fund{
		ID:                   515,
		FundCode:             -113,
		SecurityID:           "36240A101",
		FundShortName:        "Utility Trust",
		NAVDate:              datesMap["priceDate"],
		NAV:                  "4.27",
		PriorNAV:             "4.25",
		Change:               "0.02",
		PctChange:            "0.004706",
		Sort:                 "43.0",
		YtdReturn:            0.0767238547,
		YtdReturnMonthly:     0.0716806516,
		YtdReturnQuarterly:   0.0716806516,
		OneYrReturn:          0.3799871231,
		OneYrReturnMonthly:   0.2908244253,
		OneYrReturnQuarterly: 0.2908244253,
		ThreeYrAvg:           0.0821346464,
		ThreeYrAvgMonthly:    0.0806954728,
		ThreeYrAvgQuarterly:  0.0806954728,
		FiveYrAvg:            0.0628642424,
		FiveYrAvgMonthly:     0.0618578521,
		FiveYrAvgQuarterly:   0.0618578521,
		TenYrAvg:             0.0852284004,
		TenYrAvgMonthly:      0.086077124,
		TenYrAvgQuarterly:    0.086077124,
		InceptAvg:            0.0851655141,
		InceptAvgMonthly:     0.085533555,
		InceptAvgQuarterly:   0.0869319686,
		Symbol:               "GUT",
		AssetType:            "Equity",
		InceptionDate:        datesMap["inceptionDate"],
		LegalName2:           "The Gabelli Utility Trust",
		SeriesName:           "",
		DisplayName:          "Gabelli Utility Trust",
		DisplayName_:         "The Gabelli Utility Trust",
		Category:             "value",
		AnnualReport:         "https://gab-annual-reports.s3.us-east-2.amazonaws.com/GUTFundWebReady12312020.pdf",
		SemiAnnualReport:     "https://gab-semi-annuals.s3.us-east-2.amazonaws.com/TheGabelliUtilityTrust606302020.pdf",
		Cusip:                "36240A101",
		QuarterlyReport:      "https://gab-reports.s3.us-east-2.amazonaws.com/2006q3/-113.pdf",
		Prospectus:           "https://gab-prospectus.s3.us-east-2.amazonaws.com/-113.pdf",
		Sai:                  "https://gab-sai.s3.us-east-2.amazonaws.com/-113_sai.pdf",
		Soi:                  "",
		Factsheet:            "https://gab-factsheets.s3.us-east-2.amazonaws.com/closedEnd_FactSheets4Q2020DRAFT_GUT12312020.pdf",
		Commentary:           "https://gab-commentary-pdf.s3.us-east-2.amazonaws.com/WEB_CEF_4Q2012312020.pdf",
		LastMonthEnd:         datesMap["lastMonthEnd"],
		LastQtrEnd2:          datesMap["lastQtrEnd"],
	},
	gamco.Fund{
		ID:                   515,
		FundCode:             -113,
		SecurityID:           "36240A101",
		FundShortName:        "Utility Trust",
		NAVDate:              datesMap["priceDate"],
		NAV:                  "4.27",
		PriorNAV:             "4.25",
		Change:               "0.02",
		PctChange:            "0.004706",
		Sort:                 "43.0",
		YtdReturn:            0.0767238547,
		YtdReturnMonthly:     0.0716806516,
		YtdReturnQuarterly:   0.0716806516,
		OneYrReturn:          0.3799871231,
		OneYrReturnMonthly:   0.2908244253,
		OneYrReturnQuarterly: 0.2908244253,
		ThreeYrAvg:           0.0821346464,
		ThreeYrAvgMonthly:    0.0806954728,
		ThreeYrAvgQuarterly:  0.0806954728,
		FiveYrAvg:            0.0628642424,
		FiveYrAvgMonthly:     0.0618578521,
		FiveYrAvgQuarterly:   0.0618578521,
		TenYrAvg:             0.0852284004,
		TenYrAvgMonthly:      0.086077124,
		TenYrAvgQuarterly:    0.086077124,
		InceptAvg:            0.0851655141,
		InceptAvgMonthly:     0.085533555,
		InceptAvgQuarterly:   0.0869319686,
		Symbol:               "GGT",
		AssetType:            "Equity",
		InceptionDate:        datesMap["inceptionDate"],
		LegalName2:           "The Gabelli Utility Trust",
		SeriesName:           "",
		DisplayName:          "Gabelli Utility Trust",
		DisplayName_:         "The Gabelli Utility Trust",
		Category:             "value",
		AnnualReport:         "https://gab-annual-reports.s3.us-east-2.amazonaws.com/GGTFundWebReady12312020.pdf",
		SemiAnnualReport:     "https://gab-semi-annuals.s3.us-east-2.amazonaws.com/TheGabelliUtilityTrust606302020.pdf",
		Cusip:                "36240A101",
		QuarterlyReport:      "https://gab-reports.s3.us-east-2.amazonaws.com/2006q3/-113.pdf",
		Prospectus:           "https://gab-prospectus.s3.us-east-2.amazonaws.com/-113.pdf",
		Sai:                  "https://gab-sai.s3.us-east-2.amazonaws.com/-113_sai.pdf",
		Soi:                  "",
		Factsheet:            "https://gab-factsheets.s3.us-east-2.amazonaws.com/closedEnd_FactSheets4Q2020DRAFT_GGT12312020.pdf",
		Commentary:           "https://gab-commentary-pdf.s3.us-east-2.amazonaws.com/WEB_CEF_4Q2012312020.pdf",
		LastMonthEnd:         datesMap["lastMonthEnd"],
		LastQtrEnd2:          datesMap["lastQtrEnd"],
	},
}

func TestExtractNAVs(t *testing.T) {
	tests := map[string]struct {
		fundList []gamco.Fund
		want     map[string]string
	}{

		"two funds": {
			fundList: fl,
			want:     map[string]string{"GUT": fl[0].NAV, "GGT": fl[1].NAV},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := extractNAVs(tt.fundList)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: got %v, want %v", name, got, tt.want)
			}
		})
	}
}

func TestExtractPrices(t *testing.T) {
	tests := map[string]struct {
		fundList []gamco.Fund
		want     map[string]string
	}{
		"two funds": {
			fundList: fl,
			want:     map[string]string{"GUT": fl[0].NAV, "GGT": fl[1].NAV},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := extractPrices(tt.fundList)
			if err != nil {
				t.Fatalf("%s: %s", name, err)
			}
			for _, v := range tt.fundList {
				_, ok := got[v.Symbol]
				if !ok {
					t.Errorf("%s: unable to get price for %s", name, v.Symbol)
				}
			}

		})
	}
}

func TestGetDiscount(t *testing.T) {
	tests := map[string]struct {
		price string
		nav   string
		want  int
	}{
		"underpriced": {
			price: "1.00",
			nav:   "1.20",
			want:  20,
		},
		"overpriced": {
			price: "1.00",
			nav:   "0.80",
			want:  -20,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := getDiscount(tt.nav, tt.price)
			if err != nil {
				t.Fatalf("%s: %s", name, err)
			}

			if got != tt.want {
				t.Errorf("%s: got %v, want %v", name, got, tt.want)
			}
		})
	}
}

func TestComparePriceNAVMaps(t *testing.T) {
	tests := map[string]struct {
		navMap   map[string]string
		priceMap map[string]string
		want     []Stock
	}{
		"1 underpriced": {
			navMap:   map[string]string{"GUT": "1.20"},
			priceMap: map[string]string{"GUT": "1.00"},
			want: []Stock{{
				Symbol:   "GUT",
				NAV:      "1.20",
				Price:    "1.00",
				Discount: 20,
			}},
		},
		"1 overpriced": {
			navMap:   map[string]string{"GUT": "0.80"},
			priceMap: map[string]string{"GUT": "1.00"},
			want:     []Stock{},
		},
		"2 overpriced": {
			navMap:   map[string]string{"GUT": "0.80", "GGT": "0.80"},
			priceMap: map[string]string{"GUT": "1.00", "GGT": "1.00"},
			want:     []Stock{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := getDiscounts(tt.navMap, tt.priceMap)
			if err != nil {
				t.Fatalf("%s: %s", name, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: got %v, want %v", name, got, tt.want)
			}
		})
	}

}
