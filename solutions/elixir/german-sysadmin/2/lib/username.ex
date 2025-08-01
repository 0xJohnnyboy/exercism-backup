defmodule Username do
  def sanitize([]), do: []
  def sanitize([h | t]) do
    sanitized =
      case h do
        ?ä -> ~c"ae"
        ?ö -> ~c"oe"
        ?ü -> ~c"ue"
        ?ß -> ~c"ss"
        x when x in ?a..?z or h == ?_ -> [x]
        _ -> ~c""
      end

    sanitized ++ sanitize(t)
  end
end
