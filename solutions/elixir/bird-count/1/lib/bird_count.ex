defmodule BirdCount do
  def today(list), do: List.first(list)

  def increment_day_count(list) when length(list) == 0, do: [1]
  def increment_day_count([h | t]), do: [h + 1 | t]

  def has_day_without_birds?(list) when length(list) == 0, do: false
  def has_day_without_birds?(list) do
    cond do
      today(list) == 0 -> true
      true -> has_day_without_birds?(tl(list))
    end
  end

  def total(list) when length(list) == 0, do: 0
  def total(list) when length(list) == 1, do: today(list)
  def total([h | t]), do: h + total(t)

  def busy_days(list) when length(list) == 0, do: 0
  def busy_days([h | t]) do
   cond do
     h >= 5 -> 1 + busy_days(t)
     true -> 0 + busy_days(t)
   end 
  end 
end
