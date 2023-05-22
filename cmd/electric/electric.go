// logik for electric
package electric

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/webcognito/appKalkulacka/pkg/conversion"

	"github.com/google/uuid"
)

type Result struct {
	Uuid             string
	Company          string
	ContractDuration string
	DistCode         string
	UsageVT          string
	UsageNT          string
	NumberOfMonths   string
	MainBreaker      string
	Phases           string
	CostVT           string
	CostNT           string
	CostMonth        string
	Poze             string
	Total            string
	Advance          string
}

func Einput(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/index.html", "web/templates/content_einput.html")
	t.ExecuteTemplate(w, "index", "")
}

func Eoutput(w http.ResponseWriter, r *http.Request) {
	// Parse Form Values
	company := r.FormValue("company")
	contractDuration := r.FormValue("contractDuration")
	distCode := r.FormValue("distCode")
	usageVT := r.FormValue("usageVT")
	usageNT := r.FormValue("usageNT")
	nMonths := r.FormValue("numberOfMonths")
	priceVT := r.FormValue("priceVT")
	priceNT := r.FormValue("priceNT")
	constPay := r.FormValue("constPay")
	priceInputBreaker := r.FormValue("priceInputBreaker")
	oTE := r.FormValue("OTE")
	byConsumption := r.FormValue("byConsumption")
	byBreaker := r.FormValue("byBreaker")
	mainBreaker := r.FormValue("mainBreaker")
	phases := r.FormValue("Phases")

	// generate Uuid
	nuuid := uuid.New()

	// String to Float
	floatUsageVT := conversion.StringToFloat64(usageVT)
	floatUsageNT := conversion.StringToFloat64(usageNT)
	floatNMonths := conversion.StringToFloat64(nMonths)
	floatPriceVT := conversion.StringToFloat64(priceVT)
	floatPriceNT := conversion.StringToFloat64(priceNT)
	floatConstPay := conversion.StringToFloat64(constPay)
	floatPriceInputBreaker := conversion.StringToFloat64(priceInputBreaker)
	floatOTE := conversion.StringToFloat64(oTE)
	floatByConsumption := conversion.StringToFloat64(byConsumption)
	floatByBreaker := conversion.StringToFloat64(byBreaker)
	floatMainBreaker := conversion.StringToFloat64(mainBreaker)
	floatPhases := conversion.StringToFloat64(phases)

	// Usage Cost
	calVT := floatUsageVT * floatPriceVT
	calNT := floatUsageNT * floatPriceNT

	// Monthly Payments
	calMonthly := floatConstPay + floatPriceInputBreaker + floatOTE

	// POZE
	aPoze := floatByConsumption * (floatUsageNT + floatUsageVT)
	bPoze := floatByBreaker * floatMainBreaker * floatPhases * 12
	calPoze := calPoze(aPoze, bPoze)

	// Total Cost
	total := calVT + calNT + (floatNMonths * calMonthly) + calPoze
	// Monthly advances
	advance := total / floatNMonths

	// Data for result (string)
	uuid := nuuid.String()
	valNT := fmt.Sprintf("%.2f", calNT)
	valVT := fmt.Sprintf("%.2f", calVT)
	valMonthly := fmt.Sprintf("%.2f", calMonthly)
	valPoze := fmt.Sprintf("%.2f", calPoze)
	dTotal := fmt.Sprintf("%.2f", total)
	dAdvance := fmt.Sprintf("%.2f", advance)

	t, _ := template.ParseFiles("web/templates/index.html", "web/templates/content_eoutput.html")

	result := Result{
		Uuid:             uuid,
		Company:          company,
		ContractDuration: contractDuration,
		DistCode:         distCode,
		UsageVT:          usageVT,
		UsageNT:          usageNT,
		NumberOfMonths:   nMonths,
		MainBreaker:      mainBreaker,
		Phases:           phases,
		CostVT:           valVT,
		CostNT:           valNT,
		CostMonth:        valMonthly,
		Poze:             valPoze,
		Total:            dTotal,
		Advance:          dAdvance,
	}

	t.ExecuteTemplate(w, "index", result)
}

func calPoze(aPoze float64, bPoze float64) (poze float64) {
	if aPoze >= bPoze {
		poze = bPoze
	} else {
		poze = aPoze
	}
	return poze
}
