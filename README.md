# Problems

## Cause panic

```
{
  me {
    id
    histories(first: 10) {
      edges {
        node {
          ... on HistoryEntityB {
            id
            textB
          }
        }
      }
      nodes {
        ... on HistoryEntityB {
          id
          textB
        }
      }
    }
  }
}
```

```
2021/08/23 18:14:54 http: panic serving [::1]:54855: runtime error: index out of range [5] with length 5
goroutine 47 [running]:
net/http.(*conn).serve.func1(0xc0012f6140)
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:1804 +0x153
panic(0x170d140, 0xc0001553e0)
	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/panic.go:971 +0x499
github.com/jensneuse/graphql-go-tools/pkg/ast.(*Document).AddSelection(...)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/ast/ast_selection.go:107
github.com/jensneuse/graphql-go-tools/pkg/engine/datasource/graphql_datasource.(*Planner).EnterInlineFragment(0xc0005c6d00, 0x1)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/engine/datasource/graphql_datasource/graphql_datasource.go:254 +0xb8a
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkInlineFragment(0xc001374300, 0x1)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:2004 +0x165
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkSelectionSet(0xc001374300, 0x4)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1719 +0x385
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkField(0xc001374300, 0x7)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1809 +0x645
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkSelectionSet(0xc001374300, 0x5)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1715 +0x407
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkField(0xc001374300, 0x8)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1809 +0x645
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkSelectionSet(0xc001374300, 0x6)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1715 +0x407
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkField(0xc001374300, 0x9)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1809 +0x645
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkSelectionSet(0xc001374300, 0x7)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1715 +0x407
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkOperationDefinition(0xc001374300, 0x0)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1588 +0x46c
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walk(0xc001374300)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1475 +0x60b
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).Walk(0xc001374300, 0xc001313938, 0xc0002db430, 0xc000613f20)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1301 +0xa5
github.com/jensneuse/graphql-go-tools/pkg/engine/plan.(*Planner).Plan(0xc0001ca1c0, 0xc001313938, 0xc0002db430, 0x0, 0x0, 0xc000613f20, 0x0, 0x100000001)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/engine/plan/plan.go:228 +0x97b
github.com/jensneuse/graphql-go-tools/pkg/graphql.(*ExecutionEngineV2).Execute(0xc0005f0160, 0x1b22318, 0xc00140c040, 0xc001313900, 0x1b1d368, 0xc000020f00, 0x0, 0x0, 0x0, 0x0, ...)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/graphql/execution_engine_v2.go:229 +0x24a
gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/gateway/http.(*GraphQLHTTPRequestHandler).handleHTTP(0xc00068ffb0, 0x1b20d40, 0xc00130a540, 0xc0005c6a00)
	/tmp/graphql-go-tools/examples/federation/gateway/http/http.go:46 +0x565
gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/gateway/http.(*GraphQLHTTPRequestHandler).ServeHTTP(0xc00068ffb0, 0x1b20d40, 0xc00130a540, 0xc0005c6a00)
	/tmp/graphql-go-tools/examples/federation/gateway/http/handler.go:49 +0x245
main.(*Gateway).ServeHTTP(0xc000094120, 0x1b20d40, 0xc00130a540, 0xc0005c6a00)
	/tmp/graphql-go-tools/examples/federation/gateway/gateway.go:67 +0x7f
net/http.(*ServeMux).ServeHTTP(0xc000160340, 0x1b20d40, 0xc00130a540, 0xc0005c6a00)
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:2428 +0x1ad
net/http.serverHandler.ServeHTTP(0xc0001ef7a0, 0x1b20d40, 0xc00130a540, 0xc0005c6a00)
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:2867 +0xa3
net/http.(*conn).serve(0xc0012f6140, 0x1b223c0, 0xc0005e4dc0)
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:1932 +0x8cd
created by net/http.(*Server).Serve
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:2993 +0x39b
2021/08/23 18:14:54 http: panic serving [::1]:54864: runtime error: index out of range [5] with length 5
goroutine 49 [running]:
net/http.(*conn).serve.func1(0xc0012f6320)
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:1804 +0x153
panic(0x170d140, 0xc000025368)
	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/panic.go:971 +0x499
github.com/jensneuse/graphql-go-tools/pkg/ast.(*Document).AddSelection(...)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/ast/ast_selection.go:107
github.com/jensneuse/graphql-go-tools/pkg/engine/datasource/graphql_datasource.(*Planner).EnterInlineFragment(0xc0005e0a00, 0x1)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/engine/datasource/graphql_datasource/graphql_datasource.go:254 +0xb8a
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkInlineFragment(0xc001503500, 0x1)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:2004 +0x165
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkSelectionSet(0xc001503500, 0x4)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1719 +0x385
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkField(0xc001503500, 0x7)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1809 +0x645
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkSelectionSet(0xc001503500, 0x5)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1715 +0x407
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkField(0xc001503500, 0x8)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1809 +0x645
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkSelectionSet(0xc001503500, 0x6)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1715 +0x407
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkField(0xc001503500, 0x9)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1809 +0x645
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkSelectionSet(0xc001503500, 0x7)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1715 +0x407
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walkOperationDefinition(0xc001503500, 0x0)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1588 +0x46c
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).walk(0xc001503500)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1475 +0x60b
github.com/jensneuse/graphql-go-tools/pkg/astvisitor.(*Walker).Walk(0xc001503500, 0xc0013d6038, 0xc0002db430, 0xc0012f5200)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/astvisitor/visitor.go:1301 +0xa5
github.com/jensneuse/graphql-go-tools/pkg/engine/plan.(*Planner).Plan(0xc000188380, 0xc0013d6038, 0xc0002db430, 0x0, 0x0, 0xc0012f5200, 0x0, 0x100000001)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/engine/plan/plan.go:228 +0x97b
github.com/jensneuse/graphql-go-tools/pkg/graphql.(*ExecutionEngineV2).Execute(0xc0005f0160, 0x1b22318, 0xc0005e5480, 0xc0013d6000, 0x1b1d368, 0xc0013c2390, 0x0, 0x0, 0x0, 0x0, ...)
	/Users/vvakame/work/gopath/pkg/mod/github.com/jensneuse/graphql-go-tools@v1.20.2/pkg/graphql/execution_engine_v2.go:229 +0x24a
gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/gateway/http.(*GraphQLHTTPRequestHandler).handleHTTP(0xc00068ffb0, 0x1b20d40, 0xc0001efb20, 0xc0005e0700)
	/tmp/graphql-go-tools/examples/federation/gateway/http/http.go:46 +0x565
gist.github.com/vvakame/7a08b7f517b012cca09f82ec9ab4d1fd/gateway/http.(*GraphQLHTTPRequestHandler).ServeHTTP(0xc00068ffb0, 0x1b20d40, 0xc0001efb20, 0xc0005e0700)
	/tmp/graphql-go-tools/examples/federation/gateway/http/handler.go:49 +0x245
main.(*Gateway).ServeHTTP(0xc000094120, 0x1b20d40, 0xc0001efb20, 0xc0005e0700)
	/tmp/graphql-go-tools/examples/federation/gateway/gateway.go:67 +0x7f
net/http.(*ServeMux).ServeHTTP(0xc000160340, 0x1b20d40, 0xc0001efb20, 0xc0005e0700)
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:2428 +0x1ad
net/http.serverHandler.ServeHTTP(0xc0001ef7a0, 0x1b20d40, 0xc0001efb20, 0xc0005e0700)
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:2867 +0xa3
net/http.(*conn).serve(0xc0012f6320, 0x1b223c0, 0xc0005e5380)
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:1932 +0x8cd
created by net/http.(*Server).Serve
	/usr/local/Cellar/go/1.16.6/libexec/src/net/http/server.go:2993 +0x39b
```

## Cause parse error

```
{
  me {
    id
    histories(first: 10) {
      edges {
        node {
          ... on HistoryEntityB {
            id
            textB
          }
        }
      }
    }
  }
}
```

```
{
  "errors": [
    {
      "message": "Expected {, found }",
      "locations": [
        {
          "line": 1,
          "column": 169
        }
      ]
    }
  ],
  "data": {
    "me": {
      "id": "1234",
      "histories": {
        "edges": []
      }
    }
  }
}
```

```
GraphQL Request {
  query($representations: [_Any!]!, $a: Int){_entities(representations: $representations){... on User {histories(first: $a){edges {node} __typename ... on HistoryEntityB }}}}
  resp: {
    "errors": [
      {
        "message": "Expected {, found }",
        "locations": [
          {
            "line": 1,
            "column": 169
          }
        ],
        "extensions": {
          "code": "GRAPHQL_PARSE_FAILED"
        }
      }
    ],
    "data": null
  }
  error: : Expected {, found }
}
```
