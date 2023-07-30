package main

// This code is generated by github.com/parsyl/parquet.

type ValuationDetail struct {
	Valuationdetail                 *float64 `parquet:"valuationdetail"`
	Valuation                       *string  `parquet:"valuation"`
	Position                        *string  `parquet:"position"`
	Posdetail                       *string  `parquet:"posdetail"`
	Valuetime                       *int64   `parquet:"valuetime"`
	Begtime                         *int64   `parquet:"begtime"`
	Endtime                         *int64   `parquet:"endtime"`
	Quantitystatus                  *string  `parquet:"quantitystatus"`
	Description                     *string  `parquet:"description"`
	Quantity                        *float64 `parquet:"quantity"`
	Pricequantity                   *float64 `parquet:"pricequantity"`
	Exposurequantity                *float64 `parquet:"exposurequantity"`
	Price                           *float64 `parquet:"price"`
	Value                           *float64 `parquet:"value"`
	Marketprice                     *float64 `parquet:"marketprice"`
	Marketvalue                     *float64 `parquet:"marketvalue"`
	Unit                            *string  `parquet:"unit"`
	Quantitytype                    *string  `parquet:"quantitytype"`
	Positionstatus                  *string  `parquet:"positionstatus"`
	Evergreenstatus                 *string  `parquet:"evergreenstatus"`
	Settlementstatus                *bool    `parquet:"settlementstatus"`
	Expectedduedate                 *int64   `parquet:"expectedduedate"`
	Exposure                        *string  `parquet:"exposure"`
	Pricestatus                     *string  `parquet:"pricestatus"`
	Exposuretype                    *string  `parquet:"exposuretype"`
	Priceindex                      *string  `parquet:"priceindex"`
	Currency                        *string  `parquet:"currency"`
	Priceunit                       *string  `parquet:"priceunit"`
	Trade                           *string  `parquet:"trade"`
	Npvfactor                       *float64 `parquet:"npvfactor"`
	Validation                      *string  `parquet:"validation"`
	Pricedate                       *int64   `parquet:"pricedate"`
	Strikeprice                     *float64 `parquet:"strikeprice"`
	Optrefprice                     *float64 `parquet:"optrefprice"`
	Intrate                         *float64 `parquet:"intrate"`
	Optionvalue                     *float64 `parquet:"optionvalue"`
	Delta                           *float64 `parquet:"delta"`
	Gamma                           *float64 `parquet:"gamma"`
	Theta                           *float64 `parquet:"theta"`
	Vega                            *float64 `parquet:"vega"`
	Rho                             *float64 `parquet:"rho"`
	Timezone                        *string  `parquet:"timezone"`
	Volatility                      *float64 `parquet:"volatility"`
	Promptvolatility                *float64 `parquet:"promptvolatility"`
	Var                             *float64 `parquet:"var"`
	Product                         *string  `parquet:"product"`
	Fee                             *string  `parquet:"fee"`
	Feetype                         *string  `parquet:"feetype"`
	Shipment                        *string  `parquet:"shipment"`
	Measure                         *string  `parquet:"measure"`
	Tsperiod                        *string  `parquet:"tsperiod"`
	Futuremonth                     *int64   `parquet:"futuremonth"`
	Expirationdate                  *string  `parquet:"expirationdate"`
	Settlementdate                  *int64   `parquet:"settlementdate"`
	Daylightsaving                  *bool    `parquet:"daylightsaving"`
	Marketdayhour                   *float64 `parquet:"marketdayhour"`
	Cycle                           *string  `parquet:"cycle"`
	Taxlocation                     *string  `parquet:"taxlocation"`
	Optionstatus                    *string  `parquet:"optionstatus"`
	Contract                        *string  `parquet:"contract"`
	Compositeindex                  *string  `parquet:"compositeindex"`
	Company                         *string  `parquet:"company"`
	Counterparty                    *string  `parquet:"counterparty"`
	Trader                          *string  `parquet:"trader"`
	Tradebook                       *string  `parquet:"tradebook"`
	Currencyfactor                  *float64 `parquet:"currencyfactor"`
	Paymentterms                    *string  `parquet:"paymentterms"`
	Transactiontype                 *string  `parquet:"transactiontype"`
	Component                       *string  `parquet:"component"`
	Block                           *string  `parquet:"block"`
	Marketarea                      *string  `parquet:"marketarea"`
	Carrier                         *string  `parquet:"carrier"`
	Location                        *string  `parquet:"location"`
	Property                        *string  `parquet:"property"`
	Pile                            *string  `parquet:"pile"`
	Quality                         *string  `parquet:"quality"`
	Tradetype                       *string  `parquet:"tradetype"`
	Tier                            *float64 `parquet:"tier"`
	Strategy                        *string  `parquet:"strategy"`
	Strategydetail                  *string  `parquet:"strategydetail"`
	Hedge                           *string  `parquet:"hedge"`
	Hypothetical                    *bool    `parquet:"hypothetical"`
	Producttype                     *string  `parquet:"producttype"`
	Movement                        *bool    `parquet:"movement"`
	Pricebegtime                    *int64   `parquet:"pricebegtime"`
	Priceendtime                    *int64   `parquet:"priceendtime"`
	Settlementcurrexchrate          *float64 `parquet:"settlementcurrexchrate"`
	Settlementcurrency              *string  `parquet:"settlementcurrency"`
	Cashnpvfactor                   *float64 `parquet:"cashnpvfactor"`
	Fintransact                     *string  `parquet:"fintransact"`
	Fxstatus                        *string  `parquet:"fxstatus"`
	Fxcurrency                      *string  `parquet:"fxcurrency"`
	Fxexposure                      *float64 `parquet:"fxexposure"`
	Fxvalue                         *float64 `parquet:"fxvalue"`
	Fxrisk                          *float64 `parquet:"fxrisk"`
	Fxvolatility                    *float64 `parquet:"fxvolatility"`
	Fxbegtime                       *int64   `parquet:"fxbegtime"`
	Fxendtime                       *int64   `parquet:"fxendtime"`
	Quantityattribution             *string  `parquet:"quantityattribution"`
	Deltaattribution                *string  `parquet:"deltaattribution"`
	Valueattribution                *string  `parquet:"valueattribution"`
	Gammaattribution                *float64 `parquet:"gammaattribution"`
	Marketvalueattribution          *string  `parquet:"marketvalueattribution"`
	Transactionattribution          *string  `parquet:"transactionattribution"`
	Newposition                     *bool    `parquet:"newposition"`
	Nonlinearattribution            *string  `parquet:"nonlinearattribution"`
	Tpseason                        *string  `parquet:"tpseason"`
	Tpyear                          *string  `parquet:"tpyear"`
	Tpquarter                       *string  `parquet:"tpquarter"`
	Tpmonth                         *string  `parquet:"tpmonth"`
	Tpweek                          *string  `parquet:"tpweek"`
	Tpday                           *string  `parquet:"tpday"`
	Greekfactor                     *float64 `parquet:"greekfactor"`
	Exposuremonth                   *int64   `parquet:"exposuremonth"`
	Exposureindex                   *string  `parquet:"exposureindex"`
	Accountname                     *string  `parquet:"accountname"`
	Creationdate                    *int64   `parquet:"creationdate"`
	Revisionname                    *string  `parquet:"revisionname"`
	Creationname                    *string  `parquet:"creationname"`
	Revisiondate                    *string  `parquet:"revisiondate"`
	Amendedattribution              *string  `parquet:"amendedattribution"`
	Currexchrate                    *float64 `parquet:"currexchrate"`
	Currexchratecash                *float64 `parquet:"currexchratecash"`
	Discountedmarketvalue           *float64 `parquet:"discountedmarketvalue"`
	Discountedvalue                 *float64 `parquet:"discountedvalue"`
	Feeattribution                  *string  `parquet:"feeattribution"`
	Fxattribution                   *string  `parquet:"fxattribution"`
	Intransitstatus                 *string  `parquet:"intransitstatus"`
	Irattribution                   *string  `parquet:"irattribution"`
	Losstype                        *string  `parquet:"losstype"`
	Netrate                         *float64 `parquet:"netrate"`
	Rateunit                        *string  `parquet:"rateunit"`
	Thetaattribution                *string  `parquet:"thetaattribution"`
	Vegaattribution                 *string  `parquet:"vegaattribution"`
	Archive                         *string  `parquet:"archive"`
	Deltaexposure                   *float64 `parquet:"deltaexposure"`
	Deltanpvfactor                  *float64 `parquet:"deltanpvfactor"`
	Timeunit                        *string  `parquet:"timeunit"`
	Duration                        *float64 `parquet:"duration"`
	Expcurrency                     *string  `parquet:"expcurrency"`
	Expprice                        *float64 `parquet:"expprice"`
	Exppriceunit                    *string  `parquet:"exppriceunit"`
	Priceexpquantity                *float64 `parquet:"priceexpquantity"`
	Recordtype                      *string  `parquet:"recordtype"`
	Positiontype                    *string  `parquet:"positiontype"`
	Jeragm_destinationtradebook     *string  `parquet:"jeragm_destinationtradebook"`
	Jeragm_internalposition         *bool    `parquet:"jeragm_internalposition"`
	Vessel                          *string  `parquet:"vessel"`
	Jeragm_cargospecification       *string  `parquet:"jeragm_cargospecification"`
	Jeragm_cargostatus              *string  `parquet:"jeragm_cargostatus"`
	Jeragm_flexcategory             *string  `parquet:"jeragm_flexcategory"`
	Jeragm_jeraltcname              *string  `parquet:"jeragm_jeraltcname"`
	Tradestrategy                   *string  `parquet:"tradestrategy"`
	Jeragm_buysell                  *string  `parquet:"jeragm_buysell"`
	Jeragm_cash                     *float64 `parquet:"jeragm_cash"`
	Jeragm_cashexp                  *string  `parquet:"jeragm_cashexp"`
	Jeragm_cashvalue                *float64 `parquet:"jeragm_cashvalue"`
	Jeragm_disccash                 *float64 `parquet:"jeragm_disccash"`
	Jeragm_discountedcashvalue      *float64 `parquet:"jeragm_discountedcashvalue"`
	Jeragm_exposure                 *string  `parquet:"jeragm_exposure"`
	Jeragm_fxstate                  *string  `parquet:"jeragm_fxstate"`
	Jeragm_pricemarketprice         *float64 `parquet:"jeragm_pricemarketprice"`
	Jeragm_pricestate               *string  `parquet:"jeragm_pricestate"`
	Jeragm_product                  *string  `parquet:"jeragm_product"`
	Jeragm_qtyunit                  *string  `parquet:"jeragm_qtyunit"`
	Jeragm_quant                    *float64 `parquet:"jeragm_quant"`
	Jeragm_rate                     *float64 `parquet:"jeragm_rate"`
	Rate                            *float64 `parquet:"rate"`
	Swaptype                        *string  `parquet:"swaptype"`
	Expproduct                      *string  `parquet:"expproduct"`
	Expproducttype                  *string  `parquet:"expproducttype"`
	Expreportingqty                 *string  `parquet:"expreportingqty"`
	Expreportingunit                *string  `parquet:"expreportingunit"`
	Marketareapi                    *string  `parquet:"marketareapi"`
	Jeragm_freightshipment          *string  `parquet:"jeragm_freightshipment"`
	Exchange                        *string  `parquet:"exchange"`
	Jeragm_postingtype              *string  `parquet:"jeragm_postingtype"`
	Positionmode                    *string  `parquet:"positionmode"`
	Jeragm_fullypriced              *bool    `parquet:"jeragm_fullypriced"`
	Jeragm_marktoIntentprice        *float64 `parquet:"Jeragm_marktoIntentprice"`
	Jeragm_marktointentvalue        *float64 `parquet:"Jeragm_marktointentvalue"`
	Jeragm_intransitinventory       *bool    `parquet:"Jeragm_intransitinventory"`
	Jeragm_PosEndDateInFuture       *bool    `parquet:"Jeragm_PosEndDateInFuture"`
	Jeragm_settlementstatus         *bool    `parquet:"Jeragm_settlementstatus"`
	Jeragm_fullypricedpos           *bool    `parquet:"jeragm_fullypricedpos"`
	Jeragm_duedate                  *int64   `parquet:"jeragm_duedate"`
	Container                       *string  `parquet:"container"`
	Containertype                   *string  `parquet:"containertype"`
	Exchangematchid                 *string  `parquet:"exchangematchid"`
	Lossquantity                    *float64 `parquet:"lossquantity"`
	Unitseq                         *float64 `parquet:"unitseq"`
	Expreportingqty2                *string  `parquet:"expreportingqty2"`
	Expreportingqty3                *string  `parquet:"expreportingqty3"`
	Expreportingunit2               *string  `parquet:"expreportingunit2"`
	Expreportingunit3               *string  `parquet:"expreportingunit3"`
	Tl_sourcetrade                  *string  `parquet:"tl_sourcetrade"`
	Jeragm_brand                    *string  `parquet:"jeragm_brand"`
	Jeragm_includeincvbaseprice     *bool    `parquet:"jeragm_includeincvbaseprice"`
	Jeragm_referencecv              *float64 `parquet:"jeragm_referencecv"`
	Incoterms                       *string  `parquet:"incoterms"`
	Jeragm_final                    *bool    `parquet:"jeragm_final"`
	Jeragm_bestavailablecv          *float64 `parquet:"jeragm_bestavailablecv"`
	Jeragm_marktointentaccrualvalue *float64 `parquet:"jeragm_marktointentaccrualvalue"`
	Jeragm_fxconversion             *bool    `parquet:"jeragm_fxconversion"`
	Jeragm_fxcurrencybegtime        *int64   `parquet:"jeragm_fxcurrencybegtime"`
	Jeragm_fxcurrencyendtime        *int64   `parquet:"jeragm_fxcurrencyendtime"`
	Jeragm_bldate                   *int64   `parquet:"jeragm_bldate"`
	Jeragm_matchno                  *string  `parquet:"jeragm_matchno"`
	Bunkertype                      *string  `parquet:"bunkertype"`
	Gammacurve                      *float64 `parquet:"gammacurve"`
	Hedgecurve                      *float64 `parquet:"hedgecurve"`
	Vegacurve                       *float64 `parquet:"vegacurve"`
	Jeragm_fxrate                   *float64 `parquet:"jeragm_fxrate"`
	Jeragm_originaldelta            *string  `parquet:"jeragm_originaldelta"`
	Jeragm_yearquarter              *string  `parquet:"jeragm_yearquarter"`
	Jeragm_deliverymonth            *int64   `parquet:"jeragm_deliverymonth"`
	Jeragm_deliveryquarter          *int64   `parquet:"jeragm_deliveryquarter"`
	Jeragm_deliveryyear             *int64   `parquet:"jeragm_deliveryyear"`
	Jeragm_deltaposition            *float64 `parquet:"jeragm_deltaposition"`
	Jeragm_discexp                  *string  `parquet:"jeragm_discexp"`
	Jeragm_discexp2                 *string  `parquet:"jeragm_discexp2"`
	Jeragm_discexp3                 *string  `parquet:"jeragm_discexp3"`
	Jeragm_expmonth                 *string  `parquet:"jeragm_expmonth"`
	Jeragm_expmonth_coal            *string  `parquet:"jeragm_expmonth_coal"`
	Jeragm_expseasons               *string  `parquet:"jeragm_expseasons"`
	Jeragm_expyears                 *string  `parquet:"jeragm_expyears"`
	Jeragm_expquarters              *string  `parquet:"jeragm_expquarters"`
	Jeragm_quality_                 *string  `parquet:"jeragm_quality_"`
	Jeragm_reppostype               *string  `parquet:"jeragm_reppostype"`
	Jeragm_exposurestatus           *string  `parquet:"jeragm_exposurestatus"`
	Jeragm_exprolloffdate           *int64   `parquet:"jeragm_exprolloffdate"`
	Jeragm_signedpriceqty           *float64 `parquet:"jeragm_signedpriceqty"`
	Rangepricingtype                *string  `parquet:"rangepricingtype"`
	Jeragm_customdisaggkey          *string  `parquet:"jeragm_customdisaggkey"`
	Jeragm_customdisagglabel        *string  `parquet:"jeragm_customdisagglabel"`
	Jeragm_concatenation            *string  `parquet:"jeragm_concatenation"`
	Jeragm_customposref             *string  `parquet:"jeragm_customposref"`
	Jeragm_customtraderef           *string  `parquet:"jeragm_customtraderef"`
	Jeragm_getraderef               *string  `parquet:"jeragm_getraderef"`
	Jeragm_priceindex               *string  `parquet:"jeragm_priceindex"`
	Intrinsicprice                  *string  `parquet:"intrinsicprice"`
	Jeragm_mmvolannual              *string  `parquet:"jeragm_mmvolannual"`
	Jeragm_mmvoltomaturity          *string  `parquet:"jeragm_mmvoltomaturity"`
	Jeragm_nfixings                 *string  `parquet:"jeragm_nfixings"`
	Jeragm_settlementprice          *string  `parquet:"jeragm_settlementprice"`
	Jeragm_settlementtype           *string  `parquet:"jeragm_settlementtype"`
	Jeragm_cashnpvfactororiginal    *string  `parquet:"jeragm_cashnpvfactororiginal"`
	Jeragm_cashnpvfactorusd         *string  `parquet:"jeragm_cashnpvfactorusd"`
	Jeragm_currexchratecashusd      *string  `parquet:"jeragm_currexchratecashusd"`
	Jeragm_currexchratespot         *string  `parquet:"jeragm_currexchratespot"`
	Jeragm_currexchratespotusd      *string  `parquet:"jeragm_currexchratespotusd"`
	Jeragm_timetoexpiry             *string  `parquet:"jeragm_timetoexpiry"`
	Jeragm_fxdiscountdate           *string  `parquet:"jeragm_fxdiscountdate"`
}