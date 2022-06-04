package main

type Options struct {
	Age int
	Sex string
}

type Option func(*Options)

func WithAge(age int) Option {
	return func(o *Options) {
		o.Age = age
	}
}

func WithSex(sex string) Option {
	return func(o *Options) {
		o.Sex = sex
	}
}

func Create(name string, opts ...Option) {
	options := Options{
		Age: 12,
		Sex: "male",
	}
	for _, o := range opts {
		o(&options)
	}

}

func Delete(name string, opts *Options) {

}

func main() {
	Create("Alice")
	Create("Bob", WithAge(22))

	Delete("Alice", nil)
	Delete("Bob", &Options{Age: 22})
}
