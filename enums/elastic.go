package enums

type ElasticQueryType string

const (
	ElasticQueryTypeMatch    ElasticQueryType = "match"
	ElasticQueryTypeTerm     ElasticQueryType = "term"
	ElasticQueryTypeTerms    ElasticQueryType = "terms"
	ElasticQueryTypeRange    ElasticQueryType = "range"
	ElasticQueryTypeExists   ElasticQueryType = "exists"
	ElasticQueryTypePrefix   ElasticQueryType = "prefix"
	ElasticQueryTypeWildcard ElasticQueryType = "wildcard"
)
