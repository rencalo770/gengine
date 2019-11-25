package base

import "gengine/context"

type Initializer interface {
	Initialize(kc *KnowledgeContext, dc *context.DataContext)
}