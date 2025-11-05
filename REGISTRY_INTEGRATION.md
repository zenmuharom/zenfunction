# Function Registry Usage for Bianca GraphQL Client

This document shows how to integrate the zenfunction registry with Bianca GraphQL client.

## Overview

The registry provides metadata about all available zenfunction functions, including:

- Function names and descriptions
- Parameter definitions with types
- Return types
- Usage examples
- Categories

## Basic Usage

```go
package main

import (
    "encoding/json"
    "fmt"
    zenfunction "github.com/zenmuharom/zenfunction"
)

func main() {
    // Get all available functions
    functions := zenfunction.GetAvailableFunctions()
    fmt.Printf("Total functions: %d\n", len(functions))

    // Get specific function info
    md5Info := zenfunction.GetFunctionInfo("md5")
    if md5Info != nil {
        fmt.Printf("\nFunction: %s\n", md5Info.Name)
        fmt.Printf("Description: %s\n", md5Info.Description)
        fmt.Printf("Return Type: %s\n", md5Info.ReturnType)
        for _, param := range md5Info.Parameters {
            fmt.Printf("  - %s (%s): %s\n", param.Name, param.Type, param.Description)
        }
    }

    // Get functions by category
    cryptoFunctions := zenfunction.GetFunctionsByCategory("crypto")
    fmt.Printf("\nCrypto functions: %d\n", len(cryptoFunctions))

    // Get all categories
    categories := zenfunction.GetCategories()
    fmt.Printf("Categories: %v\n", categories)

    // JSON serialization for API responses
    jsonData, _ := json.MarshalIndent(functions, "", "  ")
    fmt.Printf("\nJSON Output (first 500 chars):\n%s...\n", string(jsonData[:500]))
}
```

## GraphQL Schema Example

```graphql
type Parameter {
  name: String!
  type: String!
  required: Boolean!
  description: String!
  example: String!
}

type FunctionInfo {
  name: String!
  description: String!
  parameters: [Parameter!]!
  returnType: String!
  examples: [String!]!
  category: String!
}

type Query {
  # Get all available functions
  listFunctions: [FunctionInfo!]!

  # Get specific function by name
  getFunction(name: String!): FunctionInfo

  # Get functions by category
  getFunctionsByCategory(category: String!): [FunctionInfo!]!

  # Get all categories
  getCategories: [String!]!

  # Get all function names
  getFunctionNames: [String!]!
}
```

## GraphQL Resolver Example (Go)

```go
package graphql

import (
    zenfunction "github.com/zenmuharom/zenfunction"
)

type Resolver struct{}

// Query resolvers
func (r *Resolver) Query() QueryResolver {
    return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ListFunctions(ctx context.Context) ([]*FunctionInfo, error) {
    functions := zenfunction.GetAvailableFunctions()
    result := make([]*FunctionInfo, len(functions))
    for i, fn := range functions {
        result[i] = toGraphQLFunctionInfo(fn)
    }
    return result, nil
}

func (r *queryResolver) GetFunction(ctx context.Context, name string) (*FunctionInfo, error) {
    fn := zenfunction.GetFunctionInfo(name)
    if fn == nil {
        return nil, fmt.Errorf("function not found: %s", name)
    }
    return toGraphQLFunctionInfo(*fn), nil
}

func (r *queryResolver) GetFunctionsByCategory(ctx context.Context, category string) ([]*FunctionInfo, error) {
    functions := zenfunction.GetFunctionsByCategory(category)
    result := make([]*FunctionInfo, len(functions))
    for i, fn := range functions {
        result[i] = toGraphQLFunctionInfo(fn)
    }
    return result, nil
}

func (r *queryResolver) GetCategories(ctx context.Context) ([]string, error) {
    return zenfunction.GetCategories(), nil
}

func (r *queryResolver) GetFunctionNames(ctx context.Context) ([]string, error) {
    return zenfunction.GetFunctionNames(), nil
}

// Helper function to convert registry types to GraphQL types
func toGraphQLFunctionInfo(fn zenfunction.FunctionInfo) *FunctionInfo {
    params := make([]*Parameter, len(fn.Parameters))
    for i, p := range fn.Parameters {
        params[i] = &Parameter{
            Name:        p.Name,
            Type:        string(p.Type),
            Required:    p.Required,
            Description: p.Description,
            Example:     p.Example,
        }
    }

    return &FunctionInfo{
        Name:        fn.Name,
        Description: fn.Description,
        Parameters:  params,
        ReturnType:  string(fn.ReturnType),
        Examples:    fn.Examples,
        Category:    fn.Category,
    }
}
```

## GraphQL Query Examples

### Get all functions

```graphql
query {
  listFunctions {
    name
    description
    category
    returnType
    parameters {
      name
      type
      required
      description
    }
    examples
  }
}
```

### Get specific function

```graphql
query {
  getFunction(name: "md5") {
    name
    description
    parameters {
      name
      type
      required
      description
      example
    }
    returnType
    examples
  }
}
```

### Get functions by category

```graphql
query {
  getFunctionsByCategory(category: "crypto") {
    name
    description
    parameters {
      name
      type
      required
    }
  }
}
```

### Get all categories

```graphql
query {
  getCategories
}
```

## REST API Alternative

If you prefer REST over GraphQL:

```go
package api

import (
    "encoding/json"
    "net/http"

    zenfunction "github.com/zenmuharom/zenfunction"
    "github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/api/functions", listFunctionsHandler).Methods("GET")
    r.HandleFunc("/api/functions/{name}", getFunctionHandler).Methods("GET")
    r.HandleFunc("/api/categories", getCategoriesHandler).Methods("GET")
    r.HandleFunc("/api/categories/{category}/functions", getFunctionsByCategoryHandler).Methods("GET")
}

func listFunctionsHandler(w http.ResponseWriter, r *http.Request) {
    functions := zenfunction.GetAvailableFunctions()
    json.NewEncoder(w).Encode(functions)
}

func getFunctionHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]

    fn := zenfunction.GetFunctionInfo(name)
    if fn == nil {
        http.Error(w, "Function not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(fn)
}

func getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
    categories := zenfunction.GetCategories()
    json.NewEncoder(w).Encode(categories)
}

func getFunctionsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    category := vars["category"]

    functions := zenfunction.GetFunctionsByCategory(category)
    json.NewEncoder(w).Encode(functions)
}
```

## Testing the Integration

```bash
# Run the registry tests
go test -v registry_test.go registry.go

# Example output:
# PASS: TestGetAvailableFunctions
# PASS: TestGetFunctionInfo
# PASS: TestGetFunctionsByCategory
# PASS: TestGetCategories
# PASS: TestGetFunctionNames
# PASS: TestJSONSerialization
```

## Next Steps

1. Integrate registry into Bianca's GraphQL server
2. Create GraphQL schema and resolvers
3. Add function execution endpoint (separate from discovery)
4. Add authentication/authorization if needed
5. Document the GraphQL API for end users
