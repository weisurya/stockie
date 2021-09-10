package extract

// BalanceSheetGeneral - General Industry - 1210000
type BalanceSheetGeneral struct {
	Assets               Assets               `loc:"B5" name:"Assets"   is_title:"true"`
	LiabilitiesAndEquity LiabilitiesAndEquity `loc:"B124" name:"Liabilities and equity"   is_title:"true"`
}

type Assets struct {
	CurrentAssets    CurrentAssets    `loc:"B6" name:"Current assets"   is_title:"true"`
	NonCurrentAssets NonCurrentAssets `loc:"B57" name:"Non-current assets"   is_title:"true"`
	TotalAssets      int              `loc:"B123" name:"Total assets" `
}

type CurrentAssets struct {
	CashAndCashEquivalent                                int                        `loc:"B7" name:"Cash and cash equivalents" `
	NotesReceivable                                      int                        `loc:"B8" name:"Notes receivable" `
	ShortTermInvestments                                 int                        `loc:"B9" name:"Short-term investments" `
	CurrentRestrictedFunds                               int                        `loc:"B10" name:"Current restricted funds" `
	CurrentFinancialAssets                               CurrentFinancialAssets     `loc:"B11" name:"Current financial assets"   is_title:"true"`
	TradeReceivables                                     TradeReceivables           `loc:"B17" name:"Trade receivables"   is_title:"true"`
	CurrentFinanceLeaseReceivable                        int                        `loc:"B20" name:"Current finance lease receivables" `
	RetentionReceivable                                  RetentionReceivable        `loc:"B21" name:"Retention receivables"   is_title:"true"`
	UnbilledReceivables                                  UnbilledReceivables        `loc:"B24" name:"Unbilled receivables"   is_title:"true"`
	ReceivablesOnSubsidy                                 int                        `loc:"B27" name:"Receivables on subsidy" `
	CurrentCustomerReceivables                           CurrentCustomerReceivables `loc:"B28" name:"Current customer receivables"   is_title:"true"`
	MarginReceivables                                    int                        `loc:"B31" name:"Margin receivables" `
	FromClearingAndSettlement                            int                        `loc:"B32" name:"Receivables from clearing and settlement guarantee institution" `
	PremiumAndReinsuranceReceivable                      int                        `loc:"B33" name:"Premium and reinsurance receivables" `
	DividendsAndInterestReceivables                      int                        `loc:"B34" name:"Dividends and interest receivables" `
	OtherReceivables                                     OtherReceivables           `loc:"B35" name:"Other receivables"   is_title:"true"`
	CurrentInventories                                   CurrentInventories         `loc:"B38" name:"Current inventories"   is_title:"true"`
	CurrentPrepaidExpenses                               int                        `loc:"B42" name:"Current prepaid expenses" `
	Guarantees                                           int                        `loc:"B43" name:"Guarantees" `
	CurrentAdvances                                      CurrentAdvances            `loc:"B44" name:"Current advances"   is_title:"true"`
	CurrentPrepaidTaxes                                  int                        `loc:"B48" name:"Current prepaid taxes" `
	CurrentClaimsForTaxRefund                            int                        `loc:"B49" name:"Current claims for tax refund" `
	CurrentDeferredStrippingCosts                        int                        `loc:"B50" name:"Current deferred stripping costs" `
	CurrentDeferredMobilsationCosts                      int                        `loc:"B51" name:"Current deferred mobilisation costs" `
	CurrentTaxAmnestyAssets                              int                        `loc:"B52" name:"Current tax amnesty assets" `
	OtherCurrentNonFinancialAssets                       int                        `loc:"B53" name:"Other current non-financial assets" `
	DisposalGroupClassifiedAsHeldForSale                 int                        `loc:"B54" name:"Non-current assets or disposal groups classified as held-for-sale" `
	DisposalGroupClassifiedAsHeldForDistributionToOwners int                        `loc:"B55" name:"Non-current assets or disposal groups classified as held-for-distribution to owners" `
	TotalCurrentAssets                                   int                        `loc:"B56" name:"Total current assets" `
}

type CurrentFinancialAssets struct {
	AtFairValueThroughProfitOrLoss int `loc:"B12" name:"Current financial assets at fair value through profit or loss" `
	HeldToMaturityInvestments      int `loc:"B13" name:"Current financial assets held-to-maturity investments" `
	AvailableForSale               int `loc:"B14" name:"Current financial assets available-for-sale" `
	Others                         int `loc:"B15" name:"Other current financial assets" `
	Total                          int
}

type TradeReceivables struct {
	ThirdParties   int `loc:"B18" name:"Trade receivables third parties" `
	RelatedParties int `loc:"B19" name:"Trade receivables related parties" `
	Total          int
}

type RetentionReceivable struct {
	ThirdParties   int `loc:"B22" name:"Retention receivables third parties" `
	RelatedParties int `loc:"B23" name:"Retention receivables related parties" `
	Total          int
}

type UnbilledReceivables struct {
	ThirdParties   int `loc:"B25" name:"Unbilled receivables third parties" `
	RelatedParties int `loc:"B26" name:"Unbilled receivables related parties" `
	Total          int
}

type CurrentCustomerReceivables struct {
	ThirdParties   int `loc:"B29" name:"Current customer receivables third parties" `
	RelatedParties int `loc:"B30" name:"Current customer receivables related parties" `
	Total          int
}

type OtherReceivables struct {
	ThirdParties   int `loc:"B36" name:"Other receivables third parties" `
	RelatedParties int `loc:"B37" name:"Other receivables related parties" `
	Total          int
}

type CurrentInventories struct {
	LivestockInventories int `loc:"B39" name:"Current livestock inventories" `
	RealEstateAssets     int `loc:"B40" name:"Current real estate assets" `
	Inventories          int `loc:"B41" name:"Current inventories" `
	Total                int
}

type CurrentAdvances struct {
	OnInvestments               int `loc:"B45" name:"Current advances on investments" `
	OnPurchaseOfPropertyPlantEq int `loc:"B46" name:"Current advances on purchase of property, plant and equipment" `
	Others                      int `loc:"B47" name:"Other current advances" `
	Total                       int
}

type NonCurrentAssets struct {
	FinanceLeaseReceivables             int                          `loc:"B58" name:"Non-current finance lease receivables" `
	RestrictedFunds                     int                          `loc:"B59" name:"Non-current restricted funds" `
	AircraftMaintenanceReserveFunds     int                          `loc:"B60" name:"Aircraft maintenance reserve funds" `
	ReceivablesFromRelatedParties       int                          `loc:"B61" name:"Receivables from related parties" `
	ReceivablesFromStockholders         int                          `loc:"B62" name:"Receivables from stockholders" `
	NonCurrectCustomerReceivable        NonCurrectCustomerReceivable `loc:"B63" name:"Non-current customer receivables"   is_title:"true"`
	OtherNonCurrentAssets               OtherNonCurrentAssets        `loc:"B66" name:"Other non-current receivables"   is_title:"true"`
	InvestmentsAccountedForUsingEquity  int                          `loc:"B69" name:"Investments accounted for using equity method" `
	InvestmentsInSubsidiaries           InvestmentsInSubsidiaries    `loc:"B70" name:"Investments in subsidiaries, joint ventures and associates"   is_title:"true"`
	NonCurrentAdvances                  NonCurrentAdvances           `loc:"B74" name:"Non-current advances"   is_title:"true"`
	NonCurrentFinancialAssets           NonCurrentFinancialAssets    `loc:"B78" name:"Non-current financial assets"   is_title:"true"`
	NonCurrentDerivativeFinancialAssets int                          `loc:"B83" name:"Non-current derivative financial assets" `
	NonCurrentPrepaidExpenses           int                          `loc:"B84" name:"Non-current prepaid expenses" `
	NonCurrentPrepaidTaxes              int                          `loc:"B85" name:"Non-current prepaid taxes" `
	DeferredTaxAssets                   int                          `loc:"B86" name:"Deferred tax assets" `
	NonCurrentInventories               NonCurrentInventories        `loc:"B87" name:"Non-current inventories"   is_title:"true"`
	LivestockProduction                 int                          `loc:"B91" name:"Livestock production" `
	IndustrialTimberPlantations         IndustrialTimberPlantations  `loc:"B92" name:"Industrial timber plantations"   is_title:"true"`
	PlantationAssets                    PlantationAssets             `loc:"B95" name:"Plantation assets"   is_title:"true"`
	PlasmaPlantations                   int                          `loc:"B98" name:"Plasma plantations" `
	ReinsuranceAssets                   int                          `loc:"B99" name:"Reinsurance assets" `
	InvestmentProperties                int                          `loc:"B100" name:"Investment properties" `
	PropertyPlanAndEquipment            int                          `loc:"B101" name:"Property, plant and equipment" `
	IjarahAssets                        int                          `loc:"B102" name:"Ijarah assets" `
	ForeclosedAssets                    int                          `loc:"B103" name:"Foreclosed assets" `
	OilAndGasAssets                     int                          `loc:"B104" name:"Oil and gas assets" `
	ExplorationAndEvaluationAssets      int                          `loc:"B105" name:"Exploration and evaluation assets" `
	TollRoadConcessionRights            int                          `loc:"B106" name:"Toll road concession rights" `
	MiningProperties                    int                          `loc:"B107" name:"Mining properties" `
	NonCurrentDeferredStrippingCosts    int                          `loc:"B108" name:"Non-current deferred stripping costs" `
	NonCurrentDeferredMobilisationCosts int                          `loc:"B109" name:"Non-current deferred mobilisation costs" `
	DeferredCharges                     DeferredCharges              `loc:"B110" name:"Deferred charges"   is_title:"true"`
	NonCurrentClaimsForTaxRefund        int                          `loc:"B116" name:"Non-current claims for tax refund" `
	PostEmploymentBenefitAsets          int                          `loc:"B117" name:"Post-employment benefit assets" `
	Goodwill                            int                          `loc:"B118" name:"Goodwill" `
	IntangibleAssetsOtherThanGoodwill   int                          `loc:"B119" name:"Intangible assets other than goodwill" `
	NonCurrentTaxAMnestyAssets          int                          `loc:"B120" name:"Non-current tax amnesty assets" `
	OtherNonCurrentNonFinancialAssets   int                          `loc:"B121" name:"Other non-current non-financial assets" `
	TotalNonCurrentAssets               int                          `loc:"B122" name:"Total non-current assets" `
}

type NonCurrectCustomerReceivable struct {
	ThirdParties   int `loc:"B64" name:"Non-current customer receivables" `
	RelatedParties int `loc:"B65" name:"Non-current customer receivables related parties" `
	Total          int
}

type OtherNonCurrentAssets struct {
	ThirdParties   int `loc:"B67" name:"Other non-current receivables third parties" `
	Relatedparties int `loc:"B68" name:"Other non-current receivables related parties" `
	Total          int
}

type InvestmentsInSubsidiaries struct {
	InSubsidiaries  int `loc:"B71" name:"Investments in subsidiaries" `
	InJointVentures int `loc:"B72" name:"Investments in joint ventures" `
	InAssociates    int `loc:"B73" name:"Investments in associates" `
	Total           int
}

type NonCurrentAdvances struct {
	OnInvestments                      int `loc:"B75" name:"Non-current advances on investments" `
	OnPurchasesOfPropertyPlanEquipment int `loc:"B76" name:"Non-current advances on purchase of property, plant and equipment" `
	Others                             int `loc:"B77" name:"Other non-current advances" `
	Total                              int
}

type NonCurrentFinancialAssets struct {
	AtFairValueThroughProfitOrLoss int `loc:"B79" name:"Non-current financial assets at fair value through profit or loss" `
	HeldToMaturity                 int `loc:"B80" name:"Non-current financial assets held-to-maturity" `
	AvailableForSale               int `loc:"B81" name:"Non-current financial assets available-for-sale" `
	Others                         int `loc:"B82" name:"Other non-current financial assets" `
	Total                          int
}

type NonCurrentInventories struct {
	LivestockInventories int `loc:"B88" name:"Non-current livestock inventories" `
	RealEstateAssets     int `loc:"B89" name:"Non-current real estate assets" `
	Inventories          int `loc:"B90" name:"Non-current inventories" `
	Total                int
}

type IndustrialTimberPlantations struct {
	Mature   int `loc:"B93" name:"Industrial timber plantations mature" `
	Immature int `loc:"B94" name:"Industrial timber plantations immature" `
	Total    int
}

type PlantationAssets struct {
	Mature   int `loc:"B96" name:"Plantation assets mature" `
	Immature int `loc:"B97" name:"Plantation assets immature" `
	Total    int
}

type DeferredCharges struct {
	OnLangrights                              int `loc:"B111" name:"Deferred charges on landrights" `
	OnExplorationAndDevelopment               int `loc:"B112" name:"Deferred charges on exploration and development expenditures" `
	OnCostOnForest                            int `loc:"B113" name:"Deferred charges on cost on forest" `
	OnEnvironmentalAndReclamationExpenditures int `loc:"B114" name:"Deferred charges on environmental and reclamation expenditures" `
	Others                                    int `loc:"B115" name:"Other deferred charges" `
	Total                                     int
}

type LiabilitiesAndEquity struct {
	Liabilities               Liabilities `loc:"B125" name:"Liabilities"   is_title:"true"`
	Equity                    Equity      `loc:"B233" name:"Equity"   is_title:"true"`
	TotalLiabilitiesAndEquity int         `loc:"B257" name:"Total liabilities and equity" `
}

type Liabilities struct {
	CurrentLiabilities    CurrentLiabilities    `loc:"B126" name:"Current liabilities"   is_title:"true"`
	NonCurrentLiabilities NonCurrentLiabilities `loc:"B189" name:"Non-current liabilities"   is_title:"true"`
	TotalLiabilities      int                   `loc:"B232" name:"Total liabilities" `
}

type CurrentLiabilities struct {
	ShortTermLoans                                     int                                    `loc:"B127" name:"Short-term loans" `
	TrustReceiptsPayables                              int                                    `loc:"B128" name:"Trust receipts payables" `
	TradePayables                                      TradePayables                          `loc:"B129" name:"Trade payables"   is_title:"true"`
	OtherPayables                                      OtherPayables                          `loc:"B132" name:"Other payables"   is_title:"true"`
	CurrentAdvancesFromCustomers                       CurrentAdvancesFromCustomers           `loc:"B135" name:"Current advances from customers"   is_title:"true"`
	DividentsPayable                                   int                                    `loc:"B138" name:"Dividends payable" `
	OtherCurrentFinancialLiabilities                   int                                    `loc:"B139" name:"Other current financial liabilities" `
	CurrentAccruedExpenses                             int                                    `loc:"B140" name:"Current accrued expenses" `
	ShortTermPostEmploymentBenefitObligations          int                                    `loc:"B141" name:"Short-term post-employment benefit obligations" `
	TaxesPayable                                       int                                    `loc:"B142" name:"Short-term post-employment benefit obligations" `
	ExcisePayable                                      int                                    `loc:"B143" name:"Excise payable" `
	ProjectPayable                                     int                                    `loc:"B144" name:"Project payables" `
	PayableToClearingAndSettlementGuaranteeInstitution int                                    `loc:"B145" name:"Payables to clearing and settlement guarantee institution" `
	PayablesToCustomers                                PayablesToCustomers                    `loc:"B146" name:"Payables to customers"   is_title:"true"`
	ReinsurancePayables                                int                                    `loc:"B149" name:"Reinsurance payables" `
	FactoringPayables                                  int                                    `loc:"B150" name:"Factoring payables" `
	CurrentDeposits                                    int                                    `loc:"B151" name:"Current deposits" `
	CurrentUnearnedRevenue                             int                                    `loc:"B152" name:"Current unearned revenue" `
	DueToCustomers                                     DueToCustomers                         `loc:"B153" name:"Due to customers"   is_title:"true"`
	CurrentDeferredRevenue                             int                                    `loc:"B156" name:"Current deferred revenue" `
	CurrentProvisions                                  CurrentProvisions                      `loc:"B157" name:"Current provisions"   is_title:"true"`
	ShareBasedPaymentLiabilities                       int                                    `loc:"B164" name:"Share-based payment liabilities" `
	CurrentMaturitiesOfLongTermLiabilities             CurrentMaturitiesOfLongTermLiabilities `loc:"B165" name:"Current maturities of long-term liabilities"   is_title:"true"`
	CurrentDueToRelatedParties                         int                                    `loc:"B182" name:"Current due to related parties" `
	CurrentDueToStakeholders                           int                                    `loc:"B183" name:"Current due to stockholders" `
	ShortTermDerivativeFinancialLiabilities            int                                    `loc:"B184" name:"Short-term derivative financial liabilities" `
	CurrentTaxAmnestyLiabilities                       int                                    `loc:"B185" name:"Current tax amnesty liabilities" `
	OtherCurrentNonFinancialLiabilities                int                                    `loc:"B186" name:"Other current non-financial liabilities" `
	LiabilitiesDirectlyAssociatedWithNonCurrentAssets  int                                    `loc:"B187" name:"Liabilities directly associated with non-current assets or disposal groups classified as held-for-sale or as held-for-distribution to owners" `
	TotalCurrentLiabilities                            int                                    `loc:"B188" name:"Total current liabilities" `
}

type TradePayables struct {
	ThirdParties   int `loc:"B130" name:"Trade payables third parties" `
	RelatedParties int `loc:"B131" name:"Trade payables related parties" `
	Total          int
}

type OtherPayables struct {
	ThirdParties   int `loc:"B133" name:"Other payables third parties" `
	RelatedParties int `loc:"B134" name:"Other payables related parties" `
	Total          int
}

type CurrentAdvancesFromCustomers struct {
	ThirdParties   int `loc:"B136" name:"Current advances from customers third parties" `
	RelatedParties int `loc:"B137" name:"Current advances from customers related parties" `
	Total          int
}

type PayablesToCustomers struct {
	ThirdParties   int `loc:"B147" name:"Payables to customers third parties" `
	RelatedParties int `loc:"B148" name:"Payables to customers related parties" `
	Total          int
}

type DueToCustomers struct {
	ThirdParties   int `loc:"B154" name:"Due to customers third parties" `
	RelatedParties int `loc:"B155" name:"Due to customers related parties" `
	Total          int
}

type CurrentProvisions struct {
	ForOverlay                                            int `loc:"B158" name:"Current provisions for overlay" `
	ForAircraftReturnAndMaintenanceFunds                  int `loc:"B159" name:"Current provisions for aircraft return and maintenance funds" `
	ForInfrastructureDevelopmentPublicAndSocialFacilities int `loc:"B160" name:"Current provisions for infrastructure development, public and social facilities" `
	ForAssetDismantlingCosts                              int `loc:"B161" name:"Current provisions for asset dismantling costs" `
	ForRestorationAndRehabilitation                       int `loc:"B162" name:"Current provisions for restoration and rehabilitation" `
	Others                                                int `loc:"B163" name:"Other current provisions" `
	Total                                                 int
}

type CurrentMaturitiesOfLongTermLiabilities struct {
	BankLoans                    int `loc:"B166" name:"Current maturities of bank loans" `
	SecuredLoans                 int `loc:"B167" name:"Current maturities of secured loans" `
	UnsecuredLoans               int `loc:"B168" name:"Current maturities of unsecured loans" `
	StepLoans                    int `loc:"B169" name:"Current maturities of step loans" `
	LoansFromGovernmentIndonesia int `loc:"B170" name:"Current maturities of loans from government of the republic of indonesia" `
	SubordinatedLoans            int `loc:"B171" name:"Current maturities of subordinated loans" `
	JointOperationsLiabilities   int `loc:"B172" name:"Current maturities of joint operations liabilities" `
	LandAcquisitionLiabilities   int `loc:"B173" name:"Current maturities of land acquisition liabilities" `
	ConsumerFinancingPayables    int `loc:"B174" name:"Current maturities of consumer financing payables" `
	FinanceLeaseLiabilities      int `loc:"B175" name:"Current maturities of finance lease liabilities" `
	ElectricityPurchasePayables  int `loc:"B176" name:"Current maturities of electricity purchase payables" `
	RetentionPayables            int `loc:"B177" name:"Current maturities of retention payables" `
	NotesPayable                 int `loc:"B178" name:"Current maturities of notes payable" `
	BondsPayable                 int `loc:"B179" name:"Current maturities of bonds payable" `
	Sukuk                        int `loc:"B180" name:"Current maturities of sukuk" `
	OtherBorrowings              int `loc:"B181" name:"Current maturities of other borrowings" `
	Total                        int
}

type NonCurrentLiabilities struct {
	LongTermDerivativeFinancialLiabilities    int                                       `loc:"B190" name:"Long-term derivative financial liabilities" `
	DeferredTaxLiabilities                    int                                       `loc:"B191" name:"Deferred tax liabilities" `
	NonCurrentDueToRelatedParties             int                                       `loc:"B192" name:"Non-current due to related parties" `
	NonCurrentDueToStockHolders               int                                       `loc:"B193" name:"Non-current due to stockholders" `
	LongTermLiabilitiesNetOfCurrentMaturities LongTermLiabilitiesNetOfCurrentMaturities `loc:"B194" name:"Long-term liabilities net of current maturities"   is_title:"true"`
	ConvertibleBonds                          int                                       `loc:"B211" name:"Convertible bonds" `
	NonCurrentUnearnedRevenue                 int                                       `loc:"B212" name:"Non-current unearned revenue" `
	NonCurrentDeposit                         int                                       `loc:"B213" name:"Non-current deposits" `
	NonCurrentAdvancesFromCustomers           NonCurrentAdvancesFromCustomers           `loc:"B214" name:"Non-current advances from customers"   is_title:"true"`
	NonCurrentDeferredRevenue                 int                                       `loc:"B217" name:"Non-current deferred revenue" `
	NonCurrentProvisions                      NonCurrentProvisions                      `loc:"B218" name:"Non-current provisions"   is_title:"true"`
	AccruedStrippingCosts                     int                                       `loc:"B225" name:"Accrued stripping costs" `
	LiabilitiesToPolicyholders                int                                       `loc:"B226" name:"Liabilities to policyholders" `
	LongTermPostEmploymentBenefitObligations  int                                       `loc:"B227" name:"Long-term post-employment benefit obligations" `
	NonCurrentTaxAmnestyLiabilities           int                                       `loc:"B228" name:"Non-current tax amnesty liabilities" `
	OtherNonCurrentFinancialLiabilities       int                                       `loc:"B229" name:"Other non-current financial liabilities" `
	OtherNonCurrentNonFinancialLiabilities    int                                       `loc:"B230" name:"Other non-current non-financial liabilities" `
	TotalNonCurrentLiabilities                int                                       `loc:"B231" name:"Total non-current liabilities" `
}

type LongTermLiabilitiesNetOfCurrentMaturities struct {
	BankLoans                   int `loc:"B195" name:"Long-term bank loans" `
	StepLoans                   int `loc:"B196" name:"Long-term step loans" `
	SecuredLoans                int `loc:"B197" name:"Long-term secured loans" `
	UnsecuredLoans              int `loc:"B198" name:"Long-term unsecured loans" `
	FromGovernment              int `loc:"B199" name:"Long-term loans from government of the republic of indonesia" `
	SubordinatedLoans           int `loc:"B200" name:"Long-term subordinated loans" `
	JointOperationsLiabilities  int `loc:"B201" name:"Long-term joint operations liabilities" `
	LandAcquisitionLiabilities  int `loc:"B202" name:"Long-term land acquisition liabilities" `
	ConsumerFinancingPayabels   int `loc:"B203" name:"Long-term consumer financing payables" `
	LeaseLiabilities            int `loc:"B204" name:"Long-term finance lease liabilities" `
	ElectricityPurchasePayables int `loc:"B205" name:"Long-term electricity purchase payables" `
	RetentionPayables           int `loc:"B206" name:"Long-term retention payables" `
	NotesPayable                int `loc:"B207" name:"Long-term notes payable" `
	BondPayable                 int `loc:"B208" name:"Long-term bonds payable" `
	Sukuk                       int `loc:"B209" name:"Long-term sukuk" `
	Others                      int `loc:"B210" name:"Long-term other borrowings" `
	Total                       int
}

type NonCurrentAdvancesFromCustomers struct {
	ThirdParties   int `loc:"B215" name:"Non-current advances from customers third parties" `
	RelatedParties int `loc:"B216" name:"Non-current advances from customers related parties" `
	Total          int
}

type NonCurrentProvisions struct {
	ForOverlay                           int `loc:"B219" name:"Non-current provisions for overlay" `
	ForAircraftReturnAndMaintenanceFunds int `loc:"B220" name:"Non-current provisions for aircraft return and maintenance funds" `
	ForInfrastructureDevelopment         int `loc:"B221" name:"Non-current provisions for infrastructure development, public and social facilities" `
	ForAssetDismantlingCosts             int `loc:"B222" name:"Non-current provisions for asset dismantling costs" `
	ForRestorationAndRehabilitation      int `loc:"B223" name:"Non-current provisions for restoration and rehabilitation" `
	Others                               int `loc:"B224" name:"Other non-current provisions" `
	Total                                int
}

type Equity struct {
	EquityAttributeToEquityOwnersOfParentEntity EquityAttributeToEquityOwnersOfParentEntity `loc:"B234" name:"Equity attributable to equity owners of parent entity"   is_title:"true"`
	ProformaEquity                              int                                         `loc:"B254" name:"Proforma equity" `
	NonControllingInterests                     int                                         `loc:"B255" name:"Non-controlling interests" `
	TotalEquity                                 int                                         `loc:"B256" name:"Total equity" `
}

type EquityAttributeToEquityOwnersOfParentEntity struct {
	CommonStocks                                                  int              `loc:"B235" name:"Common stocks" `
	PreferredStock                                                int              `loc:"B236" name:"Preferred stocks" `
	AdditionalPaidInCapital                                       int              `loc:"B237" name:"Additional paid-in capital" `
	TresuryStocks                                                 int              `loc:"B238" name:"Treasury stocks" `
	AdvancesInCapitalStock                                        int              `loc:"B239" name:"Advances in capital stock" `
	StockOptions                                                  int              `loc:"B240" name:"Stock options" `
	RevaluationReserves                                           int              `loc:"B241" name:"Revaluation reserves" `
	ReserveOfExchangDifferencesOnTranslation                      int              `loc:"B242" name:"Reserve of exchange differences on translation" `
	ReserveForChangesInFairValueOfAvailableForSaleFinancialAssets int              `loc:"B243" name:"Reserve for changes in fair value of available-for-sale financial assets" `
	ReserveForGainLossesFromInvestments                           int              `loc:"B244" name:"Reserve of gains (losses) from investments in equity instruments" `
	ReserveOfShareBasedPayments                                   int              `loc:"B245" name:"Reserve of share-based payments" `
	ReserveOfCashFLowHedges                                       int              `loc:"B246" name:"Reserve of cash flow hedges" `
	ReserveOfRemeasurementsOfDefinedBenefitPlans                  int              `loc:"B247" name:"Reserve of remeasurements of defined benefit plans" `
	OtherReserves                                                 int              `loc:"B248" name:"Other reserves" `
	OtherComponentsOfEquity                                       int              `loc:"B249" name:"Other components of equity" `
	RetainedEarnings                                              RetainedEarnings `loc:"B250" name:"Retained earnings (deficit)"   is_title:"true"`
	TotalEquityAttributableToEquityOwnersOfParentEntity           int              `loc:"B253" name:"Total equity attributable to equity owners of parent entity" `
}

type RetainedEarnings struct {
	AppropriateRetainedEarnings    int `loc:"B251" name:"Appropriated retained earnings" `
	UnappropriatedRetainedEarnings int `loc:"B252" name:"Unappropriated retained earnings" `
	Total                          int
}

func (r statementReader) readBalanceSheetTypeOfGeneral() (res BalanceSheetGeneral, err error) {
	var currentFinancialAssets CurrentFinancialAssets
	if err = r.setValueToSubjectFromSheet(&currentFinancialAssets, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var tradeReceivables TradeReceivables
	if err = r.setValueToSubjectFromSheet(&tradeReceivables, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var retentionReceivable RetentionReceivable
	if err = r.setValueToSubjectFromSheet(&retentionReceivable, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var unbilledReceivables UnbilledReceivables
	if err = r.setValueToSubjectFromSheet(&unbilledReceivables, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var currentCustomerReceivables CurrentCustomerReceivables
	if err = r.setValueToSubjectFromSheet(&currentCustomerReceivables, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var otherReceivables OtherReceivables
	if err = r.setValueToSubjectFromSheet(&otherReceivables, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var currentInventories CurrentInventories
	if err = r.setValueToSubjectFromSheet(&currentInventories, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var currentAdvances CurrentAdvances
	if err = r.setValueToSubjectFromSheet(&currentAdvances, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var currentAssets CurrentAssets
	currentAssets.CurrentFinancialAssets = currentFinancialAssets
	currentAssets.TradeReceivables = tradeReceivables
	currentAssets.RetentionReceivable = retentionReceivable
	currentAssets.UnbilledReceivables = unbilledReceivables
	currentAssets.CurrentCustomerReceivables = currentCustomerReceivables
	currentAssets.OtherReceivables = otherReceivables
	currentAssets.CurrentInventories = currentInventories
	currentAssets.CurrentAdvances = currentAdvances

	if err = r.setValueToSubjectFromSheet(&currentAssets, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var nonCurrectCustomerReceivable NonCurrectCustomerReceivable
	if err = r.setValueToSubjectFromSheet(&nonCurrectCustomerReceivable, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var otherNonCurrentAssets OtherNonCurrentAssets
	if err = r.setValueToSubjectFromSheet(&otherNonCurrentAssets, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var investmentsInSubsidiaries InvestmentsInSubsidiaries
	if err = r.setValueToSubjectFromSheet(&investmentsInSubsidiaries, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var nonCurrentAdvances NonCurrentAdvances
	if err = r.setValueToSubjectFromSheet(&nonCurrentAdvances, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var nonCurrentFinancialAssets NonCurrentFinancialAssets
	if err = r.setValueToSubjectFromSheet(&nonCurrentFinancialAssets, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var nonCurrentInventories NonCurrentInventories
	if err = r.setValueToSubjectFromSheet(&nonCurrentInventories, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var industrialTimberPlantations IndustrialTimberPlantations
	if err = r.setValueToSubjectFromSheet(&industrialTimberPlantations, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var plantationAssets PlantationAssets
	if err = r.setValueToSubjectFromSheet(&plantationAssets, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var deferredCharges DeferredCharges
	if err = r.setValueToSubjectFromSheet(&deferredCharges, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var nonCurrentAssets NonCurrentAssets
	nonCurrentAssets.NonCurrectCustomerReceivable = nonCurrectCustomerReceivable
	nonCurrentAssets.OtherNonCurrentAssets = otherNonCurrentAssets
	nonCurrentAssets.InvestmentsInSubsidiaries = investmentsInSubsidiaries
	nonCurrentAssets.NonCurrentAdvances = nonCurrentAdvances
	nonCurrentAssets.NonCurrentFinancialAssets = nonCurrentFinancialAssets
	nonCurrentAssets.NonCurrentInventories = nonCurrentInventories
	nonCurrentAssets.IndustrialTimberPlantations = industrialTimberPlantations
	nonCurrentAssets.PlantationAssets = plantationAssets
	nonCurrentAssets.DeferredCharges = deferredCharges

	if err = r.setValueToSubjectFromSheet(&nonCurrentAssets, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var assets Assets
	assets.CurrentAssets = currentAssets
	assets.NonCurrentAssets = nonCurrentAssets

	if err = r.setValueToSubjectFromSheet(&assets, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var tradePayables TradePayables

	if err = r.setValueToSubjectFromSheet(&tradePayables, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var otherPayables OtherPayables

	if err = r.setValueToSubjectFromSheet(&otherPayables, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var currentAdvancesFromCustomers CurrentAdvancesFromCustomers

	if err = r.setValueToSubjectFromSheet(&currentAdvancesFromCustomers, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var payablesToCustomers PayablesToCustomers

	if err = r.setValueToSubjectFromSheet(&payablesToCustomers, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var dueToCustomers DueToCustomers

	if err = r.setValueToSubjectFromSheet(&dueToCustomers, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var currentProvisions CurrentProvisions

	if err = r.setValueToSubjectFromSheet(&currentProvisions, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var currentMaturitiesOfLongTermLiabilities CurrentMaturitiesOfLongTermLiabilities

	if err = r.setValueToSubjectFromSheet(&currentMaturitiesOfLongTermLiabilities, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var currentLiabilities CurrentLiabilities
	currentLiabilities.TradePayables = tradePayables
	currentLiabilities.OtherPayables = otherPayables
	currentLiabilities.CurrentAdvancesFromCustomers = currentAdvancesFromCustomers
	currentLiabilities.PayablesToCustomers = payablesToCustomers
	currentLiabilities.DueToCustomers = dueToCustomers
	currentLiabilities.CurrentProvisions = currentProvisions
	currentLiabilities.CurrentMaturitiesOfLongTermLiabilities = currentMaturitiesOfLongTermLiabilities

	if err = r.setValueToSubjectFromSheet(&currentLiabilities, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var longTermLiabilitiesNetOfCurrentMaturities LongTermLiabilitiesNetOfCurrentMaturities

	if err = r.setValueToSubjectFromSheet(&longTermLiabilitiesNetOfCurrentMaturities, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var nonCurrentAdvancesFromCustomers NonCurrentAdvancesFromCustomers

	if err = r.setValueToSubjectFromSheet(&nonCurrentAdvancesFromCustomers, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var nonCurrentProvisions NonCurrentProvisions

	if err = r.setValueToSubjectFromSheet(&nonCurrentProvisions, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var nonCurrentLiabilities NonCurrentLiabilities
	nonCurrentLiabilities.LongTermLiabilitiesNetOfCurrentMaturities = longTermLiabilitiesNetOfCurrentMaturities
	nonCurrentLiabilities.NonCurrentAdvancesFromCustomers = nonCurrentAdvancesFromCustomers
	nonCurrentLiabilities.NonCurrentProvisions = nonCurrentProvisions

	if err = r.setValueToSubjectFromSheet(&nonCurrentLiabilities, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var liabilities Liabilities
	liabilities.CurrentLiabilities = currentLiabilities
	liabilities.NonCurrentLiabilities = nonCurrentLiabilities

	if err = r.setValueToSubjectFromSheet(&liabilities, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var retainedEarnings RetainedEarnings

	if err = r.setValueToSubjectFromSheet(&retainedEarnings, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var equityAttributeToEquityOwnersOfParentEntity EquityAttributeToEquityOwnersOfParentEntity
	equityAttributeToEquityOwnersOfParentEntity.RetainedEarnings = retainedEarnings

	if err = r.setValueToSubjectFromSheet(&equityAttributeToEquityOwnersOfParentEntity, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var equity Equity
	equity.EquityAttributeToEquityOwnersOfParentEntity = equityAttributeToEquityOwnersOfParentEntity

	if err = r.setValueToSubjectFromSheet(&equity, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var liabilitiesAndEquity LiabilitiesAndEquity

	if err = r.setValueToSubjectFromSheet(&liabilitiesAndEquity, SheetCodeGeneralIndustry); err != nil {
		return
	}

	var balanceSheetGeneral BalanceSheetGeneral
	balanceSheetGeneral.Assets = assets
	balanceSheetGeneral.LiabilitiesAndEquity = liabilitiesAndEquity

	return balanceSheetGeneral, nil
}

// BalanceSheetGeneral - General Industry - 1220000
type BalanceSheetGeneral2 struct {
	Assets               Assets2              `loc:"B5" name:"Assets"   is_title:"true"`
	LiabilitiesAndEquity LiabilitiesAndEquity `loc:"B90" name:"Liabilities and equity"   is_title:"true"`
}

type Assets2 struct {
	CashAndCashEquivalent     int                  `loc:"B6" name:"Cash and cash equivalents"`
	NotesReceivable           int                  `loc:"B7" name:"Notes receivable"`
	ShortTermInvestments      int                  `loc:"B8" name:"Short-term investments"`
	FinancialAssets           FinancialAssets      `loc:"B9" name:"Financial assets" is_title:"true"`
	DerivativeFinancialAssets int                  `loc:"B14" name:"Derivative financial assets"`
	TradeReceivables          TradeReceivables2    `loc:"B15" name:"Trade receivables"`
	RetentionReceivable       RetentionReceivable2 `loc:"B18" name:"Retention receivables"   is_title:"true"`
}

type FinancialAssets struct {
	AtFairValueThroughProfitOrLoss int `loc:"B10" name:"Financial assets at fair value through profit or loss" `
	HeldToMaturity                 int `loc:"B11" name:"Financial assets held-to-maturity" `
	AvailableForSale               int `loc:"B12" name:"Financial assets available-for-sale" `
	Others                         int `loc:"B13" name:"Other financial assets" `
	Total                          int
}

type TradeReceivables2 struct {
	ThirdParties   int `loc:"B16" name:"Trade receivables third parties" `
	RelatedParties int `loc:"B17" name:"Trade receivables related parties" `
	Total          int
}

type RetentionReceivable2 struct {
	ThirdParties   int `loc:"B19" name:"Retention receivables third parties" `
	RelatedParties int `loc:"B20" name:"Retention receivables related parties" `
	Total          int
}
