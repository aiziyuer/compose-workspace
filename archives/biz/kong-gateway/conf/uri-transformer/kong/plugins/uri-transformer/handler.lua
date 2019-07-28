-- 继承BasePlugin
local BasePlugin = require "kong.plugins.base_plugin"
local access = require "kong.plugins.uri-transformer.access"

local UriTransformerHandler = BasePlugin:extend()

-- 插件构造函数
function UriTransformerHandler:new()
  UriTransformerHandler.super.new(self, "uri-transformer")
end

function UriTransformerHandler:access(conf)
  UriTransformerHandler.super.access(self)
  access.execute(conf)
end

UriTransformerHandler.PRIORITY = 100

return UriTransformerHandler
