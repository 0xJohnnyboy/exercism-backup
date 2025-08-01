defmodule LanguageList do
  def new(), do: []
  def add(list, language), do: [language | list]
  def remove([_ | t]), do: t
  def first([h | _]), do: h
  def count(list), do: length(list)
  def functional_list?(list), do: "Elixir" in list
end
