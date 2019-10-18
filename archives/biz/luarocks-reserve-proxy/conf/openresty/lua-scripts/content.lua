package.path = package.path .. ";/opt/zbstudio/lualibs/?/?.lua;/opt/zbstudio/lualibs/?.lua"
package.cpath = package.cpath .. ";/opt/zbstudio/bin/linux/x64/clibs/?/?.so;/opt/zbstudio/bin/linux/x64/clibs/?.so"

local debug = require("mobdebug")
debug.start("192.168.50.135")
debug.coro()

tmp_str = "Debug"

print("Lua Lapis")

ngx.say("Openresty")

require('mobdebug').done()