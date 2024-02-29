{
  new(
    params={},
  ):: {
    params: params,
  },
  params:: {},
  // prevent unknown field exception, require init utils
  Get(paramsField, defaultVal=null):
    local val = if defaultVal != null then
      std.get(self.params, paramsField, defaultVal)
    else
      std.get(self.params, paramsField);
    val,
}
