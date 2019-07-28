package = "kong-plugin-uri-transformer"
version = "0.0.1-1"
source = {
   url = "https://github.com/aiziyuer/compose-workspace/tree/master/archives/biz/kong-gateway",
}
description = {
   homepage = "http://aiziyuer.github.io/2019/07/20/kong-study.html",
   license = "Apache-2.0"
}
dependencies = {}
build = {
  type = "builtin",
  modules = {
    ["kong.plugins.uri-transformer.handler"] = "kong/plugins/uri-transformer/handler.lua",
    ["kong.plugins.uri-transformer.schema"] = "kong/plugins/uri-transformer/schema.lua",
    ["kong.plugins.uri-transformer.access"] = "kong/plugins/uri-transformer/access.lua",
  }
}