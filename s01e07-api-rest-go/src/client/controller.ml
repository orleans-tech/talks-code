open Types_t
open Types
open Action
open Model

let rec update a ((r, f) : rp) =
  let m = React.S.value r in
  let m =
    match a with
    | Refresh ->
      { m with field = "" }
    | Update_field field ->
      { m with field = Js.to_string field }
    | Like _
    | Add -> m
  in
  let _ =
    match a with
    | Refresh ->
      let f req =
        let code = req.XmlHttpRequest.code in 
        if (code = 200) then (
          let content = req.XmlHttpRequest.content in
          let l = Types_j.topic_list_of_string content in
          let m = Model.set_topics l m in
          f m
        ) ;
        Lwt.return ()
      in
      Xhr.perform_empty_get ~url:(Xhr.url ["api"; "topics"]) ~f
    | Add ->
      let f req =
        let code = req.XmlHttpRequest.code in 
        let content = req.XmlHttpRequest.content in
        Js.Opt.case content 
          (fun () -> ())
          (fun v -> 
             if (code = 201) then (
               update Refresh (r, f) ;
             )
          ) ;
        Lwt.return ()
      in
      let json = Types_j.string_of_topic(Model.new_topic m.field) in
      Xhr.perform_json ~url:(Xhr.url ["api"; "topics"]) ~meth:"POST" ~f ~json:(Js.string json)
    | Like id ->
      let f req =
        let code = req.XmlHttpRequest.code in 
        let content = req.XmlHttpRequest.content in
        Js.Opt.case content 
          (fun () -> ())
          (fun v -> 
             if (code = 200) then (
               update Refresh (r, f) ;
             )
          ) ;
        Lwt.return ()
      in
      Xhr.perform_json ~url:(Xhr.url ["api"; "topics"; "like"; Int64.to_string id]) ~meth:"PUT" ~f ~json:(Js.string "")
    | Update_field _ -> Lwt.return ()
  in
  f m