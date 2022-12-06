module github.com/gtforge/go-zooz

go 1.19

require (
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

retract v1.3.1 // Contains errors with deserialization

retract v1.4.0 // Mistake in get customer by reference response format

retract v1.8.0 // three_d_secure_attributes.internal / .external / .sca_exemptions should be omited if empty
