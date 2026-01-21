# Pets

Params Types:

- <a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go">petproject30</a>.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#NewPetParam">NewPetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go">petproject30</a>.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#NewPet">NewPet</a>
- <a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go">petproject30</a>.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#Pet">Pet</a>

Methods:

- <code title="post /pets">client.Pets.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go">petproject30</a>.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#PetNewParams">PetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go">petproject30</a>.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pets/{id}">client.Pets.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go">petproject30</a>.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pets">client.Pets.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#PetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go">petproject30</a>.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#PetListParams">PetListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go">petproject30</a>.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pets/{id}">client.Pets.<a href="https://pkg.go.dev/github.com/miriambudayr/pet-project-30-go#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
