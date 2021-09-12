### Swagger
Swagger is a set of tools for implementation of the OpenAPI specification which was previously known as the Swagger Specification. OpenAPI is a specification for machine-readable interface files for describing, producing, consuming, and visualizing RESTful web services.

Swagger and OpenAPI are two separate independent communities. OpenAPI 3.1 is the latest version of the specification. 

There are other tools that support OpenAPI 3.0 specification but Swagger tools, which are supported by SmartBear Software, are among the most popular tools for implementing the OpenAPI Specification and the swagger suit include Swagger Editor, Swagger UI, SwaggerHub, etc.

### Simple-Server Implementation
Following the [goswagger official tutorial](https://goswagger.io/tutorial/todo-list.html), the simple-server (todo-list) is implemented as shown below.

I created a simple-server directory and inside added a swagger.yml file that can produce a complete skeleton for my project including endpoints. I validated the file for OpenAPI specification as following;
![step-1](https://i.imgur.com/PQB1kcj.png)

Then I generated the structure of the todo-list project with command: <i>swagger generate server -A todo-list -f ./swagger.yml</i>
It prompts to add the following packages;
![](https://i.imgur.com/hD8SH9x.png)

The packages are added using <i>go get</i> command;

![](https://i.imgur.com/fLmIDoN.png)

The following is the produced structure;

![](https://i.imgur.com/64eRIUl.png)

I edited the configure_todo_list.go file as per the tutorial and also main.go file to cope with tls-cert issue. Successfully generated the final outputs similar to the tutorial results;

![](https://i.imgur.com/Q3BNdqS.png)

![](https://i.imgur.com/3XfkoQ9.png)

### Custom-Server Implementation

Custom-Server is implemented following the official [goswagger custom-server tutorial](https://goswagger.io/tutorial/custom-server.html).

Following the tutorial, I faced GOPATH issue and handled by creating module of the custom-server task through <i>go mod init directory-name</i> and <i>go mod tidy</i> commands.

Finally got the similar results as shown in tutorial;

![](https://i.imgur.com/jaiThVq.png)

![](https://i.imgur.com/Igw6J7g.png)

![](https://i.imgur.com/rIlvVPh.png)









