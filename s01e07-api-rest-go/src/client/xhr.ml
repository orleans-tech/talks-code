
let (>>=) = Lwt.bind

let url path = 
  match Url.Current.get() with
  | None -> assert false
  | Some u -> (
      match u with
      | Url.Http h -> Url.Http { h with Url.hu_path = path }
      | _ -> assert false
    )

let perform_json ~url ~meth ~f ~json = 
  let ((res : Js.js_string Js.t Js.Opt.t XmlHttpRequest.generic_http_frame Lwt.t), w) = Lwt.task () in
  let req = XmlHttpRequest.create() in
  let () = req##_open (Js.string meth, Js.string (Url.string_of_url url), Js._true) in
  let () = req##setRequestHeader (Js.string "Cache-Control", Js.string "no-cache") in
  let () = req##setRequestHeader (Js.string "Content-type", Js.string "application/json; charset=utf-8") in
  let headers s =
    Js.Opt.case
      (req##getResponseHeader (Js.bytestring s))
      (fun () -> None)
      (fun v -> Some (Js.to_string v))
  in
  let json_response url code headers req = {
    XmlHttpRequest.url = Url.string_of_url url;
    code = code;
    content = File.CoerceTo.json (req##response);
    content_xml = (fun () -> assert false);
    headers = headers
  } in		
  let callback _ = match req##readyState with
    | XmlHttpRequest.DONE -> (
        let response : Js.js_string Js.t Js.Opt.t XmlHttpRequest.generic_http_frame = 
          json_response url (req##status) headers req
        in
        Lwt.wakeup w response
      )
    | _ -> ()
  in
  let () = req##onreadystatechange <- Js.wrap_callback callback in
  let () = req##send (Js.some json) in
  let () = Lwt.on_cancel res (fun () -> req##abort ()) in
  res >>= f

let perform_empty_get ~url ~f = 
  let h = [("Cache-Control", "no-cache")] in	
  XmlHttpRequest.perform ~headers:h ~get_args:[] url >>= f
