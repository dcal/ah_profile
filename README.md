##Bugs
In some cases, the code generator will set a defined attribute defined as an underscore to a camel-case within a function (valid_until):

```
func (c *Client) CreateRecords(path string, email string, phone string, valid_until string) (*http.Response, error) {
	var body io.Reader
	u :# url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	values :# u.Query()
	values.Set("email", email)
	values.Set("phone", phone)
	values.Set("valid_until", validUntil)
---
}
```

To avoid this, make your attributes camelCase, go style, in the attribute definitions, not like you would make them in the database.

##Observations
###goagen
Goagen expects to be executed from the top-level src directory (i.e. src/github.com/alliancehealth.....)
I found myself generating main and controllers a single time, and regenerating app and client frequently.

###API DSL
Have to admit, after using it for a while, I like it. It's straightforward and clean.

###Database
There is a need for a way to convert context payloads to the db models that allows null values.
Likewise there is a need to deal with null values

I found the best way to do this is one action / resource at a time, and not to try and define the entire DSL in one go.
This is mainly because it's generating code and it won't work if it's not right. It's not just attempting to describe an API semantically.

###controllers
I initally tried moving controllers actions into their own package, but decided to let the framework guide me and left them in the main package. I might move them since they work now.

##Pros
###Fast
I got a dummy service stubbed out in the DSL and returning canned responses in about 10 minutes.

###Web Contexts
Requests implement the go web context interface, so we can send signals across microservices --- build in concurrency

###Middleware
Lots of built-in middlewares

###Validations
Validations on request params/payloads are super easy and out of the box.

###Swagger
We get this for free

###OpenSource!!!
This project has a ton of work being done on it...we could be in and contributing to open source at the ground level

##Cons
###Verbosity
Generated code makes as few assumptions as possible about how you will use your app. It is not meant to be "dry" it is meant to be expendible.
In the case of goa, each request has a pipeline, and becuase it's statically typed, it's difficult to share functions across that pipeline out of the box.

###Data layer
Unless we want to use an ORM, we would probably want to explore creating a lightweight tool for converting Media to our models and then generating sql queries dynamically.

###
