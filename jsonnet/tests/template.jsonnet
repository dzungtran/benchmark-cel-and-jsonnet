local __params = import 'data.json';
local strings = import 'strings.libsonnet';
// local Params = import 'params.libsonnet';
// local params = Params.new(__params);

local json_params = std.parseJson('{"json": "this is json document"}');

// functions
local strToBool = function(s) std.asciiLower(s) == 'true';
local params(paramsField, defaultVal=null) =
  if defaultVal != null then
    std.get(__params, paramsField, defaultVal)
  else
    std.get(__params, paramsField)
;

{
  a_value: params('param_A'),
  b_value: 'value of B is: %s & %s' % [params('param_B'), params('param_a', false)],
  b_valus_as_bool: params('param_B'),
  No_value: 'value of B is: %s' % [params('param_No')],
  No_valus_is_not_zero: params('param_No') != 0,
  No_valus_is_odd_number: params('param_No') % 2 == 0,
  bool_string: strings.ToBool(params('bool_string')),
  unknown_field: params('unknown_field', true),
  'json_params.json': json_params.json,
}
