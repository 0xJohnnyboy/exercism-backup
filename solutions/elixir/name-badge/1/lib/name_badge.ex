defmodule NameBadge do
  def print(id, name, department) do
    "#{if id, do: "[#{id}] - ", else: ""}#{name}#{if department, do: " - #{department |> String.upcase()}", else: " - OWNER"}"
  end
end
