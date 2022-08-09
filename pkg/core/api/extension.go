package api

type ExtensionComponent interface {
	DependencyInterceptors() []Interceptor
}
