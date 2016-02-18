type action =
  | Update_field of Js.js_string Js.t
  | Add
  | Like of Int64.t
  | Refresh