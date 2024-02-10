package code

type Builder struct {
	cat Cat
}

func NewBuilder() *Builder {
	return &Builder{
		cat: Cat{},
	}
}

func (builder *Builder) Name(name string) *Builder {
	builder.cat.name = name
	return builder
}

func (builder *Builder) Age(age uint8) *Builder {
	builder.cat.age = age
	return builder
}

func (builder *Builder) Height(height uint8) *Builder {
	builder.cat.height = height
	return builder
}

func (builder *Builder) Build() *Cat {
	return &builder.cat
}
