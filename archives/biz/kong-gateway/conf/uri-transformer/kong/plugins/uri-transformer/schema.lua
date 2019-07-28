return {
  no_consumer = true,
  fields = {
    remotedebug_enable = {type = "boolean", default = false},
    remotedebug_host = {type = "string", default = "192.168.224.1"},
    match_type = {type = "string", default = "regex"},
    uri_template = {type = "string", default = "/"}
  }
}
