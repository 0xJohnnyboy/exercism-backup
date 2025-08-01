defmodule LibraryFees do
  def datetime_from_string(string) do
    {:ok, dt} = NaiveDateTime.from_iso8601(string)
    dt
  end

  def before_noon?(datetime) do
    datetime
    |> NaiveDateTime.to_time()
    |> Time.before?(~T[12:00:00])
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
    cond do
      Date.diff(planned_return_date, actual_return_datetime) >= 0 -> 0
      true -> -Date.diff(planned_return_date, actual_return_datetime)
    end
  end

  def monday?(datetime) do
    datetime |> Date.day_of_week() == 1
  end

  defp do_get_discounted_rate(true, rate), do: 0.5 * rate
  defp do_get_discounted_rate(false, rate), do: rate

  def calculate_late_fee(checkout, actual_return, rate) do
    dt_checkout = datetime_from_string(checkout)
    dt_return = datetime_from_string(actual_return)
    dt_planned_return_date = return_date(dt_checkout)

    is_monday = monday?(dt_return)
    nb_days_late = days_late(dt_planned_return_date, dt_return)
    dis_rate = do_get_discounted_rate(is_monday, rate)
        
    nb_days_late * dis_rate |> :math.floor()

  end
end
