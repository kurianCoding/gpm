## What is this about?

Go is a versatile and verbose language, the learning curve is quite
linear and therefore a lot of frameworks have sprung around it.

What is missing is perhaps a proper package manager. I am trying to 
build a package manager that answers following two  questions,

- will it get all the required files with one command(possibly like
  npm)?
- will it ensure version management with git commit, while ensuring
  space economy?

Obviously once they are downloaded, these packages have to be imported
and used by the go program just like if it were obtained with `go get`.
