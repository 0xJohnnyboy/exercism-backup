defmodule FreelancerRates do
  def daily_rate(hourly_rate) do
    8.0 * hourly_rate
  end

  def apply_discount(before_discount, discount) do
    before_discount - (before_discount * discount / 100)
  end

  def monthly_rate(hourly_rate, discount) do
    hourly_rate * 22
    |> daily_rate()
    |> apply_discount(discount)
    |> Float.ceil()
    |> trunc()
  end

  def days_in_budget(budget, hourly_rate, discount) do
    discounted_rate = apply_discount(hourly_rate, discount) |> daily_rate()
    budget / discounted_rate |> Float.floor(1)
  end
end
