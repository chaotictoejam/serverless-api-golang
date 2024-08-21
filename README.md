# aws-toronto-serverless-api-golang
Example code from DEV101 "Build a Serverless API in Go"

## Build

```
go build -o bin/handlers/createRecipe ./handlers/createRecipe/createRecipe.go
go build -o bin/handlers/getRecipeById ./handlers/getRecipeById/getRecipeById.go
go build -o bin/handlers/getRecipes ./handlers/getRecipes/getRecipes.go
go build -o bin/handlers/updateRecipe ./handlers/updateRecipe/updateRecipe.go
```