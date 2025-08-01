defmodule Username do
  @charmap %{?ä => ~c"ae", ?ö => ~c"oe", ?ü => ~c"ue", ?ß => ~c"ss"}

  def sanitize(username, acc \\ [])
  def sanitize([], acc), do: acc |> List.flatten() |> Enum.reverse()
  def sanitize([h | t], acc) when h in ?a..?z or h == ?_, do: sanitize(t, [h | acc])

  def sanitize([h | t], acc) do
    case Map.get(@charmap, h) do
      nil -> sanitize(t, acc)
      replacement -> sanitize(t, (replacement |> Enum.reverse) ++ acc)
    end
  end

end
