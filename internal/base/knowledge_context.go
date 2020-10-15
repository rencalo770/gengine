package base

type KnowledgeContext struct {
	// ruleName - RuleEntity
	RuleEntities map[string]*RuleEntity
	SortRules    []*RuleEntity
}

func NewKnowledgeContext() *KnowledgeContext {
	return &KnowledgeContext{
		RuleEntities: make(map[string]*RuleEntity),
		SortRules:    make([]*RuleEntity, 0),
	}
}

func (k *KnowledgeContext) ClearRules() {
	k.RuleEntities = make(map[string]*RuleEntity)
	k.SortRules = make([]*RuleEntity, 0)
}
