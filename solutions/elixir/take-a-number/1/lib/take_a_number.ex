defmodule TakeANumber do
  def start() do
    spawn(&loop/0)
  end

  def loop(number \\ 0) do
    receive do
      {:report_state, sender_pid} ->
        send(sender_pid, number)
        loop(number)

      {:take_a_number, sender_pid} ->
        send(sender_pid, number + 1)
        loop(number + 1)

      :stop ->
        nil

      _ ->
        loop(number)
    end
  end
end
