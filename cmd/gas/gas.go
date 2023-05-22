// logik for gas usage calculation
package gas

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/webcognito/appKalkulacka/pkg/conversion"
)

type Result struct {
	Uuid             string
	Company          string
	ContractDuration string
	UsageRange       string
	UsageM3          string
	UsageMWh         string
	NumberOfMonths   string
	CostUsage        string
	CostDist         string
	CostDistxMonths  string
	Total            string
	Advance          string
}

func Ginput(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/index.html", "web/templates/content_ginput.html")
	t.ExecuteTemplate(w, "index", "")
}

func Goutput(w http.ResponseWriter, r *http.Request) {
	// Parse Form Values
	company := r.FormValue("company")
	contractDuration := r.FormValue("contractDuration")
	usageRange := r.FormValue("usageRange")
	m3Usage := r.FormValue("m3Usage")
	nMonths := r.FormValue("numberOfMonths")
	priceGas := r.FormValue("priceGas")
	monthlyGas := r.FormValue("monthlyGas")
	priceDist := r.FormValue("priceDist")
	monthDist := r.FormValue("monthlyDist")

	// generate uuid
	nuuid := uuid.New()

	// convert m3 to MWh
	var volumeCoeficient float64 = 0.9968
	var mwhtom3 float64 = 0.0108987
	u := conversion.StringToFloat64(m3Usage)
	convUsage := u * volumeCoeficient * mwhtom3

	// String to Float
	floatPriceGas := conversion.StringToFloat64(priceGas)
	floatMonthlyGas := conversion.StringToFloat64(monthlyGas)
	floatPriceDist := conversion.StringToFloat64(priceDist)
	floatMonthlyDist := conversion.StringToFloat64(monthDist)
	floatNMonths := conversion.StringToFloat64(nMonths)

	// Usage Cost
	usageCost := convUsage * (floatPriceGas + floatPriceDist)

	// Dist Cost
	distCost := floatMonthlyGas + floatMonthlyDist
	distCostxMonths := floatNMonths * (floatMonthlyGas + floatMonthlyDist)

	// Total Cost
	total := usageCost + distCostxMonths
	// Monthly advanc
	advance := total / floatNMonths

	//Data for result (string)
	uuid := nuuid.String()
	mwhUsage := fmt.Sprintf("%.2f", convUsage)
	cUsage := fmt.Sprintf("%.2f", usageCost)
	cDist := fmt.Sprintf("%.2f", distCost)
	cDistxMonths := fmt.Sprintf("%.2f", distCostxMonths)
	dTotal := fmt.Sprintf("%.2f", total)
	dAdvance := fmt.Sprintf("%.2f", advance)

	t, _ := template.ParseFiles("web/templates/index.html", "web/templates/content_goutput.html")

	result := Result{
		Uuid:             uuid,
		Company:          company,
		ContractDuration: contractDuration,
		UsageRange:       usageRange,
		UsageM3:          m3Usage,
		UsageMWh:         mwhUsage,
		NumberOfMonths:   nMonths,
		CostUsage:        cUsage,
		CostDist:         cDist,
		CostDistxMonths:  cDistxMonths,
		Total:            dTotal,
		Advance:          dAdvance,
	}

	t.ExecuteTemplate(w, "index", result)

}
