# Variables
GOOS ?= linux
GOARCH ?= arm64
HANDLERS = createRecipe getRecipeById getRecipes updateRecipe

.PHONY: all test clean

all: build_createRecipe build_getRecipeById build_getRecipes build_updateRecipe

test: ## Run unit tests
	go test ./...

build_createRecipe:
ifeq ($(OS),Windows_NT)
	set GOOS=linux
	set GOARCH=arm64
	go build -o build/createRecipe/bootstrap ./handlers/createRecipe/createRecipe.go
else 
	env GOOS=linux GOARCH=arm64 go build -o build/createRecipe/bootstrap ./handlers/createRecipe/createRecipe.go
endif

build_getRecipeById:
ifeq ($(OS),Windows_NT)
	set GOOS=linux
	set GOARCH=arm64
	go build -o build/getRecipeById/bootstrap ./handlers/getRecipeById/getRecipeById.go
else 
	env GOOS=linux GOARCH=arm64 go build -o build/getRecipeById/bootstrap ./handlers/getRecipeById/getRecipeById.go
endif

build_getRecipes:
ifeq ($(OS),Windows_NT)
	set GOOS=linux
	set GOARCH=arm64
	go build -o build/getRecipes/bootstrap ./handlers/getRecipes/getRecipes.go
else 
	env GOOS=linux GOARCH=arm64 go build -o build/getRecipes/bootstrap ./handlers/getRecipes/getRecipes.go
endif

build_updateRecipe:
ifeq ($(OS),Windows_NT)
	set GOOS=linux
	set GOARCH=arm64
	go build -o build/updateRecipe/bootstrap ./handlers/updateRecipe/updateRecipe.go
else 
	env GOOS=linux GOARCH=arm64 go build -o build/updateRecipe/bootstrap ./handlers/updateRecipe/updateRecipe.go
endif

zip:
	zip -j build/createRecipe.zip ./build/createRecipe/bootstrap
	zip -j build/getRecipeById.zip ./build/getRecipeById/bootstrap
	zip -j build/getRecipes.zip ./build/getRecipes/bootstrap
	zip -j build/updateRecipe.zip ./build/updateRecipe/bootstrap

clean:
ifeq ($(OS),Windows_NT)
	rmdir /S /Q build
else
	rm -rf build
endif