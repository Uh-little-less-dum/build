package config_loc_strategies

type ConfigLocationStrategyId int

const (
	WaitForUserBeforeBuild ConfigLocationStrategyId = iota
	CopyFromLocation
)

type ConfigLocationStrategy struct {
	Id    ConfigLocationStrategyId
	Value string
}
