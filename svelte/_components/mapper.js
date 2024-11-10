
// arr is something like this:
// [{"name":"id","label":"ID","description":"","type":"int","inputType":"hidden","readOnly":true,"validations":null,"refEndpoint":""},{"name":"uniqPropKey","label":"Prop Key","description":"","type":"string","inputType":"text","readOnly":false,"validations":null,"refEndpoint":""}]
/**
 * @param {Array<any>} arr 
 * @returns {Object.<string, any>}
 */
function fieldsArrToMap(arr) {
  let map = {}
  for(const idx in arr) {
    let el = arr[idx]
    map[el.name] = el;
    el.idx = idx;
  }
  return map;
}

module.exports = {
  fieldsArrToMap: fieldsArrToMap,
};