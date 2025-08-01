defmodule LibraryFees do
  def datetime_from_string(string) do
    NaiveDateTime.from_iso8601!(string)
  end

  def before_noon?(datetime) do
    datetime.hour < 12
  end

  defp do_get_return_days(true), do: 28
  defp do_get_return_days(false), do: 29

  def return_date(checkout_datetime) do
    checkout_datetime
    |> before_noon?()
    |> do_get_return_days()
    |> (&Date.shift(checkout_datetime, day: &1)).()
  end

  def days_late(planned_return_date, actual_return_datetime) do
    diff = Date.diff(planned_return_date, actual_return_datetime)
    if diff >= 0, do: 0, else: -diff
  end

  def monday?(datetime) do
    datetime |> Date.day_of_week() == 1
  end

  defp do_get_discounted_rate(true, rate), do: 0.5 * rate
  defp do_get_discounted_rate(false, rate), do: rate

  def calculate_late_fee(checkout, actual_return, rate) do
    nb_days_late =
      checkout
      |> datetime_from_string()
      |> return_date()
      |> days_late(datetime_from_string(actual_return))

    discounted_rate =
      actual_return
      |> datetime_from_string()
      |> monday?()
      |> do_get_discounted_rate(rate)

    (nb_days_late * discounted_rate) |> :math.floor()
  end
end
