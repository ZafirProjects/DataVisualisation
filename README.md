# QuodOrbisChallenge

Solution to Quod Orbis Technical Challenge

download and install go 1.22 or higher

This project uses [templ](templ.guide) as a depenendency because it [allows go data to be passed to scripts](https://templ.guide/syntax-and-usage/script-templates#script-templates).

```bash
git clone https://github.com/ZafirProjects/QuodOrbisChallenge.git
```

cd into the repository

```bash
cd QuodOrbisChallenge
```

To run the project

```bash
go run cmd/main.go
```

The project should run on port 3000 and you can see all routes in the /server/routes.go file.

commands that may help if something doesn't work

```bash
go install github.com/a-h/templ/cmd/templ@latest
go get github.com/a-h/templ
go mod tidy
```

One possible improvement I could have made is my data generation algorithm.
