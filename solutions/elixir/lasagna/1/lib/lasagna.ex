defmodule Lasagna do
  def expected_minutes_in_oven() do
    40
  end

  def remaining_minutes_in_oven(min) do
    40 - min
  end

  def preparation_time_in_minutes(layers) do 
   layers * 2
  end

  def total_time_in_minutes(layers, time_in_oven) do 
    time_in_oven + preparation_time_in_minutes(layers)
  end

  def alarm() do 
    "Ding!"
  end
end
