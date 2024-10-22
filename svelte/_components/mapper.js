
// arr is something like this:
// [{"name":"id","label":"ID","description":"","type":"int","inputType":"hidden","readOnly":true,"validations":null,"refEndpoint":""},{"name":"uniqPropKey","label":"Prop Key","description":"","type":"string","inputType":"text","readOnly":false,"validations":null,"refEndpoint":""}]
function fieldsArrToMap(arr) {
  let map = {}
  for(const idx in arr) {
    let el = arr[idx]
    console.log(el)
    map[el.name] = el;
    el.idx = idx;
  }
  return map;
}

module.exports = {
  fieldsArrToMap: fieldsArrToMap,
};