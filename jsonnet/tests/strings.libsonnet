{
    // new(
    //     params={},
    // )::{
    //     __params: params,
    // },
    // __params:: {},
    ToBool(str): std.asciiLower(str) ==  'true',
    // prevent unknown field exception, require init utils
    // params: function(paramsField, defaultVal = null) {
    //     local v = if defaultVal != null then
    //         std.get(self.__params, paramsField, defaultVal)
    //     else
    //         std.get(self.__params, paramsField),
    //     val: v,
    // }
}