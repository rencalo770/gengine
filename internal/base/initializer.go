package base

import "gengine/context"

type Initializer interface {
	Initialize(dc *context.DataContext)
}