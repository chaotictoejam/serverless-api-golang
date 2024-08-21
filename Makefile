# Variables
GOOS ?= linux
GOBIN = ./bin
HANDLERS = createRecipe getRecipeById getRecipes updateRecipe

.PHONY: all test

all: build_createRecipe build_getRecipeById build_getRecipes build_updateRecipe

test: ## Run unit tests
	go test ./...

build_createRecipe:
	go build -o build/createRecipe/bootstrap ./handlers/createRecipe/createRecipe.go

build_getRecipeById:
	go build -o build/getRecipeById/bootstrap ./handlers/getRecipeById/getRecipeById.go

build_getRecipes:
	go build -o build/getRecipes/bootstrap ./handlers/getRecipes/getRecipes.go

build_updateRecipe:
	go build -o build/updateRecipe/bootstrap ./handlers/updateRecipe/updateRecipe.go

zip:
	zip -j build/createRecipe.zip ./build/createRecipe/bootstrap
	zip -j build/getRecipeById.zip ./build/getRecipeById/bootstrap
	zip -j build/getRecipes.zip ./build/getRecipes/bootstrap
	zip -j build/updateRecipe.zip ./build/updateRecipe/bootstrap