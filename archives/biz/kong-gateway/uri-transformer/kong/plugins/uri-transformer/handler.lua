-- 继承BasePlugin
local BasePlugin = require "kong.plugins.base_plugin"
local access = require "kong.plugins.uri-transformer.access"
local UriTransformerHandler = BasePlugin:extend()

-- 插件构造函数
function UriTransformerHandler:new()
  UriTransformerHandler.super.new(self, "uri-transformer")
end

function UriTransformerHandler:init_worker()
  UriTransformerHandler.super.init_worker(self)
  -- 在这里实现自定义的逻辑
end

function UriTransformerHandler:certificate(config)
  UriTransformerHandler.super.certificate(self)
  -- 在这里实现自定义的逻辑
end

function UriTransformerHandler:rewrite(config)
  UriTransformerHandler.super.rewrite(self)
  -- 在这里实现自定义的逻辑
end

function UriTransformerHandler:access(config)
  UriTransformerHandler.super.access(self)
  -- 在这里实现自定义的逻辑
end

function UriTransformerHandler:header_filter(config)
  UriTransformerHandler.super.header_filter(self)
  -- 在这里实现自定义的逻辑
end

function UriTransformerHandler:body_filter(config)
  UriTransformerHandler.super.body_filter(self)
  -- 在这里实现自定义的逻辑
end

function UriTransformerHandler:log(config)
  UriTransformerHandler.super.log(self)
  -- 在这里实现自定义的逻辑
end

UriTransformerHandler.PRIORITY = 3131

return UriTransformerHandler
