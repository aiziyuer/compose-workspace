
local _M  = {}

function _M.go()

    require("mobdebug").start("192.168.50.135")

    -- 根据访问的host:port来查询对应的后端服务
    local key = "backend:"..ngx.var.http_host..":"..ngx.var.server_port
    if not key then
        ngx.log(ngx.ERR, "no user-agent found")
        return ngx.exit(400)
    end

    -- 初始化redis连接
    local redis = require "resty.redis"
    local red = redis:new()
    red:set_timeout(1000) -- 1 second

    -- 这里的"11.11.11.3"是在compose文件中redis服务的ip, 是定死的, 避免走DNS的解析降低性能
    local ok, err = red:connect("11.11.11.3", 6379)
    if not ok then
        ngx.log(ngx.ERR, "failed to connect to redis: ", err)
        return ngx.exit(500)
    end

    -- 获取后端的信息(ip:port的形式)
    local value, err = red:get(key)
    if not value then
        ngx.log(ngx.ERR, "failed to get redis key: ", err)
        return ngx.exit(500)
    end

    -- 返回后端信息
    if value == ngx.null then
        ngx.log(ngx.ERR, "no backend found for key ", key)
        return ngx.exit(400)
    end

    -- 正则获取value中的后端和域名[=[${不用转义的正则}]=]
    -- 参考: <<OpenResty最佳实践-正则表达式>> https://moonbingbing.gitbooks.io/openresty-best-practices/lua/re.html
    local regex = [=[^(?<host>[\d.:]+)(?:\s+(?<domain>[a-z.]+))?$]=]
    local m = ngx.re.match(value, regex)
    if m then
        ngx.var.target_host = m["host"]
        -- domain是非必选的, 如果没有选就用原始域名替换
        if m["domain"] ~= nil then
            ngx.var.target_domain = m["domain"]
        end
    else
        ngx.log(ngx.ERR, "unexpected value"..value.." for key", key)
        return ngx.exit(500)
    end
    
end

return _M

