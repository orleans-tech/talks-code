type rs = Types_t.t React.signal
type rf = ?step:React.step -> Types_t.t -> unit
type rp = rs * rf