return {
  no_consumer = true,
  fields = {
    remotedebug_enable = {type = "boolean", default = false},
    remotedebug_host = {type = "string", default = "172.17.0.1"},
    preserve_schema = {type = "boolean", default = false },
    match_type = {type = "string", default = "regex"},
    uri_template = {type = "string", default = "/"}
  }
}
